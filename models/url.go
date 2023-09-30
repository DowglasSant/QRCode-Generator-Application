package models

import "time"

type Url struct {
	ID      string    `bson:"_id,omitempty"`
	URL     string    `bson:"url"`
	AddTime time.Time `bson:"addTime"`
	Image   []byte    `bson:"image"`
}
