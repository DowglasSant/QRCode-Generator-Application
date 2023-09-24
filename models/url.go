package models

import "time"

type Url struct {
	ID           string    `bson:"_id,omitempty"`
	URL          string    `bson:"url"`
	DataDeAdicao time.Time `bson:"dataDeAdicao"`
	Imagem       []byte    `bson:"imagem"`
}
