package structure

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type HeroLog struct {
	Action   int    `json:"action"`
	Function string `json:"function"`
	Location string `json:"location"`
	QuestId  int    `json:"questId"`
}

type HeroLogs struct {
	Hlogs []HeroLog `json:"hlogs"`
}

type HeroLogInt interface {
	HLog(h *Hero, action int, function string, location string) *HeroLog
	HWrite(data HeroLog, filename string)
	HRead(filename string) HeroLogs
}

type HeroLogsInt interface {
	AppendHeroLog(hl *HeroLog)
}

func NewHeroLog(action int, function string, location string) *HeroLog {
	log := &HeroLog{
		Action:   action,
		Function: function,
		Location: location,
	}
	return log
}

func NewHeroLogs(hl []HeroLog) *HeroLogs {
	log := &HeroLogs{
		Hlogs: hl,
	}
	return log
}

func (hl *HeroLog) HLog(h *Hero, action int, function string, location string) *HeroLog {
	hl.Action = action
	hl.Function = function
	hl.Location = location
	return hl
}

func (hl *HeroLog) HLogQuest(h *Hero, action int, function string, questId int) *HeroLog {
	hl.Action = action
	hl.Function = function
	hl.Location = h.GetLocationAsString()
	hl.QuestId = questId
	return hl
}

func (hls *HeroLogs) AppendHeroLog(hl *HeroLog) {
	if hls.Hlogs != nil {
		hls.Hlogs = append(hls.Hlogs, *hl)
	} else {
		hls.Hlogs = make([]HeroLog, 0)
		hls.Hlogs = append(hls.Hlogs, *hl)
	}
	fmt.Println(hls)

}

func (hl *HeroLog) HWrite(data HeroLogs, filename string) {
	file, err := os.OpenFile(filename+".json", os.O_CREATE|os.O_RDWR, 0644)
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

func (hl *HeroLog) HRead(filename string) HeroLogs {
	file, err := os.Open(filename + ".json")
	var output HeroLogs
	if err != io.EOF {
		// Decode the JSON data into a user object
		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&output); err != nil && err != io.EOF {
			panic(err)
		}

		//fmt.Println("decode: ", output)
		defer file.Close()
		return output
	}
	return output
}
