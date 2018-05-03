// Copyright © 2018 Anthony J Mirabella <a9@aneurysm9.com>
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
	"math/rand"
	"os"
	"time"

	"github.com/aneurysm9/evoword/model"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var target string
var popsize int
var maxgen int
var mutrate float32

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "evoword",
	Short: "Attempt to evolve a random string into a target string",
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		c := model.Config{
			Target:       []byte(target),
			Population:   popsize,
			MaxGens:      maxgen,
			MutationRate: mutrate,
		}
		m := model.New(c)
		m.Evolve()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.evoword.yaml)")
	rootCmd.PersistentFlags().StringVarP(&target, "target", "t", "Hello, World!", "target word")
	rootCmd.PersistentFlags().IntVarP(&popsize, "population", "p", 1e4, "population size")
	rootCmd.PersistentFlags().IntVarP(&maxgen, "maxgen", "g", 1e3, "maximum generations")
	rootCmd.PersistentFlags().Float32VarP(&mutrate, "mutation", "m", .001, "mutation rate")

	rand.Seed(time.Now().Unix())
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
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".evoword" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".evoword")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
