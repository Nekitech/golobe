package controllers

import (
	"golobe/model"
	"gorm.io/gorm"
)

type BookingScheme struct {
	gorm.Model
	DB      *gorm.DB
	Booking model.Booking
}
