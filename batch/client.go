// This file was auto-generated by Fern from our API Definition.

package batch

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

func (c *Client) VerifyTransaction(
	ctx context.Context,
	request *flagrightgo.TransactionBatchRequest,
	opts ...option.RequestOption,
) (*flagrightgo.BatchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/batch/transactions"
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

	var response *flagrightgo.BatchResponse
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

func (c *Client) Get(
	ctx context.Context,
	// Unique Batch Identifier
	batchId string,
	request *flagrightgo.BatchGetRequest,
	opts ...option.RequestOption,
) (*flagrightgo.BatchBusinessUserEventsWithRulesResult, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/batch/events/business/user/%v",
		batchId,
	)
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
	errorCodes := internal.ErrorCodes{
		401: func(apiError *core.APIError) error {
			return &flagrightgo.UnauthorizedError{
				APIError: apiError,
			}
		},
		404: func(apiError *core.APIError) error {
			return &flagrightgo.NotFoundError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &flagrightgo.TooManyRequestsError{
				APIError: apiError,
			}
		},
	}

	var response *flagrightgo.BatchBusinessUserEventsWithRulesResult
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

func (c *Client) CreateTransactionEvents(
	ctx context.Context,
	request *flagrightgo.TransactionEventBatchRequest,
	opts ...option.RequestOption,
) (*flagrightgo.BatchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/batch/events/transaction"
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
		409: func(apiError *core.APIError) error {
			return &flagrightgo.ConflictError{
				APIError: apiError,
			}
		},
		429: func(apiError *core.APIError) error {
			return &flagrightgo.TooManyRequestsError{
				APIError: apiError,
			}
		},
	}

	var response *flagrightgo.BatchResponse
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

func (c *Client) CreateConsumerUsers(
	ctx context.Context,
	request *flagrightgo.UserBatchRequest,
	opts ...option.RequestOption,
) (*flagrightgo.BatchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/batch/consumer/users"
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

	var response *flagrightgo.BatchResponse
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

func (c *Client) CreateBusinessUsers(
	ctx context.Context,
	request *flagrightgo.BusinessBatchRequest,
	opts ...option.RequestOption,
) (*flagrightgo.BatchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/batch/business/users"
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

	var response *flagrightgo.BatchResponse
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

func (c *Client) CreateConsumerUserEvents(
	ctx context.Context,
	request *flagrightgo.ConsumerUserEventBatchRequest,
	opts ...option.RequestOption,
) (*flagrightgo.BatchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/batch/events/consumer/user"
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

	var response *flagrightgo.BatchResponse
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

func (c *Client) CreateBusinessUserEvents(
	ctx context.Context,
	request *flagrightgo.BusinessUserEventBatchRequest,
	opts ...option.RequestOption,
) (*flagrightgo.BatchResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/batch/events/business/user"
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

	var response *flagrightgo.BatchResponse
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
