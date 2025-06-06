// This file was auto-generated by Fern from our API Definition.

package transactions

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

// ## POST Transactions
//
// `/transactions` endpoint allows you to operate on the [Transaction entity.](/guides/overview/entities#transaction)
//
// In order to pass the payload of a transaction to Flagright and verify the transaction, you will need to call this endpoint with the transaction payload. Not all fields are mandatory, you will only need to pass in the fields that you have and are relevant for your compliance setup.
//
// ### Payload
//
// Here are some of the most used payload fields explained (you can find the full payload [schema below](/api-reference/api-reference/transactions/verify#request) with 1 line descriptions):
//
// * `type`: Type of transaction (Ex: `WITHDRAWAL`, `DEPOSIT`, `TRANSFER` etc).
// * `transactionId` - Unique Identifier for the transaction.
// * `timestamp` - UNIX timestamp in *milliseconds* of when the transaction took place
// * `transactionState` - The state of the transaction, set to `CREATED` by default. [More details here](/guides/overview/entities#transaction-lifecycle-through-transaction-events)
// * `originUserId` - Unique identifier (if any) of the user who is sending the money. This user must be created within the Flagright system before using the [create a consumer user](/api-reference/api-reference/consumer-users/create) or [create a business user](/api-reference/api-reference/business-users/create) endpoint
// * `destinationUserId` - Unique identifier (if any) of the user who is receiving the money. This user must be created within the Flagright system before using the [create a consumer user](/api-reference/api-reference/consumer-users/create) or [create a business user](/api-reference/api-reference/business-users/create) endpoint
// * `originAmountDetails` - Details of the amount being sent from the origin
// * `destinationAmountDetails` - Details of the amount being received at the destination
// * `originPaymentDetails` - Payment details (if any) used at the origin (ex: `CARD`, `IBAN`, `WALLET` etc). You can click on the dropdown next to the field in the schema below to view all supported payment types.
// * `destinationPaymentDetails` - Payment details (if any) used at the destination (ex: `CARD`, `IBAN`, `WALLET` etc). You can click on the dropdown next to the field in the schema below to view all supported payment types.
func (c *Client) Verify(
	ctx context.Context,
	request *flagrightgo.TransactionsVerifyRequest,
	opts ...option.RequestOption,
) (*flagrightgo.TransactionsVerifyResponse, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := baseURL + "/transactions"
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

	var response *flagrightgo.TransactionsVerifyResponse
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

// ### GET Transactions
//
// `/transactions` endpoint allows you to operate on the [Transaction entity](/guides/overview/entities#transaction).
//
// Calling `GET /transactions/{transactionId}` will return the entire transaction payload and rule execution results for the transaction with the corresponding `transactionId`
func (c *Client) Get(
	ctx context.Context,
	// Unique Transaction Identifier
	transactionId string,
	opts ...option.RequestOption,
) (*flagrightgo.TransactionWithRulesResult, error) {
	options := core.NewRequestOptions(opts...)
	baseURL := internal.ResolveBaseURL(
		options.BaseURL,
		c.baseURL,
		"https://sandbox.api.flagright.com",
	)
	endpointURL := internal.EncodeURL(
		baseURL+"/transactions/%v",
		transactionId,
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

	var response *flagrightgo.TransactionWithRulesResult
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
