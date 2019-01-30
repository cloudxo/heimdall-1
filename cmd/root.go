// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
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
	"git01.bravofly.com/n7/heimdall/pkg/logging"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "heimdall",
	Short: "Heimdall, all-seeing and all-hearing",
	Long:  `Metrics importer for Cloudflare API`,

	Run: func(cmd *cobra.Command, args []string) {
		Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error("%v", err)
		//logger.Printf("%v",err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Error("%v", err)
			//logger.Printf("%v",err)
			os.Exit(1)
		}

		// Search config in home directory with name ".heimdall" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".heimdall")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Error("Using config file: %s", viper.ConfigFileUsed())
		//logger.Printf("Using config file: %s", viper.ConfigFileUsed())
	}
}
