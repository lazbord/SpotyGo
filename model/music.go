package model

type Music struct {
	ID       string `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string `bson:"name" json:"name"`
	Artist   string `bson:"artist" json:"artist"`
	Duration string `bson:"duration" json:"duration"`
	Data     []byte `bson:"data" json:"data"`
	Filename string `bson:"filename" json:"filename"`
}
