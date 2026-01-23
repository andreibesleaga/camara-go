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

// TenureService contains methods and other services that help with interacting
// with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTenureService] method instead.
type TenureService struct {
	Options []option.RequestOption
}

// NewTenureService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewTenureService(opts ...option.RequestOption) (r TenureService) {
	r = TenureService{}
	r.Options = opts
	return
}

// Verifies a specified length of tenure, based on a provided date, for a network
// subscriber to establish a level of trust for the network subscription
// identifier.
func (r *TenureService) Verify(ctx context.Context, params TenureVerifyParams, opts ...option.RequestOption) (res *TenureVerifyResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "tenure/check-tenure"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type TenureVerifyResponse struct {
	// `true` when the identified mobile subscription has had valid tenure since
	// `tenureDate`, otherwise `false`
	TenureDateCheck bool `json:"tenureDateCheck,required"`
	// If exists, populated with:
	//
	// - `PAYG` - prepaid (pay-as-you-go) account
	// - `PAYM` - contract account
	// - `Business` - Business (enterprise) account
	//
	// This attribute may be omitted from the response set if the information is not
	// available
	//
	// Any of "PAYG", "PAYM", "Business".
	ContractType TenureVerifyResponseContractType `json:"contractType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TenureDateCheck respjson.Field
		ContractType    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TenureVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *TenureVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// If exists, populated with:
//
// - `PAYG` - prepaid (pay-as-you-go) account
// - `PAYM` - contract account
// - `Business` - Business (enterprise) account
//
// This attribute may be omitted from the response set if the information is not
// available
type TenureVerifyResponseContractType string

const (
	TenureVerifyResponseContractTypePayg     TenureVerifyResponseContractType = "PAYG"
	TenureVerifyResponseContractTypePaym     TenureVerifyResponseContractType = "PAYM"
	TenureVerifyResponseContractTypeBusiness TenureVerifyResponseContractType = "Business"
)

type TenureVerifyParams struct {
	// The date, in RFC 3339 / ISO 8601 compliant format "YYYY-MM-DD", from which
	// continuous tenure of the identified network subscriber is required to be
	// confirmed
	TenureDate time.Time `json:"tenureDate,required" format:"date"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r TenureVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow TenureVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TenureVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
