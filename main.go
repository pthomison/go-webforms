package main

import (
	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
	// "os/exec"
	// "io/fs"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
)

var (
	// message string
	// name    string

	rootCmd = &cobra.Command{
		Use:   "go-webforms",
		Short: "go-webforms",
		Run:   webforms,
	}
)

func init() {
}

func main() {

	// rootCmd.PersistentFlags().StringVarP(&message, "message", "m", "hello world", "message the program will output")
	// rootCmd.PersistentFlags().StringVarP(&name, "name", "n", "patrick", "name the program will output to")

	err := rootCmd.Execute()

	utils.Check(err)
}

func webforms(cmd *cobra.Command, args []string) {
	a := &App{}

	utils.Check(a.init())
	utils.Check(a.runServer())
}
