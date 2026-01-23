// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/stainless-sdks/camara-go/internal/requestconfig"
	"github.com/stainless-sdks/camara-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the camara API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options                         []option.RequestOption
	Customerinsights                CustomerinsightService
	Deviceswap                      DeviceswapService
	Knowyourcustomerageverification KnowyourcustomerageverificationService
	KnowyourcustomerfillIn          KnowyourcustomerfillInService
	Knowyourcustomermatch           KnowyourcustomermatchService
	Tenure                          TenureService
	Numberrecycling                 NumberrecyclingService
	Otpvalidation                   OtpvalidationService
	Callforwardingsignal            CallforwardingsignalService
	Devicelocation                  DevicelocationService
	Populationdensitydata           PopulationdensitydataService
	Regiondevicecount               RegiondevicecountService
	Webrtc                          WebrtcService
	Connectivityinsights            ConnectivityinsightService
	Qualityondemand                 QualityondemandService
	Deviceidentifier                DeviceidentifierService
	Simswap                         SimswapService
	Deviceroamingstatus             DeviceroamingstatusService
	Devicereachabilitystatus        DevicereachabilitystatusService
	Connectednetworktype            ConnectednetworktypeService
}

// DefaultClientOptions read from the environment
// (CAMARA_DEVICE_LOCATION_NOTIFICATIONS_API_KEY, CAMARA_NOTIFICATIONS_API_KEY,
// CAMARA_POPULATION_DENSITY_DATA_NOTIFICATIONS_API_KEY,
// CAMARA_REGION_DEVICE_COUNT_NOTIFICATIONS_API_KEY,
// CAMARA_CONNECTIVITY_INSIGHTS_NOTIFICATIONS_API_KEY,
// CAMARA_SIM_SWAP_NOTIFICATIONS_API_KEY,
// CAMARA_DEVICE_ROAMING_STATUS_NOTIFICATIONS_API_KEY,
// CAMARA_DEVICE_REACHABILITY_STATUS_NOTIFICATIONS_API_KEY,
// CAMARA_CONNECTED_NETWORK_TYPE_NOTIFICATIONS_API_KEY, CAMARA_BASE_URL). This
// should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("CAMARA_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("CAMARA_DEVICE_LOCATION_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithDeviceLocationNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_POPULATION_DENSITY_DATA_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithPopulationDensityDataNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_REGION_DEVICE_COUNT_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithRegionDeviceCountNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_CONNECTIVITY_INSIGHTS_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithConnectivityInsightsNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_SIM_SWAP_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithSimSwapNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_DEVICE_ROAMING_STATUS_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithDeviceRoamingStatusNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_DEVICE_REACHABILITY_STATUS_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithDeviceReachabilityStatusNotificationsAPIKey(o))
	}
	if o, ok := os.LookupEnv("CAMARA_CONNECTED_NETWORK_TYPE_NOTIFICATIONS_API_KEY"); ok {
		defaults = append(defaults, option.WithConnectedNetworkTypeNotificationsAPIKey(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (CAMARA_DEVICE_LOCATION_NOTIFICATIONS_API_KEY,
// CAMARA_NOTIFICATIONS_API_KEY,
// CAMARA_POPULATION_DENSITY_DATA_NOTIFICATIONS_API_KEY,
// CAMARA_REGION_DEVICE_COUNT_NOTIFICATIONS_API_KEY,
// CAMARA_CONNECTIVITY_INSIGHTS_NOTIFICATIONS_API_KEY,
// CAMARA_SIM_SWAP_NOTIFICATIONS_API_KEY,
// CAMARA_DEVICE_ROAMING_STATUS_NOTIFICATIONS_API_KEY,
// CAMARA_DEVICE_REACHABILITY_STATUS_NOTIFICATIONS_API_KEY,
// CAMARA_CONNECTED_NETWORK_TYPE_NOTIFICATIONS_API_KEY, CAMARA_BASE_URL). The
// option passed in as arguments are applied after these default arguments, and all
// option will be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Customerinsights = NewCustomerinsightService(opts...)
	r.Deviceswap = NewDeviceswapService(opts...)
	r.Knowyourcustomerageverification = NewKnowyourcustomerageverificationService(opts...)
	r.KnowyourcustomerfillIn = NewKnowyourcustomerfillInService(opts...)
	r.Knowyourcustomermatch = NewKnowyourcustomermatchService(opts...)
	r.Tenure = NewTenureService(opts...)
	r.Numberrecycling = NewNumberrecyclingService(opts...)
	r.Otpvalidation = NewOtpvalidationService(opts...)
	r.Callforwardingsignal = NewCallforwardingsignalService(opts...)
	r.Devicelocation = NewDevicelocationService(opts...)
	r.Populationdensitydata = NewPopulationdensitydataService(opts...)
	r.Regiondevicecount = NewRegiondevicecountService(opts...)
	r.Webrtc = NewWebrtcService(opts...)
	r.Connectivityinsights = NewConnectivityinsightService(opts...)
	r.Qualityondemand = NewQualityondemandService(opts...)
	r.Deviceidentifier = NewDeviceidentifierService(opts...)
	r.Simswap = NewSimswapService(opts...)
	r.Deviceroamingstatus = NewDeviceroamingstatusService(opts...)
	r.Devicereachabilitystatus = NewDevicereachabilitystatusService(opts...)
	r.Connectednetworktype = NewConnectednetworktypeService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
