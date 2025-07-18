package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("📦 Batch Processing Examples")
	fmt.Println("=============================")

	// Create output directory
	os.MkdirAll("output", 0755)
	os.MkdirAll("sample_files", 0755)

	// Run batch examples
	createSampleFiles()
	batchFileProcessingExample()
	multiFormatBatchExample()
	concurrentProcessingExample()
	directoryProcessingExample()

	fmt.Println("\n✅ Batch processing examples completed!")
	fmt.Println("📁 Check the 'output' directory for generated files.")
}

// Create sample files for batch processing
func createSampleFiles() {
	fmt.Println("\n📝 Creating Sample Files")
	fmt.Println("------------------------")

	sampleFiles := map[string]string{
		"sample_files/main.go": `package main

import (
	"fmt"
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})
	
	fmt.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}`,

		"sample_files/utils.py": `import json
import logging
from typing import Dict, List, Any, Optional
from datetime import datetime

logger = logging.getLogger(__name__)

def load_config(config_path: str) -> Dict[str, Any]:
    """Load configuration from JSON file."""
    try:
        with open(config_path, 'r') as f:
            return json.load(f)
    except FileNotFoundError:
        logger.error(f"Config file not found: {config_path}")
        return {}
    except json.JSONDecodeError as e:
        logger.error(f"Invalid JSON in config file: {e}")
        return {}

def format_timestamp(timestamp: Optional[datetime] = None) -> str:
    """Format timestamp to ISO string."""
    if timestamp is None:
        timestamp = datetime.now()
    return timestamp.isoformat()

class DataProcessor:
    def __init__(self, config: Dict[str, Any]):
        self.config = config
        self.processed_count = 0
    
    def process_batch(self, items: List[Any]) -> List[Any]:
        """Process a batch of items."""
        results = []
        for item in items:
            processed = self.process_item(item)
            if processed:
                results.append(processed)
                self.processed_count += 1
        return results
    
    def process_item(self, item: Any) -> Optional[Any]:
        """Process a single item."""
        # Implementation depends on item type
        return item`,

		"sample_files/api.js": `const express = require('express');
const cors = require('cors');
const helmet = require('helmet');
const rateLimit = require('express-rate-limit');

const app = express();
const PORT = process.env.PORT || 3000;

// Middleware
app.use(helmet());
app.use(cors());
app.use(express.json({ limit: '10mb' }));
app.use(express.urlencoded({ extended: true }));

// Rate limiting
const limiter = rateLimit({
  windowMs: 15 * 60 * 1000, // 15 minutes
  max: 100, // limit each IP to 100 requests per windowMs
  message: 'Too many requests from this IP'
});
app.use('/api/', limiter);

// Routes
app.get('/api/health', (req, res) => {
  res.json({
    status: 'healthy',
    timestamp: new Date().toISOString(),
    uptime: process.uptime()
  });
});

app.get('/api/users', async (req, res) => {
  try {
    const { page = 1, limit = 10 } = req.query;
    const users = await getUsersPaginated(page, limit);
    
    res.json({
      data: users,
      pagination: {
        page: parseInt(page),
        limit: parseInt(limit),
        total: await getTotalUsers()
      }
    });
  } catch (error) {
    console.error('Error fetching users:', error);
    res.status(500).json({ error: 'Internal server error' });
  }
});

app.listen(PORT, () => {
  console.log('Server running on port \${PORT}\');
});`,

		"sample_files/styles.css": `/* Modern CSS Reset and Base Styles */
*,
*::before,
*::after {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
}

:root {
  --primary-color: #3b82f6;
  --secondary-color: #64748b;
  --success-color: #10b981;
  --warning-color: #f59e0b;
  --error-color: #ef4444;
  --background-color: #ffffff;
  --surface-color: #f8fafc;
  --text-primary: #1e293b;
  --text-secondary: #64748b;
  --border-color: #e2e8f0;
  --shadow: 0 1px 3px 0 rgba(0, 0, 0, 0.1);
}

body {
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  line-height: 1.6;
  color: var(--text-primary);
  background-color: var(--background-color);
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 1rem;
}

.card {
  background: var(--surface-color);
  border: 1px solid var(--border-color);
  border-radius: 8px;
  padding: 1.5rem;
  box-shadow: var(--shadow);
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.card:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px 0 rgba(0, 0, 0, 0.15);
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0.75rem 1.5rem;
  border: none;
  border-radius: 6px;
  font-weight: 500;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.2s ease;
}

.btn-primary {
  background-color: var(--primary-color);
  color: white;
}

.btn-primary:hover {
  background-color: #2563eb;
  transform: translateY(-1px);
}`,

		"sample_files/config.json": `{
  "database": {
    "host": "localhost",
    "port": 5432,
    "name": "myapp",
    "user": "postgres",
    "password": "password",
    "ssl": false,
    "pool": {
      "min": 2,
      "max": 10,
      "idle_timeout": "30s"
    }
  },
  "redis": {
    "host": "localhost",
    "port": 6379,
    "password": "",
    "db": 0,
    "pool_size": 10
  },
  "server": {
    "host": "0.0.0.0",
    "port": 8080,
    "read_timeout": "30s",
    "write_timeout": "30s",
    "idle_timeout": "60s"
  },
  "logging": {
    "level": "info",
    "format": "json",
    "output": "stdout"
  },
  "features": {
    "enable_metrics": true,
    "enable_tracing": true,
    "enable_profiling": false
  }
}`,
	}

	for filename, content := range sampleFiles {
		err := os.WriteFile(filename, []byte(content), 0644)
		if err != nil {
			fmt.Printf("❌ Error creating %s: %v\n", filename, err)
			continue
		}
		fmt.Printf("✅ Created: %s\n", filename)
	}
}

