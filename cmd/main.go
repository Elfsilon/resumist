package main

import (
	"fmt"
	"github.com/Elfsilon/resumist/internal/pkg/analyser"
)

func main() {
	skills := [][]string{
		{"React", "Реакт", "ReactJS"},
		{"Redux", "Редукс", "Рдакс"},
		{"Saga", "Сага", "Sagas", "Саги"},
	}
	config := analyser.NewConfig(skills...)

	a := analyser.New(config)
	a.LoadDocsFromPath("data.json")

	fmt.Println("Config", a.Config)
	for _, d := range a.Docs {
		fmt.Printf("\nDoc #%v:\n", d.ID)
		for _, s := range d.About {
			fmt.Printf("%v\n", s)
		}
	}
}
