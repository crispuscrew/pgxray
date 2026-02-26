// Package cfg provides configuration by config file and CLI overrides for pgxray.
package cfg

import (
	"github.com/crispuscrew/pgxray/src/internal/opt"

	"log"
	"os"
	"fmt"

	toml "github.com/pelletier/go-toml/v2"
)

func BuildConfig(cliOverride CliConfig) (Profile, Keybinds, []string) {
	profile := defaultProfile
	keybinds := defaultKeybinds
	var warnings []string

	configPath, err := resolveConfigPath(cliOverride.ConfigPath)
	if err != nil {
		warnings = append(warnings, err.Error())
	}

	keybindsPath, err := resolveKeybindsPath(cliOverride.KeybindsPath)
	if err != nil {
		warnings = append(warnings, err.Error())
	}

	if opt.IsSet(cliOverride.ProfileName) {
		profile.Name = opt.Get(cliOverride.ProfileName)
	}

	prFromFile, err := profileFromFile(configPath, profile.Name)
	if err != nil {
		log.Fatalf("could not load profile from file: %v", err)
	}
	profile = merge(merge(prFromFile, profile), cliOverride.ProfileOverride)

	kbFromFile, err := keybindsFromFile(keybindsPath)
	if err != nil {
		warnings = append(warnings, err.Error())
	} else {
		keybinds = merge(keybinds, kbFromFile)
	}

	return profile, keybinds, warnings
}

func resolveConfigPath(cliPath opt.Opt[string]) (string, error) {
	if opt.IsSet(cliPath) {
		return resolvePath(opt.Get(cliPath))
	}
	if os.Getenv(configPathEnvVar) != "" {
		return resolvePath(os.Getenv(configPathEnvVar))
	}
	//Always can resolve to default path, so ignore error
	path, _ := resolvePath(defaultConfigPath)
	return path, fmt.Errorf("warning: could not resolve config path: %v, using defaults and CLI overrides", cliPath)
}

func resolveKeybindsPath(cliPath opt.Opt[string]) (string, error) {
	if opt.IsSet(cliPath) {
		return resolvePath(opt.Get(cliPath))
	}
	if os.Getenv(keybindsPathEnvVar) != "" {
		return resolvePath(os.Getenv(keybindsPathEnvVar))
	}
	//Always can resolve to default path, so ignore error
	path, _ := resolvePath(defaultKeybindsPath)
	return path, fmt.Errorf("warning: could not resolve keybinds path: %v, using defaults and CLI overrides", cliPath)
}

func profileFromFile(path, profileName string) (Profile, error) {
	type fileConfig struct {
		Connections map[string]Profile `toml:"connections"`
	}
	var fc fileConfig
	data, err := os.ReadFile(path)
	if err != nil {                                                                                                                                    
		return Profile{}, fmt.Errorf("could not read config file: %w", err)
	}
	if err := toml.Unmarshal(data, &fc); err != nil {
		return Profile{}, fmt.Errorf("could not parse config file: %w", err)
	}

	profile, exists := fc.Connections[profileName]
	if !exists {
		return Profile{}, fmt.Errorf("profile name (%s) not found in config file", profileName)
	}

	return profile, nil
}

func keybindsFromFile(path string) (Keybinds, error) {
	var kb Keybinds
	data, err := os.ReadFile(path)
	if err != nil {                                                                                                                                    
		return Keybinds{}, fmt.Errorf("could not read keybinds file: %w", err)
	}
	if err := toml.Unmarshal(data, &kb); err != nil {
		return Keybinds{}, fmt.Errorf("could not parse keybinds file: %w", err)
	}
	return kb, nil
}

func merge[T any](base, override T) T {
	ForEachFieldPair(&base, &override, func(field, overrideField Field) {
		if set := overrideField.Value.FieldByName("Set"); set.IsValid() && set.Bool() {
			field.Value.Set(overrideField.Value)
		}
	})
	return base
}