package module

import (
	"fmt"
	"slices"

	"github.com/pgmod/ngo/internal/system"
	"github.com/pgmod/ngo/internal/templates"
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
