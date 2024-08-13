package httpbs

// CreateRoute let us to create routes for handlers
func CreateRoute(method string, route string) string {
	return method + " " + route
}
