// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/andreibesleaga/camara-go/option"
)

// DevicereachabilitystatusService contains methods and other services that help
// with interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicereachabilitystatusService] method instead.
type DevicereachabilitystatusService struct {
	Options []option.RequestOption
	// Device Reachability Status Subscriptions
	Subscriptions DevicereachabilitystatusSubscriptionService
}

// NewDevicereachabilitystatusService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicereachabilitystatusService(opts ...option.RequestOption) (r DevicereachabilitystatusService) {
	r = DevicereachabilitystatusService{}
	r.Options = opts
	r.Subscriptions = NewDevicereachabilitystatusSubscriptionService(opts...)
	return
}
