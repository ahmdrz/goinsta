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

package cmd

import (
	"fmt"
	"os"

	"github.com/ahmdrz/goinsta/goinsta/cmd/account"
	"github.com/ahmdrz/goinsta/goinsta/cmd/logout"
	"github.com/ahmdrz/goinsta/goinsta/cmd/search"
	"github.com/ahmdrz/goinsta/goinsta/cmd/timeline"
	"github.com/ahmdrz/goinsta/goinsta/cmd/user"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(user.RootCmd)
	rootCmd.AddCommand(search.RootCmd)
	rootCmd.AddCommand(account.RootCmd)
	rootCmd.AddCommand(timeline.RootCmd)
	rootCmd.AddCommand(logout.RootCmd)
}

var (
	username string
)

var rootCmd = &cobra.Command{
	Use:   "goinsta",
	Short: "Command line tool for Instagram",
}

// Execute is called from main as an entrance to the command line interaction tool.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "", "Output directory")
}
