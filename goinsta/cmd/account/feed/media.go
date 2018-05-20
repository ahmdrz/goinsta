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

package feed

import (
	"fmt"

	"github.com/cheggaaa/pb"
	"github.com/spf13/cobra"
	"github.com/ahmdrz/goinsta/utils"
)

// mediaCmd represents the media command
var mediaCmd = &cobra.Command{
	Use:     "media",
	Short:   "Download media in output directory",
	Example: "goinsta account feed media",
	Run: func(cmd *cobra.Command, args []string) {
		cmd = cmd.Root()

		inst := utils.New()

		output, err := cmd.Flags().GetString("output")
		if err != nil || output == "" {
			output = "./" + inst.Account.Username + "/"
		}

		media := inst.Account.Feed(nil)

		fmt.Println("Downloading your feed")
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

func init() {
	RootCmd.AddCommand(mediaCmd)
}
