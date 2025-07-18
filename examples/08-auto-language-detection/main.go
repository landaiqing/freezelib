package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("🔍 FreezeLib Auto Language Detection Examples")
	fmt.Println("=============================================")

	// Create output directory
	err := os.MkdirAll("output", 0755)
	if err != nil {
		fmt.Printf("❌ Error creating output directory: %v\n", err)
		return
	}

	// Run examples
	basicAutoDetectionExample()
	languageDetectionAPIExample()
	customLanguageDetectorExample()
	batchAutoDetectionExample()
	languageAnalysisExample()
}

// Basic auto detection example
func basicAutoDetectionExample() {
	fmt.Println("\n🎯 Basic Auto Detection")
	fmt.Println("-----------------------")

	freeze := freezelib.New()

	// Different code samples without specifying language
	codeSamples := []struct {
		name string
		code string
	}{
		{
			"go_example",
			`package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
    
    // Goroutine example
    go func() {
        fmt.Println("Running in goroutine")
    }()
}`,
		},
		{
			"python_example",
			`#!/usr/bin/env python3

def fibonacci(n):
    """Calculate fibonacci number recursively."""
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

if __name__ == "__main__":
    print(f"Fibonacci(10) = {fibonacci(10)}")`,
		},
		{
			"javascript_example",
			`// Modern JavaScript with async/await
async function fetchUserData(userId) {
    try {
        const response = await fetch('/api/users/' + userId);
        const userData = await response.json();

        return {
            ...userData,
            lastUpdated: new Date().toISOString()
        };
    } catch (error) {
        console.error('Failed to fetch user data:', error);
        throw error;
    }
}`,
		},
		{
			"rust_example",
			`use std::collections::HashMap;

fn main() {
    let mut scores = HashMap::new();
    
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    
    for (key, value) in &scores {
        println!("{}: {}", key, value);
    }
}`,
		},
	}

	for _, sample := range codeSamples {
		fmt.Printf("🔍 Processing %s (auto-detecting language)...\n", sample.name)

		// Detect language first
		detectedLang := freeze.DetectLanguage(sample.code)
		fmt.Printf("   Detected language: %s\n", detectedLang)

		// Generate with auto detection
		svgData, err := freeze.GenerateFromCodeAuto(sample.code)
		if err != nil {
			fmt.Printf("❌ Error generating %s: %v\n", sample.name, err)
			continue
		}

		filename := fmt.Sprintf("output/auto_%s.svg", sample.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Language detection API example
func languageDetectionAPIExample() {
	fmt.Println("\n🔧 Language Detection API")
	fmt.Println("-------------------------")

	freeze := freezelib.New()

	// Test different detection methods
	testCodes := []struct {
		name     string
		code     string
		filename string
	}{
		{
			"config_file",
			`{
    "name": "my-app",
    "version": "1.0.0",
    "dependencies": {
        "express": "^4.18.0",
        "lodash": "^4.17.21"
    }
}`,
			"package.json",
		},
		{
			"shell_script",
			`#!/bin/bash

# Deploy script
set -e

echo "Starting deployment..."

if [ ! -d "dist" ]; then
    echo "Building project..."
    npm run build
fi

echo "Deploying to server..."
rsync -av dist/ user@server:/var/www/html/

echo "Deployment complete!"`,
			"deploy.sh",
		},
		{
			"dockerfile",
			`FROM node:18-alpine

WORKDIR /app

COPY package*.json ./
RUN npm ci --only=production

COPY . .

EXPOSE 3000

CMD ["npm", "start"]`,
			"Dockerfile",
		},
	}

	for _, test := range testCodes {
		fmt.Printf("🔍 Analyzing %s...\n", test.name)

		// Test different detection methods
		langFromContent := freeze.DetectLanguage(test.code)
		langFromFilename := freeze.DetectLanguageFromFilename(test.filename)
		langFromBoth := freeze.DetectLanguageFromFile(test.filename, test.code)

		fmt.Printf("   Content-based: %s\n", langFromContent)
		fmt.Printf("   Filename-based: %s\n", langFromFilename)
		fmt.Printf("   Combined: %s\n", langFromBoth)

		// Generate screenshot using the best detection
		svgData, err := freeze.GenerateFromCodeAuto(test.code)
		if err != nil {
			fmt.Printf("❌ Error generating %s: %v\n", test.name, err)
			continue
		}

		filename := fmt.Sprintf("output/detection_%s.svg", test.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Custom language detector example
func customLanguageDetectorExample() {
	fmt.Println("\n⚙️ Custom Language Detector")
	fmt.Println("---------------------------")

	freeze := freezelib.New()

	// Get the language detector and customize it
	detector := freeze.GetLanguageDetector()

	// Add custom mappings
	detector.AddCustomMapping(".myext", "python")
	detector.AddCustomMapping(".config", "json")

	// Test custom mappings
	customTests := []struct {
		filename string
		content  string
	}{
		{
			"script.myext",
			`def custom_function():
    print("This is a custom extension file")
    return True`,
		},
		{
			"app.config",
			`{
    "database": {
        "host": "localhost",
        "port": 5432
    }
}`,
		},
	}

	for _, test := range customTests {
		fmt.Printf("🔍 Testing custom mapping for %s...\n", test.filename)

		detectedLang := freeze.DetectLanguageFromFile(test.filename, test.content)
		fmt.Printf("   Detected language: %s\n", detectedLang)

		svgData, err := freeze.GenerateFromCodeAuto(test.content)
		if err != nil {
			fmt.Printf("❌ Error generating screenshot: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/custom_%s.svg", filepath.Base(test.filename))
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Batch auto detection example
func batchAutoDetectionExample() {
	fmt.Println("\n📦 Batch Auto Detection")
	fmt.Println("-----------------------")

	// Create sample files with different languages
	sampleFiles := map[string]string{
		"hello.go": `package main

import "fmt"

func main() {
    fmt.Println("Hello from Go!")
}`,
		"hello.py": `def main():
    print("Hello from Python!")

if __name__ == "__main__":
    main()`,
		"hello.js": `function main() {
    console.log("Hello from JavaScript!");
}

main();`,
		"hello.rs": `fn main() {
    println!("Hello from Rust!");
}`,
		"style.css": `body {
    font-family: 'Arial', sans-serif;
    background-color: #f0f0f0;
    margin: 0;
    padding: 20px;
}

.container {
    max-width: 800px;
    margin: 0 auto;
    background: white;
    border-radius: 8px;
    box-shadow: 0 2px 10px rgba(0,0,0,0.1);
}`,
	}

	// Create temporary files
	tempDir := "temp_files"
	err := os.MkdirAll(tempDir, 0755)
	if err != nil {
		fmt.Printf("❌ Error creating temp directory: %v\n", err)
		return
	}
	defer os.RemoveAll(tempDir)

	// Write sample files
	for filename, content := range sampleFiles {
		filePath := filepath.Join(tempDir, filename)
		err := os.WriteFile(filePath, []byte(content), 0644)
		if err != nil {
			fmt.Printf("❌ Error writing %s: %v\n", filename, err)
			continue
		}
	}

	freeze := freezelib.New()

	// Process each file with auto detection
	for filename := range sampleFiles {
		filePath := filepath.Join(tempDir, filename)
		fmt.Printf("📄 Processing %s...\n", filename)

		svgData, err := freeze.GenerateFromFile(filePath)
		if err != nil {
			fmt.Printf("❌ Error processing %s: %v\n", filename, err)
			continue
		}

		outputFile := fmt.Sprintf("output/batch_%s.svg", filename)
		err = os.WriteFile(outputFile, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", outputFile, err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", outputFile)
	}
}

// Language analysis example
func languageAnalysisExample() {
	fmt.Println("\n📊 Language Analysis")
	fmt.Println("--------------------")

	freeze := freezelib.New()

	// Get supported languages
	supportedLangs := freeze.GetSupportedLanguages()
	fmt.Printf("📈 Total supported languages: %d\n", len(supportedLangs))

	// Show first 20 languages
	fmt.Println("🔤 First 20 supported languages:")
	for i, lang := range supportedLangs {
		if i >= 20 {
			break
		}
		fmt.Printf("   %d. %s\n", i+1, lang)
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

	// Create a summary file
	summaryContent := fmt.Sprintf(`# FreezeLib Language Support Summary

Total supported languages: %d

## Sample of supported languages:
`, len(supportedLangs))

	for i, lang := range supportedLangs {
		if i >= 50 {
			summaryContent += "... and more\n"
			break
		}
		summaryContent += fmt.Sprintf("- %s\n", lang)
	}

	svgData, err := freeze.GenerateFromCode(summaryContent, "markdown")
	if err != nil {
		fmt.Printf("❌ Error generating summary: %v\n", err)
		return
	}

	err = os.WriteFile("output/language_summary.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving summary: %v\n", err)
		return
	}

	fmt.Printf("✅ Generated language summary: output/language_summary.svg\n")
}
