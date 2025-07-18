package main

import (
	"fmt"
	"os"

	"github.com/landaiqing/freezelib"
)

func main() {
	fmt.Println("🔧 Advanced Configuration Examples")
	fmt.Println("===================================")

	// Create output directory
	os.MkdirAll("output", 0755)

	// Run advanced examples
	customFontExample()
	advancedLayoutExample()
	performanceOptimizationExample()
	responsiveDesignExample()
	brandingExample()

	fmt.Println("\n✅ Advanced examples completed!")
	fmt.Println("📁 Check the 'output' directory for generated files.")
}

// Custom font example
func customFontExample() {
	fmt.Println("\n🔤 Custom Font Examples")
	fmt.Println("-----------------------")

	code := `interface UserRepository {
  findById(id: string): Promise<User | null>;
  findByEmail(email: string): Promise<User | null>;
  create(user: CreateUserDto): Promise<User>;
  update(id: string, updates: UpdateUserDto): Promise<User>;
  delete(id: string): Promise<void>;
}

class PostgresUserRepository implements UserRepository {
  constructor(private db: Database) {}

  async findById(id: string): Promise<User | null> {
    const result = await this.db.query(
      'SELECT * FROM users WHERE id = $1',
      [id]
    );
    return result.rows[0] || null;
  }

  async create(user: CreateUserDto): Promise<User> {
    const { name, email, role } = user;
    const result = await this.db.query(
      'INSERT INTO users (name, email, role) VALUES ($1, $2, $3) RETURNING *',
      [name, email, role]
    );
    return result.rows[0];
  }
}`

	// Different font configurations
	fontConfigs := []struct {
		name   string
		family string
		size   float64
		desc   string
	}{
		{"monospace_small", "Courier New", 12, "Small monospace font"},
		{"monospace_large", "JetBrains Mono", 16, "Large modern monospace"},
		{"system_font", "system-ui", 14, "System default font"},
		{"serif_font", "Georgia", 14, "Serif font for readability"},
		{"condensed_font", "SF Mono", 13, "Condensed font for more content"},
	}

	for _, config := range fontConfigs {
		fmt.Printf("🔤 Generating %s example...\n", config.name)

		freeze := freezelib.New().
			WithTheme("github-dark").
			WithFont(config.family, config.size).
			WithWindow(true).
			WithLineNumbers(true).
			WithShadow(15, 0, 8).
			WithPadding(25)

		svgData, err := freeze.GenerateFromCode(code, "typescript")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/font_%s.svg", config.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		fmt.Printf("✅ Generated: %s - %s\n", filename, config.desc)
	}
}

// Advanced layout example
func advancedLayoutExample() {
	fmt.Println("\n📐 Advanced Layout Examples")
	fmt.Println("---------------------------")

	code := `from dataclasses import dataclass
from typing import List, Optional, Dict, Any
from datetime import datetime
import asyncio
import aiohttp

@dataclass
class APIResponse:
    status_code: int
    data: Dict[str, Any]
    headers: Dict[str, str]
    timestamp: datetime

class AsyncAPIClient:
    def __init__(self, base_url: str, timeout: int = 30):
        self.base_url = base_url.rstrip('/')
        self.timeout = aiohttp.ClientTimeout(total=timeout)
        self.session: Optional[aiohttp.ClientSession] = None
    
    async def __aenter__(self):
        self.session = aiohttp.ClientSession(timeout=self.timeout)
        return self
    
    async def __aexit__(self, exc_type, exc_val, exc_tb):
        if self.session:
            await self.session.close()
    
    async def get(self, endpoint: str, params: Optional[Dict] = None) -> APIResponse:
        if not self.session:
            raise RuntimeError("Client not initialized. Use async context manager.")
        
        url = f"{self.base_url}/{endpoint.lstrip('/')}"
        
        async with self.session.get(url, params=params) as response:
            data = await response.json()
            return APIResponse(
                status_code=response.status,
                data=data,
                headers=dict(response.headers),
                timestamp=datetime.now()
            )

# Usage example
async def main():
    async with AsyncAPIClient("https://api.example.com") as client:
        response = await client.get("/users", {"page": 1, "limit": 10})
        print(f"Status: {response.status_code}")
        print(f"Data: {response.data}")

if __name__ == "__main__":
    asyncio.run(main())`

	// Different layout configurations
	layouts := []struct {
		name   string
		config func() *freezelib.Freeze
		desc   string
	}{
		{
			"compact",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("SF Mono", 11).
					WithPadding(10).
					WithMargin(5).
					WithWindow(false).
					WithLineNumbers(true).
					WithDimensions(600, 800)
			},
			"Compact layout for maximum content",
		},
		{
			"spacious",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github-dark").
					WithFont("JetBrains Mono", 16).
					WithPadding(40).
					WithMargin(30).
					WithWindow(true).
					WithLineNumbers(true).
					WithShadow(25, 0, 15).
					WithDimensions(1000, 1200)
			},
			"Spacious layout for presentations",
		},
		{
			"mobile_friendly",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("dracula").
					WithFont("Menlo", 13).
					WithPadding(15).
					WithMargin(10).
					WithWindow(false).
					WithLineNumbers(false).
					WithDimensions(400, 600)
			},
			"Mobile-friendly narrow layout",
		},
		{
			"print_optimized",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("Times New Roman", 12).
					WithPadding(20).
					WithMargin(15).
					WithWindow(false).
					WithLineNumbers(true).
					WithShadow(0, 0, 0). // No shadow for print
					WithBackground("#ffffff").
					WithDimensions(800, 1000)
			},
			"Print-optimized layout",
		},
	}

	for _, layout := range layouts {
		fmt.Printf("📐 Creating %s layout...\n", layout.name)

		freeze := layout.config()
		svgData, err := freeze.GenerateFromCode(code, "python")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/layout_%s.svg", layout.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		fmt.Printf("✅ Generated: %s - %s\n", filename, layout.desc)
	}
}

