package module

import "os"

type InitConfig struct {
	OutputDir string
	FullName  string
	Modules   []string
	Name      string
}

func InitDefault(conf InitConfig) {
	os.MkdirAll(conf.OutputDir, 0755)
	os.Chdir(conf.OutputDir)
}

var Modules = map[string]func(conf InitConfig) error{
	"go":        InitGo,
	"vscode":    InitVscode,
	"air":       InitAir,
	"git":       InitGit,
	"defer.git": DeferGit,
	"docker":    InitDocker,
	"gitlab-ci": InitGitlabCI,
	"guhbap": func(conf InitConfig) error {
		conf.Modules = append(conf.Modules, "go", "vscode", "air", "git", "defer.git", "docker", "gitlab-ci")
		err := InitGo(conf)
		if err != nil {
			return err
		}
		err = InitVscode(conf)
		if err != nil {
			return err
		}
		err = InitAir(conf)
		if err != nil {
			return err
		}
		err = InitGit(conf)
		if err != nil {
			return err
		}
		err = DeferGit(conf)
		if err != nil {
			return err
		}
		err = InitDocker(conf)
		if err != nil {
			return err
		}
		err = InitGitlabCI(conf)
		if err != nil {
			return err
		}
		err = DeferGit(conf)
		if err != nil {
			return err
		}
		return nil
	},
}
