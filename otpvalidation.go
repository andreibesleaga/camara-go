// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/camara-go/internal/apijson"
	"github.com/stainless-sdks/camara-go/internal/requestconfig"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/stainless-sdks/camara-go/packages/param"
	"github.com/stainless-sdks/camara-go/packages/respjson"
)

// One Time Password SMS
//
// OtpvalidationService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewOtpvalidationService] method instead.
type OtpvalidationService struct {
	Options []option.RequestOption
}

// NewOtpvalidationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewOtpvalidationService(opts ...option.RequestOption) (r OtpvalidationService) {
	r = OtpvalidationService{}
	r.Options = opts
	return
}

// Sends an SMS with the desired message and an OTP code to the received phone
// number.
func (r *OtpvalidationService) SendCode(ctx context.Context, params OtpvalidationSendCodeParams, opts ...option.RequestOption) (res *OtpvalidationSendCodeResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "otpvalidation/send-code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Verifies the code is valid for the received authenticationId
func (r *OtpvalidationService) ValidateCode(ctx context.Context, params OtpvalidationValidateCodeParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	path := "otpvalidation/validate-code"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, nil, opts...)
	return err
}

// Structure to provide authentication identifier
type OtpvalidationSendCodeResponse struct {
	// unique id of the verification attempt the code belongs to.
	AuthenticationID string `json:"authenticationId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthenticationID respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OtpvalidationSendCodeResponse) RawJSON() string { return r.JSON.raw }
func (r *OtpvalidationSendCodeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtpvalidationSendCodeParams struct {
	// Message template used to compose the content of the SMS sent to the phone
	// number. It must include the following label indicating where to include the
	// short code `{{code}}`
	Message string `json:"message" api:"required"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber string            `json:"phoneNumber" api:"required"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r OtpvalidationSendCodeParams) MarshalJSON() (data []byte, err error) {
	type shadow OtpvalidationSendCodeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OtpvalidationSendCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OtpvalidationValidateCodeParams struct {
	// unique id of the verification attempt the code belongs to.
	AuthenticationID string `json:"authenticationId" api:"required"`
	// temporal, short code to be validated
	Code        string            `json:"code" api:"required"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r OtpvalidationValidateCodeParams) MarshalJSON() (data []byte, err error) {
	type shadow OtpvalidationValidateCodeParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OtpvalidationValidateCodeParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
