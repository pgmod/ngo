package module

import (
	"fmt"
	"ngo/internal/system"
	"ngo/internal/templates"
	"slices"
)

func InitAir(conf InitConfig) error {
	fmt.Println("Installing air...")
	out, code := system.Exec([]string{"go", "install", "github.com/air-verse/air@v1.61.7"})
	if code != 0 {
		return fmt.Errorf("failed to install air: %s", out)
	}
	fmt.Println(out)
	templates.InsertTemplate("air", map[string]string{})
	if slices.Contains(conf.Modules, "vscode") {
		templates.InsertTemplate("air.vscode", map[string]string{})
	}
	return nil
}
