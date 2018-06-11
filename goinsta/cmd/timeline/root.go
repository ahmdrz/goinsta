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

package timeline

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta/goinsta/cmd/timeline/feed"
	"github.com/ahmdrz/goinsta/goinsta/cmd/timeline/stories"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(feed.RootCmd)
	RootCmd.AddCommand(stories.RootCmd)
}

//RootCmd is used as a command line interaction with Instagram Timeline related methods.
var RootCmd = &cobra.Command{
	Use:   "timeline",
	Short: "Get timeline feed or stories",
}

//Execute the method to to start the command execution of Instagram Timeline related methods.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
