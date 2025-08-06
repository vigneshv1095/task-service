package pagination

import (
    "net/http"
    "strconv"
)

// Parse extracts limit and offset from query params.
func Parse(r *http.Request) (limit, offset int) {
    limit = 10
    offset = 0
    if l := r.URL.Query().Get("limit"); l != "" {
        if i, err := strconv.Atoi(l); err == nil && i > 0 {
            limit = i
        }
    }
    if o := r.URL.Query().Get("offset"); o != "" {
        if i, err := strconv.Atoi(o); err == nil && i >= 0 {
            offset = i
        }
    }
    return
}
