// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/camara-go"
	"github.com/stainless-sdks/camara-go/internal/testutil"
	"github.com/stainless-sdks/camara-go/option"
)

func TestDeviceidentifierGetIdentifierWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := camara.NewClient(
		option.WithBaseURL(baseURL),
		option.WithDeviceLocationNotificationsAPIKey("My Device Location Notifications API Key"),
		option.WithNotificationsAPIKey("My Notifications API Key"),
		option.WithPopulationDensityDataNotificationsAPIKey("My Population Density Data Notifications API Key"),
		option.WithRegionDeviceCountNotificationsAPIKey("My Region Device Count Notifications API Key"),
		option.WithConnectivityInsightsNotificationsAPIKey("My Connectivity Insights Notifications API Key"),
		option.WithSimSwapNotificationsAPIKey("My Sim Swap Notifications API Key"),
		option.WithDeviceRoamingStatusNotificationsAPIKey("My Device Roaming Status Notifications API Key"),
		option.WithDeviceReachabilityStatusNotificationsAPIKey("My Device Reachability Status Notifications API Key"),
		option.WithConnectedNetworkTypeNotificationsAPIKey("My Connected Network Type Notifications API Key"),
	)
	_, err := client.Deviceidentifier.GetIdentifier(context.TODO(), camara.DeviceidentifierGetIdentifierParams{
		DeviceIdentifierRequestBody: camara.DeviceIdentifierRequestBodyParam{
			Device: camara.DeviceIdentifierRequestBodyDeviceParam{
				Ipv4Address: camara.DeviceIdentifierRequestBodyDeviceIpv4AddressParam{
					PrivateAddress: camara.String("84.125.93.10"),
					PublicAddress:  camara.String("84.125.93.10"),
					PublicPort:     camara.Int(59765),
				},
				Ipv6Address:             camara.String("2001:db8:85a3:8d3:1319:8a2e:370:7344"),
				NetworkAccessIdentifier: camara.String("123456789@example.com"),
				PhoneNumber:             camara.String("+123456789"),
			},
		},
		XCorrelator: camara.String("b4333c46-49c0-4f62-80d7-f0ef930f1c46"),
	})
	if err != nil {
		var apierr *camara.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceidentifierGetPpidWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := camara.NewClient(
		option.WithBaseURL(baseURL),
		option.WithDeviceLocationNotificationsAPIKey("My Device Location Notifications API Key"),
		option.WithNotificationsAPIKey("My Notifications API Key"),
		option.WithPopulationDensityDataNotificationsAPIKey("My Population Density Data Notifications API Key"),
		option.WithRegionDeviceCountNotificationsAPIKey("My Region Device Count Notifications API Key"),
		option.WithConnectivityInsightsNotificationsAPIKey("My Connectivity Insights Notifications API Key"),
		option.WithSimSwapNotificationsAPIKey("My Sim Swap Notifications API Key"),
		option.WithDeviceRoamingStatusNotificationsAPIKey("My Device Roaming Status Notifications API Key"),
		option.WithDeviceReachabilityStatusNotificationsAPIKey("My Device Reachability Status Notifications API Key"),
		option.WithConnectedNetworkTypeNotificationsAPIKey("My Connected Network Type Notifications API Key"),
	)
	_, err := client.Deviceidentifier.GetPpid(context.TODO(), camara.DeviceidentifierGetPpidParams{
		DeviceIdentifierRequestBody: camara.DeviceIdentifierRequestBodyParam{
			Device: camara.DeviceIdentifierRequestBodyDeviceParam{
				Ipv4Address: camara.DeviceIdentifierRequestBodyDeviceIpv4AddressParam{
					PrivateAddress: camara.String("84.125.93.10"),
					PublicAddress:  camara.String("84.125.93.10"),
					PublicPort:     camara.Int(59765),
				},
				Ipv6Address:             camara.String("2001:db8:85a3:8d3:1319:8a2e:370:7344"),
				NetworkAccessIdentifier: camara.String("123456789@example.com"),
				PhoneNumber:             camara.String("+123456789"),
			},
		},
		XCorrelator: camara.String("b4333c46-49c0-4f62-80d7-f0ef930f1c46"),
	})
	if err != nil {
		var apierr *camara.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestDeviceidentifierGetTypeWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := camara.NewClient(
		option.WithBaseURL(baseURL),
		option.WithDeviceLocationNotificationsAPIKey("My Device Location Notifications API Key"),
		option.WithNotificationsAPIKey("My Notifications API Key"),
		option.WithPopulationDensityDataNotificationsAPIKey("My Population Density Data Notifications API Key"),
		option.WithRegionDeviceCountNotificationsAPIKey("My Region Device Count Notifications API Key"),
		option.WithConnectivityInsightsNotificationsAPIKey("My Connectivity Insights Notifications API Key"),
		option.WithSimSwapNotificationsAPIKey("My Sim Swap Notifications API Key"),
		option.WithDeviceRoamingStatusNotificationsAPIKey("My Device Roaming Status Notifications API Key"),
		option.WithDeviceReachabilityStatusNotificationsAPIKey("My Device Reachability Status Notifications API Key"),
		option.WithConnectedNetworkTypeNotificationsAPIKey("My Connected Network Type Notifications API Key"),
	)
	_, err := client.Deviceidentifier.GetType(context.TODO(), camara.DeviceidentifierGetTypeParams{
		DeviceIdentifierRequestBody: camara.DeviceIdentifierRequestBodyParam{
			Device: camara.DeviceIdentifierRequestBodyDeviceParam{
				Ipv4Address: camara.DeviceIdentifierRequestBodyDeviceIpv4AddressParam{
					PrivateAddress: camara.String("84.125.93.10"),
					PublicAddress:  camara.String("84.125.93.10"),
					PublicPort:     camara.Int(59765),
				},
				Ipv6Address:             camara.String("2001:db8:85a3:8d3:1319:8a2e:370:7344"),
				NetworkAccessIdentifier: camara.String("123456789@example.com"),
				PhoneNumber:             camara.String("+123456789"),
			},
		},
		XCorrelator: camara.String("b4333c46-49c0-4f62-80d7-f0ef930f1c46"),
	})
	if err != nil {
		var apierr *camara.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
