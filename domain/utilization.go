package domain

import "time"

type Utilization struct {
	ID           int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Date         time.Time `json:"date" gorm:"type:date;notNull"`
	Contract     string    `json:"contract" gorm:"type:varchar(100);notNull"`
	MoveType     string    `json:"move_type" gorm:"type:varchar(100);notNull"`
	UnitCategory string    `json:"unit_category" gorm:"type:varchar(100);notNull"`
	UnitName     string    `json:"unit_name" gorm:"type:varchar(100);notNull"`
	Unit         string    `json:"unit" gorm:"type:varchar(100);notNull"`
	Condition    string    `json:"condition" gorm:"type:varchar(50);notNull"`
	CreatedAt    time.Time `json:"created_at" gorm:"notNull"`
}

type Utilizations struct {
	Data []MoveType `json:"data"`
}

type MoveType struct {
	MoveType    string        `json:"move_type"`
	Contract    string        `json:"contract"`
	Realization []Realization `json:"realization"`
	Units       UnitData      `json:"units"`
}

type Realization struct {
	UnitName         string   `json:"unit_name"`
	TotalAvailable   int      `json:"total_available"`
	AlocationFromMPE int      `json:"alocation_from_mpe"`
	Realization      int      `json:"realization"`
	AdditionalInfo   []string `json:"additional_info"`
	BD               int      `json:"bd"`
	Accident         int      `json:"accident"`
	TLO              int      `json:"tlo"`
	CMS              int      `json:"cms"`
	GassDiff         int      `json:"gas_diff"`
	Standby          int      `json:"standby"`
}

type UnitData struct {
	OprEngineOn              int `json:"opr_engine_on"`
	IdleEngineOn             int `json:"idle_engine_on"`
	StdEngineOff             int `json:"std_engine_off"`
	StdEngineUnknown         int `json:"std_engine_unkown"`
	BdEngineOn               int `json:"bd_engine_on"`
	BdEngineOff              int `json:"bg_engine_off"`
	BdEngineUnknown          int `json:"bd_engine_unkown"`
	AcdEngineUnknown         int `json:"acd_engine_unkown"`
	ChasisCrackEngineUnknown int `json:"chasis_crack_engine_unkown"`
}

func NewUtilization(date time.Time, contract, moveType, unitCategory, unitName, unit, condition string) Utilization {
	return Utilization{
		Date:         date,
		Contract:     contract,
		MoveType:     moveType,
		UnitCategory: unitCategory,
		UnitName:     unitName,
		Unit:         unit,
		Condition:    condition,
		CreatedAt:    time.Now().UTC(),
	}
}
