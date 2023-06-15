package repository

import (
	"fmt"
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

func (tr *TagRepository) CheckTagExistByName(tagName, userId string) bool {
	var count int64
	tr.db.Model(&db.Tag{}).Where(db.Tag{Title: tagName, UserID: userId}).Count(&count)
	fmt.Println(count)
	return count > 0
}

func (tr *TagRepository) CreateTag(tag db.Tag) (*db.Tag, error) {
	results := tr.db.Clauses(clause.Returning{}).Select("ID", "Title", "UserID", "Color", "Description").Create(&tag)
	if results.Error != nil {
		return nil, results.Error
	}
	return &tag, nil
}

func (tr *TagRepository) GetAllTags(userId string) ([]db.Tag, error) {
	var tags []db.Tag
	rows := tr.db.Where(db.Tag{UserID: userId}).Find(&tags)
	if rows.Error != nil {
		return nil, rows.Error
	}
	return tags, nil
}

func (tr *TagRepository) GetTagByName(name, userId string) ([]db.Tag, error) {
	var tags []db.Tag
	tr.db.Where("title LIKE ?", "%"+name+"%").Where(db.Tag{UserID: userId}).Find(&tags)
	//fmt.Println(tag)
	return tags, nil
}
func (tr *TagRepository) DeleteTagById(id, userId string) error {
	dt := tr.db.Where(db.Tag{UserID: userId}).Where("id", id).Delete(&db.Tag{})
	if dt.Error != nil {
		return dt.Error
	}
	return nil
}
