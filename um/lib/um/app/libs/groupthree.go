package libs

type State struct {
	Opened   string `json:"opened"`
	Selected string `json:"selected"`
}

type GroupTree struct {
	Id     string `json:"id"`
	Parent string `json:"parent"`
	Text   string `json:"text"`
	Aattr  string `json:"a_attr"`
	State  State  `json:"state"`
}
