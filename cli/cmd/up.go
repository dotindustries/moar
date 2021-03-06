package cmd

import (
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
	"go.elastic.co/apm/module/apmhttp"
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

		twirpHandler := moarpb.NewModuleRegistryServer(server, twirp.WithServerPathPrefix(""))
		tracedHandler := apmhttp.Wrap(twirpHandler)
		loggingHandler := handlers.CombinedLoggingHandler(os.Stdout, tracedHandler)
		logrus.Infof("Registry listening on http://%s/", host)
		if err := http.ListenAndServe(host, loggingHandler); err != nil {
			logrus.Fatal(err)
		}

		server.Shutdown()
	},
}

func init() {
	upCmd.Flags().StringVar(&moduleStorageType, "storage_type", "s3", "Defines what storage type to use. Possible values: s3")
	upCmd.Flags().StringVar(&storageAddress, "storage_addr", "", "The address to reach the storage")
	upCmd.Flags().StringVar(&host, "host", ":8000", "The address to bind the server to")
	upCmd.Flags().BoolVar(&versionOverwriteEnabled, "overwrite", false, "Toggles whether version overwrite is enabled")
	rootCmd.AddCommand(upCmd)
}
