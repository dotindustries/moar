package cmd

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	"github.com/dotindustries/moar/moarpb/v1/v1connect"
	"io/ioutil"
	"mime"
	"os"
	"path"
	"strings"

	"github.com/Masterminds/semver"
	"github.com/dotindustries/moar/client"
	moarpb "github.com/dotindustries/moar/moarpb/v1"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "Version management commands",
	Aliases: []string{"v"},
}

var (
	module    string
	version   string
	uploadCmd = &cobra.Command{
		Use:     "upload",
		Short:   "Uploads a new module version to the registry",
		Aliases: []string{"up"},
		Args: func(cmd *cobra.Command, args []string) error {
			min := 1
			if len(args) < min {
				return fmt.Errorf("at least %d arg(s) are required, received %d", min, len(args))
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

			providedPath := args[0]
			fi, err := os.Stat(providedPath)
			if err != nil {
				logrus.Fatal(err)
			}
			// grab files
			var files []*moarpb.File
			if fi.IsDir() {
				dir, err := os.ReadDir(providedPath)
				if err != nil {
					logrus.Fatal(err)
				}
				for _, entry := range dir {
					if entry.IsDir() {
						continue
					}
					efi, _ := entry.Info()
					f, err := parseFile(path.Join(providedPath, efi.Name()))
					if err != nil {
						logrus.Fatal(err)
					}
					files = append(files, f)
				}
			}
			// TODO validate that at least 1 js file is available
			_, err = client.UploadVersion(context.Background(), connect.NewRequest(&moarpb.UploadVersionRequest{
				ModuleName: module,
				Version:    version,
				Files:      files,
			}))

			if err != nil {
				fmt.Println("ERROR: ", err)
				os.Exit(1)
			}
			fmt.Printf("Successfully uploaded new version: %s\n", ver.String())
		},
	}
)

func parseFile(filepath string) (*moarpb.File, error) {
	fileBytes := MustReadFileBytes(filepath)
	fileName := path.Base(filepath)

	var mtype string
	if mtype = mime.TypeByExtension(path.Ext(filepath)); mtype == "" {
		mtype = fmt.Sprintf("application/%s", strings.TrimPrefix(path.Ext(filepath), "."))
	}
	return &moarpb.File{
		Name:     fileName,
		Data:     fileBytes,
		MimeType: mtype,
	}, nil
}

func MustReadFileBytes(path string) []byte {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return bytes
}

func protobufClient() v1connect.ModuleRegistryServiceClient {
	return client.New(client.Config{Url: backendAddr})
}

func init() {
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
