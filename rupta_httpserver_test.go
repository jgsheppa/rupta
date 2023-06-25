package rupta_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/jgsheppa/rupta"
)

func TestRoute_NotFound(t *testing.T) {
	t.Parallel()
	r := rupta.NewRouter()

	t.Run("httpServer", func(t *testing.T) {
		s := httptest.NewServer(r)
		defer s.Close()
		resp, _ := s.Client().Get(s.URL + "/notfound")

		got := resp.StatusCode
		want := http.StatusNotFound

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(want, got))
		}
	})
}

func TestRoute_InternalServerError(t *testing.T) {
	t.Parallel()
	r := rupta.NewRouter()

	r.Route(http.MethodGet, "/panic", func(w http.ResponseWriter, r *http.Request) {
		panic("something bad happened!")
	})

	t.Run("httpServer", func(t *testing.T) {
		s := httptest.NewServer(r)
		defer s.Close()
		resp, _ := s.Client().Get(s.URL + "/panic")

		got := resp.StatusCode
		want := http.StatusInternalServerError

		if !cmp.Equal(got, want) {
			t.Error(cmp.Diff(want, got))
		}
	})
}
