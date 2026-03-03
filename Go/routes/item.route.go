package routes

import (
	"OpenList/Go/sqlite"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddItem: Adds a new item to a specific list
func AddItem(c *gin.Context) {
	var item sqlite.Item
	listID, err := strconv.Atoi(c.Param("idList"))
	if err != nil || listID <= 0 {
		SendResponse(c, "error", "invalid list ID", nil)
		return
	}

	if err := c.ShouldBindJSON(&item); err != nil {
		println("JSON Binding Error:", err.Error())
		SendResponse(c, "error", "invalid payload", nil)
		return
	}

	item.ListID = uint(listID)

	if err := sqlite.DB.Create(&item).Error; err != nil {
		println("GORM Error:", err.Error())
		SendResponse(c, "error", "database error", nil)
		return
	}

	SendResponse(c, "success", "item added successfully", item)
}

// DeleteItem: Deletes an item by ID from a specific list
func DeleteItem(c *gin.Context) {
	listID, _ := strconv.Atoi(c.Param("idList"))
	itemID, err := strconv.Atoi(c.Param("idItem"))

	if err != nil || listID <= 0 || itemID <= 0 {
		SendResponse(c, "error", "invalid IDs", nil)
		return
	}

	result := sqlite.DB.Where("id = ? AND list_id = ?", itemID, listID).Delete(&sqlite.Item{})

	if result.Error != nil {
		println("GORM Error:", result.Error.Error())
		SendResponse(c, "error", "database error", nil)
		return
	}

	if result.RowsAffected == 0 {
		SendResponse(c, "error", "item not found in this list", nil)
		return
	}

	SendResponse(c, "success", "item deleted successfully", nil)
}

// ValidateItemID: Toggles the validated status of an item by ID in a specific list
func ValidateItemID(c *gin.Context) {
	listID, _ := strconv.Atoi(c.Param("idList"))
	itemID, err := strconv.Atoi(c.Param("idItem"))

	if err != nil || listID <= 0 || itemID <= 0 {
		SendResponse(c, "error", "invalid IDs", nil)
		return
	}

	var item sqlite.Item
	if err := sqlite.DB.Select("id", "validated").Where("id = ? AND list_id = ?", itemID, listID).First(&item).Error; err != nil {
		SendResponse(c, "error", "item not found in this list", nil)
		return
	}

	newStatus := !item.Validated

	if err := sqlite.DB.Model(&item).UpdateColumn("validated", newStatus).Error; err != nil {
		println("GORM Error:", err.Error())
		SendResponse(c, "error", "database update failed", nil)
		return
	}

	item.Validated = newStatus
	SendResponse(c, "success", "item toggled successfully", item)
}
