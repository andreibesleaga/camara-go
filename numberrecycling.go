// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/camara-go/internal/apijson"
	"github.com/stainless-sdks/camara-go/internal/requestconfig"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/stainless-sdks/camara-go/packages/param"
	"github.com/stainless-sdks/camara-go/packages/respjson"
)

// Number Recycling
//
// NumberrecyclingService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNumberrecyclingService] method instead.
type NumberrecyclingService struct {
	Options []option.RequestOption
}

// NewNumberrecyclingService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNumberrecyclingService(opts ...option.RequestOption) (r NumberrecyclingService) {
	r = NumberrecyclingService{}
	r.Options = opts
	return
}

// Check whether the subscriber of the phone number has changed.
func (r *NumberrecyclingService) CheckSubscriberChange(ctx context.Context, params NumberrecyclingCheckSubscriberChangeParams, opts ...option.RequestOption) (res *NumberrecyclingCheckSubscriberChangeResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "numberrecycling/check"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type NumberrecyclingCheckSubscriberChangeResponse struct {
	// Set to true (Boolean, not string) when there has been a change in the subscriber
	// associated with the specific phone number after “specifiedDate”.
	PhoneNumberRecycled bool `json:"phoneNumberRecycled" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumberRecycled respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r NumberrecyclingCheckSubscriberChangeResponse) RawJSON() string { return r.JSON.raw }
func (r *NumberrecyclingCheckSubscriberChangeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NumberrecyclingCheckSubscriberChangeParams struct {
	// Specified date to check whether there has been a change in the subscriber
	// associated with the specific phone number, in RFC 3339 calendar date format
	// (YYYY-MM-DD).
	SpecifiedDate time.Time `json:"specifiedDate" api:"required" format:"date"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r NumberrecyclingCheckSubscriberChangeParams) MarshalJSON() (data []byte, err error) {
	type shadow NumberrecyclingCheckSubscriberChangeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NumberrecyclingCheckSubscriberChangeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
