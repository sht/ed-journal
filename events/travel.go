package events

import (
	"encoding/json"
	"github.com/sht/ed-journal/event"
)

const (
	ApproachBody     = "ApproachBody"
	Docked           = "Docked"
	DockingCancelled = "DockingCancelled"
	DockingDenied    = "DockingDenied"
	DockingGranted   = "DockingGranted"
	DockingRequested = "DockingRequested"
	DockingTimeout   = "DockingTimeout"
	FSDJump          = "FSDJump"
	FSDTarget        = "FSDTarget"
	LeaveBody        = "LeaveBody"
	Liftoff          = "Liftoff"
	Location         = "Location"
	StartJump        = "StartJump"
	SupercruiseEntry = "SupercruiseEntry"
	SupercruiseExit  = "SupercruiseExit"
	Touchdown        = "Touchdown"
	Undocked         = "Undocked"
	Route            = "Route"
)

type ApproachBodyEvent struct {
	event.Event
	Body          string `json:"Body"`
	BodyID        int    `json:"BodyID"`
	StarSystem    string `json:"StarSystem"`
	SystemAddress int    `json:"SystemAddress"`
}

func ApproachBodyEventHandler(b []byte) {
	var e ApproachBodyEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type DockedEvent struct {
	event.Event
	ActiveFine        bool    `json:"ActiveFine,omitempty"`
	DistFromStarLS    float64 `json:"DistFromStarLS"`
	MarketID          int     `json:"MarketID"`
	StarSystem        string  `json:"StarSystem"`
	StationAllegiance string  `json:"StationAllegiance,omitempty"`
	StationEconomies  []*struct {
		Name          string  `json:"Name"`
		NameLocalised string  `json:"Name_Localised"`
		Proportion    float64 `json:"Proportion"`
	} `json:"StationEconomies"`
	StationEconomy          string `json:"StationEconomy"`
	StationEconomyLocalised string `json:"StationEconomy_Localised"`
	StationFaction          *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name"`
	} `json:"StationFaction,omitempty"`
	StationGovernment          string   `json:"StationGovernment"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised"`
	StationName                string   `json:"StationName"`
	StationServices            []string `json:"StationServices"`
	StationType                string   `json:"StationType"`
	SystemAddress              int      `json:"SystemAddress"`
	Wanted                     bool     `json:"Wanted,omitempty"`
}

