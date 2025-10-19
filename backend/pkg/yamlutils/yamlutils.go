package yamlutils

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadYamlData[T any](path string, data T) error {
	file, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	err = yaml.Unmarshal(file, data)

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func SaveYamlData[T any](path string, data []T) error {
	yamlData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	log.Println("Writing data", data)
	log.Println("Path", path)
	err = os.WriteFile(path, yamlData, 0644)
	if err != nil {
		return err
	}
	return nil
}
