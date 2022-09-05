/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/clantable/jukslv2-generator/config"
	"github.com/clantable/jukslv2-generator/internal/handler"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
)

// handlerCmd represents the handler command
var handlerCmd = &cobra.Command{
	Use:   "handler",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(cfgFile, &c)
		handler.BadgeGen(c)
	},
}

func init() {
	rootCmd.AddCommand(handlerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// handlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
