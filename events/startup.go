package events

import (
	"encoding/json"
	"fmt"

	"github.com/sht/ed-journal/event"
)

const (
	Cargo          = "Cargo"
	ClearSavedGame = "ClearSavedGame"
	Commander      = "Commander"
	Loadout        = "Loadout"
	Materials      = "Materials"
	Missions       = "Missions"
	NewCommander   = "NewCommander"
	LoadGame       = "LoadGame"
	Passengers     = "Passengers"
	Powerplay      = "Powerplay"
	Progress       = "Progress"
	Rank           = "Rank"
	Reputation     = "Reputation"
	Statistics     = "Statistics"
)

type CargoEvent struct {
	event.Event
	Vessel    string `json:"Vessel"`
	Count     int    `json:"Count"`
	Inventory *[]struct {
		Name          string `json:"Name"`
		Count         int    `json:"Count"`
		Stolen        int    `json:"Stolen"`
		MissionID     *int   `json:"MissionID,omitempty"`
		NameLocalised string `json:"Name_Localised,omitempty"`
	} `json:"Inventory,omitempty"`
}

func CargoEventHandler(b []byte) {
	var e CargoEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type ClearSavedGameEvent struct {
	event.Event
	Name string `json:"Name"`
	FID  string `json:"FID"`
}

func ClearSavedGameEventHandler(b []byte) {
	var e ClearSavedGameEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		return
	}

	debug(b, e)
}

type CommanderEvent struct {
	event.Event
	Name string `json:"Name"`
	FID  string `json:"FID"`
}

