package utils

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/knadh/koanf/parsers/json"
	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

func ParseConfFile(path string, target interface{}) (err error) {
	var (
		k      = koanf.New(".")
		ext    = strings.ToLower(filepath.Ext(path)) // Toujours en minuscule
		parser koanf.Parser
	)

	// Détection du parser selon l'extension.
	switch ext {
	case ".json":
		parser = json.Parser()
	case ".yaml", ".yml":
		parser = yaml.Parser()
	case ".toml":
		parser = toml.Parser()
	default:
		return errors.New("unsupported file format: " + ext)
	}

	// Validate path
	if path == "" {
		return errors.New("configuration file path cannot be empty")
	}
	if _, err := os.Stat(path); err != nil {
		return fmt.Errorf("configuration file not accessible: %w", err)
	}

	// Chargement du fichier de configuration.
	if err = k.Load(file.Provider(path), parser); err != nil {
		return fmt.Errorf("failed to load configuration file: %w", err)
	}

	// Unmarshal avec le tag "koanf".
	if err = k.UnmarshalWithConf("", &target, koanf.UnmarshalConf{Tag: "json"}); err != nil {
		return fmt.Errorf("failed to unmarshal configuration: %w", err)
	}

	return nil
}
