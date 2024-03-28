// Copyright © 2021 dotindustries
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/leebenson/conform"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// VersionInfo contains the version and commit hash
type VersionInfo struct {
	Version string
	Commit  string
}

var (
	// Version set in compile time through ldflags
	Version VersionInfo
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "moarctl",
	Short: "UMD module registry application",
	Long: `The registry manages multiple versions of umd modules. For example:

The registry can upload new VueJS or ReactJS modules, resolve the storage urls of version ranges
or specific versions.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		readGlobalConfig()

		if GlobalConfig.Debug {
			logrus.SetLevel(logrus.DebugLevel)
		}

		logrus.Debugf("moarctl version: %s (%s)", Version.Version, Version.Commit)

		return nil
	},
}

// init is called before main
func init() {
	// A custom sanitizer to redact sensitive data by defining a struct tag= named "redact".
	conform.AddSanitizer("redact", func(_ string) string { return "*****" })

	// Initialize the config and panic on failure
	if err := configInit(); err != nil {
		panic(err.Error())
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}
