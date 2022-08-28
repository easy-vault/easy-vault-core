package main

import (
	"fmt"
	"os"
)

func ExportSecrets(secrets map[string]interface{}) {

	file, err := os.Create(".vault_secrets") //os.O_CREATE, fs.FileMode(777))
	if err != nil {
		fmt.Printf("Error while writting screts:%v", err)
	}
	for key, value := range secrets {

		_, err = file.WriteString(fmt.Sprintf("export %s=\"%v\"\n", key, value))
		if err != nil {
			fmt.Println(err)
		}
	}
	file.Close()
}
