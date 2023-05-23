package utils

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Si las variables de entorno no existen en el sistema, las toma de .env
func ParseEnvironment() {
	// Valido si .env existe
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		log.Printf(".env no existe. Se intentará utilizar variables globales.")
	}

	setEnvVariable("CLIENT_ID", os.Getenv("CLIENT_ID"))
	setEnvVariable("CLIENT_SECRET", os.Getenv("CLIENT_SECRET"))
	setEnvVariable("ISSUER", os.Getenv("ISSUER"))
	setEnvVariable("AWS_ACCESS_KEY_ID", os.Getenv("AWS_ACCESS_KEY_ID"))
	setEnvVariable("AWS_SECRET_ACCESS_KEY", os.Getenv("AWS_SECRET_ACCESS_KEY"))
	setEnvVariable("AWS_DEFAULT_REGION", os.Getenv("AWS_DEFAULT_REGION"))

	if os.Getenv("CLIENT_ID") == "" {
		log.Printf("Could not resolve a CLIENT_ID environment variable.")
		os.Exit(1)
	} else if os.Getenv("CLIENT_SECRET") == "" {
		log.Printf("Could not resolve a CLIENT_SECRET environment variable.")
		os.Exit(1)
	} else if os.Getenv("ISSUER") == "" {
		log.Printf("Could not resolve a ISSUER environment variable.")
		os.Exit(1)
	}
}

func setEnvVariable(env string, current string) {
	if current != "" {
		return // Si current ya está seteado, retorno
	}

	file, _ := os.Open(".env")
	defer file.Close()

	lookInFile := bufio.NewScanner(file)
	lookInFile.Split(bufio.ScanLines)

	for lookInFile.Scan() {
		if lookInFile.Text() != "" {
			parts := strings.Split(lookInFile.Text(), "=")
			key, value := parts[0], parts[1]
			if key == env {
				os.Setenv(key, value)
			}
		}
	}

}
