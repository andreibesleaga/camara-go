// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/camara-go/internal/apijson"
	"github.com/stainless-sdks/camara-go/internal/requestconfig"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/stainless-sdks/camara-go/packages/param"
	"github.com/stainless-sdks/camara-go/packages/respjson"
)

// QualityondemandService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewQualityondemandService] method instead.
type QualityondemandService struct {
	Options []option.RequestOption
}

// NewQualityondemandService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewQualityondemandService(opts ...option.RequestOption) (r QualityondemandService) {
	r = QualityondemandService{}
	r.Options = opts
	return
}

// Returns a QoS Profile that matches the given name.
//
// The access token may be either a 2-legged or 3-legged access token. If the
// access token is 3-legged, a QoS Profile is only returned if available to all
// subjects associated with the access token.
func (r *QualityondemandService) GetQosProfile(ctx context.Context, name string, query QualityondemandGetQosProfileParams, opts ...option.RequestOption) (res *QosProfile, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if name == "" {
		err = errors.New("missing required name parameter")
		return
	}
	path := fmt.Sprintf("qualityondemand/qos-profiles/%s", name)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all QoS Profiles that match the given criteria. **NOTES:**
//
//   - The access token may be either a 2-legged or 3-legged access token.
//   - If the access token is 3-legged, all returned QoS Profiles will be available
//     to the subject (device) associated with the access token.
//   - If the access token is 2-legged and a device filter is provided, all returned
//     QoS Profiles will be available to that device. If multiple device identifiers
//     are provided within the device property, only QoS Profiles available to the
//     device identifier chosen by the implementation will be returned, even if the
//     identifiers do not match the same device. API provider does not perform any
//     logic to validate/correlate that the indicated device identifiers match the
//     same device. No error should be returned if the identifiers are otherwise
//     valid to prevent API consumers correlating different identifiers with a given
//     end user.
//   - This call uses the POST method instead of GET to comply with the CAMARA
//     Commonalities guidelines for sending sensitive or complex data in API calls.
//     Since the device field may contain personally identifiable information, it
//     should not be sent via GET. Additionally, this call may include complex data
//     structures.
//     [CAMARA API Design Guidelines](https://github.com/camaraproject/Commonalities/blob/r3.3/documentation/API-design-guidelines.md#post-or-get-for-transferring-sensitive-or-complex-data)
func (r *QualityondemandService) GetQosProfiles(ctx context.Context, params QualityondemandGetQosProfilesParams, opts ...option.RequestOption) (res *[]QosProfile, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "qualityondemand/retrieve-qos-profiles"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Specification of duration
type Duration struct {
	// Units of time
	//
	// Any of "Days", "Hours", "Minutes", "Seconds", "Milliseconds", "Microseconds",
	// "Nanoseconds".
	Unit DurationUnit `json:"unit"`
	// Quantity of duration
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Duration) RawJSON() string { return r.JSON.raw }
func (r *Duration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Units of time
type DurationUnit string

const (
	DurationUnitDays         DurationUnit = "Days"
	DurationUnitHours        DurationUnit = "Hours"
	DurationUnitMinutes      DurationUnit = "Minutes"
	DurationUnitSeconds      DurationUnit = "Seconds"
	DurationUnitMilliseconds DurationUnit = "Milliseconds"
	DurationUnitMicroseconds DurationUnit = "Microseconds"
	DurationUnitNanoseconds  DurationUnit = "Nanoseconds"
)

// Data type with attributes of a QosProfile
type QosProfile struct {
	// A unique name for identifying a specific QoS profile. This may follow different
	// formats depending on the service providers implementation. Some options
	// addresses:
	//
	//   - A UUID style string
	//   - Support for predefined profile names like `QOS_E`, `QOS_S`, `QOS_M`, and
	//     `QOS_L`
	//   - A searchable descriptive name
	Name string `json:"name,required" format:"string"`
	// The current status of the QoS Profile
	//
	//   - `ACTIVE`- QoS Profile is available to be used
	//   - `INACTIVE`- QoS Profile is not currently available to be deployed
	//   - `DEPRECATED`- QoS profile is actively being used in a QoD session, but can not
	//     be deployed in new QoD sessions
	//
	// Any of "ACTIVE", "INACTIVE", "DEPRECATED".
	Status QosProfileStatus `json:"status,required"`
	// A list of countries, and optionally networks, for which the API provider makes
	// the profile available
	CountryAvailability []QosProfileCountryAvailability `json:"countryAvailability"`
	// A description of the QoS profile.
	Description string `json:"description"`
	// Specification of duration
	Jitter Duration `json:"jitter"`
	// **NOTE**: l4sQueueType is experimental and could change or be removed in a
	// future release.
	//
	// Specifies the type of queue for L4S (Low Latency, Low Loss, Scalable Throughput)
	// traffic management. L4S is an advanced queue management approach designed to
	// provide ultra-low latency and high throughput for internet traffic, particularly
	// beneficial for interactive applications such as gaming, video conferencing, and
	// virtual reality.
	//
	// **Queue Type Descriptions:**
	//
	//   - **non-l4s-queue**: A traditional queue used for legacy internet traffic that
	//     does not utilize L4S enhancements. It provides standard latency and throughput
	//     levels.
	//
	//   - **l4s-queue**: A dedicated queue optimized for L4S traffic, delivering
	//     ultra-low latency, low loss, and scalable throughput to support
	//     latency-sensitive applications.
	//
	//   - **mixed-queue**: A shared queue that can handle both L4S and traditional
	//     traffic, offering a balance between ultra-low latency for L4S flows and
	//     compatibility with non-L4S flows.
	//
	// Any of "non-l4s-queue", "l4s-queue", "mixed-queue".
	L4sQueueType QosProfileL4sQueueType `json:"l4sQueueType"`
	// Specification of rate
	MaxDownstreamBurstRate Rate `json:"maxDownstreamBurstRate"`
	// Specification of rate
	MaxDownstreamRate Rate `json:"maxDownstreamRate"`
	// Specification of duration
	MaxDuration Duration `json:"maxDuration"`
	// Specification of rate
	MaxUpstreamBurstRate Rate `json:"maxUpstreamBurstRate"`
	// Specification of rate
	MaxUpstreamRate Rate `json:"maxUpstreamRate"`
	// Specification of duration
	MinDuration Duration `json:"minDuration"`
	// Specification of duration
	PacketDelayBudget Duration `json:"packetDelayBudget"`
	// This field specifies the acceptable level of data loss during transmission. The
	// value is an exponent of 10, so a value of 3 means that up to 10⁻³, or 0.1%, of
	// the data packets may be lost. This setting is part of a broader system that
	// categorizes different types of network traffic (like phone calls, video streams,
	// or data transfers) to ensure they perform reliably on the network.
	PacketErrorLossRate int64 `json:"packetErrorLossRate"`
	// Priority levels allow efficient resource allocation and ensure optimal
	// performance for various services in each technology, with the highest priority
	// traffic receiving preferential treatment. The lower value the higher priority.
	// Not all access networks use the same priority range, so this priority will be
	// scaled to the access network's priority range.
	Priority int64 `json:"priority"`
	// **NOTE**: serviceClass is experimental and could change or be removed in a
	// future release.
	//
	// The name of a Service Class, representing a QoS Profile designed to provide
	// optimized behavior for a specific application type. While DSCP values are
	// commonly associated with Service Classes, their use may vary across network
	// segments and may not be applied throughout the entire end-to-end QoS session.
	// This aligns with the serviceClass concept used in HomeDevicesQoQ for consistent
	// terminology.
	//
	// Service classes define specific QoS behaviors that map to DSCP (Differentiated
	// Services Code Point) values or Microsoft QoS traffic types.
	//
	// The supported mappings are:
	//
	//  1. Values aligned with the
	//     [RFC4594](https://datatracker.ietf.org/doc/html/rfc4594) guidelines for
	//     differentiated traffic classes.
	//  2. Microsoft
	//     [QOS_TRAFFIC_TYPE](https://learn.microsoft.com/en-us/windows/win32/api/qos2/ne-qos2-qos_traffic_type)
	//     values for Windows developers.
	//
	// **Supported Service Classes**:
	//
	// | Service Class Name    | DSCP Name | DSCP value (decimal) | DCSP value (binary) | Microsoft Value | Application Examples                                                 |
	// | --------------------- | --------- | -------------------- | ------------------- | --------------- | -------------------------------------------------------------------- |
	// | Microsoft Voice       | CS7       | 56                   | 111000              | 4,5             | Microsoft QOSTrafficTypeVoice and QOSTrafficTypeControl              |
	// | Microsoft Audio/Video | CS5       | 40                   | 101000              | 2,3             | Microsoft QOSTrafficTypeExcellentEffort and QOSTrafficTypeAudioVideo |
	// | Real-Time Interactive | CS4       | 32                   | 100000              |                 | Video conferencing and Interactive gaming                            |
	// | Multimedia Streaming  | AF31      | 26                   | 011010              |                 | Streaming video and audio on demand                                  |
	// | Broadcast Video       | CS3       | 24                   | 011000              |                 | Broadcast TV & live events                                           |
	// | Low-Latency Data      | AF21      | 18                   | 010010              |                 | Client/server transactions Web-based ordering                        |
	// | High-Throughput Data  | AF11      | 10                   | 001010              |                 | Store and forward applications                                       |
	// | Low-Priority Data     | CS1       | 8                    | 001000              | 1               | Any flow that has no BW assurance - also:                            |
	// |                       |           |                      |                     |                 | Microsoft QOSTrafficTypeBackground                                   |
	// | Standard              | DF(CS0)   | 0                    | 000000              | 0               | Undifferentiated applications - also:                                |
	// |                       |           |                      |                     |                 | Microsoft QOSTrafficTypeBestEffort                                   |
	//
	// Any of "microsoft_voice", "microsoft_audio_video", "real_time_interactive",
	// "multimedia_streaming", "broadcast_video", "low_latency_data",
	// "high_throughput_data", "low_priority_data", "standard".
	ServiceClass QosProfileServiceClass `json:"serviceClass"`
	// Specification of rate
	TargetMinDownstreamRate Rate `json:"targetMinDownstreamRate"`
	// Specification of rate
	TargetMinUpstreamRate Rate `json:"targetMinUpstreamRate"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name                    respjson.Field
		Status                  respjson.Field
		CountryAvailability     respjson.Field
		Description             respjson.Field
		Jitter                  respjson.Field
		L4sQueueType            respjson.Field
		MaxDownstreamBurstRate  respjson.Field
		MaxDownstreamRate       respjson.Field
		MaxDuration             respjson.Field
		MaxUpstreamBurstRate    respjson.Field
		MaxUpstreamRate         respjson.Field
		MinDuration             respjson.Field
		PacketDelayBudget       respjson.Field
		PacketErrorLossRate     respjson.Field
		Priority                respjson.Field
		ServiceClass            respjson.Field
		TargetMinDownstreamRate respjson.Field
		TargetMinUpstreamRate   respjson.Field
		ExtraFields             map[string]respjson.Field
		raw                     string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QosProfile) RawJSON() string { return r.JSON.raw }
func (r *QosProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type QosProfileCountryAvailability struct {
	// The two letter ISO 3166-2 country code for the country in which the QoS profile
	// is available in at least one network
	CountryName string `json:"countryName,required"`
	// A list of networks within the country for which the QoS profile is available
	// from the API provider
	Networks []string `json:"networks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CountryName respjson.Field
		Networks    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r QosProfileCountryAvailability) RawJSON() string { return r.JSON.raw }
func (r *QosProfileCountryAvailability) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **NOTE**: l4sQueueType is experimental and could change or be removed in a
// future release.
//
// Specifies the type of queue for L4S (Low Latency, Low Loss, Scalable Throughput)
// traffic management. L4S is an advanced queue management approach designed to
// provide ultra-low latency and high throughput for internet traffic, particularly
// beneficial for interactive applications such as gaming, video conferencing, and
// virtual reality.
//
// **Queue Type Descriptions:**
//
//   - **non-l4s-queue**: A traditional queue used for legacy internet traffic that
//     does not utilize L4S enhancements. It provides standard latency and throughput
//     levels.
//
//   - **l4s-queue**: A dedicated queue optimized for L4S traffic, delivering
//     ultra-low latency, low loss, and scalable throughput to support
//     latency-sensitive applications.
//
//   - **mixed-queue**: A shared queue that can handle both L4S and traditional
//     traffic, offering a balance between ultra-low latency for L4S flows and
//     compatibility with non-L4S flows.
type QosProfileL4sQueueType string

const (
	QosProfileL4sQueueTypeNonL4sQueue QosProfileL4sQueueType = "non-l4s-queue"
	QosProfileL4sQueueTypeL4sQueue    QosProfileL4sQueueType = "l4s-queue"
	QosProfileL4sQueueTypeMixedQueue  QosProfileL4sQueueType = "mixed-queue"
)

// **NOTE**: serviceClass is experimental and could change or be removed in a
// future release.
//
// The name of a Service Class, representing a QoS Profile designed to provide
// optimized behavior for a specific application type. While DSCP values are
// commonly associated with Service Classes, their use may vary across network
// segments and may not be applied throughout the entire end-to-end QoS session.
// This aligns with the serviceClass concept used in HomeDevicesQoQ for consistent
// terminology.
//
// Service classes define specific QoS behaviors that map to DSCP (Differentiated
// Services Code Point) values or Microsoft QoS traffic types.
//
// The supported mappings are:
//
//  1. Values aligned with the
//     [RFC4594](https://datatracker.ietf.org/doc/html/rfc4594) guidelines for
//     differentiated traffic classes.
//  2. Microsoft
//     [QOS_TRAFFIC_TYPE](https://learn.microsoft.com/en-us/windows/win32/api/qos2/ne-qos2-qos_traffic_type)
//     values for Windows developers.
//
// **Supported Service Classes**:
//
// | Service Class Name    | DSCP Name | DSCP value (decimal) | DCSP value (binary) | Microsoft Value | Application Examples                                                 |
// | --------------------- | --------- | -------------------- | ------------------- | --------------- | -------------------------------------------------------------------- |
// | Microsoft Voice       | CS7       | 56                   | 111000              | 4,5             | Microsoft QOSTrafficTypeVoice and QOSTrafficTypeControl              |
// | Microsoft Audio/Video | CS5       | 40                   | 101000              | 2,3             | Microsoft QOSTrafficTypeExcellentEffort and QOSTrafficTypeAudioVideo |
// | Real-Time Interactive | CS4       | 32                   | 100000              |                 | Video conferencing and Interactive gaming                            |
// | Multimedia Streaming  | AF31      | 26                   | 011010              |                 | Streaming video and audio on demand                                  |
// | Broadcast Video       | CS3       | 24                   | 011000              |                 | Broadcast TV & live events                                           |
// | Low-Latency Data      | AF21      | 18                   | 010010              |                 | Client/server transactions Web-based ordering                        |
// | High-Throughput Data  | AF11      | 10                   | 001010              |                 | Store and forward applications                                       |
// | Low-Priority Data     | CS1       | 8                    | 001000              | 1               | Any flow that has no BW assurance - also:                            |
// |                       |           |                      |                     |                 | Microsoft QOSTrafficTypeBackground                                   |
// | Standard              | DF(CS0)   | 0                    | 000000              | 0               | Undifferentiated applications - also:                                |
// |                       |           |                      |                     |                 | Microsoft QOSTrafficTypeBestEffort                                   |
type QosProfileServiceClass string

const (
	QosProfileServiceClassMicrosoftVoice      QosProfileServiceClass = "microsoft_voice"
	QosProfileServiceClassMicrosoftAudioVideo QosProfileServiceClass = "microsoft_audio_video"
	QosProfileServiceClassRealTimeInteractive QosProfileServiceClass = "real_time_interactive"
	QosProfileServiceClassMultimediaStreaming QosProfileServiceClass = "multimedia_streaming"
	QosProfileServiceClassBroadcastVideo      QosProfileServiceClass = "broadcast_video"
	QosProfileServiceClassLowLatencyData      QosProfileServiceClass = "low_latency_data"
	QosProfileServiceClassHighThroughputData  QosProfileServiceClass = "high_throughput_data"
	QosProfileServiceClassLowPriorityData     QosProfileServiceClass = "low_priority_data"
	QosProfileServiceClassStandard            QosProfileServiceClass = "standard"
)

// The current status of the QoS Profile
//
//   - `ACTIVE`- QoS Profile is available to be used
//   - `INACTIVE`- QoS Profile is not currently available to be deployed
//   - `DEPRECATED`- QoS profile is actively being used in a QoD session, but can not
//     be deployed in new QoD sessions
type QosProfileStatus string

const (
	QosProfileStatusActive     QosProfileStatus = "ACTIVE"
	QosProfileStatusInactive   QosProfileStatus = "INACTIVE"
	QosProfileStatusDeprecated QosProfileStatus = "DEPRECATED"
)

// Specification of rate
type Rate struct {
	// Units of rate
	//
	// Any of "bps", "kbps", "Mbps", "Gbps", "Tbps".
	Unit RateUnit `json:"unit"`
	// Quantity of rate
	Value int64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Unit        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Rate) RawJSON() string { return r.JSON.raw }
func (r *Rate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Units of rate
type RateUnit string

const (
	RateUnitBps  RateUnit = "bps"
	RateUnitKbps RateUnit = "kbps"
	RateUnitMbps RateUnit = "Mbps"
	RateUnitGbps RateUnit = "Gbps"
	RateUnitTbps RateUnit = "Tbps"
)

type QualityondemandGetQosProfileParams struct {
	// Value for the x-correlator
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type QualityondemandGetQosProfilesParams struct {
	// A unique name for identifying a specific QoS profile. This may follow different
	// formats depending on the service providers implementation. Some options
	// addresses:
	//
	//   - A UUID style string
	//   - Support for predefined profile names like `QOS_E`, `QOS_S`, `QOS_M`, and
	//     `QOS_L`
	//   - A searchable descriptive name
	Name param.Opt[string] `json:"name,omitzero" format:"string"`
	// Value for the x-correlator
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	// End-user equipment able to connect to a mobile network. Examples of devices
	// include smartphones or IoT sensors/actuators.
	//
	// The developer can choose to provide the below specified device identifiers:
	//
	//   - `ipv4Address`
	//   - `ipv6Address`
	//   - `phoneNumber` NOTE1: the network operator might support only a subset of these
	//     options. The API consumer can provide multiple identifiers to be compatible
	//     across different operators. In this case the identifiers MUST belong to the
	//     same device. NOTE2: as for this Commonalities release, we are enforcing that
	//     the networkAccessIdentifier is only part of the schema for future-proofing,
	//     and CAMARA does not currently allow its use. After the CAMARA meta-release
	//     work is concluded and the relevant issues are resolved, its use will need to
	//     be explicitly documented in the guidelines.
	Device QualityondemandGetQosProfilesParamsDevice `json:"device,omitzero"`
	// The current status of the QoS Profile
	//
	//   - `ACTIVE`- QoS Profile is available to be used
	//   - `INACTIVE`- QoS Profile is not currently available to be deployed
	//   - `DEPRECATED`- QoS profile is actively being used in a QoD session, but can not
	//     be deployed in new QoD sessions
	//
	// Any of "ACTIVE", "INACTIVE", "DEPRECATED".
	Status QosProfileStatus `json:"status,omitzero"`
	paramObj
}

func (r QualityondemandGetQosProfilesParams) MarshalJSON() (data []byte, err error) {
	type shadow QualityondemandGetQosProfilesParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QualityondemandGetQosProfilesParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// End-user equipment able to connect to a mobile network. Examples of devices
// include smartphones or IoT sensors/actuators.
//
// The developer can choose to provide the below specified device identifiers:
//
//   - `ipv4Address`
//   - `ipv6Address`
//   - `phoneNumber` NOTE1: the network operator might support only a subset of these
//     options. The API consumer can provide multiple identifiers to be compatible
//     across different operators. In this case the identifiers MUST belong to the
//     same device. NOTE2: as for this Commonalities release, we are enforcing that
//     the networkAccessIdentifier is only part of the schema for future-proofing,
//     and CAMARA does not currently allow its use. After the CAMARA meta-release
//     work is concluded and the relevant issues are resolved, its use will need to
//     be explicitly documented in the guidelines.
type QualityondemandGetQosProfilesParamsDevice struct {
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
	Ipv4Address QualityondemandGetQosProfilesParamsDeviceIpv4Address `json:"ipv4Address,omitzero"`
	paramObj
}

func (r QualityondemandGetQosProfilesParamsDevice) MarshalJSON() (data []byte, err error) {
	type shadow QualityondemandGetQosProfilesParamsDevice
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QualityondemandGetQosProfilesParamsDevice) UnmarshalJSON(data []byte) error {
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
type QualityondemandGetQosProfilesParamsDeviceIpv4Address struct {
	// A single IPv4 address with no subnet mask
	PrivateAddress param.Opt[string] `json:"privateAddress,omitzero" format:"ipv4"`
	// A single IPv4 address with no subnet mask
	PublicAddress param.Opt[string] `json:"publicAddress,omitzero" format:"ipv4"`
	// TCP or UDP port number
	PublicPort param.Opt[int64] `json:"publicPort,omitzero"`
	paramObj
}

func (r QualityondemandGetQosProfilesParamsDeviceIpv4Address) MarshalJSON() (data []byte, err error) {
	type shadow QualityondemandGetQosProfilesParamsDeviceIpv4Address
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *QualityondemandGetQosProfilesParamsDeviceIpv4Address) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
