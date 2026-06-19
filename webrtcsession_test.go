// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/andreibesleaga/camara-go"
	"github.com/andreibesleaga/camara-go/internal/testutil"
	"github.com/andreibesleaga/camara-go/option"
)

func TestWebrtcSessionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Webrtc.Sessions.New(context.TODO(), camara.WebrtcSessionNewParams{
		RegistrationID: "registrationId",
		MediaSessionInformation: camara.MediaSessionInformationParam{
			Answer: camara.SdpDescriptorParam{
				Sdp: camara.String("sdp"),
			},
			CallType: camara.MediaSessionInformationCallTypeRegular,
			LocationDetails: camara.WebRtcLocationDetailsParam{
				Confidence: camara.WebRtcLocationDetailsConfidenceParam{
					Pdf:   "normal",
					Value: camara.Float(0),
				},
				Coordinates: camara.WebRtcLocationDetailsCoordinatesUnionParam{
					OfWebRtcCircleCoordinates: &camara.WebRtcCircleCoordinatesParam{
						Latitude:  0,
						Longitude: 0,
						Radius:    0,
					},
				},
				Method:    camara.WebRtcLocationDetailsMethodGps,
				Shape:     camara.WebRtcLocationDetailsShapeCircle,
				Timestamp: camara.Time(time.Now()),
			},
			MediaSessionID: camara.String("0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF"),
			Offer: camara.SdpDescriptorParam{
				Sdp: camara.String("v=0\r\no=- 8066321617929821805 2 IN IP4 127.0.0.1\r\ns=-\r\nt=0 0\r\nm=audio 42988 RTP/SAVPF 102 113\r\nc=IN IP6 2001:e0:410:2448:7a05:9b11:66f2:c9e\r\nb=AS:64\r\na=rtcp:9 IN IP4 0.0.0.0\r\na=candidate:1645903805 1 udp 2122262783 2001:e0:410:2448:7a05:9b11:66f2:c9e 42988 typ host generation 0 network-id 3 network-cost 900\r\na=ice-ufrag:4eKp\r\na=ice-pwd:D4sF5Pv9vx9ggaqxBlHbAFMx\r\na=ice-options:trickle renomination\r\na=mid:audio\r\na=extmap:2 http://www.ietf.org/id/draft-holmer-rmcat-transport-wide-cc-extensions-01\r\na=sendrecv\r\na=rtcp-mux\r\na=crypto:1 AES_CM_128_HMAC_SHA1_80 inline:Xm3YciqVIWFNSwy19e9MvfZ2YOdAZil7oT/tHjdf\r\na=rtpmap:102 AMR-WB/16000\r\na=fmtp:102 octet-align=0; mode-set=0,1,2; mode-change-capability=2\r\na=rtpmap:113 telephone-event/16000\r\n"),
			},
			OriginatorAddress: camara.String("tel:+17085852753"),
			OriginatorName:    camara.String("tel:+17085852753"),
			ReceiverAddress:   camara.String("tel:+17085854000"),
			ReceiverName:      camara.String("tel:+17085854000"),
			Status:            camara.MediaSessionInformationStatusRinging,
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

func TestWebrtcSessionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Webrtc.Sessions.Get(
		context.TODO(),
		"mediaSessionId",
		camara.WebrtcSessionGetParams{
			XCorrelator: camara.String("b4333c46-49c0-4f62-80d7-f0ef930f1c46"),
		},
	)
	if err != nil {
		var apierr *camara.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebrtcSessionDeleteWithOptionalParams(t *testing.T) {
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
	err := client.Webrtc.Sessions.Delete(
		context.TODO(),
		"mediaSessionId",
		camara.WebrtcSessionDeleteParams{
			XCorrelator: camara.String("b4333c46-49c0-4f62-80d7-f0ef930f1c46"),
		},
	)
	if err != nil {
		var apierr *camara.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestWebrtcSessionUpdateStatusWithOptionalParams(t *testing.T) {
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
	_, err := client.Webrtc.Sessions.UpdateStatus(
		context.TODO(),
		"mediaSessionId",
		camara.WebrtcSessionUpdateStatusParams{
			MediaSessionInformation: camara.MediaSessionInformationParam{
				Answer: camara.SdpDescriptorParam{
					Sdp: camara.String("sdp"),
				},
				CallType: camara.MediaSessionInformationCallTypeRegular,
				LocationDetails: camara.WebRtcLocationDetailsParam{
					Confidence: camara.WebRtcLocationDetailsConfidenceParam{
						Pdf:   "normal",
						Value: camara.Float(0),
					},
					Coordinates: camara.WebRtcLocationDetailsCoordinatesUnionParam{
						OfWebRtcCircleCoordinates: &camara.WebRtcCircleCoordinatesParam{
							Latitude:  0,
							Longitude: 0,
							Radius:    0,
						},
					},
					Method:    camara.WebRtcLocationDetailsMethodGps,
					Shape:     camara.WebRtcLocationDetailsShapeCircle,
					Timestamp: camara.Time(time.Now()),
				},
				MediaSessionID: camara.String("0AEE1B58BAEEDA3EABA42B32EBB3DFE07E9CFF402EAF9EED8EF"),
				Offer: camara.SdpDescriptorParam{
					Sdp: camara.String("sdp"),
				},
				OriginatorAddress: camara.String("tel:+11234567899"),
				OriginatorName:    camara.String("Alice"),
				ReceiverAddress:   camara.String("tel:+11234567899"),
				ReceiverName:      camara.String("Bob"),
				Status:            camara.MediaSessionInformationStatusRinging,
			},
			XCorrelator: camara.String("b4333c46-49c0-4f62-80d7-f0ef930f1c46"),
		},
	)
	if err != nil {
		var apierr *camara.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
