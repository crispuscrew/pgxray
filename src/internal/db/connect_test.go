package db                                                    
				
import (
	"github.com/crispuscrew/pgxray/internal/opt"
	"github.com/crispuscrew/pgxray/internal/cfg"

	"os"
	"testing"
	"fmt"
)

var testConn *Conn

func TestMain(m *testing.M) {
	config := cfg.BuildConfig(cfg.CliConfig{                      
		ConfigPath: opt.Set("testdata/test-config.toml"),         
	})

	conn, err := Connect(config.Profile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect failed: %v\n", err)
		os.Exit(1)
	}

	testConn = conn
	os.Exit(m.Run())
}

func TestConnect_WithPassword(t *testing.T) {
	os.Setenv("PGPASSWORD", "secret")
	defer os.Setenv("PGPASSWORD", "test") // restore original

	config := cfg.BuildConfig(cfg.CliConfig{
		ConfigPath: opt.Set("testdata/test-config-pw.toml"),
	})
	conn, err := Connect(config.Profile)
	if err != nil { t.Fatalf("Expected connect with correct password to succeed: %v", err) }
	_ = conn
}

func TestConnect_NoPassword(t *testing.T) {
	os.Unsetenv("PGPASSWORD")
	defer os.Setenv("PGPASSWORD", "test") // restore original

	config := cfg.BuildConfig(cfg.CliConfig{
		ConfigPath: opt.Set("testdata/test-config-pw.toml"),
	})
	_, err := Connect(config.Profile)
	if err == nil { t.Errorf("Expected error when no password provided, got nil") }
}

func TestConnect_WrongPassword(t *testing.T) {
	os.Setenv("PGPASSWORD", "wrongpassword")
	defer os.Setenv("PGPASSWORD", "test") // restore original

	config := cfg.BuildConfig(cfg.CliConfig{
		ConfigPath: opt.Set("testdata/test-config-pw.toml"),
	})
	_, err := Connect(config.Profile)
	if err == nil { t.Errorf("Expected error with wrong password, got nil") }
}

func TestConnect_WrongDatabase(t *testing.T) {
	config := cfg.BuildConfig(cfg.CliConfig{
		ConfigPath: opt.Set("testdata/test-config.toml"),
	})
	config.Profile.Database = opt.Set("nonexistent_db")
	_, err := Connect(config.Profile)
	if err == nil { t.Errorf("Expected error connecting to nonexistent database, got nil") }
}

func TestConnect_WrongHost(t *testing.T) {
	config := cfg.BuildConfig(cfg.CliConfig{
		ConfigPath: opt.Set("testdata/test-config.toml"),
	})
	config.Profile.Host = opt.Set("nonexistent-host")
	_, err := Connect(config.Profile)
	if err == nil { t.Errorf("Expected error connecting to nonexistent host, got nil") }
}