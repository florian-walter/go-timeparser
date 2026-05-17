/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: parseRun,
}

var unit string

func parseRun(cmd *cobra.Command, args []string) {
	fmt.Println("parse called")

	for _, x := range args {
		num, err := strconv.Atoi(x)
		if (err != nil) {
			fmt.Println("failed")
			break
		}
		fmt.Printf("%v %s = ", num, unit)

		strings := []string{}
		switch unit {
		case "ms":
			if (num <= 0) {
				break
			}
			strings = append([]string{strconv.Itoa(num % 1000) + "ms "}, strings...)
			num = num / 1000
			fallthrough
		case "s":
			if (num <= 0) {
				break
			}
			strings = append([]string{strconv.Itoa(num % 60) + "s "}, strings...)
			num = num / 60
			fallthrough
		case "m":
			if (num <= 0) {
				break
			}
			strings = append([]string{strconv.Itoa(num % 60) + "m "}, strings...)
			num = num / 60
			fallthrough
		case "h":
			if (num <= 0) {
				break
			}
			strings = append([]string{strconv.Itoa(num % 24) + "h "}, strings...)
			num = num / 24
			fallthrough
		case "d":
			if (num <= 0) {
				break
			}
			strings = append([]string{strconv.Itoa(num) + "d "}, strings...)
		}

		for i := range strings {
			fmt.Print(strings[i])
		}
		fmt.Print("\n")
	}
}

func init() {
	rootCmd.AddCommand(parseCmd)

	parseCmd.Flags().StringVarP(&unit, "unit", "u", "ms", "Specifies the unit for the time value")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// parseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// parseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
