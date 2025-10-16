package functions

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
	"taskManager/data"
)

func ExistDir() error {

	_, err := os.Stat(data.Route)

	if os.IsNotExist(err) {
		err = os.Mkdir("./tasks", 0755)

		if err != nil {
			return errors.New("error al crear el directorio")
		}

		file, err := os.Create(data.Route)
		if err != nil {
			return errors.New("error al crear el archivo")
		}
		defer file.Close()

		configID := data.Config{ActualID: 0}
		jsonData, err := json.MarshalIndent(configID, "", "  ")
		if err != nil {
			return errors.New("error al convertir a JSON")
		}
		err = os.WriteFile(data.ConfigRoute, jsonData, 0644)
		if err != nil {
			return errors.New("error al guardar la tarea")
		}

	}

	if err != nil {
		return errors.New("error al verificar existencia del directorio")
	}

	return nil
}

func ReadJson(route string, structData any) error {

	file, err := os.ReadFile(route)
	if err != nil {
		return errors.New("error: encontrando el directorio")
	}

	if len(file) == 0 {
		return nil
	}
	err = json.Unmarshal(file, &structData)
	if err != nil {
		return errors.New("error: decodificado el json")
	}
	return nil
}

func ChangeJson(route string, structData any) error {
	jsonBytes, err := json.MarshalIndent(structData, "", " ")
	if err != nil {
		return errors.New("error: fallo al convertir en json")
	}

	err = os.WriteFile(route, jsonBytes, 0644)
	if err != nil {
		return errors.New("error al guardar la tarea")
	}
	return nil
}

func GetID() (int, error) {
	var config data.Config
	err := ReadJson(data.ConfigRoute, &config)
	if err != nil {
		return 0, errors.New("error: extrayendo datos")
	}
	config.ActualID++
	err = ChangeJson(data.ConfigRoute, config)
	if err != nil {
		return 0, errors.New("error: actualizando ID")
	}
	return config.ActualID, nil
}

func GetIndex(id string, tasks []data.Task) (int, error) {
	idN, err := strconv.Atoi(id)
	if err != nil {
		return 0, errors.New("error en el id")
	}

	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == idN {
			return i, nil
		}
	}
	return 0, errors.New("indice no encontrado")
}
