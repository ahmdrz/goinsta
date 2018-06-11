// Copyright © 2018 Mester
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package following

import (
	"fmt"
	"strconv"

	"github.com/ahmdrz/goinsta/utils"
	"github.com/spf13/cobra"
)

//RootCmd is used as a command line interaction with Instagram get user following method.
var RootCmd = &cobra.Command{
	Use:     "following",
	Short:   "Get user following",
	Example: "goinsta user following robpike",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing args. See example.")
			return
		}
		inst := utils.New()

		user, err := inst.Profiles.ByName(args[0])
		if err != nil {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			user, err = inst.Profiles.ByID(id)
			if err != nil {
				fmt.Printf("Invalid username or id: %s\n", args[0])
				return
			}
		}

		users := user.Following()

		i := 0
		fmt.Println("Followers:\n  ID\t\tUsername")
		for users.Next() {
			for _, u := range users.Users {
				i++
				fmt.Printf("  %d\t%s\n", u.ID, u.Username)
			}
		}
		fmt.Printf("Total: %d\n", i)
	},
}
