package handlers

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// base reponse object
type FileResponse struct {
	File string `json:"file"`
	Url  string `json:"url"`
}

func UploadFiles(c *fiber.Ctx) error {
	// filter if it is multipart form
	form, err := c.MultipartForm()
	if err != nil {
		c.SendString("Form is invalid.")
		return c.SendStatus(403)
	}

	// get all files from "documents" key:
	files := form.File["file"]

	// making sure not someone that i dont one uses the bucket
	if token := form.Value["token"]; token[0] != os.Getenv("TOKEN") {
		fmt.Printf("token used was %s", token[0])
		return c.SendStatus(403)
	}

	// Loop through files:
	for _, file := range files {

		// making sure file is not that big
		if file.Size > 5000000 {
			c.SendString("file is to large!.")
			return c.SendStatus(415)
		}

		// saving content response
		fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])
		// => "tutorial.pdf" 360641 "application/pdf"

		// changing file name
		id := uuid.New().String()
		filename := fmt.Sprintf("%s_%s", id, file.Filename)

		// making sure folder exits
		err := os.MkdirAll("./files/", 0755)
		if err != nil {
			c.SendString("Error creating folder.")
			fmt.Println(err)
			return c.SendStatus(403)
		}

		// creating response object
		response := FileResponse{File: filename}
		response.Url = fmt.Sprintf("%s/%s", os.Getenv("APP_URL"), filename)

		// save the files to disk
		if err := c.SaveFile(file, fmt.Sprintf("./files/%s", filename)); err != nil {
			c.SendString("Error loading the file into the system.")
			fmt.Println(err)
			return c.SendStatus(403)
		} else {
			c.SendStatus(200)
			return c.JSON(response)
		}
	}
	return nil
}
