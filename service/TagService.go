package service

import (
	"errors"
	"fmt"
	dto "github.com/cna-so/todo-api/models/DTO"
	"github.com/cna-so/todo-api/models/db"
	"github.com/cna-so/todo-api/repository"
	"net/http"
)

type TagService struct {
	repository repository.TagRepository
}

func TagServiceProvider(tagRepository repository.TagRepository) TagService {
	return TagService{
		repository: tagRepository,
	}
}

func (ts *TagService) CreateTagService(tag dto.TagCreateRequestDto) (*db.Tags, int, error) {
	if isExist := ts.repository.CheckTagExistByName(tag.Title); isExist == true {
		return nil, http.StatusConflict, errors.New(fmt.Sprintf("tag with name '%s' aleardy exist", tag.Title))
	}
	tagModel := db.Tags{Title: tag.Title, Color: tag.Color, UserID: tag.UserID, Description: tag.Description}
	createTag, err := ts.repository.CreateTag(tagModel)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	return createTag, http.StatusCreated, nil
}
func (ts *TagService) FindAllTags(userID string) ([]dto.TagResponseDto, error) {
	tags, err := ts.repository.GetAllTags(userID)
	if err != nil {
		return nil, err
	}
	var resTags []dto.TagResponseDto
	for _, tag := range tags {
		resTags = append(resTags, dto.TagResponseDto{ID: tag.ID, Title: tag.Title, Description: tag.Description, Color: tag.Color})
	}
	return resTags, nil
}
