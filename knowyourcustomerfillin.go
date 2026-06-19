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

// Know Your Customer Fill-in
//
// KnowyourcustomerfillInService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKnowyourcustomerfillInService] method instead.
type KnowyourcustomerfillInService struct {
	Options []option.RequestOption
}

// NewKnowyourcustomerfillInService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewKnowyourcustomerfillInService(opts ...option.RequestOption) (r KnowyourcustomerfillInService) {
	r = KnowyourcustomerfillInService{}
	r.Options = opts
	return
}

// Providing information related to a customer identity stored the account data
// bound to the customer's phone number.
func (r *KnowyourcustomerfillInService) New(ctx context.Context, params KnowyourcustomerfillInNewParams, opts ...option.RequestOption) (res *KnowyourcustomerfillInNewResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "knowyourcustomerfill-in/fill-in"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type KnowyourcustomerfillInNewResponse struct {
	// Complete address of the customer stored on the Operator's system. For some
	// countries, it is built following the usual concatenation of parameters in a
	// country, but for other countries, this is not the case. For some countries, it
	// can use streetName, streetNumber and/or houseNumberExtension. For example, in
	// ESP, streetName+streetNumber; in NLD, it can be streetName+streetNumber or
	// streetName+streetNumber+houseNumberExtension.
	Address string `json:"address"`
	// Birthdate of the customer, in ISO 8601 calendar date format (YYYY-MM-DD), stored
	// on the Operator's system.
	Birthdate time.Time `json:"birthdate" format:"date"`
	// City where the customer was born.
	CityOfBirth string `json:"cityOfBirth"`
	// Country of the customer's address stored on the Operator's system. Format ISO
	// 3166-1 alpha-2.
	Country string `json:"country"`
	// Country where the customer was born. Format ISO 3166-1 alpha-2.
	CountryOfBirth string `json:"countryOfBirth"`
	// Email address of the customer in the RFC specified format (local-part@domain),
	// stored on the Operator's system.
	Email string `json:"email" format:"email"`
	// Last name, family name, or surname of the customer stored on the Operator's
	// system.
	FamilyName string `json:"familyName"`
	// Last/family/sur- name at birth of the customer stored on the Operator's system.
	FamilyNameAtBirth string `json:"familyNameAtBirth"`
	// Gender of the customer stored on the Operator's system (Male/Female/Other).
	//
	// Any of "MALE", "FEMALE", "OTHER".
	Gender KnowyourcustomerfillInNewResponseGender `json:"gender"`
	// First/given name or compound first/given name of the customer on the Operator's
	// system.
	GivenName string `json:"givenName"`
	// House number extension of the customer stored on the Operator's system. Specific
	// identifier of the house needed depending on the property type. For example,
	// number of apartment in an apartment building.
	HouseNumberExtension string `json:"houseNumberExtension"`
	// Id number associated to the id_document of the customer stored on the Operator's
	// system.
	IDDocument string `json:"idDocument"`
	// Expiration date of the identity document (ISO 8601).
	IDDocumentExpiryDate time.Time `json:"idDocumentExpiryDate" format:"date"`
	// Type of the official identity document provided.
	//
	// Any of "passport", "national_id_card", "residence_permit", "diplomatic_id",
	// "driver_licence", "social_security_id", "other".
	IDDocumentType KnowyourcustomerfillInNewResponseIDDocumentType `json:"idDocumentType"`
	// Locality of the customer's address, stored on the Operator's system.
	Locality string `json:"locality"`
	// Middle name/s of the customer stored on the Operator's system.
	MiddleNames string `json:"middleNames"`
	// Complete name of the customer stored on the Operator's system. It is usually
	// composed of first/given name and last/family/sur- name in a country. Depending
	// on the country, the order of first/give name and last/family/sur- name varies,
	// and middle name could be included. It can use givenName, middleNames, familyName
	// and/or familyNameAtBirth. For example, in ESP, name+familyName; in NLD, it can
	// be name+middleNames+familyName or name+middleNames+familyNameAtBirth, etc.
	Name string `json:"name"`
	// Complete name of the customer in Hankaku-Kana format (reading of name) for
	// Japan, stored on the Operator's system.
	NameKanaHankaku string `json:"nameKanaHankaku"`
	// Complete name of the customer in Zenkaku-Kana format (reading of name) for
	// Japan, stored on the Operator's system.
	NameKanaZenkaku string `json:"nameKanaZenkaku"`
	// ISO 3166-1 alpha-2 code of the customer’s nationality. In the case a customer
	// has more than one nationality, it is supposed to be the nationality related to
	// the ID document provided in the match request.
	Nationality string `json:"nationality"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber string `json:"phoneNumber"`
	// The postal code or Zip code of the customer's address, stored on the Operator's
	// system.
	PostalCode string `json:"postalCode"`
	// Region/prefecture of the customer's address, stored on the Operator's system.
	Region string `json:"region"`
	// Name of the street of the customer's address on the Operator's system. It should
	// not include the type of the street.
	StreetName string `json:"streetName"`
	// The street number of the customer's address on the Operator's system. Number
	// identifying a specific property on the 'streetName'.
	StreetNumber string `json:"streetNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Address              respjson.Field
		Birthdate            respjson.Field
		CityOfBirth          respjson.Field
		Country              respjson.Field
		CountryOfBirth       respjson.Field
		Email                respjson.Field
		FamilyName           respjson.Field
		FamilyNameAtBirth    respjson.Field
		Gender               respjson.Field
		GivenName            respjson.Field
		HouseNumberExtension respjson.Field
		IDDocument           respjson.Field
		IDDocumentExpiryDate respjson.Field
		IDDocumentType       respjson.Field
		Locality             respjson.Field
		MiddleNames          respjson.Field
		Name                 respjson.Field
		NameKanaHankaku      respjson.Field
		NameKanaZenkaku      respjson.Field
		Nationality          respjson.Field
		PhoneNumber          respjson.Field
		PostalCode           respjson.Field
		Region               respjson.Field
		StreetName           respjson.Field
		StreetNumber         respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowyourcustomerfillInNewResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowyourcustomerfillInNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gender of the customer stored on the Operator's system (Male/Female/Other).
type KnowyourcustomerfillInNewResponseGender string

const (
	KnowyourcustomerfillInNewResponseGenderMale   KnowyourcustomerfillInNewResponseGender = "MALE"
	KnowyourcustomerfillInNewResponseGenderFemale KnowyourcustomerfillInNewResponseGender = "FEMALE"
	KnowyourcustomerfillInNewResponseGenderOther  KnowyourcustomerfillInNewResponseGender = "OTHER"
)

// Type of the official identity document provided.
type KnowyourcustomerfillInNewResponseIDDocumentType string

const (
	KnowyourcustomerfillInNewResponseIDDocumentTypePassport         KnowyourcustomerfillInNewResponseIDDocumentType = "passport"
	KnowyourcustomerfillInNewResponseIDDocumentTypeNationalIDCard   KnowyourcustomerfillInNewResponseIDDocumentType = "national_id_card"
	KnowyourcustomerfillInNewResponseIDDocumentTypeResidencePermit  KnowyourcustomerfillInNewResponseIDDocumentType = "residence_permit"
	KnowyourcustomerfillInNewResponseIDDocumentTypeDiplomaticID     KnowyourcustomerfillInNewResponseIDDocumentType = "diplomatic_id"
	KnowyourcustomerfillInNewResponseIDDocumentTypeDriverLicence    KnowyourcustomerfillInNewResponseIDDocumentType = "driver_licence"
	KnowyourcustomerfillInNewResponseIDDocumentTypeSocialSecurityID KnowyourcustomerfillInNewResponseIDDocumentType = "social_security_id"
	KnowyourcustomerfillInNewResponseIDDocumentTypeOther            KnowyourcustomerfillInNewResponseIDDocumentType = "other"
)

type KnowyourcustomerfillInNewParams struct {
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r KnowyourcustomerfillInNewParams) MarshalJSON() (data []byte, err error) {
	type shadow KnowyourcustomerfillInNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KnowyourcustomerfillInNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
