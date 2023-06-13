package api

import (
	"github.com/cna-so/todo-api/controllers/authorization/jwt"
	dto "github.com/cna-so/todo-api/models/DTO"
	"github.com/cna-so/todo-api/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type TagApi struct {
	s service.TagService
}

func (us *TagApi) CreateTag(ctx *gin.Context) {
	var tag dto.TagCreateRequestDto
	err := ctx.ShouldBindJSON(&tag)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "title are require and color is optional",
		})
		return
	}
	// get user details from jwt
	token := strings.Split(ctx.GetHeader("Authorization"), " ")
	tokenValues, _ := jwt.GetValueFromJWT(token[1], []string{"user_id"})
	tag.UserID = tokenValues["user_id"]
	tagResponse, statusCode, err := us.s.CreateTagService(tag)
	if err != nil {
		ctx.AbortWithStatusJSON(statusCode, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(statusCode, gin.H{
		"tag": dto.TagResponseDto{
			ID:          tagResponse.ID,
			Title:       tagResponse.Title,
			Color:       tagResponse.Color,
			Description: tagResponse.Description,
		},
	})
}

func (us *TagApi) GetAllTags(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")
	tokenValues, _ := jwt.GetValueFromJWT(token[1], []string{"user_id"})
	tags, err := us.s.FindAllTags(tokenValues["user_id"])
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNoContent, gin.H{
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"tags": tags,
	})

}
