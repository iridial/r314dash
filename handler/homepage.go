package handler

import (
	"html/template"
	"net/http"

	"r314dash/config"
	"r314dash/model"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	username := r.Header.Get(config.Main.Settings.UserHeaderKey)
	useremail := r.Header.Get(config.Main.Settings.UserEmailHeaderKey)
	fullname := r.Header.Get(config.Main.Settings.UserFullnameHeaderKey)
	groups := r.Header.Get(config.Main.Settings.GroupsHeaderKey)

	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	if username == "" {
		username = "Unknown"
	}

	if useremail == "" {
		useremail = "@"
	}

	if groups == "" {
		groups = "-"
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := model.Homepage{
		Links:              config.Main.Links,
		ShowUserManagement: !config.Main.Settings.HideUserManagement,
		Username:           username,
		Email:              useremail,
		Fullname:           fullname,
		Groups:             groups,
		StaticPath:         config.Main.Settings.StaticPath,
		SiteName:           config.Main.Settings.SiteName,
		Title:              config.Main.Settings.Title,
		SignOutUrl:         config.Main.Settings.SignOutUrl,
		SettingsUrl:        config.Main.Settings.SettingsUrl,
		ShowStats:          config.Main.Settings.Stats.Enabled,
		StatsInterval:      config.Main.Settings.Stats.RefreshInterval,
		ShowVpn:            config.Main.Settings.Stats.Vpn.Enabled,
		VpnInterval:        config.Main.Settings.Stats.Vpn.RefreshInterval,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
