/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type Unit int

type UnitObject struct {
	name 		string
	maxTime int
}

// parseCmd represents the parse command
var parseCmd = &cobra.Command{
	Use:   "parse",
	Short: "Converts a time value provided to a time representation",
	Long: `Parse converts the time value provided to a readable time representation`,
	Run: parseRun,
}

var unitFlag string
var unit Unit

const (
	Microseconds Unit = iota
	Milliseconds
	Seconds
	Minutes
	Hours
	Days
	MAX_TIME_UNIT
)

var unitObjects = map[Unit]UnitObject{
	Microseconds: {name: "us", maxTime: 1000},
	Milliseconds: {name: "ms", maxTime: 1000},
	Seconds: 			{name:  "s", maxTime:   60},
	Minutes:			{name:  "m", maxTime:   60},
	Hours:				{name:  "h", maxTime:   24},
	Days:					{name:  "d", maxTime:    0},
}

func parseRun(cmd *cobra.Command, args []string) {
	switch unitFlag {
	case "us":
		unit = Microseconds
	case "ms":
		unit = Milliseconds
	case "s":
		unit = Seconds
	case "m":
		unit = Minutes
	case "h":
		unit = Hours
	case "d":
		unit = Days
	default:
		unit = Milliseconds
	}

	for _, x := range args {
		rest, err := strconv.Atoi(x)
		if (err != nil) {
			fmt.Println("failed")
			break
		}
		fmt.Printf("%v%s = ", rest, unitObjects[unit].name)

		var timeValue int
		strings := []string{"\n"}
		for i := unit; i < MAX_TIME_UNIT; i++ {
			if (rest <= 0) {
				break
			}

			// if last possible time unit
			if (i+1 == MAX_TIME_UNIT) {
				timeValue = rest
			} else {
				timeValue = rest % unitObjects[i].maxTime
				rest 			= rest / unitObjects[i].maxTime
			}
			strings = append(strings, strconv.Itoa(timeValue) + unitObjects[i].name + " ")
		}

		// print string-Array from back to front
		for i := range strings {
			fmt.Print(strings[len(strings)-1-i])
		}
	}
}

func init() {
	rootCmd.AddCommand(parseCmd)
	parseCmd.Flags().StringVarP(&unitFlag, "unit", "u", "ms", "Specifies the unit for the time value")
}
