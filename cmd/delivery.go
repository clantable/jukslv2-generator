/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/clantable/jukslv2-generator/config"
	"github.com/clantable/jukslv2-generator/internal/delivery"
	"github.com/clantable/jukslv2-generator/util"
	"github.com/spf13/cobra"
	"github.com/zeromicro/go-zero/core/conf"
)

// deliveryCmd represents the delivery command
var deliveryCmd = &cobra.Command{
	Use:   "delivery",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		var c config.Config
		conf.MustLoad(cfgFile, &c)
		err := util.GenDir(c.DeliveryConfig.Path)
		if err != nil {
			panic(err)
		}
		delivery.BadgeGen(c)
	},
}

func init() {
	rootCmd.AddCommand(deliveryCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deliveryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deliveryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
