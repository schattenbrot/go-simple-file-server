package files

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/schattenbrot/go-simple-upload-server/internal/config"
	"github.com/schattenbrot/go-simple-upload-server/packages/explerror"
	"github.com/schattenbrot/go-simple-upload-server/packages/responder"
)

// uploadResponse represents the response structure for file upload.
// swagger:model
type uploadResponse struct {
	// Filepath represents the URL where the uploaded file can be accessed.
	// Example: http://example.com/api/v1/files/filename.png
	Filepath string `json:"filepath"`
}

// @Summary Upload a file
// @Description Uploads a file
// @Tags files
// @Accept mpfd
// @Produce json
// @Success 201 {object} uploadResponse
// @Failure 400 {object} explerror.ErrorResponse
// @Router /files/ [post]
// @Security BearerAuth
func uploadFile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		explerror.BadRequest(w, err)
		return
	}

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)

	defer file.Close()
	workDir, err := os.Getwd()
	if err != nil {
		explerror.InternalServerError(w, err)
		return
	}

	filesDir := filepath.Join(workDir, "data", "files", filename)
	out, err := os.Create(filesDir)
	if err != nil {
		explerror.InternalServerError(w, err)
		return
	}
	defer out.Close()
	io.Copy(out, file)

	domain := config.Domain
	if config.Env != "prod" {
		domain = fmt.Sprintf("%s:%d", domain, config.Port)
	}
	responder.Send(w, http.StatusCreated, uploadResponse{
		Filepath: fmt.Sprintf("%s/api/v1/files/%s", domain, filename),
	})
}

// FileResponse represents a file resource.
// swagger:model
type FileResponse struct {
	// Content represents the content of the file.
	Content []byte `json:"contentAsBlob"`
}

// @Summary Get a file
// @Description Retrieves a file by filename
// @Tags files
// @Param filename path string true "File name"
// @Produce octet-stream
// @Success 200 {object} FileResponse "File content"
// @Failure 400 {object} explerror.ErrorResponse
// @Failure 404 {object} explerror.ErrorResponse
// @Router /files/{filename} [get]
// @Security BearerAuth
func getFile(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename")

	workDir, err := os.Getwd()
	if err != nil {
		explerror.InternalServerError(w, err)
		return
	}

	filesDir := filepath.Join(workDir, "data", "files", filename)

	file, err := os.Open(filesDir)
	if err != nil {
		if os.IsNotExist(err) {
			explerror.NotFound(w, errors.New("file not found"))
			return
		}
		explerror.InternalServerError(w, err)
		return
	}
	defer file.Close()

	file.WriteTo(w)
}

// FileInfo represents the response structure for a single file info.
// swagger:model
type FileInfo struct {
	// Filename represents the file's name.
	Filename string `json:"filename"`
	// Filetype represents the file's type.
	Filetype string `json:"filetype"`
}

// @Summary Get files
// @Description Retrieves a list of available files
// @Tags files
// @Produce json
// @Success 200 {array} FileInfo "List of files"
// @Failure 404 {object} explerror.ErrorResponse
// @Failure 500 {object} explerror.ErrorResponse
// @Router /files/ [get]
// @Security BearerAuth
func getFiles(w http.ResponseWriter, r *http.Request) {
	workDir, err := os.Getwd()
	if err != nil {
		explerror.InternalServerError(w, err)
		return
	}
	dir := filepath.Join(workDir, "data", "files")

	var files []FileInfo = []FileInfo{}

	// Walk through the directory
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file is not a directory and has the desired extension
		if !info.IsDir() {
			files = append(files, FileInfo{
				Filename: info.Name(),
				Filetype: filepath.Ext(info.Name()),
			})
		}
		return nil
	})
	if err != nil {
		explerror.InternalServerError(w, err)
		return
	}

	responder.Send(w, http.StatusOK, files)
}
