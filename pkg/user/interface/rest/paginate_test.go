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

	ctx := context.WithValue(context.TODO(), contextKeyPage, "test")
	page := PageParam(ctx)
	assert.Equal(page, defaultPage, "fall back in case of wrong input")

	ctx = context.WithValue(context.TODO(), contextKeyPage, 1)
	page = PageParam(ctx)
	assert.Equal(page, 1, "we should get page parameter correctly set")
}

func TestPaginateLimitParam(t *testing.T) {
	assert := assert.New(t)

	ctx := context.WithValue(context.TODO(), contextKeyLimit, "test")
	limit := LimitParam(ctx)
	assert.Equal(limit, defaultLimit, "fall back in case of wrong input")

	ctx = context.WithValue(context.TODO(), contextKeyLimit, 1)
	limit = LimitParam(ctx)
	assert.Equal(limit, 1, "we should get page parameter correctly set")
}

func Test_PaginateMiddleware(t *testing.T) {
	type Expected struct {
		page  int
		limit int
	}
	tests := []struct {
		name     string
		query    string
		expected Expected
		page     string
		limit    string
		status   int
	}{
		{
			name:  "OK",
			query: "/?page=10&limit=100",
			expected: Expected{
				page:  10,
				limit: 100,
			},
			status: http.StatusOK,
		},
		{
			name:  "default values",
			query: "/",
			expected: Expected{
				page:  0,
				limit: 10,
			},
			status: http.StatusOK,
		},
		{
			name:  "invalid page value",
			query: "/?page=-1",
			expected: Expected{
				page:  0,
				limit: 10,
			},
			status: http.StatusBadRequest,
		},
		{
			name:  "invalid page value (not number)",
			query: "/?page=xy",
			expected: Expected{
				page:  0,
				limit: 0,
			},
			status: http.StatusBadRequest,
		},
		{
			name:  "invalid limit value",
			query: "/?limit=-1",
			expected: Expected{
				page:  0,
				limit: 0,
			},
			status: http.StatusBadRequest,
		},
		{
			name:  "invalid limit value (not number)",
			query: "/?limit=aa",
			expected: Expected{
				page:  0,
				limit: 10,
			},
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

				page := PageParam(r.Context())
				assert.Equal(tt.expected.page, page, "page parameter not in request context: got %q want %q", page, tt.expected.page)

				limit := LimitParam(r.Context())
				assert.Equal(tt.expected.limit, limit, "limit parameter not in request context: got %q want %q", limit, tt.expected.limit)
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
