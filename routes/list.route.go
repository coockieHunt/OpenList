package routes

import (
    "OpenList/sqlite"
    "strconv"
    "github.com/gin-gonic/gin"
)

func GetAllLists(c *gin.Context) {
    var lists []sqlite.List

	//preload for get items with list content
    if err := sqlite.DB.Preload("Items").Find(&lists).Error; err != nil {
        SendResponse(c, "error", "Erreur lors de la récupération", nil)
        return
    }

    SendResponse(c, "success", "Lists retrieved successfully", lists)
}

func GetListByID(c *gin.Context) {
	listStr := c.Param("id")
	listID, err := strconv.ParseUint(listStr, 10, 32)
	if err != nil {
		SendResponse(c, "error", "ID invalide", nil)
		return
	}

	var list sqlite.List
	if err := sqlite.DB.Preload("Items").First(&list, listID).Error; err != nil {
		SendResponse(c, "error", "List not found", nil)
		return
	}

	SendResponse(c, "success", "List retrieved successfully", list)
}

func NewList(c *gin.Context) {
    var newList sqlite.List

    if err := c.ShouldBindJSON(&newList); err != nil {
        SendResponse(c, "error", "Invalid request payload", nil)
        return
    }

    if newList.Title == "" {
        SendResponse(c, "error", "Title is required", nil)
        return
    }

	if len(newList.Items) > 0 {
		if newList.Items[0].Name == "" {
            newList.Items = []sqlite.Item{}
		}
	} else {
        newList.Items = []sqlite.Item{}
	}

    if err := sqlite.DB.Create(&newList).Error; err != nil {
        SendResponse(c, "error", err.Error(), nil)
        return
    }

    SendResponse(c, "success", "New list created successfully", newList)
}

func DeleteList(c *gin.Context) {
	listStr := c.Param("id")
	listID, err := strconv.ParseUint(listStr, 10, 32)
	if err != nil {
		SendResponse(c, "error", "id error", nil)
		return
	}

    if err := sqlite.DB.Delete(&sqlite.List{}, listID).Error; err != nil {
		SendResponse(c, "error", "error delete list db", nil)
		return
	}

	SendResponse(c, "success", "List deleted successfully", nil)
}