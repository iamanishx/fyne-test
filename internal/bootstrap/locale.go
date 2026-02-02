package bootstrap

import "os"

func init() {
	fixLocaleEnv("LANG")
	fixLocaleEnv("LC_ALL")
	fixLocaleEnv("LC_CTYPE")
}

func fixLocaleEnv(key string) {
	val := os.Getenv(key)
	if val == "" || val == "C" || val == "POSIX" {
		_ = os.Setenv(key, "en_US.UTF-8")
	}
}
