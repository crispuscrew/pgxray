package cfg

import (
	"github.com/crispuscrew/pgxray/src/internal/opt"

	"os"
	"testing"
	"path/filepath"
)

func TestBuildConfig_Defaults(t *testing.T) {
	config := BuildConfig(CliConfig{})
	if config.Profile.Host.Get() != "localhost" {
		t.Errorf("expected localhost, got %s", config.Profile.Host.Get())
	}
	if config.Profile != defaultProfile {
		t.Errorf("expected default profile, got %+v", config.Profile)
	}
}

func TestBuildConfig_FileOverridesDefaults(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.toml")
	os.WriteFile(path, []byte(`
	[connections.default]
	host = "myserver"
	`), 0644)

	config := BuildConfig(CliConfig{
		ConfigPath: opt.Set(path),
	})
	if config.Profile.Host.Get() != "myserver" {
		t.Errorf("expected myserver, got %s", config.Profile.Host.Get())
	}
}