// Batch file processing example
func batchFileProcessingExample() {
	fmt.Println("\n📦 Batch File Processing")
	fmt.Println("------------------------")

	// Get all sample files
	files, err := filepath.Glob("sample_files/*")
	if err != nil {
		fmt.Printf("❌ Error finding files: %v\n", err)
		return
	}

	// Create a consistent freeze instance for all files
	freeze := freezelib.New().
		WithTheme("github-dark").
		WithFont("JetBrains Mono", 14).
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(15, 0, 8).
		WithPadding(20)

	fmt.Printf("🔄 Processing %d files...\n", len(files))

	successCount := 0
	for _, file := range files {
		fmt.Printf("📄 Processing: %s\n", file)

		// Detect language from file extension
		ext := filepath.Ext(file)
		lang := detectLanguage(ext)

		svgData, err := freeze.GenerateFromFile(file)
		if err != nil {
			fmt.Printf("❌ Error processing %s: %v\n", file, err)
			continue
		}

		// Create output filename
		baseName := strings.TrimSuffix(filepath.Base(file), ext)
		outputFile := fmt.Sprintf("output/batch_%s.svg", baseName)

		err = os.WriteFile(outputFile, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", outputFile, err)
			continue
		}

		fmt.Printf("✅ Generated: %s (language: %s)\n", outputFile, lang)
		successCount++
	}

	fmt.Printf("📊 Batch processing completed: %d/%d files successful\n",
		successCount, len(files))
}

