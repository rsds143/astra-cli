//  Copyright 2021 Ryan Svihla
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

package cmd

import (
    "github.com/spf13/cobra"
	"github.com/rsds143/astra-cli/cmd/db"
)

func init(){
    dbCmd.AddCommand(db.CreateCmd)
    dbCmd.AddCommand(db.DeleteCmd)
    dbCmd.AddCommand(db.ParkCmd)
    dbCmd.AddCommand(db.UnparkCmd)
    dbCmd.AddCommand(db.ResizeCmd)
    dbCmd.AddCommand(db.GetCmd)
    dbCmd.AddCommand(db.ListCmd)
    dbCmd.AddCommand(db.TiersCmd)
}

var dbCmd =  &cobra.Command{
  Use:   "db",
  Short: "Shows all the db commands",
  Long:  `Shows all other db commands. Create, Delete, Get information on your databases`,
  Run: func(cobraCmd *cobra.Command, args []string) {
	cobraCmd.Usage()
    },
}
