package ginapiserver

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/intelops/go-common/logging"
	"github.com/kube-tarian/kad/capten/agent/gin-api-server/api"
	"github.com/kube-tarian/kad/capten/agent/internal/config"
)

// StartRestServer starts a REST server with the given API interface, service configuration, and logger.
//
// It initializes a Gin router with default middleware, registers the API handlers, and starts the server
// with the specified host, port, certificate file, and key file.
//
// Parameters:
// - rpcapi: The API interface to register the handlers with.
// - cfg: The service configuration.
// - log: The logger.
//
// Returns:
// - error: An error if the server fails to start.
func StartRestServer(rpcapi api.ServerInterface, cfg *config.SericeConfig, log logging.Logger) error {
	r := gin.Default()
	api.RegisterHandlers(r, rpcapi)

	return r.RunTLS(fmt.Sprintf("%s:%d", cfg.Host, cfg.RestPort), cfg.CertFileName, cfg.KeyFileName)
}
