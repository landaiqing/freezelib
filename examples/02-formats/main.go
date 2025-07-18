package main

import (
	"fmt"
	"github.com/landaiqing/freezelib"
	"os"
)

func main() {
	fmt.Println("📊 Output Format Examples")
	fmt.Println("=========================")

	// Create output directory
	os.MkdirAll("output", 0755)

	// Run format examples
	svgVsPngExample()
	qualityComparisonExample()
	dimensionExamples()
	formatOptimizationExample()

	fmt.Println("\n✅ Format examples completed!")
	fmt.Println("📁 Check the 'output' directory for generated files.")
	fmt.Println("📏 Compare file sizes and visual quality between formats.")
}

// SVG vs PNG comparison
func svgVsPngExample() {
	fmt.Println("\n🆚 SVG vs PNG Comparison")
	fmt.Println("-------------------------")

	freeze := freezelib.New().
		WithTheme("github-dark").
		WithFont("JetBrains Mono", 14).
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(15, 0, 8)

	code := `package main

import (
    "fmt"
    "net/http"
    "log"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}`

	// Generate SVG
	svgData, err := freeze.GenerateFromCode(code, "go")
	if err != nil {
		fmt.Printf("❌ Error generating SVG: %v\n", err)
		return
	}

	// Generate PNG
	pngData, err := freeze.GeneratePNGFromCode(code, "go")
	if err != nil {
		fmt.Printf("❌ Error generating PNG: %v\n", err)
		return
	}

	// Save files
	svgPath := "output/comparison.svg"
	pngPath := "output/comparison.png"

	os.WriteFile(svgPath, svgData, 0644)
	os.WriteFile(pngPath, pngData, 0644)

	// Show file size comparison
	svgInfo, _ := os.Stat(svgPath)
	pngInfo, _ := os.Stat(pngPath)

	fmt.Printf("✅ Generated: %s (%d bytes)\n", svgPath, svgInfo.Size())
	fmt.Printf("✅ Generated: %s (%d bytes)\n", pngPath, pngInfo.Size())
	fmt.Printf("📊 Size ratio: PNG is %.1fx larger than SVG\n",
		float64(pngInfo.Size())/float64(svgInfo.Size()))
}

