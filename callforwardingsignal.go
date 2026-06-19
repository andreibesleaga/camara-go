// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/camara-go/internal/apijson"
	shimjson "github.com/stainless-sdks/camara-go/internal/encoding/json"
	"github.com/stainless-sdks/camara-go/internal/requestconfig"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/stainless-sdks/camara-go/packages/param"
	"github.com/stainless-sdks/camara-go/packages/respjson"
)

// Call Forwarding Signal
//
// CallforwardingsignalService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCallforwardingsignalService] method instead.
type CallforwardingsignalService struct {
	Options []option.RequestOption
}

// NewCallforwardingsignalService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCallforwardingsignalService(opts ...option.RequestOption) (r CallforwardingsignalService) {
	r = CallforwardingsignalService{}
	r.Options = opts
	return
}

// This endpoint provides information about which type of call forwarding service
// is active. More than one service can be active, e.g. conditional and
// unconditional. This endpoint exceeds the main scope of the Call Forwarding
// Signal API, for this reason an error code 501 can be returned.
func (r *CallforwardingsignalService) CheckActiveForwardings(ctx context.Context, params CallforwardingsignalCheckActiveForwardingsParams, opts ...option.RequestOption) (res *[]CallforwardingsignalCheckActiveForwardingsResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "callforwardingsignal/call-forwardings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// This endpoint provides information about the status of the unconditional call
// forwarding, being active or not.
func (r *CallforwardingsignalService) CheckUnconditionalForwarding(ctx context.Context, params CallforwardingsignalCheckUnconditionalForwardingParams, opts ...option.RequestOption) (res *CallforwardingsignalCheckUnconditionalForwardingResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "callforwardingsignal/unconditional-call-forwardings"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// resource containing the phone number (PhoneNumber) regarding which the Call
// Forwarding Service must be checked. To be provided/valued only in case of
// two-legged authentication. If provided/valued with three-legged authentication a
// 422-UNNECESSARY_IDENTIFIER error code is returned.
type CreateCallForwardingSignalParam struct {
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	paramObj
}

func (r CreateCallForwardingSignalParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateCallForwardingSignalParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateCallForwardingSignalParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallforwardingsignalCheckActiveForwardingsResponse string

const (
	CallforwardingsignalCheckActiveForwardingsResponseInactive                CallforwardingsignalCheckActiveForwardingsResponse = "inactive"
	CallforwardingsignalCheckActiveForwardingsResponseUnconditional           CallforwardingsignalCheckActiveForwardingsResponse = "unconditional"
	CallforwardingsignalCheckActiveForwardingsResponseConditionalBusy         CallforwardingsignalCheckActiveForwardingsResponse = "conditional_busy"
	CallforwardingsignalCheckActiveForwardingsResponseConditionalNotReachable CallforwardingsignalCheckActiveForwardingsResponse = "conditional_not_reachable"
	CallforwardingsignalCheckActiveForwardingsResponseConditionalNoAnswer     CallforwardingsignalCheckActiveForwardingsResponse = "conditional_no_answer"
)

// resource containing the information about the Unconditional Call Forwarding
// Service for the given phone number (PhoneNumber)
type CallforwardingsignalCheckUnconditionalForwardingResponse struct {
	// Indicates if the unconditional call forwarding service is active.
	Active bool `json:"active"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CallforwardingsignalCheckUnconditionalForwardingResponse) RawJSON() string { return r.JSON.raw }
func (r *CallforwardingsignalCheckUnconditionalForwardingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallforwardingsignalCheckActiveForwardingsParams struct {
	// resource containing the phone number (PhoneNumber) regarding which the Call
	// Forwarding Service must be checked. To be provided/valued only in case of
	// two-legged authentication. If provided/valued with three-legged authentication a
	// 422-UNNECESSARY_IDENTIFIER error code is returned.
	CreateCallForwardingSignal CreateCallForwardingSignalParam
	XCorrelator                param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r CallforwardingsignalCheckActiveForwardingsParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCallForwardingSignal)
}
func (r *CallforwardingsignalCheckActiveForwardingsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CallforwardingsignalCheckUnconditionalForwardingParams struct {
	// resource containing the phone number (PhoneNumber) regarding which the Call
	// Forwarding Service must be checked. To be provided/valued only in case of
	// two-legged authentication. If provided/valued with three-legged authentication a
	// 422-UNNECESSARY_IDENTIFIER error code is returned.
	CreateCallForwardingSignal CreateCallForwardingSignalParam
	XCorrelator                param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r CallforwardingsignalCheckUnconditionalForwardingParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateCallForwardingSignal)
}
func (r *CallforwardingsignalCheckUnconditionalForwardingParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
