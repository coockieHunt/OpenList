package routes

import (
	"OpenList/Go/sqlite"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllLists: Retrieves all lists with their items
func GetAllLists(c *gin.Context) {
	var lists []sqlite.List

	if err := sqlite.DB.Preload("Items").Find(&lists).Error; err != nil {
		println("GORM Error:", err.Error())
		SendResponse(c, "error", "failed to retrieve lists", nil)
		return
	}

	SendResponse(c, "success", "lists retrieved successfully", lists)
}

// GetListByID: Retrieves a specific list by ID with its items
func GetListByID(c *gin.Context) {
	listID, _ := strconv.Atoi(c.Param("idList"))
	if listID <= 0 {
		SendResponse(c, "error", "invalid list ID", nil)
		return
	}

	var list sqlite.List
	if err := sqlite.DB.Preload("Items").First(&list, listID).Error; err != nil {
		SendResponse(c, "error", "list not found", nil)
		return
	}

	SendResponse(c, "success", "list retrieved successfully", list)
}

// NewList: Creates a new list with optional items
func NewList(c *gin.Context) {
	var newList sqlite.List

	if err := c.ShouldBindJSON(&newList); err != nil {
		SendResponse(c, "error", "invalid payload", nil)
		return
	}

	if newList.Title == "" {
		SendResponse(c, "error", "title is required", nil)
		return
	}

	if len(newList.Items) > 0 && newList.Items[0].Name == "" {
		newList.Items = []sqlite.Item{}
	}

	if err := sqlite.DB.Create(&newList).Error; err != nil {
		println("GORM Error:", err.Error())
		SendResponse(c, "error", "failed to create list", nil)
		return
	}

	SendResponse(c, "success", "list created successfully", newList)
}

// DeleteList: Deletes a list by ID
func DeleteList(c *gin.Context) {
	listID, err := strconv.Atoi(c.Param("idList"))
	if err != nil || listID <= 0 {
		SendResponse(c, "error", "invalid list ID", nil)
		return
	}

	err = sqlite.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("list_id = ?", listID).Delete(&sqlite.Item{}).Error; err != nil {
			return err
		}

		result := tx.Delete(&sqlite.List{}, listID)
		if result.Error != nil {
			return result.Error
		}

		if result.RowsAffected == 0 {
			return gorm.ErrRecordNotFound
		}

		return nil
	})

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			SendResponse(c, "error", "list not found", nil)
			return
		}

		println("GORM Error:", err.Error())
		SendResponse(c, "error", "database error", nil)
		return
	}

	SendResponse(c, "success", "list deleted successfully", nil)
}
