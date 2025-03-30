package models

import (
	"errors"
	"time"

	"gorm.io/gorm"
)

type RequestStatus string

const (
	Pending  RequestStatus = "pending"
	Approved RequestStatus = "approved"
	Rejected RequestStatus = "rejected"
)

var validStatuses = map[RequestStatus]bool{
	Pending:  true,
	Approved: true,
	Rejected: true,
}

type ClockRequest struct {
	ID         uint          `gorm:"primaryKey"`
	EmployeeID uint          `gorm:"not null"`
	ClockID    uint          `gorm:"not null"`
	Type       TimeClockType `gorm:"type:varchar(20);not null"`
	Time       time.Time     `gorm:"not null"`
	Status     RequestStatus `gorm:"type:varchar(20);default:'pending'"`
	Reason     string        `gorm:"type:text"`
	ReviewedBy *uint
	ReviewedAt *time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (r *ClockRequest) BeforeCreate(tx *gorm.DB) error {
	return r.validate()
}

func (r *ClockRequest) BeforeUpdate(tx *gorm.DB) error {
	return r.validate()
}

func (r *ClockRequest) validate() error {
	if !validStatuses[r.Status] {
		return errors.New("invalid request status: " + string(r.Status))
	}
	return nil
}
