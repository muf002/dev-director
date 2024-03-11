package director

import (
	"fmt"

	"github.com/muf002/dev-director/pkg/directory"
	"github.com/spf13/cobra"
)

var directories []string
var isOpen bool
var isClose bool

var DirectorCmd = &cobra.Command{
	Use:   "director",
	Short: "Main command for the program",
	Run: func(cmd *cobra.Command, args []string) {
		if isClose {
			fmt.Println("then came here for false")
		} else {
			fmt.Println("then came here")
			directory.OpenDirectories(directories)
		}
	},
}

func init() {
	DirectorCmd.Flags().StringSliceVarP(&directories, "directory", "d", []string{}, "store all the directories")
	DirectorCmd.Flags().BoolVar(&isOpen, "open", true, "to indicate the open flag")
	DirectorCmd.Flags().BoolVar(&isClose, "close", false, "to indicate the close flag")
}
