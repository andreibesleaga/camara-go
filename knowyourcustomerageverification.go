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

// Know Your Customer Age Verification
//
// KnowyourcustomerageverificationService contains methods and other services that
// help with interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKnowyourcustomerageverificationService] method instead.
type KnowyourcustomerageverificationService struct {
	Options []option.RequestOption
}

// NewKnowyourcustomerageverificationService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewKnowyourcustomerageverificationService(opts ...option.RequestOption) (r KnowyourcustomerageverificationService) {
	r = KnowyourcustomerageverificationService{}
	r.Options = opts
	return
}

// Verify that the age of the subscriber associated with a phone number is equal to
// or greater than the specified age threshold value.
//
// As it is possible that the person holding the contract and the end-user of the
// subscription may not be the same, the endpoint also admits a list of optional
// properties to be included in the request to improve the identification. The
// response may optionally include the `identityMatchScore` property with a value
// that indicates how certain it is that the information returned relates to the
// person that the API Client is requesting. To increase the reliability of the
// information returned, the API Provider may include in the response the
// `verifiedStatus` property, indicating whether the identity information in its
// possession has been verified against an identification document legally accepted
// as an age verification document (Note). Note: Depending on the country,
// credit-check or other mechanism can be used instead of official identification
// for Age Verification. For details, please contact API Provider.
//
// If the API Client indicates request properties `includeContentLock` or
// `includeParentalControl` with value `true` and the API Provider implements this
// functionality, then the response will also include `contentLock` and
// `parentalControl` properties to indicate if the subscription has any kind of
// content filtering enabled. On the other hand, if the request properties are not
// included or the API Client specifies value `false`, then the response properties
// will not be returned. If the API Provider doesn't implement this functionality,
// request properties will be ignored and response properties won't be returned in
// any case.
func (r *KnowyourcustomerageverificationService) Verify(ctx context.Context, params KnowyourcustomerageverificationVerifyParams, opts ...option.RequestOption) (res *KnowyourcustomerageverificationVerifyResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "knowyourcustomerageverification/verify"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Response to an age verification request
type KnowyourcustomerageverificationVerifyResponse struct {
	// Indicate `"true"` when the age of the user is the same age or older than the age
	// threshold (age >= age threshold), and `"false"` if not (age < age threshold). If
	// the API Provider doesn't have enough information to perform the validation, a
	// `not_available` can be returned.
	//
	// Any of "true", "false", "not_available".
	AgeCheck KnowyourcustomerageverificationVerifyResponseAgeCheck `json:"ageCheck" api:"required"`
	// Indicate `"true"` if the subscription associated with the phone number has any
	// kind of content lock (i.e certain web content blocked) and `"false"` if not. If
	// the information is not available the value `not_available` can be returned.
	//
	// Any of "true", "false", "not_available".
	ContentLock KnowyourcustomerageverificationVerifyResponseContentLock `json:"contentLock"`
	// The overall score of identity information available in the API Provider,
	// information either provided in the request body comparing it to the one that the
	// API Provider holds or directly using internal API Provider's information. It is
	// optional for the API Provider to return the Identity match score.
	IdentityMatchScore int64 `json:"identityMatchScore"`
	// Indicate `"true"` if the subscription associated with the phone number has any
	// kind of parental control activated and `"false"` if not. If the information is
	// not available the value `not_available` can be returned.
	//
	// Any of "true", "false", "not_available".
	ParentalControl KnowyourcustomerageverificationVerifyResponseParentalControl `json:"parentalControl"`
	// Indicate `true` if the information provided has been compared against
	// information based on an identification document legally accepted as an age
	// verification document (Note), otherwise indicate `false`. Note: Depending on the
	// country, credit-check or other mechanism can be used instead of official
	// identification for Age Verification. For details, please contact API Provider.
	VerifiedStatus bool `json:"verifiedStatus"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AgeCheck           respjson.Field
		ContentLock        respjson.Field
		IdentityMatchScore respjson.Field
		ParentalControl    respjson.Field
		VerifiedStatus     respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowyourcustomerageverificationVerifyResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowyourcustomerageverificationVerifyResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Indicate `"true"` when the age of the user is the same age or older than the age
// threshold (age >= age threshold), and `"false"` if not (age < age threshold). If
// the API Provider doesn't have enough information to perform the validation, a
// `not_available` can be returned.
type KnowyourcustomerageverificationVerifyResponseAgeCheck string

const (
	KnowyourcustomerageverificationVerifyResponseAgeCheckTrue         KnowyourcustomerageverificationVerifyResponseAgeCheck = "true"
	KnowyourcustomerageverificationVerifyResponseAgeCheckFalse        KnowyourcustomerageverificationVerifyResponseAgeCheck = "false"
	KnowyourcustomerageverificationVerifyResponseAgeCheckNotAvailable KnowyourcustomerageverificationVerifyResponseAgeCheck = "not_available"
)

// Indicate `"true"` if the subscription associated with the phone number has any
// kind of content lock (i.e certain web content blocked) and `"false"` if not. If
// the information is not available the value `not_available` can be returned.
type KnowyourcustomerageverificationVerifyResponseContentLock string

const (
	KnowyourcustomerageverificationVerifyResponseContentLockTrue         KnowyourcustomerageverificationVerifyResponseContentLock = "true"
	KnowyourcustomerageverificationVerifyResponseContentLockFalse        KnowyourcustomerageverificationVerifyResponseContentLock = "false"
	KnowyourcustomerageverificationVerifyResponseContentLockNotAvailable KnowyourcustomerageverificationVerifyResponseContentLock = "not_available"
)

// Indicate `"true"` if the subscription associated with the phone number has any
// kind of parental control activated and `"false"` if not. If the information is
// not available the value `not_available` can be returned.
type KnowyourcustomerageverificationVerifyResponseParentalControl string

const (
	KnowyourcustomerageverificationVerifyResponseParentalControlTrue         KnowyourcustomerageverificationVerifyResponseParentalControl = "true"
	KnowyourcustomerageverificationVerifyResponseParentalControlFalse        KnowyourcustomerageverificationVerifyResponseParentalControl = "false"
	KnowyourcustomerageverificationVerifyResponseParentalControlNotAvailable KnowyourcustomerageverificationVerifyResponseParentalControl = "not_available"
)

type KnowyourcustomerageverificationVerifyParams struct {
	// The age to be verified. The indicated range is a global definition of maximum
	// and minimum values allowed to be requested. It is important to note that this
	// range might be more restrictive in some implementations due to local regulations
	// of a country i.e. A country does not allow to request for an age under 18. This
	// limitation must be informed during the onboarding process.
	AgeThreshold int64 `json:"ageThreshold" api:"required"`
	// The birthdate of the customer, in RFC 3339 / ISO 8601 calendar date format
	// (YYYY-MM-DD).
	Birthdate param.Opt[time.Time] `json:"birthdate,omitzero" format:"date"`
	// Email address of the customer in the RFC specified format (local-part@domain).
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Last name, family name, or surname of the customer.
	FamilyName param.Opt[string] `json:"familyName,omitzero"`
	// Last/family/sur- name at birth of the customer.
	FamilyNameAtBirth param.Opt[string] `json:"familyNameAtBirth,omitzero"`
	// First/given name or compound first/given name of the customer.
	GivenName param.Opt[string] `json:"givenName,omitzero"`
	// Id number associated to the official identity document in the country. It may
	// contain alphanumeric characters.
	IDDocument param.Opt[string] `json:"idDocument,omitzero"`
	// If this parameter is included in the request with value `true`, the response
	// property `contentLock` will be returned. If it is not included or its value is
	// `false`, the response property will not be returned.
	IncludeContentLock param.Opt[bool] `json:"includeContentLock,omitzero"`
	// If this parameter is included in the request with value `true`, the response
	// property `parentalControl` will be returned. If it is not included or its value
	// is `false`, the response property will not be returned.
	IncludeParentalControl param.Opt[bool] `json:"includeParentalControl,omitzero"`
	// Middle name/s of the customer.
	MiddleNames param.Opt[string] `json:"middleNames,omitzero"`
	// Complete name of the customer, usually composed of first/given name and
	// last/family/sur- name in a country. Depending on the country, the order of
	// first/give name and last/family/sur- name varies, and middle name could be
	// included. It can use givenName, middleNames, familyName and/or
	// familyNameAtBirth. For example, in ESP, name+familyName; in NLD, it can be
	// name+middleNames+familyName or name+middleNames+familyNameAtBirth, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r KnowyourcustomerageverificationVerifyParams) MarshalJSON() (data []byte, err error) {
	type shadow KnowyourcustomerageverificationVerifyParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KnowyourcustomerageverificationVerifyParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
