// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package camara

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/camara-go/internal/apijson"
	shimjson "github.com/stainless-sdks/camara-go/internal/encoding/json"
	"github.com/stainless-sdks/camara-go/internal/requestconfig"
	"github.com/stainless-sdks/camara-go/option"
	"github.com/stainless-sdks/camara-go/packages/param"
	"github.com/stainless-sdks/camara-go/packages/respjson"
)

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
		opts = append(opts, option.WithHeader("registrationId", fmt.Sprintf("%s", params.RegistrationID)))
	}
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	path := "webrtc/sessions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Get the media Session description based on `mediaSessionId`.
//
// ** The client shall construct the API path using the `mediaSessionId` supplied
// in the session creation response (origination) or in the invitation notification
// (termination). **
func (r *WebrtcSessionService) Get(ctx context.Context, mediaSessionID string, query WebrtcSessionGetParams, opts ...option.RequestOption) (res *MediaSessionInformation, err error) {
	if !param.IsOmitted(query.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", query.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mediaSessionID == "" {
		err = errors.New("missing required mediaSessionId parameter")
		return
	}
	path := fmt.Sprintf("webrtc/sessions/%s", mediaSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Cancel a 1-1 media session (as originator), Decline a 1-1 media session (as
// receiver), Terminate a 1-1 an ongoing media session ** The client shall
// construct the API path using the mediaSessionId supplied in the session creation
// response (origination) or in the invitation notification (termination). **'
func (r *WebrtcSessionService) Delete(ctx context.Context, mediaSessionID string, body WebrtcSessionDeleteParams, opts ...option.RequestOption) (err error) {
	if !param.IsOmitted(body.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", body.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if mediaSessionID == "" {
		err = errors.New("missing required mediaSessionId parameter")
		return
	}
	path := fmt.Sprintf("webrtc/sessions/%s", mediaSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Update the status of the media session, this may include updating SDP media
//
// The API consumer shall construct the API path using the `mediaSessionId`
// supplied in the session creation response (origination) or in the invitation
// notification (termination).
func (r *WebrtcSessionService) UpdateStatus(ctx context.Context, mediaSessionID string, params WebrtcSessionUpdateStatusParams, opts ...option.RequestOption) (res *MediaSessionInformation, err error) {
	if !param.IsOmitted(params.XCorrelator) {
		opts = append(opts, option.WithHeader("x-correlator", fmt.Sprintf("%s", params.XCorrelator.Value)))
	}
	opts = slices.Concat(r.Options, opts)
	if mediaSessionID == "" {
		err = errors.New("missing required mediaSessionId parameter")
		return
	}
	path := fmt.Sprintf("webrtc/sessions/%s/status", mediaSessionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
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

type WebrtcSessionNewParams struct {
	RegistrationID          string            `header:"registrationId,required" json:"-"`
	XCorrelator             param.Opt[string] `header:"x-correlator,omitzero" json:"-"`
	MediaSessionInformation MediaSessionInformationParam
	paramObj
}

func (r WebrtcSessionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.MediaSessionInformation)
}
func (r *WebrtcSessionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.MediaSessionInformation)
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
	return json.Unmarshal(data, &r.MediaSessionInformation)
}
