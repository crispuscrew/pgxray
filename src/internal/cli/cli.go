// Package cli provides command-line interface parsing for pgxray. 
// It uses the cobra library to define commands and flags, 
// and it dynamically generates flags based on the fields of the cfg.Profile struct. 
// The Execute function runs the CLI and returns a cfg.Profile with any overrides specified by the user.
package cli

import (
	"github.com/crispuscrew/pgxray/src/internal/cfg"
	"github.com/crispuscrew/pgxray/src/internal/opt"

	"github.com/spf13/cobra"
	"log"
	"reflect"
	"strings"
)

var rootCmd = &cobra.Command{                                                                                          
	Use:   "pgxray",                                                                                                   
	Short: "TUI PostgreSQL viewer",                                                                                    
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	t := reflect.TypeOf(cfg.Profile{}) 
	flags := rootCmd.Flags()                                         
																	
	for i := range t.NumField() {
		field := t.Field(i)
		tag := field.Tag.Get("cli")
		if tag == "" {
			continue
		}

		parts := strings.SplitN(tag, ",", 3)
		name, shorthand, desc := parts[0], parts[1], parts[2]

		innerKind := field.Type.Field(0).Type.Kind()
		switch innerKind {
		case reflect.String:
			flags.StringP(name, shorthand, "", desc)
		case reflect.Int:
			flags.IntP(name, shorthand, 0, desc)
		}
	}

	flags.StringP("config",  "c", "", "path to config file")
	flags.StringP("profile", "p", "", "connection profile to use")
}

func Execute() (cfg.CliConfig) {
	var config cfg.CliConfig
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		config = buildConfig(cmd)
		return nil
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("could not parse CLI: %v", err)
	}
	return config
}

func buildConfig(cmd *cobra.Command) cfg.CliConfig {
	var config cfg.CliConfig
	overrideVal := reflect.ValueOf(&config.ProfileOverride).Elem()                     
	fields := reflect.VisibleFields(reflect.TypeOf(cfg.CliConfig.ProfileOverride{}))       
									
	for i, field := range fields {
		tag := field.Tag.Get("cli")
		name := strings.SplitN(tag, ",", 3)[0]
		if tag == "" || !cmd.Flags().Changed(name) { continue }

		innerKind := field.Type.Field(0).Type.Kind()
		switch innerKind {
		case reflect.String:
			v, _ := cmd.Flags().GetString(name)
			overrideVal.Field(i).Set(reflect.ValueOf(opt.Set(v)))
		case reflect.Int:
			v, _ := cmd.Flags().GetInt(name)
			overrideVal.Field(i).Set(reflect.ValueOf(opt.Set(v)))
		}
	}

	if cmd.Flags().Changed("config") {
		cfgPath, _ := cmd.Flags().GetString("config")
		config.ConfigPath = opt.Set(cfgPath)
	}
	if cmd.Flags().Changed("profile") {
		profileName, _ := cmd.Flags().GetString("profile")
		config.ProfileName = opt.Set(profileName)
	}
	return config
}