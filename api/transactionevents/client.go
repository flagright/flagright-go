// This file was auto-generated by Fern from our API Definition.

package transactionevents

import (
	context "context"
	http "net/http"

	sdk "github.com/flagright/flagright-go/api"
	core "github.com/flagright/flagright-go/api/core"
	internal "github.com/flagright/flagright-go/api/internal"
	option "github.com/flagright/flagright-go/api/option"
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

// ## POST Transaction Events
//
// `/events/transaction` endpoint allows you to operate on the [Transaction Events entity.](/guides/overview/entities#transaction-event)
//
// Transaction events are created after the initial `POST /transactions` call (which creates a transaction) and are used to:
//
// * Update the STATE of the transaction, using the `transactionState` field and manage the [Transaction Lifecycle](/guides/overview/entities#transaction-lifecycle-through-transaction-events)
// * Update the transaction details, using the `updatedTransactionAttributes` field.
//
// > If you have neither of the above two use cases, you do not need to use transaction events.
//
// ### Payload
//
// Each transaction event needs three mandatory fields:
//
// * `transactionState` - STATE of the transaction -> value is set to `CREATED` after `POST /transactions` call
// * `timestamp`- the timestamp of when the event was created or occured in your system
// * `transactionId` - The ID of the transaction for which this event is generated.
//
// In order to make individual events retrievable, you also need to pass in a unique `eventId` to the request body.
func (c *Client) Create(
	ctx context.Context,
	request *sdk.TransactionEvent,
	opts ...option.RequestOption,
) (*sdk.TransactionEventMonitoringResult, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/events/transaction"
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
		429: func(apiError *core.APIError) error {
			return &sdk.TooManyRequestsError{
				APIError: apiError,
			}
		},
	}

	var response *sdk.TransactionEventMonitoringResult
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

// ### GET Transaction Events
//
// `/events/transaction` endpoint allows you to operate on the [Transaction Events entity.](/guides/overview/entities#transaction-event).
//
// You can retrieve any transaction event you created using the [POST Transaction Events](/api-reference/api-reference/transaction-events/create) call.
func (c *Client) Get(
	ctx context.Context,
	// Unique Transaction Identifier
	eventId string,
	opts ...option.RequestOption,
) (*sdk.TransactionEventWithRulesResult, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/events/transaction/%v",
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

	var response *sdk.TransactionEventWithRulesResult
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
