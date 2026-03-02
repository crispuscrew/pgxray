// Package cli provides command-line interface parsing for pgxray. 
// It uses the cobra library to define commands and flags, 
// and it dynamically generates flags based on the fields of the cfg.Profile struct. 
// The Execute function runs the CLI and returns a cfg.Profile with any overrides specified by the user.
package cli

import (
	"github.com/crispuscrew/pgxray/internal/cfg"
	"github.com/crispuscrew/pgxray/internal/opt"

	"log"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{                                                                                          
	Use:   "pgxray",                                                                                                   
	Short: "TUI PostgreSQL viewer",                                                                                    
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}
// TODO Implement init CLI command to generate default config and keybinds files
func init() {
	flags := rootCmd.Flags()

	
	forEachCliField(&cfg.Profile{}, func(field cfg.Field, name, shorthand, desc string) {
		innerKind := field.Meta.Type.Field(0).Type.Kind()
		switch innerKind {
		case reflect.String:
			flags.StringP(name, shorthand, "", desc)
		case reflect.Int:
			flags.IntP(name, shorthand, 0, desc)
		}
	})

	flags.StringP("config",  	"c", "", "path to config file")
	flags.StringP("profile", 	"p", "", "connection profile to use")
	flags.StringP("keybinds",	"k", "", "path to keybinds file")
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
	forEachCliField(&config.ProfileOverride, func(field cfg.Field, name, shorthand, desc string) {
		if !cmd.Flags().Changed(name) { return }

		innerKind := field.Meta.Type.Field(0).Type.Kind()
		switch innerKind {
		case reflect.String:
			v, _ := cmd.Flags().GetString(name) // Ignore error since flag existence is already checked
			field.Value.Set(reflect.ValueOf(opt.Set(v)))
		case reflect.Int:
			v, _ := cmd.Flags().GetInt(name) // Ignore error since flag existence is already checked
			field.Value.Set(reflect.ValueOf(opt.Set(v)))
		}
	})

	if cmd.Flags().Changed("config") {
		cfgPath, _ := cmd.Flags().GetString("config") // Ignore error since flag existence is already checked
		config.ConfigPath = opt.Set(cfgPath)
	}
	if cmd.Flags().Changed("profile") {
		profileName, _ := cmd.Flags().GetString("profile") // Ignore error since flag existence is already checked
		config.ProfileName = opt.Set(profileName)
	}
	if cmd.Flags().Changed("keybinds") {
		keybindsPath, _ := cmd.Flags().GetString("keybinds") // Ignore error since flag existence is already checked
		config.KeybindsPath = opt.Set(keybindsPath)
	}
	return config
}

func forEachCliField(p *cfg.Profile, fn func(field cfg.Field, name, shorthand, desc string)) {
	cfg.ForEachField(p, func(field cfg.Field) {
		tag := field.Meta.Tag.Get("cli")
		if tag == "" { return }

		parts := strings.SplitN(tag, ",", 3)
		fn(field, parts[0], parts[1], parts[2])
	})
}