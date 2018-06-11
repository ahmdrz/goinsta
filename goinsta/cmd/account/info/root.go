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

	"github.com/ahmdrz/goinsta/utils"
	"github.com/spf13/cobra"
)

//RootCmd is used as a command line interaction with Instagram Profile info.
var RootCmd = &cobra.Command{
	Use:     "info",
	Short:   "Get partial info about your account",
	Example: "goinsta account info",
	Run: func(cmd *cobra.Command, args []string) {
		inst := utils.New()

		user, _ := inst.Profiles.ByID(inst.Account.ID)

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
`, user.Username, user.FullName, user.ID, user.ProfilePicURL,
			user.PublicEmail, user.Gender, user.Biography, user.FollowerCount,
			user.FollowingCount)
	},
}
