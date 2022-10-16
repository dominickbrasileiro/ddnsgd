package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/dominickbrasileiro/ddnsgd/internal"
	"github.com/spf13/cobra"
)

var interval int
var username string
var password string
var hostname string

var rootCmd = &cobra.Command{
	Use:   "ddnsgd",
	Short: "TODO",
	Long:  "TODO",
	Run: func(cmd *cobra.Command, args []string) {
		config := internal.AppConfig{
			Interval: interval,
			Username: username,
			Password: password,
			Hostname: hostname,
		}

		logger := log.Default()

		internal.Run(&config, logger)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&interval, "interval", "i", 0, "The interval in seconds to fetch the IPv4 address and update the DNS record.")
	rootCmd.Flags().StringVarP(&username, "username", "u", "", "The generated username associated with the host that is to be updated.")
	rootCmd.Flags().StringVarP(&password, "password", "p", "", "The generated password associated with the host that is to be updated.")
	rootCmd.Flags().StringVarP(&hostname, "hostname", "H", "", "subdomain.yourdomain.com")

	rootCmd.MarkFlagRequired("interval")
	rootCmd.MarkFlagRequired("username")
	rootCmd.MarkFlagRequired("password")
	rootCmd.MarkFlagRequired("hostname")
}