// Performance optimization example
func performanceOptimizationExample() {
	fmt.Println("\n⚡ Performance Optimization")
	fmt.Println("---------------------------")

	// Short code for performance testing
	shortCode := `fn quicksort<T: Ord>(arr: &mut [T]) {
    if arr.len() <= 1 {
        return;
    }
    
    let pivot_index = partition(arr);
    let (left, right) = arr.split_at_mut(pivot_index);
    
    quicksort(left);
    quicksort(&mut right[1..]);
}

fn partition<T: Ord>(arr: &mut [T]) -> usize {
    let pivot_index = arr.len() - 1;
    let mut i = 0;
    
    for j in 0..pivot_index {
        if arr[j] <= arr[pivot_index] {
            arr.swap(i, j);
            i += 1;
        }
    }
    
    arr.swap(i, pivot_index);
    i
}`

	// Performance-optimized configurations
	perfConfigs := []struct {
		name   string
		config func() *freezelib.Freeze
		desc   string
	}{
		{
			"minimal_overhead",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("monospace", 12).
					WithWindow(false).
					WithLineNumbers(false).
					WithShadow(0, 0, 0).
					WithPadding(10).
					WithMargin(0)
			},
			"Minimal processing overhead",
		},
		{
			"optimized_svg",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("system-ui", 13).
					WithWindow(false).
					WithLineNumbers(true).
					WithShadow(0, 0, 0).
					WithPadding(15).
					WithDimensions(600, 400) // Fixed dimensions
			},
			"SVG-optimized configuration",
		},
		{
			"batch_processing",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("monokai").
					WithFont("Courier", 12).
					WithWindow(false).
					WithLineNumbers(false).
					WithPadding(12).
					WithDimensions(500, 300)
			},
			"Optimized for batch processing",
		},
	}

	for _, config := range perfConfigs {
		fmt.Printf("⚡ Testing %s...\n", config.name)

		freeze := config.config()

		// Generate multiple times to test performance
		for i := 0; i < 3; i++ {
			svgData, err := freeze.GenerateFromCode(shortCode, "rust")
			if err != nil {
				fmt.Printf("❌ Error: %v\n", err)
				break
			}

			filename := fmt.Sprintf("output/perf_%s_%d.svg", config.name, i+1)
			err = os.WriteFile(filename, svgData, 0644)
			if err != nil {
				fmt.Printf("❌ Error saving: %v\n", err)
				break
			}
		}

		fmt.Printf("✅ Generated 3 files for %s - %s\n", config.name, config.desc)
	}
}

