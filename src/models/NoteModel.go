package models

import (
	db "github.com/ebarquero85/link-backend/src/database"
)

type NoteModel struct {
	NoteId int    `gorm:"column:note_id;type:integer;primaryKey"`
	Note   string `gorm:"column:note;type:text;not null"`
}

func (NoteModel) TableName() string {
	return "public.notes"
}

func (Note *NoteModel) Create() error {

	if err := db.Databases.DBPostgresql.Instance.Create(Note).Error; err != nil {
		return err
	}

	return nil

}

func (Note *NoteModel) Update(id int) error {

	if err := db.Databases.DBPostgresql.Instance.Where("note_id = ?", id).Updates(&Note); err != nil {
		return err.Error
	}

	return nil

}
