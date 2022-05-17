package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	router()

}
func router() {
	fmt.Println("Router Run")
	routes := gin.Default()
	routes.POST("/", uploadFile)
	routes.DELETE("/", deleteFile)
	routes.Run("127.0.0.1:8000")
}

func uploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	ext1 := filepath.Ext(file.Filename)

	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 1000000
	milisToString := strconv.FormatInt(int64(millis), 10)

	fileUploaded := milisToString + ext1

	err := os.Mkdir("temp", 0755)
	if err != nil {
		log.Fatal(err)
	}
	c.SaveUploadedFile(file, "temp/"+fileUploaded)
	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}

type inputModel struct {
	DataInput string `json:"asdfasdf"`
}

func deleteFile(c *gin.Context) {
	var inpuAnuan inputModel

	hasil := c.PostForm("input")
	err := c.ShouldBindJSON(&inpuAnuan)

	if err != nil {

		errorMessages := []string{}
		for _, e := range err.(validator.ValidationErrors) {
			errorMessage := fmt.Sprintf("Error on filed %s, condition: %s", e.Field(), e.ActualTag())
			errorMessages = append(errorMessages, errorMessage)
		}
		c.JSON(http.StatusBadRequest, gin.H{
			"errors": errorMessages,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"hasil":    inpuAnuan.DataInput,
		"hasilnya": hasil,
	})

}
