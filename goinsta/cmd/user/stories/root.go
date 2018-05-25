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

package story

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ahmdrz/goinsta/utils"
	"github.com/cheggaaa/pb"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "stories",
	Short:   "Get stories of a user",
	Example: "goinsta user stories robpike",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing arguments. See example.")
			return
		}
		cmd = cmd.Root()

		output, err := cmd.Flags().GetString("output")
		if err != nil || output == "" {
			output = "./" + args[0] + "/stories/"
		}
		inst := utils.New()

		user, err := inst.Profiles.ByName(args[0])
		if err != nil {
			id, _ := strconv.ParseInt(args[0], 10, 64)
			user, err = inst.Profiles.ByID(id)
			if err != nil {
				fmt.Printf("Invalid username or id: %s\n", args[0])
				os.Exit(1)
			}
		}

		media := user.Stories()

		fmt.Println("Downloading stories of", user.Username)
		for media.Next() {
			pgb := pb.StartNew(len(media.Items))
			for _, item := range media.Items {
				err := item.Download(output, "")
				if err != nil {
					fmt.Println(err)
				}
				pgb.Add(1)
			}
			pgb.Finish()
		}
	},
}
