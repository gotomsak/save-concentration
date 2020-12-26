package main

// SaveConcentration postされてきたdataのbind
type SaveConcentration struct {
	TypeData          string      `json:"type"`
	ID                string      `json:"id"`
	ConcentrationData interface{} `json:"concentration_data"`
}

// GetSaveImagesID postされてきたdataのbind
type GetSaveImagesID struct {
	TypeData string `json:"type"`
}
