// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/stainless-sdks/camara-go/option"
)

// DevicelocationService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicelocationService] method instead.
type DevicelocationService struct {
	Options []option.RequestOption
	// Device Geofencing Subscriptions
	Subscriptions DevicelocationSubscriptionService
}

// NewDevicelocationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDevicelocationService(opts ...option.RequestOption) (r DevicelocationService) {
	r = DevicelocationService{}
	r.Options = opts
	r.Subscriptions = NewDevicelocationSubscriptionService(opts...)
	return
}
