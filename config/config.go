package config

import (
	"os"

	"r314dash/model"

	"strings"

	"github.com/google/uuid"
	"gopkg.in/yaml.v2"
)

var Main model.Config

func LoadConfig() error {
	file, err := os.Open("config.yaml")
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&Main); err != nil {
		return err
	}

	// Set default values if necessary
	if Main.Settings.SiteName == "" {
		Main.Settings.SiteName = "Relay314"
	}
	if Main.Settings.Title == "" {
		Main.Settings.Title = Main.Settings.SiteName
	}
	if Main.Settings.Path == "" {
		Main.Settings.Path = "/"
	}
	if Main.Settings.StaticPath == "" {
		Main.Settings.StaticPath = Main.Settings.Path + "static/"
	}
	if Main.Settings.StatsPath == "" {
		Main.Settings.StatsPath = Main.Settings.Path + "stats"
	}
	if Main.Settings.Port == 0 {
		Main.Settings.Port = 8080
	}
	if Main.Settings.SignOutUrl == "" {
		Main.Settings.SignOutUrl = "#"
	}
	if Main.Settings.SettingsUrl == "" {
		Main.Settings.SettingsUrl = "#"
	}
	// create ids
	for idx, link := range Main.Links {
		if link.Href == "" {
			Main.Links[idx].Href = "#"
		}
		Main.Links[idx].Id = uuid.New().String()
	}

	for _, metric := range Main.Settings.Stats.Metrics {
		if len(strings.Split(metric.Label, "_")) != 2 {
			panic("Invalid metric label config: " + metric.Label)
		}
	}

	if Main.Settings.Stats.RefreshInterval <= 0 {
		Main.Settings.Stats.RefreshInterval = 5
	}

	return nil
}
