package types

// Specimen represents a historical biological specimen with genetic information
type Specimen struct {
	ID          string            `json:"id"`
	Name        string            `json:"name"`
	Location    string            `json:"location"`
	Date        string            `json:"date"`
	GeneticData map[string]string `json:"genetic_data"`
}

// Result represents the outcome of a genetic analysis
type Result struct {
	SpecimenID      string  `json:"specimen_id"`
	SimilarityScore float64 `json:"similarity_score"`
	MatchedSpecies  string  `json:"matched_species"`
	Confidence      float64 `json:"confidence"`
}