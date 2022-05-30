package envvars

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Get(name EnvVarName) string {
	envVarVal, envVarExists := os.LookupEnv(string(name))

	if !envVarExists {
		log.Panicf("the env var \"%s\" doesn't exist", name)
	}

	if len(envVarVal) == 0 {
		log.Panicf("the env var \"%s\" is empty", name)
	}

	return envVarVal
}

func Ensure(distFilePath string) {
	envVarsInDistFile, err := godotenv.Read(distFilePath)

	if err != nil {
		log.Panicf("%v", err)
	}

	for envVarName := range envVarsInDistFile {
		envVarVal, envVarExists := os.LookupEnv(envVarName)

		if !envVarExists {
			log.Panicf("the env var \"%s\" doesn't exist", envVarName)
		}

		if len(envVarVal) == 0 {
			log.Panicf("the env var \"%s\" is empty", envVarName)
		}
	}
}
