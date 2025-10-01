package main

import (
	"os"

	"github.com/Kansuler/configura"
	"github.com/Kansuler/configura/_example/config"
	"github.com/Kansuler/configura/_example/subpackage"
)

func main() {
	// --- Simulate setting environment variables (for example purposes) ---
	// In a real scenario, these would be set in your shell, Dockerfile, K8s manifest, etc.
	os.Setenv(string(config.DATABASE_URL), "postgres://user:pass@host:port/dbname")
	os.Setenv(string(config.PORT), "8080")
	os.Setenv(string(config.API_KEY), "supersecretapikey")
	os.Setenv(string(config.ENABLE_FEATURE_X), "true")
	os.Setenv(string(subpackage.SUBPACKAGE_DEFINED_CONFIG), "some_value")
	// TIMEOUT_SECONDS is not set, so its fallback will be used.

	// --- Initialize Configura ---
	cfg := configura.New()

	// Load environment variables with fallbacks
	configura.Load(cfg, config.DATABASE_URL, "postgres://fallback_user:fallback_pass@localhost:5432/fallback_db")
	configura.Load(cfg, config.PORT, 3000)  // Fallback port 3000
	configura.Load(cfg, config.API_KEY, "") // Fallback empty string if not set
	configura.Load(cfg, config.ENABLE_FEATURE_X, false)
	configura.Load(cfg, config.TIMEOUT_SECONDS, int64(30)) // Fallback 30 seconds
	configura.Load(cfg, subpackage.SUBPACKAGE_DEFINED_CONFIG, "default_value")

	// Set the configuration by yourself
	configura.Write(cfg, map[configura.Variable[int64]]int64{config.TIMEOUT_SECONDS: 25})

	err := subpackage.Initialize(cfg)
	if err != nil {
		panic(err) // Handle error appropriately in your application
	}
}
