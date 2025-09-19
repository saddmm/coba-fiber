package model

type Post struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title" gorm:"not null, type:varchar(100)"`
	Content string `json:"content" gorm:"not null, type:text"`
	UserID  uint   `json:"user_id" gorm:"foreignKey:UserID, onDelete:cascade, not null"`
}
