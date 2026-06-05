# Reference
## Transactions
<details><summary><code>client.Transactions.Verify(request) -> *flagrightgo.TransactionsVerifyResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

## POST Transactions

`/transactions` endpoint allows you to operate on the [Transaction entity.](/guides/overview/entities#transaction)

In order to pass the payload of a transaction to Flagright and verify the transaction, you will need to call this endpoint with the transaction payload. Not all fields are mandatory, you will only need to pass in the fields that you have and are relevant for your compliance setup.


### Payload

Here are some of the most used payload fields explained (you can find the full payload [schema below](/api-reference/api-reference/transactions/verify#request) with 1 line descriptions):

* `type`: Type of transaction (Ex: `WITHDRAWAL`, `DEPOSIT`, `TRANSFER` etc).
* `transactionId` - Unique Identifier for the transaction.
* `timestamp` - UNIX timestamp in *milliseconds* of when the transaction took place
* `transactionState` - The state of the transaction, set to `CREATED` by default. [More details here](/guides/overview/entities#transaction-lifecycle-through-transaction-events)
* `originUserId` - Unique identifier (if any) of the user who is sending the money. This user must be created within the Flagright system before using the [create a consumer user](/api-reference/api-reference/consumer-users/create) or [create a business user](/api-reference/api-reference/business-users/create) endpoint
* `destinationUserId` - Unique identifier (if any) of the user who is receiving the money. This user must be created within the Flagright system before using the [create a consumer user](/api-reference/api-reference/consumer-users/create) or [create a business user](/api-reference/api-reference/business-users/create) endpoint
* `originAmountDetails` - Details of the amount being sent from the origin
* `destinationAmountDetails` - Details of the amount being received at the destination
* `originPaymentDetails` - Payment details (if any) used at the origin (ex: `CARD`, `IBAN`, `WALLET` etc). You can click on the dropdown next to the field in the schema below to view all supported payment types.
* `destinationPaymentDetails` - Payment details (if any) used at the destination (ex: `CARD`, `IBAN`, `WALLET` etc). You can click on the dropdown next to the field in the schema below to view all supported payment types.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
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
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**validateOriginUserId:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should validate if provided originUserId exist. True by default
    
</dd>
</dl>

<dl>
<dd>

**validateDestinationUserId:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should validate if provided destinationUserId exist. True by default
    
</dd>
</dl>

<dl>
<dd>

**request:** `*flagrightgo.Transaction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Transactions.Get(TransactionId) -> *flagrightgo.TransactionWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

### GET Transactions

`/transactions` endpoint allows you to operate on the [Transaction entity](/guides/overview/entities#transaction).

Calling `GET /transactions/{transactionId}` will return the entire transaction payload and rule execution results for the transaction with the corresponding `transactionId`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.Transactions.Get(
        context.TODO(),
        "transactionId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**transactionId:** `string` — Unique Transaction Identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## Batch
<details><summary><code>client.Batch.VerifyTransaction(request) -> *flagrightgo.BatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.TransactionBatchRequest{
        ValidateOriginUserId: flagrightgo.BooleanStringTrue.Ptr(),
        ValidateDestinationUserId: flagrightgo.BooleanStringTrue.Ptr(),
        Data: []*flagrightgo.Transaction{
            &flagrightgo.Transaction{
                Type: "type",
                TransactionId: "transactionId",
                Timestamp: 1.1,
            },
        },
    }
client.Batch.VerifyTransaction(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**validateOriginUserId:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should validate if provided originUserId exist. True by default
    
</dd>
</dl>

<dl>
<dd>

**validateDestinationUserId:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should validate if provided destinationUserId exist. True by default
    
</dd>
</dl>

<dl>
<dd>

**batchId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `[]*flagrightgo.Transaction` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.GetTransactions(BatchId) -> *flagrightgo.BatchTransactionMonitoringResults</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BatchGetTransactionsRequest{
        PageSize: flagrightgo.Float64(
            1.1,
        ),
        Page: flagrightgo.Float64(
            1.1,
        ),
    }
client.Batch.GetTransactions(
        context.TODO(),
        "batchId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `string` — Unique Batch Identifier
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*flagrightgo.PageSize` — Page size (default 20)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*flagrightgo.Page` — Page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.CreateTransactionEvents(request) -> *flagrightgo.BatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.TransactionEventBatchRequest{
        Data: []*flagrightgo.TransactionEvent{
            &flagrightgo.TransactionEvent{
                TransactionState: flagrightgo.TransactionStateCreated,
                Timestamp: 1.1,
                TransactionId: "transactionId",
            },
        },
    }
client.Batch.CreateTransactionEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `[]*flagrightgo.TransactionEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.GetTransactionEvents(BatchId) -> *flagrightgo.BatchTransactionEventMonitoringResults</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BatchGetTransactionEventsRequest{
        PageSize: flagrightgo.Float64(
            1.1,
        ),
        Page: flagrightgo.Float64(
            1.1,
        ),
    }
client.Batch.GetTransactionEvents(
        context.TODO(),
        "batchId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `string` — Unique Batch Identifier
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*flagrightgo.PageSize` — Page size (default 20)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*flagrightgo.Page` — Page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.CreateConsumerUsers(request) -> *flagrightgo.BatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.UserBatchRequest{
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        Data: []*flagrightgo.User{
            &flagrightgo.User{
                UserId: "userId",
                CreatedTimestamp: 1.1,
            },
        },
    }
client.Batch.CreateConsumerUsers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**batchId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `[]*flagrightgo.User` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.GetConsumerUsers(BatchId) -> *flagrightgo.BatchConsumerUsersWithRulesResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BatchGetConsumerUsersRequest{
        PageSize: flagrightgo.Float64(
            1.1,
        ),
        Page: flagrightgo.Float64(
            1.1,
        ),
    }
client.Batch.GetConsumerUsers(
        context.TODO(),
        "batchId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `string` — Unique Batch Identifier
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*flagrightgo.PageSize` — Page size (default 20)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*flagrightgo.Page` — Page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.CreateBusinessUsers(request) -> *flagrightgo.BatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BusinessBatchRequest{
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        Data: []*flagrightgo.Business{
            &flagrightgo.Business{
                UserId: "userId",
                CreatedTimestamp: 1.1,
                LegalEntity: &flagrightgo.LegalEntity{
                    CompanyGeneralDetails: &flagrightgo.CompanyGeneralDetails{
                        LegalName: flagrightgo.String(
                            "Ozkan Hazelnut Export JSC",
                        ),
                        BusinessIndustry: []string{
                            "Farming",
                        },
                        SecondaryBusinessIndustry: []string{
                            "Food Processing",
                        },
                        MainProductsServicesSold: []string{
                            "Hazelnut",
                        },
                    },
                },
            },
        },
    }
client.Batch.CreateBusinessUsers(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**batchId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `[]*flagrightgo.Business` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.GetBusinessUsers(BatchId) -> *flagrightgo.BatchBusinessUsersWithRulesResults</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BatchGetBusinessUsersRequest{
        PageSize: flagrightgo.Float64(
            1.1,
        ),
        Page: flagrightgo.Float64(
            1.1,
        ),
    }
client.Batch.GetBusinessUsers(
        context.TODO(),
        "batchId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `string` — Unique Batch Identifier
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*flagrightgo.PageSize` — Page size (default 20)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*flagrightgo.Page` — Page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.CreateConsumerUserEvents(request) -> *flagrightgo.BatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.ConsumerUserEventBatchRequest{
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        Data: []*flagrightgo.ConsumerUserEvent{
            &flagrightgo.ConsumerUserEvent{
                Timestamp: 1.1,
                UserId: "userId",
            },
        },
    }
client.Batch.CreateConsumerUserEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**batchId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `[]*flagrightgo.ConsumerUserEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.GetConsumerUserEvents(BatchId) -> *flagrightgo.BatchConsumerUserEventsRulesResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BatchGetConsumerUserEventsRequest{
        PageSize: flagrightgo.Float64(
            1.1,
        ),
        Page: flagrightgo.Float64(
            1.1,
        ),
    }
client.Batch.GetConsumerUserEvents(
        context.TODO(),
        "batchId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `string` — Unique Batch Identifier
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*flagrightgo.PageSize` — Page size (default 20)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*flagrightgo.Page` — Page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.CreateBusinessUserEvents(request) -> *flagrightgo.BatchResponse</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BusinessUserEventBatchRequest{
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        Data: []*flagrightgo.BusinessUserEvent{
            &flagrightgo.BusinessUserEvent{
                Timestamp: 1.1,
                UserId: "userId",
            },
        },
    }
client.Batch.CreateBusinessUserEvents(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**batchId:** `*string` 
    
</dd>
</dl>

<dl>
<dd>

**data:** `[]*flagrightgo.BusinessUserEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.Batch.GetBusinessUserEvents(BatchId) -> *flagrightgo.BatchBusinessUserEventsWithRulesResult</code></summary>
<dl>
<dd>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BatchGetBusinessUserEventsRequest{
        PageSize: flagrightgo.Float64(
            1.1,
        ),
        Page: flagrightgo.Float64(
            1.1,
        ),
    }
client.Batch.GetBusinessUserEvents(
        context.TODO(),
        "batchId",
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**batchId:** `string` — Unique Batch Identifier
    
</dd>
</dl>

<dl>
<dd>

**pageSize:** `*flagrightgo.PageSize` — Page size (default 20)
    
</dd>
</dl>

<dl>
<dd>

**page:** `*flagrightgo.Page` — Page
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## TransactionEvents
<details><summary><code>client.TransactionEvents.Create(request) -> *flagrightgo.TransactionEventMonitoringResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

## POST Transaction Events

`/events/transaction` endpoint allows you to operate on the [Transaction Events entity.](/guides/overview/entities#transaction-event)

Transaction events are created after the initial `POST /transactions` call (which creates a transaction) and are used to:

* Update the STATE of the transaction, using the `transactionState` field and manage the [Transaction Lifecycle](/guides/overview/entities#transaction-lifecycle-through-transaction-events)
* Update the transaction details, using the `updatedTransactionAttributes` field.

> If you have neither of the above two use cases, you do not need to use transaction events.

### Payload

Each transaction event needs three mandatory fields:

* `transactionState` - STATE of the transaction -> value is set to `CREATED` after `POST /transactions` call
* `timestamp`- the timestamp of when the event was created or occured in your system
* `transactionId` - The ID of the transaction for which this event is generated.

In order to make individual events retrievable, you also need to pass in a unique `eventId` to the request body.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.TransactionEvent{
        TransactionState: flagrightgo.TransactionStateSuccessful,
        Timestamp: 1752526580000,
        TransactionId: "443dea26147a406b957d9ee3a1247b11",
        EventId: flagrightgo.String(
            "aaeeb166147a406b957dd9147a406b957",
        ),
        EventDescription: flagrightgo.String(
            "Transaction created",
        ),
        MetaData: &flagrightgo.DeviceData{
            BatteryLevel: flagrightgo.Float64(
                76.3,
            ),
            DeviceLatitude: flagrightgo.Float64(
                13.009711,
            ),
            DeviceLongitude: flagrightgo.Float64(
                76.102898,
            ),
            IpAddress: flagrightgo.String(
                "79.144.2.20",
            ),
            VpnUsed: flagrightgo.Bool(
                true,
            ),
        },
    }
client.TransactionEvents.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**request:** `*flagrightgo.TransactionEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.TransactionEvents.Get(EventId) -> *flagrightgo.TransactionEventWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

### GET Transaction Events

`/events/transaction` endpoint allows you to operate on the [Transaction Events entity.](/guides/overview/entities#transaction-event).

You can retrieve any transaction event you created using the [POST Transaction Events](/api-reference/api-reference/transaction-events/create) call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.TransactionEvents.Get(
        context.TODO(),
        "eventId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**eventId:** `string` — Unique Transaction Identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ConsumerUsers
<details><summary><code>client.ConsumerUsers.Create(request) -> *flagrightgo.ConsumerUsersCreateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

## POST Consumer User

`/consumer/user` endpoint allows you to operate on the Consumer user entity.

In order to pass the payload of a User to Flagright and verify the User, you will need to call this endpoint with the User payload. Not all fields are mandatory, you will only need to pass in the fields that you have and are relevant for your compliance setup.

### Payload

Each consumer user needs two mandatory fields:

* `userId` - Unique identifier for the user
* `createdTimestamp` - UNIX timestamp in *milliseconds* for when the User is created in your system
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.ConsumerUsersCreateRequest{
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        ValidateUserId: flagrightgo.BooleanStringTrue.Ptr(),
        Body: &flagrightgo.User{
            UserId: "96647cfd9e8fe66ee0f3362e011e34e8",
            CreatedTimestamp: 1641654664000,
            UserDetails: &flagrightgo.UserDetails{
                Name: &flagrightgo.ConsumerName{
                    FirstName: "Baran",
                    MiddleName: flagrightgo.String(
                        "Realblood",
                    ),
                    LastName: flagrightgo.String(
                        "Ozkan",
                    ),
                },
                DateOfBirth: flagrightgo.String(
                    "1991-01-01",
                ),
                CountryOfResidence: flagrightgo.CountryCodeUs.Ptr(),
                CountryOfNationality: flagrightgo.CountryCodeDe.Ptr(),
            },
            LegalDocuments: []*flagrightgo.LegalDocument{
                &flagrightgo.LegalDocument{
                    DocumentType: "passport",
                    DocumentNumber: "Z9431P",
                    DocumentIssuedDate: flagrightgo.Float64(
                        1639939034000,
                    ),
                    DocumentExpirationDate: flagrightgo.Float64(
                        1839939034000,
                    ),
                    DocumentIssuedCountry: flagrightgo.CountryCodeDe.Ptr(),
                    Tags: []*flagrightgo.Tag{
                        &flagrightgo.Tag{
                            Key: "customerType",
                            Value: "wallet",
                        },
                    },
                },
            },
            ContactDetails: &flagrightgo.ContactDetails{
                EmailIds: []string{
                    "baran@flagright.com",
                },
                ContactNumbers: []string{
                    "+37112345432",
                },
                Websites: []string{
                    "flagright.com",
                },
                Addresses: []*flagrightgo.Address{
                    &flagrightgo.Address{
                        AddressLines: []string{
                            "Klara-Franke Str 20",
                        },
                        Postcode: flagrightgo.String(
                            "10557",
                        ),
                        City: flagrightgo.String(
                            "Berlin",
                        ),
                        State: flagrightgo.String(
                            "Berlin",
                        ),
                        Country: flagrightgo.String(
                            "Germany",
                        ),
                        Tags: []*flagrightgo.Tag{
                            &flagrightgo.Tag{
                                Key: "customKey",
                                Value: "customValue",
                            },
                        },
                    },
                },
            },
            Tags: []*flagrightgo.UserTag{
                &flagrightgo.UserTag{
                    Key: "customKey",
                    Value: "customValue",
                },
            },
        },
    }
client.ConsumerUsers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**validateUserId:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should validate the userId
    
</dd>
</dl>

<dl>
<dd>

**request:** `*flagrightgo.User` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConsumerUsers.Get(UserId) -> *flagrightgo.UserWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

### GET Consumer User

`/consumer/user` endpoint allows you to operate on the Consumer User entity.

Calling `GET /consumer/user/{userId}` will return the entire user payload and rule execution results for the user with the corresponding `userId`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ConsumerUsers.Get(
        context.TODO(),
        "userId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BusinessUsers
<details><summary><code>client.BusinessUsers.Create(request) -> *flagrightgo.BusinessUsersCreateResponse</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

## POST Business User

`/business/user` endpoint allows you to operate on the Business user entity.

In order to pass the payload of a User to Flagright and verify the User, you will need to call this endpoint with the User payload. Not all fields are mandatory, you will only need to pass in the fields that you have and are relevant for your compliance setup.

### Payload


Each business user needs three mandatory fields:

* `userId` - Unique identifier for the user
* `legalEntity` - Details of the business legal entity (CompanyGeneralDetails, FinancialDetails etc) - only `legalName`in `CompanyGeneralDetails` is mandatory
* `createdTimestamp` - UNIX timestamp in *milliseconds* for when the User is created in your system
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BusinessUsersCreateRequest{
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        ValidateUserId: flagrightgo.BooleanStringTrue.Ptr(),
        Body: &flagrightgo.Business{
            UserId: "userId",
            CreatedTimestamp: 1.1,
            LegalEntity: &flagrightgo.LegalEntity{
                CompanyGeneralDetails: &flagrightgo.CompanyGeneralDetails{
                    LegalName: flagrightgo.String(
                        "Ozkan Hazelnut Export JSC",
                    ),
                    BusinessIndustry: []string{
                        "Farming",
                    },
                    SecondaryBusinessIndustry: []string{
                        "Food Processing",
                    },
                    MainProductsServicesSold: []string{
                        "Hazelnut",
                    },
                },
            },
        },
    }
client.BusinessUsers.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**validateUserId:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should validate the userId
    
</dd>
</dl>

<dl>
<dd>

**request:** `*flagrightgo.Business` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BusinessUsers.Get(UserId) -> *flagrightgo.BusinessWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

### GET Business User

`/business/user` endpoint allows you to operate on the Business User entity.

Calling `GET /business/user/{userId}` will return the entire User payload and rule execution results for the User with the corresponding `userId`
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BusinessUsers.Get(
        context.TODO(),
        "userId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**userId:** `string` — 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## ConsumerUserEvents
<details><summary><code>client.ConsumerUserEvents.Create(request) -> *flagrightgo.UserWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

## POST Consumer User Events

`/events/consumer/user` endpoint allows you to operate on the Consumer User Events entity.

User events are created after the initial `POST /consumer/users` call (which creates a user) and are used to:

* Update the STATE and KYC Status of the user, using the `userStateDetails` or `kycStatusDetails` field
* Update the user details, using the `updatedConsumerUserAttributes` field.

> If you have neither of the above two use cases, you do not need to use user events.

### Payload

Each user event needs three mandatory fields:

* `timestamp`- the timestamp of when the event was created or occured in your system
* `userId` - The ID of the transaction for which this event is generated.

In order to make individual events retrievable, you also need to pass in a unique `eventId` to the request body.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.ConsumerUserEventsCreateRequest{
        AllowUserTypeConversion: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        Body: &flagrightgo.ConsumerUserEvent{
            Timestamp: 1.1,
            UserId: "userId",
        },
    }
client.ConsumerUserEvents.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**allowUserTypeConversion:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should allow a Consumer user event to be applied to a Business user with the same user ID. This will converts a Business user to a Consumer user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*flagrightgo.ConsumerUserEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.ConsumerUserEvents.Get(EventId) -> *flagrightgo.ConsumerUserEventWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

### GET a Consumer User Event
You can retrieve any consumer user event you created using the [POST Consumer User Events](/api-reference/api-reference/consumer-user-events/create) call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.ConsumerUserEvents.Get(
        context.TODO(),
        "eventId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**eventId:** `string` — Unique Consumer User Event Identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

## BusinessUserEvents
<details><summary><code>client.BusinessUserEvents.Create(request) -> *flagrightgo.BusinessWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

## POST Business User Events

`/events/business/user` endpoint allows you to operate on the Business User Events entity.

User events are created after the initial `POST /business/users` call (which creates a user) and are used to:

* Update the STATE and KYC Status of the user, using the `userStateDetails` or `kycStatusDetails` field
* Update the user details, using the `updatedBusinessUserAttributes` field.

> If you have neither of the above two use cases, you do not need to use user events.

### Payload

Each user event needs three mandatory fields:

* `timestamp`- the timestamp of when the event was created or occured in your system
* `userId` - The ID of the transaction for which this event is generated.

In order to make individual events retrievable, you also need to pass in a unique `eventId` to the request body.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
request := &flagrightgo.BusinessUserEventsCreateRequest{
        AllowUserTypeConversion: flagrightgo.BooleanStringTrue.Ptr(),
        LockKycRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        LockCraRiskLevel: flagrightgo.BooleanStringTrue.Ptr(),
        Body: &flagrightgo.BusinessUserEvent{
            Timestamp: 1.1,
            UserId: "userId",
        },
    }
client.BusinessUserEvents.Create(
        context.TODO(),
        request,
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**allowUserTypeConversion:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should allow a Business user event to be applied to a Consumer user with the same user ID. This will converts a Consumer user to a Business user.
    
</dd>
</dl>

<dl>
<dd>

**lockKycRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the KYC risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**lockCraRiskLevel:** `*flagrightgo.BooleanString` — Boolean string whether Flagright should lock the CRA risk level for the user.
    
</dd>
</dl>

<dl>
<dd>

**request:** `*flagrightgo.BusinessUserEvent` 
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

<details><summary><code>client.BusinessUserEvents.Get(EventId) -> *flagrightgo.BusinessUserEventWithRulesResult</code></summary>
<dl>
<dd>

#### 📝 Description

<dl>
<dd>

<dl>
<dd>

### GET a Business User Event
You can retrieve any business user event you created using the [POST Business User Events](/api-reference/api-reference/business-user-events/create) call.
</dd>
</dl>
</dd>
</dl>

#### 🔌 Usage

<dl>
<dd>

<dl>
<dd>

```go
client.BusinessUserEvents.Get(
        context.TODO(),
        "eventId",
    )
}
```
</dd>
</dl>
</dd>
</dl>

#### ⚙️ Parameters

<dl>
<dd>

<dl>
<dd>

**eventId:** `string` — Unique Business User Event Identifier
    
</dd>
</dl>
</dd>
</dl>


</dd>
</dl>
</details>

