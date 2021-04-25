package main

import (
	utils "github.com/pthomison/golang-utils"
	"github.com/spf13/cobra"
)

const (
	HOST = "0.0.0.0"
	PORT = "8080"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-webforms",
		Short: "go-webforms",
		Run:   webforms,
	}
)

func main() {
	utils.Check(rootCmd.Execute())
}

func webforms(cmd *cobra.Command, args []string) {
	a := &App{}
	utils.Check(a.runServer())
}
