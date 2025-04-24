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
