package generate

import (
	"fmt"

	"github.com/flow-cli/internal/generate"
	"github.com/spf13/cobra"
)

var includeSpecial bool
var length int
var defaultPasswordLength int = 8

var passwordCmd = &cobra.Command{
	Use:     "password",
	Aliases: []string{"pass"},
	Short:   "Generates a random password",
	Run: func(cmd *cobra.Command, args []string) {
		password := generate.GeneratePassword(length, includeSpecial)
		fmt.Println(password)
	},
}

func init() {
	passwordCmd.Flags().IntVarP(&length, "length", "l", defaultPasswordLength, "Length of the password. Default to 8")
	passwordCmd.Flags().BoolVarP(&includeSpecial, "specials", "s", false, "Include specials in generated password. Defaulted to false")
}
