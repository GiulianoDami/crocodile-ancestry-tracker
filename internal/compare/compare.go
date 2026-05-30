package compare

import (
	"errors"
)

// Specimen represents a genetic specimen with its markers
type Specimen struct {
	Name    string
	Markers []byte
}

// ComparePopulations compares genetic markers between two populations
// and returns the similarity ratio as a float64 between 0 and 1
func ComparePopulations(pop1, pop2 []Specimen) (float64, error) {
	if len(pop1) == 0 || len(pop2) == 0 {
		return 0, errors.New("populations cannot be empty")
	}

	totalComparisons := 0
	matchingMarkers := 0

	for _, spec1 := range pop1 {
		for _, spec2 := range pop2 {
			if len(spec1.Markers) != len(spec2.Markers) {
				continue
			}
			
			totalComparisons++
			
			for i, marker1 := range spec1.Markers {
				if marker1 == spec2.Markers[i] {
					matchingMarkers++
				}
			}
		}
	}

	if totalComparisons == 0 {
		return 0, errors.New("no valid comparisons could be made")
	}

	similarity := float64(matchingMarkers) / (float64(totalComparisons) * float64(len(pop1[0].Markers)))
	return similarity, nil
}