package cmd

import (
	"net/http"
	"os"

	s32 "github.com/nadilas/moar/internal/storage/s3"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/gorilla/handlers"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
	"go.elastic.co/apm/module/apmhttp"

	"github.com/nadilas/moar/internal/registry"
	"github.com/nadilas/moar/moarpb"
	"github.com/nadilas/moar/rpc"
)

var (
	moduleStorageType string
	host              string
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
			moduleStorage = s32.New()
		default:
			logrus.Fatalf("invalid module storage type: '%s'", moduleStorageType)
		}
		reverseProxy := viper.GetString("proxy")
		registry := registry.New(moduleStorage, reverseProxy)
		server := rpc.NewServer(registry)

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
	upCmd.Flags().StringVar(&host, "host", ":8000", "The address to bind the server to")
	rootCmd.AddCommand(upCmd)
}
