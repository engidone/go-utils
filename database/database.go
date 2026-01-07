package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/engidone/go-utils/common"
	"github.com/engidone/go-utils/log"

	_ "github.com/lib/pq"
)

// DatabaseConfig represents the database configuration structure
type DatabaseConfig struct {
	DSN      string `yaml:"dsn"`
	Engine   string `yaml:"engine"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	SSLMode  string `yaml:"ssl_mode"`
	DBName   string `yaml:"db_name"`
}

// openDb creates and returns a database connection using the app configuration
func Open() (*sql.DB, error) {
	common.LoadEnvFile()
	appConfig, err := common.LoadFile[struct {
		Logging string `yaml:"Logging"`
		Database DatabaseConfig `yaml:"database"`
	}](common.GetConfigPath("app.yaml"))

	os.Setenv("LOGGING", appConfig.Logging)

	if err != nil {
		log.Fatalf("Error loading config: %v", err.Error())
	}

	database := appConfig.Database
	dsn := BuildDSN(database)

	db, err := sql.Open(database.Engine, dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	return db, nil
}

// TruncateTables truncates the specified database tables
func TruncateTables(db *sql.DB, tableNames []string) {
	log.Info("🧹 Starting table truncation...")

	for _, tableName := range tableNames {
		truncateSingleTable(db, tableName)
	}

	restoreForeignKeyConstraints(db)
	log.Success("🎉 Table truncation completed")
}

// truncateSingleTable truncates a single table and shows count
func truncateSingleTable(db *sql.DB, tableName string) {
	// Disable foreign key constraints temporarily
	disableFKQuery := "SET session_replication_role = replica;"
	_, err := db.Exec(disableFKQuery)
	if err != nil {
		log.Fatalf("Error desactivando restricciones de clave externa: %v", err)
	}

	// Count records before truncating
	showTableRecordCount(db, tableName)

	// Truncate table
	truncateQuery := fmt.Sprintf("TRUNCATE TABLE %s RESTART IDENTITY CASCADE", tableName)
	_, err = db.Exec(truncateQuery)
	if err != nil {
		log.Fatalf("Error truncando tabla %s: %v", tableName, err)
	}

	log.Successf("✅ Table %s truncated successfully\n", tableName)
}

// showTableRecordCount displays the number of records in a table
func showTableRecordCount(db *sql.DB, tableName string) {
	var count int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)
	err := db.QueryRow(countQuery).Scan(&count)
	if err != nil {
		log.Errorf("⚠️  Could not count records in table %s: %v\n", tableName, err)
	} else {
		log.Infof("📊 Table %s: %d records found\n", tableName, count)
	}
}

// restoreForeignKeyConstraints restores normal foreign key constraint checking
func restoreForeignKeyConstraints(db *sql.DB) {
	enableFKQuery := "SET session_replication_role = DEFAULT;"
	_, err := db.Exec(enableFKQuery)
	if err != nil {
		log.Fatalf("Error restaurando restricciones de clave externa: %v", err)
	}
}
