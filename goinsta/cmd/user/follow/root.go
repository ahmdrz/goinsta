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

package follow

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"gopkg.in/ahmdrz/goinsta.v2/utils"
)

var RootCmd = &cobra.Command{
	Use:     "follow",
	Short:   "Start following a user",
	Example: "goinsta user follow robpike",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing arguments. See example.")
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
		user.FriendShip()

		if user.Friendship.Following {
			fmt.Printf("You are currently following %s.\n", user.Username)
			return
		}

		err = user.Follow()
		if err != nil {
			fmt.Printf("error following %s: %s\n", user.Username, err)
			return
		}
		user.FriendShip()

		if user.IsPrivate {
			fmt.Printf("%s is private. Requested.\n", user.Username)
			return
		}

		fmt.Printf("Following %s: %v\n", user.Username, user.Friendship.Following)
	},
}
