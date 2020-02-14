package rest

type contextKey string

func (c contextKey) String() string {
	return "eventival " + string(c)
}
