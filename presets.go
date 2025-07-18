package freezelib

// BasePreset returns a basic configuration for simple code screenshots
func BasePreset() *Config {
	config := DefaultConfig()
	config.Background = "#171717"
	config.SetPadding(20)
	config.SetMargin(0)
	config.Window = false
	config.Border = Border{Radius: 0, Width: 0, Color: "#515151"}
	config.Shadow = Shadow{Blur: 0, X: 0, Y: 0}
	config.ShowLineNumbers = false
	return config
}

// FullPreset returns a macOS-like configuration with window controls and shadow
func FullPreset() *Config {
	config := DefaultConfig()
	config.Background = "#282c34"
	config.SetPadding(20, 40, 20, 20)
	config.SetMargin(20)
	config.Window = true
	config.Border = Border{Radius: 8, Width: 0, Color: "#515151"}
	config.Shadow = Shadow{Blur: 20, X: 0, Y: 10}
	config.ShowLineNumbers = false
	config.Theme = "github-dark"
	return config
}

// TerminalPreset returns a configuration optimized for terminal output
func TerminalPreset() *Config {
	config := DefaultConfig()
	config.Background = "#0d1117"
	config.SetPadding(15)
	config.SetMargin(10)
	config.Window = false
	config.Border = Border{Radius: 6, Width: 1, Color: "#30363d"}
	config.Shadow = Shadow{Blur: 15, X: 0, Y: 5}
	config.ShowLineNumbers = false
	config.Theme = "github-dark"
	config.Font.Family = "JetBrains Mono"
	config.Font.Size = 13
	return config
}

// PresentationPreset returns a configuration suitable for presentations
func PresentationPreset() *Config {
	config := DefaultConfig()
	config.Background = "#ffffff"
	config.SetPadding(40)
	config.SetMargin(30)
	config.Window = true
	config.Border = Border{Radius: 12, Width: 2, Color: "#e1e4e8"}
	config.Shadow = Shadow{Blur: 30, X: 0, Y: 15}
	config.ShowLineNumbers = true
	config.Theme = "github"
	config.Font.Size = 16
	config.LineHeight = 1.4
	return config
}

// MinimalPreset returns a minimal configuration with no decorations
func MinimalPreset() *Config {
	config := DefaultConfig()
	config.Background = "#ffffff"
	config.SetPadding(10)
	config.SetMargin(0)
	config.Window = false
	config.Border = Border{Radius: 0, Width: 0, Color: ""}
	config.Shadow = Shadow{Blur: 0, X: 0, Y: 0}
	config.ShowLineNumbers = false
	config.Theme = "github"
	return config
}

// DarkPreset returns a dark theme configuration
func DarkPreset() *Config {
	config := DefaultConfig()
	config.Background = "#1e1e1e"
	config.SetPadding(25)
	config.SetMargin(15)
	config.Window = false
	config.Border = Border{Radius: 8, Width: 1, Color: "#3c3c3c"}
	config.Shadow = Shadow{Blur: 20, X: 0, Y: 8}
	config.ShowLineNumbers = false
	config.Theme = "dracula"
	config.Font.Family = "Fira Code"
	config.Font.Size = 14
	config.Font.Ligatures = true
	return config
}

// LightPreset returns a light theme configuration
func LightPreset() *Config {
	config := DefaultConfig()
	config.Background = "#fafbfc"
	config.SetPadding(25)
	config.SetMargin(15)
	config.Window = false
	config.Border = Border{Radius: 8, Width: 1, Color: "#d1d5da"}
	config.Shadow = Shadow{Blur: 20, X: 0, Y: 8}
	config.ShowLineNumbers = false
	config.Theme = "github"
	config.Font.Family = "SF Mono"
	config.Font.Size = 14
	return config
}

// RetroPreset returns a retro terminal-style configuration
func RetroPreset() *Config {
	config := DefaultConfig()
	config.Background = "#000000"
	config.SetPadding(20)
	config.SetMargin(10)
	config.Window = false
	config.Border = Border{Radius: 0, Width: 2, Color: "#00ff00"}
	config.Shadow = Shadow{Blur: 0, X: 0, Y: 0}
	config.ShowLineNumbers = false
	config.Theme = "monokai"
	config.Font.Family = "Courier New"
	config.Font.Size = 12
	config.Font.Ligatures = false
	return config
}

// NeonPreset returns a neon-style configuration
func NeonPreset() *Config {
	config := DefaultConfig()
	config.Background = "#0a0a0a"
	config.SetPadding(30)
	config.SetMargin(20)
	config.Window = false
	config.Border = Border{Radius: 10, Width: 2, Color: "#ff00ff"}
	config.Shadow = Shadow{Blur: 25, X: 0, Y: 0}
	config.ShowLineNumbers = false
	config.Theme = "vim"
	config.Font.Family = "Fira Code"
	config.Font.Size = 14
	config.Font.Ligatures = true
	return config
}

// CompactPreset returns a compact configuration for small code snippets
func CompactPreset() *Config {
	config := DefaultConfig()
	config.Background = "#f6f8fa"
	config.SetPadding(10)
	config.SetMargin(5)
	config.Window = false
	config.Border = Border{Radius: 4, Width: 1, Color: "#d0d7de"}
	config.Shadow = Shadow{Blur: 5, X: 0, Y: 2}
	config.ShowLineNumbers = false
	config.Theme = "github"
	config.Font.Size = 12
	config.LineHeight = 1.1
	return config
}

// PresetMap contains all available presets
var PresetMap = map[string]func() *Config{
	"base":         BasePreset,
	"full":         FullPreset,
	"terminal":     TerminalPreset,
	"presentation": PresentationPreset,
	"minimal":      MinimalPreset,
	"dark":         DarkPreset,
	"light":        LightPreset,
	"retro":        RetroPreset,
	"neon":         NeonPreset,
	"compact":      CompactPreset,
}

// GetPreset returns a preset configuration by name
func GetPreset(name string) *Config {
	if preset, exists := PresetMap[name]; exists {
		return preset()
	}
	return DefaultConfig()
}

// ListPresets returns a list of available preset names
func ListPresets() []string {
	presets := make([]string, 0, len(PresetMap))
	for name := range PresetMap {
		presets = append(presets, name)
	}
	return presets
}

// IsValidPreset checks if a preset name is valid
func IsValidPreset(name string) bool {
	_, exists := PresetMap[name]
	return exists
}
