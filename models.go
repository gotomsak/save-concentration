package main

import "go.mongodb.org/mongo-driver/bson/primitive"

// SaveConcentration postされてきたdataのbind
type SaveConcentration struct {
	Type          string             `json:"type"`
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Measurement   string             `json:"measurement"`
	Concentration []interface{}      `json:"concentration"`
}

// GetSaveImagesID postされてきたdataのbind
type GetSaveImagesID struct {
	Type string `json:"type"`
}

// GetID postされてきたdataのbind
type GetID struct {
	Type string `json:"type"`
	// ID       primitive.ObjectID `json:"id" bson:"_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Concentration []interface{} `json:"concentration"`
}

// GetIDSave 保存
type GetIDSave struct {
	Type string             `json:"type"`
	ID   primitive.ObjectID `json:"id" bson:"_id"`
	// ID                string        `json:"id"`
	Measurement   string        `json:"measurement"`
	Concentration []interface{} `json:"concentration"`
}

type getIDRes struct {
	ID primitive.ObjectID `json:"id" bson:"_id"`
}

type getSaveImagesIDRes struct {
	ID uint64 `json:"id"`
}
