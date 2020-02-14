package rest

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaginatePageParam(t *testing.T) {
	assert := assert.New(t)
	ctx := context.TODO()
	ctx = context.WithValue(ctx, contextKeyPage, "1")

	page, ok := PageParam(ctx)

	assert.True(ok, "we did set the context")
	assert.Equal(page, "1", "we should get page parameter correctly set")
}

func TestPaginateLimitParam(t *testing.T) {
	assert := assert.New(t)
	ctx := context.TODO()
	ctx = context.WithValue(ctx, contextKeyLimit, "1")

	limit, ok := LimitParam(ctx)

	assert.True(ok, "we did set the context")
	assert.Equal(limit, "1", "we should get page parameter correctly set")
}

func Test_PaginateMiddleware(t *testing.T) {
	tests := []struct {
		name   string
		page   string
		limit  string
		query  string
		status int
	}{
		{
			name:   "OK",
			page:   "10",
			limit:  "100",
			query:  "/?page=10&limit=100",
			status: http.StatusOK,
		},
		{
			name:   "default values",
			page:   "0",
			limit:  "10",
			query:  "/",
			status: http.StatusOK,
		},
		{
			name:   "invalid page value",
			page:   "",
			limit:  "10",
			query:  "/?page=-1",
			status: http.StatusBadRequest,
		},
		{
			name:   "invalid page value (not number)",
			page:   "",
			limit:  "",
			query:  "/?page=xy",
			status: http.StatusBadRequest,
		},
		{
			name:   "invalid limit value",
			page:   "",
			limit:  "",
			query:  "/?limit=-1",
			status: http.StatusBadRequest,
		},
		{
			name:   "invalid limit value (not number)",
			page:   "",
			limit:  "",
			query:  "/?limit=aa",
			status: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert := assert.New(t)
			req, err := http.NewRequest("GET", tt.query, nil)
			if err != nil {
				t.Fatal(err)
			}

			testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

				page, ok := PageParam(r.Context())
				assert.True(ok, "should parse page parameter correctly")
				assert.Equal(tt.page, page, "page parameter not in request context: got %q want %q", page, tt.page)

				limit, ok := LimitParam(r.Context())
				assert.True(ok, "should parse limit parameter correctly")
				assert.Equal(tt.limit, limit, "limit parameter not in request context: got %q want %q", limit, tt.limit)
			})

			rr := httptest.NewRecorder()
			handler := paginate(testHandler)

			handler.ServeHTTP(rr, req)

			// Check the status code is what we expect.
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.status)
			}
		})
	}

}
