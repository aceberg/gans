package models

// Conf - web gui config
type Conf struct {
	DB		string
	Host     string
	Port     string
	Theme    string
	YamlPath string
	Timeout  string
	Quit     chan bool
}

// Repo - git repository
type Repo struct {
	Name  string   `yaml:"name"`
	Path  string   `yaml:"path"`
	Head  string   `yaml:"head"`
	Inv   string   `yaml:"inventory"`
	Hosts []string `yaml:"hosts"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Icon   string
	Repos  []Repo
	Themes []string
}
