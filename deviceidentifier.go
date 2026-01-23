// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/stainless-sdks/camara-go/internal/apijson"
	shimjson "github.com/stainless-sdks/camara-go/internal/encoding/json"
	"github.com/stainless-sdks/camara-go/internal/requestconfig"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/stainless-sdks/camara-go/packages/param"
	"github.com/stainless-sdks/camara-go/packages/respjson"
)

// DeviceidentifierService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDeviceidentifierService] method instead.
type DeviceidentifierService struct {
	Options []option.RequestOption
}

// NewDeviceidentifierService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewDeviceidentifierService(opts ...option.RequestOption) (r DeviceidentifierService) {
	r = DeviceidentifierService{}
	r.Options = opts
	return
}

// Get details about the specific device being used by a given mobile subscriber
func (r *DeviceidentifierService) GetIdentifier(ctx context.Context, params DeviceidentifierGetIdentifierParams, opts ...option.RequestOption) (res *DeviceidentifierGetIdentifierResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "deviceidentifier/retrieve-identifier"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get a pseudonymous identifier for device being used by a given mobile subscriber
func (r *DeviceidentifierService) GetPpid(ctx context.Context, params DeviceidentifierGetPpidParams, opts ...option.RequestOption) (res *DeviceidentifierGetPpidResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "deviceidentifier/retrieve-ppid"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get details about the type of device being used by a given mobile subscriber
func (r *DeviceidentifierService) GetType(ctx context.Context, params DeviceidentifierGetTypeParams, opts ...option.RequestOption) (res *DeviceidentifierGetTypeResponse, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "deviceidentifier/retrieve-type"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Common request body to allow optional Device object to be passed
type DeviceIdentifierRequestBodyParam struct {
	// End-user equipment able to connect to a mobile network. Examples of devices
	// include smartphones or IoT sensors/actuators. The developer can choose to
	// provide the below specified device identifiers:
	//
	//   - `ipv4Address`
	//   - `ipv6Address`
	//   - `phoneNumber`
	//   - `networkAccessIdentifier` NOTE 1: The MNO might support only a subset of these
	//     options. The API invoker can provide multiple identifiers to be compatible
	//     across different MNOs. In this case the identifiers MUST belong to the same
	//     device. NOTE 2: For the current Commonalities release, we are enforcing that
	//     the networkAccessIdentifier is only part of the schema for future-proofing,
	//     and CAMARA does not currently allow its use. After the CAMARA meta-release
	//     work is concluded and the relevant issues are resolved, its use will need to
	//     be explicitly documented in the guidelines.
	Device DeviceIdentifierRequestBodyDeviceParam `json:"device,omitzero"`
	paramObj
}

func (r DeviceIdentifierRequestBodyParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceIdentifierRequestBodyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceIdentifierRequestBodyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user equipment able to connect to a mobile network. Examples of devices
// include smartphones or IoT sensors/actuators. The developer can choose to
// provide the below specified device identifiers:
//
//   - `ipv4Address`
//   - `ipv6Address`
//   - `phoneNumber`
//   - `networkAccessIdentifier` NOTE 1: The MNO might support only a subset of these
//     options. The API invoker can provide multiple identifiers to be compatible
//     across different MNOs. In this case the identifiers MUST belong to the same
//     device. NOTE 2: For the current Commonalities release, we are enforcing that
//     the networkAccessIdentifier is only part of the schema for future-proofing,
//     and CAMARA does not currently allow its use. After the CAMARA meta-release
//     work is concluded and the relevant issues are resolved, its use will need to
//     be explicitly documented in the guidelines.
type DeviceIdentifierRequestBodyDeviceParam struct {
	// The device should be identified by the observed IPv6 address, or by any single
	// IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to
	// the /64 prefix).
	Ipv6Address param.Opt[string] `json:"ipv6Address,omitzero" format:"ipv6"`
	// A public identifier addressing a subscription in a mobile network. In 3GPP
	// terminology, it corresponds to the GPSI formatted with the External Identifier
	// ({Local Identifier}@{Domain Identifier}). Unlike the telephone number, the
	// network access identifier is not subjected to portability ruling in force, and
	// is individually managed by each operator.
	NetworkAccessIdentifier param.Opt[string] `json:"networkAccessIdentifier,omitzero"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	// The device should be identified by either the public (observed) IP address and
	// port as seen by the application server, or the private (local) and any public
	// (observed) IP addresses in use by the device (this information can be obtained
	// by various means, for example from some DNS servers).
	//
	// If the allocated and observed IP addresses are the same (i.e. NAT is not in use)
	// then the same address should be specified for both publicAddress and
	// privateAddress.
	//
	// If NAT64 is in use, the device should be identified by its publicAddress and
	// publicPort, or separately by its allocated IPv6 address (field ipv6Address of
	// the Device object)
	//
	// In all cases, publicAddress must be specified, along with at least one of either
	// privateAddress or publicPort, dependent upon which is known. In general, mobile
	// devices cannot be identified by their public IPv4 address alone.
	Ipv4Address DeviceIdentifierRequestBodyDeviceIpv4AddressParam `json:"ipv4Address,omitzero"`
	paramObj
}

func (r DeviceIdentifierRequestBodyDeviceParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceIdentifierRequestBodyDeviceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceIdentifierRequestBodyDeviceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The device should be identified by either the public (observed) IP address and
// port as seen by the application server, or the private (local) and any public
// (observed) IP addresses in use by the device (this information can be obtained
// by various means, for example from some DNS servers).
//
// If the allocated and observed IP addresses are the same (i.e. NAT is not in use)
// then the same address should be specified for both publicAddress and
// privateAddress.
//
// If NAT64 is in use, the device should be identified by its publicAddress and
// publicPort, or separately by its allocated IPv6 address (field ipv6Address of
// the Device object)
//
// In all cases, publicAddress must be specified, along with at least one of either
// privateAddress or publicPort, dependent upon which is known. In general, mobile
// devices cannot be identified by their public IPv4 address alone.
type DeviceIdentifierRequestBodyDeviceIpv4AddressParam struct {
	// A single IPv4 address with no subnet mask
	PrivateAddress param.Opt[string] `json:"privateAddress,omitzero" format:"ipv4"`
	// A single IPv4 address with no subnet mask
	PublicAddress param.Opt[string] `json:"publicAddress,omitzero" format:"ipv4"`
	// TCP or UDP port number
	PublicPort param.Opt[int64] `json:"publicPort,omitzero"`
	paramObj
}

func (r DeviceIdentifierRequestBodyDeviceIpv4AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceIdentifierRequestBodyDeviceIpv4AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceIdentifierRequestBodyDeviceIpv4AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceidentifierGetIdentifierResponse struct {
	// The device subscription identifier that was used to identify the device whose
	// identifier is being returned. If this property is not present, then the device
	// subscription identifier specified in the request was used.
	Device DeviceidentifierGetIdentifierResponseDevice `json:"device"`
	// IMEI of the device
	Imei string `json:"imei"`
	// IMEISV of the device
	Imeisv string `json:"imeisv"`
	// Date and time that the information was last confirmed by the mobile operator to
	// be correct. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	LastChecked time.Time `json:"lastChecked" format:"date-time"`
	// Manufacturer of the device
	Manufacturer string `json:"manufacturer"`
	// Model of the device
	Model string `json:"model"`
	// IMEI TAC of the device
	Tac string `json:"tac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Device       respjson.Field
		Imei         respjson.Field
		Imeisv       respjson.Field
		LastChecked  respjson.Field
		Manufacturer respjson.Field
		Model        respjson.Field
		Tac          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceidentifierGetIdentifierResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceidentifierGetIdentifierResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The device subscription identifier that was used to identify the device whose
// identifier is being returned. If this property is not present, then the device
// subscription identifier specified in the request was used.
type DeviceidentifierGetIdentifierResponseDevice struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceidentifierGetIdentifierResponseDevice) RawJSON() string { return r.JSON.raw }
func (r *DeviceidentifierGetIdentifierResponseDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceidentifierGetPpidResponse struct {
	// The device subscription identifier that was used to identify the device whose
	// identifier is being returned. If this property is not present, then the device
	// subscription identifier specified in the request was used.
	Device DeviceidentifierGetPpidResponseDevice `json:"device"`
	// Date and time that the information was last confirmed by the mobile operator to
	// be correct. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	LastChecked time.Time `json:"lastChecked" format:"date-time"`
	// A PPID for the identified physical device
	Ppid string `json:"ppid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Device      respjson.Field
		LastChecked respjson.Field
		Ppid        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceidentifierGetPpidResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceidentifierGetPpidResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The device subscription identifier that was used to identify the device whose
// identifier is being returned. If this property is not present, then the device
// subscription identifier specified in the request was used.
type DeviceidentifierGetPpidResponseDevice struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceidentifierGetPpidResponseDevice) RawJSON() string { return r.JSON.raw }
func (r *DeviceidentifierGetPpidResponseDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceidentifierGetTypeResponse struct {
	// The device subscription identifier that was used to identify the device whose
	// identifier is being returned. If this property is not present, then the device
	// subscription identifier specified in the request was used.
	Device DeviceidentifierGetTypeResponseDevice `json:"device"`
	// Date and time that the information was last confirmed by the mobile operator to
	// be correct. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	LastChecked time.Time `json:"lastChecked" format:"date-time"`
	// Manufacturer of the device
	Manufacturer string `json:"manufacturer"`
	// Model of the device
	Model string `json:"model"`
	// IMEI TAC of the device
	Tac string `json:"tac"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Device       respjson.Field
		LastChecked  respjson.Field
		Manufacturer respjson.Field
		Model        respjson.Field
		Tac          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceidentifierGetTypeResponse) RawJSON() string { return r.JSON.raw }
func (r *DeviceidentifierGetTypeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The device subscription identifier that was used to identify the device whose
// identifier is being returned. If this property is not present, then the device
// subscription identifier specified in the request was used.
type DeviceidentifierGetTypeResponseDevice struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceidentifierGetTypeResponseDevice) RawJSON() string { return r.JSON.raw }
func (r *DeviceidentifierGetTypeResponseDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DeviceidentifierGetIdentifierParams struct {
	// Common request body to allow optional Device object to be passed
	DeviceIdentifierRequestBody DeviceIdentifierRequestBodyParam
	XCorrelator                 param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r DeviceidentifierGetIdentifierParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DeviceIdentifierRequestBody)
}
func (r *DeviceidentifierGetIdentifierParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DeviceIdentifierRequestBody)
}

type DeviceidentifierGetPpidParams struct {
	// Common request body to allow optional Device object to be passed
	DeviceIdentifierRequestBody DeviceIdentifierRequestBodyParam
	XCorrelator                 param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r DeviceidentifierGetPpidParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DeviceIdentifierRequestBody)
}
func (r *DeviceidentifierGetPpidParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DeviceIdentifierRequestBody)
}

type DeviceidentifierGetTypeParams struct {
	// Common request body to allow optional Device object to be passed
	DeviceIdentifierRequestBody DeviceIdentifierRequestBodyParam
	XCorrelator                 param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

func (r DeviceidentifierGetTypeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.DeviceIdentifierRequestBody)
}
func (r *DeviceidentifierGetTypeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.DeviceIdentifierRequestBody)
}
