package services

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"server/app/firebase"
	"server/app/models"
	"slices"
)

type FileOptions struct {
	Path              string
	ExtensionsAllowed []string
}

// UploadFile uploads a file to Firebase Cloud Storage and returns a BaseFile model.
//
// Parameters:
//   - store: A pointer to the Firebase CloudStorage instance.
//   - file: A pointer to the multipart.FileHeader containing the file to be uploaded.
//   - options: FileOptions struct containing the upload path and allowed extensions.
//
// Returns:
//   - models.BaseFile: A BaseFile struct containing information about the uploaded file.
//   - error: An error if the upload fails, nil otherwise.
func UploadFile(store *firebase.CloudStorage, file *multipart.FileHeader, options FileOptions) (models.BaseFile, error) {

	var baseFile models.BaseFile
	attrs, err := store.UploadFile(file, options.Path)
	if err != nil {
		return baseFile, err
	}

	ext := filepath.Ext(file.Filename)

	if options.ExtensionsAllowed != nil && len(options.ExtensionsAllowed) > 0 {
		if !slices.Contains(options.ExtensionsAllowed, ext) {
			return baseFile, errors.New("file extension not allowed")
		}
	}

	baseFile = models.BaseFile{
		FileName:     attrs.Name,
		FileUrl:      attrs.MediaLink,
		Extension:    models.FileExtension(ext),
		UserFileName: file.Filename,
	}

	return baseFile, nil
}
