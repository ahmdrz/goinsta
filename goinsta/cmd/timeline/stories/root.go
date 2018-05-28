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

package stories

import (
	"fmt"

	"github.com/ahmdrz/goinsta/utils"
	"github.com/cheggaaa/pb"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:     "stories",
	Short:   "Get stories of a user",
	Example: "goinsta user stories robpike",
	Run: func(cmd *cobra.Command, args []string) {
		cmd = cmd.Root()

		output, err := cmd.Flags().GetString("output")
		if err != nil || output == "" {
			output = "./timeline/stories/"
		}
		inst := utils.New()

		tray, err := inst.Timeline.Stories()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Downloading your timeline stories")
		for _, media := range tray.Stories {
			out := fmt.Sprintf("%s/%s/", output, media.User.Username)
			pgb := pb.StartNew(len(media.Items))
			for _, item := range media.Items {
				_, _, err := item.Download(out, "")
				if err != nil {
					fmt.Println(err)
				}
				pgb.Add(1)
			}
			pgb.Finish()
		}
	},
}
