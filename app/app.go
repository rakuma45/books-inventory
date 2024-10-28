package app

import "github.com/jinzhu/gorm"

type Handler struct{
	DB *gorm.DB
}

func New(db *gorm.DB) Handler {
	return Handler{DB: db}
}

func (h *Handler) GetBooks(c *gin) {
	
}
