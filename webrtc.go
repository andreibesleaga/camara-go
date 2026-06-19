// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"github.com/andreibesleaga/camara-go/option"
)

// WebrtcService contains methods and other services that help with interacting
// with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebrtcService] method instead.
type WebrtcService struct {
	Options []option.RequestOption
	// WebRTC Call Handling
	Sessions WebrtcSessionService
}

// NewWebrtcService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewWebrtcService(opts ...option.RequestOption) (r WebrtcService) {
	r = WebrtcService{}
	r.Options = opts
	r.Sessions = NewWebrtcSessionService(opts...)
	return
}
