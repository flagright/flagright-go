// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	internal "github.com/flagright/flagright-go/internal"
)

type BusinessUserEventBatchRequest struct {
	BatchId *string              `json:"batchId,omitempty" url:"-"`
	Data    []*BusinessUserEvent `json:"data,omitempty" url:"-"`
}

type BusinessBatchRequest struct {
	BatchId *string     `json:"batchId,omitempty" url:"-"`
	Data    []*Business `json:"data,omitempty" url:"-"`
}

type ConsumerUserEventBatchRequest struct {
	BatchId *string              `json:"batchId,omitempty" url:"-"`
	Data    []*ConsumerUserEvent `json:"data,omitempty" url:"-"`
}

type UserBatchRequest struct {
	BatchId *string `json:"batchId,omitempty" url:"-"`
	Data    []*User `json:"data,omitempty" url:"-"`
}

type TransactionEventBatchRequest struct {
	BatchId *string             `json:"batchId,omitempty" url:"-"`
	Data    []*TransactionEvent `json:"data,omitempty" url:"-"`
}

// Response from creation of a batch
type BatchResponse struct {
	Status        BatchResponseStatus          `json:"status" url:"status"`
	BatchId       string                       `json:"batchId" url:"batchId"`
	Successful    float64                      `json:"successful" url:"successful"`
	Failed        float64                      `json:"failed" url:"failed"`
	FailedRecords []*BatchResponseFailedRecord `json:"failedRecords,omitempty" url:"failedRecords,omitempty"`
	Message       *string                      `json:"message,omitempty" url:"message,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BatchResponse) GetStatus() BatchResponseStatus {
	if b == nil {
		return ""
	}
	return b.Status
}

func (b *BatchResponse) GetBatchId() string {
	if b == nil {
		return ""
	}
	return b.BatchId
}

func (b *BatchResponse) GetSuccessful() float64 {
	if b == nil {
		return 0
	}
	return b.Successful
}

func (b *BatchResponse) GetFailed() float64 {
	if b == nil {
		return 0
	}
	return b.Failed
}

func (b *BatchResponse) GetFailedRecords() []*BatchResponseFailedRecord {
	if b == nil {
		return nil
	}
	return b.FailedRecords
}

func (b *BatchResponse) GetMessage() *string {
	if b == nil {
		return nil
	}
	return b.Message
}

func (b *BatchResponse) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchResponse(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchResponse) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Failed record in a batch response
type BatchResponseFailedRecord struct {
	Id         *string `json:"id,omitempty" url:"id,omitempty"`
	ReasonCode *string `json:"reasonCode,omitempty" url:"reasonCode,omitempty"`

	extraProperties map[string]interface{}
	rawJSON         json.RawMessage
}

func (b *BatchResponseFailedRecord) GetId() *string {
	if b == nil {
		return nil
	}
	return b.Id
}

func (b *BatchResponseFailedRecord) GetReasonCode() *string {
	if b == nil {
		return nil
	}
	return b.ReasonCode
}

func (b *BatchResponseFailedRecord) GetExtraProperties() map[string]interface{} {
	return b.extraProperties
}

func (b *BatchResponseFailedRecord) UnmarshalJSON(data []byte) error {
	type unmarshaler BatchResponseFailedRecord
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*b = BatchResponseFailedRecord(value)
	extraProperties, err := internal.ExtractExtraProperties(data, *b)
	if err != nil {
		return err
	}
	b.extraProperties = extraProperties
	b.rawJSON = json.RawMessage(data)
	return nil
}

func (b *BatchResponseFailedRecord) String() string {
	if len(b.rawJSON) > 0 {
		if value, err := internal.StringifyJSON(b.rawJSON); err == nil {
			return value
		}
	}
	if value, err := internal.StringifyJSON(b); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", b)
}

// Status of the batch response
type BatchResponseStatus string

const (
	BatchResponseStatusSuccess        BatchResponseStatus = "SUCCESS"
	BatchResponseStatusPartialFailure BatchResponseStatus = "PARTIAL_FAILURE"
	BatchResponseStatusFailure        BatchResponseStatus = "FAILURE"
)

func NewBatchResponseStatusFromString(s string) (BatchResponseStatus, error) {
	switch s {
	case "SUCCESS":
		return BatchResponseStatusSuccess, nil
	case "PARTIAL_FAILURE":
		return BatchResponseStatusPartialFailure, nil
	case "FAILURE":
		return BatchResponseStatusFailure, nil
	}
	var t BatchResponseStatus
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (b BatchResponseStatus) Ptr() *BatchResponseStatus {
	return &b
}

type TransactionBatchRequest struct {
	// Boolean string whether Flagright should validate if provided originUserId exist. True by default
	ValidateOriginUserId *BooleanString `json:"-" url:"validateOriginUserId,omitempty"`
	// Boolean string whether Flagright should validate if provided destinationUserId exist. True by default
	ValidateDestinationUserId *BooleanString `json:"-" url:"validateDestinationUserId,omitempty"`
	BatchId                   *string        `json:"batchId,omitempty" url:"-"`
	Data                      []*Transaction `json:"data,omitempty" url:"-"`
}
