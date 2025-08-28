# Flagright Go Library

[![fern shield](https://img.shields.io/badge/%F0%9F%8C%BF-Built%20with%20Fern-brightgreen)](https://buildwithfern.com?utm_source=github&utm_medium=github&utm_campaign=readme&utm_source=https%3A%2F%2Fgithub.com%2Fflagright%2Fflagright-go)

The Flagright Go library provides convenient access to the Flagright APIs from Go.

## Usage

Instantiate and use the client with the following:

```go
package example

import (
    client "github.com/flagright/flagright-go/client"
    option "github.com/flagright/flagright-go/option"
    context "context"
    flagright "github.com/flagright/flagright-go"
)

func do() () {
    client := client.NewClient(
        option.WithApiKey(
            "<value>",
        ),
    )
    client.Transactions.Verify(
        context.TODO(),
        &flagright.TransactionsVerifyRequest{
            Body: &flagright.Transaction{
                Type: "DEPOSIT",
                TransactionId: "7b80a539eea6e78acbd6d458e5971482",
                Timestamp: 1641654664000,
                OriginUserId: flagright.String(
                    "8650a2611d0771cba03310f74bf6",
                ),
                DestinationUserId: flagright.String(
                    "9350a2611e0771cba03310f74bf6",
                ),
                OriginAmountDetails: &flagright.TransactionAmountDetails{
                    TransactionAmount: 800,
                    TransactionCurrency: flagright.CurrencyCodeEur,
                    Country: flagright.CountryCodeDe.Ptr(),
                },
                DestinationAmountDetails: &flagright.TransactionAmountDetails{
                    TransactionAmount: 68351.34,
                    TransactionCurrency: flagright.CurrencyCodeInr,
                    Country: flagright.CountryCodeIn.Ptr(),
                },
                OriginPaymentDetails: &flagright.TransactionOriginPaymentDetails{
                    Card: &flagright.CardDetails{
                        CardFingerprint: flagright.String(
                            "20ac00fed8ef913aefb17cfae1097cce",
                        ),
                        CardIssuedCountry: flagright.CountryCodeTr.Ptr(),
                        TransactionReferenceField: flagright.String(
                            "Deposit",
                        ),
                    },
                },
                DestinationPaymentDetails: &flagright.TransactionDestinationPaymentDetails{
                    Card: &flagright.CardDetails{
                        CardFingerprint: flagright.String(
                            "20ac00fed8ef913aefb17cfae1097cce",
                        ),
                        CardIssuedCountry: flagright.CountryCodeTr.Ptr(),
                        TransactionReferenceField: flagright.String(
                            "Deposit",
                        ),
                    },
                },
                PromotionCodeUsed: flagright.Bool(
                    true,
                ),
                Reference: flagright.String(
                    "loan repayment",
                ),
                OriginDeviceData: &flagright.DeviceData{
                    BatteryLevel: flagright.Float64(
                        95,
                    ),
                    DeviceLatitude: flagright.Float64(
                        13.0033,
                    ),
                    DeviceLongitude: flagright.Float64(
                        76.1004,
                    ),
                    IpAddress: flagright.String(
                        "10.23.191.2",
                    ),
                    DeviceIdentifier: flagright.String(
                        "3c49f915d04485e34caba",
                    ),
                    VpnUsed: flagright.Bool(
                        false,
                    ),
                    OperatingSystem: flagright.String(
                        "Android 11.2",
                    ),
                    DeviceMaker: flagright.String(
                        "ASUS",
                    ),
                    DeviceModel: flagright.String(
                        "Zenphone M2 Pro Max",
                    ),
                    DeviceYear: flagright.String(
                        "2018",
                    ),
                    AppVersion: flagright.String(
                        "1.1.0",
                    ),
                },
                DestinationDeviceData: &flagright.DeviceData{
                    BatteryLevel: flagright.Float64(
                        95,
                    ),
                    DeviceLatitude: flagright.Float64(
                        13.0033,
                    ),
                    DeviceLongitude: flagright.Float64(
                        76.1004,
                    ),
                    IpAddress: flagright.String(
                        "10.23.191.2",
                    ),
                    DeviceIdentifier: flagright.String(
                        "3c49f915d04485e34caba",
                    ),
                    VpnUsed: flagright.Bool(
                        false,
                    ),
                    OperatingSystem: flagright.String(
                        "Android 11.2",
                    ),
                    DeviceMaker: flagright.String(
                        "ASUS",
                    ),
                    DeviceModel: flagright.String(
                        "Zenphone M2 Pro Max",
                    ),
                    DeviceYear: flagright.String(
                        "2018",
                    ),
                    AppVersion: flagright.String(
                        "1.1.0",
                    ),
                },
                Tags: []*flagright.Tag{
                    &flagright.Tag{
                        Key: "customKey",
                        Value: "customValue",
                    },
                },
            },
        },
    )
}
```

## Environments

You can choose between different environments by using the `option.WithBaseURL` option. You can configure any arbitrary base
URL, which is particularly useful in test environments.

```go
client := client.NewClient(
    option.WithBaseURL(flagright.Environments.SandboxApiServerEu1),
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

### Retries

The SDK is instrumented with automatic retries with exponential backoff. A request will be retried as long
as the request is deemed retryable and the number of retry attempts has not grown larger than the configured
retry limit (default: 2).

A request is deemed retryable when any of the following HTTP status codes is returned:

- [408](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408) (Timeout)
- [429](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429) (Too Many Requests)
- [5XX](https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500) (Internal Server Errors)

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

## Contributing

While we value open-source contributions to this SDK, this library is generated programmatically.
Additions made directly to this library would have to be moved over to our generation code,
otherwise they would be overwritten upon the next generated release. Feel free to open a PR as
a proof of concept, but know that we will not be able to merge it as-is. We suggest opening
an issue first to discuss with us!

On the other hand, contributions to the README are always very welcome!