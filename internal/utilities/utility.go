package utilities

import "os/user"

// GetHomeDir returns current user's home directory
func GetHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}
	return usr.HomeDir
}

// CheckError checks if error is nil and report to Sentry
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
