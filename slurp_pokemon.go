package gopherdex

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Pokedex struct {
	ID int
	Name Name
	Type []string
	Base Base
}

type Name struct {
	English string
	Japanese string
	Chinese string
}

type Base struct {
	HP int
	Attack int
	Defense int
	SpAttack int
	SpDefense int
	Speed int
}


func getPokedex() []Pokedex {
	var m []Pokedex
	data, err := ioutil.ReadFile("pokemon.json/pokedex.json")

	if err != nil {
		panic(err)
	}

    unmarshalErr := json.Unmarshal(data, &m)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	return(m)
}
