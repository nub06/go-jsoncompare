/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/nub06/go-jsoncompare/conf"
	"github.com/nub06/go-jsoncompare/service"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "go-jsoncompare",
	Short: "Compares two files and displays the differences",
	Long: `It is used to compare two files and display the differences between them. The command takes two file paths as input arguments and usage of using input:
	e.g. "go-jsoncompare file1 file2"`,

	Run: func(cmd *cobra.Command, args []string) {

		compareFiles(args)

	},
}

func compareFiles(args []string) {

	if len(args) != 2 {
		log.Fatal("You must enter exactly two inputs to run the program \n Number of inputs you are trying to enter:", len(args))
	}

	for i, s := range args {
		if !strings.Contains(s, ".json") {
			var sb strings.Builder
			sb.WriteString(args[i])
			sb.WriteString(".json")
			args[i] = sb.String()
		}

		args[i] = filepath.FromSlash(args[i])

	}

	conf.FirstInput, conf.SecondInput = args[0], args[1]

	service.Run()

}
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
