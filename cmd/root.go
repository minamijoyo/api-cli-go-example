package cmd

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var RootCmd = &cobra.Command{
	Use:           "hoge",
	Short:         "A hoge CLI written in Go.",
	SilenceErrors: true,
	SilenceUsage:  true,
}

func init() {
	cobra.OnInitialize(initConfig)

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default $HOME/.hoge.yml)")
	RootCmd.PersistentFlags().StringP("url", "", "https://hoge.example.com/api", "hoge endpoint URL")

	viper.BindPFlag("url", RootCmd.PersistentFlags().Lookup("url"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	}

	viper.SetConfigName(".hoge")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()

	viper.ReadInConfig()
}

func newDefaultClient() (*Client, error) {
	endpointURL := viper.GetString("url")
	httpClient := &http.Client{}
	userAgent := fmt.Sprintf("hoge/%s (%s)", Version, runtime.Version())
	return newClient(endpointURL, httpClient, userAgent)
}
