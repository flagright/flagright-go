# Flagright Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=https%3A%2F%2Fgithub.com%2Fflagright%2Fflagright-go)

The Flagright Go library provides convenient access to the Flagright APIs from Go.

## Table of Contents

- [Reference](#reference)
- [Usage](#usage)
- [Environments](#environments)
- [Errors](#errors)
- [Request Options](#request-options)
- [Advanced](#advanced)
  - [Response Headers](#response-headers)
  - [Retries](#retries)
  - [Timeouts](#timeouts)
  - [Explicit Null](#explicit-null)
- [Contributing](#contributing)

## Reference

A full reference for this library is available [here](https://github.com/flagright/flagright-go/blob/HEAD/./reference.md).

## Usage

Instantiate and use the client with the following:

```go
package example

import (
    context "context"

    flagrightgo "github.com/flagright/flagright-go"
    client "github.com/flagright/flagright-go/client"
    option "github.com/flagright/flagright-go/option"
)

func do() {
    client := client.NewClient(
        option.WithApiKey(
            "<value>",
        ),
    )
    request := &flagrightgo.TransactionsVerifyRequest{
        ValidateOriginUserId: flagrightgo.BooleanStringTrue.Ptr(),
        ValidateDestinationUserId: flagrightgo.BooleanStringTrue.Ptr(),
        Body: &flagrightgo.Transaction{
            Type: "DEPOSIT",
            TransactionId: "7b80a539eea6e78acbd6d458e5971482",
            Timestamp: 1641654664000,
            OriginUserId: flagrightgo.String(
                "8650a2611d0771cba03310f74bf6",
            ),
            DestinationUserId: flagrightgo.String(
                "9350a2611e0771cba03310f74bf6",
            ),
            OriginAmountDetails: &flagrightgo.TransactionAmountDetails{
                TransactionAmount: 2000,
                TransactionCurrency: flagrightgo.CurrencyCodeEur,
                Country: flagrightgo.CountryCodeDe.Ptr(),
            },
            DestinationAmountDetails: &flagrightgo.TransactionAmountDetails{
                TransactionAmount: 68351.34,
                TransactionCurrency: flagrightgo.CurrencyCodeInr,
                Country: flagrightgo.CountryCodeIn.Ptr(),
            },
            OriginPaymentDetails: &flagrightgo.TransactionOriginPaymentDetails{
                Card: &flagrightgo.CardDetails{
                    CardFingerprint: flagrightgo.String(
                        "20ac00fed8ef913aefb17cfae1097cce",
                    ),
                    CardIssuedCountry: flagrightgo.CountryCodeTr.Ptr(),
                    TransactionReferenceField: flagrightgo.String(
                        "Deposit",
                    ),
                    Field3DsDone: flagrightgo.Bool(
                        true,
                    ),
                },
            },
            DestinationPaymentDetails: &flagrightgo.TransactionDestinationPaymentDetails{
                Card: &flagrightgo.CardDetails{
                    CardFingerprint: flagrightgo.String(
                        "20ac00fed8ef913aefb17cfae1097cce",
                    ),
                    CardIssuedCountry: flagrightgo.CountryCodeTr.Ptr(),
                    TransactionReferenceField: flagrightgo.String(
                        "Deposit",
                    ),
                    Field3DsDone: flagrightgo.Bool(
                        true,
                    ),
                },
            },
            PromotionCodeUsed: flagrightgo.Bool(
                true,
            ),
            Reference: flagrightgo.String(
                "loan repayment",
            ),
            OriginDeviceData: &flagrightgo.DeviceData{
                BatteryLevel: flagrightgo.Float64(
                    95,
                ),
                DeviceLatitude: flagrightgo.Float64(
                    13.0033,
                ),
                DeviceLongitude: flagrightgo.Float64(
                    76.1004,
                ),
                IpAddress: flagrightgo.String(
                    "10.23.191.2",
                ),
                DeviceIdentifier: flagrightgo.String(
                    "3c49f915d04485e34caba",
                ),
                VpnUsed: flagrightgo.Bool(
                    false,
                ),
                OperatingSystem: flagrightgo.String(
                    "Android 11.2",
                ),
                DeviceMaker: flagrightgo.String(
                    "ASUS",
                ),
                DeviceModel: flagrightgo.String(
                    "Zenphone M2 Pro Max",
                ),
                DeviceYear: flagrightgo.String(
                    "2018",
                ),
                AppVersion: flagrightgo.String(
                    "1.1.0",
                ),
            },
            DestinationDeviceData: &flagrightgo.DeviceData{
                BatteryLevel: flagrightgo.Float64(
                    95,
                ),
                DeviceLatitude: flagrightgo.Float64(
                    13.0033,
                ),
                DeviceLongitude: flagrightgo.Float64(
                    76.1004,
                ),
                IpAddress: flagrightgo.String(
                    "10.23.191.2",
                ),
                DeviceIdentifier: flagrightgo.String(
                    "3c49f915d04485e34caba",
                ),
                VpnUsed: flagrightgo.Bool(
                    false,
                ),
                OperatingSystem: flagrightgo.String(
                    "Android 11.2",
                ),
                DeviceMaker: flagrightgo.String(
                    "ASUS",
                ),
                DeviceModel: flagrightgo.String(
                    "Zenphone M2 Pro Max",
                ),
                DeviceYear: flagrightgo.String(
                    "2018",
                ),
                AppVersion: flagrightgo.String(
                    "1.1.0",
                ),
            },
            Tags: []*flagrightgo.Tag{
                &flagrightgo.Tag{
                    Key: "customKey",
                    Value: "customValue",
                },
            },
        },
    }
    client.Transactions.Verify(
        context.TODO(),
        request,
    )
}
```

## Environments

You can choose between different environments by using the `option.WithBaseURL` option. You can configure any arbitrary base
URL, which is particularly useful in test environments.

```go
client := client.NewClient(
    option.WithBaseURL(api.Environments.SandboxApiServerEu1),
)
```

## Errors

Structured error types are returned from API calls that return non-success status codes. These errors are compatible
with the `errors.Is` and `errors.As` APIs, so you can access the error like so:

```go
response, err := client.Transactions.Verify(...)
if err != nil {
    var apiError *core.APIError
    if errors.As(err, apiError) {
        // Do something with the API error ...
    }
    return err
}
```

## Request Options

A variety of request options are included to adapt the behavior of the library, which includes configuring
authorization tokens, or providing your own instrumented `*http.Client`.

These request options can either be
specified on the client so that they're applied on every request, or for an individual request, like so:

> Providing your own `*http.Client` is recommended. Otherwise, the `http.DefaultClient` will be used,
> and your client will wait indefinitely for a response (unless the per-request, context-based timeout
> is used).

```go
// Specify default options applied on every request.
client := client.NewClient(
    option.WithToken("<YOUR_API_KEY>"),
    option.WithHTTPClient(
        &http.Client{
            Timeout: 5 * time.Second,
        },
    ),
)

// Specify options for an individual request.
response, err := client.Transactions.Verify(
    ...,
    option.WithToken("<YOUR_API_KEY>"),
)
```

## Advanced

### Response Headers

You can access the raw HTTP response data by using the `WithRawResponse` field on the client. This is useful
when you need to examine the response headers received from the API call. (When the endpoint is paginated,
the raw HTTP response data will be included automatically in the Page response object.)

```go
response, err := client.Transactions.WithRawResponse.Verify(...)
if err != nil {
    return err
}
fmt.Printf("Got response headers: %v", response.Header)
fmt.Printf("Got status code: %d", response.StatusCode)
```

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as long
as the request is deemed retryable and the number of retry attempts has not grown larger than the configured
retry limit (default: 2).

Which status codes are retried depends on the `retryStatusCodes` generator configuration:

**`legacy`** (current default): retries on
- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status#server_error_responses) (All server errors, including 500)

**`recommended`**: retries on
- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [502](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502) (Bad Gateway)
- [503](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503) (Service Unavailable)
- [504](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504) (Gateway Timeout)

If the `Retry-After` header is present in the response, the SDK will prioritize respecting its value exactly
over the default exponential backoff.

Use the `option.WithMaxAttempts` option to configure this behavior for the entire client or an individual request:

```go
client := client.NewClient(
    option.WithMaxAttempts(1),
)

response, err := client.Transactions.Verify(
    ...,
    option.WithMaxAttempts(1),
)
```

### Timeouts

Setting a timeout for each individual request is as simple as using the standard context library. Setting a one second timeout for an individual API call looks like the following:

```go
ctx, cancel := context.WithTimeout(ctx, time.Second)
defer cancel()

response, err := client.Transactions.Verify(ctx, ...)
```

### Explicit Null

If you want to send the explicit `null` JSON value through an optional parameter, you can use the setters\
that come with every object. Calling a setter method for a property will flip a bit in the `explicitFields`
bitfield for that setter's object; during serialization, any property with a flipped bit will have its
omittable status stripped, so zero or `nil` values will be sent explicitly rather than omitted altogether:

```go
type ExampleRequest struct {
    // An optional string parameter.
    Name *string `json:"name,omitempty" url:"-"`

    // Private bitmask of fields set to an explicit value and therefore not to be omitted
    explicitFields *big.Int `json:"-" url:"-"`
}

request := &ExampleRequest{}
request.SetName(nil)

response, err := client.Transactions.Verify(ctx, request, ...)
```

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!
