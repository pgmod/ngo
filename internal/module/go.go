package module

import (
	"fmt"

	"github.com/pgmod/ngo/internal/system"
	"github.com/pgmod/ngo/internal/templates"
)

func InitGo(conf InitConfig) error {
	res, code := system.Exec([]string{"go", "mod", "init", conf.FullName})
	if code != 0 {
		return fmt.Errorf("failed to init go: %s", res)
	}
	templates.InsertTemplate("go", map[string]string{})
	return nil
}
