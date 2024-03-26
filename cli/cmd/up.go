package cmd

import (
	"github.com/dotindustries/moar/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/dotindustries/moar/internal/registry"
	"github.com/dotindustries/moar/internal/storage/s3"
	"github.com/dotindustries/moar/moarpb"
	"github.com/dotindustries/moar/rpc"
	"github.com/gorilla/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
)

var (
	moduleStorageType       string
	storageAddress          string
	host                    string
	versionOverwriteEnabled bool
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "Starts the registry service",
	Run: func(cmd *cobra.Command, args []string) {
		var moduleStorage registry.Storage
		switch moduleStorageType {
		// case "etcd":
		// moduleStorage = internal.NewDatabase()
		case "s3":
			moduleStorage = s3.New(storageAddress)
		default:
			logrus.Fatalf("invalid module storage type: '%s'", moduleStorageType)
		}
		registry := registry.New(moduleStorage)

		if reverseProxyAddr == "" {
			reverseProxyAddr = os.Getenv("S3_PROXY_URL")
			if reverseProxyAddr == "" {
				reverseProxyAddr = defaultReverseProxyAddr
			}
		}
		logrus.Infof("Using reverse proxy for content with address: %s", reverseProxyAddr)
		server := rpc.NewServer(registry, reverseProxyAddr, rpc.Opts{
			VersionOverwriteEnabled: versionOverwriteEnabled,
		})

		// Echo instance
		app, err := apm()
		if err != nil {
			panic(err)
		}

		twirpHandler := moarpb.NewModuleRegistryServer(server,
			twirp.WithServerPathPrefix(""),
		)
		loggingHandler := handlers.CombinedLoggingHandler(os.Stdout, twirpHandler)

		e := echo.New()

		s := NewStats()
		e.Use(s.Process)
		e.GET("/stats", s.Handle) // Endpoint to get stats

		e.Use(middleware.RequestID())
		e.Use(nrecho.Middleware(app))
		e.GET("/", func(c echo.Context) error {
			return c.JSON(http.StatusOK, "I'm up")
		})
		e.Any("*", echo.WrapHandler(loggingHandler), middleware.KeyAuthWithConfig(middleware.KeyAuthConfig{
			// Allow for both Authorization and X-Api-Key header
			KeyLookup: "header:" + echo.HeaderAuthorization + ",header:X-Api-Key",
			Validator: auth.KeyValidator,
		}))
		logrus.Infof("Registry listening on http://%s/", host)
		if err := http.ListenAndServe(host, e); err != nil {
			logrus.Fatal(err)
		}

		server.Shutdown()
	},
}

func apm() (*newrelic.Application, error) {
	return newrelic.NewApplication(
		newrelic.ConfigAppLogForwardingEnabled(true),
		newrelic.ConfigFromEnvironment(),
	)
}

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: map[string]int{},
	}
}

// Process is the middleware function.
func (s *Stats) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

// Handle is the endpoint to get stats.
func (s *Stats) Handle(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

func init() {
	upCmd.Flags().StringVar(&moduleStorageType, "storage_type", "s3", "Defines what storage type to use. Possible values: s3")
	upCmd.Flags().StringVar(&storageAddress, "storage_addr", "", "The address to reach the storage")
	upCmd.Flags().StringVar(&host, "host", ":8000", "The address to bind the server to")
	upCmd.Flags().BoolVar(&versionOverwriteEnabled, "overwrite", false, "Toggles whether version overwrite is enabled")
	rootCmd.AddCommand(upCmd)
}
