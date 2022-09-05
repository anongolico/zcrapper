package schemas

type Hilo struct {
	Id     string `json:"id,omitempty"`
	Titulo string `json:"titulo,omitempty"`
	Media  Media  `json:"media,omitempty"`
}
