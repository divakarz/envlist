package envlist

import (
	"fmt"
	"os"
	"strings"
)

// ListEnvVars lists environment variables based on the option provided ("all" or "keys").
func ListEnvVars(option string) {
	switch option {
	case "all":
		for _, env := range os.Environ() {
			fmt.Println(env)
		}
	case "keys":
		for _, env := range os.Environ() {
			pair := strings.SplitN(env, "=", 2)
			fmt.Println(pair[0])
		}
	default:
		fmt.Println("Invalid option. Use 'all' to list all environment variables or 'keys' to list only the keys.")
	}
}
