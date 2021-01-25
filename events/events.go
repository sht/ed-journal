package events

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"

	"github.com/sht/ed-journal/dispatcher"
)

// Event represents an incoming event from the journal
type Event struct {
	Event     string    `json:"event"`
	Timestamp time.Time `json:"timestamp"`
}

func debug(b1 []byte, e interface{}) {
	b2, err := json.Marshal(e)
	if err != nil {
		fmt.Println(err)
		return
	}

	b1map := make(map[string]interface{})
	err = json.Unmarshal(b1, &b1map)
	if err != nil {
		fmt.Println(err)
		return
	}

	b2map := make(map[string]interface{})
	err = json.Unmarshal(b2, &b2map)
	if err != nil {
		fmt.Println(err)
		return
	}

	b1, err = json.Marshal(b1map)
	if err != nil {
		fmt.Println(err)
		return
	}

	b2, err = json.Marshal(b2map)
	if err != nil {
		fmt.Println(err)
		return
	}

	if bytes.Compare(b1, b2) != 0 {
		fmt.Printf("\n invalid event struct definition\n\noriginal:\n%s\n\nparsed:\n%s\n\n", b1, b2)
		return
	}

	//fmt.Printf("received %s event\n", b2map["event"])
}

func AddListeners(d *dispatcher.Dispatcher) {
	// startup
	{
		d.On(Cargo, CargoEventHandler)
		d.On(ClearSavedGame, ClearSavedGameEventHandler)
		d.On(Commander, CommanderEventHandler)
		d.On(Loadout, LoadoutEventHandler)
		d.On(Materials, MaterialsEventHandler)
		d.On(Missions, MissionsEventHandler)
		d.On(NewCommander, NewCommanderEventHandler)
		d.On(LoadGame, LoadGameEventHandler)
		d.On(Passengers, PassengersEventHandler)
		d.On(Powerplay, PowerplayEventHandler)
		d.On(Progress, ProgressEventHandler)
		d.On(Rank, RankEventHandler)
		d.On(Reputation, ReputationEventHandler)
		d.On(Statistics, StatisticsEventHandler)
	}
}
