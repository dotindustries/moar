package cmd

import (
	"connectrpc.com/connect"
	"context"
	"errors"
	"fmt"
	"github.com/sanity-io/litter"
	"os"

	moarpb "github.com/dotindustries/moar/moarpb/v1"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/spf13/cobra"
)

var moduleCmd = &cobra.Command{
	Use:     "module",
	Short:   "Module management commands",
	Aliases: []string{"m"},
}

var getAll bool
var getModuleCmd = &cobra.Command{
	Use:     "get module_name",
	Short:   "Get module details",
	Aliases: []string{"g", "read", "r"},
	Run: func(cmd *cobra.Command, args []string) {
		client := protobufClient()

		request := &moarpb.GetModuleRequest{}

		if !getAll {
			if len(args) < 1 {
				fmt.Println("ERROR: module name not provided")
				_ = cmd.Help()
				os.Exit(1)
			}
			request.ModuleName = args[0]
		}
		rsp, err := client.GetModule(context.Background(), connect.NewRequest(request))
		if err != nil {
			fmt.Println(connect.CodeOf(err))
			if connectErr := new(connect.Error); errors.As(err, &connectErr) {
				fmt.Println(connectErr.Message())
				litter.Dump(connectErr.Details())
			}
			os.Exit(1)
		}
		response := rsp.Msg
		t := table.NewWriter()
		t.SetOutputMirror(os.Stdout)
		t.AppendHeader(table.Row{"Module", "Author", "Language", "Version"})
		rowConfigAutoMerge := table.RowConfig{AutoMerge: true}
		for _, module := range response.Module {
			for _, v := range module.Versions {
				t.AppendRow(table.Row{module.Name, module.Author, module.Language, v.Name}, rowConfigAutoMerge)
				// t.AppendSeparator()
			}
		}
		t.Render()
	},
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
		_, err = client.CreateModule(context.Background(), connect.NewRequest(&moarpb.CreateModuleRequest{
			ModuleName: moduleName,
			Author:     author,
			Language:   language,
		}))
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

	getModuleCmd.Flags().BoolVarP(&getAll, "all", "a", false, "Gets all modules within the system ignoring provided module name arguments")
	moduleCmd.AddCommand(getModuleCmd)
	rootCmd.AddCommand(moduleCmd)
}
