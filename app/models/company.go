package models

import "time"

type Company struct {
	ID          uint       `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt" gorm:"index"`
	Name        string     `json:"name"`
	FantasyName string     `json:"fantasyName"`
	UserID      uint       `json:"userId"`
	User        *User      `json:"user"`
}
