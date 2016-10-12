// Package correlation provides a context-based means to correlation a transaction
// and correlate events spanning API boundaries.
package correlation

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-uuid"
)

type key int

const (
	// KeyName should be used to ensure consistent log message keying
	KeyName string = "CORRELATION-ID"
	ctxKey  key    = 0
)

// DeriveCtxWithID returns a new context with correlation IDKey added.
func DeriveCtxWithID(ctx context.Context, correlationID string) context.Context {
	return context.WithValue(ctx, ctxKey, correlationID)
}

// ID retrieves correlation ID from context.
func ID(ctx context.Context) string {
	return ctx.Value(ctxKey).(string)
}

// InitID middleware correlation ID to context from request header or GenerateUUID
func InitID(ctx context.Context, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get(KeyName)
		if correlationID == "" {
			correlationID, _ = uuid.GenerateUUID()
		}

		ctx := DeriveCtxWithID(ctx, correlationID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
