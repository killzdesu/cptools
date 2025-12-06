package config

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"gopkg.in/yaml.v3"
)

// LanguageConfig holds configuration for a single programming language
type LanguageConfig struct {
	Name      string `yaml:"name"`
	Extension string `yaml:"extension"`
	Compile   string `yaml:"compile"`
	Run       string `yaml:"run"`
	Template  string `yaml:"template"`
	Compiled  bool   `yaml:"compiled"`
}

// Config holds all language configurations
type Config struct {
	Languages map[string]LanguageConfig `yaml:"languages"`
}

// AppConfig is the global configuration accessible to all commands
var AppConfig *Config

// LoadConfig loads the configuration from embedded defaults and user overrides
func LoadConfig() (*Config, error) {
	// Load embedded default config
	config := &Config{}
	if err := yaml.Unmarshal(defaultConfigYAML, config); err != nil {
		return nil, fmt.Errorf("failed to parse embedded config: %w", err)
	}

	// Try to load user config if it exists
	userConfigPath, err := getUserConfigPath()
	if err != nil {
		// User config path error, just use defaults
		return config, nil
	}

	if _, err := os.Stat(userConfigPath); err == nil {
		// User config exists, merge it
		userConfigData, err := os.ReadFile(userConfigPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read user config: %w", err)
		}

		userConfig := &Config{}
		if err := yaml.Unmarshal(userConfigData, userConfig); err != nil {
			return nil, fmt.Errorf("failed to parse user config: %w", err)
		}

		// Merge user config over defaults
		for lang, langConfig := range userConfig.Languages {
			config.Languages[lang] = langConfig
		}
	}

	return config, nil
}

// GetLanguageConfig returns the configuration for a specific language
func GetLanguageConfig(lang string) (*LanguageConfig, error) {
	if AppConfig == nil {
		return nil, fmt.Errorf("configuration not loaded")
	}

	langConfig, ok := AppConfig.Languages[lang]
	if !ok {
		return nil, fmt.Errorf("language '%s' not found in configuration", lang)
	}

	return &langConfig, nil
}

// GetAvailableLanguages returns a list of all configured languages
func GetAvailableLanguages() []string {
	if AppConfig == nil {
		return []string{}
	}

	langs := make([]string, 0, len(AppConfig.Languages))
	for lang := range AppConfig.Languages {
		langs = append(langs, lang)
	}
	return langs
}

// getUserConfigPath returns the path to the user's config file
func getUserConfigPath() (string, error) {
	var configDir string

	// Get platform-specific config directory
	if runtime.GOOS == "windows" {
		appData := os.Getenv("APPDATA")
		if appData == "" {
			return "", fmt.Errorf("APPDATA environment variable not set")
		}
		configDir = filepath.Join(appData, "cptools")
	} else {
		// Unix-like systems
		home := os.Getenv("HOME")
		if home == "" {
			return "", fmt.Errorf("HOME environment variable not set")
		}
		configDir = filepath.Join(home, ".config", "cptools")
	}

	return filepath.Join(configDir, "languages.yaml"), nil
}

// GetUserConfigPath returns the path where user config should be placed (exported for help messages)
func GetUserConfigPath() string {
	path, err := getUserConfigPath()
	if err != nil {
		return ""
	}
	return path
}

// SubstituteVariables replaces template variables in a command string
func SubstituteVariables(command, filename string) string {
	basename := strings.TrimSuffix(filename, filepath.Ext(filename))
	ext := strings.TrimPrefix(filepath.Ext(filename), ".")

	// Output name depends on platform
	output := basename
	if runtime.GOOS == "windows" && ext != "py" && ext != "js" {
		output = basename + ".exe"
	}

	replacements := map[string]string{
		"{{filename}}": filename,
		"{{basename}}": basename,
		"{{output}}":   output,
		"{{ext}}":      ext,
	}

	result := command
	for placeholder, value := range replacements {
		result = strings.ReplaceAll(result, placeholder, value)
	}

	return result
}
