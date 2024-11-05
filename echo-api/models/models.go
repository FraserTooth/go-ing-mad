package models

import "gorm.io/gorm"

type Note struct {
  gorm.Model
	Title string `json:"title"`
	Content string `json:"content"`
	Done bool `json:"done"`
}