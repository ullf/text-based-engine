package structure

import (
	"encoding/json"
	"fmt"
	"os"
)

type HeroLog struct {
	Action   int    `json:"action"`
	Function string `json:"function"`
	Location string `json:"location"`
}

type HeroLogInt interface {
	HLog(h *Hero, action int, function string, location string) *HeroLog
	HWrite(data HeroLog, filename string)
	HRead(filename string)
}

func NewHeroLog(action int, function string, location string) *HeroLog {
	log := &HeroLog{
		Action:   action,
		Function: function,
		Location: location,
	}
	return log
}

func (hl *HeroLog) HLog(h *Hero, action int, function string, location string) *HeroLog {
	hl.Action = action
	hl.Function = function
	hl.Location = location
	return hl
}

func (hl *HeroLog) HWrite(data HeroLog, filename string) {
	file, err := os.OpenFile(filename+".json", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}

	// Encode the user object to JSON and write it to the file
	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}

	defer file.Close()
}

func (hl *HeroLog) HRead(filename string) {
	file, err := os.Open(filename + ".json")
	if err != nil {
		panic(err)
	}

	// Decode the JSON data into a user object
	decoder := json.NewDecoder(file)
	var output HeroLog
	if err := decoder.Decode(&output); err != nil {
		panic(err)
	}
	fmt.Println(output)
	defer file.Close()
}
