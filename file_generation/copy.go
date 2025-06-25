package filegeneration

import (
	"fmt"
	"io"
	"os"
)

func copy(srcPath, destPath string) error {
	srcFile, err := os.Open(srcPath)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(destPath)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	// Ensure that the file is stored to disk instead of memory.
	err = dstFile.Sync()
	if err != nil {
		return err
	}

	srcInfo, err := os.Stat(srcPath)
	if err != nil {
		return err
	}
	err = os.Chmod(destPath, srcInfo.Mode())
	if err != nil {
		return err
	}
	return nil
}
