package schemas

import "time"

type Comentario struct {
	Id        string    `json:"id,omitempty"`
	Creacion  time.Time `json:"creacion"`
	Contenido string    `json:"contenido,omitempty"`
	Media     Media     `json:"media"`
}
