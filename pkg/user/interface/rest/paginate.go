package rest

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/go-chi/render"
	"golang.org/x/net/context"
)

const defaultPage = "0"
const defaultLimit = "10"

var (
	contextKeyPage  = contextKey("page")
	contextKeyLimit = contextKey("limit")
)

// PageParam extracts the `page` parameter from the context
func PageParam(ctx context.Context) (string, bool) {
	pageStr, ok := ctx.Value(contextKeyPage).(string)
	return pageStr, ok
}

// LimitParam extracts the `page` parameter from the context
func LimitParam(ctx context.Context) (string, bool) {
	limitStr, ok := ctx.Value(contextKeyLimit).(string)
	return limitStr, ok
}

// paginate is a middleware to implement paginated request
// we support `page` and `limit` parameters
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		if page == "" {
			page = defaultPage
		}

		pageInt, err := strconv.Atoi(page)
		if err != nil {
			render.Render(w, r, ErrBadRequest(err))
			return
		}

		err = validation.Validate(pageInt, validation.Min(0))
		if err != nil {
			render.Render(w, r, ErrBadRequest(err))
			return
		}

		limit := r.URL.Query().Get("limit")
		if limit == "" {
			limit = defaultLimit
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			render.Render(w, r, ErrBadRequest(err))
			return
		}

		err = validation.Validate(limitInt, validation.Min(0))
		if err != nil {
			render.Render(w, r, ErrBadRequest(err))
			return
		}

		ctx := context.WithValue(r.Context(), contextKeyPage, page)
		ctx = context.WithValue(ctx, contextKeyLimit, limit)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
