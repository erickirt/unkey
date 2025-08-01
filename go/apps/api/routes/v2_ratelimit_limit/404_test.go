package v2RatelimitLimit_test

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/unkeyed/unkey/go/apps/api/openapi"
	handler "github.com/unkeyed/unkey/go/apps/api/routes/v2_ratelimit_limit"
	"github.com/unkeyed/unkey/go/pkg/db"
	"github.com/unkeyed/unkey/go/pkg/testutil"
	"github.com/unkeyed/unkey/go/pkg/uid"
)

func TestNamespaceNotFound(t *testing.T) {
	h := testutil.NewHarness(t)

	route := &handler.Handler{
		DB:                      h.DB,
		Keys:                    h.Keys,
		Logger:                  h.Logger,
		Ratelimit:               h.Ratelimit,
		RatelimitNamespaceCache: h.Caches.RatelimitNamespace,
	}

	h.Register(route)

	rootKey := h.CreateRootKey(h.Resources().UserWorkspace.ID)

	headers := http.Header{
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", rootKey)},
	}

	// Test with non-existent namespace
	t.Run("namespace not found", func(t *testing.T) {
		t.Skip()
		nonExistentNamespace := "nonexistent_namespace"
		req := handler.Request{
			Namespace:  nonExistentNamespace,
			Identifier: "user_123",
			Limit:      100,
			Duration:   60000,
		}

		res := testutil.CallRoute[handler.Request, openapi.NotFoundErrorResponse](h, route, headers, req)
		require.Equal(t, http.StatusNotFound, res.Status)
		require.NotNil(t, res.Body)
		require.Equal(t, "https://unkey.com/docs/errors/not_found", res.Body.Error.Type)
		require.Equal(t, http.StatusNotFound, res.Body.Error.Status)
	})

	// Test with deleted namespace
	t.Run("deleted namespace", func(t *testing.T) {
		// Create a namespace and then delete it
		deletedNamespace := uid.New("test")
		ctx := context.Background()

		// Create namespace
		namespaceID := uid.New(uid.RatelimitNamespacePrefix)
		err := db.Query.InsertRatelimitNamespace(ctx, h.DB.RW(), db.InsertRatelimitNamespaceParams{
			ID:          namespaceID,
			WorkspaceID: h.Resources().UserWorkspace.ID,
			Name:        deletedNamespace,
			CreatedAt:   time.Now().UnixMilli(),
		})
		require.NoError(t, err)

		// Soft delete the namespace
		err = db.Query.SoftDeleteRatelimitNamespace(ctx, h.DB.RW(), db.SoftDeleteRatelimitNamespaceParams{
			ID:  namespaceID,
			Now: sql.NullInt64{Valid: true, Int64: time.Now().UnixMilli()},
		})
		require.NoError(t, err)

		req := handler.Request{
			Namespace:  deletedNamespace,
			Identifier: "user_123",
			Limit:      100,
			Duration:   60000,
		}

		res := testutil.CallRoute[handler.Request, openapi.NotFoundErrorResponse](h, route, headers, req)
		require.Equal(t, http.StatusNotFound, res.Status)
		require.NotNil(t, res.Body)
		require.Equal(t, "https://unkey.com/docs/errors/unkey/data/ratelimit_namespace_not_found", res.Body.Error.Type)
	})
}
