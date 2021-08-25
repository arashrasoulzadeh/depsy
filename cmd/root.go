package cmd

import (
	"arashrasoulzadeh/deepzy/functions"
	"arashrasoulzadeh/deepzy/logger"
	"arashrasoulzadeh/deepzy/structs"
	"log"
	"os"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deepzy",
	Short: "A Simple Deployment manager",
	Long: `Deployment process made easy

By Arash Rasoulzadeh.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
		os.Exit(1)
	}
	str,err_cfg_read := functions.ReadFileToString(cfgFile)
	if err_cfg_read != nil{
		logger.StepBreak("Cannot load config file",cfgFile)
	}
	config := structs.Config{}
	err := yaml.Unmarshal([]byte(str), &config)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	functions.Execute(config)
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "deploy.yaml", "config file (default is deploy.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".deepzy" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".deepzy")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		log.Println("Using config file:", viper.ConfigFileUsed())
	}
}
