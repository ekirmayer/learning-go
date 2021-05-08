/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var name string
var greeting string
var preview bool
var prompt bool
var debug bool = false

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "motd",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		if prompt == false && (name == "" || greeting == "") {
			cmd.Usage()
			os.Exit(1)
		}

		if prompt {
			name, greeting = renderPrompt()
		}

		if debug {
			fmt.Println("Name", name)
			fmt.Println("greeting", greeting)
			fmt.Println("prompt", prompt)
			os.Exit(0)
		}

		m := buildMessage(name, greeting)

		if preview {
			fmt.Println(m)
		} else {
			f, err := os.OpenFile("./output/openfile.txt", os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error: unable to open the file for writing.", err)
				os.Exit(1)
			}
			defer f.Close()

			_, err = f.Write([]byte(m))
			if err != nil {
				fmt.Println("Error: unable to write to the file")
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

// This can be in any file and will be executed on file load
func init() {
	// cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.motd.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringVarP(&name, "name", "n", "", "name to use inthe message")
	rootCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "phrase to use in the greeting message")
	rootCmd.Flags().BoolVarP(&prompt, "prompt", "p", false, "use the prompt for input")
	rootCmd.Flags().BoolVarP(&preview, "preview", "v", false, "Show preview of the message without writing to file")
	rootCmd.Flags().BoolVarP(&debug, "debug", "d", false, "Show debug")
}

func renderPrompt() (name, greeting string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your greeting? ")
	greeting, _ = reader.ReadString('\n')
	greeting = strings.TrimSpace(greeting)

	fmt.Print("Name please? ")
	name, _ = reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return
}

func buildMessage(name string, message string) string {
	return fmt.Sprintf("%s, %s", message, name)
}
