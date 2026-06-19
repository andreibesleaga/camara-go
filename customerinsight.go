// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/andreibesleaga/camara-go/option"
)

// CustomerinsightService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerinsightService] method instead.
type CustomerinsightService struct {
	Options []option.RequestOption
	// Customer Insights
	Scoring CustomerinsightScoringService
}

// NewCustomerinsightService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomerinsightService(opts ...option.RequestOption) (r CustomerinsightService) {
	r = CustomerinsightService{}
	r.Options = opts
	r.Scoring = NewCustomerinsightScoringService(opts...)
	return
}
