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

package info

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/ahmdrz/goinsta/utils"
)

var RootCmd = &cobra.Command{
	Use:     "info",
	Short:   "Get partial info about user",
	Example: "goinsta user info robpike",
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
				fmt.Printf("Invalid username or id '%s': %s\n", args[0], err)
				return
			}
		}
		user.FriendShip()

		fmt.Printf(`
Username: %s
Fullname: %s
ID: %d
ProfilePicURL: %s
Email: %s
Gender: %d
Biography: %s
Followers: %d
Following: %d
You follow him/her: %v
`, user.Username, user.FullName, user.ID, user.ProfilePicURL,
			user.PublicEmail, user.Gender, user.Biography, user.FollowerCount,
			user.FollowingCount, user.Friendship.Following)
	},
}
