/*
decentro-in-collections

Collections

API version: 1.0.0
Contact: admin@decentro.tech
*/

// Code generated by Konfig (https://konfigthis.com); DO NOT EDIT.

package decentro_in_collections

import (
	"encoding/json"
)

// IssueUpiRefundResponseData struct for IssueUpiRefundResponseData
type IssueUpiRefundResponseData struct {
	TransactionId *string `json:"transactionId,omitempty"`
	TransactionStatus *string `json:"transactionStatus,omitempty"`
	TransactionStatusDescription *string `json:"transactionStatusDescription,omitempty"`
}

// NewIssueUpiRefundResponseData instantiates a new IssueUpiRefundResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueUpiRefundResponseData() *IssueUpiRefundResponseData {
	this := IssueUpiRefundResponseData{}
	return &this
}

// NewIssueUpiRefundResponseDataWithDefaults instantiates a new IssueUpiRefundResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueUpiRefundResponseDataWithDefaults() *IssueUpiRefundResponseData {
	this := IssueUpiRefundResponseData{}
	return &this
}

// GetTransactionId returns the TransactionId field value if set, zero value otherwise.
func (o *IssueUpiRefundResponseData) GetTransactionId() string {
	if o == nil || isNil(o.TransactionId) {
		var ret string
		return ret
	}
	return *o.TransactionId
}

// GetTransactionIdOk returns a tuple with the TransactionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueUpiRefundResponseData) GetTransactionIdOk() (*string, bool) {
	if o == nil || isNil(o.TransactionId) {
    return nil, false
	}
	return o.TransactionId, true
}

// HasTransactionId returns a boolean if a field has been set.
func (o *IssueUpiRefundResponseData) HasTransactionId() bool {
	if o != nil && !isNil(o.TransactionId) {
		return true
	}

	return false
}

// SetTransactionId gets a reference to the given string and assigns it to the TransactionId field.
func (o *IssueUpiRefundResponseData) SetTransactionId(v string) {
	o.TransactionId = &v
}

// GetTransactionStatus returns the TransactionStatus field value if set, zero value otherwise.
func (o *IssueUpiRefundResponseData) GetTransactionStatus() string {
	if o == nil || isNil(o.TransactionStatus) {
		var ret string
		return ret
	}
	return *o.TransactionStatus
}

// GetTransactionStatusOk returns a tuple with the TransactionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueUpiRefundResponseData) GetTransactionStatusOk() (*string, bool) {
	if o == nil || isNil(o.TransactionStatus) {
    return nil, false
	}
	return o.TransactionStatus, true
}

// HasTransactionStatus returns a boolean if a field has been set.
func (o *IssueUpiRefundResponseData) HasTransactionStatus() bool {
	if o != nil && !isNil(o.TransactionStatus) {
		return true
	}

	return false
}

// SetTransactionStatus gets a reference to the given string and assigns it to the TransactionStatus field.
func (o *IssueUpiRefundResponseData) SetTransactionStatus(v string) {
	o.TransactionStatus = &v
}

// GetTransactionStatusDescription returns the TransactionStatusDescription field value if set, zero value otherwise.
func (o *IssueUpiRefundResponseData) GetTransactionStatusDescription() string {
	if o == nil || isNil(o.TransactionStatusDescription) {
		var ret string
		return ret
	}
	return *o.TransactionStatusDescription
}

// GetTransactionStatusDescriptionOk returns a tuple with the TransactionStatusDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueUpiRefundResponseData) GetTransactionStatusDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.TransactionStatusDescription) {
    return nil, false
	}
	return o.TransactionStatusDescription, true
}

// HasTransactionStatusDescription returns a boolean if a field has been set.
func (o *IssueUpiRefundResponseData) HasTransactionStatusDescription() bool {
	if o != nil && !isNil(o.TransactionStatusDescription) {
		return true
	}

	return false
}

// SetTransactionStatusDescription gets a reference to the given string and assigns it to the TransactionStatusDescription field.
func (o *IssueUpiRefundResponseData) SetTransactionStatusDescription(v string) {
	o.TransactionStatusDescription = &v
}

func (o IssueUpiRefundResponseData) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.TransactionId) {
		toSerialize["transactionId"] = o.TransactionId
	}
	if !isNil(o.TransactionStatus) {
		toSerialize["transactionStatus"] = o.TransactionStatus
	}
	if !isNil(o.TransactionStatusDescription) {
		toSerialize["transactionStatusDescription"] = o.TransactionStatusDescription
	}
	return json.Marshal(toSerialize)
}

type NullableIssueUpiRefundResponseData struct {
	value *IssueUpiRefundResponseData
	isSet bool
}

func (v NullableIssueUpiRefundResponseData) Get() *IssueUpiRefundResponseData {
	return v.value
}

func (v *NullableIssueUpiRefundResponseData) Set(val *IssueUpiRefundResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueUpiRefundResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueUpiRefundResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueUpiRefundResponseData(val *IssueUpiRefundResponseData) *NullableIssueUpiRefundResponseData {
	return &NullableIssueUpiRefundResponseData{value: val, isSet: true}
}

func (v NullableIssueUpiRefundResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueUpiRefundResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


