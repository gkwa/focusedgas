package core

import (
	"github.com/go-logr/logr"
	openbao "github.com/openbao/openbao/api"
)

func Hello(logger logr.Logger) {
	logger.V(1).Info("Debug: Entering Hello function")
	logger.Info("Hello, World!")
	logger.V(1).Info("Debug: Exiting Hello function")
}
