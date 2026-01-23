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

// ConnectednetworktypeSubscriptionService contains methods and other services that
// help with interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConnectednetworktypeSubscriptionService] method instead.
type ConnectednetworktypeSubscriptionService struct {
	Options []option.RequestOption
}

// NewConnectednetworktypeSubscriptionService generates a new service that applies
// the given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewConnectednetworktypeSubscriptionService(opts ...option.RequestOption) (r ConnectednetworktypeSubscriptionService) {
	r = ConnectednetworktypeSubscriptionService{}
	r.Options = opts
	return
}

// Create a subscription for receiving notifications on changes to the connected
// network type of a device.
func (r *ConnectednetworktypeSubscriptionService) New(ctx context.Context, params ConnectednetworktypeSubscriptionNewParams, opts ...option.RequestOption) (res *ConnectedNetworkTypeSubscription, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectednetworktype/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// retrieve ConnectedNetworkType subscription information for a given subscription
// ID.
func (r *ConnectednetworktypeSubscriptionService) Get(ctx context.Context, subscriptionID string, query ConnectednetworktypeSubscriptionGetParams, opts ...option.RequestOption) (res *ConnectedNetworkTypeSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("connectednetworktype/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Retrieve a list of device connected network type event subscription(s)
func (r *ConnectednetworktypeSubscriptionService) List(ctx context.Context, query ConnectednetworktypeSubscriptionListParams, opts ...option.RequestOption) (res *[]ConnectedNetworkTypeSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "connectednetworktype/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// delete a given ConnectedNetworkType subscription.
func (r *ConnectednetworktypeSubscriptionService) Delete(ctx context.Context, subscriptionID string, body ConnectednetworktypeSubscriptionDeleteParams, opts ...option.RequestOption) (res *ConnectednetworktypeSubscriptionDeleteResponse, err error) {
	if !param.IsOmitted(body.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", body.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return
	}
	path := fmt.Sprintf("connectednetworktype/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Implementation-specific configuration parameters needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
// type attributes must be defined in `subscriptionDetail` Note: if a request is
// performed for several event type, all subscribed event will use same `config`
// parameters.
type ConnectedNetworkTypeConfig struct {
	// The detail of the requested event subscription.
	SubscriptionDetail ConnectedNetworkTypeConfigSubscriptionDetail `json:"subscriptionDetail,required"`
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request. Example:
	// Consumer request area entered event. If consumer sets initialEvent to true and
	// device is already in the geofence, an event is triggered
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
func (r ConnectedNetworkTypeConfig) RawJSON() string { return r.JSON.raw }
func (r *ConnectedNetworkTypeConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ConnectedNetworkTypeConfig to a
// ConnectedNetworkTypeConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ConnectedNetworkTypeConfigParam.Overrides()
func (r ConnectedNetworkTypeConfig) ToParam() ConnectedNetworkTypeConfigParam {
	return param.Override[ConnectedNetworkTypeConfigParam](json.RawMessage(r.RawJSON()))
}

// The detail of the requested event subscription.
type ConnectedNetworkTypeConfigSubscriptionDetail struct {
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
	Device ConnectedNetworkTypeConfigSubscriptionDetailDevice `json:"device"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Device      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectedNetworkTypeConfigSubscriptionDetail) RawJSON() string { return r.JSON.raw }
func (r *ConnectedNetworkTypeConfigSubscriptionDetail) UnmarshalJSON(data []byte) error {
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
type ConnectedNetworkTypeConfigSubscriptionDetailDevice struct {
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
	Ipv4Address ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4Address `json:"ipv4Address"`
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
func (r ConnectedNetworkTypeConfigSubscriptionDetailDevice) RawJSON() string { return r.JSON.raw }
func (r *ConnectedNetworkTypeConfigSubscriptionDetailDevice) UnmarshalJSON(data []byte) error {
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
type ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4Address struct {
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
func (r ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4Address) RawJSON() string {
	return r.JSON.raw
}
func (r *ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4Address) UnmarshalJSON(data []byte) error {
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
type ConnectedNetworkTypeConfigParam struct {
	// The detail of the requested event subscription.
	SubscriptionDetail ConnectedNetworkTypeConfigSubscriptionDetailParam `json:"subscriptionDetail,omitzero,required"`
	// Set to `true` by API consumer if consumer wants to get an event as soon as the
	// subscription is created and current situation reflects event request. Example:
	// Consumer request area entered event. If consumer sets initialEvent to true and
	// device is already in the geofence, an event is triggered
	InitialEvent param.Opt[bool] `json:"initialEvent,omitzero"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer.
	SubscriptionExpireTime param.Opt[time.Time] `json:"subscriptionExpireTime,omitzero" format:"date-time"`
	// Identifies the maximum number of event reports to be generated (>=1) requested
	// by the API consumer - Once this number is reached, the subscription ends.
	SubscriptionMaxEvents param.Opt[int64] `json:"subscriptionMaxEvents,omitzero"`
	paramObj
}

func (r ConnectedNetworkTypeConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectedNetworkTypeConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectedNetworkTypeConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The detail of the requested event subscription.
type ConnectedNetworkTypeConfigSubscriptionDetailParam struct {
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
	Device ConnectedNetworkTypeConfigSubscriptionDetailDeviceParam `json:"device,omitzero"`
	paramObj
}

func (r ConnectedNetworkTypeConfigSubscriptionDetailParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectedNetworkTypeConfigSubscriptionDetailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectedNetworkTypeConfigSubscriptionDetailParam) UnmarshalJSON(data []byte) error {
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
type ConnectedNetworkTypeConfigSubscriptionDetailDeviceParam struct {
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
	Ipv4Address ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4AddressParam `json:"ipv4Address,omitzero"`
	paramObj
}

func (r ConnectedNetworkTypeConfigSubscriptionDetailDeviceParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectedNetworkTypeConfigSubscriptionDetailDeviceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectedNetworkTypeConfigSubscriptionDetailDeviceParam) UnmarshalJSON(data []byte) error {
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
type ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4AddressParam struct {
	// A single IPv4 address with no subnet mask
	PrivateAddress param.Opt[string] `json:"privateAddress,omitzero" format:"ipv4"`
	// A single IPv4 address with no subnet mask
	PublicAddress param.Opt[string] `json:"publicAddress,omitzero" format:"ipv4"`
	// TCP or UDP port number
	PublicPort param.Opt[int64] `json:"publicPort,omitzero"`
	paramObj
}

func (r ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4AddressParam) MarshalJSON() (data []byte, err error) {
	type shadow ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4AddressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectedNetworkTypeConfigSubscriptionDetailDeviceIpv4AddressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of a delivery protocol. Only HTTP is allowed for now
type ConnectedNetworkTypeProtocol string

const (
	ConnectedNetworkTypeProtocolHTTP  ConnectedNetworkTypeProtocol = "HTTP"
	ConnectedNetworkTypeProtocolMqtt3 ConnectedNetworkTypeProtocol = "MQTT3"
	ConnectedNetworkTypeProtocolMqtt5 ConnectedNetworkTypeProtocol = "MQTT5"
	ConnectedNetworkTypeProtocolAmqp  ConnectedNetworkTypeProtocol = "AMQP"
	ConnectedNetworkTypeProtocolNats  ConnectedNetworkTypeProtocol = "NATS"
	ConnectedNetworkTypeProtocolKafka ConnectedNetworkTypeProtocol = "KAFKA"
)

// Represents a event-type subscription.
type ConnectedNetworkTypeSubscription struct {
	// The unique identifier of the subscription in the scope of the subscription
	// manager. When this information is contained within an event notification, this
	// concept SHALL be referred as subscriptionId as per Commonalities Event
	// Notification Model.
	ID string `json:"id,required"`
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
	// type attributes must be defined in `subscriptionDetail` Note: if a request is
	// performed for several event type, all subscribed event will use same `config`
	// parameters.
	Config ConnectedNetworkTypeConfig `json:"config,required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol ConnectedNetworkTypeProtocol `json:"protocol,required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink,required" format:"uri"`
	// Camara Event types eligible to be delivered by this subscription. Note: For the
	// current Commonalities API design guidelines, only one event type per
	// subscription is allowed
	Types []ConnectedNetworkTypeSubscriptionEventType `json:"types,required"`
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
	Status ConnectedNetworkTypeSubscriptionStatus `json:"status"`
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
func (r ConnectedNetworkTypeSubscription) RawJSON() string { return r.JSON.raw }
func (r *ConnectedNetworkTypeSubscription) UnmarshalJSON(data []byte) error {
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
type ConnectedNetworkTypeSubscriptionStatus string

const (
	ConnectedNetworkTypeSubscriptionStatusActivationRequested ConnectedNetworkTypeSubscriptionStatus = "ACTIVATION_REQUESTED"
	ConnectedNetworkTypeSubscriptionStatusActive              ConnectedNetworkTypeSubscriptionStatus = "ACTIVE"
	ConnectedNetworkTypeSubscriptionStatusExpired             ConnectedNetworkTypeSubscriptionStatus = "EXPIRED"
	ConnectedNetworkTypeSubscriptionStatusInactive            ConnectedNetworkTypeSubscriptionStatus = "INACTIVE"
	ConnectedNetworkTypeSubscriptionStatusDeleted             ConnectedNetworkTypeSubscriptionStatus = "DELETED"
)

// network-type-changed - Event triggered when the connected network type of the
// device changes.
type ConnectedNetworkTypeSubscriptionEventType string

const (
	ConnectedNetworkTypeSubscriptionEventTypeOrgCamaraprojectConnectedNetworkTypeSubscriptionsV0NetworkTypeChanged ConnectedNetworkTypeSubscriptionEventType = "org.camaraproject.connected-network-type-subscriptions.v0.network-type-changed"
)

// Response for a event-type subscription request managed asynchronously (Creation
// or Deletion)
type ConnectednetworktypeSubscriptionDeleteResponse struct {
	// The unique identifier of the subscription in the scope of the subscription
	// manager. When this information is contained within an event notification, this
	// concept SHALL be referred as subscriptionId as per Commonalities Event
	// Notification Model.
	ID string `json:"id,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectednetworktypeSubscriptionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ConnectednetworktypeSubscriptionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectednetworktypeSubscriptionNewParams struct {
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime`, `subscriptionMaxEvents`, `initialEvent` Specific event
	// type attributes must be defined in `subscriptionDetail` Note: if a request is
	// performed for several event type, all subscribed event will use same `config`
	// parameters.
	Config ConnectedNetworkTypeConfigParam `json:"config,omitzero,required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol ConnectedNetworkTypeProtocol `json:"protocol,omitzero,required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink,required" format:"uri"`
	// Camara Event types eligible to be delivered by this subscription. Note: As of
	// now we enforce to have only event type per subscription.
	Types       []ConnectedNetworkTypeSubscriptionEventType `json:"types,omitzero,required"`
	XCorrelator param.Opt[string]                           `header:"x-correlator,omitzero" json:"-"`
	// A sink credential provides authentication or authorization information necessary
	// to enable delivery of events to a target.
	SinkCredential ConnectednetworktypeSubscriptionNewParamsSinkCredential `json:"sinkCredential,omitzero"`
	paramObj
}

func (r ConnectednetworktypeSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ConnectednetworktypeSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectednetworktypeSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sink credential provides authentication or authorization information necessary
// to enable delivery of events to a target.
//
// The property CredentialType is required.
type ConnectednetworktypeSubscriptionNewParamsSinkCredential struct {
	// The type of the credential. Note: Type of the credential - MUST be set to
	// ACCESSTOKEN for now
	//
	// Any of "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN".
	CredentialType string `json:"credentialType,omitzero,required"`
	paramObj
}

func (r ConnectednetworktypeSubscriptionNewParamsSinkCredential) MarshalJSON() (data []byte, err error) {
	type shadow ConnectednetworktypeSubscriptionNewParamsSinkCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ConnectednetworktypeSubscriptionNewParamsSinkCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ConnectednetworktypeSubscriptionNewParamsSinkCredential](
		"credentialType", "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN",
	)
}

type ConnectednetworktypeSubscriptionGetParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type ConnectednetworktypeSubscriptionListParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type ConnectednetworktypeSubscriptionDeleteParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}
