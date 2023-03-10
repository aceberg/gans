package models

// Conf - web gui config
type Conf struct {
	DB       string
	Host     string
	Port     string
	Theme    string
	Show     string
	ConfPath string
	YamlPath string
	KeyPath  string
	LogPath  string
	Interval string
	Quit     chan bool
	GrMap    map[string][]string
}

// Host - one host and its groups
type Host struct {
	Name   string
	Groups []string
}

// Repo - git repository
type Repo struct {
	Path  string `yaml:"path"`
	Head  string `yaml:"head"`
	Inv   string `yaml:"inventory"`
	Hosts []Host `yaml:"hosts"`
}

// Play - info about ansible-playbook run
type Play struct {
	ID    int    `db:"ID"`
	Date  string `db:"DATE"`
	Host  string `db:"HOST"`
	File  string `db:"FILE"`
	Head  string `db:"HEAD"`
	Inv   string `db:"INV"`
	Out   string `db:"OUT"`
	Error string `db:"ERROR"`
	Color string `db:"COLOR"`
}

// Key - ssh key
type Key struct {
	ID    int    `db:"ID"`
	Date  string `db:"DATE"`
	Name  string `db:"NAME"`
	File  string `db:"FILE"`
	State string `db:"STATE"`
}

// GuiData - web gui data
type GuiData struct {
	Config Conf
	Icon   string
	Repo   Repo
	Plays  []Play
	Themes []string
	Len    int
	Keys   []Key
	File   string
}
