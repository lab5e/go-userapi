package apitools

import (
	"context"
	"time"

	"github.com/lab5e/go-userapi"
)

// NewAuthenticatedContext creates a context with a token authentication.
// The done function returned should be called when the operation is finished
// as usual.
func NewAuthenticatedContext(token string, timeout time.Duration) (ctx context.Context, done func()) {
	var reqCtx context.Context
	reqCtx, done = context.WithTimeout(context.Background(), 30*time.Second)

	keys := make(map[string]userapi.APIKey)
	keys["APIToken"] = userapi.APIKey{Key: token, Prefix: ""}
	ctx = context.WithValue(reqCtx, userapi.ContextAPIKeys, keys)

	return
}
