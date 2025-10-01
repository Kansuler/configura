package config

import "github.com/Kansuler/configura"

// Define your application's configuration variables
const (
	DATABASE_URL     configura.Variable[string] = "DATABASE_URL"
	PORT             configura.Variable[int]    = "PORT"
	API_KEY          configura.Variable[string] = "API_KEY"
	ENABLE_FEATURE_X configura.Variable[bool]   = "ENABLE_FEATURE_X"
	TIMEOUT_SECONDS  configura.Variable[int64]  = "TIMEOUT_SECONDS"
)
