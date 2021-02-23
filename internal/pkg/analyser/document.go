package analyser

import (
	"github.com/Elfsilon/resumist/internal/pkg/utils"
)

// Document ...
type Document struct {
	ID    int    `json:"id"`
	About string `json:"about"`
}

// NewParsedDocument ...
func NewParsedDocument() ParsedDocument {
	return ParsedDocument{}
}

// ParsedDocument ...
type ParsedDocument struct {
	ID    int
	About [][]string
}

// FromDoc ...
func (a *ParsedDocument) FromDoc(doc Document) {
	sentances := utils.SplitAny(doc.About, ".\n")

	parsedSentances := make([][]string, 0)
	for _, s := range sentances {
		tokens := utils.SplitAny(s, " ")
		parsedSentances = append(parsedSentances, tokens)
	}

	a.ID = doc.ID
	a.About = parsedSentances
}
