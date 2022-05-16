package main

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/niftyConnect/order-synchronizer/command"
	cmm "github.com/niftyConnect/order-synchronizer/common"
)

// Bind all flags and read the config into viper
func bindFlagsLoadViper(cmd *cobra.Command, args []string) error {
	// cmd.Flags() includes flags from this command and all persistent flags from the parent
	if err := viper.BindPFlags(cmd.Flags()); err != nil {
		return err
	}
	return nil
}

type cobraCmdFunc func(cmd *cobra.Command, args []string) error

func concatCobraCmdFuncs(fs ...cobraCmdFunc) cobraCmdFunc {
	return func(cmd *cobra.Command, args []string) error {
		for _, f := range fs {
			if f != nil {
				if err := f(cmd, args); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
func main() {
	rootCmd := &cobra.Command{
		Use:   "order-synchronizer",
		Short: "Command line interface for order synchronizer",
	}

	rootCmd.PersistentFlags().String(cmm.FlagConfigPath, "", "config file path")

	rootCmd.AddCommand(
		command.Start(),
	)
	// prepare and add flags
	rootCmd.PersistentPreRunE = concatCobraCmdFuncs(bindFlagsLoadViper, rootCmd.PersistentPreRunE)
	rootCmd.SilenceUsage = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
