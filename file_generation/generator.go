package filegeneration

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/a-h/templ"
)

func CreateStaticFiles(routePath string, component templ.Component) {
	rootPath := "public"
	err := os.Mkdir(rootPath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create output directory %v", err)
	}
	filePath := getFileName(rootPath, routePath)

	f, err := os.Create(filePath)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create file %v", err)
	}

	err = component.Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to render component: %v", err)
	}

	assetsPath := filepath.Join(rootPath, "assets")
	err = os.Mkdir(assetsPath, 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatalf("failed to create assets output directory %v", err)
	}

	tailwindPath := filepath.Join(assetsPath, "tailwind.css")
	err = Copy("./assets/tailwind.css", tailwindPath)
	if err != nil {
		log.Fatalf("failed to copy asset files: %v", err)
	}
}

func getFileName(outputDirPath, route string) string {
	if route == "/" {
		return filepath.Join(outputDirPath, "index.html")
	}
	cleanPath := strings.TrimPrefix(route, "/")
	return filepath.Join(outputDirPath, cleanPath+".html")
}
