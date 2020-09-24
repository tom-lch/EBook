package cmd

import (
	"admin/global"
	"admin/tools"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use: 	"admin",
	Short:  "admin",
	SilenceUsage: true,
	Long: 	"admin",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			tip()
			return errors.New("requires at least one arg")
		}
		return nil
	},
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		tip()
	},
}

func tip() {
	usageStr := `欢迎使用 `+ tools.Green( `admin ` +global.Version) + ` 可以使用 ` + tools.Red(`-h`) + ` 查看命令`
	fmt.Printf("%s\n", usageStr)
}


func init() {
	rootCmd.AddCommand()
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}