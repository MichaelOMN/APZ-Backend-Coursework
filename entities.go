package entity

import (
	"errors"
)

type Visitor struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Coach struct {
	Id       int    `json:"-"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type PhysicalInfo struct {
	Id        int     `json:"id" db:"id"`
	VisitorId int     `json:"visitor_id" db:"visitor_id"`
	Weight    float64 `json:"weight" db:"weight"`
	Height    float64 `json:"height" db:"height"`
}

type Club struct {
	Id      int    `json:"id" db:"id"`
	Address string `json:"address" db:"address"`
	Name    string `json:"name" binding:"required" db:"club_name"`
}

type ClubUpdateForm struct {
	Address *string `json:"address"`
	Name    *string `json:"name"`
}

type PhysicalInfoUpdateForm struct {
	Weight *string `json:"weight"`
	Height *string `json:"height"`
}

func (piuf PhysicalInfoUpdateForm) Validate() error {
	if piuf.Height == nil && piuf.Weight == nil {
		return errors.New("weight and height fields abscent")
	}
	return nil
}

func (cuf ClubUpdateForm) Validate() error {
	if cuf.Address == nil && cuf.Name == nil {
		return errors.New("address and name field abscent")
	}

	return nil
}

type Training struct {
	Id      int    `json:"id" db:"id"`
	Start   string `json:"start" db:"start_datetime"`
	End     string `json:"end" db:"end_datetime"`
	CoachId int    `json:"coach_id" db:"coach_id"`
	ClubId  int    `json:"club_id" binding:"required" db:"club_id"`
}

type StatesTypes struct {
	Id         int    `json:"id" db:"id"`
	Name       string `json:"name" binding:"required" db:"type_name"`
	Decription string `json:"description" db:"type_desc"`
	Unit       string `json:"unit" binding:"required" db:"type_unit"`
}

type StatesTypesUpdateForm struct {
	Name       *string `json:"name"`
	Decription *string `json:"description"`
	Unit       *string `json:"unit"`
}

type Activity struct {
	Id          int    `json:"id" db:"id"`
	Name        string `json:"name" binding:"required" db:"activity_name"`
	Description string `json:"description" db:"activity_desc"`
	ClubId      int    `json:"club_id" db:"club_id"`
}

type ActivityUpdateForm struct {
	Name        *string `json:"name" binding:"required"`
	Description *string `json:"description" binding:"required"`
	ClubId      *int    `json:"club_id" binding:"required"`
}

type Attendance struct {
	Id         int `json:"id" db:"id"`
	VisitorId  int `json:"visitor_id" db:"visitor_id"`
	TrainingId int `json:"training_id" binding:"required" db:"training_id"`
}

type AttendanceUpdateForm struct {
	TrainingId *int `json:"training_id"`
}

func (a *AttendanceUpdateForm) Validate() error {
	if a.TrainingId == nil {
		return errors.New("TrainingId field missing")
	}
	return nil
}

type ActivityUsage struct {
	Id           int    `json:"id" db:"id"`
	VisitorId    int    `json:"visitor_id" db:"visitor_id"`
	ActivityName string `json:"activity_name" binding:"required" db:"activity_name"`
	Start        string `json:"start" db:"usage_start_datetime"`
	End          string `json:"end" db:"usage_end_datetime"`
	TrainingId   int    `json:"training_id" binding:"required" db:"training_id"`
}

type ActivityUsageUpdateForm struct {
	Start *string `json:"start"`
	End   *string `json:"end"`
}

func (a *ActivityUsageUpdateForm) Validate() error {
	if a.Start == nil && a.End == nil {
		return errors.New("fields are abscent")
	}
	return nil
}

type PhysicalState struct {
	Id              int     `json:"id" db:"id"`
	ActivityUsageId int     `json:"activity_usage_id" binding:"required" db:"activity_usage_id"`
	UnitAmount      float64 `json:"unit_amount" binding:"required" db:"unit_amount"`
	StateTypeId     int     `json:"state_type_id" binding:"required" db:"state_type_id"`
	At              string  `json:"at" db:"at_datetime"`
	Secs            float64 `json:"secs"  db:"duration_secs"`
}

type ActivityState struct {
	Id           int     `json:"id" db:"id"`
	ActivityName string  `json:"activity_name" db:"activity_name"`
	UnitAmount   float64 `json:"unit_amount" binding:"required" db:"unit_amount"`
	StateTypeId  int     `json:"state_type_id" binding:"required" db:"state_type_id"`
	At           string  `json:"at" db:"at_datetime"`
	Secs         float64 `json:"secs" db:"duration_secs"`
}
