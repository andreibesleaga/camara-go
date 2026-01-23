// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/stainless-sdks/camara-go"
	"github.com/stainless-sdks/camara-go/internal/testutil"
	"github.com/stainless-sdks/camara-go/option"
)

func TestKnowyourcustomerageverificationVerifyWithOptionalParams(t *testing.T) {
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
	_, err := client.Knowyourcustomerageverification.Verify(context.TODO(), camara.KnowyourcustomerageverificationVerifyParams{
		AgeThreshold:           18,
		Birthdate:              camara.Time(time.Now()),
		Email:                  camara.String("federicaSanchez.Arjona@example.com"),
		FamilyName:             camara.String("Sanchez Arjona"),
		FamilyNameAtBirth:      camara.String("YYYY"),
		GivenName:              camara.String("Federica"),
		IDDocument:             camara.String("66666666q"),
		IncludeContentLock:     camara.Bool(true),
		IncludeParentalControl: camara.Bool(true),
		MiddleNames:            camara.String("Sanchez"),
		Name:                   camara.String("Federica Sanchez Arjona"),
		PhoneNumber:            camara.String("+34629255833"),
		XCorrelator:            camara.String("b4333c46-49c0-4f62-80d7-f0ef930f1c46"),
	})
	if err != nil {
		var apierr *camara.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
