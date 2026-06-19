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

func TestTenureVerifyWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := camara.NewClient(
		option.WithBaseURL(baseURL),
		option.WithBearerToken("My Bearer Token"),
		option.WithCustomerInsightsToken("My Customer Insights Token"),
		option.WithDeviceSwapToken("My Device Swap Token"),
		option.WithKYCAgeVerificationToken("My KYC Age Verification Token"),
		option.WithKYCFillInToken("My KYC Fill In Token"),
		option.WithKYCMatchToken("My KYC Match Token"),
		option.WithTenureToken("My Tenure Token"),
		option.WithNumberRecyclingToken("My Number Recycling Token"),
		option.WithOtpValidationToken("My Otp Validation Token"),
		option.WithCallForwardingSignalToken("My Call Forwarding Signal Token"),
		option.WithDeviceLocationToken("My Device Location Token"),
		option.WithPopulationDensityDataToken("My Population Density Data Token"),
		option.WithRegionDeviceCountToken("My Region Device Count Token"),
		option.WithWebRtcToken("My Web Rtc Token"),
		option.WithConnectivityInsightsToken("My Connectivity Insights Token"),
		option.WithQualityOnDemandToken("My Quality On Demand Token"),
		option.WithDeviceIdentifierToken("My Device Identifier Token"),
		option.WithSimSwapToken("My Sim Swap Token"),
		option.WithDeviceRoamingStatusToken("My Device Roaming Status Token"),
		option.WithDeviceReachabilityStatusToken("My Device Reachability Status Token"),
		option.WithConnectedNetworkTypeToken("My Connected Network Type Token"),
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
	_, err := client.Tenure.Verify(context.TODO(), camara.TenureVerifyParams{
		TenureDate:  time.Now(),
		PhoneNumber: camara.String("+123456789"),
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
