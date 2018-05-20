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

package facebook

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/ahmdrz/goinsta/utils"
)

var RootCmd = &cobra.Command{
	Use:     "facebook",
	Short:   "Search facebook users on Instagram",
	Example: "goinsta search facebook rob pike",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Missing arguments. See example")
			return
		}

		inst := utils.New()

		res, err := inst.Search.Facebook(strings.Join(args, " "))
		if err != nil {
			fmt.Println(err)
			return
		}
		if res.NumResults == 0 {
			fmt.Println("No results.")
			return
		}

		fmt.Printf("Results: %d\n", res.NumResults)
		for i := range res.Users {
			fmt.Printf("  %s\n", res.Users[i].Username)
		}
	},
}
