package main

import (
    "fmt"
    "log"
    "os"

    "github.com/hashicorp/vault/api"
)

func main() {
    // Retrieve the Vault address and token from environment variables
    vaultAddr := os.Getenv("VAULT_ADDR")
    vaultToken := os.Getenv("VAULT_TOKEN")

    // Create a new Vault client
    client, err := api.NewClient(&api.Config{
        Address: vaultAddr,
    })
    if err != nil {
        log.Fatal(err)
    }

    // Set the Vault token
    client.SetToken(vaultToken)

    // Retrieve the secret from Vault
    secret, err := client.Logical().Read("secret/data/myapp")
    if err != nil {
        log.Fatal(err)
    }

    // Parse the secret data into a struct
    var data struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if secret != nil {
        if secret.Data != nil {
            err = secret.Data.DecodeJSON(&data)
            if err != nil {
                log.Fatal(err)
            }
        }
    }

    // Print the secret values
    fmt.Printf("Username: %s\n", data.Username)
    fmt.Printf("Password: %s\n", data.Password)
}
