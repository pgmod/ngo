package module

import (
	"github.com/pgmod/ngo/internal/templates"
)

func InitGitlabCI(conf InitConfig) error {
	err := templates.InsertTemplate("gitlab-ci", map[string]string{
		"repo": conf.FullName,
	})
	if err != nil {
		return err
	}
	return nil
}
