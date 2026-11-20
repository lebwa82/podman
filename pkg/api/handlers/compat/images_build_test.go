//go:build !remote && (linux || freebsd)

package compat

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateBuildOptionsUsesServerTmpDir(t *testing.T) {
	tmpDir := t.TempDir()
	t.Setenv("TMPDIR", tmpDir)

	for _, path := range []string{"/v1.41/build", "/v6.0.0/libpod/build"} {
		t.Run(path, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, path, nil)
			options, cleanup, err := createBuildOptions(&BuildQuery{}, &BuildContext{}, nil, request)
			if cleanup != nil {
				t.Cleanup(cleanup)
			}
			require.NoError(t, err)
			require.Equal(t, tmpDir, options.SystemContext.BigFilesTemporaryDir)
		})
	}
}
