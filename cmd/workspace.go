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
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/spf13/cobra"
)

// workspaceCmd represents the workspace command
var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "Generates a Go workspace",
	Long:  `Usage: 'go-start workspace <workspace-folder-name>'`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide a <workspace-folder-name> after calling the workspace command.")
			return
		}

		dir := args[0]
		if err := os.Mkdir(dir, 0700); err != nil {
			log.Fatal("ERROR:", err)
			return
		}

		pwd, err := os.Getwd()
		if err != nil {
			log.Fatal("ERROR:", err)
			return
		}

		if err := os.Chdir(path.Join(pwd, dir)); err != nil {
			log.Fatal("ERROR:", err)
			return
		}

		gosrc := os.Getenv("GOSTARTDIR")
		workdir := path.Join(gosrc, dir)

		if err := goModInit(workdir); err != nil {
			log.Fatal("ERROR:", err)
			return
		}

		if err := gitInit(); err != nil {
			log.Fatal("ERROR:", err)
			return
		}

		if err := mainGoInit(); err != nil {
			log.Fatal("ERROR:", err)
			return
		}

		fmt.Println("Go workspace created successfully at:", workdir)
	},
}

func goModInit(dir string) error {
	cmd := exec.Command("go", "mod", "init", dir)
	return cmd.Run()
}

func gitInit() error {
	return exec.Command("git", "init", ".").Run()
}

func mainGoInit() error {
	data := `
package main

import "fmt"

func main() {
	fmt.Println("these pretzels are making me thirsty")
}
	`
	if err := os.WriteFile("main.go", []byte(data), 0700); err != nil {
		return err
	}

	return exec.Command("gofmt", "main.go").Run()
}

func init() {
	rootCmd.AddCommand(workspaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// workspaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// workspaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
