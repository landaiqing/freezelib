package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("🎨 Theme Showcase Examples")
	fmt.Println("===========================")

	// Create output directory
	os.MkdirAll("output", 0755)

	// Run theme examples
	popularThemesExample()
	lightVsDarkExample()
	themeComparisonExample()
	customThemeExample()

	fmt.Println("\n✅ Theme examples completed!")
	fmt.Println("📁 Check the 'output' directory for generated files.")
	fmt.Println("🎨 Compare different themes and their visual styles.")
}

// Popular themes showcase
func popularThemesExample() {
	fmt.Println("\n🌟 Popular Themes Showcase")
	fmt.Println("--------------------------")

	code := `class DataProcessor:
    def __init__(self, data_source):
        self.data_source = data_source
        self.processed_data = []
    
    def process(self):
        """Process the data with validation and transformation."""
        try:
            raw_data = self.load_data()
            validated_data = self.validate(raw_data)
            self.processed_data = self.transform(validated_data)
            return True
        except Exception as e:
            print(f"Processing failed: {e}")
            return False
    
    def validate(self, data):
        # Remove null values and duplicates
        clean_data = [item for item in data if item is not None]
        return list(set(clean_data))
    
    def transform(self, data):
        # Apply business logic transformations
        return [item.upper() if isinstance(item, str) else item 
                for item in data]`

	// Popular themes to showcase
	themes := []struct {
		name        string
		description string
	}{
		{"github", "GitHub light theme - clean and professional"},
		{"github-dark", "GitHub dark theme - modern and sleek"},
		{"dracula", "Dracula theme - purple and pink accents"},
		{"monokai", "Monokai theme - classic dark with vibrant colors"},
		{"solarized-dark", "Solarized dark - easy on the eyes"},
		{"solarized-light", "Solarized light - warm and readable"},
		{"nord", "Nord theme - arctic, north-bluish color palette"},
		{"one-dark", "One Dark theme - Atom's signature theme"},
		{"material", "Material theme - Google's material design"},
		{"vim", "Vim theme - classic terminal colors"},
	}

	for _, theme := range themes {
		fmt.Printf("🎨 Generating %s theme...\n", theme.name)

		freeze := freezelib.New().
			WithTheme(theme.name).
			WithFont("JetBrains Mono", 14).
			WithWindow(true).
			WithLineNumbers(true).
			WithShadow(15, 0, 8).
			WithPadding(20)

		svgData, err := freeze.GenerateFromCode(code, "python")
		if err != nil {
			fmt.Printf("❌ Error with theme %s: %v\n", theme.name, err)
			continue
		}

		filename := fmt.Sprintf("output/theme_%s.svg", theme.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", filename, err)
			continue
		}

		fmt.Printf("✅ Generated: %s - %s\n", filename, theme.description)
	}
}

