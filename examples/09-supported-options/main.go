package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("📋 FreezeLib Supported Options Examples")
	fmt.Println("=======================================")

	// Create output directory
	err := os.MkdirAll("output", 0755)
	if err != nil {
		fmt.Printf("❌ Error creating output directory: %v\n", err)
		return
	}

	// Run examples
	languageListExample()
	themeListExample()
	presetListExample()
	generateExamples()
}

// Language list example
func languageListExample() {
	fmt.Println("\n💻 Supported Languages")
	fmt.Println("---------------------")

	freeze := freezelib.New()

	// Get all supported languages
	allLanguages := freeze.GetSupportedLanguages()
	fmt.Printf("📈 Total supported languages: %d\n", len(allLanguages))

	// Show first 20 languages as examples
	fmt.Println("\n📋 First 20 supported languages:")
	for i, lang := range allLanguages[:min(20, len(allLanguages))] {
		fmt.Printf("   %2d. %s\n", i+1, lang)
	}

	// Test language support
	testLanguages := []string{"go", "python", "javascript", "rust", "unknown-lang"}
	fmt.Println("\n🧪 Testing language support:")
	for _, lang := range testLanguages {
		supported := freeze.IsLanguageSupported(lang)
		status := "❌"
		if supported {
			status = "✅"
		}
		fmt.Printf("   %s %s\n", status, lang)
	}
}

// Theme list example
func themeListExample() {
	fmt.Println("\n🎨 Supported Themes")
	fmt.Println("------------------")

	freeze := freezelib.New()

	// Get all supported themes
	allThemes := freeze.GetSupportedThemes()
	fmt.Printf("📈 Total supported themes: %d\n", len(allThemes))

	// Show first 20 themes as examples
	fmt.Println("\n📋 First 20 supported themes:")
	for i, theme := range allThemes[:min(20, len(allThemes))] {
		fmt.Printf("   %2d. %s\n", i+1, theme)
	}

	// Test theme support
	testThemes := []string{"github", "dracula", "monokai", "unknown-theme"}
	fmt.Println("\n🧪 Testing theme support:")
	for _, theme := range testThemes {
		supported := freeze.IsThemeSupported(theme)
		status := "❌"
		if supported {
			status = "✅"
		}
		fmt.Printf("   %s %s\n", status, theme)
	}
}

// Preset list example
func presetListExample() {
	fmt.Println("\n⚙️ Available Presets")
	fmt.Println("-------------------")

	freeze := freezelib.New()

	// Get all available presets
	presets := freeze.GetAvailablePresets()
	fmt.Printf("📈 Total available presets: %d\n", len(presets))

	// Show all presets
	fmt.Println("\n📋 Available presets:")
	for i, preset := range presets {
		fmt.Printf("   %2d. %s\n", i+1, preset)
	}

	// Test preset validity
	testPresets := []string{"base", "full", "terminal", "dark", "light", "unknown-preset"}
	fmt.Println("\n🧪 Testing preset validity:")
	for _, preset := range testPresets {
		valid := freezelib.IsValidPreset(preset)
		status := "❌"
		if valid {
			status = "✅"
		}
		fmt.Printf("   %s %s\n", status, preset)
	}
}

// Generate examples with different options
func generateExamples() {
	fmt.Println("\n🎨 Generating Examples")
	fmt.Println("---------------------")

	freeze := freezelib.New()

	// Sample code for examples
	sampleCode := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")

    // This is a comment
    for i := 0; i < 10; i++ {
        fmt.Printf("Count: %d\n", i)
    }
}`

	// Generate examples with different themes
	themes := []string{"github", "github-dark", "dracula", "monokai"}
	for _, theme := range themes {
		if !freeze.IsThemeSupported(theme) {
			continue
		}

		fmt.Printf("🎨 Generating example with %s theme...\n", theme)

		svgData, err := freeze.WithTheme(theme).
			WithFont("JetBrains Mono", 14).
			WithWindow(true).
			WithLineNumbers(true).
			WithShadow(15, 0, 8).
			WithPadding(20).
			GenerateFromCode(sampleCode, "go")

		if err != nil {
			fmt.Printf("❌ Error with theme %s: %v\n", theme, err)
			continue
		}

		filename := fmt.Sprintf("output/theme_%s.svg", theme)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}

	// Generate examples with different presets
	presets := []string{"base", "full", "terminal", "dark", "light"}
	bashCode := `#!/bin/bash

echo "Starting deployment..."

if [ ! -d "dist" ]; then
    echo "Building project..."
    npm run build
fi

echo "Deploying to server..."
rsync -av dist/ user@server:/var/www/html/

echo "Deployment complete!"`

	for _, preset := range presets {
		if !freezelib.IsValidPreset(preset) {
			continue
		}

		fmt.Printf("⚙️ Generating example with %s preset...\n", preset)

		presetFreeze := freezelib.NewWithPreset(preset)
		svgData, err := presetFreeze.GenerateFromCode(bashCode, "bash")
		if err != nil {
			fmt.Printf("❌ Error with preset %s: %v\n", preset, err)
			continue
		}

		filename := fmt.Sprintf("output/preset_%s.svg", preset)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Helper function
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
