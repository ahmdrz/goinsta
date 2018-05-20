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

package user

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/block"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/feed"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/follow"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/followers"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/following"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/info"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/stories"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/unblock"
	"gopkg.in/ahmdrz/goinsta.v2/goinsta/cmd/user/unfollow"
)

func init() {
	RootCmd.AddCommand(feed.RootCmd)
	RootCmd.AddCommand(story.RootCmd)
	RootCmd.AddCommand(info.RootCmd)
	RootCmd.AddCommand(follow.RootCmd)
	RootCmd.AddCommand(followers.RootCmd)
	RootCmd.AddCommand(following.RootCmd)
	RootCmd.AddCommand(unfollow.RootCmd)
	RootCmd.AddCommand(block.RootCmd)
	RootCmd.AddCommand(unblock.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "user",
	Short: "Get downloads specified Instagram user's object.",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
