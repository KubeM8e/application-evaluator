package utils

import (
	"archive/zip"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"
)

const destinationFolder = "./tmp"

func CopyUploadedFile(fileHeader *multipart.FileHeader) string {
	src, err := fileHeader.Open()
	if err != nil {
		log.Fatalf("Could not open file: %v", err)
	}

	dst, err := os.Create(fileHeader.Filename)
	if err != nil {
		log.Fatalf("Could not create file: %v", err)
	}

	_, err = io.Copy(dst, src)
	if err != nil {
		log.Fatalf("Could not copy file: %v", err)
	}

	return fileHeader.Filename
}

func UnZipFile(sourceFile string) string {
	r, err := zip.OpenReader(sourceFile)
	if err != nil {
		log.Fatalf("Unzip - could not open reader: %v", err)
	}
	defer r.Close()

	for _, f := range r.File {
		rc, errF := f.Open()
		if errF != nil {
			log.Fatalf("Unzip - could not open file: %v", err)
		}
		defer rc.Close()

		path := filepath.Join(destinationFolder, f.Name)
		if f.FileInfo().IsDir() {
			os.MkdirAll(path, f.Mode())
		} else {
			file, errOpen := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if errOpen != nil {
				log.Fatalf("Unzip - could not open dst file: %v", err)
			}
			defer file.Close()

			_, err = io.Copy(file, rc)
			if err != nil {
				log.Fatalf("Unzip - could not copy files: %v", err)
			}
		}
	}

	return destinationFolder
}
