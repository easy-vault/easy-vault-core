package config

type Environment struct {
	// Vault address, approle login credentials, and secret locations
	VaultAddress       string `env:"VAULT_ADDRESS"                 default:"localhost:8200"               description:"Vault address"                                          long:"vault-address"`
	VaultApproleRoleID string `env:"VAULT_APPROLE_ROLE_ID"         required:"true"                        description:"AppRole RoleID to log in to Vault"                      long:"vault-approle-role-id"`
}
type VaultParameters struct {
	// connection parameters
	Address         string
	ApproleRoleID   string
	ApproleSecretID string

	// the locations / field names of our two secrets
	SecretPath string
}
