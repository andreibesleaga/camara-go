// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"encoding/json"
	"errors"
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

// ConnectivityinsightSubscriptionService contains methods and other services that
// help with interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectivityinsightSubscriptionService] method instead.
type ConnectivityinsightSubscriptionService struct {
	Options []option.RequestOption
}

// NewConnectivityinsightSubscriptionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectivityinsightSubscriptionService(opts ...option.RequestOption) (r ConnectivityinsightSubscriptionService) {
	r = ConnectivityinsightSubscriptionService{}
	r.Options = opts
	return
}

// Create a Connectivity insights subscription for a device
func (r *ConnectivityinsightSubscriptionService) New(ctx context.Context, params ConnectivityinsightSubscriptionNewParams, opts ...option.RequestOption) (res *Subscription, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectivityinsights/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Retrieve a given subscription by ID
func (r *ConnectivityinsightSubscriptionService) Get(ctx context.Context, subscriptionID string, query ConnectivityinsightSubscriptionGetParams, opts ...option.RequestOption) (res *Subscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("connectivityinsights/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Operation to list subscriptions authorized to be retrieved by the provided
// access token.
func (r *ConnectivityinsightSubscriptionService) List(ctx context.Context, query ConnectivityinsightSubscriptionListParams, opts ...option.RequestOption) (res *[]Subscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectivityinsights/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a given subscription by ID
func (r *ConnectivityinsightSubscriptionService) Delete(ctx context.Context, subscriptionID string, body ConnectivityinsightSubscriptionDeleteParams, opts ...option.RequestOption) (res *ConnectivityinsightSubscriptionDeleteResponse, err error) {
	if !param.IsOmitted(body.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", body.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("connectivityinsights/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Implementation-specific configuration parameters needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
// type attributes must be defined in `subscriptionDetail` Note: if a request is
// performed for several event type, all subscribed event will use same `config`
// parameters.
type Config struct {
	// The detail of the requested event subscription
	SubscriptionDetail ConfigSubscriptionDetail `json:"subscriptionDetail,required"`
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request.
	InitialEvent bool `json:"initialEvent"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer. Up to API project decision to keep it. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	SubscriptionExpireTime time.Time `json:"subscriptionExpireTime" format:"date-time"`
	// Identifies the maximum number of event reports to be generated (>=1) requested
	// by the API consumer - Once this number is reached, the subscription ends. Note
	// on combined usage of `initialEvent` and `subscriptionMaxEvents`: If an event is
	// triggered following `initialEvent` set to `true`, this event will be counted
	// towards `subscriptionMaxEvents`.
	SubscriptionMaxEvents int64 `json:"subscriptionMaxEvents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionDetail     respjson.Field
		InitialEvent           respjson.Field
		SubscriptionExpireTime respjson.Field
		SubscriptionMaxEvents  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Config) RawJSON() string { return r.JSON.raw }
func (r *Config) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Config to a ConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConfigParam.Overrides()
func (r Config) ToParam() ConfigParam {
	return param.Override[ConfigParam](json.RawMessage(r.RawJSON()))
}

// The detail of the requested event subscription
type ConfigSubscriptionDetail struct {
	// Identifier for the Application Profile
	ApplicationProfileID string `json:"applicationProfileId,required" format:"uuid"`
	// End-user equipment able to connect to a mobile network. Examples of devices
	// include smartphones or IoT sensors/actuators. The developer can choose to
	// provide the below specified device identifiers: _ `ipv4Address` _ `ipv6Address`
	// _ `phoneNumber` _ `networkAccessIdentifier` NOTE1: the network operator might
	// support only a subset of these options. The API invoker can provide multiple
	// identifiers to be compatible across different network operators. In this case
	// the identifiers MUST belong to the same device. NOTE2: as for this Commonalities
	// release, we are enforcing that the networkAccessIdentifier is only part of the
	// schema for future-proofing, and CAMARA does not currently allow its use. After
	// the CAMARA meta-release work is concluded and the relevant issues are resolved,
	// its use will need to be explicitly documented in the guidelines.
	Device ConfigSubscriptionDetailDevice `json:"device,required"`
	// A server hosting backend applications to deliver some business logic to clients.
	//
	// The developer can choose to provide the below specified device identifiers:
	//
	// - `ipv4Address`
	// - `ipv6Address`
	//
	// The Operator will use this information to calculate the end to end network
	// performance in scenarios where its feasible.
	ApplicationServer ConfigSubscriptionDetailApplicationServer `json:"applicationServer"`
	// Specification of several TCP or UDP ports
	ApplicationServerPorts ConfigSubscriptionDetailApplicationServerPorts `json:"applicationServerPorts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ApplicationProfileID   respjson.Field
		Device                 respjson.Field
		ApplicationServer      respjson.Field
		ApplicationServerPorts respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSubscriptionDetail) RawJSON() string { return r.JSON.raw }
func (r *ConfigSubscriptionDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user equipment able to connect to a mobile network. Examples of devices
// include smartphones or IoT sensors/actuators. The developer can choose to
// provide the below specified device identifiers: _ `ipv4Address` _ `ipv6Address`
// _ `phoneNumber` _ `networkAccessIdentifier` NOTE1: the network operator might
// support only a subset of these options. The API invoker can provide multiple
// identifiers to be compatible across different network operators. In this case
// the identifiers MUST belong to the same device. NOTE2: as for this Commonalities
// release, we are enforcing that the networkAccessIdentifier is only part of the
// schema for future-proofing, and CAMARA does not currently allow its use. After
// the CAMARA meta-release work is concluded and the relevant issues are resolved,
// its use will need to be explicitly documented in the guidelines.
type ConfigSubscriptionDetailDevice struct {
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
	Ipv4Address ConfigSubscriptionDetailDeviceIpv4Address `json:"ipv4Address"`
	// The device should be identified by the observed IPv6 address, or by any single
	// IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to
	// the /64 prefix).
	//
	// The session shall apply to all IP flows between the device subnet and the
	// specified application server, unless further restricted by the optional
	// parameters devicePorts or applicationServerPorts.
	Ipv6Address string `json:"ipv6Address" format:"ipv6"`
	// A public identifier addressing a subscription in a mobile network. In 3GPP
	// terminology, it corresponds to the GPSI formatted with the External Identifier
	// ({Local Identifier}@{Domain Identifier}). Unlike the telephone number, the
	// network access identifier is not subjected to portability ruling in force, and
	// is individually managed by each operator.
	NetworkAccessIdentifier string `json:"networkAccessIdentifier"`
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber string `json:"phoneNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ipv4Address             respjson.Field
		Ipv6Address             respjson.Field
		NetworkAccessIdentifier respjson.Field
		PhoneNumber             respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSubscriptionDetailDevice) RawJSON() string { return r.JSON.raw }
func (r *ConfigSubscriptionDetailDevice) UnmarshalJSON(data []byte) error {
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
type ConfigSubscriptionDetailDeviceIpv4Address struct {
	// A single IPv4 address with no subnet mask
	PrivateAddress string `json:"privateAddress" format:"ipv4"`
	// A single IPv4 address with no subnet mask
	PublicAddress string `json:"publicAddress" format:"ipv4"`
	// TCP or UDP port number
	PublicPort int64 `json:"publicPort"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PrivateAddress respjson.Field
		PublicAddress  respjson.Field
		PublicPort     respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSubscriptionDetailDeviceIpv4Address) RawJSON() string { return r.JSON.raw }
func (r *ConfigSubscriptionDetailDeviceIpv4Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A server hosting backend applications to deliver some business logic to clients.
//
// The developer can choose to provide the below specified device identifiers:
//
// - `ipv4Address`
// - `ipv6Address`
//
// The Operator will use this information to calculate the end to end network
// performance in scenarios where its feasible.
type ConfigSubscriptionDetailApplicationServer struct {
	// IPv4 address may be specified in form <address/mask> as:
	//
	//   - address - an IPv4 number in dotted-quad form 1.2.3.4. Only this exact IP
	//     number will match the flow control rule.
	//   - address/mask - an IP number as above with a mask width of the form 1.2.3.4/24.
	//     In this case, all IP numbers from 1.2.3.0 to 1.2.3.255 will match. The bit
	//     width MUST be valid for the IP version.
	Ipv4Address string `json:"ipv4Address"`
	// IPv6 address may be specified in form <address/mask> as:
	//
	// - address - The /128 subnet is optional for single addresses:
	//   - 2001:db8:85a3:8d3:1319:8a2e:370:7344
	//   - 2001:db8:85a3:8d3:1319:8a2e:370:7344/128
	//
	// - address/mask - an IP v6 number with a mask:
	//   - 2001:db8:85a3:8d3::0/64
	//   - 2001:db8:85a3:8d3::/64
	Ipv6Address string `json:"ipv6Address"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ipv4Address respjson.Field
		Ipv6Address respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSubscriptionDetailApplicationServer) RawJSON() string { return r.JSON.raw }
func (r *ConfigSubscriptionDetailApplicationServer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification of several TCP or UDP ports
type ConfigSubscriptionDetailApplicationServerPorts struct {
	// Array of TCP or UDP ports
	Ports []int64 `json:"ports"`
	// Range of TCP or UDP ports
	Ranges []ConfigSubscriptionDetailApplicationServerPortsRange `json:"ranges"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Ports       respjson.Field
		Ranges      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSubscriptionDetailApplicationServerPorts) RawJSON() string { return r.JSON.raw }
func (r *ConfigSubscriptionDetailApplicationServerPorts) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConfigSubscriptionDetailApplicationServerPortsRange struct {
	// TCP or UDP port number
	From int64 `json:"from,required"`
	// TCP or UDP port number
	To int64 `json:"to,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		From        respjson.Field
		To          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigSubscriptionDetailApplicationServerPortsRange) RawJSON() string { return r.JSON.raw }
func (r *ConfigSubscriptionDetailApplicationServerPortsRange) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Implementation-specific configuration parameters needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
// type attributes must be defined in `subscriptionDetail` Note: if a request is
// performed for several event type, all subscribed event will use same `config`
// parameters.
//
// The property SubscriptionDetail is required.
type ConfigParam struct {
	// The detail of the requested event subscription
	SubscriptionDetail ConfigSubscriptionDetailParam `json:"subscriptionDetail,omitzero,required"`
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request.
	InitialEvent param.Opt[bool] `json:"initialEvent,omitzero"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer. Up to API project decision to keep it. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	SubscriptionExpireTime param.Opt[time.Time] `json:"subscriptionExpireTime,omitzero" format:"date-time"`
	// Identifies the maximum number of event reports to be generated (>=1) requested
	// by the API consumer - Once this number is reached, the subscription ends. Note
	// on combined usage of `initialEvent` and `subscriptionMaxEvents`: If an event is
	// triggered following `initialEvent` set to `true`, this event will be counted
	// towards `subscriptionMaxEvents`.
	SubscriptionMaxEvents param.Opt[int64] `json:"subscriptionMaxEvents,omitzero"`
	paramObj
}

func (r ConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The detail of the requested event subscription
//
// The properties ApplicationProfileID, Device are required.
type ConfigSubscriptionDetailParam struct {
	// Identifier for the Application Profile
	ApplicationProfileID string `json:"applicationProfileId,required" format:"uuid"`
	// End-user equipment able to connect to a mobile network. Examples of devices
	// include smartphones or IoT sensors/actuators. The developer can choose to
	// provide the below specified device identifiers: _ `ipv4Address` _ `ipv6Address`
	// _ `phoneNumber` _ `networkAccessIdentifier` NOTE1: the network operator might
	// support only a subset of these options. The API invoker can provide multiple
	// identifiers to be compatible across different network operators. In this case
	// the identifiers MUST belong to the same device. NOTE2: as for this Commonalities
	// release, we are enforcing that the networkAccessIdentifier is only part of the
	// schema for future-proofing, and CAMARA does not currently allow its use. After
	// the CAMARA meta-release work is concluded and the relevant issues are resolved,
	// its use will need to be explicitly documented in the guidelines.
	Device ConfigSubscriptionDetailDeviceParam `json:"device,omitzero,required"`
	// A server hosting backend applications to deliver some business logic to clients.
	//
	// The developer can choose to provide the below specified device identifiers:
	//
	// - `ipv4Address`
	// - `ipv6Address`
	//
	// The Operator will use this information to calculate the end to end network
	// performance in scenarios where its feasible.
	ApplicationServer ConfigSubscriptionDetailApplicationServerParam `json:"applicationServer,omitzero"`
	// Specification of several TCP or UDP ports
	ApplicationServerPorts ConfigSubscriptionDetailApplicationServerPortsParam `json:"applicationServerPorts,omitzero"`
	paramObj
}

func (r ConfigSubscriptionDetailParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSubscriptionDetailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSubscriptionDetailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user equipment able to connect to a mobile network. Examples of devices
// include smartphones or IoT sensors/actuators. The developer can choose to
// provide the below specified device identifiers: _ `ipv4Address` _ `ipv6Address`
// _ `phoneNumber` _ `networkAccessIdentifier` NOTE1: the network operator might
// support only a subset of these options. The API invoker can provide multiple
// identifiers to be compatible across different network operators. In this case
// the identifiers MUST belong to the same device. NOTE2: as for this Commonalities
// release, we are enforcing that the networkAccessIdentifier is only part of the
// schema for future-proofing, and CAMARA does not currently allow its use. After
// the CAMARA meta-release work is concluded and the relevant issues are resolved,
// its use will need to be explicitly documented in the guidelines.
type ConfigSubscriptionDetailDeviceParam struct {
	// The device should be identified by the observed IPv6 address, or by any single
	// IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to
	// the /64 prefix).
	//
	// The session shall apply to all IP flows between the device subnet and the
	// specified application server, unless further restricted by the optional
	// parameters devicePorts or applicationServerPorts.
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
	Ipv4Address ConfigSubscriptionDetailDeviceIpv4AddressParam `json:"ipv4Address,omitzero"`
	paramObj
}

func (r ConfigSubscriptionDetailDeviceParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSubscriptionDetailDeviceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSubscriptionDetailDeviceParam) UnmarshalJSON(data []byte) error {
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
type ConfigSubscriptionDetailDeviceIpv4AddressParam struct {
	// A single IPv4 address with no subnet mask
	PrivateAddress param.Opt[string] `json:"privateAddress,omitzero" format:"ipv4"`
	// A single IPv4 address with no subnet mask
	PublicAddress param.Opt[string] `json:"publicAddress,omitzero" format:"ipv4"`
	// TCP or UDP port number
	PublicPort param.Opt[int64] `json:"publicPort,omitzero"`
	paramObj
}

func (r ConfigSubscriptionDetailDeviceIpv4AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSubscriptionDetailDeviceIpv4AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSubscriptionDetailDeviceIpv4AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A server hosting backend applications to deliver some business logic to clients.
//
// The developer can choose to provide the below specified device identifiers:
//
// - `ipv4Address`
// - `ipv6Address`
//
// The Operator will use this information to calculate the end to end network
// performance in scenarios where its feasible.
type ConfigSubscriptionDetailApplicationServerParam struct {
	// IPv4 address may be specified in form <address/mask> as:
	//
	//   - address - an IPv4 number in dotted-quad form 1.2.3.4. Only this exact IP
	//     number will match the flow control rule.
	//   - address/mask - an IP number as above with a mask width of the form 1.2.3.4/24.
	//     In this case, all IP numbers from 1.2.3.0 to 1.2.3.255 will match. The bit
	//     width MUST be valid for the IP version.
	Ipv4Address param.Opt[string] `json:"ipv4Address,omitzero"`
	// IPv6 address may be specified in form <address/mask> as:
	//
	// - address - The /128 subnet is optional for single addresses:
	//   - 2001:db8:85a3:8d3:1319:8a2e:370:7344
	//   - 2001:db8:85a3:8d3:1319:8a2e:370:7344/128
	//
	// - address/mask - an IP v6 number with a mask:
	//   - 2001:db8:85a3:8d3::0/64
	//   - 2001:db8:85a3:8d3::/64
	Ipv6Address param.Opt[string] `json:"ipv6Address,omitzero"`
	paramObj
}

func (r ConfigSubscriptionDetailApplicationServerParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSubscriptionDetailApplicationServerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSubscriptionDetailApplicationServerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification of several TCP or UDP ports
type ConfigSubscriptionDetailApplicationServerPortsParam struct {
	// Array of TCP or UDP ports
	Ports []int64 `json:"ports,omitzero"`
	// Range of TCP or UDP ports
	Ranges []ConfigSubscriptionDetailApplicationServerPortsRangeParam `json:"ranges,omitzero"`
	paramObj
}

func (r ConfigSubscriptionDetailApplicationServerPortsParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSubscriptionDetailApplicationServerPortsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSubscriptionDetailApplicationServerPortsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties From, To are required.
type ConfigSubscriptionDetailApplicationServerPortsRangeParam struct {
	// TCP or UDP port number
	From int64 `json:"from,required"`
	// TCP or UDP port number
	To int64 `json:"to,required"`
	paramObj
}

func (r ConfigSubscriptionDetailApplicationServerPortsRangeParam) MarshalJSON() (data []byte, err error) {
	type shadow ConfigSubscriptionDetailApplicationServerPortsRangeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConfigSubscriptionDetailApplicationServerPortsRangeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// event-type - Event triggered when an event-type event occurred
type EventType string

const (
	EventTypeOrgCamaraprojectConnectivityInsightsSubscriptionsV0NetworkQuality EventType = "org.camaraproject.connectivity-insights-subscriptions.v0.network-quality"
)

// Identifier of a delivery protocol. Only HTTP is allowed for now
type Protocol string

const (
	ProtocolHTTP  Protocol = "HTTP"
	ProtocolMqtt3 Protocol = "MQTT3"
	ProtocolMqtt5 Protocol = "MQTT5"
	ProtocolAmqp  Protocol = "AMQP"
	ProtocolNats  Protocol = "NATS"
	ProtocolKafka Protocol = "KAFKA"
)

// Represents a event-type subscription.
type Subscription struct {
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
	// type attributes must be defined in `subscriptionDetail` Note: if a request is
	// performed for several event type, all subscribed event will use same `config`
	// parameters.
	Config Config `json:"config,required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol Protocol `json:"protocol,required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink,required" format:"uri"`
	// Date when the event subscription will begin/began It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	StartsAt time.Time `json:"startsAt,required" format:"date-time"`
	// Camara Event types eligible to be delivered by this subscription.
	Types []EventType `json:"types,required"`
	// Date when the event subscription will expire. Only provided when
	// `subscriptionExpireTime` is indicated by API client or Telco Operator has
	// specific policy about that. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// Current status of the subscription - Management of Subscription State engine is
	// not mandatory for now. Note not all statuses may be considered to be
	// implemented. Details:
	//
	//   - `ACTIVATION_REQUESTED`: Subscription creation (POST) is triggered but
	//     subscription creation process is not finished yet.
	//   - `ACTIVE`: Subscription creation process is completed. Subscription is fully
	//     operative.
	//   - `DEACTIVE`: Subscription is temporarily inactive, but its workflow logic is
	//     not deleted.
	//   - `EXPIRED`: Subscription is ended (no longer active). This status applies when
	//     subscription is ended due to `SUBSCRIPTION_EXPIRED` or `ACCESS_TOKEN_EXPIRED`
	//     event.
	//   - `DELETED`: Subscription is ended as deleted (no longer active). This status
	//     applies when subscription information is kept (i.e. subscription workflow is
	//     no longer active but its metainformation is kept).
	//
	// Any of "ACTIVATION_REQUESTED", "ACTIVE", "EXPIRED", "DEACTIVE", "DELETED".
	Status SubscriptionStatus `json:"status"`
	// When this information is contained within an event notification, it SHALL be
	// referred to as `subscriptionId` as per the Commonalities Event Notification
	// Model.
	SubscriptionID string `json:"subscriptionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config         respjson.Field
		Protocol       respjson.Field
		Sink           respjson.Field
		StartsAt       respjson.Field
		Types          respjson.Field
		ExpiresAt      respjson.Field
		Status         respjson.Field
		SubscriptionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subscription) RawJSON() string { return r.JSON.raw }
func (r *Subscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current status of the subscription - Management of Subscription State engine is
// not mandatory for now. Note not all statuses may be considered to be
// implemented. Details:
//
//   - `ACTIVATION_REQUESTED`: Subscription creation (POST) is triggered but
//     subscription creation process is not finished yet.
//   - `ACTIVE`: Subscription creation process is completed. Subscription is fully
//     operative.
//   - `DEACTIVE`: Subscription is temporarily inactive, but its workflow logic is
//     not deleted.
//   - `EXPIRED`: Subscription is ended (no longer active). This status applies when
//     subscription is ended due to `SUBSCRIPTION_EXPIRED` or `ACCESS_TOKEN_EXPIRED`
//     event.
//   - `DELETED`: Subscription is ended as deleted (no longer active). This status
//     applies when subscription information is kept (i.e. subscription workflow is
//     no longer active but its metainformation is kept).
type SubscriptionStatus string

const (
	SubscriptionStatusActivationRequested SubscriptionStatus = "ACTIVATION_REQUESTED"
	SubscriptionStatusActive              SubscriptionStatus = "ACTIVE"
	SubscriptionStatusExpired             SubscriptionStatus = "EXPIRED"
	SubscriptionStatusDeactive            SubscriptionStatus = "DEACTIVE"
	SubscriptionStatusDeleted             SubscriptionStatus = "DELETED"
)

// Response for a event-type subscription request managed asynchronously (Creation
// or Deletion)
type ConnectivityinsightSubscriptionDeleteResponse struct {
	// When this information is contained within an event notification, it SHALL be
	// referred to as `subscriptionId` as per the Commonalities Event Notification
	// Model.
	SubscriptionID string `json:"subscriptionId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionID respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectivityinsightSubscriptionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectivityinsightSubscriptionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectivityinsightSubscriptionNewParams struct {
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
	// type attributes must be defined in `subscriptionDetail` Note: if a request is
	// performed for several event type, all subscribed event will use same `config`
	// parameters.
	Config ConfigParam `json:"config,omitzero,required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol Protocol `json:"protocol,omitzero,required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink,required" format:"uri"`
	// Camara Event types eligible to be delivered by this subscription.
	Types       []EventType       `json:"types,omitzero,required"`
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	// A sink credential provides authentication or authorization information
	SinkCredential ConnectivityinsightSubscriptionNewParamsSinkCredential `json:"sinkCredential,omitzero"`
	paramObj
}

func (r ConnectivityinsightSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectivityinsightSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectivityinsightSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sink credential provides authentication or authorization information
//
// The property CredentialType is required.
type ConnectivityinsightSubscriptionNewParamsSinkCredential struct {
	// The type of the credential. Note: Type of the credential - MUST be set to
	// ACCESSTOKEN for now
	//
	// Any of "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN".
	CredentialType string `json:"credentialType,omitzero,required"`
	paramObj
}

func (r ConnectivityinsightSubscriptionNewParamsSinkCredential) MarshalJSON() (data []byte, err error) {
	type shadow ConnectivityinsightSubscriptionNewParamsSinkCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectivityinsightSubscriptionNewParamsSinkCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConnectivityinsightSubscriptionNewParamsSinkCredential](
		"credentialType", "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN",
	)
}

type ConnectivityinsightSubscriptionGetParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type ConnectivityinsightSubscriptionListParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type ConnectivityinsightSubscriptionDeleteParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}
