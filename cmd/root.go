// Copyright 2018 Axel Springer SE
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
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ssm"
	config "github.com/axelspringer/templeton/cfg"
	"github.com/axelspringer/templeton/env"
	"github.com/axelspringer/templeton/store"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	cfg     *config.Config
)

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "templeton",
	Short: "Executes a subprocess with environment variables from SSM",
	RunE:  runE,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// new AWS Session
	session := session.Must(session.NewSession())

	// Init a new Config
	cfg = config.Must(session)

	// Init Cobra config
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.templeton.yaml)")

	// Set the verbosity
	RootCmd.PersistentFlags().BoolVarP(&cfg.Verbose, "verbose", "v", cfg.Verbose, "enable verbose logging")

	// SSMPAth
	RootCmd.Flags().StringVarP(&cfg.SSMPath, "path", "p", "", "path in the SSM")

	// Recursive lookup
	RootCmd.Flags().BoolVarP(&cfg.Recursive, "recursive", "r", cfg.Recursive, "recursive lookup")

	// With decryption
	RootCmd.Flags().BoolVarP(&cfg.WithDecryption, "decrypt", "d", cfg.WithDecryption, "disable decryption")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	// read in files
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".templeton" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".templeton")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

// runE is running the root command of Templeton
func runE(cmd *cobra.Command, args []string) error {
	var err error

	svc := ssm.New(cfg.Session)

	_, err = store.New(svc, config.String(cfg.SSMPath), config.Bool(cfg.Recursive), config.Bool(cfg.WithDecryption))
	_ = env.New()

	return err
}