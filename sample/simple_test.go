package main

import (
	"fmt"
	"github.com/landaiqing/freezelib"
	"os"
)

func simpleTest() {
	fmt.Println("🧪 Simple Test")
	fmt.Println("==============")

	// Create a new freeze instance
	freeze := freezelib.New()

	// Simple Go code to test
	code := `package main

import "fmt"

func main() {
	fmt.Println("Hello from FreezeLib!")
}`

	// Generate SVG
	svgData, err := freeze.GenerateFromCode(code, "go")
	if err != nil {
		fmt.Printf("❌ Error: %v\n", err)
		return
	}

	// Save to file
	err = os.WriteFile("simple_test.svg", svgData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving file: %v\n", err)
		return
	}

	fmt.Printf("✅ Generated simple_test.svg (%d bytes)\n", len(svgData))

	// Test QuickFreeze API
	qf := freezelib.NewQuickFreeze()
	svgData2, err := qf.WithTheme("github").CodeToSVG(code)
	if err != nil {
		fmt.Printf("❌ QuickFreeze Error: %v\n", err)
		return
	}

	err = os.WriteFile("quickfreeze_test.svg", svgData2, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving QuickFreeze file: %v\n", err)
		return
	}

	fmt.Printf("✅ Generated quickfreeze_test.svg (%d bytes)\n", len(svgData2))

	// Test ANSI output
	ansiOutput := "\033[32m✓ SUCCESS\033[0m: Test passed\n\033[31m✗ ERROR\033[0m: Test failed"
	ansiData, err := freeze.GenerateFromANSI(ansiOutput)
	if err != nil {
		fmt.Printf("❌ ANSI Error: %v\n", err)
		return
	}

	err = os.WriteFile("ansi_test.svg", ansiData, 0644)
	if err != nil {
		fmt.Printf("❌ Error saving ANSI file: %v\n", err)
		return
	}

	fmt.Printf("✅ Generated ansi_test.svg (%d bytes)\n", len(ansiData))
	fmt.Println("🎉 All tests passed!")
}

func init() {
	// Run simple test instead of full examples
	if len(os.Args) > 1 && os.Args[1] == "test" {
		simpleTest()
		os.Exit(0)
	}
}
