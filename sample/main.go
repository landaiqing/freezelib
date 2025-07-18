package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("🎨 Freeze Library Examples")
	fmt.Println("========================")

	// Run all examples
	basicExample()
	quickFreezeExample()
	terminalExample()
	customConfigExample()
	fileExample()
	presetExample()
	chainedExample()

	fmt.Println("\n✅ All examples completed successfully!")
	fmt.Println("Check the generated files in the current directory.")
}

func basicExample() {
	fmt.Println("\n📝 Basic Example")
	fmt.Println("----------------")

	// Create a new freeze instance
	freeze := freezelib.New()

	// Go code to screenshot
	code := `package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println("This is a beautiful code screenshot!")
}`

	// Generate SVG
	svgData, err := freeze.GenerateFromCode(code, "go")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Save to file
	err = os.WriteFile("basic_example.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Println("✓ Generated basic_example.svg")
}

func quickFreezeExample() {
	fmt.Println("\n⚡ QuickFreeze Example")
	fmt.Println("---------------------")

	// Use QuickFreeze for simplified API
	qf := freezelib.NewQuickFreeze()

	// JavaScript code with styling
	code := `function fibonacci(n) {
	if (n <= 1) return n;
	return fibonacci(n - 1) + fibonacci(n - 2);
}

console.log('Fibonacci sequence:');
for (let i = 0; i < 10; i++) {
	console.log('F(' + i + ') = ' + fibonacci(i));
}`

	// Chain styling options
	svgData, err := qf.WithTheme("dracula").
		WithFont("Fira Code", 14).
		WithWindow().
		WithShadow().
		WithLineNumbers().
		WithLanguage("javascript").
		CodeToSVG(code)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.WriteFile("quickfreeze_example.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Println("✓ Generated quickfreeze_example.svg")
}

func terminalExample() {
	fmt.Println("\n💻 Terminal Example")
	fmt.Println("-------------------")

	// Use terminal preset for ANSI output
	freeze := freezelib.NewWithConfig(freezelib.TerminalPreset())

	// Colored terminal output
	terminalOutput := "\033[32m✓ SUCCESS\033[0m: Build completed successfully\n" +
		"\033[33m⚠ WARNING\033[0m: Deprecated function used in main.go:42\n" +
		"\033[31m✗ ERROR\033[0m: File not found: config.json\n" +
		"\033[36mINFO\033[0m: Starting server on port 8080\n" +
		"\033[35mDEBUG\033[0m: Loading configuration from ~/.config/app"

	svgData, err := freeze.GenerateFromANSI(terminalOutput)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.WriteFile("terminal_example.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Println("✓ Generated terminal_example.svg")
}

func customConfigExample() {
	fmt.Println("\n⚙️  Custom Config Example")
	fmt.Println("-------------------------")

	// Create custom configuration
	config := freezelib.DefaultConfig()

	// Customize appearance
	config.Theme = "github"
	config.Background = "#f6f8fa"
	config.Font.Family = "JetBrains Mono"
	config.Font.Size = 16
	config.LineHeight = 1.4

	// Layout settings
	config.SetPadding(30)
	config.SetMargin(20)
	config.Width = 800

	// Effects
	config.Window = true
	config.ShowLineNumbers = true
	config.Border.Radius = 12
	config.Border.Width = 2
	config.Border.Color = "#d1d9e0"
	config.Shadow.Blur = 25
	config.Shadow.Y = 15

	// Create freeze instance with custom config
	freeze := freezelib.NewWithConfig(config)

	// Python code
	code := `import numpy as np
import matplotlib.pyplot as plt

def plot_sine_wave():
	x = np.linspace(0, 2 * np.pi, 100)
	y = np.sin(x)
	
	plt.figure(figsize=(10, 6))
	plt.plot(x, y, 'b-', linewidth=2, label='sin(x)')
	plt.xlabel('x')
	plt.ylabel('sin(x)')
	plt.title('Sine Wave')
	plt.grid(True, alpha=0.3)
	plt.legend()
	plt.show()

if __name__ == "__main__":
	plot_sine_wave()`

	svgData, err := freeze.GenerateFromCode(code, "python")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.WriteFile("custom_config_example.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Println("✓ Generated custom_config_example.svg")
}

func fileExample() {
	fmt.Println("\n📁 File Example")
	fmt.Println("---------------")

	// Create a sample Rust file
	sampleCode := `use std::collections::HashMap;

fn main() {
    let mut scores = HashMap::new();
    
    scores.insert(String::from("Blue"), 10);
    scores.insert(String::from("Yellow"), 50);
    
    for (key, value) in &scores {
        println!("{}: {}", key, value);
    }
}`

	// Create sample file
	err := os.WriteFile("sample.rs", []byte(sampleCode), 0644)
	if err != nil {
		fmt.Printf("Error creating sample file: %v\n", err)
		return
	}

	// Use presentation preset
	freeze := freezelib.NewWithConfig(freezelib.PresentationPreset())

	// Generate from file
	svgData, err := freeze.GenerateFromFile("sample.rs")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.WriteFile("file_example.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Println("✓ Generated file_example.svg")
	fmt.Println("✓ Created sample.rs")
}

func presetExample() {
	fmt.Println("\n🎨 Preset Example")
	fmt.Println("-----------------")

	code := `const express = require('express');
const app = express();

app.get('/', (req, res) => {
  res.json({ message: 'Hello, World!' });
});

app.listen(3000, () => {
  console.log('Server running on port 3000');
});`

	// Try different presets
	presets := []string{"dark", "light", "minimal", "retro"}

	for _, preset := range presets {
		freeze := freezelib.NewWithPreset(preset)
		svgData, err := freeze.GenerateFromCode(code, "javascript")
		if err != nil {
			fmt.Printf("Error with preset %s: %v\n", preset, err)
			continue
		}

		filename := fmt.Sprintf("preset_%s_example.svg", preset)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✓ Generated %s\n", filename)
	}
}

func chainedExample() {
	fmt.Println("\n🔗 Chained Methods Example")
	fmt.Println("---------------------------")

	// Create base freeze instance
	freeze := freezelib.New()

	code := `#include <iostream>
#include <vector>
#include <algorithm>

int main() {
    std::vector<int> numbers = {5, 2, 8, 1, 9};
    
    std::sort(numbers.begin(), numbers.end());
    
    std::cout << "Sorted numbers: ";
    for (const auto& num : numbers) {
        std::cout << num << " ";
    }
    std::cout << std::endl;
    
    return 0;
}`

	// Chain multiple styling methods
	svgData, err := freeze.
		WithTheme("monokai").
		WithFont("Cascadia Code", 15).
		WithBackground("#2d2d2d").
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(20, 0, 10).
		WithBorder(1, 10, "#444444").
		WithPadding(25).
		WithMargin(15).
		GenerateFromCode(code, "cpp")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = os.WriteFile("chained_example.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("Error saving file: %v\n", err)
		return
	}

	fmt.Println("✓ Generated chained_example.svg")
}