// Responsive design example
func responsiveDesignExample() {
	fmt.Println("\n📱 Responsive Design")
	fmt.Println("--------------------")

	code := `@media (max-width: 768px) {
  .container {
    padding: 1rem;
    margin: 0;
  }
  
  .grid {
    grid-template-columns: 1fr;
    gap: 1rem;
  }
  
  .card {
    margin-bottom: 1rem;
  }
  
  .navigation {
    flex-direction: column;
  }
  
  .nav-item {
    width: 100%;
    text-align: center;
    padding: 0.75rem;
  }
}

@media (min-width: 769px) and (max-width: 1024px) {
  .container {
    max-width: 750px;
    padding: 2rem;
  }
  
  .grid {
    grid-template-columns: repeat(2, 1fr);
    gap: 1.5rem;
  }
}

@media (min-width: 1025px) {
  .container {
    max-width: 1200px;
    padding: 3rem;
  }
  
  .grid {
    grid-template-columns: repeat(3, 1fr);
    gap: 2rem;
  }
  
  .hero {
    height: 60vh;
    display: flex;
    align-items: center;
    justify-content: center;
  }
}`

	// Different screen size simulations
	screenSizes := []struct {
		name   string
		width  float64
		height float64
		desc   string
	}{
		{"mobile", 375, 600, "Mobile phone size"},
		{"tablet", 768, 800, "Tablet size"},
		{"desktop", 1200, 800, "Desktop size"},
		{"ultrawide", 1600, 900, "Ultrawide monitor"},
	}

	for _, size := range screenSizes {
		fmt.Printf("📱 Creating %s responsive example...\n", size.name)

		freeze := freezelib.New().
			WithTheme("github").
			WithFont("system-ui", 13).
			WithDimensions(size.width, size.height).
			WithWindow(true).
			WithLineNumbers(true).
			WithPadding(20).
			WithShadow(10, 0, 5)

		svgData, err := freeze.GenerateFromCode(code, "css")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/responsive_%s.svg", size.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		fmt.Printf("✅ Generated: %s (%dx%d) - %s\n",
			filename, size.width, size.height, size.desc)
	}
}

// Branding example
func brandingExample() {
	fmt.Println("\n🎨 Branding Examples")
	fmt.Println("--------------------")

	code := `public class BrandService {
    private final Logger logger = LoggerFactory.getLogger(BrandService.class);
    private final BrandRepository brandRepository;
    private final CacheManager cacheManager;
    
    public BrandService(BrandRepository brandRepository, CacheManager cacheManager) {
        this.brandRepository = brandRepository;
        this.cacheManager = cacheManager;
    }
    
    @Cacheable("brands")
    public Brand getBrandById(Long id) {
        logger.info("Fetching brand with id: {}", id);
        
        return brandRepository.findById(id)
            .orElseThrow(() -> new BrandNotFoundException("Brand not found: " + id));
    }
    
    @Transactional
    public Brand createBrand(CreateBrandRequest request) {
        validateBrandRequest(request);
        
        Brand brand = Brand.builder()
            .name(request.getName())
            .description(request.getDescription())
            .logoUrl(request.getLogoUrl())
            .primaryColor(request.getPrimaryColor())
            .secondaryColor(request.getSecondaryColor())
            .createdAt(Instant.now())
            .build();
            
        Brand savedBrand = brandRepository.save(brand);
        cacheManager.evictCache("brands");
        
        logger.info("Created new brand: {}", savedBrand.getName());
        return savedBrand;
    }
    
    private void validateBrandRequest(CreateBrandRequest request) {
        if (StringUtils.isBlank(request.getName())) {
            throw new ValidationException("Brand name is required");
        }
        
        if (brandRepository.existsByName(request.getName())) {
            throw new ValidationException("Brand name already exists");
        }
    }
}`

	// Different brand styles
	brandStyles := []struct {
		name   string
		config func() *freezelib.Freeze
		desc   string
	}{
		{
			"corporate_blue",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github").
					WithFont("Arial", 14).
					WithBackground("#f8f9fa").
					WithWindow(true).
					WithLineNumbers(true).
					WithShadow(8, 2, 4).
					WithBorder(2, 8, "#0066cc").
					WithPadding(30)
			},
			"Corporate blue branding",
		},
		{
			"startup_green",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("github-dark").
					WithFont("Inter", 14).
					WithBackground("#0d1117").
					WithWindow(true).
					WithLineNumbers(true).
					WithShadow(15, 0, 10).
					WithBorder(1, 12, "#00d084").
					WithPadding(25)
			},
			"Startup green branding",
		},
		{
			"creative_purple",
			func() *freezelib.Freeze {
				return freezelib.New().
					WithTheme("dracula").
					WithFont("Poppins", 14).
					WithBackground("#1a1a2e").
					WithWindow(true).
					WithLineNumbers(true).
					WithShadow(20, 0, 15).
					WithBorder(2, 16, "#8b5cf6").
					WithPadding(35)
			},
			"Creative purple branding",
		},
	}

	for _, style := range brandStyles {
		fmt.Printf("🎨 Creating %s style...\n", style.name)

		freeze := style.config()
		svgData, err := freeze.GenerateFromCode(code, "java")
		if err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			continue
		}

		filename := fmt.Sprintf("output/brand_%s.svg", style.name)
		err = os.WriteFile(filename, svgData, 0644)
		if err != nil {
			fmt.Printf("❌ Error saving: %v\n", err)
			continue
		}

		fmt.Printf("✅ Generated: %s - %s\n", filename, style.desc)
	}
}
