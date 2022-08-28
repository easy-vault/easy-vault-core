package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Environment struct {
	// Vault address, approle login credentials, and secret locations
	VaultAddress       string `env:"VAULT_ADDRESS"                 default:"localhost:8200"               description:"Vault address"                                          long:"vault-address"`
	VaultApproleRoleID string `env:"VAULT_APPROLE_ROLE_ID"         required:"true"                        description:"AppRole RoleID to log in to Vault"                      long:"vault-approle-role-id"`
}
type VaultParameters struct {
	// connection parameters
	Address         string `yaml:"address"`
	ApproleRoleID   string `yaml:"approleRoleId"`
	ApproleSecretID string `yaml:"approleSecretId"`

	// the locations / field names of our two secrets
	Secrets []Secret `yaml:"secretsPath"`
}

type Secret struct {
	Path   string `yaml:"path"`
	Engine string `yaml:"secretEngine"`
}

func Load(path string) VaultParameters {

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("Unable to load configuration from file :%s\n", path)
	}
	vaultParameters := VaultParameters{}
	if err = yaml.Unmarshal(data, &vaultParameters); err != nil {
		fmt.Printf("error while parsing config file=%s \n error = %v\n", path, err)
	}
	return vaultParameters
}
