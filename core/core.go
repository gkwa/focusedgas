package core

import (
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

	// Example: Write a secret to OpenBao
	secretData := map[string]interface{}{
		"foo": "bar",
	}
	_, err = client.Logical().Write("secret/data/my-secret", secretData)
	if err != nil {
		logger.Error(err, "Failed to write secret to OpenBao")
	} else {
		logger.Info("Successfully wrote secret to OpenBao")
	}

	logger.V(1).Info("Debug: Exiting Hello function")
}
