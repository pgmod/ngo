package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/pgmod/ngo/internal/module"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init [параметры]",
	Short: "Инициализация с параметрами",
	Args:  cobra.MinimumNArgs(1), // Требуется хотя бы один аргумент после init
	Run: func(cmd *cobra.Command, args []string) {
		if rm {
			fmt.Println("Удаление папки", outputDir)
			os.RemoveAll(outputDir)
		}
		conf := module.InitConfig{
			OutputDir: outputDir,
			FullName:  fullName,
			Modules:   args,
			Name:      name,
		}
		module.InitDefault(conf)
		for _, arg := range args {
			modFunc, ok := module.Modules[arg]
			if !ok {
				fmt.Printf("Модуль %s не найден\n", arg)
				continue
			}
			deferFunc, ok := module.Modules["defer."+arg]
			if ok {
				defer func() {
					err := deferFunc(conf)
					if err != nil {
						fmt.Printf("Ошибка при завершении модуля %s: %s\n", arg, err)
					}
				}()
			}
			err := modFunc(conf)
			if err != nil {
				fmt.Printf("Ошибка при инициализации модуля %s: %s\n", arg, err)
				continue
			}
		}
	},
}

var outputDir string
var name string
var fullName string
var rm bool

func init() {
	initCmd.Flags().StringVarP(&outputDir, "output", "o", "./", "Папка для вывода")
	initCmd.Flags().StringVarP(&fullName, "fullname", "f", "github.com/user/app", "Полное имя проекта")
	initCmd.Flags().BoolVarP(&rm, "rm", "", false, "Удалить папку перед работой")
	if fullName != "" {
		parts := strings.Split(fullName, "/")
		name = parts[len(parts)-1]
	}
	rootCmd.AddCommand(initCmd)
}
