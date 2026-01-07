package database

import "strings"

// buildDSN constructs the database connection string from config
func BuildDSN(database DatabaseConfig) string {
	dsn := database.DSN
	dsn = strings.ReplaceAll(dsn, "{engine}", database.Engine)
	dsn = strings.ReplaceAll(dsn, "{user}", database.Username)
	dsn = strings.ReplaceAll(dsn, "{password}", database.Password)
	dsn = strings.ReplaceAll(dsn, "{host}", database.Host)
	dsn = strings.ReplaceAll(dsn, "{port}", database.Port)
	dsn = strings.ReplaceAll(dsn, "{db_name}", database.Password)
	dsn = strings.ReplaceAll(dsn, "{ssl_mode}", database.SSLMode)
	return dsn
}