package module

import (
	"fmt"
	"ngo/internal/system"
	"ngo/internal/templates"
)

func InitGit(conf InitConfig) error {
	res, code := system.Exec([]string{"git", "init"})
	if code != 0 {
		return fmt.Errorf("failed to init git: %s", res)
	}
	templates.InsertTemplate("git", map[string]string{})
	return nil
}

func DeferGit(conf InitConfig) error {
	res, code := system.Exec([]string{"git", "add", "."})
	if code != 0 {
		return fmt.Errorf("failed to add git: %s", res)
	}
	res, code = system.Exec([]string{"git", "commit", "-m", "Initial commit"})
	if code != 0 {
		return fmt.Errorf("failed to commit git: %s", res)
	}
	return nil
}
