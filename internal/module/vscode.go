package module

import (
	"ngo/internal/templates"
	"slices"
)

func InitVscode(conf InitConfig) error {
	if slices.Contains(conf.Modules, "air") {
		return nil
	}
	if err := templates.InsertTemplate("vscode", map[string]string{}); err != nil {
		return err
	}
	return nil
}
