package analyser

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// New ...
func New(c *Config) *Analyser {
	return &Analyser{
		Config: c,
	}
}

// Analyser ...
type Analyser struct {
	Config *Config
	Docs   []ParsedDocument
}

// LoadDocsFromPath ...
func (a *Analyser) LoadDocsFromPath(path string) {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var docs []Document
	json.Unmarshal(byteValue, &docs)

	a.LoadDocs(docs...)
}

// LoadDocs ...
func (a *Analyser) LoadDocs(docs ...Document) {
	newDocs := make([]ParsedDocument, 0)
	for _, d := range docs {
		pDoc := NewParsedDocument()
		pDoc.FromDoc(d)
		newDocs = append(newDocs, pDoc)
	}
	a.Docs = newDocs
}

// Start ...
func (a *Analyser) Start() {

}
