// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/stainless-sdks/camara-go/option"
)

// DeviceroamingstatusService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceroamingstatusService] method instead.
type DeviceroamingstatusService struct {
	Options       []option.RequestOption
	Subscriptions DeviceroamingstatusSubscriptionService
}

// NewDeviceroamingstatusService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceroamingstatusService(opts ...option.RequestOption) (r DeviceroamingstatusService) {
	r = DeviceroamingstatusService{}
	r.Options = opts
	r.Subscriptions = NewDeviceroamingstatusSubscriptionService(opts...)
	return
}
