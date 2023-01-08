package models

// Conf - web gui config
type Conf struct {
	Host  string
	Port  string
	Theme string
}

// Repo - git repository
type Repo struct {
	Name    string `yaml:"name"`
	Path    string `yaml:"path"`
	Timeout string `yaml:"timeout"`
	Hosts	[]string `yaml:"hosts"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Icon   string
	Repos  []Repo
	Themes []string
}