// Multi-format batch example
func multiFormatBatchExample() {
	fmt.Println("\n🎨 Multi-format Batch Processing")
	fmt.Println("--------------------------------")

	code := `#include <iostream>
#include <vector>
#include <algorithm>
#include <memory>

template<typename T>
class SmartVector {
private:
    std::unique_ptr<T[]> data;
    size_t size_;
    size_t capacity_;

public:
    SmartVector(size_t initial_capacity = 10) 
        : data(std::make_unique<T[]>(initial_capacity))
        , size_(0)
        , capacity_(initial_capacity) {}

    void push_back(const T& value) {
        if (size_ >= capacity_) {
            resize();
        }
        data[size_++] = value;
    }

    T& operator[](size_t index) {
        if (index >= size_) {
            throw std::out_of_range("Index out of range");
        }
        return data[index];
    }

    size_t size() const { return size_; }
    
    void sort() {
        std::sort(data.get(), data.get() + size_);
    }

private:
    void resize() {
        capacity_ *= 2;
        auto new_data = std::make_unique<T[]>(capacity_);
        std::copy(data.get(), data.get() + size_, new_data.get());
        data = std::move(new_data);
    }
};

int main() {
    SmartVector<int> vec;
    
    for (int i = 0; i < 15; ++i) {
        vec.push_back(rand() % 100);
    }
    
    vec.sort();
    
    std::cout << "Sorted vector: ";
    for (size_t i = 0; i < vec.size(); ++i) {
        std::cout << vec[i] << " ";
    }
    std::cout << std::endl;
    
    return 0;
}`

	// Different format configurations
	formats := []struct {
		name   string
		format string
		theme  string
	}{
		{"svg_light", "svg", "github"},
		{"svg_dark", "svg", "github-dark"},
		{"png_presentation", "png", "dracula"},
		{"png_print", "png", "github"},
	}

	freeze := freezelib.New().
		WithFont("Cascadia Code", 14).
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(15, 0, 8).
		WithPadding(25)

	for _, format := range formats {
		fmt.Printf("🎨 Generating %s format...\n", format.name)

		freeze.WithTheme(format.theme)

		var data []byte
		var err error
		var filename string

		if format.format == "svg" {
			data, err = freeze.GenerateFromCode(code, "cpp")
			filename = fmt.Sprintf("output/multiformat_%s.svg", format.name)
		} else {
			data, err = freeze.GeneratePNGFromCode(code, "cpp")
			filename = fmt.Sprintf("output/multiformat_%s.png", format.name)
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

		// Show file size
		info, _ := os.Stat(filename)
		fmt.Printf("✅ Generated: %s (%d bytes)\n", filename, info.Size())
	}
}

// Concurrent processing example
func concurrentProcessingExample() {
	fmt.Println("\n⚡ Concurrent Processing")
	fmt.Println("-----------------------")

	// Sample code snippets for concurrent processing
	codeSnippets := []struct {
		name string
		code string
		lang string
	}{
		{
			"snippet1",
			`def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

print([fibonacci(i) for i in range(10)])`,
			"python",
		},
		{
			"snippet2",
			`function quickSort(arr) {
    if (arr.length <= 1) return arr;
    
    const pivot = arr[Math.floor(arr.length / 2)];
    const left = arr.filter(x => x < pivot);
    const middle = arr.filter(x => x === pivot);
    const right = arr.filter(x => x > pivot);
    
    return [...quickSort(left), ...middle, ...quickSort(right)];
}`,
			"javascript",
		},
		{
			"snippet3",
			`public class BinarySearch {
    public static int search(int[] arr, int target) {
        int left = 0, right = arr.length - 1;
        
        while (left <= right) {
            int mid = left + (right - left) / 2;
            
            if (arr[mid] == target) return mid;
            if (arr[mid] < target) left = mid + 1;
            else right = mid - 1;
        }
        
        return -1;
    }
}`,
			"java",
		},
		{
			"snippet4",
			`use std::collections::HashMap;

fn word_count(text: &str) -> HashMap<String, usize> {
    text.split_whitespace()
        .map(|word| word.to_lowercase())
        .fold(HashMap::new(), |mut acc, word| {
            *acc.entry(word).or_insert(0) += 1;
            acc
        })
}`,
			"rust",
		},
	}

	// Create freeze instance
	freeze := freezelib.New().
		WithTheme("nord").
		WithFont("JetBrains Mono", 13).
		WithWindow(true).
		WithLineNumbers(true).
		WithShadow(10, 0, 5).
		WithPadding(20)

	// Use goroutines for concurrent processing
	var wg sync.WaitGroup
	results := make(chan string, len(codeSnippets))

	fmt.Printf("🚀 Processing %d snippets concurrently...\n", len(codeSnippets))

	for _, snippet := range codeSnippets {
		wg.Add(1)
		go func(s struct {
			name string
			code string
			lang string
		}) {
			defer wg.Done()

			svgData, err := freeze.GenerateFromCode(s.code, s.lang)
			if err != nil {
				results <- fmt.Sprintf("❌ Error processing %s: %v", s.name, err)
				return
			}

			filename := fmt.Sprintf("output/concurrent_%s.svg", s.name)
			err = os.WriteFile(filename, svgData, 0644)
			if err != nil {
				results <- fmt.Sprintf("❌ Error saving %s: %v", filename, err)
				return
			}

			results <- fmt.Sprintf("✅ Generated: %s", filename)
		}(snippet)
	}

	// Wait for all goroutines to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Collect results
	for result := range results {
		fmt.Println(result)
	}
}

// Directory processing example
func directoryProcessingExample() {
	fmt.Println("\n📁 Directory Processing")
	fmt.Println("-----------------------")

	// Process all files in sample_files directory
	err := filepath.Walk("sample_files", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories
		if info.IsDir() {
			return nil
		}

		// Only process certain file types
		ext := filepath.Ext(path)
		if !isSupportedFile(ext) {
			return nil
		}

		fmt.Printf("📄 Processing directory file: %s\n", path)

		// Create themed freeze instance based on file type
		theme := getThemeForFile(ext)
		freeze := freezelib.New().
			WithTheme(theme).
			WithFont("SF Mono", 13).
			WithWindow(true).
			WithLineNumbers(true).
			WithPadding(20)

		svgData, err := freeze.GenerateFromFile(path)
		if err != nil {
			fmt.Printf("❌ Error processing %s: %v\n", path, err)
			return nil
		}

		// Create output filename
		baseName := strings.TrimSuffix(filepath.Base(path), ext)
		outputFile := fmt.Sprintf("output/directory_%s.svg", baseName)

		err = os.WriteFile(outputFile, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving %s: %v\n", outputFile, err)
			return nil
		}

		fmt.Printf("✅ Generated: %s (theme: %s)\n", outputFile, theme)
		return nil
	})

	if err != nil {
		fmt.Printf("❌ Error walking directory: %v\n", err)
	}
}

// Helper functions
func detectLanguage(ext string) string {
	switch ext {
	case ".go":
		return "go"
	case ".py":
		return "python"
	case ".js":
		return "javascript"
	case ".css":
		return "css"
	case ".json":
		return "json"
	default:
		return "text"
	}
}

func isSupportedFile(ext string) bool {
	supported := []string{".go", ".py", ".js", ".css", ".json", ".md", ".txt"}
	for _, s := range supported {
		if ext == s {
			return true
		}
	}
	return false
}

func getThemeForFile(ext string) string {
	switch ext {
	case ".go":
		return "github-dark"
	case ".py":
		return "monokai"
	case ".js":
		return "dracula"
	case ".css":
		return "github"
	case ".json":
		return "nord"
	default:
		return "github"
	}
}
