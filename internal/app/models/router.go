package models

type Route struct {
	Name       string      `json:"name"`
	Path       string      `json:"path"`
	Methods    []string    `json:"methods"`
	Roles      []string    `json:"roles"`
	Handler    string      `json:"handler"`
	PublicApi  bool        `json:"isPublic"`
	Prefix     string      `json:"prefix"`
	Key        string      `json:"key"`
	Permission interface{} `json:"permissions"`
}

type GroupRoute struct {
	GroupAPI string  `json:"group"`
	Data     []Route `json:"data"`
}