// Light vs Dark theme comparison
func lightVsDarkExample() {
	fmt.Println("\n☀️🌙 Light vs Dark Comparison")
	fmt.Println("------------------------------")

	code := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Modern Web App</title>
    <style>
        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 2rem;
        }
        
        .card {
            background: white;
            border-radius: 8px;
            box-shadow: 0 2px 10px rgba(0,0,0,0.1);
            padding: 1.5rem;
            margin-bottom: 1rem;
        }
        
        .btn-primary {
            background: #007bff;
            color: white;
            border: none;
            padding: 0.75rem 1.5rem;
            border-radius: 4px;
            cursor: pointer;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>Welcome to Our App</h1>
        <div class="card">
            <h2>Features</h2>
            <p>This is a modern web application with responsive design.</p>
            <button class="btn-primary">Get Started</button>
        </div>
    </div>
</body>
</html>`

	// Light and dark theme pairs
	themePairs := []struct {
		light string
		dark  string
		name  string
	}{
		{"github", "github-dark", "GitHub"},
		{"solarized-light", "solarized-dark", "Solarized"},
		{"material", "one-dark", "Material vs One Dark"},
	}

	for _, pair := range themePairs {
		fmt.Printf("🔄 Comparing %s themes...\n", pair.name)

		// Light theme
		lightFreeze := freezelib.New().
			WithTheme(pair.light).
			WithFont("SF Mono", 13).
			WithWindow(true).
			WithLineNumbers(true).
			WithShadow(10, 0, 5).
			WithPadding(25)

		lightData, err := lightFreeze.GenerateFromCode(code, "html")
		if err != nil {
			fmt.Printf("❌ Error with light theme: %v\n", err)
			continue
		}

		// Dark theme
		darkFreeze := freezelib.New().
			WithTheme(pair.dark).
			WithFont("SF Mono", 13).
			WithWindow(true).
			WithLineNumbers(true).
			WithShadow(15, 0, 10).
			WithPadding(25)

		darkData, err := darkFreeze.GenerateFromCode(code, "html")
		if err != nil {
			fmt.Printf("❌ Error with dark theme: %v\n", err)
			continue
		}

		// Save files
		lightFile := fmt.Sprintf("output/comparison_%s_light.svg",
			sanitizeFilename(pair.name))
		darkFile := fmt.Sprintf("output/comparison_%s_dark.svg",
			sanitizeFilename(pair.name))

		os.WriteFile(lightFile, lightData, 0644)
		os.WriteFile(darkFile, darkData, 0644)

		fmt.Printf("✅ Generated: %s (light)\n", lightFile)
		fmt.Printf("✅ Generated: %s (dark)\n", darkFile)
	}
}

// Theme comparison grid
func themeComparisonExample() {
	fmt.Println("\n📊 Theme Comparison Grid")
	fmt.Println("------------------------")

	// Short code snippet for comparison
	code := `fn main() {
    let numbers = vec![1, 2, 3, 4, 5];
    
    let doubled: Vec<i32> = numbers
        .iter()
        .map(|x| x * 2)
        .collect();
    
    println!("Original: {:?}", numbers);
    println!("Doubled: {:?}", doubled);
    
    // Pattern matching
    match doubled.len() {
        0 => println!("Empty vector"),
        1..=5 => println!("Small vector"),
        _ => println!("Large vector"),
    }
}`

	// Themes for comparison
	comparisonThemes := []string{
		"github", "github-dark", "dracula", "monokai",
		"nord", "one-dark", "material", "vim",
	}

	for i, theme := range comparisonThemes {
		fmt.Printf("🎯 Creating comparison sample %d: %s\n", i+1, theme)

		freeze := freezelib.New().
			WithTheme(theme).
			WithFont("Fira Code", 12).
			WithWindow(false). // No window for cleaner comparison
			WithLineNumbers(false).
			WithPadding(15).
			WithDimensions(600, 400) // Consistent size

		svgData, err := freeze.GenerateFromCode(code, "rust")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/comparison_grid_%02d_%s.svg", i+1, theme)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		fmt.Printf("✅ Generated: %s\n", filename)
	}
}

// Custom theme example
func customThemeExample() {
	fmt.Println("\n🎨 Custom Theme Example")
	fmt.Println("------------------------")

	code := `package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type Response struct {
    Message string            'json:"message"'
    Data    map[string]interface{} 'json:"data"'
    Status  int               'json:"status"'
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
    response := Response{
        Message: "API is working!",
        Data: map[string]interface{}{
            "timestamp": "2024-01-15T10:30:00Z",
            "version":   "1.0.0",
            "healthy":   true,
        },
        Status: 200,
    }
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

func main() {
    http.HandleFunc("/api/status", apiHandler)
    fmt.Println("API server starting on :8080")
    http.ListenAndServe(":8080", nil)
}`

	// Create custom-styled versions
	customStyles := []struct {
		name   string
		config func() *freezelib.Freeze
		desc   string
	}{
		{
			"corporate",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("Arial", 14).
					WithBackground("#f8f9fa").
					WithWindow(true).
					WithLineNumbers(true).
					WithShadow(5, 2, 3).
					WithBorder(1, 4, "#dee2e6").
					WithPadding(30)
			},
			"Corporate style - clean and professional",
		},
		{
			"cyberpunk",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("dracula").
					WithFont("Courier New", 13).
					WithBackground("#0d1117").
					WithWindow(true).
					WithLineNumbers(true).
					WithShadow(20, 0, 15).
					WithBorder(2, 0, "#ff79c6").
					WithPadding(25)
			},
			"Cyberpunk style - neon and futuristic",
		},
		{
			"minimal",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("system-ui", 13).
					WithBackground("#ffffff").
					WithWindow(false).
					WithLineNumbers(false).
					WithShadow(0, 0, 0).
					WithPadding(20)
			},
			"Minimal style - clean and distraction-free",
		},
	}

	for _, style := range customStyles {
		fmt.Printf("✨ Creating %s style...\n", style.name)

		freeze := style.config()
		svgData, err := freeze.GenerateFromCode(code, "go")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/custom_%s.svg", style.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		fmt.Printf("✅ Generated: %s - %s\n", filename, style.desc)
	}
}

// Helper function to sanitize filenames
func sanitizeFilename(name string) string {
	// Replace spaces and special characters with underscores
	result := ""
	for _, char := range name {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') ||
			(char >= '0' && char <= '9') {
			result += string(char)
		} else {
			result += "_"
		}
	}
	return result
}
