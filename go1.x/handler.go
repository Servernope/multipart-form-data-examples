package function

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	// Save multipart/form-data file to disk
	fileName, err := FileUpload(r, "image", "/tmp/")

	if err != nil {
		response := map[string]interface{}{
			"message": "Failed to save file",
			"error":   err.Error(),
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Serve the file
	http.ServeFile(w, r, "/tmp/"+fileName)
}

func FileUpload(r *http.Request, fileParameter string, pathToSave string) (string, error) {
	// ParseMultipartForm parses a request body as multipart/form-data
	r.ParseMultipartForm(32 << 20)

	// Retrieve the file from form data
	file, handler, err := r.FormFile(fileParameter)

	if err != nil {
		return "", err
	}
	// Close the file when we finish
	defer file.Close()

	// This is the path which we want to store the file
	f, err := os.OpenFile(pathToSave+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		return "", err
	}

	// Copy the file to the destination path
	io.Copy(f, file)

	return handler.Filename, nil
}
