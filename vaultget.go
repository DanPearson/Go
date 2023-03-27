package main

import (
    "fmt"
    "log"
    "os"

    "github.com/hashicorp/vault/api"
)

func main() {
    // Retrieve the Vault address and token from environment variables
    vaultAddr, ok := os.LookupEnv("VAULT_ADDR")
    if !ok {
        log.Fatal("Vault address not set in VAULT_ADDR environment variable")
    }
    vaultToken, ok := os.LookupEnv("VAULT_TOKEN")
    if !ok {
        log.Fatal("Vault token not set in VAULT_TOKEN environment variable")
    }

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
