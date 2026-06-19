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

	"github.com/andreibesleaga/camara-go/internal/apijson"
	"github.com/andreibesleaga/camara-go/internal/requestconfig"
	"github.com/andreibesleaga/camara-go/option"
	"github.com/andreibesleaga/camara-go/packages/param"
	"github.com/andreibesleaga/camara-go/packages/respjson"
)

// Device Geofencing Subscriptions
//
// DevicelocationSubscriptionService contains methods and other services that help
// with interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicelocationSubscriptionService] method instead.
type DevicelocationSubscriptionService struct {
	Options []option.RequestOption
}

// NewDevicelocationSubscriptionService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewDevicelocationSubscriptionService(opts ...option.RequestOption) (r DevicelocationSubscriptionService) {
	r = DevicelocationSubscriptionService{}
	r.Options = opts
	return
}

// Create a subscription for a device to receive notifications when the device
// enters or exits a specified area.
func (r *DevicelocationSubscriptionService) New(ctx context.Context, params DevicelocationSubscriptionNewParams, opts ...option.RequestOption) (res *DeviceLocationSubscription, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "devicelocation/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve Geofencing subscription information for a given subscription ID.
func (r *DevicelocationSubscriptionService) Get(ctx context.Context, subscriptionID string, query DevicelocationSubscriptionGetParams, opts ...option.RequestOption) (res *DeviceLocationSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devicelocation/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a list of geofencing event subscription(s).
func (r *DevicelocationSubscriptionService) List(ctx context.Context, query DevicelocationSubscriptionListParams, opts ...option.RequestOption) (res *[]DeviceLocationSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "devicelocation/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a given Geofencing subscription.
func (r *DevicelocationSubscriptionService) Delete(ctx context.Context, subscriptionID string, body DevicelocationSubscriptionDeleteParams, opts ...option.RequestOption) (res *DevicelocationSubscriptionDeleteResponse, err error) {
	if !param.IsOmitted(body.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", body.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devicelocation/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// The geofencing area where the monitor is active. This area is specified by API
// consumers in the subscription request. The same area definition is included in
// event notifications without any modifications.
type DeviceLocationArea struct {
	// Type of this area. CIRCLE - The area is defined as a circle.
	//
	// Any of "CIRCLE".
	AreaType DeviceLocationAreaAreaType `json:"areaType" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AreaType    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceLocationArea) RawJSON() string { return r.JSON.raw }
func (r *DeviceLocationArea) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceLocationArea to a DeviceLocationAreaParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceLocationAreaParam.Overrides()
func (r DeviceLocationArea) ToParam() DeviceLocationAreaParam {
	return param.Override[DeviceLocationAreaParam](json.RawMessage(r.RawJSON()))
}

// Type of this area. CIRCLE - The area is defined as a circle.
type DeviceLocationAreaAreaType string

const (
	DeviceLocationAreaAreaTypeCircle DeviceLocationAreaAreaType = "CIRCLE"
)

// The geofencing area where the monitor is active. This area is specified by API
// consumers in the subscription request. The same area definition is included in
// event notifications without any modifications.
//
// The property AreaType is required.
type DeviceLocationAreaParam struct {
	// Type of this area. CIRCLE - The area is defined as a circle.
	//
	// Any of "CIRCLE".
	AreaType DeviceLocationAreaAreaType `json:"areaType,omitzero" api:"required"`
	paramObj
}

func (r DeviceLocationAreaParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceLocationAreaParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceLocationAreaParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Implementation-specific configuration parameters are needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`.
type DeviceLocationConfig struct {
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request. Example:
	// Consumer request area entered event. If consumer sets initialEvent to true and
	// device is already in the geofence, an event is triggered.
	InitialEvent bool `json:"initialEvent"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer. It must follow
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
		InitialEvent           respjson.Field
		SubscriptionExpireTime respjson.Field
		SubscriptionMaxEvents  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceLocationConfig) RawJSON() string { return r.JSON.raw }
func (r *DeviceLocationConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceLocationConfig to a DeviceLocationConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceLocationConfigParam.Overrides()
func (r DeviceLocationConfig) ToParam() DeviceLocationConfigParam {
	return param.Override[DeviceLocationConfigParam](json.RawMessage(r.RawJSON()))
}

// Implementation-specific configuration parameters are needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`.
type DeviceLocationConfigParam struct {
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request. Example:
	// Consumer request area entered event. If consumer sets initialEvent to true and
	// device is already in the geofence, an event is triggered.
	InitialEvent param.Opt[bool] `json:"initialEvent,omitzero"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer. It must follow
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

func (r DeviceLocationConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceLocationConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceLocationConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user device able to connect to a mobile network. Examples of devices include
// smartphones or IoT sensors/actuators.
//
// The developer can choose to provide the below specified device identifiers:
//
// - `ipv4Address`
// - `ipv6Address`
// - `phoneNumber`
// - `networkAccessIdentifier`
//
// NOTE1: the API provider might support only a subset of these options. The API
// consumer can provide multiple identifiers to be compatible across different API
// providers. In this case the identifiers MUST belong to the same device. Where
// more than one device identifier is provided, only one identifier will be
// selected by the implementation and this choice indicated to the API consumer in
// the response or event. NOTE2: as for this Commonalities release, we are
// enforcing that the networkAccessIdentifier is only part of the schema for
// future-proofing, and CAMARA does not currently allow its use. After the CAMARA
// meta-release work is concluded and the relevant issues are resolved, its use
// will need to be explicitly documented in the guidelines.
type DeviceLocationDevice struct {
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
	Ipv4Address DeviceLocationDeviceIpv4Address `json:"ipv4Address"`
	// The device should be identified by the observed IPv6 address, or by any single
	// IPv6 address from within the subnet allocated to the device (e.g. adding ::0 to
	// the /64 prefix).
	Ipv6Address string `json:"ipv6Address" format:"ipv6"`
	// A public identifier addressing a subscription in a mobile network. In 3GPP
	// terminology, it corresponds to the GPSI formatted with the External Identifier
	// ({Local Identifier}@{Domain Identifier}). Unlike the telephone number, the
	// network access identifier is not subjected to portability ruling in force, and
	// is individually managed by each operator.
	NetworkAccessIdentifier string `json:"networkAccessIdentifier"`
	// A public identifier addressing a telephone subscription. In mobile networks, it
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
func (r DeviceLocationDevice) RawJSON() string { return r.JSON.raw }
func (r *DeviceLocationDevice) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceLocationDevice to a DeviceLocationDeviceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceLocationDeviceParam.Overrides()
func (r DeviceLocationDevice) ToParam() DeviceLocationDeviceParam {
	return param.Override[DeviceLocationDeviceParam](json.RawMessage(r.RawJSON()))
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
type DeviceLocationDeviceIpv4Address struct {
	// A single IPv4 address with no subnet mask.
	PrivateAddress string `json:"privateAddress" format:"ipv4"`
	// A single IPv4 address with no subnet mask.
	PublicAddress string `json:"publicAddress" format:"ipv4"`
	// TCP or UDP port number.
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
func (r DeviceLocationDeviceIpv4Address) RawJSON() string { return r.JSON.raw }
func (r *DeviceLocationDeviceIpv4Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user device able to connect to a mobile network. Examples of devices include
// smartphones or IoT sensors/actuators.
//
// The developer can choose to provide the below specified device identifiers:
//
// - `ipv4Address`
// - `ipv6Address`
// - `phoneNumber`
// - `networkAccessIdentifier`
//
// NOTE1: the API provider might support only a subset of these options. The API
// consumer can provide multiple identifiers to be compatible across different API
// providers. In this case the identifiers MUST belong to the same device. Where
// more than one device identifier is provided, only one identifier will be
// selected by the implementation and this choice indicated to the API consumer in
// the response or event. NOTE2: as for this Commonalities release, we are
// enforcing that the networkAccessIdentifier is only part of the schema for
// future-proofing, and CAMARA does not currently allow its use. After the CAMARA
// meta-release work is concluded and the relevant issues are resolved, its use
// will need to be explicitly documented in the guidelines.
type DeviceLocationDeviceParam struct {
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
	// A public identifier addressing a telephone subscription. In mobile networks, it
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
	Ipv4Address DeviceLocationDeviceIpv4AddressParam `json:"ipv4Address,omitzero"`
	paramObj
}

func (r DeviceLocationDeviceParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceLocationDeviceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceLocationDeviceParam) UnmarshalJSON(data []byte) error {
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
type DeviceLocationDeviceIpv4AddressParam struct {
	// A single IPv4 address with no subnet mask.
	PrivateAddress param.Opt[string] `json:"privateAddress,omitzero" format:"ipv4"`
	// A single IPv4 address with no subnet mask.
	PublicAddress param.Opt[string] `json:"publicAddress,omitzero" format:"ipv4"`
	// TCP or UDP port number.
	PublicPort param.Opt[int64] `json:"publicPort,omitzero"`
	paramObj
}

func (r DeviceLocationDeviceIpv4AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceLocationDeviceIpv4AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceLocationDeviceIpv4AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of a delivery protocol. Only HTTP is allowed for now.
type DeviceLocationProtocol string

const (
	DeviceLocationProtocolHTTP  DeviceLocationProtocol = "HTTP"
	DeviceLocationProtocolMqtt3 DeviceLocationProtocol = "MQTT3"
	DeviceLocationProtocolMqtt5 DeviceLocationProtocol = "MQTT5"
	DeviceLocationProtocolAmqp  DeviceLocationProtocol = "AMQP"
	DeviceLocationProtocolNats  DeviceLocationProtocol = "NATS"
	DeviceLocationProtocolKafka DeviceLocationProtocol = "KAFKA"
)

// Represents a event-type subscription.
type DeviceLocationSubscription struct {
	// The unique identifier of the subscription in the scope of the subscription
	// manager. When this information is contained within an event notification, this
	// concept SHALL be referred as subscriptionId as per Commonalities Event
	// Notification Model.
	ID string `json:"id" api:"required"`
	// Implementation-specific configuration parameters are needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`.
	Config DeviceLocationSubscriptionConfig `json:"config" api:"required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now.
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol DeviceLocationProtocol `json:"protocol" api:"required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink" api:"required" format:"uri"`
	// Date when the event subscription will begin/began It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	StartsAt time.Time `json:"startsAt" api:"required" format:"date-time"`
	// Camara Event types eligible to be delivered by this subscription. Note: As of
	// now we enforce to have only event type per subscription.
	Types []DeviceLocationSubscriptionEventType `json:"types" api:"required"`
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
	//   - `INACTIVE`: Subscription is temporarily inactive, but its workflow logic is
	//     not deleted.
	//   - `EXPIRED`: Subscription is ended (no longer active). This status applies when
	//     subscription is ended due to `SUBSCRIPTION_EXPIRED` or `ACCESS_TOKEN_EXPIRED`
	//     event.
	//   - `DELETED`: Subscription is ended as deleted (no longer active). This status
	//     applies when subscription information is kept (i.e. subscription workflow is
	//     no longer active but its meta-information is kept).
	//
	// Any of "ACTIVATION_REQUESTED", "ACTIVE", "EXPIRED", "INACTIVE", "DELETED".
	Status DeviceLocationSubscriptionStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Config      respjson.Field
		Protocol    respjson.Field
		Sink        respjson.Field
		StartsAt    respjson.Field
		Types       respjson.Field
		ExpiresAt   respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceLocationSubscription) RawJSON() string { return r.JSON.raw }
func (r *DeviceLocationSubscription) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Implementation-specific configuration parameters are needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`.
type DeviceLocationSubscriptionConfig struct {
	// The detail of the event subscription granted by the implementation.
	SubscriptionDetail DeviceLocationSubscriptionConfigSubscriptionDetail `json:"subscriptionDetail" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionDetail respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
	DeviceLocationConfig
}

// Returns the unmodified JSON received from the API
func (r DeviceLocationSubscriptionConfig) RawJSON() string { return r.JSON.raw }
func (r *DeviceLocationSubscriptionConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The detail of the event subscription granted by the implementation.
type DeviceLocationSubscriptionConfigSubscriptionDetail struct {
	// The geofencing area where the monitor is active. This area is specified by API
	// consumers in the subscription request. The same area definition is included in
	// event notifications without any modifications.
	Area DeviceLocationArea `json:"area" api:"required"`
	// End-user device able to connect to a mobile network. Examples of devices include
	// smartphones or IoT sensors/actuators.
	//
	// The developer can choose to provide the below specified device identifiers:
	//
	// - `ipv4Address`
	// - `ipv6Address`
	// - `phoneNumber`
	// - `networkAccessIdentifier`
	//
	// NOTE1: the API provider might support only a subset of these options. The API
	// consumer can provide multiple identifiers to be compatible across different API
	// providers. In this case the identifiers MUST belong to the same device. Where
	// more than one device identifier is provided, only one identifier will be
	// selected by the implementation and this choice indicated to the API consumer in
	// the response or event. NOTE2: as for this Commonalities release, we are
	// enforcing that the networkAccessIdentifier is only part of the schema for
	// future-proofing, and CAMARA does not currently allow its use. After the CAMARA
	// meta-release work is concluded and the relevant issues are resolved, its use
	// will need to be explicitly documented in the guidelines.
	Device DeviceLocationDevice `json:"device"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Area        respjson.Field
		Device      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceLocationSubscriptionConfigSubscriptionDetail) RawJSON() string { return r.JSON.raw }
func (r *DeviceLocationSubscriptionConfigSubscriptionDetail) UnmarshalJSON(data []byte) error {
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
//   - `INACTIVE`: Subscription is temporarily inactive, but its workflow logic is
//     not deleted.
//   - `EXPIRED`: Subscription is ended (no longer active). This status applies when
//     subscription is ended due to `SUBSCRIPTION_EXPIRED` or `ACCESS_TOKEN_EXPIRED`
//     event.
//   - `DELETED`: Subscription is ended as deleted (no longer active). This status
//     applies when subscription information is kept (i.e. subscription workflow is
//     no longer active but its meta-information is kept).
type DeviceLocationSubscriptionStatus string

const (
	DeviceLocationSubscriptionStatusActivationRequested DeviceLocationSubscriptionStatus = "ACTIVATION_REQUESTED"
	DeviceLocationSubscriptionStatusActive              DeviceLocationSubscriptionStatus = "ACTIVE"
	DeviceLocationSubscriptionStatusExpired             DeviceLocationSubscriptionStatus = "EXPIRED"
	DeviceLocationSubscriptionStatusInactive            DeviceLocationSubscriptionStatus = "INACTIVE"
	DeviceLocationSubscriptionStatusDeleted             DeviceLocationSubscriptionStatus = "DELETED"
)

// area-entered - Event triggered when the device enters the given area
//
// area-left - Event triggered when the device leaves the given area
type DeviceLocationSubscriptionEventType string

const (
	DeviceLocationSubscriptionEventTypeOrgCamaraprojectGeofencingSubscriptionsV0AreaEntered DeviceLocationSubscriptionEventType = "org.camaraproject.geofencing-subscriptions.v0.area-entered"
	DeviceLocationSubscriptionEventTypeOrgCamaraprojectGeofencingSubscriptionsV0AreaLeft    DeviceLocationSubscriptionEventType = "org.camaraproject.geofencing-subscriptions.v0.area-left"
)

// Response for an event-type subscription request managed asynchronously (Creation
// or Deletion).
type DevicelocationSubscriptionDeleteResponse struct {
	// The unique identifier of the subscription in the scope of the subscription
	// manager. When this information is contained within an event notification, this
	// concept SHALL be referred as subscriptionId as per Commonalities Event
	// Notification Model.
	ID string `json:"id" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DevicelocationSubscriptionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DevicelocationSubscriptionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DevicelocationSubscriptionNewParams struct {
	// Implementation-specific configuration parameters are needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`.
	Config DevicelocationSubscriptionNewParamsConfig `json:"config,omitzero" api:"required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now.
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol DeviceLocationProtocol `json:"protocol,omitzero" api:"required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink" api:"required" format:"uri"`
	// Camara Event types which are eligible to be delivered by this subscription.
	// Note: As of now we enforce to have only event type per subscription.
	Types       []DeviceLocationSubscriptionEventType `json:"types,omitzero" api:"required"`
	XCorrelator param.Opt[string]                     `header:"x-correlator,omitzero" json:"-"`
	// A sink credential provides authentication or authorization information necessary
	// to enable delivery of events to a target.
	SinkCredential DevicelocationSubscriptionNewParamsSinkCredential `json:"sinkCredential,omitzero"`
	paramObj
}

func (r DevicelocationSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DevicelocationSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DevicelocationSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Implementation-specific configuration parameters are needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent`.
type DevicelocationSubscriptionNewParamsConfig struct {
	// The detail of the requested event subscription.
	SubscriptionDetail DevicelocationSubscriptionNewParamsConfigSubscriptionDetail `json:"subscriptionDetail,omitzero" api:"required"`
	DeviceLocationConfigParam
}

func (r DevicelocationSubscriptionNewParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow struct {
		*DevicelocationSubscriptionNewParamsConfig
		MarshalJSON bool `json:"-"` // Prevent inheriting [json.Marshaler] from the embedded field
	}
	return param.MarshalObject(r, shadow{&r, false})
}

// The detail of the requested event subscription.
//
// The property Area is required.
type DevicelocationSubscriptionNewParamsConfigSubscriptionDetail struct {
	// The geofencing area where the monitor is active. This area is specified by API
	// consumers in the subscription request. The same area definition is included in
	// event notifications without any modifications.
	Area DeviceLocationAreaParam `json:"area,omitzero" api:"required"`
	// End-user device able to connect to a mobile network. Examples of devices include
	// smartphones or IoT sensors/actuators.
	//
	// The developer can choose to provide the below specified device identifiers:
	//
	// - `ipv4Address`
	// - `ipv6Address`
	// - `phoneNumber`
	// - `networkAccessIdentifier`
	//
	// NOTE1: the API provider might support only a subset of these options. The API
	// consumer can provide multiple identifiers to be compatible across different API
	// providers. In this case the identifiers MUST belong to the same device. Where
	// more than one device identifier is provided, only one identifier will be
	// selected by the implementation and this choice indicated to the API consumer in
	// the response or event. NOTE2: as for this Commonalities release, we are
	// enforcing that the networkAccessIdentifier is only part of the schema for
	// future-proofing, and CAMARA does not currently allow its use. After the CAMARA
	// meta-release work is concluded and the relevant issues are resolved, its use
	// will need to be explicitly documented in the guidelines.
	Device DeviceLocationDeviceParam `json:"device,omitzero"`
	paramObj
}

func (r DevicelocationSubscriptionNewParamsConfigSubscriptionDetail) MarshalJSON() (data []byte, err error) {
	type shadow DevicelocationSubscriptionNewParamsConfigSubscriptionDetail
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DevicelocationSubscriptionNewParamsConfigSubscriptionDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sink credential provides authentication or authorization information necessary
// to enable delivery of events to a target.
//
// The property CredentialType is required.
type DevicelocationSubscriptionNewParamsSinkCredential struct {
	// The type of the credential. Note: Type of the credential - MUST be set to
	// ACCESSTOKEN for now
	//
	// Any of "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN".
	CredentialType string `json:"credentialType,omitzero" api:"required"`
	paramObj
}

func (r DevicelocationSubscriptionNewParamsSinkCredential) MarshalJSON() (data []byte, err error) {
	type shadow DevicelocationSubscriptionNewParamsSinkCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DevicelocationSubscriptionNewParamsSinkCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DevicelocationSubscriptionNewParamsSinkCredential](
		"credentialType", "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN",
	)
}

type DevicelocationSubscriptionGetParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type DevicelocationSubscriptionListParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type DevicelocationSubscriptionDeleteParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}
