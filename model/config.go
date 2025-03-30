package model

import "r314dash/model/shared"

type Settings struct {
	Port                  int    `yaml:"port"`
	Path                  string `yaml:"path"`
	StaticPath            string `yaml:"static_path"`
	StatsPath             string `yaml:"stats_path"`
	HideUserManagement    bool   `yaml:"hide_user_management"`
	UserHeaderKey         string `yaml:"user_header_key"`
	UserEmailHeaderKey    string `yaml:"user_email_header_key"`
	UserFullnameHeaderKey string `yaml:"user_fullname_header_key"`
	GroupsHeaderKey       string `yaml:"groups_header_key"`
	SiteName              string `yaml:"site_name"`
	Title                 string `yaml:"title"`
	SignOutUrl            string `yaml:"sign_out_url"`
	SettingsUrl           string `yaml:"settings_url"`
	Stats                 Stats  `yaml:"stats"`
}

type Stats struct {
	Enabled         bool     `yaml:"enabled"`
	BackendUrl      string   `yaml:"backend_url"`
	RefreshInterval int16    `yaml:"refresh_interval"`
	Metrics         []Metric `yaml:"metrics"`
	Vpn             Vpn      `yaml:"vpn"`
}

type Vpn struct {
	Enabled         bool   `yaml:"enabled"`
	BackendUrl      string `yaml:"backend_url"`
	ApiKey          string `yaml:"api_key"`
	RefreshInterval int16  `yaml:"refresh_interval"`
}

// label format: "net_" for network resources, "cpu_", "mem_"
type Metric struct {
	Label string `yaml:"label"`
	Query string `yaml:"query"`
}

type Config struct {
	Settings Settings      `yaml:"settings"`
	Links    []shared.Link `yaml:"links"`
}
