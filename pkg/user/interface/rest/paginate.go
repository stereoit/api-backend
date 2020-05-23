package rest

import (
	"net/http"
	"strconv"

	validation "github.com/go-ozzo/ozzo-validation/v4"

	"github.com/go-chi/render"
	"golang.org/x/net/context"
)

const defaultPage = 0
const defaultPageStr = "0"
const defaultLimit = 10
const defaultLimitStr = "10"

var (
	contextKeyPage  = contextKey("page")
	contextKeyLimit = contextKey("limit")
)

// PageParam extracts the `page` parameter from the context,
// return `defaultPage` if not set
func PageParam(ctx context.Context) int {
	page, ok := ctx.Value(contextKeyPage).(int)
	if !ok {
		return defaultPage
	}

	return page
}

// LimitParam extracts the `page` parameter from the context,
// return default value if not set
func LimitParam(ctx context.Context) int {
	limit, ok := ctx.Value(contextKeyLimit).(int)
	if !ok {
		return defaultLimit
	}

	return limit
}

// paginate is a middleware to implement paginated request
// we support `page` and `limit` parameters
func paginate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		page := r.URL.Query().Get("page")
		if page == "" {
			page = defaultPageStr
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
			limit = defaultLimitStr
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

		ctx := context.WithValue(r.Context(), contextKeyPage, pageInt)
		ctx = context.WithValue(ctx, contextKeyLimit, limitInt)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
