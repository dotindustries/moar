package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/moarpb"
	"github.com/spf13/cobra"
	"github.com/twitchtv/twirp"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Version management commands",
	Aliases: []string{"v"},
}

var (
	backendAddr string
	module      string
	version     string
	uploadCmd   = &cobra.Command{
		Use:   "upload",
		Short: "Uploads a new module version to the registry",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			ver, err := semver.NewVersion(version)
			if err != nil {
				fmt.Println(err, ":", version)
				os.Exit(5)
			}

			client := protobufClient()

			filePath := args[0]
			bytes, err := ioutil.ReadFile(filePath)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}

			_, err = client.UploadVersion(context.Background(), &moarpb.UploadVersionRequest{
				ModuleName: module,
				Version:    version,
				FileData:   bytes,
			})

			if err != nil {
				fmt.Println("ERROR: ", err)
				os.Exit(1)
			}
			fmt.Printf("Successfully uploaded new version: %s\n", ver.String())
		},
	}
)

func protobufClient() moarpb.ModuleRegistry {
	return moarpb.NewModuleRegistryProtobufClient(
		backendAddr,
		http.DefaultClient,
		twirp.WithClientPathPrefix(""),
	)
}

func init() {
	uploadCmd.Flags().StringVarP(&backendAddr, "addr", "a", "http://localhost:8000", "The backend service address")
	uploadCmd.Flags().StringVarP(&module, "module", "m", "", "The target module name")
	uploadCmd.Flags().StringVarP(&version, "version", "v", "", "The target version to upload")
	err := uploadCmd.MarkFlagRequired("module")
	if err != nil {
		panic(err)
	}
	err = uploadCmd.MarkFlagRequired("version")
	if err != nil {
		panic(err)
	}
	versionCmd.AddCommand(uploadCmd)
	rootCmd.AddCommand(versionCmd)
}
