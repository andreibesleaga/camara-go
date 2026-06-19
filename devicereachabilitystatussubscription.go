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

// Device Reachability Status Subscriptions
//
// DevicereachabilitystatusSubscriptionService contains methods and other services
// that help with interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDevicereachabilitystatusSubscriptionService] method instead.
type DevicereachabilitystatusSubscriptionService struct {
	Options []option.RequestOption
}

// NewDevicereachabilitystatusSubscriptionService generates a new service that
// applies the given options to each request. These options are applied after the
// parent client's options (if there is one), and before any request-specific
// options.
func NewDevicereachabilitystatusSubscriptionService(opts ...option.RequestOption) (r DevicereachabilitystatusSubscriptionService) {
	r = DevicereachabilitystatusSubscriptionService{}
	r.Options = opts
	return
}

// Create a device reachability status event subscription for a device
func (r *DevicereachabilitystatusSubscriptionService) New(ctx context.Context, params DevicereachabilitystatusSubscriptionNewParams, opts ...option.RequestOption) (res *DeviceReachabilityStatusSubscription, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "devicereachabilitystatus/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Retrieve a given subscription by ID
func (r *DevicereachabilitystatusSubscriptionService) Get(ctx context.Context, subscriptionID string, query DevicereachabilitystatusSubscriptionGetParams, opts ...option.RequestOption) (res *DeviceReachabilityStatusSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devicereachabilitystatus/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a list of device reachability status event subscription(s)
func (r *DevicereachabilitystatusSubscriptionService) List(ctx context.Context, query DevicereachabilitystatusSubscriptionListParams, opts ...option.RequestOption) (res *[]DeviceReachabilityStatusSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "devicereachabilitystatus/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a given subscription by ID
func (r *DevicereachabilitystatusSubscriptionService) Delete(ctx context.Context, subscriptionID string, body DevicereachabilitystatusSubscriptionDeleteParams, opts ...option.RequestOption) (res *DevicereachabilitystatusSubscriptionDeleteResponse, err error) {
	if !param.IsOmitted(body.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", body.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("devicereachabilitystatus/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Implementation-specific configuration parameters needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
// type attributes must be defined in `subscriptionDetail` Note: if a request is
// performed for several event type, all subscribed event will use same `config`
// parameters.
type DeviceReachabilityStatusConfig struct {
	// The detail of the requested event subscription.
	SubscriptionDetail DeviceReachabilityStatusConfigSubscriptionDetail `json:"subscriptionDetail" api:"required"`
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request. Example:
	// Consumer subscribes to reachability SMS. If consumer sets initialEvent to true
	// and device is already reachable by SMS, an event is triggered.
	InitialEvent bool `json:"initialEvent"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer.
	SubscriptionExpireTime time.Time `json:"subscriptionExpireTime" format:"date-time"`
	// Identifies the maximum number of event reports to be generated (>=1) requested
	// by the API consumer - Once this number is reached, the subscription ends.
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
func (r DeviceReachabilityStatusConfig) RawJSON() string { return r.JSON.raw }
func (r *DeviceReachabilityStatusConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DeviceReachabilityStatusConfig to a
// DeviceReachabilityStatusConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DeviceReachabilityStatusConfigParam.Overrides()
func (r DeviceReachabilityStatusConfig) ToParam() DeviceReachabilityStatusConfigParam {
	return param.Override[DeviceReachabilityStatusConfigParam](json.RawMessage(r.RawJSON()))
}

// The detail of the requested event subscription.
type DeviceReachabilityStatusConfigSubscriptionDetail struct {
	// End-user equipment able to connect to a mobile network. Examples of devices
	// include smartphones or IoT sensors/actuators.
	//
	// The developer can choose to provide the below specified device identifiers:
	//
	// - `ipv4Address`
	// - `ipv6Address`
	// - `phoneNumber`
	// - `networkAccessIdentifier`
	//
	// NOTE: the MNO might support only a subset of these options. The API invoker can
	// provide multiple identifiers to be compatible across different MNOs. In this
	// case the identifiers MUST belong to the same device.
	Device DeviceReachabilityStatusConfigSubscriptionDetailDevice `json:"device"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Device      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceReachabilityStatusConfigSubscriptionDetail) RawJSON() string { return r.JSON.raw }
func (r *DeviceReachabilityStatusConfigSubscriptionDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user equipment able to connect to a mobile network. Examples of devices
// include smartphones or IoT sensors/actuators.
//
// The developer can choose to provide the below specified device identifiers:
//
// - `ipv4Address`
// - `ipv6Address`
// - `phoneNumber`
// - `networkAccessIdentifier`
//
// NOTE: the MNO might support only a subset of these options. The API invoker can
// provide multiple identifiers to be compatible across different MNOs. In this
// case the identifiers MUST belong to the same device.
type DeviceReachabilityStatusConfigSubscriptionDetailDevice struct {
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
	Ipv4Address DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4Address `json:"ipv4Address"`
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
func (r DeviceReachabilityStatusConfigSubscriptionDetailDevice) RawJSON() string { return r.JSON.raw }
func (r *DeviceReachabilityStatusConfigSubscriptionDetailDevice) UnmarshalJSON(data []byte) error {
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
type DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4Address struct {
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
func (r DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4Address) RawJSON() string {
	return r.JSON.raw
}
func (r *DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4Address) UnmarshalJSON(data []byte) error {
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
type DeviceReachabilityStatusConfigParam struct {
	// The detail of the requested event subscription.
	SubscriptionDetail DeviceReachabilityStatusConfigSubscriptionDetailParam `json:"subscriptionDetail,omitzero" api:"required"`
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request. Example:
	// Consumer subscribes to reachability SMS. If consumer sets initialEvent to true
	// and device is already reachable by SMS, an event is triggered.
	InitialEvent param.Opt[bool] `json:"initialEvent,omitzero"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer.
	SubscriptionExpireTime param.Opt[time.Time] `json:"subscriptionExpireTime,omitzero" format:"date-time"`
	// Identifies the maximum number of event reports to be generated (>=1) requested
	// by the API consumer - Once this number is reached, the subscription ends.
	SubscriptionMaxEvents param.Opt[int64] `json:"subscriptionMaxEvents,omitzero"`
	paramObj
}

func (r DeviceReachabilityStatusConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceReachabilityStatusConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceReachabilityStatusConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The detail of the requested event subscription.
type DeviceReachabilityStatusConfigSubscriptionDetailParam struct {
	// End-user equipment able to connect to a mobile network. Examples of devices
	// include smartphones or IoT sensors/actuators.
	//
	// The developer can choose to provide the below specified device identifiers:
	//
	// - `ipv4Address`
	// - `ipv6Address`
	// - `phoneNumber`
	// - `networkAccessIdentifier`
	//
	// NOTE: the MNO might support only a subset of these options. The API invoker can
	// provide multiple identifiers to be compatible across different MNOs. In this
	// case the identifiers MUST belong to the same device.
	Device DeviceReachabilityStatusConfigSubscriptionDetailDeviceParam `json:"device,omitzero"`
	paramObj
}

func (r DeviceReachabilityStatusConfigSubscriptionDetailParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceReachabilityStatusConfigSubscriptionDetailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceReachabilityStatusConfigSubscriptionDetailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user equipment able to connect to a mobile network. Examples of devices
// include smartphones or IoT sensors/actuators.
//
// The developer can choose to provide the below specified device identifiers:
//
// - `ipv4Address`
// - `ipv6Address`
// - `phoneNumber`
// - `networkAccessIdentifier`
//
// NOTE: the MNO might support only a subset of these options. The API invoker can
// provide multiple identifiers to be compatible across different MNOs. In this
// case the identifiers MUST belong to the same device.
type DeviceReachabilityStatusConfigSubscriptionDetailDeviceParam struct {
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
	Ipv4Address DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4AddressParam `json:"ipv4Address,omitzero"`
	paramObj
}

func (r DeviceReachabilityStatusConfigSubscriptionDetailDeviceParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceReachabilityStatusConfigSubscriptionDetailDeviceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceReachabilityStatusConfigSubscriptionDetailDeviceParam) UnmarshalJSON(data []byte) error {
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
type DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4AddressParam struct {
	// A single IPv4 address with no subnet mask
	PrivateAddress param.Opt[string] `json:"privateAddress,omitzero" format:"ipv4"`
	// A single IPv4 address with no subnet mask
	PublicAddress param.Opt[string] `json:"publicAddress,omitzero" format:"ipv4"`
	// TCP or UDP port number
	PublicPort param.Opt[int64] `json:"publicPort,omitzero"`
	paramObj
}

func (r DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DeviceReachabilityStatusConfigSubscriptionDetailDeviceIpv4AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of a delivery protocol. Only HTTP is allowed for now
type DeviceReachabilityStatusProtocol string

const (
	DeviceReachabilityStatusProtocolHTTP  DeviceReachabilityStatusProtocol = "HTTP"
	DeviceReachabilityStatusProtocolMqtt3 DeviceReachabilityStatusProtocol = "MQTT3"
	DeviceReachabilityStatusProtocolMqtt5 DeviceReachabilityStatusProtocol = "MQTT5"
	DeviceReachabilityStatusProtocolAmqp  DeviceReachabilityStatusProtocol = "AMQP"
	DeviceReachabilityStatusProtocolNats  DeviceReachabilityStatusProtocol = "NATS"
	DeviceReachabilityStatusProtocolKafka DeviceReachabilityStatusProtocol = "KAFKA"
)

// Represents a event-type subscription.
type DeviceReachabilityStatusSubscription struct {
	// The unique identifier of the subscription in the scope of the subscription
	// manager. When this information is contained within an event notification, this
	// concept SHALL be referred as subscriptionId as per Commonalities Event
	// Notification Model.
	ID string `json:"id" api:"required"`
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
	// type attributes must be defined in `subscriptionDetail` Note: if a request is
	// performed for several event type, all subscribed event will use same `config`
	// parameters.
	Config DeviceReachabilityStatusConfig `json:"config" api:"required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol DeviceReachabilityStatusProtocol `json:"protocol" api:"required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink" api:"required" format:"uri"`
	// Camara Event types eligible to be delivered by this subscription. Note: For the
	// current Commonalities API design guidelines, only one event type per
	// subscription is allowed
	Types []DeviceReachabilityStatusSubscriptionEventType `json:"types" api:"required"`
	// Date when the event subscription will expire. Only provided when
	// `subscriptionExpireTime` is indicated by API client or Telco Operator has
	// specific policy about that. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ (i.e. which
	// allows 2023-07-03T14:27:08.312+02:00 or 2023-07-03T12:27:08.312Z)
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// Date when the event subscription will begin/began It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone. Recommended format is yyyy-MM-dd'T'HH:mm:ss.SSSZ (i.e. which
	// allows 2023-07-03T14:27:08.312+02:00 or 2023-07-03T12:27:08.312Z)
	StartsAt time.Time `json:"startsAt" format:"date-time"`
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
	Status DeviceReachabilityStatusSubscriptionStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Config      respjson.Field
		Protocol    respjson.Field
		Sink        respjson.Field
		Types       respjson.Field
		ExpiresAt   respjson.Field
		StartsAt    respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DeviceReachabilityStatusSubscription) RawJSON() string { return r.JSON.raw }
func (r *DeviceReachabilityStatusSubscription) UnmarshalJSON(data []byte) error {
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
type DeviceReachabilityStatusSubscriptionStatus string

const (
	DeviceReachabilityStatusSubscriptionStatusActivationRequested DeviceReachabilityStatusSubscriptionStatus = "ACTIVATION_REQUESTED"
	DeviceReachabilityStatusSubscriptionStatusActive              DeviceReachabilityStatusSubscriptionStatus = "ACTIVE"
	DeviceReachabilityStatusSubscriptionStatusExpired             DeviceReachabilityStatusSubscriptionStatus = "EXPIRED"
	DeviceReachabilityStatusSubscriptionStatusInactive            DeviceReachabilityStatusSubscriptionStatus = "INACTIVE"
	DeviceReachabilityStatusSubscriptionStatusDeleted             DeviceReachabilityStatusSubscriptionStatus = "DELETED"
)

// reachability-data - Event triggered when the device is connected to the network
// for Data usage (regardless of the SMS reachability).
//
// reachability-sms - Event triggered when the device is connected to the network
// only for SMS usage
//
// reachability-disconnected - Event triggered when the device is not connected.
type DeviceReachabilityStatusSubscriptionEventType string

const (
	DeviceReachabilityStatusSubscriptionEventTypeOrgCamaraprojectDeviceReachabilityStatusSubscriptionsV0ReachabilityData         DeviceReachabilityStatusSubscriptionEventType = "org.camaraproject.device-reachability-status-subscriptions.v0.reachability-data"
	DeviceReachabilityStatusSubscriptionEventTypeOrgCamaraprojectDeviceReachabilityStatusSubscriptionsV0ReachabilitySMS          DeviceReachabilityStatusSubscriptionEventType = "org.camaraproject.device-reachability-status-subscriptions.v0.reachability-sms"
	DeviceReachabilityStatusSubscriptionEventTypeOrgCamaraprojectDeviceReachabilityStatusSubscriptionsV0ReachabilityDisconnected DeviceReachabilityStatusSubscriptionEventType = "org.camaraproject.device-reachability-status-subscriptions.v0.reachability-disconnected"
)

// Response for a device reachability status operation managed asynchronously
// (Creation or Deletion)
type DevicereachabilitystatusSubscriptionDeleteResponse struct {
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
func (r DevicereachabilitystatusSubscriptionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DevicereachabilitystatusSubscriptionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DevicereachabilitystatusSubscriptionNewParams struct {
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
	// type attributes must be defined in `subscriptionDetail` Note: if a request is
	// performed for several event type, all subscribed event will use same `config`
	// parameters.
	Config DeviceReachabilityStatusConfigParam `json:"config,omitzero" api:"required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol DeviceReachabilityStatusProtocol `json:"protocol,omitzero" api:"required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink" api:"required" format:"uri"`
	// Camara Event types eligible to be delivered by this subscription. Note: For the
	// current Commonalities API design guidelines, only one event type per
	// subscription is allowed, yet in the following releases use of array of event
	// types SHALL be specified without changing this definition.
	Types       []DeviceReachabilityStatusSubscriptionEventType `json:"types,omitzero" api:"required"`
	XCorrelator param.Opt[string]                               `header:"x-correlator,omitzero" json:"-"`
	// A sink credential provides authentication or authorization information necessary
	// to enable delivery of events to a target.
	SinkCredential DevicereachabilitystatusSubscriptionNewParamsSinkCredential `json:"sinkCredential,omitzero"`
	paramObj
}

func (r DevicereachabilitystatusSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DevicereachabilitystatusSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DevicereachabilitystatusSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sink credential provides authentication or authorization information necessary
// to enable delivery of events to a target.
//
// The property CredentialType is required.
type DevicereachabilitystatusSubscriptionNewParamsSinkCredential struct {
	// The type of the credential. Note: Type of the credential - MUST be set to
	// ACCESSTOKEN for now
	//
	// Any of "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN".
	CredentialType string `json:"credentialType,omitzero" api:"required"`
	paramObj
}

func (r DevicereachabilitystatusSubscriptionNewParamsSinkCredential) MarshalJSON() (data []byte, err error) {
	type shadow DevicereachabilitystatusSubscriptionNewParamsSinkCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DevicereachabilitystatusSubscriptionNewParamsSinkCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DevicereachabilitystatusSubscriptionNewParamsSinkCredential](
		"credentialType", "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN",
	)
}

type DevicereachabilitystatusSubscriptionGetParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type DevicereachabilitystatusSubscriptionListParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type DevicereachabilitystatusSubscriptionDeleteParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}
