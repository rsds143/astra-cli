//   Copyright 2021 Ryan Svihla
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

//Package db provides the sub-commands for the db command
package db

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/rsds143/astra-cli/pkg"
	"github.com/rsds143/astra-devops-sdk-go/astraops"
)

var limit int
var include string
var provider string
var startingAfter string
var listFmt string

func init(){
    ListCmd.Flags().IntVarP(&limit, "limit", "l", 10, "limit of databases retrieved")
    ListCmd.Flags().StringVarP(&include, "include","i", "", "the type of filter to apply")
    ListCmd.Flags().StringVarP(&provider, "provider", "p", "", "provider to filter by")
    ListCmd.Flags().StringVarP(&startingAfter, "startingAfter", "a", "", "timestamp filter, ie only show databases created after this timestamp")
    ListCmd.Flags().StringVarP(&listFmt, "format", "f", "text", "Output format for report default is json")
}
//ListCmd provides the list databases command
var ListCmd =  &cobra.Command{
  Use:   "list",
  Short: "lists all databases",
  Long: `lists all databases in your Astra account`,
  Run: func(cmd *cobra.Command, args []string) {
    client, err := pkg.LoginClient()
	if err != nil {
	    fmt.Fprintln(os.Stderr, fmt.Sprintf("unable to login with error %v", err))
        os.Exit(1)
    }
	var dbs []astraops.Database
	if dbs, err = client.ListDb(include, provider, startingAfter, int32(limit)); err != nil {
	    fmt.Fprintln(os.Stderr, fmt.Sprintf("unable to get list of dbs with error %v", err))
	    os.Exit(1)
	}
	switch listFmt {
	case "text":
		var rows [][]string
		rows = append(rows, []string{"name", "id", "status"})
		for _, db := range dbs {
			rows = append(rows, []string{db.Info.Name, db.ID, string(db.Status)})
		}
		for _, row := range pkg.PadColumns(rows) {
			fmt.Println(strings.Join(row, " "))
		}
	case "json":
		b, err := json.MarshalIndent(dbs, "", "  ")
		if err != nil {
	        fmt.Fprintln(os.Stderr, fmt.Sprintf("unexpected error marshaling to json: '%v', Try -output text instead", err))
	        os.Exit(1)
		}
		fmt.Println(string(b))
	default:
	    fmt.Fprintln(os.Stderr, fmt.Sprintf("-output %q is not valid option.", getFmt))
	    os.Exit(1)
	}
},
}
