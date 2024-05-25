package bws

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func GetSecret(env_var string) string {
	// Load env var for token UUID
	godotenv.Load()
	env_var_val := os.Getenv(env_var)
	BWS_ACCESS_TOKEN := os.Getenv("BWS_ACCESS_TOKEN")
	fmt.Println("ACCESS_TOKEN: " + BWS_ACCESS_TOKEN)

	// Get secret from Bitwarden secret manager
	cmd_string := "bws secret get " + env_var_val + " --access-token " + BWS_ACCESS_TOKEN + " | jq -r '.value'"
	out, err := exec.Command("bash", "-c", cmd_string).Output()
	if err != nil {
		log.Fatal(err)
	}

	return string(out)
}
