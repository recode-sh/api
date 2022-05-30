package envvars

type EnvVarName string

const (
	EnvVarNamePort                    EnvVarName = "PORT"
	EnvVarNameGitHubOAuthClientID     EnvVarName = "GITHUB_OAUTH_CLIENT_ID"
	EnvVarNameGitHubOAuthClientSecret EnvVarName = "GITHUB_OAUTH_CLIENT_SECRET"
	EnvVarNameGinMode                 EnvVarName = "GIN_MODE"
)
