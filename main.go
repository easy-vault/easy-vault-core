package main

import (
	"context"
	"easy-vault/config"
	"easy-vault/vault"
	"fmt"
)

func main() {

	vaultParam := config.VaultParameters{
		Address:         "http://127.0.0.1:1234",
		ApproleRoleID:   "c33d9ce8-0594-e038-ff07-dec972282ca9",
		ApproleSecretID: "aef9fb14-fa3d-0b5b-a016-64658475a692",
		SecretPath:      "kv/test/demo",
	}
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
}
