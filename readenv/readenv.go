package readenv

import (
	"os"
	"strings"
)

var envNames = make(map[string]string)

func RegisterEnvFile(fileName string) error {
	file, err := os.ReadFile(fileName)
	fileStr := []string{}
	if err != nil {
		// fmt.Println(err.Error())
		return err
	}

	for _, val := range file {
		if string(val) != "\r" {
			fileStr = append(fileStr, string(val))
		}
	}
	envVars := strings.Split(strings.Join(fileStr, ""), "\n")

	return envMap(envVars)
}

// TODO: Validar que en data no haya espacios vacios
func envMap(envVars []string) error {
	for _, data := range envVars {
		pairKeyValue := strings.Split(data, "=")
		envNames[pairKeyValue[0]] = pairKeyValue[1]
	}
	return registerEnvVars(envNames)
}

func registerEnvVars(envVars map[string]string) error {

	for key, value := range envVars {
		os.Setenv(key, value)
	}
	return nil
}
