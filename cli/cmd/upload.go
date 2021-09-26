package cmd

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/nadilas/moar/client"
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
		Use:     "upload",
		Short:   "Uploads a new module version to the registry",
		Aliases: []string{"up"},
		Args: func(cmd *cobra.Command, args []string) error {
			min := 1
			max := 2
			if len(args) < min || len(args) > max {
				return fmt.Errorf("accepts between %d and %d arg(s), received %d", min, max, len(args))
			}
			// first arg must be .js
			if !strings.HasSuffix(args[0], ".js") {
				return fmt.Errorf("first argument must be a javascript file with .js extension. Got: %s", args[0])
			}
			if len(args) < 2 {
				return nil
			}
			// second arg if provided must be .css
			if !strings.HasSuffix(args[1], ".css") {
				return fmt.Errorf("second argument must be a stylesheet with .css extension. Got: %s", args[1])
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			ver, err := semver.NewVersion(version)
			if err != nil {
				fmt.Println(err, ":", version)
				os.Exit(5)
			}

			client := protobufClient()

			scriptPath := args[0]
			bytes := MustReadFileBytes(scriptPath)
			var styleBytes []byte
			if len(args) == 2 {
				stylePath := args[1]
				styleBytes = MustReadFileBytes(stylePath)
			}
			_, err = client.UploadVersion(context.Background(), &moarpb.UploadVersionRequest{
				ModuleName: module,
				Version:    version,
				FileData:   bytes,
				StyleData:  styleBytes,
			})

			if err != nil {
				fmt.Println("ERROR: ", err)
				os.Exit(1)
			}
			fmt.Printf("Successfully uploaded new version: %s\n", ver.String())
		},
	}
)

func MustReadFileBytes(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bytes
}

func protobufClient() moarpb.ModuleRegistry {
	return client.New(client.Config{Url: backendAddr}, twirp.WithClientPathPrefix(""))
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
