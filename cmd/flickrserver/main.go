package main

import (
	"log"
	"os"

	"github.com/anz-bank/sysl-go/config"
	"github.com/anz-bank/sysl-template/internal/flickr"
	flickr2 "github.com/anz-bank/sysl-template/internal/gen/flickr"
	"github.com/spf13/cobra"
)

var configFile string

func main() {
	cmdFlickr := &cobra.Command{
		Use:   "flickr",
		Short: "flickr",
		Run:   func(cmd *cobra.Command, args []string) { startServer() },
	}

	cmdFlickr.PersistentFlags().
		StringVarP(&configFile, "config", "c", "config-flickr.yaml", "Config file to read")

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmdFlickr)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func startServer() {
	// load the configuration from the config file
	rootConfig := flickr.RootConfig{}
	defaultConfig := flickr2.NewDefaultConfig()
	if err := config.LoadConfig(configFile, &defaultConfig, &rootConfig); err != nil {
		log.Fatal(err)
	}

	// start the server
	cfg := flickr.ConfigContainer{DefaultConfig: defaultConfig, CustomConfig: rootConfig.CustomConfig}
	if err := flickr.ListenAndServe(&cfg); err != nil {
		log.Fatal(err)
	}
}
