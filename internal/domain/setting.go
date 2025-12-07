package domain

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

type Setting struct {
	Quality           VideoQuality `json:"quality"`
	ParallelDownloads int          `json:"parallel_downloads"`
	DownloadPath      string       `json:"download_path"`
}

func getSettingsFilePath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Printf("Error getting config dir: %v", err)
		return "byto_settings.json"
	}

	bytoDir := filepath.Join(configDir, "byto")
	if err := os.MkdirAll(bytoDir, 0755); err != nil {
		log.Printf("Error creating config dir: %v", err)
		return "byto_settings.json"
	}

	return filepath.Join(bytoDir, "settings.json")
}

func getDefaultDownloadPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		return "./downloads"
	}
	return filepath.Join(home, "Downloads")
}

func NewSetting() *Setting {
	settings := loadSettings()
	if settings != nil {
		return settings
	}

	return &Setting{
		Quality:           Quality1080p,
		ParallelDownloads: 1,
		DownloadPath:      getDefaultDownloadPath(),
	}
}

func loadSettings() *Setting {
	filePath := getSettingsFilePath()

	data, err := os.ReadFile(filePath)
	if err != nil {
		if !os.IsNotExist(err) {
			log.Printf("Error reading settings file: %v", err)
		}
		return nil
	}

	var settings Setting
	if err := json.Unmarshal(data, &settings); err != nil {
		log.Printf("Error parsing settings file: %v", err)
		return nil
	}

	log.Printf("Loaded settings from %s", filePath)
	return &settings
}

func (s *Setting) Save() error {
	filePath := getSettingsFilePath()

	data, err := json.MarshalIndent(s, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}

	log.Printf("Settings saved to %s", filePath)
	return nil
}

func (s *Setting) Update(quality VideoQuality, parallelDownloads int, downloadPath string) {
	s.Quality = quality
	s.ParallelDownloads = parallelDownloads
	s.DownloadPath = downloadPath
}
