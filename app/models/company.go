package models

import "time"

type Company struct {
	ID          uint       `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt,omitempty" gorm:"index"`
	Name        string     `json:"name"`
	FantasyName string     `json:"fantasyName"`
	UserID      uint       `json:"userId,omitempty"`
	User        *User      `json:"user,omitempty"`
}
