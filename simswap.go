// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/stainless-sdks/camara-go/option"
)

// SimswapService contains methods and other services that help with interacting
// with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimswapService] method instead.
type SimswapService struct {
	Options       []option.RequestOption
	Subscriptions SimswapSubscriptionService
}

// NewSimswapService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSimswapService(opts ...option.RequestOption) (r SimswapService) {
	r = SimswapService{}
	r.Options = opts
	r.Subscriptions = NewSimswapSubscriptionService(opts...)
	return
}
