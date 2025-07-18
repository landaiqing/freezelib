package main

import (
	"fmt"
	"github.com/landaiqing/freezelib"
	"os"
)

func main() {
	fmt.Println("🎯 Basic Usage Examples")
	fmt.Println("=======================")

	// Create output directory
	os.MkdirAll("./output", 0755)

	// Run basic examples
	simpleExample()
	helloWorldExample()
	quickStartExample()
	defaultConfigExample()

	fmt.Println("\n✅ Basic examples completed!")
	fmt.Println("📁 Check the 'output' directory for generated files.")
}

// Simple example - minimal code
func simpleExample() {
	fmt.Println("\n📝 Simple Example")
	fmt.Println("------------------")

	freeze := freezelib.New()

	code := `fmt.Println("Hello, FreezeLib!")`

	// Generate SVG
	svgData, err := freeze.GenerateFromCode(code, "go")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/simple.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/simple.svg")
}

// Hello World example - classic first program
func helloWorldExample() {
	fmt.Println("\n👋 Hello World Example")
	fmt.Println("-----------------------")

	freeze := freezelib.New()

	code := `package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("Welcome to FreezeLib!")
    
    // This is a comment
    for i := 1; i <= 3; i++ {
        fmt.Printf("Count: %d\n", i)
    }
}`

	// Generate both SVG and PNG
	svgData, err := freeze.GenerateFromCode(code, "go")
	if err != nil {
		fmt.Printf("❌ Error generating SVG: %v\n", err)
		return
	}

	pngData, err := freeze.GeneratePNGFromCode(code, "go")
	if err != nil {
		fmt.Printf("❌ Error generating PNG: %v\n", err)
		return
	}

	// Save files
	os.WriteFile("output/hello_world.svg", svgData, 0644)
	os.WriteFile("output/hello_world.png", pngData, 0644)

	fmt.Println("✅ Generated: output/hello_world.svg")
	fmt.Println("✅ Generated: output/hello_world.png")
}

// Quick start example - using QuickFreeze API
func quickStartExample() {
	fmt.Println("\n⚡ Quick Start Example")
	fmt.Println("----------------------")

	qf := freezelib.NewQuickFreeze()

	code := `function greet(name) {
    return "Hello, " + name + "!";
}

const message = greet("FreezeLib");
console.log(message);

// Arrow function example
const multiply = (a, b) => a * b;
console.log("5 * 3 =", multiply(5, 3));`

	// Use QuickFreeze with basic styling
	svgData, err := qf.WithTheme("github").
		WithFont("JetBrains Mono", 14).
		WithLineNumbers().
		CodeToSVG(code)

	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/quick_start.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/quick_start.svg")
}

// Default configuration example
func defaultConfigExample() {
	fmt.Println("\n⚙️  Default Configuration Example")
	fmt.Println("---------------------------------")

	// Show what default configuration looks like
	config := freezelib.DefaultConfig()
	freeze := freezelib.NewWithConfig(config)

	code := `# Python Example
def fibonacci(n):
    """Calculate fibonacci number recursively."""
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

# Generate first 10 fibonacci numbers
print("Fibonacci sequence:")
for i in range(10):
    print(f"F({i}) = {fibonacci(i)}")`

	svgData, err := freeze.GenerateFromCode(code, "python")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	err = os.WriteFile("output/default_config.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Println("✅ Generated: output/default_config.svg")

	// Print configuration details
	fmt.Println("\n📋 Default Configuration:")
	fmt.Printf("   Theme: %s\n", config.Theme)
	fmt.Printf("   Font: %s, %dpt\n", config.Font.Family, config.Font.Size)
	fmt.Printf("   Background: %s\n", config.Background)
	fmt.Printf("   Window: %t\n", config.Window)
	fmt.Printf("   Line Numbers: %t\n", config.ShowLineNumbers)
}
