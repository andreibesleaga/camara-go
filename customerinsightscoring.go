// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"fmt"
	"net/http"
	"slices"

	"github.com/andreibesleaga/camara-go/internal/apijson"
	"github.com/andreibesleaga/camara-go/internal/requestconfig"
	"github.com/andreibesleaga/camara-go/option"
	"github.com/andreibesleaga/camara-go/packages/param"
	"github.com/andreibesleaga/camara-go/packages/respjson"
)

// Customer Insights
//
// CustomerinsightScoringService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomerinsightScoringService] method instead.
type CustomerinsightScoringService struct {
	Options []option.RequestOption
}

// NewCustomerinsightScoringService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewCustomerinsightScoringService(opts ...option.RequestOption) (r CustomerinsightScoringService) {
	r = CustomerinsightScoringService{}
	r.Options = opts
	return
}

// Retrieves Scoring information, for the user associated with the provided
// `idDocument`, `phoneNumber` or the combination of both parameters. It also
// allows to select the type of the Scoring scale measurement.
func (r *CustomerinsightScoringService) Get(ctx context.Context, params CustomerinsightScoringGetParams, opts ...option.RequestOption) (res *CustomerinsightScoringGetResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "customerinsights/scoring/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Scoring information based on the individual's profile owned by a Telco Operator.
type CustomerinsightScoringGetResponse struct {
	// Scoring measurement system.
	//
	// Allowed values are:
	//
	// - `gaugeMetric`: ranges from index 850 (lowest risk) to index 300 (highest risk)
	// - `veritasIndex`: ranges from index 0 (lowest risk) to index 19 (highest risk)
	//
	// Any of "gaugeMetric", "veritasIndex".
	ScoringType CustomerinsightScoringGetResponseScoringType `json:"scoringType" api:"required"`
	// Result of the Scoring analysis expressed in the measure indicated in the
	// `scoringType` field.
	ScoringValue int64 `json:"scoringValue" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ScoringType  respjson.Field
		ScoringValue respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomerinsightScoringGetResponse) RawJSON() string { return r.JSON.raw }
func (r *CustomerinsightScoringGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scoring measurement system.
//
// Allowed values are:
//
// - `gaugeMetric`: ranges from index 850 (lowest risk) to index 300 (highest risk)
// - `veritasIndex`: ranges from index 0 (lowest risk) to index 19 (highest risk)
type CustomerinsightScoringGetResponseScoringType string

const (
	CustomerinsightScoringGetResponseScoringTypeGaugeMetric  CustomerinsightScoringGetResponseScoringType = "gaugeMetric"
	CustomerinsightScoringGetResponseScoringTypeVeritasIndex CustomerinsightScoringGetResponseScoringType = "veritasIndex"
)

type CustomerinsightScoringGetParams struct {
	// Identification number associated to the official identity document in the
	// country. It may contain alphanumeric characters.
	IDDocument param.Opt[string] `json:"idDocument,omitzero"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	// Scoring type, i.e.: scale. API Client may use this field to indicate the Scoring
	// in one of the defined scales; if this field is not informed, the API will return
	// the Scoring in the scale configured by default in the system.
	//
	// Allowed values are:
	//
	// - `gaugeMetric`: ranges from index 850 (lowest risk) to index 300 (highest risk)
	// - `veritasIndex`: ranges from index 0 (lowest risk) to index 19 (highest risk)
	//
	// Any of "gaugeMetric", "veritasIndex".
	ScoringType CustomerinsightScoringGetParamsScoringType `json:"scoringType,omitzero"`
	paramObj
}

func (r CustomerinsightScoringGetParams) MarshalJSON() (data []byte, err error) {
	type shadow CustomerinsightScoringGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomerinsightScoringGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Scoring type, i.e.: scale. API Client may use this field to indicate the Scoring
// in one of the defined scales; if this field is not informed, the API will return
// the Scoring in the scale configured by default in the system.
//
// Allowed values are:
//
// - `gaugeMetric`: ranges from index 850 (lowest risk) to index 300 (highest risk)
// - `veritasIndex`: ranges from index 0 (lowest risk) to index 19 (highest risk)
type CustomerinsightScoringGetParamsScoringType string

const (
	CustomerinsightScoringGetParamsScoringTypeGaugeMetric  CustomerinsightScoringGetParamsScoringType = "gaugeMetric"
	CustomerinsightScoringGetParamsScoringTypeVeritasIndex CustomerinsightScoringGetParamsScoringType = "veritasIndex"
)
