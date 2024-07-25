package image

import (
	"fmt"
	"github.com/anthonynsimon/bild/blur"
	"github.com/anthonynsimon/bild/imgio"
	"log"
	"strings"
	"time"
)

func BlurImage(image string, value float64) string {
	fmt.Println(image)
	imageName := strings.Split(image, "./")[1]
	img, err := imgio.Open("./uploads/" + imageName)
	if err != nil {
		log.Fatal("Image not found / cannt open")
	}
	bluredImg := blur.Box(img, value) // blur current image

	// genearte unique name
	file_arr := strings.Split(image, "./")
	actualName := strings.Split(file_arr[1], ".")[0]
	currentDate := time.Now()
	dateString := currentDate.Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("./results/%s_%s.png", actualName, dateString)

	// save image
	if err := imgio.Save(fileName, bluredImg, imgio.PNGEncoder()); err != nil {
		log.Fatal("Error while saving error", err)
	}

	return fileName
}
