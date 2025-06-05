package filegeneration

import (
	"fmt"
	"io"
	"os"
)

func Copy(src, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer dstFile.Close()
	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}
	srcInfo, err := os.Stat(src)
	if err != nil {
		return err
	}
	err = os.Chmod(dest, srcInfo.Mode())
	if err != nil {
		return err
	}
	return nil
}
