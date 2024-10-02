package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username  string
	Password  string
	Questions []Question
}

type Question struct {
	gorm.Model
	UserID   uint
	Password string
	Content  string
}

type Answer struct {
	gorm.Model
	UserID     uint
	Password   string
	AnswerType AnswerType
}

type AnswerType struct {
	gorm.Model
	Description string
}
