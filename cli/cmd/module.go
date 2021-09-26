package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nadilas/moar/moarpb"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:     "module",
	Short:   "Module management commands",
	Aliases: []string{"m"},
}

var author, language string

var newModuleCmd = &cobra.Command{
	Use:     "new",
	Short:   "Creates a new module",
	Aliases: []string{"n", "c", "create"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := validateModuleLanguage()
		client := protobufClient()

		moduleName := args[0]
		_, err = client.CreateModule(context.Background(), &moarpb.CreateModuleRequest{
			ModuleName: moduleName,
			Author:     author,
			Language:   language,
		})
		if err != nil {
			fmt.Println("ERROR: ", err)
			os.Exit(1)
		}
		fmt.Printf("Successfully created module: %s\n", moduleName)
	},
}

func validateModuleLanguage() error {
	if language != "vue" && language != "react" {
		return fmt.Errorf("invalid module language: %s", language)
	}
	return nil
}

func init() {
	newModuleCmd.Flags().StringVarP(&author, "author", "a", "", "The author for the module")
	err := newModuleCmd.MarkFlagRequired("author")
	if err != nil {
		panic(err)
	}
	newModuleCmd.Flags().StringVarP(&language, "lang", "l", "", "The module language. Possible values are: vue|react")
	err = newModuleCmd.MarkFlagRequired("lang")
	if err != nil {
		panic(err)
	}
	moduleCmd.AddCommand(newModuleCmd)
	rootCmd.AddCommand(moduleCmd)
}
