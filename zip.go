// Package zipp provides functions to create and extract .zip archives.
// Compatible with cross-platform systems.
package zipp

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// Pack creates a .zip archive from the source directory (sourceDir)
// and saves it to the specified targetArchive path.
func Pack(sourceDir, targetArchive string) error {
	// Check if the source directory exists
	info, err := os.Stat(sourceDir)
	if err != nil {
		return fmt.Errorf("source directory does not exist: %v", err)
	}

	if !info.IsDir() {
		return fmt.Errorf("source is not a directory")
	}

	// Create a file to write the archive
	zipFile, err := os.Create(targetArchive)
	if err != nil {
		return fmt.Errorf("could not create archive file: %v", err)
	}
	defer zipFile.Close()

	// Initialize zip writer
	zipWriter := zip.NewWriter(zipFile)
	defer func() {
		if err := zipWriter.Close(); err != nil {
			fmt.Printf("error closing zip writer: %v\n", err)
		}
	}()

	// Walk through the source directory
	err = filepath.Walk(sourceDir, func(filePath string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip root dir
		if filePath == sourceDir {
			return nil
		}

		// Determine the relative path for the file
		relPath, err := filepath.Rel(sourceDir, filePath)
		if err != nil {
			return err
		}
		relPath = filepath.ToSlash(relPath)

		// Create a zip file header
		header, err := zip.FileInfoHeader(fileInfo)
		if err != nil {
			return err
		}
		header.Name = relPath
		if fileInfo.Mode().IsDir() {
			header.Method = zip.Store // Use Store method for directories
		} else {
			header.Method = zip.Deflate // Use Deflate method for files
		}

		// Write the file header to the zip archive
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Write the file content to the zip archive
		if !fileInfo.Mode().IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

// Unpack extracts a .zip archive (sourceArchive) into the target directory (targetDir).
func Unpack(sourceArchive, targetDir string) error {
	// Open zip archive for reading
	r, err := zip.OpenReader(sourceArchive)
	if err != nil {
		return fmt.Errorf("could not open archive file: %v", err)
	}
	defer r.Close()

	if err := os.MkdirAll(targetDir, 0755); err != nil {
		return err
	}

	// Iterate over each file in the archive
	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		path := filepath.Join(targetDir, f.Name)

		if f.FileInfo().IsDir() {
			if err := os.MkdirAll(path, f.Mode()); err != nil {
				return err
			}
		} else {
			if err := os.MkdirAll(filepath.Dir(path), os.ModePerm); err != nil {
				return err
			}
			outFile, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return err
			}
			defer outFile.Close()

			_, err = io.Copy(outFile, rc)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
