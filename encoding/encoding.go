package encoding

import (
	"encoding/json"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// Читаем данные из файла JSON
	data, err := os.ReadFile(j.FileInput)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, j.DockerCompose); err != nil {
		return err
	}

	yamlData, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}

	return os.WriteFile(j.FileOutput, yamlData, 0644)
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Читаем данные из файла YAML
	data, err := os.ReadFile(y.FileInput)
	if err != nil {
		return err
	}

	// Декодируем YAML в структуру
	if err := yaml.Unmarshal(data, y.DockerCompose); err != nil {
		return err
	}

	// Сериализуем структуру в JSON
	jsonData, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}

	// Записываем JSON в выходной файл
	return os.WriteFile(y.FileOutput, jsonData, 0644)
}
