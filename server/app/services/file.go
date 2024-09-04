package services

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"server/app/firebase"
	"server/app/models"
	"slices"
	"sync"
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

// DeleteFiles deletes multiple files from Firebase Cloud Storage concurrently.
//
// Parameters:
//   - store: A pointer to the Firebase CloudStorage instance.
//   - files: A slice of models.BaseFile containing information about the files to be deleted.
//
// Returns:
//   - error: An error if any file deletion fails, nil if all files are deleted successfully.
//
// The function uses goroutines to delete files concurrently, improving performance for bulk deletions.
// If any errors occur during deletion, they are collected and returned as a single error message.
func DeleteFiles(store *firebase.CloudStorage, files []models.BaseFile) error {
	if len(files) == 0 {
		return nil
	}

	errChan := make(chan error, len(files))
	var wg sync.WaitGroup

	for _, file := range files {
		wg.Add(1)
		go func(file models.BaseFile) {
			defer wg.Done()
			if file.FileName == "" {
				errChan <- errors.New("empty file name")
				return
			}
			if err := store.DeleteFile(file.FileName); err != nil {
				errChan <- fmt.Errorf("failed to delete file %s: %w", file.FileName, err)
				return
			}
		}(file)
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	var errs []error
	for err := range errChan {
		if err != nil {
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return fmt.Errorf("failed to delete %d files: %v", len(errs), errs)
	}

	return nil
}
