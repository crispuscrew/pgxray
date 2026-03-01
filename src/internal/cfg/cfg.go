// Package cfg provides configuration by config file and CLI overrides for pgxray.
package cfg

import (
	"log"
	"os"
	"fmt"
	"errors"

	toml "github.com/pelletier/go-toml/v2"
)

var ErrFileNotFound = errors.New("File not found")
func BuildConfig(cliOverride CliConfig) Config {
	profile := defaultProfile
	keybinds := defaultKeybinds
	var warnings []string

	if cliOverride.ProfileName.IsSet() {
		profile.Name = cliOverride.ProfileName.Get()
	}

	configPath, err := resolvePath(cliOverride.ConfigPath, configPathEnvVar, defaultConfigPath)
	if err != nil {
		warnings = append(warnings, err.Error())
	}

	prFromFile, err := profileFromFile(configPath, profile.Name)
	if errors.Is(err, ErrFileNotFound) {
		warnings = append(warnings, err.Error())
	} else if err != nil {
		log.Fatalf("could not load profile from file: %v", err)
	} else {
		profile = merge(profile, prFromFile)
	}

	profile = merge(profile, cliOverride.ProfileOverride)

	keybindsPath, err := resolvePath(cliOverride.KeybindsPath, keybindsPathEnvVar, defaultKeybindsPath)
	if err != nil {
		warnings = append(warnings, err.Error())
	}

	kbFromFile, err := keybindsFromFile(keybindsPath)
	if errors.Is(err, ErrFileNotFound) {
		warnings = append(warnings, err.Error())
	} else if err != nil {
		log.Fatalf("could not load keybinds from file: %v", err)
	} else {
		keybinds = merge(keybinds, kbFromFile)
	}

	return Config{
		Profile:  profile,
		Keybinds: keybinds,
		Warnings: warnings,
	}
}

func profileFromFile(path, profileName string) (Profile, error) {
	type fileConfig struct {
		Connections map[string]Profile `toml:"connections"`
	}
	var fc fileConfig
	data, err := os.ReadFile(path)
	if errors.Is(err, os.ErrNotExist) {                                                                                                                                
		return Profile{}, fmt.Errorf("could not read config file, %w: %w, fallback to default value", ErrFileNotFound, err)
	} else if err != nil {
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
	if errors.Is(err, os.ErrNotExist) {                                                                                                                                    
		return Keybinds{}, fmt.Errorf("could not read keybinds file, %w: %w", ErrFileNotFound, err)
	} else if err != nil {
		return Keybinds{}, fmt.Errorf("could not read keybinds file: %w", err)
	}

	if err := toml.Unmarshal(data, &kb); err != nil {
		return Keybinds{}, fmt.Errorf("could not parse keybinds file: %w", err)
	}
	return kb, nil
}