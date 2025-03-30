// This file was auto-generated by Fern from our API Definition.

package consumeruserevents

import (
	context "context"
	sdk "github.com/flagright/flagright-go/sdk"
	core "github.com/flagright/flagright-go/sdk/core"
	internal "github.com/flagright/flagright-go/sdk/internal"
	option "github.com/flagright/flagright-go/sdk/option"
	http "net/http"
)

type Client struct {
	baseURL string
	caller  *internal.Caller
	header  http.Header
}

func NewClient(opts ...option.RequestOption) *Client {
	options := core.NewRequestOptions(opts...)
	return &Client{
		baseURL: options.BaseURL,
		caller: internal.NewCaller(
			&internal.CallerParams{
				Client:      options.HTTPClient,
				MaxAttempts: options.MaxAttempts,
			},
		),
		header: options.ToHeader(),
	}
}

// ## POST Consumer User Events
//
// `/events/consumer/user` endpoint allows you to operate on the Consumer User Events entity.
//
// User events are created after the initial `POST /consumer/users` call (which creates a user) and are used to:
//
// * Update the STATE and KYC Status of the user, using the `userStateDetails` or `kycStatusDetails` field
// * Update the user details, using the `updatedConsumerUserAttributes` field.
//
// > If you have neither of the above two use cases, you do not need to use user events.
//
// ### Payload
//
// Each user event needs three mandatory fields:
//
// * `timestamp`- the timestamp of when the event was created or occured in your system
// * `userId` - The ID of the transaction for which this event is generated.
//
// In order to make individual events retrievable, you also need to pass in a unique `eventId` to the request body.
func (c *Client) Create(
	ctx context.Context,
	request *sdk.ConsumerUserEventsCreateRequest,
	opts ...option.RequestOption,
) (*sdk.UserWithRulesResult, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/events/consumer/user"
	queryParams, err := internal.QueryValues(request)
	if err != nil {
		return nil, err
	}
	if len(queryParams) > 0 {
		endpointURL += "?" + queryParams.Encode()
	}
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	headers.Set("Content-Type", "application/json")
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &sdk.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &sdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		409: func(apiError *core.APIError) error {
			return &sdk.ConflictError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &sdk.TooManyRequestsError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.UserWithRulesResult
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodPost,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Request:         request,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}

// ### GET a Consumer User Event
// You can retrieve any consumer user event you created using the [POST Consumer User Events](/api-reference/api-reference/consumer-user-events/create) call.
func (c *Client) Get(
	ctx context.Context,
	// Unique Consumer User Event Identifier
	eventId string,
	opts ...option.RequestOption,
) (*sdk.ConsumerUserEventWithRulesResult, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/events/consumer/user/%v",
		eventId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &sdk.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &sdk.UnauthorizedError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &sdk.TooManyRequestsError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.ConsumerUserEventWithRulesResult
	if err := c.caller.Call(
		ctx,
		&internal.CallParams{
			URL:             endpointURL,
			Method:          http.MethodGet,
			Headers:         headers,
			MaxAttempts:     options.MaxAttempts,
			BodyProperties:  options.BodyProperties,
			QueryParameters: options.QueryParameters,
			Client:          options.HTTPClient,
			Response:        &response,
			ErrorDecoder:    internal.NewErrorDecoder(errorCodes),
		},
	); err != nil {
		return nil, err
	}
	return response, nil
}
