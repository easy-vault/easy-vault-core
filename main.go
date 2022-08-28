package main

import (
	"context"
	"easy-vault/config"
	"easy-vault/vault"
	"flag"
	"fmt"
)

var (
	configFile string
)

func main() {

	/*vaultParam := config.VaultParameters{
		Address:         "http://192.168.43.20:1234",
		ApproleRoleID:   "5c5a9055-5ab0-adba-7278-2ab21a9acabc",
		ApproleSecretID: "85030def-bed5-d967-8cb4-7b719c03352c",
		SecretsPath:      []string{"kv/test/demo"},
	}*/
	flag.StringVar(&configFile, "config", "vault.yaml", "vault config file path")

	vaultParam := config.Load(configFile)
	fmt.Print(vaultParam)
	ctx := context.Background()
	vClient, _, err := vault.NewVaultAppRoleClient(ctx, vaultParam)

	if err != nil {
		panic(err)
	}
	secrets, err := vClient.GetSecretAPIKey(ctx)
	if err != nil {
		fmt.Printf("error in reading")
		panic(err)
	}
	fmt.Println(secrets)
	ExportSecrets(secrets)
}
