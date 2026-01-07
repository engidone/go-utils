package clog

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

// Colores predefinidos
var (
	red     = color.New(color.FgRed)
	green   = color.New(color.FgGreen)
	yellow  = color.New(color.FgYellow)
	cyan    = color.New(color.FgCyan)
	magenta = color.New(color.FgMagenta)
	blue    = color.New(color.FgBlue)
	white   = color.New(color.FgWhite)
)

// ===== SIN FORMATO (variadic args...) =====

// Error(args...)
func Error(args ...any) {
	red.Print("❌ ERROR: ")
	red.Print(args...)
	red.Println()
	os.Exit(1)
}

// Fatal(args...)
func Fatal(args ...any) {
	red.Print("‼️ ")
	red.Print(args...)
	red.Println()
	os.Exit(1)
}

// Success(args...)
func Success(args ...any) {
	green.Print("✅ ")
	green.Print(args...)
	green.Println()
}

// Warn(args...)
func Warn(args ...any) {
	yellow.Print("⚠️  ")
	yellow.Print(args...)
	yellow.Println()
}

// Info(args...)
func Info(args ...any) {
	cyan.Print("ℹ️  ")
	cyan.Print(args...)
	cyan.Println()
}

// Debug(args...)
func Debug(args ...any) {
	if os.Getenv("DEBUG") == "true" {
		blue.Print("🐛 ")
		blue.Print(args...)
		blue.Println()
	}
}

// ===== CON FORMATO (printf-style) =====

// Errorf(format, args...)
func Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	red.Print("")
	red.Printf("❌ ERROR: %s", msg)
	red.Println("")
	os.Exit(1)
}

// Errorf(format, args...)
func Fatalf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	red.Print("")
	red.Printf("‼️ Fatal: %s", msg)
	red.Println("")
	os.Exit(1)
}

// Successf(format, args...)
func Successf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	green.Print("")
	green.Printf("✅ ERROR: %s", msg)
	green.Println("")
}

// Warnf(format, args...)
func Warnf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	yellow.Print("")
	yellow.Printf("⚠️ WARN: %s", msg)
	yellow.Println("")
}

// Infof(format, args...)
func Infof(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	cyan.Print("")
	cyan.Printf("ℹ️ INFO: %s", msg)
	cyan.Println("")

}

// Debugf(format, args...)
func Debugf(format string, args ...any) {
	if os.Getenv("DEBUG") == "true" {
		msg := fmt.Sprintf(format, args...)
		magenta.Print("")
		magenta.Print("🐛 DEBUG: ")
		white.Print(msg)
		magenta.Println("")
	}
}
