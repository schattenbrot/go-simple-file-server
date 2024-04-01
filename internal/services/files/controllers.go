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
	"github.com/schattenbrot/go-simple-upload-server/packages/explerror"
	"github.com/schattenbrot/go-simple-upload-server/packages/responder"
)

func uploadFile(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		explerror.BadRequest(w, err)
		return
	}

	filename := fmt.Sprintf("%d-%s", time.Now().Unix(), fileHeader.Filename)

	fileType := filepath.Ext(fileHeader.Filename)
	if fileType != ".jpg" && fileType != ".jpeg" && fileType != ".png" {
		explerror.BadRequest(w, errors.New("wrong filetype"))
		return
	}

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

	responder.Send(w, http.StatusCreated, struct {
		Filename string `json:"filename"`
	}{
		Filename: filename,
	})
}

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

func getFiles(w http.ResponseWriter, r *http.Request) {
	workDir, err := os.Getwd()
	if err != nil {
		explerror.InternalServerError(w, err)
		return
	}
	dir := filepath.Join(workDir, "data", "files")

	type FileInfo struct {
		Filename string `json:"filename"`
		Filetype string `json:"filetype"`
	}

	var files []FileInfo = []FileInfo{}

	// Walk through the directory
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the file is not a directory and has the desired extension
		if !info.IsDir() && (filepath.Ext(info.Name()) == ".png" || filepath.Ext(info.Name()) == ".jpeg" || filepath.Ext(info.Name()) == ".jpg") {
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
