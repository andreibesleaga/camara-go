// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara_test

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/andreibesleaga/camara-go"
	"github.com/andreibesleaga/camara-go/internal"
	"github.com/andreibesleaga/camara-go/option"
)

type closureTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *closureTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestUserAgentHeader(t *testing.T) {
	var userAgent string
	client := camara.NewClient(
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
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					userAgent = req.Header.Get("User-Agent")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)
	_, _ = client.Customerinsights.Scoring.Get(context.Background(), camara.CustomerinsightScoringGetParams{})
	if userAgent != fmt.Sprintf("Camara/Go %s", internal.PackageVersion) {
		t.Errorf("Expected User-Agent to be correct, but got: %#v", userAgent)
	}
}

func TestRetryAfter(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := camara.NewClient(
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
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Customerinsights.Scoring.Get(context.Background(), camara.CustomerinsightScoringGetParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	attempts := len(retryCountHeaders)
	if attempts != 3 {
		t.Errorf("Expected %d attempts, got %d", 3, attempts)
	}

	expectedRetryCountHeaders := []string{"0", "1", "2"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestDeleteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := camara.NewClient(
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
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeaderDel("X-Stainless-Retry-Count"),
	)
	_, err := client.Customerinsights.Scoring.Get(context.Background(), camara.CustomerinsightScoringGetParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"", "", ""}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestOverwriteRetryCountHeader(t *testing.T) {
	retryCountHeaders := make([]string, 0)
	client := camara.NewClient(
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
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					retryCountHeaders = append(retryCountHeaders, req.Header.Get("X-Stainless-Retry-Count"))
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After"): []string{"0.1"},
						},
					}, nil
				},
			},
		}),
		option.WithHeader("X-Stainless-Retry-Count", "42"),
	)
	_, err := client.Customerinsights.Scoring.Get(context.Background(), camara.CustomerinsightScoringGetParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}

	expectedRetryCountHeaders := []string{"42", "42", "42"}
	if !reflect.DeepEqual(retryCountHeaders, expectedRetryCountHeaders) {
		t.Errorf("Expected %v retry count headers, got %v", expectedRetryCountHeaders, retryCountHeaders)
	}
}

func TestRetryAfterMs(t *testing.T) {
	attempts := 0
	client := camara.NewClient(
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
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					attempts++
					return &http.Response{
						StatusCode: http.StatusTooManyRequests,
						Header: http.Header{
							http.CanonicalHeaderKey("Retry-After-Ms"): []string{"100"},
						},
					}, nil
				},
			},
		}),
	)
	_, err := client.Customerinsights.Scoring.Get(context.Background(), camara.CustomerinsightScoringGetParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
	if want := 3; attempts != want {
		t.Errorf("Expected %d attempts, got %d", want, attempts)
	}
}

func TestContextCancel(t *testing.T) {
	client := camara.NewClient(
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
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithCancel(context.Background())
	cancel()
	_, err := client.Customerinsights.Scoring.Get(cancelCtx, camara.CustomerinsightScoringGetParams{})
	if err == nil {
		t.Error("Expected there to be a cancel error")
	}
}

func TestContextCancelDelay(t *testing.T) {
	client := camara.NewClient(
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
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		}),
	)
	cancelCtx, cancel := context.WithTimeout(context.Background(), 2*time.Millisecond)
	defer cancel()
	_, err := client.Customerinsights.Scoring.Get(cancelCtx, camara.CustomerinsightScoringGetParams{})
	if err == nil {
		t.Error("expected there to be a cancel error")
	}
}

func TestContextDeadline(t *testing.T) {
	testTimeout := time.After(3 * time.Second)
	testDone := make(chan struct{})

	deadline := time.Now().Add(100 * time.Millisecond)
	deadlineCtx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	go func() {
		client := camara.NewClient(
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
			option.WithHTTPClient(&http.Client{
				Transport: &closureTransport{
					fn: func(req *http.Request) (*http.Response, error) {
						<-req.Context().Done()
						return nil, req.Context().Err()
					},
				},
			}),
		)
		_, err := client.Customerinsights.Scoring.Get(deadlineCtx, camara.CustomerinsightScoringGetParams{})
		if err == nil {
			t.Error("expected there to be a deadline error")
		}
		close(testDone)
	}()

	select {
	case <-testTimeout:
		t.Fatal("client didn't finish in time")
	case <-testDone:
		if diff := time.Since(deadline); diff < -30*time.Millisecond || 30*time.Millisecond < diff {
			t.Fatalf("client did not return within 30ms of context deadline, got %s", diff)
		}
	}
}
