package module

import "ngo/internal/templates"

func InitDocker(conf InitConfig) error {
	err := templates.InsertTemplate("docker", map[string]string{
		"name": conf.Name,
	})
	if err != nil {
		return err
	}
	return nil
}
