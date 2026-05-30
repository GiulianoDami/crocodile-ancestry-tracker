package analysis

import (
	"errors"
	"math"
)

// Specimen represents a genetic specimen with its DNA sequence
type Specimen struct {
	ID    string
	DNA   string
	Name  string
}

// Result represents the analysis result for a specimen
type Result struct {
	SpecimenID      string
	SimilarityScore float64
	Matches         []string
}

// AnalyzeSpecimens performs DNA analysis on a collection of specimens
func AnalyzeSpecimens(specimens []Specimen) ([]Result, error) {
	if len(specimens) == 0 {
		return nil, errors.New("no specimens provided")
	}

	results := make([]Result, 0, len(specimens))

	for _, specimen := range specimens {
		result := Result{
			SpecimenID:      specimen.ID,
			SimilarityScore: 0,
			Matches:         make([]string, 0),
		}

		// Compare against all other specimens
		for _, other := range specimens {
			if specimen.ID == other.ID {
				continue
			}

			score := calculateSimilarity(specimen.DNA, other.DNA)
			if score > 0.8 { // Threshold for considering a match
				result.Matches = append(result.Matches, other.ID)
				result.SimilarityScore = math.Max(result.SimilarityScore, score)
			}
		}

		results = append(results, result)
	}

	return results, nil
}

// calculateSimilarity computes a simple similarity score between two DNA sequences
func calculateSimilarity(dna1, dna2 string) float64 {
	if len(dna1) == 0 || len(dna2) == 0 {
		return 0
	}

	matches := 0
	minLen := len(dna1)
	if len(dna2) < minLen {
		minLen = len(dna2)
	}

	for i := 0; i < minLen; i++ {
		if dna1[i] == dna2[i] {
			matches++
		}
	}

	return float64(matches) / float64(minLen)
}