func DockedEventHandler(b []byte) {
	var e DockedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type DockingCancelledEvent struct {
	event.Event
	MarketID    int    `json:"MarketID,omitempty"`
	StationName string `json:"StationName,omitempty"`
	StationType string `json:"StationType,omitempty"`
}

func DockingCancelledEventHandler(b []byte) {
	var e DockingCancelledEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type DockingDeniedEvent struct {
	event.Event
	MarketID    int    `json:"MarketID"`
	Reason      string `json:"Reason"`
	StationName string `json:"StationName"`
	StationType string `json:"StationType"`
}

func DockingDeniedEventHandler(b []byte) {
	var e DockingDeniedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type DockingGrantedEvent struct {
	event.Event
	LandingPad  int    `json:"LandingPad"`
	MarketID    int    `json:"MarketID"`
	StationName string `json:"StationName"`
	StationType string `json:"StationType"`
}

func DockingGrantedEventHandler(b []byte) {
	var e DockingGrantedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type DockingRequestedEvent struct {
	event.Event
	MarketID    int    `json:"MarketID"`
	StationName string `json:"StationName"`
	StationType string `json:"StationType"`
}

func DockingRequestedEventHandler(b []byte) {
	var e DockingRequestedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type DockingTimeoutEvent struct {
	event.Event
}

func DockingTimeoutEventHandler(b []byte) {
	var e DockingTimeoutEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type FSDJumpEvent struct {
	event.Event
	Body     string `json:"Body"`
	BodyID   int    `json:"BodyID"`
	BodyType string `json:"BodyType"`
	Factions []*struct {
		Allegiance         string  `json:"Allegiance"`
		FactionState       string  `json:"FactionState"`
		Government         string  `json:"Government"`
		Happiness          string  `json:"Happiness"`
		HappinessLocalised string  `json:"Happiness_Localised"`
		Influence          float64 `json:"Influence"`
		MyReputation       int     `json:"MyReputation"`
		Name               string  `json:"Name"`
		ActiveStates       []*struct {
			State string `json:"State"`
		} `json:"ActiveStates,omitempty"`
	} `json:"Factions,omitempty"`
	FuelLevel              float64   `json:"FuelLevel"`
	FuelUsed               float64   `json:"FuelUsed"`
	JumpDist               float64   `json:"JumpDist"`
	Population             int       `json:"Population"`
	PowerplayState         string    `json:"PowerplayState,omitempty"`
	Powers                 []string  `json:"Powers,omitempty"`
	StarPos                []float64 `json:"StarPos"`
	StarSystem             string    `json:"StarSystem"`
	SystemAddress          int       `json:"SystemAddress"`
	SystemAllegiance       string    `json:"SystemAllegiance"`
	SystemEconomy          string    `json:"SystemEconomy"`
	SystemEconomyLocalised string    `json:"SystemEconomy_Localised"`
	SystemFaction          *struct {
		FactionState string `json:"FactionState"`
		Name         string `json:"Name"`
	} `json:"SystemFaction,omitempty"`
	SystemGovernment             string `json:"SystemGovernment"`
	SystemGovernmentLocalised    string `json:"SystemGovernment_Localised"`
	SystemSecondEconomy          string `json:"SystemSecondEconomy"`
	SystemSecondEconomyLocalised string `json:"SystemSecondEconomy_Localised"`
	SystemSecurity               string `json:"SystemSecurity"`
	SystemSecurityLocalised      string `json:"SystemSecurity_Localised"`
	BoostUsed                    int    `json:"BoostUsed,omitempty"`
}

func FSDJumpEventHandler(b []byte) {
	var e FSDJumpEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type FSDTargetEvent struct {
	event.Event
	Name                  string `json:"Name"`
	RemainingJumpsInRoute int    `json:"RemainingJumpsInRoute,omitempty"`
	StarClass             string `json:"StarClass,omitempty"`
	SystemAddress         int    `json:"SystemAddress"`
}

func FSDTargetEventHandler(b []byte) {
	var e FSDTargetEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type LeaveBodyEvent struct {
	event.Event
	Body          string `json:"Body"`
	BodyID        int    `json:"BodyID"`
	StarSystem    string `json:"StarSystem"`
	SystemAddress int64  `json:"SystemAddress"`
}

func LeaveBodyEventHandler(b []byte) {
	var e LeaveBodyEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type LiftoffEvent struct {
	event.Event
	NearestDestination string  `json:"NearestDestination,omitempty"`
	Latitude           float64 `json:"Latitude,omitempty"`
	Longitude          float64 `json:"Longitude,omitempty"`
	PlayerControlled   bool    `json:"PlayerControlled"`
}

func LiftoffEventHandler(b []byte) {
	var e LiftoffEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type LocationEvent struct {
	event.Event
	Body     string `json:"Body"`
	BodyID   int    `json:"BodyID"`
	BodyType string `json:"BodyType"`
	Docked   bool   `json:"Docked"`
	Factions []*struct {
		Allegiance         string  `json:"Allegiance"`
		FactionState       string  `json:"FactionState"`
		Government         string  `json:"Government"`
		Happiness          string  `json:"Happiness"`
		HappinessLocalised string  `json:"Happiness_Localised"`
		Influence          float64 `json:"Influence"`
		MyReputation       int     `json:"MyReputation"`
		Name               string  `json:"Name"`
		RecoveringStates   []*struct {
			State string `json:"State"`
			Trend int    `json:"Trend"`
		} `json:"RecoveringStates,omitempty"`
		ActiveStates []*struct {
			State string `json:"State"`
		} `json:"ActiveStates,omitempty"`
	} `json:"Factions,omitempty"`
	MarketID          int       `json:"MarketID,omitempty"`
	Population        int       `json:"Population"`
	PowerplayState    string    `json:"PowerplayState,omitempty"`
	Powers            []string  `json:"Powers,omitempty"`
	StarPos           []float64 `json:"StarPos"`
	StarSystem        string    `json:"StarSystem"`
	StationAllegiance string    `json:"StationAllegiance,omitempty"`
	StationEconomies  []*struct {
		Name          string  `json:"Name"`
		NameLocalised string  `json:"Name_Localised"`
		Proportion    float64 `json:"Proportion"`
	} `json:"StationEconomies,omitempty"`
	StationEconomy          string `json:"StationEconomy,omitempty"`
	StationEconomyLocalised string `json:"StationEconomy_Localised,omitempty"`
	StationFaction          *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name"`
	} `json:"StationFaction,omitempty"`
	StationGovernment          string   `json:"StationGovernment,omitempty"`
	StationGovernmentLocalised string   `json:"StationGovernment_Localised,omitempty"`
	StationName                string   `json:"StationName,omitempty"`
	StationServices            []string `json:"StationServices,omitempty"`
	StationType                string   `json:"StationType,omitempty"`
	SystemAddress              int      `json:"SystemAddress"`
	SystemAllegiance           string   `json:"SystemAllegiance"`
	SystemEconomy              string   `json:"SystemEconomy"`
	SystemEconomyLocalised     string   `json:"SystemEconomy_Localised"`
	SystemFaction              *struct {
		FactionState string `json:"FactionState,omitempty"`
		Name         string `json:"Name"`
	} `json:"SystemFaction,omitempty"`
	SystemGovernment             string `json:"SystemGovernment"`
	SystemGovernmentLocalised    string `json:"SystemGovernment_Localised"`
	SystemSecondEconomy          string `json:"SystemSecondEconomy"`
	SystemSecondEconomyLocalised string `json:"SystemSecondEconomy_Localised"`
	SystemSecurity               string `json:"SystemSecurity"`
	SystemSecurityLocalised      string `json:"SystemSecurity_Localised"`
}

func LocationEventHandler(b []byte) {
	var e LocationEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type StartJumpEvent struct {
	event.Event
	JumpType      string `json:"JumpType"`
	StarClass     string `json:"StarClass,omitempty"`
	StarSystem    string `json:"StarSystem,omitempty"`
	SystemAddress int    `json:"SystemAddress,omitempty"`
}

func StartJumpEventHandler(b []byte) {
	var e StartJumpEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type SupercruiseEntryEvent struct {
	event.Event
	StarSystem    string `json:"StarSystem"`
	SystemAddress int    `json:"SystemAddress"`
}

func SupercruiseEntryEventHandler(b []byte) {
	var e SupercruiseEntryEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type SupercruiseExitEvent struct {
	event.Event
	Body          string `json:"Body"`
	BodyID        int    `json:"BodyID"`
	BodyType      string `json:"BodyType"`
	StarSystem    string `json:"StarSystem"`
	SystemAddress int    `json:"SystemAddress"`
}

func SupercruiseExitEventHandler(b []byte) {
	var e SupercruiseExitEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type TouchdownEvent struct {
	event.Event
	NearestDestination string  `json:"NearestDestination,omitempty"`
	Latitude           float64 `json:"Latitude,omitempty"`
	Longitude          float64 `json:"Longitude,omitempty"`
	PlayerControlled   bool    `json:"PlayerControlled"`
}

func TouchdownEventHandler(b []byte) {
	var e TouchdownEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type UndockedEvent struct {
	event.Event
	MarketID    int    `json:"MarketID,omitempty"`
	StationName string `json:"StationName"`
	StationType string `json:"StationType"`
}

func UndockedEventHandler(b []byte) {
	var e UndockedEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type RouteEvent struct {
	event.Event
}

func RouteEventHandler(b []byte) {
	var e RouteEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}
