package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/recode-sh/api/internal/envvars"
	"github.com/recode-sh/api/internal/routes"
)

func main() {
	// godotenv returns an error in prod
	// when the file ".env" is not present...
	_ = godotenv.Load()
	// ...so, we let it fail silently and ensure that
	// all env vars are still set.
	envvars.Ensure(".env.dist")

	gin.SetMode(envvars.Get(envvars.EnvVarNameGinMode))

	r := gin.Default()

	r.GET("/github/oauth/callback", routes.GitHubOAuthCallback)

	r.Run(":" + envvars.Get(envvars.EnvVarNamePort))
}
