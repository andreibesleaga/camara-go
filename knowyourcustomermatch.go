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

// Know Your Customer Match
//
// KnowyourcustomermatchService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewKnowyourcustomermatchService] method instead.
type KnowyourcustomermatchService struct {
	Options []option.RequestOption
}

// NewKnowyourcustomermatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewKnowyourcustomermatchService(opts ...option.RequestOption) (r KnowyourcustomermatchService) {
	r = KnowyourcustomermatchService{}
	r.Options = opts
	return
}

// Verify matching of a number of attributes related to a customer identity against
// the verified data bound to their phone number in the Operator systems.
// Regardless of whether the `phoneNumber` is explicitly stated in the request
// body, at least one of the other fields must be provided, otherwise a
// `HTTP 400 - KNOW_YOUR_CUSTOMER.INVALID_PARAM_COMBINATION` error will be
// returned.
//
// The API will return the result of the matching process for each requested
// attribute. This means that the response will **only** contain the attributes for
// which validation has been requested. Possible values are:
//
//   - **true**: the attribute provided matches with the one in the Operator systems,
//     which is equal to a `match_score` of 100.
//   - **false**: the attribute provided does not match with the one in the Operator
//     systems.
//   - **not_available**: the attribute is not available to validate.
func (r *KnowyourcustomermatchService) Match(ctx context.Context, params KnowyourcustomermatchMatchParams, opts ...option.RequestOption) (res *KnowyourcustomermatchMatchResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "knowyourcustomermatch/match"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// true - the attribute provided matches with the one in the Operator systems,
// which is equal to a `match_score` of 100. false - the attribute provided does
// not match with the one in the Operator systems. not_available - the attribute is
// not available to validate.
type MatchResult string

const (
	MatchResultTrue         MatchResult = "true"
	MatchResultFalse        MatchResult = "false"
	MatchResultNotAvailable MatchResult = "not_available"
)

type KnowyourcustomermatchMatchResponse struct {
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	AddressMatch MatchResult `json:"addressMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	AddressMatchScore int64 `json:"addressMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	BirthdateMatch MatchResult `json:"birthdateMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	CityOfBirthMatch MatchResult `json:"cityOfBirthMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	CityOfBirthMatchScore int64 `json:"cityOfBirthMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	CountryMatch MatchResult `json:"countryMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	CountryOfBirthMatch MatchResult `json:"countryOfBirthMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	EmailMatch MatchResult `json:"emailMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	EmailMatchScore int64 `json:"emailMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	FamilyNameAtBirthMatch MatchResult `json:"familyNameAtBirthMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	FamilyNameAtBirthMatchScore int64 `json:"familyNameAtBirthMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	FamilyNameMatch MatchResult `json:"familyNameMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	FamilyNameMatchScore int64 `json:"familyNameMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	GenderMatch MatchResult `json:"genderMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	GivenNameMatch MatchResult `json:"givenNameMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	GivenNameMatchScore int64 `json:"givenNameMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	HouseNumberExtensionMatch MatchResult `json:"houseNumberExtensionMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	IDDocumentExpiryDateMatch MatchResult `json:"idDocumentExpiryDateMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	IDDocumentMatch MatchResult `json:"idDocumentMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	IDDocumentTypeMatch MatchResult `json:"idDocumentTypeMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	LocalityMatch MatchResult `json:"localityMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	LocalityMatchScore int64 `json:"localityMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	MiddleNamesMatch MatchResult `json:"middleNamesMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	MiddleNamesMatchScore int64 `json:"middleNamesMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	NameKanaHankakuMatch MatchResult `json:"nameKanaHankakuMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	NameKanaHankakuMatchScore int64 `json:"nameKanaHankakuMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	NameKanaZenkakuMatch MatchResult `json:"nameKanaZenkakuMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	NameKanaZenkakuMatchScore int64 `json:"nameKanaZenkakuMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	NameMatch MatchResult `json:"nameMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	NameMatchScore int64 `json:"nameMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	NationalityMatch MatchResult `json:"nationalityMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	PostalCodeMatch MatchResult `json:"postalCodeMatch"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	RegionMatch MatchResult `json:"regionMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	RegionMatchScore int64 `json:"regionMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	StreetNameMatch MatchResult `json:"streetNameMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	StreetNameMatchScore int64 `json:"streetNameMatchScore"`
	// true - the attribute provided matches with the one in the Operator systems,
	// which is equal to a `match_score` of 100. false - the attribute provided does
	// not match with the one in the Operator systems. not_available - the attribute is
	// not available to validate.
	//
	// Any of "true", "false", "not_available".
	StreetNumberMatch MatchResult `json:"streetNumberMatch"`
	// Indicates the similarity score assigned to the input value when it does not
	// exactly match the value stored in the operator's system. This property shall
	// only be returned when the value of the corresponding match field is `false`. A
	// perfect match with a score of 100 is indicated by `match` being 'true' and no
	// `matchScore` is returned in this case.
	StreetNumberMatchScore int64 `json:"streetNumberMatchScore"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AddressMatch                respjson.Field
		AddressMatchScore           respjson.Field
		BirthdateMatch              respjson.Field
		CityOfBirthMatch            respjson.Field
		CityOfBirthMatchScore       respjson.Field
		CountryMatch                respjson.Field
		CountryOfBirthMatch         respjson.Field
		EmailMatch                  respjson.Field
		EmailMatchScore             respjson.Field
		FamilyNameAtBirthMatch      respjson.Field
		FamilyNameAtBirthMatchScore respjson.Field
		FamilyNameMatch             respjson.Field
		FamilyNameMatchScore        respjson.Field
		GenderMatch                 respjson.Field
		GivenNameMatch              respjson.Field
		GivenNameMatchScore         respjson.Field
		HouseNumberExtensionMatch   respjson.Field
		IDDocumentExpiryDateMatch   respjson.Field
		IDDocumentMatch             respjson.Field
		IDDocumentTypeMatch         respjson.Field
		LocalityMatch               respjson.Field
		LocalityMatchScore          respjson.Field
		MiddleNamesMatch            respjson.Field
		MiddleNamesMatchScore       respjson.Field
		NameKanaHankakuMatch        respjson.Field
		NameKanaHankakuMatchScore   respjson.Field
		NameKanaZenkakuMatch        respjson.Field
		NameKanaZenkakuMatchScore   respjson.Field
		NameMatch                   respjson.Field
		NameMatchScore              respjson.Field
		NationalityMatch            respjson.Field
		PostalCodeMatch             respjson.Field
		RegionMatch                 respjson.Field
		RegionMatchScore            respjson.Field
		StreetNameMatch             respjson.Field
		StreetNameMatchScore        respjson.Field
		StreetNumberMatch           respjson.Field
		StreetNumberMatchScore      respjson.Field
		ExtraFields                 map[string]respjson.Field
		raw                         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KnowyourcustomermatchMatchResponse) RawJSON() string { return r.JSON.raw }
func (r *KnowyourcustomermatchMatchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type KnowyourcustomermatchMatchParams struct {
	// Complete address of the customer. For some countries, it is built following the
	// usual concatenation of parameters in a country, but for other countries, this is
	// not the case. For some countries, it can use streetName, streetNumber and/or
	// houseNumberExtension. For example, in ESP, streetName+streetNumber; in NLD, it
	// can be streetName+streetNumber or streetName+streetNumber+houseNumberExtension.
	Address param.Opt[string] `json:"address,omitzero"`
	// The birthdate of the customer, in RFC 3339 / ISO 8601 calendar date format
	// (YYYY-MM-DD).
	Birthdate param.Opt[time.Time] `json:"birthdate,omitzero" format:"date"`
	// City where the customer was born.
	CityOfBirth param.Opt[string] `json:"cityOfBirth,omitzero"`
	// Country of the customer's address. Format ISO 3166-1 alpha-2
	Country param.Opt[string] `json:"country,omitzero"`
	// Country where the customer was born. Format ISO 3166-1 alpha-2.
	CountryOfBirth param.Opt[string] `json:"countryOfBirth,omitzero"`
	// Email address of the customer in the RFC specified format (local-part@domain).
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	// Last name, family name, or surname of the customer.
	FamilyName param.Opt[string] `json:"familyName,omitzero"`
	// Last/family/sur- name at birth of the customer.
	FamilyNameAtBirth param.Opt[string] `json:"familyNameAtBirth,omitzero"`
	// First/given name or compound first/given name of the customer.
	GivenName param.Opt[string] `json:"givenName,omitzero"`
	// Specific identifier of the house needed depending on the property type. For
	// example, number of apartment in an apartment building.
	HouseNumberExtension param.Opt[string] `json:"houseNumberExtension,omitzero"`
	// Id number associated to the official identity document in the country. It may
	// contain alphanumeric characters.
	IDDocument param.Opt[string] `json:"idDocument,omitzero"`
	// Expiration date of the identity document (ISO 8601).
	IDDocumentExpiryDate param.Opt[time.Time] `json:"idDocumentExpiryDate,omitzero" format:"date"`
	// Locality of the customer's address
	Locality param.Opt[string] `json:"locality,omitzero"`
	// Middle name/s of the customer.
	MiddleNames param.Opt[string] `json:"middleNames,omitzero"`
	// Complete name of the customer, usually composed of first/given name and
	// last/family/sur- name in a country. Depending on the country, the order of
	// first/give name and last/family/sur- name varies, and middle name could be
	// included. It can use givenName, middleNames, familyName and/or
	// familyNameAtBirth. For example, in ESP, name+familyName; in NLD, it can be
	// name+middleNames+familyName or name+middleNames+familyNameAtBirth, etc.
	Name param.Opt[string] `json:"name,omitzero"`
	// Complete name of the customer in Hankaku-Kana format (reading of name) for
	// Japan.
	NameKanaHankaku param.Opt[string] `json:"nameKanaHankaku,omitzero"`
	// Complete name of the customer in Zenkaku-Kana format (reading of name) for
	// Japan.
	NameKanaZenkaku param.Opt[string] `json:"nameKanaZenkaku,omitzero"`
	// ISO 3166-1 alpha-2 code of the customer’s nationality. In the case a customer
	// has more than one nationality, it is supposed to be the nationality related to
	// the ID document provided in the match request.
	Nationality param.Opt[string] `json:"nationality,omitzero"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// Zip code or postal code
	PostalCode param.Opt[string] `json:"postalCode,omitzero"`
	// Region/prefecture of the customer's address
	Region param.Opt[string] `json:"region,omitzero"`
	// Name of the street of the customer's address. It should not include the type of
	// the street.
	StreetName param.Opt[string] `json:"streetName,omitzero"`
	// The street number of the customer's address. Number identifying a specific
	// property on the 'streetName'.
	StreetNumber param.Opt[string] `json:"streetNumber,omitzero"`
	XCorrelator  param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	// Gender of the customer (Male/Female/Other).
	//
	// Any of "MALE", "FEMALE", "OTHER".
	Gender KnowyourcustomermatchMatchParamsGender `json:"gender,omitzero"`
	// Type of the official identity document provided.
	//
	// Any of "passport", "national_id_card", "residence_permit", "diplomatic_id",
	// "driver_licence", "social_security_id", "other".
	IDDocumentType KnowyourcustomermatchMatchParamsIDDocumentType `json:"idDocumentType,omitzero"`
	paramObj
}

func (r KnowyourcustomermatchMatchParams) MarshalJSON() (data []byte, err error) {
	type shadow KnowyourcustomermatchMatchParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *KnowyourcustomermatchMatchParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Gender of the customer (Male/Female/Other).
type KnowyourcustomermatchMatchParamsGender string

const (
	KnowyourcustomermatchMatchParamsGenderMale   KnowyourcustomermatchMatchParamsGender = "MALE"
	KnowyourcustomermatchMatchParamsGenderFemale KnowyourcustomermatchMatchParamsGender = "FEMALE"
	KnowyourcustomermatchMatchParamsGenderOther  KnowyourcustomermatchMatchParamsGender = "OTHER"
)

// Type of the official identity document provided.
type KnowyourcustomermatchMatchParamsIDDocumentType string

const (
	KnowyourcustomermatchMatchParamsIDDocumentTypePassport         KnowyourcustomermatchMatchParamsIDDocumentType = "passport"
	KnowyourcustomermatchMatchParamsIDDocumentTypeNationalIDCard   KnowyourcustomermatchMatchParamsIDDocumentType = "national_id_card"
	KnowyourcustomermatchMatchParamsIDDocumentTypeResidencePermit  KnowyourcustomermatchMatchParamsIDDocumentType = "residence_permit"
	KnowyourcustomermatchMatchParamsIDDocumentTypeDiplomaticID     KnowyourcustomermatchMatchParamsIDDocumentType = "diplomatic_id"
	KnowyourcustomermatchMatchParamsIDDocumentTypeDriverLicence    KnowyourcustomermatchMatchParamsIDDocumentType = "driver_licence"
	KnowyourcustomermatchMatchParamsIDDocumentTypeSocialSecurityID KnowyourcustomermatchMatchParamsIDDocumentType = "social_security_id"
	KnowyourcustomermatchMatchParamsIDDocumentTypeOther            KnowyourcustomermatchMatchParamsIDDocumentType = "other"
)
