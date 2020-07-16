package main

import (
	"log"
	"os"

	"github.com/anz-bank/sysl-go/config"
	petdemo2 "github.com/anz-bank/sysl-template/internal/gen/petdemo"
	"github.com/anz-bank/sysl-template/internal/petdemo"
	"github.com/spf13/cobra"
)

var configFile string

func main() {
	cmdPetdemo := &cobra.Command{
		Use:   "petdemo",
		Short: "petdemo",
		Run:   func(cmd *cobra.Command, args []string) { startServer() },
	}

	cmdPetdemo.PersistentFlags().
		StringVarP(&configFile, "config", "c", "config-petdemo.yaml", "Config file to read")

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmdPetdemo)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func startServer() {
	// load the configuration from the config file
	rootConfig := petdemo.RootConfig{}
	defaultConfig := petdemo2.NewDefaultConfig()
	if err := config.LoadConfig(configFile, &defaultConfig, &rootConfig); err != nil {
		log.Fatal(err)
	}

	// start the server
	cfg := petdemo.ConfigContainer{DefaultConfig: defaultConfig, CustomConfig: rootConfig.CustomConfig}
	if err := petdemo.ListenAndServe(&cfg); err != nil {
		log.Fatal(err)
	}
}
