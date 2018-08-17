// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spilliams/pokemon/pokedex"

	"github.com/spf13/cobra"
)

// renamefilesCmd represents the renamefiles command
var renamefilesCmd = &cobra.Command{
	Use:   "renamefiles",
	Short: "Looks in the local directory for files named with integers. Copies those files and names the new ones based on the name of the pokemon at that index",
	Long: `Looks in the local directory for files named with integers. Copies those files and names the new ones based on the name of the pokemon at that index. For example:

	$ pokemon renamefiles
	151 files renamed and placed in "renamedfiles"

will look for files in this directory, and place new files in a new "renamedfiles" directory.

Supports jpg, jpeg, gif, and png image formats.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dex, err := pokedex.AllPokemon()
		if err != nil {
			return err
		}
		dex.FilterBy("IsDefault", true)
		// fmt.Println(dex)
		_ = os.Mkdir("renamedfiles", 0755)
		count := 0
		for _, p := range dex.Pokemon {
			one, err := renameOneFile(p.SpeciesID, p.Name)
			count += one
			if err != nil {
				return err
			}
		}
		fmt.Printf("%v files renamed and placed in \"renamedfiles\"\n", count)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(renamefilesCmd)
}

// renames all files with the same name (e.g. 4.jpg, 4.gif, 4.png) to the pokemon name (e.g. renamedfiles/charmander.jpg etc)
// returns the number of files renamed in this way
func renameOneFile(species int, name string) (int, error) {
	extensions := []string{
		"jpg",
		"jpeg",
		"png",
		"gif",
	}
	count := 0
	for _, ext := range extensions {
		lookfor := fmt.Sprintf("%v.%v", species, ext)
		renameTo := fmt.Sprintf("renamedfiles/%v.%v", name, ext)
		if _, err := os.Stat(lookfor); err == nil {
			bytes, err := ioutil.ReadFile(lookfor)
			if err != nil {
				return count, err
			}
			err = ioutil.WriteFile(renameTo, bytes, 0644)
			if err != nil {
				return count, err
			}
			count++
		}
	}
	return count, nil
}
