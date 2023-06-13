package repository

import (
	"github.com/cna-so/todo-api/models/db"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TagRepository struct {
	db *gorm.DB
}

func TagRepositoryProvider(db *gorm.DB) TagRepository {
	return TagRepository{
		db: db,
	}
}

func (tr *TagRepository) CheckTagExistByName(tagName string) bool {
	var count int64
	tr.db.Model(&db.Tags{}).Where(db.Tags{Title: tagName}).Count(&count)
	return count > 0
}

func (tr *TagRepository) CreateTag(tag db.Tags) (*db.Tags, error) {
	results := tr.db.Clauses(clause.Returning{}).Select("ID", "Title", "UserID", "Color", "Description").Create(&tag)
	if results.Error != nil {
		return nil, results.Error
	}
	return &tag, nil
}
