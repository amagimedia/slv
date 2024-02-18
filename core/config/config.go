package config

import (
	"os"
	"strings"
)

func IsAdminModeEnabled() bool {
	if adminMode == nil {
		adminMode = new(bool)
		*adminMode = strings.ToUpper(os.Getenv(envar_SLV_ADMIN_MODE_ENABLED)) == "TRUE"
	}
	return *adminMode
}

func GetEnvSecretKey() string {
	if envSecretKey == nil {
		envSecretKey = new(string)
		*envSecretKey = os.Getenv(envar_SLV_ENV_SECRET_KEY)
	}
	return *envSecretKey
}

func GetEnvSecretBinding() string {
	if envSecretBinding == nil {
		envSecretBinding = new(string)
		*envSecretBinding = os.Getenv(envar_SLV_ENV_SECRET_BINDING)
	}
	return *envSecretBinding
}

func GetEnvSecretPassword() string {
	if envSecretPassword == nil {
		envSecretPassword = new(string)
		*envSecretPassword = os.Getenv(envar_SLV_ENV_SECRET_PASSWORD)
	}
	return *envSecretPassword
}

func GetGitHTTPUser() string {
	if gitHTTPUser == nil {
		gitHTTPUser = new(string)
		*gitHTTPUser = os.Getenv(envar_SLV_GIT_HTTP_USER)
	}
	return *gitHTTPUser
}

func GetGitHTTPToken() string {
	if gitHTTPToken == nil {
		gitHTTPToken = new(string)
		*gitHTTPToken = os.Getenv(envar_SLV_GIT_HTTP_TOKEN)
		if *gitHTTPToken == "" {
			*gitHTTPToken = os.Getenv(envar_SLV_GIT_HTTP_PASS)
		}
	}
	return *gitHTTPToken
}
