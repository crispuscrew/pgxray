package db                                                    
				
import (
	"github.com/crispuscrew/pgxray/src/internal/opt"
	"github.com/crispuscrew/pgxray/src/internal/cfg"

	"os"
	"testing"
	"fmt"
)

func TestMain(m *testing.M) {
	config := cfg.BuildConfig(cfg.CliConfig{                      
		ConfigPath: opt.Set("testdata/test-config.toml"),         
	})

	_, err := Connect(config.Profile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "connect failed: %v\n", err)
		os.Exit(1)
	}

	os.Exit(m.Run())
}