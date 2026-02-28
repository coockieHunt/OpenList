package routes

import (
	"OpenList/sqlite"
	"strconv"
	"github.com/gin-gonic/gin"
)

func AddItem(c *gin.Context) {
	var item sqlite.Item
    listStr := c.Param("id")
    
    listID, err := strconv.ParseUint(listStr, 10, 32)
	if err != nil {
		SendResponse(c, "error", "ID invalide", nil)
		return
	}

    if err := c.ShouldBindJSON(&item); err != nil {
		println("Erreur Binding JSON:", err.Error())
		SendResponse(c, "error", "Payload invalide: "+err.Error(), nil)
		return
	}

	item.ListID = uint(listID)

    if err := sqlite.DB.Create(&item).Error; err != nil {
		println("Erreur GORM:", err.Error())
		SendResponse(c, "error", "Erreur DB: "+err.Error(), nil)
		return
	}

	SendResponse(c, "success", "Item ajouté", item)
}

func DeleteItem(c *gin.Context) {
	itemStr := c.Param("id")
	itemID, err := strconv.ParseUint(itemStr, 10, 32)
	if err != nil {
		SendResponse(c, "error", "id invalid", nil)
		return
	}

	if err := sqlite.DB.Delete(&sqlite.Item{}, itemID).Error; err != nil {
		SendResponse(c, "error", "error delete item db", nil)
		return
	}

	SendResponse(c, "success", "Item deleted successfully", nil)
}

func ValidateItemID(c *gin.Context){
	itemStr := c.Param("id")
	listID, err := strconv.ParseUint(itemStr, 10, 32)
	if err != nil {
		SendResponse(c, "error", "id invalid", nil)
		return
	}

	type validatePayload struct {
		ID uint `json:"id"`
	}

	var payload validatePayload
	if c.Request.ContentLength > 0 {
		if err := c.ShouldBindJSON(&payload); err != nil {
			SendResponse(c, "error", "Payload invalide", nil)
			return
		}
	}

	itemIDStr := c.Query("item_id")

	var item sqlite.Item
	if payload.ID > 0 {
		if err := sqlite.DB.Where("id = ? AND list_id = ?", payload.ID, uint(listID)).First(&item).Error; err != nil {
			SendResponse(c, "error", "Item not found", nil)
			return
		}
	} else if itemIDStr != "" {
		itemID, parseErr := strconv.ParseUint(itemIDStr, 10, 32)
		if parseErr != nil {
			SendResponse(c, "error", "item_id invalid", nil)
			return
		}

		if err := sqlite.DB.Where("id = ? AND list_id = ?", uint(itemID), uint(listID)).First(&item).Error; err != nil {
			SendResponse(c, "error", "Item not found", nil)
			return
		}
	} else {
		if err := sqlite.DB.Where("list_id = ?", uint(listID)).Order("id desc").First(&item).Error; err != nil {
			SendResponse(c, "error", "Item not found", nil)
			return
		}
	}

	item.Validated = !item.Validated

	if err := sqlite.DB.Save(&item).Error; err != nil {
		SendResponse(c, "error", "error update item db", nil)
		return
	}

	SendResponse(c, "success", "Item updated successfully", item)
}
