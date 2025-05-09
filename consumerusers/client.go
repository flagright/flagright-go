// This file was auto-generated by Fern from our API Definition.

package consumerusers

import (
	context "context"
	flagrightgo "github.com/flagright/flagright-go"
	core "github.com/flagright/flagright-go/core"
	internal "github.com/flagright/flagright-go/internal"
	option "github.com/flagright/flagright-go/option"
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

// ## POST Consumer User
//
// `/consumer/user` endpoint allows you to operate on the Consumer user entity.
//
// In order to pass the payload of a User to Flagright and verify the User, you will need to call this endpoint with the User payload. Not all fields are mandatory, you will only need to pass in the fields that you have and are relevant for your compliance setup.
//
// ### Payload
//
// Each consumer user needs two mandatory fields:
//
// * `userId` - Unique identifier for the user
// * `createdTimestamp` - UNIX timestamp in *milliseconds* for when the User is created in your system
func (c *Client) Create(
	ctx context.Context,
	request *flagrightgo.ConsumerUsersCreateRequest,
	opts ...option.RequestOption,
) (*flagrightgo.ConsumerUsersCreateResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/consumer/users"
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
			return &flagrightgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &flagrightgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &flagrightgo.TooManyRequestsError{
				APIError: apiError,
			}
		},
	}

	var response *flagrightgo.ConsumerUsersCreateResponse
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

// ### GET Consumer User
//
// `/consumer/user` endpoint allows you to operate on the Consumer User entity.
//
// Calling `GET /consumer/user/{userId}` will return the entire user payload and rule execution results for the user with the corresponding `userId`
func (c *Client) Get(
	ctx context.Context,
	userId string,
	opts ...option.RequestOption,
) (*flagrightgo.UserWithRulesResult, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/consumer/users/%v",
		userId,
	)
	headers := internal.MergeHeaders(
		c.header.Clone(),
		options.ToHeader(),
	)
	errorCodes := internal.ErrorCodes{
		400: func(apiError *core.APIError) error {
			return &flagrightgo.BadRequestError{
				APIError: apiError,
			}
		},
		401: func(apiError *core.APIError) error {
			return &flagrightgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &flagrightgo.TooManyRequestsError{
				APIError: apiError,
			}
		},
	}

	var response *flagrightgo.UserWithRulesResult
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
