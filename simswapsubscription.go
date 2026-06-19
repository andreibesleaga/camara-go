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

// Sim Swap Subscriptions
//
// SimswapSubscriptionService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSimswapSubscriptionService] method instead.
type SimswapSubscriptionService struct {
	Options []option.RequestOption
}

// NewSimswapSubscriptionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSimswapSubscriptionService(opts ...option.RequestOption) (r SimswapSubscriptionService) {
	r = SimswapSubscriptionService{}
	r.Options = opts
	return
}

// Create a sim swap event subscription for a phone number
func (r *SimswapSubscriptionService) New(ctx context.Context, params SimswapSubscriptionNewParams, opts ...option.RequestOption) (res *SimSwapSubscription, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "simswap/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// retrieve event subscription information for a given subscription.
func (r *SimswapSubscriptionService) Get(ctx context.Context, subscriptionID string, query SimswapSubscriptionGetParams, opts ...option.RequestOption) (res *SimSwapSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("simswap/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Retrieve a list of sim swap event subscription(s)
func (r *SimswapSubscriptionService) List(ctx context.Context, query SimswapSubscriptionListParams, opts ...option.RequestOption) (res *[]SimSwapSubscription, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "simswap/subscriptions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// delete a given event subscription.
func (r *SimswapSubscriptionService) Delete(ctx context.Context, subscriptionID string, body SimswapSubscriptionDeleteParams, opts ...option.RequestOption) (res *SimswapSubscriptionDeleteResponse, err error) {
	if !param.IsOmitted(body.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", body.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if subscriptionID == "" {
		err = errors.New("missing required subscriptionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("simswap/subscriptions/%s", subscriptionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Implementation-specific configuration parameters needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime` or `subscriptionMaxEvents` to limit subscription
// lifetime. Event type attributes must be defined in `subscriptionDetail`
type SimSwapConfig struct {
	// The detail of the requested event subscription
	SubscriptionDetail SimSwapConfigSubscriptionDetail `json:"subscriptionDetail" api:"required"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	SubscriptionExpireTime time.Time `json:"subscriptionExpireTime" format:"date-time"`
	// Identifies the maximum number of event reports to be generated (>=1) requested
	// by the API consumer - Once this number is reached, the subscription ends.
	SubscriptionMaxEvents int64 `json:"subscriptionMaxEvents"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SubscriptionDetail     respjson.Field
		SubscriptionExpireTime respjson.Field
		SubscriptionMaxEvents  respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimSwapConfig) RawJSON() string { return r.JSON.raw }
func (r *SimSwapConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SimSwapConfig to a SimSwapConfigParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SimSwapConfigParam.Overrides()
func (r SimSwapConfig) ToParam() SimSwapConfigParam {
	return param.Override[SimSwapConfigParam](json.RawMessage(r.RawJSON()))
}

// The detail of the requested event subscription
type SimSwapConfigSubscriptionDetail struct {
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber string `json:"phoneNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimSwapConfigSubscriptionDetail) RawJSON() string { return r.JSON.raw }
func (r *SimSwapConfigSubscriptionDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Implementation-specific configuration parameters needed by the subscription
// manager for acquiring events. In CAMARA we have predefined attributes like
// `subscriptionExpireTime` or `subscriptionMaxEvents` to limit subscription
// lifetime. Event type attributes must be defined in `subscriptionDetail`
//
// The property SubscriptionDetail is required.
type SimSwapConfigParam struct {
	// The detail of the requested event subscription
	SubscriptionDetail SimSwapConfigSubscriptionDetailParam `json:"subscriptionDetail,omitzero" api:"required"`
	// The subscription expiration time (in date-time format) requested by the API
	// consumer. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	SubscriptionExpireTime param.Opt[time.Time] `json:"subscriptionExpireTime,omitzero" format:"date-time"`
	// Identifies the maximum number of event reports to be generated (>=1) requested
	// by the API consumer - Once this number is reached, the subscription ends.
	SubscriptionMaxEvents param.Opt[int64] `json:"subscriptionMaxEvents,omitzero"`
	paramObj
}

func (r SimSwapConfigParam) MarshalJSON() (data []byte, err error) {
	type shadow SimSwapConfigParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimSwapConfigParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The detail of the requested event subscription
type SimSwapConfigSubscriptionDetailParam struct {
	// A public identifier addressing a telephone subscription. In mobile networks it
	// corresponds to the MSISDN (Mobile Station International Subscriber Directory
	// Number). In order to be globally unique it has to be formatted in international
	// format, according to E.164 standard, prefixed with '+'.
	PhoneNumber param.Opt[string] `json:"phoneNumber,omitzero"`
	paramObj
}

func (r SimSwapConfigSubscriptionDetailParam) MarshalJSON() (data []byte, err error) {
	type shadow SimSwapConfigSubscriptionDetailParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimSwapConfigSubscriptionDetailParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Identifier of a delivery protocol. Only HTTP is allowed for now
type SimSwapProtocol string

const (
	SimSwapProtocolHTTP  SimSwapProtocol = "HTTP"
	SimSwapProtocolMqtt3 SimSwapProtocol = "MQTT3"
	SimSwapProtocolMqtt5 SimSwapProtocol = "MQTT5"
	SimSwapProtocolAmqp  SimSwapProtocol = "AMQP"
	SimSwapProtocolNats  SimSwapProtocol = "NATS"
	SimSwapProtocolKafka SimSwapProtocol = "KAFKA"
)

// Represents a event-type subscription.
type SimSwapSubscription struct {
	// The unique identifier of the subscription in the scope of the subscription
	// manager. When this information is contained within an event notification, this
	// concept SHALL be referred as subscriptionId as per Commonalities Event
	// Notification Model.
	ID string `json:"id" api:"required"`
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime` or `subscriptionMaxEvents` to limit subscription
	// lifetime. Event type attributes must be defined in `subscriptionDetail`
	Config SimSwapConfig `json:"config" api:"required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol SimSwapProtocol `json:"protocol" api:"required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink" api:"required" format:"uri"`
	// Camara Event types eligible for subscription:
	//
	//   - org.camaraproject.sim-swap-subscriptions.v0.swapped: receive a notification
	//     when a sim swap is performed on the line. Note: for the Commonalities
	//     meta-release v0.4 we enforce to have only event type per subscription then for
	//     following meta-release use of array MUST be decided at API project level.
	Types []SimSwapSubscriptionEventType `json:"types" api:"required"`
	// Date when the event subscription will expire. Only provided when
	// `subscriptionExpireTime` is indicated by API client or Telco Operator has
	// specific policy about that. It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
	ExpiresAt time.Time `json:"expiresAt" format:"date-time"`
	// Date when the event subscription will begin/began It must follow
	// [RFC 3339](https://datatracker.ietf.org/doc/html/rfc3339#section-5.6) and must
	// have time zone.
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
	//     subscription is ended due to `SUBSCRIPTION_EXPIRED` event.
	//   - `DELETED`: Subscription is ended as deleted (no longer active). This status
	//     applies when subscription information is kept (i.e. subscription workflow is
	//     no longer active but its metainformation is kept).
	//
	// Any of "ACTIVATION_REQUESTED", "ACTIVE", "EXPIRED", "INACTIVE", "DELETED".
	Status SimSwapSubscriptionStatus `json:"status"`
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
func (r SimSwapSubscription) RawJSON() string { return r.JSON.raw }
func (r *SimSwapSubscription) UnmarshalJSON(data []byte) error {
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
//     subscription is ended due to `SUBSCRIPTION_EXPIRED` event.
//   - `DELETED`: Subscription is ended as deleted (no longer active). This status
//     applies when subscription information is kept (i.e. subscription workflow is
//     no longer active but its metainformation is kept).
type SimSwapSubscriptionStatus string

const (
	SimSwapSubscriptionStatusActivationRequested SimSwapSubscriptionStatus = "ACTIVATION_REQUESTED"
	SimSwapSubscriptionStatusActive              SimSwapSubscriptionStatus = "ACTIVE"
	SimSwapSubscriptionStatusExpired             SimSwapSubscriptionStatus = "EXPIRED"
	SimSwapSubscriptionStatusInactive            SimSwapSubscriptionStatus = "INACTIVE"
	SimSwapSubscriptionStatusDeleted             SimSwapSubscriptionStatus = "DELETED"
)

// swapped - Event triggered when a sim swap occurs on the line
type SimSwapSubscriptionEventType string

const (
	SimSwapSubscriptionEventTypeOrgCamaraprojectSimSwapSubscriptionsV0Swapped SimSwapSubscriptionEventType = "org.camaraproject.sim-swap-subscriptions.v0.swapped"
)

// Response for a event-type subscription request managed asynchronously (Creation
// or Deletion)
type SimswapSubscriptionDeleteResponse struct {
	// The unique identifier of the subscription in the scope of the subscription
	// manager. When this information is contained within an event notification, this
	// concept SHALL be referred as subscriptionId as per Commonalities Event
	// Notification Model.
	ID string `json:"id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SimswapSubscriptionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SimswapSubscriptionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SimswapSubscriptionNewParams struct {
	// Implementation-specific configuration parameters needed by the subscription
	// manager for acquiring events. In CAMARA we have predefined attributes like
	// `subscriptionExpireTime` or `subscriptionMaxEvents` to limit subscription
	// lifetime. Event type attributes must be defined in `subscriptionDetail`
	Config SimSwapConfigParam `json:"config,omitzero" api:"required"`
	// Identifier of a delivery protocol. Only HTTP is allowed for now
	//
	// Any of "HTTP", "MQTT3", "MQTT5", "AMQP", "NATS", "KAFKA".
	Protocol SimSwapProtocol `json:"protocol,omitzero" api:"required"`
	// The address to which events shall be delivered using the selected protocol.
	Sink string `json:"sink" api:"required" format:"uri"`
	// Camara Event types eligible for subscription:
	//
	//   - org.camaraproject.sim-swap-subscriptions.v0.swapped: receive a notification
	//     when a sim swap is performed on the line.
	Types       []SimSwapSubscriptionEventType `json:"types,omitzero" api:"required"`
	XCorrelator param.Opt[string]              `header:"x-correlator,omitzero" json:"-"`
	// A sink credential provides authentication or authorization information necessary
	// to enable delivery of events to a target.
	SinkCredential SimswapSubscriptionNewParamsSinkCredential `json:"sinkCredential,omitzero"`
	paramObj
}

func (r SimswapSubscriptionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow SimswapSubscriptionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimswapSubscriptionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A sink credential provides authentication or authorization information necessary
// to enable delivery of events to a target.
//
// The property CredentialType is required.
type SimswapSubscriptionNewParamsSinkCredential struct {
	// The type of the credential. With the current API version the type MUST be set to
	// ACCESSTOKEN.
	//
	// Any of "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN".
	CredentialType string `json:"credentialType,omitzero" api:"required"`
	paramObj
}

func (r SimswapSubscriptionNewParamsSinkCredential) MarshalJSON() (data []byte, err error) {
	type shadow SimswapSubscriptionNewParamsSinkCredential
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SimswapSubscriptionNewParamsSinkCredential) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[SimswapSubscriptionNewParamsSinkCredential](
		"credentialType", "PLAIN", "ACCESSTOKEN", "REFRESHTOKEN",
	)
}

type SimswapSubscriptionGetParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type SimswapSubscriptionListParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type SimswapSubscriptionDeleteParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}
