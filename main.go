package main

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/spf13/viper"
)

type Environment struct {
	Description string `mapstructure:"description"`
	Uri         string `mapstructure:"uri"`
}
type File struct {
	Env Environment `mapstructure:"environment"`
}

func main() {
	file1 := new(File)
	file2 := new(File)

	viper1 := newViperConfig("file1.yaml")
	viperToStruct(viper1, file1)

	viper2 := newViperConfig("file2.yaml")
	viperToStruct(viper2, file2)

	result := cmp.Diff(file1, file2)
	fmt.Println(result)
}

func newViperConfig(file string) *viper.Viper {
	const basePath = "D:\\Estudos\\golang\\go-cmp"
	const yamlDEF = "yaml"

	instance := viper.New()
	instance.SetConfigType(yamlDEF)
	instance.SetConfigFile(file)
	instance.AddConfigPath(basePath)
	if err := instance.ReadInConfig(); err != nil {
		showCustomError("Problema ao instanciar o viper - ", err)
	}
	return instance
}

func showCustomError(message string, err error) error {
	return fmt.Errorf("%s %v", message, err.Error())
}

func viperToStruct(viperInstance *viper.Viper, file *File) error {
	if err := viperInstance.Unmarshal(&file); err != nil {
		return showCustomError("Problema no unmarshal do arquivo 1 - ", err)
	}
	return nil
}
