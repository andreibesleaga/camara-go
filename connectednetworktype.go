// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/stainless-sdks/camara-go/option"
)

// ConnectednetworktypeService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectednetworktypeService] method instead.
type ConnectednetworktypeService struct {
	Options []option.RequestOption
	// Connected Network Type Subscriptions
	Subscriptions ConnectednetworktypeSubscriptionService
}

// NewConnectednetworktypeService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewConnectednetworktypeService(opts ...option.RequestOption) (r ConnectednetworktypeService) {
	r = ConnectednetworktypeService{}
	r.Options = opts
	r.Subscriptions = NewConnectednetworktypeSubscriptionService(opts...)
	return
}
