package model

import "r314dash/model/shared"

type Homepage struct {
	Links              []shared.Link
	ShowUserManagement bool
	Username           string
	Email              string
	Fullname           string
	Groups             string
	StaticPath         string
	SiteName           string
	Title              string
	SignOutUrl         string
	SettingsUrl        string
	ShowStats          bool
	StatsInterval      int16
	ShowVpn            bool
	VpnInterval        int16
}
