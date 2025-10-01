package subpackage

import (
	"fmt"

	"github.com/Kansuler/configura"
	"github.com/Kansuler/configura/_example/config"
)

const (
	SUBPACKAGE_DEFINED_CONFIG configura.Variable[string] = "SUBPACKAGE_DEFINED_CONFIG"
)

// RequiredUserServiceKeys lists the configuration variables this service needs.
var RequiredUserServiceKeys = []any{
	SUBPACKAGE_DEFINED_CONFIG,
	config.DATABASE_URL,
	config.API_KEY,
}

// Initialize sets up the user service with the given configuration.
// It validates that all required configuration keys are registered.
func Initialize(cfg configura.Config) error {
	// Validate that the config instance has all the keys our service needs
	if err := cfg.Exists(RequiredUserServiceKeys...); err != nil {
		return fmt.Errorf("user service configuration validation failed: %w", err)
	}

	// Access the validated configuration
	dbURL := cfg.String(config.DATABASE_URL)
	apiKey := cfg.String(config.API_KEY)
	definedConfig := cfg.String(SUBPACKAGE_DEFINED_CONFIG)

	fmt.Printf("UserService: Initializing with DB URL: %s and API Key (present: %t), and has subpackage defined key (present: %s)\n", dbURL, apiKey != "", definedConfig)
	// ... further initialization logic for the user service ...

	return nil
}
