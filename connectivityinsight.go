// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/stainless-sdks/camara-go/option"
)

// ConnectivityinsightService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectivityinsightService] method instead.
type ConnectivityinsightService struct {
	Options []option.RequestOption
	// Connectivity Insights Subscriptions
	Subscriptions ConnectivityinsightSubscriptionService
}

// NewConnectivityinsightService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectivityinsightService(opts ...option.RequestOption) (r ConnectivityinsightService) {
	r = ConnectivityinsightService{}
	r.Options = opts
	r.Subscriptions = NewConnectivityinsightSubscriptionService(opts...)
	return
}
