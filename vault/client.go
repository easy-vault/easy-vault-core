package vault

import (
	"context"
	"easy-vault/config"
	"fmt"
	"log"

	vault "github.com/hashicorp/vault/api"
	"github.com/hashicorp/vault/api/auth/approle"
)

type Vault struct {
	client     *vault.Client
	parameters config.VaultParameters
}

// NewVaultAppRoleClient logs in to Vault using the AppRole authentication
// method, returning an authenticated client and the auth token itself, which
// can be periodically renewed.
func NewVaultAppRoleClient(ctx context.Context, parameters config.VaultParameters) (*Vault, *vault.Secret, error) {
	log.Printf("connecting to vault @ %s", parameters.Address)

	config := vault.DefaultConfig() // modify for more granular configuration
	config.Address = parameters.Address

	client, err := vault.NewClient(config)
	if err != nil {
		return nil, nil, fmt.Errorf("unable to initialize vault client: %w", err)
	}

	vault := &Vault{
		client:     client,
		parameters: parameters,
	}

	token, err := vault.login(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("vault login error: %w", err)
	}

	log.Println("connecting to vault: success!")

	return vault, token, nil
}

// A combination of a RoleID and a SecretID is required to log into Vault
// with AppRole authentication method. The SecretID is a value that needs
// to be protected, so instead of the app having knowledge of the SecretID
// directly, we have a trusted orchestrator (simulated with a script here)
// give the app access to a short-lived response-wrapping token.
//
// ref: https://www.vaultproject.io/docs/concepts/response-wrapping
// ref: https://learn.hashicorp.com/tutorials/vault/secure-introduction?in=vault/app-integration#trusted-orchestrator
// ref: https://learn.hashicorp.com/tutorials/vault/approle-best-practices?in=vault/auth-methods#secretid-delivery-best-practices
func (v *Vault) login(ctx context.Context) (*vault.Secret, error) {
	log.Printf("logging in to vault with approle auth; role id: %s", v.parameters.ApproleRoleID)

	approleSecretID := &approle.SecretID{
		FromString: v.parameters.ApproleSecretID,
	}

	appRoleAuth, err := approle.NewAppRoleAuth(
		v.parameters.ApproleRoleID,
		approleSecretID,
		approle.WithMountPath("approle"), // only required if the SecretID is response-wrapped
	)
	if err != nil {
		return nil, fmt.Errorf("unable to initialize approle authentication method: %w", err)
	}

	v.client.SetNamespace(v.parameters.Namespace)
	authInfo, err := v.client.Auth().Login(ctx, appRoleAuth)
	if err != nil {
		return nil, fmt.Errorf("unable to login using approle auth method: %w", err)
	}
	if authInfo == nil {
		return nil, fmt.Errorf("no approle info was returned after login")
	}

	log.Println("logging in to vault with approle auth: success!")

	return authInfo, nil
}

// GetSecretAPIKey fetches the latest version of secret api key from kv-v2
func (v *Vault) GetSecretAPIKey(ctx context.Context) (map[string]interface{}, error) {
	log.Printf("getting secret api key from vault %v", v.parameters.Secrets)

	secrets := make(map[string]interface{})
	for _, secretConfig := range v.parameters.Secrets {

		kv2 := v.client.KVv2(secretConfig.Engine)
		ctx := context.Background()
		secret, err := kv2.Get(ctx, secretConfig.Path)
		if err != nil {
			fmt.Printf("Error while fetching secrets %s : %v", fmt.Sprintf("%s/%s", secretConfig.Engine, secretConfig.Path), err)
		}

		for k, v := range secret.Data {
			secrets[k] = v
		}

	}
	return secrets, nil
}
