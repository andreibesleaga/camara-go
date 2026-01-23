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

// PopulationdensitydataService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPopulationdensitydataService] method instead.
type PopulationdensitydataService struct {
	Options []option.RequestOption
}

// NewPopulationdensitydataService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewPopulationdensitydataService(opts ...option.RequestOption) (r PopulationdensitydataService) {
	r = PopulationdensitydataService{}
	r.Options = opts
	return
}

// Retrieves population density estimation together with the estimation range
// related for a time slot for a given area (described as a polygon) as a data set
// consisting of a sequence of equally-sized objects covering the input polygon
// area.
func (r *PopulationdensitydataService) Get(ctx context.Context, params PopulationdensitydataGetParams, opts ...option.RequestOption) (res *PopulationdensitydataGetResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "populationdensitydata/retrieve"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Population density values is represented in time intervals for different cells
// of the requested area. Each element in `timedPopulationDensityData` array
// corresponds to a time interval, containing population density data for the grid
// cells. The intervals are 1 hour long.
type PopulationdensitydataGetResponse struct {
	// Represents the state of the response for the input polygon defined in the
	// request, the possible values are:
	//
	//   - `SUPPORTED_AREA`: The whole request area is supported. Population density data
	//     for the entire requested area is returned.
	//   - `PART_OF_AREA_NOT_SUPPORTED`: Part of the requested area is outside the MNOs
	//     coverage area, the cells outside the coverage area will have property
	//     `dataType` with value `NO_DATA`.
	//   - `AREA_NOT_SUPPORTED`: The whole requested area is outside the MNOs coverage
	//     area. No data will be returned.
	//   - `OPERATION_NOT_COMPLETED`: An error happened during asynchronous processing of
	//     the request. This status will only be returned in case the asynchronous API
	//     behaviour is used.
	//
	// Any of "SUPPORTED_AREA", "PART_OF_AREA_NOT_SUPPORTED", "AREA_NOT_SUPPORTED",
	// "OPERATION_NOT_COMPLETED".
	Status PopulationdensitydataGetResponseStatus `json:"status,required"`
	// Time ranges along with the population density data for the cells within it. The
	// request startTime or the request endTime have to be fully covered by the
	// intervals. For example, if the intervals are 1-hour long and the input date
	// range were [2024-01-03T11:25:00Z to 2024-01-03T12:45:00Z] it would contain 2
	// intervals (Interval from 2024-01-03T11:00:00Z to 2024-01-03T12:00:00Z and
	// interval from 2024-01-03T12:00:00Z to 2024-01-03T13:00:00Z).
	TimedPopulationDensityData []PopulationdensitydataGetResponseTimedPopulationDensityData `json:"timedPopulationDensityData,required"`
	// Information about the status, mandatory when property `status` is
	// `OPERATION_NOT_COMPLETED` for adding extra information about the error.
	StatusInfo string `json:"statusInfo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status                     respjson.Field
		TimedPopulationDensityData respjson.Field
		StatusInfo                 respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PopulationdensitydataGetResponse) RawJSON() string { return r.JSON.raw }
func (r *PopulationdensitydataGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Represents the state of the response for the input polygon defined in the
// request, the possible values are:
//
//   - `SUPPORTED_AREA`: The whole request area is supported. Population density data
//     for the entire requested area is returned.
//   - `PART_OF_AREA_NOT_SUPPORTED`: Part of the requested area is outside the MNOs
//     coverage area, the cells outside the coverage area will have property
//     `dataType` with value `NO_DATA`.
//   - `AREA_NOT_SUPPORTED`: The whole requested area is outside the MNOs coverage
//     area. No data will be returned.
//   - `OPERATION_NOT_COMPLETED`: An error happened during asynchronous processing of
//     the request. This status will only be returned in case the asynchronous API
//     behaviour is used.
type PopulationdensitydataGetResponseStatus string

const (
	PopulationdensitydataGetResponseStatusSupportedArea          PopulationdensitydataGetResponseStatus = "SUPPORTED_AREA"
	PopulationdensitydataGetResponseStatusPartOfAreaNotSupported PopulationdensitydataGetResponseStatus = "PART_OF_AREA_NOT_SUPPORTED"
	PopulationdensitydataGetResponseStatusAreaNotSupported       PopulationdensitydataGetResponseStatus = "AREA_NOT_SUPPORTED"
	PopulationdensitydataGetResponseStatusOperationNotCompleted  PopulationdensitydataGetResponseStatus = "OPERATION_NOT_COMPLETED"
)

type PopulationdensitydataGetResponseTimedPopulationDensityData struct {
	// Population density data for the different cells in a concrete time range.
	CellPopulationDensityData []PopulationdensitydataGetResponseTimedPopulationDensityDataCellPopulationDensityData `json:"cellPopulationDensityData,required"`
	// Interval end time. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ (i.e. which
	// allows 2023-07-03T14:27:08.312+02:00 or 2023-07-03T12:27:08.312Z)
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Interval start time. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ (i.e. which
	// allows 2023-07-03T14:27:08.312+02:00 or 2023-07-03T12:27:08.312Z)
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CellPopulationDensityData respjson.Field
		EndTime                   respjson.Field
		StartTime                 respjson.Field
		ExtraFields               map[string]respjson.Field
		raw                       string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PopulationdensitydataGetResponseTimedPopulationDensityData) RawJSON() string {
	return r.JSON.raw
}
func (r *PopulationdensitydataGetResponseTimedPopulationDensityData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Population density data of a cell in a concrete time range. In case of
// insufficient data, to guarantee an anonymized prediction due to the k-anonymity
// within a specific cell and time range, no population density data is returned
// and the property `dataType` value is "LOW_DENSITY". In case of a cell not
// supported `dataType` value is "NO_DATA"
type PopulationdensitydataGetResponseTimedPopulationDensityDataCellPopulationDensityData struct {
	// Any of "NO_DATA", "LOW_DENSITY", "DENSITY_ESTIMATION".
	DataType string `json:"dataType,required"`
	// Coordinates of the cell represented as a string using the
	// [Geohash system](https://en.wikipedia.org/wiki/Geohash). Encoding a geographic
	// location into a short string. The value length, and thus, the cell granularity,
	// is determined by the request body property `precision`.
	Geohash string `json:"geohash,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DataType    respjson.Field
		Geohash     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PopulationdensitydataGetResponseTimedPopulationDensityDataCellPopulationDensityData) RawJSON() string {
	return r.JSON.raw
}
func (r *PopulationdensitydataGetResponseTimedPopulationDensityDataCellPopulationDensityData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PopulationdensitydataGetParams struct {
	// Base schema for all areas
	Area PopulationdensitydataGetParamsArea `json:"area,omitzero,required"`
	// End date time. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ (i.e. which
	// allows 2023-07-03T14:27:08.312+02:00 or 2023-07-03T12:27:08.312Z) The maximum
	// endTime allowed is 3 months from the time of the request.
	EndTime time.Time `json:"endTime,required" format:"date-time"`
	// Start date time. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ
	StartTime time.Time `json:"startTime,required" format:"date-time"`
	// Precision required of response cells. Precision defines a geohash level and
	// corresponds to the length of the geohash for each cell. More information at
	// [Geohash system](https://en.wikipedia.org/wiki/Geohash)" If not included the
	// default precision level 7 is used by default. In case of using a not supported
	// level by the MNO, the API returns the error response
	// `POPULATION_DENSITY_DATA.UNSUPPORTED_PRECISION`.
	Precision param.Opt[int64] `json:"precision,omitzero"`
	// The address where the API response will be asynchronously delivered, using the
	// HTTP protocol.
	Sink        param.Opt[string] `json:"sink,omitzero" format:"uri"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	// A sink credential provides authentication or authorization information necessary
	// to enable delivery of events to a target.
	SinkCredential PopulationdensitydataGetParamsSinkCredential `json:"sinkCredential,omitzero"`
	paramObj
}

func (r PopulationdensitydataGetParams) MarshalJSON() (data []byte, err error) {
	type shadow PopulationdensitydataGetParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PopulationdensitydataGetParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Base schema for all areas
//
// The property AreaType is required.
type PopulationdensitydataGetParamsArea struct {
	// Type of this area. POLYGON - The area is defined as a polygon.
	//
	// Any of "POLYGON".
	AreaType string `json:"areaType,omitzero,required"`
	paramObj
}

func (r PopulationdensitydataGetParamsArea) MarshalJSON() (data []byte, err error) {
	type shadow PopulationdensitydataGetParamsArea
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PopulationdensitydataGetParamsArea) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PopulationdensitydataGetParamsArea](
		"areaType", "POLYGON",
	)
}

// A sink credential provides authentication or authorization information necessary
// to enable delivery of events to a target.
//
// The property CredentialType is required.
type PopulationdensitydataGetParamsSinkCredential struct {
	// The type of the credential. Note: Type of the credential - MUST be set to
	// ACCESSTOKEN for now
	//
	// Any of "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN".
	CredentialType string `json:"credentialType,omitzero,required"`
	paramObj
}

func (r PopulationdensitydataGetParamsSinkCredential) MarshalJSON() (data []byte, err error) {
	type shadow PopulationdensitydataGetParamsSinkCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PopulationdensitydataGetParamsSinkCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PopulationdensitydataGetParamsSinkCredential](
		"credentialType", "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN",
	)
}
