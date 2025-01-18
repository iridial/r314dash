package shared

type Link struct {
	Href   string `yaml:"href"`
	Icon   string `yaml:"icon"`
	Title  string `yaml:"title"`
	Id     string `yaml:"id"`
	IsLink bool   `yaml:"is_link"`
}
