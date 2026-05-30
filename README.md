PROJECT_NAME: crocodile-ancestry-tracker

# crocodile-ancestry-tracker

A Go-based tool that analyzes genetic markers from historical specimens to trace crocodile lineage and migration patterns, inspired by the discovery that Seychelles' lost crocodiles were actually saltwater crocodiles that drifted across the Indian Ocean.

## Description

This project demonstrates how modern DNA analysis techniques can be applied to historical biological specimens to solve long-standing mysteries about species distribution and migration. By processing genetic data from museum collections, it helps researchers understand whether seemingly unique species are actually isolated populations of known species.

The tool implements basic DNA sequence comparison algorithms to identify genetic similarities between specimens and provides a framework for analyzing historical genetic material that might otherwise be inaccessible.

## Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/crocodile-ancestry-tracker.git
cd crocodile-ancestry-tracker

# Install dependencies
go mod tidy

# Build the project
go build -o croc-tracker main.go
```

## Usage

```bash
# Analyze DNA samples from historical specimens
./croc-tracker analyze --specimen-file museum_samples.json

# Compare genetic markers between populations
./croc-tracker compare --population1 seychelles --population2 saltwater

# Generate lineage report
./croc-tracker report --output report.txt
```

### Example Input Format (museum_samples.json)

```json
{
  "specimens": [
    {
      "id": "SEY-001",
      "species": "Crocodylus porosus",
      "genetic_markers": ["ATCGATCG", "GCTAGCTA"],
      "location": "Seychelles",
      "year": 1780
    }
  ]
}
```

### Features

- DNA sequence alignment and comparison
- Historical specimen data processing
- Genetic similarity scoring
- Lineage tracking visualization
- Cross-population analysis capabilities

This tool serves as a foundation for researchers studying historical biodiversity and species migration patterns using preserved genetic material from natural history collections.