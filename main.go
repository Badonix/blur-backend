package main

import (
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/badonix/blur-backend/image"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

type imgToBlur struct {
	Radius float64               `form:"radius" binding:"required"`
	Image  *multipart.FileHeader `form:"image" binding:"required"`
}

func main() {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	router.Static("/results", "./results")
	router.POST("/blur", blurImage)

	router.Run("localhost:8080")

}

func blurImage(c *gin.Context) {
	var newImage imgToBlur

	if err := c.Bind(&newImage); err != nil {
		return
	}
	file, _ := newImage.Image.Open()
	defer file.Close()

	filename := filepath.Base(newImage.Image.Filename)
	filepath := filepath.Join("./uploads", filename)

	out, err := os.Create(filepath)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create file"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to write file"})
		return
	}
	name := image.BlurImage("./"+filename, newImage.Radius)
	c.IndentedJSON(http.StatusCreated, name)
}
