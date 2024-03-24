package cmd

import (
	"context"
	"github.com/dotindustries/moar/auth"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"net/http"
	"os"

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
		e.Use(middleware.KeyAuth(auth.KeyValidator))
		e.Use(nrecho.Middleware(app))
		e.Any("*", echo.WrapHandler(loggingHandler))
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

func newAPMInterceptor() twirp.Interceptor {
	return func(next twirp.Method) twirp.Method {
		return func(ctx context.Context, req interface{}) (interface{}, error) {

			return next(ctx, req)
		}
	}
}

func init() {
	upCmd.Flags().StringVar(&moduleStorageType, "storage_type", "s3", "Defines what storage type to use. Possible values: s3")
	upCmd.Flags().StringVar(&storageAddress, "storage_addr", "", "The address to reach the storage")
	upCmd.Flags().StringVar(&host, "host", ":8000", "The address to bind the server to")
	upCmd.Flags().BoolVar(&versionOverwriteEnabled, "overwrite", false, "Toggles whether version overwrite is enabled")
	rootCmd.AddCommand(upCmd)
}
