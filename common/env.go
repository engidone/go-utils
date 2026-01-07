package common

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func LoadEnvFile() error {
	wd, _ := os.Getwd()
    file, err := os.Open(filepath.Join(wd, ".env"))

    if err != nil {
        return err
    }

    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        
        // Skip empty lines and comments
        if line == "" || strings.HasPrefix(line, "#") {
            continue
        }

        // KEY=VALUE
        parts := strings.SplitN(line, "=", 2)
        if len(parts) != 2 {
            continue
        }

        key := strings.TrimSpace(parts[0])
        value := strings.TrimSpace(parts[1])
        
        // Set environment variable
        os.Setenv(key, value)
    }
    
    return scanner.Err()
}