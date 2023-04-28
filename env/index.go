package env

import (
	"errors"
	"github.com/gin-gonic/gin"
	"os"
	"strings"
)

// cached env variables
var (
	DevMode  bool
	UiDevUrl string
)

const (
	Version    = "0.0.1"
	dotenvFile = ".env"
)

func init() {
	// init with default values
	_ = os.Setenv("GIN_MODE", gin.ReleaseMode)
	_ = os.Setenv("ENV", "PROD")
	_ = os.Setenv("UI_DEV_URL", "")
	_ = os.Setenv("HOST", ":8080")

	if _, err := os.Stat(dotenvFile); err == nil {
		if content, err := os.ReadFile(dotenvFile); err == nil {
			lines := strings.Split(string(content), "\n")

			for _, line := range lines {
				// ignore comments (line is prefixed with #)
				if strings.HasPrefix(line, "#") {
					continue
				}

				key, value, extractErr := extractEnvVal(line)

				if extractErr != nil {
					continue
				}

				_ = os.Setenv(key, value)
			}
		}
	}

	DevMode = os.Getenv("ENV") == "DEV"
	UiDevUrl = os.Getenv("UI_DEV_URL")

	// update gin_mode in case it was already loaded
	gin.SetMode(os.Getenv("GIN_MODE"))
}

func extractEnvVal(line string) (string, string, error) {
	if !strings.Contains(line, "=") {
		return "", "", errors.New("unable to ")
	}

	separatorIndex := strings.IndexRune(line, '=')

	return line[:separatorIndex], line[separatorIndex+1:], nil
}
