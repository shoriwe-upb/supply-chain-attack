package router

import (
	"net/http"
	"testing"
)

func Test_GetEnv(t *testing.T) {

	t.Run("Succeed", func(t *testing.T) {
		expect, l := newExpect(t)
		defer l.Close()

		expect.
			GET(EnvRoute).
			Expect().
			Status(http.StatusOK)
	})
}
