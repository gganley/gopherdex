package gopherdex

import (
	"encoding/json"
	"io/ioutil"
)

type Pokedex struct {
	ID   int
	Name Name
	Type []string
	Base Base
}

type Name struct {
	English  string
	Japanese string
	Chinese  string
}

type Base struct {
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type Item struct {
	ID   int
	Name Name
}

type Skill struct {
	Accuracy int
	Ename    string
	ID       int
	Power    int
	PP       int
}

func gregUnmarshal(m interface{}, path string) {
	data, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	unmarshalErr := json.Unmarshal(data, &m)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}
}

func GetPokedex() []Pokedex {
	var m []Pokedex
	gregUnmarshal(&m, "json/pokedex.json")
	return m
}

func GetSkills() []Skill {
	var m []Skill
	gregUnmarshal(&m, "json/skills.json")
	return m
}
