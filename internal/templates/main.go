package templates

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

var templasePath = `C:\Users\guhbap\Code\pgmod\ngo\templates`

func InsertTemplate(name string, data map[string]string) error {
	templateDir := path.Join(templasePath, name)
	err := filepath.Walk(templateDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Create destination directory if it's a directory
		if info.IsDir() {
			relPath, err := filepath.Rel(templateDir, path)
			if err != nil {
				return err
			}
			if relPath == "." {
				return nil
			}
			return os.MkdirAll(relPath, 0755)
		}

		// Read template file
		content, err := os.ReadFile(path)
		if err != nil {
			return err
		}

		// Replace template variables
		strContent := string(content)
		for key, value := range data {
			strContent = strings.ReplaceAll(strContent, "{{"+key+"}}", value)
		}

		// Calculate relative path from template dir
		relPath, err := filepath.Rel(templateDir, path)
		if err != nil {
			return err
		}

		// Create destination directory if needed
		destPath := relPath
		destDir := filepath.Dir(destPath)
		if destDir != "." {
			if err := os.MkdirAll(destDir, 0755); err != nil {
				return err
			}
		}

		// Write processed file
		return os.WriteFile(destPath, []byte(strContent), 0644)
	})
	return err
}
