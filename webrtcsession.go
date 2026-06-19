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
	shimjson "github.com/andreibesleaga/camara-go/internal/encoding/json"
	"github.com/andreibesleaga/camara-go/internal/requestconfig"
	"github.com/andreibesleaga/camara-go/option"
	"github.com/andreibesleaga/camara-go/packages/param"
	"github.com/andreibesleaga/camara-go/packages/respjson"
)

// WebRTC Call Handling
//
// WebrtcSessionService contains methods and other services that help with
// interacting with the camara API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWebrtcSessionService] method instead.
type WebrtcSessionService struct {
	Options []option.RequestOption
}

// NewWebrtcSessionService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWebrtcSessionService(opts ...option.RequestOption) (r WebrtcSessionService) {
	r = WebrtcSessionService{}
	r.Options = opts
	return
}

// Creates a voice and/or video session
func (r *WebrtcSessionService) New(ctx context.Context, params WebrtcSessionNewParams, opts ...option.RequestOption) (res *MediaSessionInformation, err error) {
	if !param.IsOmitted(params.RegistrationID) {
		opts = append(opts, option.WithHeader("registrationId", fmt.Sprintf("%v", params.RegistrationID)))
	}
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "webrtc/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Get the media Session description based on `mediaSessionId`.
//
// ** The client shall construct the API path using the `mediaSessionId` supplied
// in the session creation response (origination) or in the invitation notification
// (termination). **
func (r *WebrtcSessionService) Get(ctx context.Context, mediaSessionID string, query WebrtcSessionGetParams, opts ...option.RequestOption) (res *MediaSessionInformation, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mediaSessionID == "" {
		err = errors.New("missing required mediaSessionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webrtc/sessions/%s", mediaSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Cancel a 1-1 media session (as originator), Decline a 1-1 media session (as
// receiver), Terminate a 1-1 an ongoing media session ** The client shall
// construct the API path using the mediaSessionId supplied in the session creation
// response (origination) or in the invitation notification (termination). **'
func (r *WebrtcSessionService) Delete(ctx context.Context, mediaSessionID string, body WebrtcSessionDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", body.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if mediaSessionID == "" {
		err = errors.New("missing required mediaSessionId parameter")
		return err
	}
	path := fmt.Sprintf("webrtc/sessions/%s", mediaSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Update the status of the media session, this may include updating SDP media
//
// The API consumer shall construct the API path using the `mediaSessionId`
// supplied in the session creation response (origination) or in the invitation
// notification (termination).
func (r *WebrtcSessionService) UpdateStatus(ctx context.Context, mediaSessionID string, params WebrtcSessionUpdateStatusParams, opts ...option.RequestOption) (res *MediaSessionInformation, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%v", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mediaSessionID == "" {
		err = errors.New("missing required mediaSessionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("webrtc/sessions/%s/status", mediaSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

type MediaSessionInformation struct {
	// **OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax
	// is used, the content of this element SHALL be embedded in a CDATA section.
	//
	// **ANSWER**: This type represents an answer in WebRTC Signaling. This element is
	// not present in case there is no answer yet, or the session invitation has been
	// declined by the Terminating Participant.This element MUST NOT be present in a
	// request from the application to the server to create a session.
	Answer SdpDescriptor `json:"answer"`
	// Type of call. When set to EMERGENCY, the client MAY provide locationDetails. If
	// omitted, treated as REGULAR.
	//
	// Any of "REGULAR", "EMERGENCY".
	CallType MediaSessionInformationCallType `json:"callType"`
	// Details about the caller's location and related information. This object adheres
	// to 3GPP TS 24.229, RFC 4119, RFC 5139, and RFC 5491 for PIDF-LO compatibility.
	LocationDetails WebRtcLocationDetails `json:"locationDetails"`
	// The media session ID created by the network. The mediaSessionId shall not be
	// included in POST requests by the client, but must be included in the
	// notifications from the network to the client device.
	MediaSessionID string `json:"mediaSessionId"`
	// **OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax
	// is used, the content of this element SHALL be embedded in a CDATA section.
	//
	// **ANSWER**: This type represents an answer in WebRTC Signaling. This element is
	// not present in case there is no answer yet, or the session invitation has been
	// declined by the Terminating Participant.This element MUST NOT be present in a
	// request from the application to the server to create a session.
	Offer SdpDescriptor `json:"offer"`
	// Subscriber address (Sender or Receiver)
	OriginatorAddress string `json:"originatorAddress"`
	// Friendly name of the call originator
	OriginatorName string `json:"originatorName"`
	// Subscriber address (Sender or Receiver)
	ReceiverAddress string `json:"receiverAddress"`
	// Friendly name of the call terminator
	ReceiverName string `json:"receiverName"`
	// Provides the status of the media session. During the session creation, this
	// attribute SHALL NOT be included in the request.
	//
	// Any of "Initial", "InProgress", "Ringing", "Proceeding", "Connected",
	// "Terminated", "Hold", "Resume", "SessionCancelled", "Declined", "Failed",
	// "Waiting", "NoAnswer", "NotReachable", "Busy".
	Status MediaSessionInformationStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Answer            respjson.Field
		CallType          respjson.Field
		LocationDetails   respjson.Field
		MediaSessionID    respjson.Field
		Offer             respjson.Field
		OriginatorAddress respjson.Field
		OriginatorName    respjson.Field
		ReceiverAddress   respjson.Field
		ReceiverName      respjson.Field
		Status            respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MediaSessionInformation) RawJSON() string { return r.JSON.raw }
func (r *MediaSessionInformation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this MediaSessionInformation to a MediaSessionInformationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MediaSessionInformationParam.Overrides()
func (r MediaSessionInformation) ToParam() MediaSessionInformationParam {
	return param.Override[MediaSessionInformationParam](json.RawMessage(r.RawJSON()))
}

// Type of call. When set to EMERGENCY, the client MAY provide locationDetails. If
// omitted, treated as REGULAR.
type MediaSessionInformationCallType string

const (
	MediaSessionInformationCallTypeRegular   MediaSessionInformationCallType = "REGULAR"
	MediaSessionInformationCallTypeEmergency MediaSessionInformationCallType = "EMERGENCY"
)

// Provides the status of the media session. During the session creation, this
// attribute SHALL NOT be included in the request.
type MediaSessionInformationStatus string

const (
	MediaSessionInformationStatusInitial          MediaSessionInformationStatus = "Initial"
	MediaSessionInformationStatusInProgress       MediaSessionInformationStatus = "InProgress"
	MediaSessionInformationStatusRinging          MediaSessionInformationStatus = "Ringing"
	MediaSessionInformationStatusProceeding       MediaSessionInformationStatus = "Proceeding"
	MediaSessionInformationStatusConnected        MediaSessionInformationStatus = "Connected"
	MediaSessionInformationStatusTerminated       MediaSessionInformationStatus = "Terminated"
	MediaSessionInformationStatusHold             MediaSessionInformationStatus = "Hold"
	MediaSessionInformationStatusResume           MediaSessionInformationStatus = "Resume"
	MediaSessionInformationStatusSessionCancelled MediaSessionInformationStatus = "SessionCancelled"
	MediaSessionInformationStatusDeclined         MediaSessionInformationStatus = "Declined"
	MediaSessionInformationStatusFailed           MediaSessionInformationStatus = "Failed"
	MediaSessionInformationStatusWaiting          MediaSessionInformationStatus = "Waiting"
	MediaSessionInformationStatusNoAnswer         MediaSessionInformationStatus = "NoAnswer"
	MediaSessionInformationStatusNotReachable     MediaSessionInformationStatus = "NotReachable"
	MediaSessionInformationStatusBusy             MediaSessionInformationStatus = "Busy"
)

type MediaSessionInformationParam struct {
	// The media session ID created by the network. The mediaSessionId shall not be
	// included in POST requests by the client, but must be included in the
	// notifications from the network to the client device.
	MediaSessionID param.Opt[string] `json:"mediaSessionId,omitzero"`
	// Subscriber address (Sender or Receiver)
	OriginatorAddress param.Opt[string] `json:"originatorAddress,omitzero"`
	// Friendly name of the call originator
	OriginatorName param.Opt[string] `json:"originatorName,omitzero"`
	// Subscriber address (Sender or Receiver)
	ReceiverAddress param.Opt[string] `json:"receiverAddress,omitzero"`
	// Friendly name of the call terminator
	ReceiverName param.Opt[string] `json:"receiverName,omitzero"`
	// **OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax
	// is used, the content of this element SHALL be embedded in a CDATA section.
	//
	// **ANSWER**: This type represents an answer in WebRTC Signaling. This element is
	// not present in case there is no answer yet, or the session invitation has been
	// declined by the Terminating Participant.This element MUST NOT be present in a
	// request from the application to the server to create a session.
	Answer SdpDescriptorParam `json:"answer,omitzero"`
	// Type of call. When set to EMERGENCY, the client MAY provide locationDetails. If
	// omitted, treated as REGULAR.
	//
	// Any of "REGULAR", "EMERGENCY".
	CallType MediaSessionInformationCallType `json:"callType,omitzero"`
	// Details about the caller's location and related information. This object adheres
	// to 3GPP TS 24.229, RFC 4119, RFC 5139, and RFC 5491 for PIDF-LO compatibility.
	LocationDetails WebRtcLocationDetailsParam `json:"locationDetails,omitzero"`
	// **OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax
	// is used, the content of this element SHALL be embedded in a CDATA section.
	//
	// **ANSWER**: This type represents an answer in WebRTC Signaling. This element is
	// not present in case there is no answer yet, or the session invitation has been
	// declined by the Terminating Participant.This element MUST NOT be present in a
	// request from the application to the server to create a session.
	Offer SdpDescriptorParam `json:"offer,omitzero"`
	// Provides the status of the media session. During the session creation, this
	// attribute SHALL NOT be included in the request.
	//
	// Any of "Initial", "InProgress", "Ringing", "Proceeding", "Connected",
	// "Terminated", "Hold", "Resume", "SessionCancelled", "Declined", "Failed",
	// "Waiting", "NoAnswer", "NotReachable", "Busy".
	Status MediaSessionInformationStatus `json:"status,omitzero"`
	paramObj
}

func (r MediaSessionInformationParam) MarshalJSON() (data []byte, err error) {
	type shadow MediaSessionInformationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MediaSessionInformationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// **OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax
// is used, the content of this element SHALL be embedded in a CDATA section.
//
// **ANSWER**: This type represents an answer in WebRTC Signaling. This element is
// not present in case there is no answer yet, or the session invitation has been
// declined by the Terminating Participant.This element MUST NOT be present in a
// request from the application to the server to create a session.
type SdpDescriptor struct {
	// An inlined session description in SDP format [RFC4566].If XML syntax is used,
	// the content of this element SHALL be embedded in a CDATA section
	Sdp string `json:"sdp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Sdp         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SdpDescriptor) RawJSON() string { return r.JSON.raw }
func (r *SdpDescriptor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SdpDescriptor to a SdpDescriptorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SdpDescriptorParam.Overrides()
func (r SdpDescriptor) ToParam() SdpDescriptorParam {
	return param.Override[SdpDescriptorParam](json.RawMessage(r.RawJSON()))
}

// **OFFER**: An inlined session description in SDP format [RFC4566].If XML syntax
// is used, the content of this element SHALL be embedded in a CDATA section.
//
// **ANSWER**: This type represents an answer in WebRTC Signaling. This element is
// not present in case there is no answer yet, or the session invitation has been
// declined by the Terminating Participant.This element MUST NOT be present in a
// request from the application to the server to create a session.
type SdpDescriptorParam struct {
	// An inlined session description in SDP format [RFC4566].If XML syntax is used,
	// the content of this element SHALL be embedded in a CDATA section
	Sdp param.Opt[string] `json:"sdp,omitzero"`
	paramObj
}

func (r SdpDescriptorParam) MarshalJSON() (data []byte, err error) {
	type shadow SdpDescriptorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SdpDescriptorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebRtcCircleCoordinates struct {
	// Latitude of the center point in decimal degrees (WGS84).
	Latitude float64 `json:"latitude" api:"required"`
	// Longitude of the center point in decimal degrees (WGS84).
	Longitude float64 `json:"longitude" api:"required"`
	// Radius of the circle in meters, indicating the uncertainty.
	Radius float64 `json:"radius" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude    respjson.Field
		Longitude   respjson.Field
		Radius      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebRtcCircleCoordinates) RawJSON() string { return r.JSON.raw }
func (r *WebRtcCircleCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WebRtcCircleCoordinates to a WebRtcCircleCoordinatesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WebRtcCircleCoordinatesParam.Overrides()
func (r WebRtcCircleCoordinates) ToParam() WebRtcCircleCoordinatesParam {
	return param.Override[WebRtcCircleCoordinatesParam](json.RawMessage(r.RawJSON()))
}

// The properties Latitude, Longitude, Radius are required.
type WebRtcCircleCoordinatesParam struct {
	// Latitude of the center point in decimal degrees (WGS84).
	Latitude float64 `json:"latitude" api:"required"`
	// Longitude of the center point in decimal degrees (WGS84).
	Longitude float64 `json:"longitude" api:"required"`
	// Radius of the circle in meters, indicating the uncertainty.
	Radius float64 `json:"radius" api:"required"`
	paramObj
}

func (r WebRtcCircleCoordinatesParam) MarshalJSON() (data []byte, err error) {
	type shadow WebRtcCircleCoordinatesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebRtcCircleCoordinatesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebRtcEllipsoidCoordinates struct {
	// Latitude in the WGS 84 geocentric coordinate system.
	Latitude float64 `json:"latitude" api:"required"`
	// Longitude in the WGS 84 geocentric coordinate system.
	Longitude float64 `json:"longitude" api:"required"`
	// Orientation of the ellipsoid in degrees.
	Orientation float64 `json:"orientation" api:"required"`
	// Length of the semi-major axis of the ellipsoid in meters.
	SemiMajorAxis float64 `json:"semiMajorAxis" api:"required"`
	// Length of the semi-minor axis of the ellipsoid in meters.
	SemiMinorAxis float64 `json:"semiMinorAxis" api:"required"`
	// Length of the vertical axis of the ellipsoid in meters.
	VerticalAxis float64 `json:"verticalAxis" api:"required"`
	// Altitude (optional) in the WGS 84 geocentric coordinate system.
	ZAxis float64 `json:"zAxis" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Latitude      respjson.Field
		Longitude     respjson.Field
		Orientation   respjson.Field
		SemiMajorAxis respjson.Field
		SemiMinorAxis respjson.Field
		VerticalAxis  respjson.Field
		ZAxis         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebRtcEllipsoidCoordinates) RawJSON() string { return r.JSON.raw }
func (r *WebRtcEllipsoidCoordinates) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WebRtcEllipsoidCoordinates to a
// WebRtcEllipsoidCoordinatesParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WebRtcEllipsoidCoordinatesParam.Overrides()
func (r WebRtcEllipsoidCoordinates) ToParam() WebRtcEllipsoidCoordinatesParam {
	return param.Override[WebRtcEllipsoidCoordinatesParam](json.RawMessage(r.RawJSON()))
}

// The properties Latitude, Longitude, Orientation, SemiMajorAxis, SemiMinorAxis,
// VerticalAxis, ZAxis are required.
type WebRtcEllipsoidCoordinatesParam struct {
	// Latitude in the WGS 84 geocentric coordinate system.
	Latitude float64 `json:"latitude" api:"required"`
	// Longitude in the WGS 84 geocentric coordinate system.
	Longitude float64 `json:"longitude" api:"required"`
	// Orientation of the ellipsoid in degrees.
	Orientation float64 `json:"orientation" api:"required"`
	// Length of the semi-major axis of the ellipsoid in meters.
	SemiMajorAxis float64 `json:"semiMajorAxis" api:"required"`
	// Length of the semi-minor axis of the ellipsoid in meters.
	SemiMinorAxis float64 `json:"semiMinorAxis" api:"required"`
	// Length of the vertical axis of the ellipsoid in meters.
	VerticalAxis float64 `json:"verticalAxis" api:"required"`
	// Altitude (optional) in the WGS 84 geocentric coordinate system.
	ZAxis float64 `json:"zAxis" api:"required"`
	paramObj
}

func (r WebRtcEllipsoidCoordinatesParam) MarshalJSON() (data []byte, err error) {
	type shadow WebRtcEllipsoidCoordinatesParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebRtcEllipsoidCoordinatesParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Details about the caller's location and related information. This object adheres
// to 3GPP TS 24.229, RFC 4119, RFC 5139, and RFC 5491 for PIDF-LO compatibility.
type WebRtcLocationDetails struct {
	// The confidence level of the location information.
	Confidence WebRtcLocationDetailsConfidence `json:"confidence"`
	// The coordinates of the caller's location, specific to the chosen shape.
	Coordinates WebRtcLocationDetailsCoordinatesUnion `json:"coordinates"`
	// The method used to obtain the location information.
	//
	// - **GPS:** Global Positioning System (highly accurate)
	// - **DBH:** Device-Based Hybrid
	// - **DBH_HELO:** Device-Based Hybrid using Apple Hybridized Emergency Location
	// - **Other:** Other methods (e.g., landmarks, IP Based etc.)
	//
	// Any of "GPS", "DBH", "DBH_HELO", "Other".
	Method WebRtcLocationDetailsMethod `json:"method"`
	// The shape representing the caller's location (Circle or Ellipsoid).
	//
	// Any of "Circle", "Ellipsoid".
	Shape WebRtcLocationDetailsShape `json:"shape"`
	// The timestamp (in ISO 8601 format) indicating when the location information was
	// Calculated. \nThis is crucial for emergency services to assess the timeliness of
	// the data. if not provided current timestamp will be used by default"
	Timestamp time.Time `json:"timestamp" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		Coordinates respjson.Field
		Method      respjson.Field
		Shape       respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebRtcLocationDetails) RawJSON() string { return r.JSON.raw }
func (r *WebRtcLocationDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WebRtcLocationDetails to a WebRtcLocationDetailsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WebRtcLocationDetailsParam.Overrides()
func (r WebRtcLocationDetails) ToParam() WebRtcLocationDetailsParam {
	return param.Override[WebRtcLocationDetailsParam](json.RawMessage(r.RawJSON()))
}

// The confidence level of the location information.
type WebRtcLocationDetailsConfidence struct {
	// The probability density function (PDF) associated with the confidence value.
	//
	// Any of "normal", "uniform".
	Pdf string `json:"pdf"`
	// The confidence value (percentage).
	Value float64 `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Pdf         respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WebRtcLocationDetailsConfidence) RawJSON() string { return r.JSON.raw }
func (r *WebRtcLocationDetailsConfidence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// WebRtcLocationDetailsCoordinatesUnion contains all possible properties and
// values from [WebRtcCircleCoordinates], [WebRtcEllipsoidCoordinates].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type WebRtcLocationDetailsCoordinatesUnion struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	// This field is from variant [WebRtcCircleCoordinates].
	Radius float64 `json:"radius"`
	// This field is from variant [WebRtcEllipsoidCoordinates].
	Orientation float64 `json:"orientation"`
	// This field is from variant [WebRtcEllipsoidCoordinates].
	SemiMajorAxis float64 `json:"semiMajorAxis"`
	// This field is from variant [WebRtcEllipsoidCoordinates].
	SemiMinorAxis float64 `json:"semiMinorAxis"`
	// This field is from variant [WebRtcEllipsoidCoordinates].
	VerticalAxis float64 `json:"verticalAxis"`
	// This field is from variant [WebRtcEllipsoidCoordinates].
	ZAxis float64 `json:"zAxis"`
	JSON  struct {
		Latitude      respjson.Field
		Longitude     respjson.Field
		Radius        respjson.Field
		Orientation   respjson.Field
		SemiMajorAxis respjson.Field
		SemiMinorAxis respjson.Field
		VerticalAxis  respjson.Field
		ZAxis         respjson.Field
		raw           string
	} `json:"-"`
}

func (u WebRtcLocationDetailsCoordinatesUnion) AsWebRtcCircleCoordinates() (v WebRtcCircleCoordinates) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u WebRtcLocationDetailsCoordinatesUnion) AsWebRtcEllipsoidCoordinates() (v WebRtcEllipsoidCoordinates) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u WebRtcLocationDetailsCoordinatesUnion) RawJSON() string { return u.JSON.raw }

func (r *WebRtcLocationDetailsCoordinatesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The method used to obtain the location information.
//
// - **GPS:** Global Positioning System (highly accurate)
// - **DBH:** Device-Based Hybrid
// - **DBH_HELO:** Device-Based Hybrid using Apple Hybridized Emergency Location
// - **Other:** Other methods (e.g., landmarks, IP Based etc.)
type WebRtcLocationDetailsMethod string

const (
	WebRtcLocationDetailsMethodGps     WebRtcLocationDetailsMethod = "GPS"
	WebRtcLocationDetailsMethodDbh     WebRtcLocationDetailsMethod = "DBH"
	WebRtcLocationDetailsMethodDbhHelo WebRtcLocationDetailsMethod = "DBH_HELO"
	WebRtcLocationDetailsMethodOther   WebRtcLocationDetailsMethod = "Other"
)

// The shape representing the caller's location (Circle or Ellipsoid).
type WebRtcLocationDetailsShape string

const (
	WebRtcLocationDetailsShapeCircle    WebRtcLocationDetailsShape = "Circle"
	WebRtcLocationDetailsShapeEllipsoid WebRtcLocationDetailsShape = "Ellipsoid"
)

// Details about the caller's location and related information. This object adheres
// to 3GPP TS 24.229, RFC 4119, RFC 5139, and RFC 5491 for PIDF-LO compatibility.
type WebRtcLocationDetailsParam struct {
	// The timestamp (in ISO 8601 format) indicating when the location information was
	// Calculated. \nThis is crucial for emergency services to assess the timeliness of
	// the data. if not provided current timestamp will be used by default"
	Timestamp param.Opt[time.Time] `json:"timestamp,omitzero" format:"date-time"`
	// The confidence level of the location information.
	Confidence WebRtcLocationDetailsConfidenceParam `json:"confidence,omitzero"`
	// The coordinates of the caller's location, specific to the chosen shape.
	Coordinates WebRtcLocationDetailsCoordinatesUnionParam `json:"coordinates,omitzero"`
	// The method used to obtain the location information.
	//
	// - **GPS:** Global Positioning System (highly accurate)
	// - **DBH:** Device-Based Hybrid
	// - **DBH_HELO:** Device-Based Hybrid using Apple Hybridized Emergency Location
	// - **Other:** Other methods (e.g., landmarks, IP Based etc.)
	//
	// Any of "GPS", "DBH", "DBH_HELO", "Other".
	Method WebRtcLocationDetailsMethod `json:"method,omitzero"`
	// The shape representing the caller's location (Circle or Ellipsoid).
	//
	// Any of "Circle", "Ellipsoid".
	Shape WebRtcLocationDetailsShape `json:"shape,omitzero"`
	paramObj
}

func (r WebRtcLocationDetailsParam) MarshalJSON() (data []byte, err error) {
	type shadow WebRtcLocationDetailsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebRtcLocationDetailsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The confidence level of the location information.
type WebRtcLocationDetailsConfidenceParam struct {
	// The confidence value (percentage).
	Value param.Opt[float64] `json:"value,omitzero"`
	// The probability density function (PDF) associated with the confidence value.
	//
	// Any of "normal", "uniform".
	Pdf string `json:"pdf,omitzero"`
	paramObj
}

func (r WebRtcLocationDetailsConfidenceParam) MarshalJSON() (data []byte, err error) {
	type shadow WebRtcLocationDetailsConfidenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WebRtcLocationDetailsConfidenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[WebRtcLocationDetailsConfidenceParam](
		"pdf", "normal", "uniform",
	)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type WebRtcLocationDetailsCoordinatesUnionParam struct {
	OfWebRtcCircleCoordinates    *WebRtcCircleCoordinatesParam    `json:",omitzero,inline"`
	OfWebRtcEllipsoidCoordinates *WebRtcEllipsoidCoordinatesParam `json:",omitzero,inline"`
	paramUnion
}

func (u WebRtcLocationDetailsCoordinatesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfWebRtcCircleCoordinates, u.OfWebRtcEllipsoidCoordinates)
}
func (u *WebRtcLocationDetailsCoordinatesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *WebRtcLocationDetailsCoordinatesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfWebRtcCircleCoordinates) {
		return u.OfWebRtcCircleCoordinates
	} else if !param.IsOmitted(u.OfWebRtcEllipsoidCoordinates) {
		return u.OfWebRtcEllipsoidCoordinates
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetRadius() *float64 {
	if vt := u.OfWebRtcCircleCoordinates; vt != nil {
		return &vt.Radius
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetOrientation() *float64 {
	if vt := u.OfWebRtcEllipsoidCoordinates; vt != nil {
		return &vt.Orientation
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetSemiMajorAxis() *float64 {
	if vt := u.OfWebRtcEllipsoidCoordinates; vt != nil {
		return &vt.SemiMajorAxis
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetSemiMinorAxis() *float64 {
	if vt := u.OfWebRtcEllipsoidCoordinates; vt != nil {
		return &vt.SemiMinorAxis
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetVerticalAxis() *float64 {
	if vt := u.OfWebRtcEllipsoidCoordinates; vt != nil {
		return &vt.VerticalAxis
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetZAxis() *float64 {
	if vt := u.OfWebRtcEllipsoidCoordinates; vt != nil {
		return &vt.ZAxis
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetLatitude() *float64 {
	if vt := u.OfWebRtcCircleCoordinates; vt != nil {
		return (*float64)(&vt.Latitude)
	} else if vt := u.OfWebRtcEllipsoidCoordinates; vt != nil {
		return (*float64)(&vt.Latitude)
	}
	return nil
}

// Returns a pointer to the underlying variant's property, if present.
func (u WebRtcLocationDetailsCoordinatesUnionParam) GetLongitude() *float64 {
	if vt := u.OfWebRtcCircleCoordinates; vt != nil {
		return (*float64)(&vt.Longitude)
	} else if vt := u.OfWebRtcEllipsoidCoordinates; vt != nil {
		return (*float64)(&vt.Longitude)
	}
	return nil
}

type WebrtcSessionNewParams struct {
	RegistrationID          string            `header:"registrationId" api:"required" json:"-"`
	XCorrelator             param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	MediaSessionInformation MediaSessionInformationParam
	paramObj
}

func (r WebrtcSessionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaSessionInformation)
}
func (r *WebrtcSessionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WebrtcSessionGetParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type WebrtcSessionDeleteParams struct {
	XCorrelator param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	paramObj
}

type WebrtcSessionUpdateStatusParams struct {
	XCorrelator             param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	MediaSessionInformation MediaSessionInformationParam
	paramObj
}

func (r WebrtcSessionUpdateStatusParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaSessionInformation)
}
func (r *WebrtcSessionUpdateStatusParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
