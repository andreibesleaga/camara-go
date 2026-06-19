// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/andreibesleaga/camara-go/internal/apijson"
	"github.com/andreibesleaga/camara-go/internal/requestconfig"
	"github.com/andreibesleaga/camara-go/option"
	"github.com/andreibesleaga/camara-go/packages/param"
	"github.com/andreibesleaga/camara-go/packages/respjson"
)

// Device Swap
//
// DeviceswapService contains methods and other services that help with interacting
// with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceswapService] method instead.
type DeviceswapService struct {
	Options []option.RequestOption
}

// NewDeviceswapService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDeviceswapService(opts ...option.RequestOption) (r DeviceswapService) {
	r = DeviceswapService{}
	r.Options = opts
	return
}

// Check if device swap has been performed during a past period
func (r *DeviceswapService) Check(ctx context.Context, params DeviceswapCheckParams, opts ...option.RequestOption) (res *DeviceswapCheckResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "deviceswap/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get timestamp of last device swap for a mobile user account provided with phone
// number.
func (r *DeviceswapService) GetDate(ctx context.Context, params DeviceswapGetDateParams, opts ...option.RequestOption) (res *DeviceswapGetDateResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "deviceswap/retrieve-date"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type DeviceswapCheckResponse struct {
	// Indicates whether the device has been swapped during the period within the
	// provided age.
	Swapped bool `json:"swapped" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Swapped     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceswapCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceswapCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceswapGetDateResponse struct {
	// Timestamp of latest device swap performed. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	LatestDeviceChange time.Time `json:"latestDeviceChange" api:"required" format:"date-time"`
	// Timeframe in days for device change supervision for the phone number. It could
	// be valued in the response if the latest Device swap occurred before this
	// monitored period.
	MonitoredPeriod int64 `json:"monitoredPeriod"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LatestDeviceChange respjson.Field
		MonitoredPeriod    respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceswapGetDateResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceswapGetDateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceswapCheckParams struct {
	// Period in hours to be checked for device swap.
	MaxAge param.Opt[int64] `json:"maxAge,omitzero"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r DeviceswapCheckParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceswapCheckParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceswapCheckParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceswapGetDateParams struct {
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r DeviceswapGetDateParams) MarshalJSON() (data []byte, err error) {
	type shadow DeviceswapGetDateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceswapGetDateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