// Quality comparison with different settings
func qualityComparisonExample() {
	fmt.Println("\n🎨 Quality Comparison")
	fmt.Println("---------------------")

	baseCode := `def quicksort(arr):
    if len(arr) <= 1:
        return arr
    
    pivot = arr[len(arr) // 2]
    left = [x for x in arr if x < pivot]
    middle = [x for x in arr if x == pivot]
    right = [x for x in arr if x > pivot]
    
    return quicksort(left) + middle + quicksort(right)

# Example usage
numbers = [3, 6, 8, 10, 1, 2, 1]
sorted_numbers = quicksort(numbers)
print(f"Original: {numbers}")
print(f"Sorted: {sorted_numbers}")`

	// Different quality settings
	configs := []struct {
		name     string
		width    float64
		height   float64
		fontSize float64
		theme    string
	}{
		{"low_quality", 400, 300, 10, "github"},
		{"medium_quality", 800, 600, 14, "github-dark"},
		{"high_quality", 1200, 900, 16, "dracula"},
		{"ultra_quality", 1600, 1200, 18, "monokai"},
	}

	for _, config := range configs {
		fmt.Printf("🔧 Generating %s...\n", config.name)

		freeze := freezelib.New().
			WithTheme(config.theme).
			WithFont("JetBrains Mono", config.fontSize).
			WithDimensions(config.width, config.height).
			WithWindow(true).
			WithLineNumbers(true).
			WithShadow(10, 0, 5)

		// Generate PNG for quality comparison
		pngData, err := freeze.GeneratePNGFromCode(baseCode, "python")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/quality_%s.png", config.name)
		err = os.WriteFile(filename, pngData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		// Show file info
		info, _ := os.Stat(filename)
		fmt.Printf("✅ Generated: %s (%dx%d, %d bytes)\n",
			filename, config.width, config.height, info.Size())
	}
}

// Different dimension examples
func dimensionExamples() {
	fmt.Println("\n📏 Dimension Examples")
	fmt.Println("---------------------")

	code := `SELECT u.name, u.email, COUNT(o.id) as order_count
FROM users u
LEFT JOIN orders o ON u.id = o.user_id
WHERE u.created_at >= '2024-01-01'
GROUP BY u.id, u.name, u.email
HAVING COUNT(o.id) > 0
ORDER BY order_count DESC
LIMIT 10;`

	dimensions := []struct {
		name   string
		width  float64
		height float64
		desc   string
	}{
		{"square", 600, 600, "Square format"},
		{"wide", 1000, 400, "Wide format (presentations)"},
		{"tall", 400, 800, "Tall format (mobile)"},
		{"standard", 800, 600, "Standard 4:3 ratio"},
		{"widescreen", 1200, 675, "Widescreen 16:9 ratio"},
	}

	for _, dim := range dimensions {
		fmt.Printf("📐 Creating %s format (%fx%f)...\n", dim.name, dim.width, dim.height)

		freeze := freezelib.New().
			WithTheme("nord").
			WithFont("Cascadia Code", 13).
			WithDimensions(dim.width, dim.height).
			WithWindow(true).
			WithLineNumbers(true).
			WithPadding(20)

		svgData, err := freeze.GenerateFromCode(code, "sql")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/dimension_%s.svg", dim.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		fmt.Printf("✅ Generated: %s - %s\n", filename, dim.desc)
	}
}

// Format optimization examples
func formatOptimizationExample() {
	fmt.Println("\n⚡ Format Optimization")
	fmt.Println("----------------------")

	code := `import React, { useState, useEffect } from 'react';

const TodoApp = () => {
    const [todos, setTodos] = useState([]);
    const [input, setInput] = useState('');

    useEffect(() => {
        // Load todos from localStorage
        const saved = localStorage.getItem('todos');
        if (saved) {
            setTodos(JSON.parse(saved));
        }
    }, []);

    const addTodo = () => {
        if (input.trim()) {
            const newTodo = {
                id: Date.now(),
                text: input,
                completed: false
            };
            setTodos([...todos, newTodo]);
            setInput('');
        }
    };

    return (
        <div className="todo-app">
            <h1>Todo List</h1>
            <input 
                value={input}
                onChange={(e) => setInput(e.target.value)}
                placeholder="Add a todo..."
            />
            <button onClick={addTodo}>Add</button>
        </div>
    );
};

export default TodoApp;`

	// Optimized for different use cases
	optimizations := []struct {
		name        string
		format      string
		description string
		config      func() *freezelib.Freeze
	}{
		{
			"web_optimized",
			"svg",
			"Optimized for web (small, scalable)",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("system-ui", 12).
					WithPadding(15).
					WithWindow(false) // No window for smaller size
			},
		},
		{
			"print_optimized",
			"png",
			"Optimized for print (high DPI)",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("Times New Roman", 14).
					WithDimensions(1200, 900).
					WithWindow(true).
					WithLineNumbers(true).
					WithShadow(0, 0, 0) // No shadow for print
			},
		},
		{
			"social_optimized",
			"png",
			"Optimized for social media",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("dracula").
					WithFont("Fira Code", 16).
					WithDimensions(1080, 1080). // Square for Instagram
					WithWindow(true).
					WithShadow(20, 0, 15).
					WithPadding(30)
			},
		},
	}

	for _, opt := range optimizations {
		fmt.Printf("🎯 Creating %s (%s)...\n", opt.name, opt.description)

		freeze := opt.config()

		var data []byte
		var err error
		var filename string

		if opt.format == "svg" {
			data, err = freeze.GenerateFromCode(code, "javascript")
			filename = fmt.Sprintf("output/optimized_%s.svg", opt.name)
		} else {
			data, err = freeze.GeneratePNGFromCode(code, "javascript")
			filename = fmt.Sprintf("output/optimized_%s.png", opt.name)
		}

		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		err = os.WriteFile(filename, data, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		info, _ := os.Stat(filename)
		fmt.Printf("✅ Generated: %s (%d bytes)\n", filename, info.Size())
	}
}
