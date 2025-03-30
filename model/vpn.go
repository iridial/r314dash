package model

type VpnResponse struct {
	PublicIP     string `json:"public_ip"`
	Region       string `json:"region"`
	Country      string `json:"country"`
	City         string `json:"city"`
	Hostname     string `json:"hostname"`
	Location     string `json:"location"`
	Organization string `json:"organization"`
	PostalCode   string `json:"postal_code"`
	Timezone     string `json:"timezone"`
}

type VpnConnection struct {
	Connected bool   `json:"connected"`
	PublicIP  string `json:"public_ip"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	City      string `json:"city"`
}
