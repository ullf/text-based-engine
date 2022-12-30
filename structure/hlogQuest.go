package structure

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type QuestLog struct {
	Action   int    `json:"action"`
	Function string `json:"function"`
	Location string `json:"location"`
	QuestId  int    `json:"questId"`
}

type QuestLogs struct {
	Hlogs []QuestLog `json:"hlogs"`
}

type HeroLogInt interface {
	HLog(h *Hero, action int, function string, location string) *QuestLog
	HWrite(data QuestLog, filename string) QuestLogs
	HRead(filename string) QuestLogs
}

type HeroLogsInt interface {
	AppendHeroLog(hl *QuestLog)
}

func NewHeroLog(action int, function string, location string) *QuestLog {
	log := &QuestLog{
		Action:   action,
		Function: function,
		Location: location,
	}
	return log
}

func NewHeroLogs(hl []QuestLog) *QuestLogs {
	log := &QuestLogs{
		Hlogs: hl,
	}
	return log
}

/*func (hl *HeroLog) HLog(h *Hero, action int, function string, location string) *HeroLog {
	hl.Action = action
	hl.Function = function
	hl.Location = location
	return hl
}*/

func (hl *QuestLog) HLogQuest(h *Hero, action int, function string, questId int) *QuestLog {
	hl.Action = action
	hl.Function = function
	hl.Location = h.GetLocationAsString()
	hl.QuestId = questId
	return hl
}

func (hls *QuestLogs) AppendHeroLog(hl *QuestLog) {
	if hls.Hlogs != nil {
		hls.Hlogs = append(hls.Hlogs, *hl)
	} else {
		hls.Hlogs = make([]QuestLog, 0)
		hls.Hlogs = append(hls.Hlogs, *hl)
	}
	fmt.Println(hls)

}

func (hl *QuestLogs) HWrite(data QuestLogs, filename string) {
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

func (hl *QuestLogs) HRead(filename string) QuestLogs {
	file, _ := os.Open(filename)
	var output QuestLogs
	//if err != io.EOF {
	// Decode the JSON data into a user object
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&output); err != nil && err != io.EOF {
		panic(err)
	}

	fmt.Println("decode: ", output)
	defer file.Close()
	fmt.Println("decode: ", output)
	return output
	//}
	//return output
}
