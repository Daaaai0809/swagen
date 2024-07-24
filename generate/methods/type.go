package methods

type PathRoot struct {
	Get    GetPathSchema    `yaml:"get,omitempty"`
	Post   PostPathSchema   `yaml:"post,omitempty"`
	Put    PutPathSchema    `yaml:"put,omitempty"`
	Delete DeletePathSchema `yaml:"delete,omitempty"`
}
