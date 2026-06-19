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

// Region Device Count
//
// RegiondevicecountService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewRegiondevicecountService] method instead.
type RegiondevicecountService struct {
	Options []option.RequestOption
}

// NewRegiondevicecountService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewRegiondevicecountService(opts ...option.RequestOption) (r RegiondevicecountService) {
	r = RegiondevicecountService{}
	r.Options = opts
	return
}

// Get the number of devices in the specified area during a certain time interval.
//
//   - The query area can be a circle or a polygon composed of longitude and latitude
//     points.
//   - If the areaType is circle, the circleCenter and circleRadius must be provided;
//     if the area is a polygon, the point list must be provided.
//   - If starttime and endtime are not passed in,this api should return the current
//     number of devices in the area.
//   - If the device appears in the specified area at least once during the certain
//     time interval, it should be counted.
func (r *RegiondevicecountService) GetCount(ctx context.Context, params RegiondevicecountGetCountParams, opts ...option.RequestOption) (res *RegiondevicecountGetCountResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "regiondevicecount/count"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// RegionDeviceCount result
type RegiondevicecountGetCountResponse struct {
	// Device Count
	Count float64 `json:"count"`
	// SUPPORTED_AREA: The whole requested area is supported Region Device Count for
	// the entire requested area is returned - Telco Coverage = 100 %
	//
	// PART_OF_AREA_NOT_SUPPORTED: Part of the requested area is outside the MNOs
	// coverage area, the area outside the coverage area are not returned - 100% >Telco
	// Coverage >=50%
	//
	// AREA_NOT_SUPPORTED: The whole requested area is outside the MNO coverage area No
	// data will be returned- Telco Coverage <50%
	//
	// DENSITY_BELOW_PRIVACY_THRESHOLD: The number of connected devices is below
	// privacy threshold of local regulation
	//
	// TIME_INTERVAL_NO_DATA_FOUND: Unable to find device count data within the
	// requested time interval
	//
	// Any of "SUPPORTED_AREA", "PART_OF_AREA_NOT_SUPPORTED", "AREA_NOT_SUPPORTED",
	// "DENSITY_BELOW_PRIVACY_THRESHOLD", "TIME_INTERVAL_NO_DATA_FOUND".
	Status RegiondevicecountGetCountResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RegiondevicecountGetCountResponse) RawJSON() string { return r.JSON.raw }
func (r *RegiondevicecountGetCountResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SUPPORTED_AREA: The whole requested area is supported Region Device Count for
// the entire requested area is returned - Telco Coverage = 100 %
//
// PART_OF_AREA_NOT_SUPPORTED: Part of the requested area is outside the MNOs
// coverage area, the area outside the coverage area are not returned - 100% >Telco
// Coverage >=50%
//
// AREA_NOT_SUPPORTED: The whole requested area is outside the MNO coverage area No
// data will be returned- Telco Coverage <50%
//
// DENSITY_BELOW_PRIVACY_THRESHOLD: The number of connected devices is below
// privacy threshold of local regulation
//
// TIME_INTERVAL_NO_DATA_FOUND: Unable to find device count data within the
// requested time interval
type RegiondevicecountGetCountResponseStatus string

const (
	RegiondevicecountGetCountResponseStatusSupportedArea                RegiondevicecountGetCountResponseStatus = "SUPPORTED_AREA"
	RegiondevicecountGetCountResponseStatusPartOfAreaNotSupported       RegiondevicecountGetCountResponseStatus = "PART_OF_AREA_NOT_SUPPORTED"
	RegiondevicecountGetCountResponseStatusAreaNotSupported             RegiondevicecountGetCountResponseStatus = "AREA_NOT_SUPPORTED"
	RegiondevicecountGetCountResponseStatusDensityBelowPrivacyThreshold RegiondevicecountGetCountResponseStatus = "DENSITY_BELOW_PRIVACY_THRESHOLD"
	RegiondevicecountGetCountResponseStatusTimeIntervalNoDataFound      RegiondevicecountGetCountResponseStatus = "TIME_INTERVAL_NO_DATA_FOUND"
)

type RegiondevicecountGetCountParams struct {
	// Ending timestamp for counting the number of devices in the area. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	Endtime param.Opt[time.Time] `json:"endtime,omitzero" format:"date-time"`
	// Starting timestamp for counting the number of devices in the area. It must
	// follow [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and
	// must have time zone.
	Starttime param.Opt[time.Time] `json:"starttime,omitzero" format:"date-time"`
	// The URL where the API response will be asynchronously delivered, using the HTTP
	// protocol.
	Sink        param.Opt[string]                   `json:"sink,omitzero" format:"uri"`
	XCorrelator param.Opt[string]                   `header:"x-correlator,omitzero" json:"-"`
	Area        RegiondevicecountGetCountParamsArea `json:"area,omitzero"`
	// This parameter is used to filter devices. Currently, two filtering criteria are
	// defined, `roamingStatus` and `deviceType`, which can be expanded in the future.
	// `IN` logic is used used for multiple filtering items within a single filtering
	// criterion, `AND` logic is used between multiple filtering criteria.
	//
	//   - If a filtering critera is not provided, it means that there is no need to
	//     filter this item.
	//   - At least one of the criteria must be provided,a filter without any criteria is
	//     not allowed.
	//   - If no filtering is required, this parameter does not need to be provided. For
	//     example
	//     ,`"filter":{"roamingStatus": ["roaming"],"deviceType": ["human device","IoT device"]}`
	//     means the API need to return the count of human network devices and IoT
	//     devices that are in roaming mode.`"filter":{"roamingStatus": ["non-roaming"]}`
	//     means that the API need to return the count of all devices that are not in
	//     roaming mode.
	Filter RegiondevicecountGetCountParamsFilter `json:"filter,omitzero"`
	// A sink credential provides authentication or authorization information necessary
	// to enable delivery of events to a target.
	SinkCredential RegiondevicecountGetCountParamsSinkCredential `json:"sinkCredential,omitzero"`
	paramObj
}

func (r RegiondevicecountGetCountParams) MarshalJSON() (data []byte, err error) {
	type shadow RegiondevicecountGetCountParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegiondevicecountGetCountParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property AreaType is required.
type RegiondevicecountGetCountParamsArea struct {
	// Type of this area. CIRCLE - The area is defined as a circle. POLYGON - The area
	// is defined as a polygon.
	//
	// Any of "CIRCLE", "POLYGON".
	AreaType string `json:"areaType,omitzero" api:"required"`
	paramObj
}

func (r RegiondevicecountGetCountParamsArea) MarshalJSON() (data []byte, err error) {
	type shadow RegiondevicecountGetCountParamsArea
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegiondevicecountGetCountParamsArea) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RegiondevicecountGetCountParamsArea](
		"areaType", "CIRCLE", "POLYGON",
	)
}

// This parameter is used to filter devices. Currently, two filtering criteria are
// defined, `roamingStatus` and `deviceType`, which can be expanded in the future.
// `IN` logic is used used for multiple filtering items within a single filtering
// criterion, `AND` logic is used between multiple filtering criteria.
//
//   - If a filtering critera is not provided, it means that there is no need to
//     filter this item.
//   - At least one of the criteria must be provided,a filter without any criteria is
//     not allowed.
//   - If no filtering is required, this parameter does not need to be provided. For
//     example
//     ,`"filter":{"roamingStatus": ["roaming"],"deviceType": ["human device","IoT device"]}`
//     means the API need to return the count of human network devices and IoT
//     devices that are in roaming mode.`"filter":{"roamingStatus": ["non-roaming"]}`
//     means that the API need to return the count of all devices that are not in
//     roaming mode.
type RegiondevicecountGetCountParamsFilter struct {
	// Filtering by device type, 'human device' represents the need to filter for human
	// network devices, 'IoT device' represents the need to filter for IoT devices, and
	// 'other' represents the need to filter for other types of devices.
	//
	// Any of "human device", "IoT device", "other".
	DeviceType []string `json:"deviceType,omitzero"`
	// Filter whether the device is in roaming mode,'roaming' represents the need to
	// filter devices that are in roaming mode,'non-roaming' represents the need to
	// filter devices that are not roaming.
	//
	// Any of "roaming", "non-roaming".
	RoamingStatus []string `json:"roamingStatus,omitzero"`
	paramObj
}

func (r RegiondevicecountGetCountParamsFilter) MarshalJSON() (data []byte, err error) {
	type shadow RegiondevicecountGetCountParamsFilter
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegiondevicecountGetCountParamsFilter) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sink credential provides authentication or authorization information necessary
// to enable delivery of events to a target.
//
// The property CredentialType is required.
type RegiondevicecountGetCountParamsSinkCredential struct {
	// The type of the credential. Note: Type of the credential - MUST be set to
	// ACCESSTOKEN for now
	//
	// Any of "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN".
	CredentialType string `json:"credentialType,omitzero" api:"required"`
	paramObj
}

func (r RegiondevicecountGetCountParamsSinkCredential) MarshalJSON() (data []byte, err error) {
	type shadow RegiondevicecountGetCountParamsSinkCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RegiondevicecountGetCountParamsSinkCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[RegiondevicecountGetCountParamsSinkCredential](
		"credentialType", "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN",
	)
}
