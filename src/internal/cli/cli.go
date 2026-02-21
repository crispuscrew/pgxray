package cli

import (
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{                                                                                          
	Use:   "pgxray",                                                                                                   
	Short: "TUI PostgreSQL viewer",                                                                                    
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

type CliConfig struct {
	ConfigPath string
	Profile    string

	Host     string
	Port     int
	User     string
	Database string
}

func init() {
	rootCmd.Flags().StringP("config"	, "c", ""	, "config file path"	)
	rootCmd.Flags().StringP("profile"	, "p", ""	, "connection profile"	)

	rootCmd.Flags().StringP("host"		, "H", ""	, "override host"		)
	rootCmd.Flags().IntP(	"port"		, "P", 5432	, "override port"		)
	rootCmd.Flags().StringP("user"		, "u", ""	, "connection user"		)
	rootCmd.Flags().StringP("database"	, "d", ""	, "connection database"	)
}

func Execute() (CliConfig) {
	var cfg CliConfig
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		cfg = buildConfig(cmd)
		return nil
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("could not parse CLI: %v", err)
	}
	return cfg
}

func buildConfig(cmd *cobra.Command) CliConfig {        
	mustString := func (name string) string {
		v, _ := cmd.Flags().GetString(name)
		return v
	}                                                               
	port, _ := cmd.Flags().GetInt("port")                                                                              
	return CliConfig{
		ConfigPath: mustString("config"),
		Profile:    mustString("profile"),
		Host:       mustString("host"),
		Port:       port,
		User:       mustString("user"),
		Database:   mustString("database"),
	}
}