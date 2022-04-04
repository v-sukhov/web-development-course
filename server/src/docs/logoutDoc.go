package docs

import (
	security "resagg/security"
)

// swagger:route POST /api/protectes/logout logout-tag idOfLogoutEdnpoint
// Logout user
// responses:
//   200: logoutResponse

// swagger:parameters idOfLogoutEdnpoint
type logoutRequestWrapper struct {
	// Empty request
	// in:body
	Body security.LogoutRequest
}

// User logout success
// swagger:response logoutResponse
type logoutResponseWrapper struct {
	// in:body
	Body security.LogoutResponse
}
