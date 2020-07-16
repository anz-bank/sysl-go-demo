package main

import (
	"log"
	"os"

	"github.com/anz-bank/sysl-go/config"
	petstore2 "github.com/anz-bank/sysl-template/internal/gen/petstore"
	"github.com/anz-bank/sysl-template/internal/petstore"
	"github.com/spf13/cobra"
)

var configFile string

func main() {
	cmdPetstore := &cobra.Command{
		Use:   "petstore",
		Short: "petstore",
		Run:   func(cmd *cobra.Command, args []string) { startServer() },
	}

	cmdPetstore.PersistentFlags().
		StringVarP(&configFile, "config", "c", "config-petstore.yaml", "Config file to read")

	rootCmd := &cobra.Command{}
	rootCmd.AddCommand(cmdPetstore)
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func startServer() {
	// load the configuration from the config file
	rootConfig := petstore.RootConfig{}
	defaultConfig := petstore2.NewDefaultConfig()
	if err := config.LoadConfig(configFile, &defaultConfig, &rootConfig); err != nil {
		log.Fatal(err)
	}

	// start the server
	cfg := petstore.ConfigContainer{DefaultConfig: defaultConfig, CustomConfig: rootConfig.CustomConfig}
	if err := petstore.ListenAndServe(&cfg); err != nil {
		log.Fatal(err)
	}
}
