package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"tagd/models"
)

type TagHandler struct {
	db *gorm.DB
}

func NewTagHandler(db *gorm.DB) *TagHandler {
	return &TagHandler{db: db}
}

// GetTags Get tag list
// @Summary Get tag list
// @Description Query tag list with conditions
// @Tags tags
// @Accept json
// @Produce json
// @Param scope query string false "Code scope"
// @Param scope_type query string false "Scope type"
// @Param subject query string false "Subject category"
// @Param key_code query string false "Key code segment"
// @Success 200 {array} models.Tag
// @Router /tags [get]
func (h *TagHandler) GetTags(c *gin.Context) {
	var query models.TagPosition
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var tags []models.Tag
	if result := h.db.Where(&query).Find(&tags); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// GetTag Get tag details
// @Summary Get tag details
// @Description Get tag details by ID
// @Tags tags
// @Accept json
// @Produce json
// @Param tagid path int true "Tag ID"
// @Success 200 {object} models.Tag
// @Router /tags/{tagid} [get]
func (h *TagHandler) GetTag(c *gin.Context) {
	id := c.Param("tagid")

	var tag models.Tag
	if result := h.db.First(&tag, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// AddTag Add new tag
// @Summary Add new tag
// @Description Add a new tag
// @Tags tags
// @Accept json
// @Produce json
// @Param tag body models.Tag true "Tag information"
// @Success 201 {object} models.Tag
// @Router /tags [post]
func (h *TagHandler) AddTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.db.Create(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// UpdateTag Update tag
// @Summary Update tag
// @Description Update all tag fields
// @Tags tags
// @Accept json
// @Produce json
// @Param tagid path int true "Tag ID"
// @Param tag body models.Tag true "Updated tag information"
// @Success 200 {object} models.Tag
// @Router /tags/{tagid} [put]
func (h *TagHandler) UpdateTag(c *gin.Context) {
	id := c.Param("tagid")

	var existingTag models.Tag
	if result := h.db.First(&existingTag, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := h.db.Save(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// UpdateTagPair Update tag key-value pair
// @Summary Update tag key-value pair
// @Description Update specified key-value pair of the tag
// @Tags tags
// @Accept json
// @Produce json
// @Param tagid path int true "Tag ID"
// @Param key path string true "Key name"
// @Param value body string true "New value"
// @Success 200 {object} models.Tag
// @Router /tags/{tagid}/{key} [put]
func (h *TagHandler) UpdateTagPair(c *gin.Context) {
	id := c.Param("tagid")
	key := c.Param("key")

	var tag models.Tag
	if result := h.db.First(&tag, id); result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tag not found"})
		return
	}

	var value string
	if err := c.ShouldBindJSON(&value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tag.Pairs == nil {
		tag.Pairs = make(map[string]string)
	}
	tag.Pairs[key] = value

	if result := h.db.Save(&tag); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}

// DeleteTag Delete tag
// @Summary Delete tag
// @Description Delete tag by ID
// @Tags tags
// @Accept json
// @Produce json
// @Param tagid path int true "Tag ID"
// @Success 204
// @Router /tags/{tagid} [delete]
func (h *TagHandler) DeleteTag(c *gin.Context) {
	id := c.Param("tagid")

	if result := h.db.Delete(&models.Tag{}, id); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
