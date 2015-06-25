package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

func check_version() string {
	return "check_psqlcompare - v1.1.0 (c)2015 Martin Seener - Cash Payment Solutions GmbH"
}

func show_help() string {
	return `check_psqlcompare

			You can compare two or more SQL queries with each other. The check will exit OK if all queries have the same output. If not it will exit CRITICAL
			`
}

var (
	compare    = kingpin.Flag("compare", "Enable comparison mode for 2 or more SQL Queries.").Short('c').Bool()
	sqlqueries = kingpin.Arg("sqlqueries", "Complete sets of strings which contain a user, password, server, port, database and the query to execute on. Format: \"postgres(ql)://username:password@hostname:5432/database#query\"").Strings()
)

func main() {
	kingpin.Version(check_version())
	kingpin.CommandLine.Help = show_help()
	kingpin.Parse()

	if *compare {
		if len(*sqlqueries) >= 2 {
			// Check if psql is available
			cmd := exec.Command("which", "psql")
			_, err := cmd.Output()
			if err != nil {
				fmt.Println("UNKNOWN - Unable to find psql")
				os.Exit(3)
			}

			// Creating slice for the output
			output := make([]string, len(*sqlqueries))
			// Going through each entry to get its value from the query
			for i := 0; i < len(*sqlqueries); i++ {
				// Split up the connection information and the query and validate if there was a split event
				args := strings.Split((*sqlqueries)[i], "#")
				if len(args) != 2 {
					i += 1
					fmt.Printf("UNKNOWN - Query %v does not have the right format! It probably misses the '#' delimiter.\n", i)
					os.Exit(3)
				}
				// Execute the query and save the output to the new slice - Exit if an error occures
				cmd := exec.Command("psql", "-At", args[0], "-c", args[1])
				out, err := cmd.Output()
				if err != nil {
					fmt.Println("UNKNOWN - Error occured:", err.Error())
					os.Exit(3)
				}
				// Writing the psql output to the output slice
				output[i] = strings.TrimSpace(string(out))
			}

			// Let us compare the output and act accordingly - using first slice value as reference
			for i := 1; i < len(*sqlqueries); i++ {
				if output[0] != output[i] {
					querynum := i + 1
					fmt.Printf("CRITICAL - Output of given Queries does not match! Reference Query 1 (Result: %v) does not match Query %v (Result: %v)\n", output[0], querynum, output[i])
					os.Exit(2)
				}
			}
			fmt.Printf("OK - Output of %v Queries is identical (Result: %v).\n", len(*sqlqueries), output[0])
			os.Exit(0)
		} else {
			fmt.Println("UNKNOWN - You have given less than 2 SQL Queries. Aborting.")
			os.Exit(3)
		}
	}
	kingpin.Usage()
	os.Exit(3)
}
