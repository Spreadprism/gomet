package configs

type Plugin struct {
	Cmd          string
	Dependencies []string
	Build        string
}

type Configs struct {
	Plugins map[string]Plugin
}