func CommanderEventHandler(b []byte) {
	var e CommanderEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type LoadoutEvent struct {
	event.Event
	Ship         string  `json:"Ship"`
	ShipID       int     `json:"ShipID"`
	ShipName     string  `json:"ShipName"`
	ShipIdent    string  `json:"ShipIdent"`
	HullValue    *int    `json:"HullValue,omitempty"`
	ModulesValue *int    `json:"ModulesValue,omitempty"`
	HullHealth   float64 `json:"HullHealth"`
	UnladenMass  float64 `json:"UnladenMass"`
	FuelCapacity struct {
		Main    float64 `json:"Main"`
		Reserve float64 `json:"Reserve"`
	} `json:"FuelCapacity"`
	CargoCapacity int     `json:"CargoCapacity"`
	MaxJumpRange  float64 `json:"MaxJumpRange"`
	Rebuy         int     `json:"Rebuy"`
	Hot           *bool   `json:"Hot,omitempty"`
	Modules       []*struct {
		Slot         string   `json:"Slot"`
		Item         string   `json:"Item"`
		On           bool     `json:"On"`
		Priority     int      `json:"Priority"`
		Health       float64  `json:"Health"`
		Value        *float64 `json:"Value,omitempty"`
		AmmoInClip   *int     `json:"AmmoInClip,omitempty"`
		AmmoInHopper *int     `json:"AmmoInHopper,omitempty"`
		Engineering  *struct {
			Engineer                    string  `json:"Engineer"`
			EngineerID                  uint64  `json:"EngineerID"`
			BlueprintName               string  `json:"BlueprintName"`
			BlueprintID                 int     `json:"BlueprintID"`
			Level                       int     `json:"Level"`
			Quality                     float64 `json:"Quality"`
			ExperimentalEffect          *string `json:"ExperimentalEffect,omitempty"`
			ExperimentalEffectLocalised *string `json:"ExperimentalEffect_Localised,omitempty"`
			Modifiers                   []*struct {
				Label         string   `json:"Label"`
				Value         *float64 `json:"Value,omitempty"`
				OriginalValue float64  `json:"OriginalValue"`
				LessIsGood    int      `json:"LessIsGood"`
			} `json:"Modifiers"`
		} `json:"Engineering,omitempty"`
	} `json:"Modules"`
}

func LoadoutEventHandler(b []byte) {
	var e LoadoutEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type MaterialsEvent struct {
	event.Event
	Raw          []*Material `json:"Raw"`
	Manufactured []*Material `json:"Manufactured"`
	Encoded      []*Material `json:"Encoded"`
}

type Material struct {
	Name          string  `json:"Name"`
	NameLocalised *string `json:"Name_Localised,omitempty"`
	Count         int     `json:"Count"`
}

func MaterialsEventHandler(b []byte) {
	var e MaterialsEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type MissionsEvent struct {
	event.Event
	Active   []*Mission `json:"Active"`
	Failed   []*Mission `json:"Failed"`
	Complete []*Mission `json:"Complete"`
}

type Mission struct {
	MissionID        int    `json:"MissionID"`
	Name             string `json:"Name"`
	PassengerMission bool   `json:"PassengerMission"`
	Expires          int    `json:"Expires"`
}

func MissionsEventHandler(b []byte) {
	var e MissionsEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type NewCommanderEvent struct {
	event.Event
	Name    string `json:"Name"`
	FID     string `json:"FID"`
	Package string `json:"Package"`
}

func NewCommanderEventHandler(b []byte) {
	var e NewCommanderEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type LoadGameEvent struct {
	event.Event
	FID           string   `json:"FID"`
	Commander     string   `json:"Commander"`
	Horizons      bool     `json:"Horizons"`
	Ship          *string  `json:"Ship,omitempty"`
	ShipLocalised *string  `json:"Ship_Localised,omitempty"`
	ShipID        *int     `json:"ShipID,omitempty"`
	StartLanded   *bool    `json:"StartLanded,omitempty"`
	StartDead     *bool    `json:"StartDead,omitempty"`
	GameMode      *string  `json:"GameMode,omitempty"`
	Group         *string  `json:"Group,omitempty"`
	Credits       int      `json:"Credits"`
	Loan          int      `json:"Loan"`
	ShipName      *string  `json:"ShipName,omitempty"`
	ShipIdent     *string  `json:"ShipIdent,omitempty"`
	FuelLevel     *float64 `json:"FuelLevel,omitempty"`
	FuelCapacity  *float64 `json:"FuelCapacity,omitempty"`
}

func LoadGameEventHandler(b []byte) {
	var e LoadGameEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type PassengersEvent struct {
	event.Event
	MissionID int    `json:"MissionID"`
	Type      string `json:"Type"`
	VIP       bool   `json:"VIP"`
	Wanted    bool   `json:"Wanted"`
	Count     int    `json:"Count"`
}

func PassengersEventHandler(b []byte) {
	var e PassengersEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type PowerplayEvent struct {
	event.Event
	Power       string `json:"Power"`
	Rank        int    `json:"Rank"`
	Merits      int    `json:"Merits"`
	Votes       int    `json:"Votes"`
	TimePledged int    `json:"TimePledged"`
}

func PowerplayEventHandler(b []byte) {
	var e PowerplayEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type ProgressEvent struct {
	event.Event
	Combat     int `json:"Combat"`
	Trade      int `json:"Trade"`
	Explore    int `json:"Explore"`
	Empire     int `json:"Empire"`
	Federation int `json:"Federation"`
	CQC        int `json:"CQC"`
}

func ProgressEventHandler(b []byte) {
	var e ProgressEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type RankEvent struct {
	event.Event
	Combat     int `json:"Combat"`
	Trade      int `json:"Trade"`
	Explore    int `json:"Explore"`
	Empire     int `json:"Empire"`
	Federation int `json:"Federation"`
	CQC        int `json:"CQC"`
}

func RankEventHandler(b []byte) {
	var e RankEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type ReputationEvent struct {
	event.Event
	Empire      float64  `json:"Empire"`
	Federation  float64  `json:"Federation"`
	Independent *float64 `json:"Independent,omitempty"`
	Alliance    float64  `json:"Alliance"`
}

func ReputationEventHandler(b []byte) {
	var e ReputationEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}

type StatisticsEvent struct {
	event.Event
	BankAccount *struct {
		CurrentWealth          int64 `json:"Current_Wealth"`
		SpentOnShips           int   `json:"Spent_On_Ships"`
		SpentOnOutfitting      int64 `json:"Spent_On_Outfitting"`
		SpentOnRepairs         int   `json:"Spent_On_Repairs"`
		SpentOnFuel            int   `json:"Spent_On_Fuel"`
		SpentOnAmmoConsumables int   `json:"Spent_On_Ammo_Consumables"`
		InsuranceClaims        int   `json:"Insurance_Claims"`
		SpentOnInsurance       int   `json:"Spent_On_Insurance"`
		OwnedShipCount         int   `json:"Owned_Ship_Count"`
	} `json:"Bank_Account,omitempty"`
	Combat *struct {
		BountiesClaimed      int `json:"Bounties_Claimed"`
		BountyHuntingProfit  int `json:"Bounty_Hunting_Profit"`
		CombatBonds          int `json:"Combat_Bonds"`
		CombatBondProfits    int `json:"Combat_Bond_Profits"`
		Assassinations       int `json:"Assassinations"`
		AssassinationProfits int `json:"Assassination_Profits"`
		HighestSingleReward  int `json:"Highest_Single_Reward"`
		SkimmersKilled       int `json:"Skimmers_Killed"`
	} `json:"Combat,omitempty"`
	Crime *struct {
		Notoriety        int `json:"Notoriety"`
		Fines            int `json:"Fines"`
		TotalFines       int `json:"Total_Fines"`
		BountiesReceived int `json:"Bounties_Received"`
		TotalBounties    int `json:"Total_Bounties"`
		HighestBounty    int `json:"Highest_Bounty"`
	} `json:"Crime,omitempty"`
	Smuggling *struct {
		BlackMarketsTradedWith   int     `json:"Black_Markets_Traded_With"`
		BlackMarketsProfits      int     `json:"Black_Markets_Profits"`
		ResourcesSmuggled        int     `json:"Resources_Smuggled"`
		AverageProfit            float64 `json:"Average_Profit"`
		HighestSingleTransaction int     `json:"Highest_Single_Transaction"`
	} `json:"Smuggling,omitempty"`
	Trading *struct {
		MarketsTradedWith        int     `json:"Markets_Traded_With"`
		MarketProfits            int64   `json:"Market_Profits"`
		ResourcesTraded          int     `json:"Resources_Traded"`
		AverageProfit            float64 `json:"Average_Profit"`
		HighestSingleTransaction int     `json:"Highest_Single_Transaction"`
	} `json:"Trading,omitempty"`
	Mining *struct {
		MiningProfits      int64 `json:"Mining_Profits"`
		QuantityMined      int   `json:"Quantity_Mined"`
		MaterialsCollected int   `json:"Materials_Collected"`
	} `json:"Mining,omitempty"`
	Exploration *struct {
		SystemsVisited            int     `json:"Systems_Visited"`
		ExplorationProfits        int     `json:"Exploration_Profits"`
		PlanetsScannedToLevel2    int     `json:"Planets_Scanned_To_Level_2"`
		PlanetsScannedToLevel3    int     `json:"Planets_Scanned_To_Level_3"`
		EfficientScans            int     `json:"Efficient_Scans"`
		HighestPayout             int     `json:"Highest_Payout"`
		TotalHyperspaceDistance   int     `json:"Total_Hyperspace_Distance"`
		TotalHyperspaceJumps      int     `json:"Total_Hyperspace_Jumps"`
		GreatestDistanceFromStart float64 `json:"Greatest_Distance_From_Start"`
		TimePlayed                int     `json:"Time_Played"`
	} `json:"Exploration,omitempty"`
	Passengers *struct {
		PassengersMissionsBulk      int `json:"Passengers_Missions_Bulk"`
		PassengersMissionsVIP       int `json:"Passengers_Missions_VIP"`
		PassengersMissionsDelivered int `json:"Passengers_Missions_Delivered"`
		PassengersMissionsEjected   int `json:"Passengers_Missions_Ejected"`
	} `json:"Passengers,omitempty"`
	SearchAndRescue *struct {
		SearchRescueTraded int `json:"SearchRescue_Traded"`
		SearchRescueProfit int `json:"SearchRescue_Profit"`
		SearchRescueCount  int `json:"SearchRescue_Count"`
	} `json:"Search_And_Rescue,omitempty"`
	Crafting *struct {
		CountOfUsedEngineers  int `json:"Count_Of_Used_Engineers"`
		RecipesGenerated      int `json:"Recipes_Generated"`
		RecipesGeneratedRank1 int `json:"Recipes_Generated_Rank_1"`
		RecipesGeneratedRank2 int `json:"Recipes_Generated_Rank_2"`
		RecipesGeneratedRank3 int `json:"Recipes_Generated_Rank_3"`
		RecipesGeneratedRank4 int `json:"Recipes_Generated_Rank_4"`
		RecipesGeneratedRank5 int `json:"Recipes_Generated_Rank_5"`
	} `json:"Crafting,omitempty"`
	Crew *struct {
		NpcCrewTotalWages int `json:"NpcCrew_TotalWages"`
		NpcCrewHired      int `json:"NpcCrew_Hired"`
		NpcCrewFired      int `json:"NpcCrew_Fired"`
		NpcCrewDied       int `json:"NpcCrew_Died"`
	} `json:"Crew,omitempty"`
	Multicrew *struct {
		MulticrewTimeTotal        int `json:"Multicrew_Time_Total"`
		MulticrewGunnerTimeTotal  int `json:"Multicrew_Gunner_Time_Total"`
		MulticrewFighterTimeTotal int `json:"Multicrew_Fighter_Time_Total"`
		MulticrewCreditsTotal     int `json:"Multicrew_Credits_Total"`
		MulticrewFinesTotal       int `json:"Multicrew_Fines_Total"`
	} `json:"Multicrew,omitempty"`
	MaterialTraderStats *struct {
		TradesCompleted        int `json:"Trades_Completed"`
		MaterialsTraded        int `json:"Materials_Traded"`
		EncodedMaterialsTraded int `json:"Encoded_Materials_Traded"`
		RawMaterialsTraded     int `json:"Raw_Materials_Traded"`
		Grade1MaterialsTraded  int `json:"Grade_1_Materials_Traded"`
		Grade2MaterialsTraded  int `json:"Grade_2_Materials_Traded"`
		Grade3MaterialsTraded  int `json:"Grade_3_Materials_Traded"`
		Grade4MaterialsTraded  int `json:"Grade_4_Materials_Traded"`
		Grade5MaterialsTraded  int `json:"Grade_5_Materials_Traded"`
	} `json:"Material_Trader_Stats,omitempty"`
	CQC *struct {
		CQCCreditsEarned int     `json:"CQC_Credits_Earned"`
		CQCTimePlayed    int     `json:"CQC_Time_Played"`
		CQCKD            float64 `json:"CQC_KD"`
		CQCKills         int     `json:"CQC_Kills"`
		CQCWL            int     `json:"CQC_WL"`
	} `json:"CQC,omitempty"`
}

func StatisticsEventHandler(b []byte) {
	var e StatisticsEvent
	err := json.Unmarshal(b, &e)
	if err != nil {
		fmt.Println(err)
		return
	}

	debug(b, e)
}
