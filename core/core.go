package core

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"
	openbao "github.com/hashicorp/vault/api"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")

	// Initialize OpenBao client
	config := openbao.DefaultConfig()
	config.Address = "http://127.0.0.1:8200"

	client, err := openbao.NewClient(config)
	if err != nil {
		logger.Error(err, "Unable to initialize OpenBao client")
		return
	}

	client.SetToken("dev-only-token")

	// Step 4: Store a secret
	secretData := map[string]interface{}{
		"password": "OpenBao123",
	}

	_, err = client.KVv2("secret").Put(context.Background(), "my-secret-password", secretData)
	if err != nil {
		logger.Error(err, "Unable to write secret")
		return
	}

	logger.Info("Secret written successfully.")

	// Step 5: Retrieve a secret
	secret, err := client.KVv2("secret").Get(context.Background(), "my-secret-password")
	if err != nil {
		logger.Error(err, "Unable to read secret")
		return
	}

	value, ok := secret.Data["password"].(string)
	if !ok {
		logger.Error(
			fmt.Errorf(
				"value type assertion failed: %T %#v",
				secret.Data["password"],
				secret.Data["password"],
			),
			"Type assertion failed",
		)
		return
	}

	if value != "OpenBao123" {
		logger.Error(
			fmt.Errorf("unexpected password value %q retrieved from openbao", value),
			"Unexpected password value",
		)
		return
	}

	logger.Info("Access granted!")

	logger.V(1).Info("Debug: Exiting Hello function")
}
