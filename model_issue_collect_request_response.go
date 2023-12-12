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

// IssueCollectRequestResponse struct for IssueCollectRequestResponse
type IssueCollectRequestResponse struct {
	DecentroTxnId *string `json:"decentroTxnId,omitempty"`
	Status *string `json:"status,omitempty"`
	ResponseCode *string `json:"responseCode,omitempty"`
	Message *string `json:"message,omitempty"`
	Data *IssueCollectRequestResponseData `json:"data,omitempty"`
}

// NewIssueCollectRequestResponse instantiates a new IssueCollectRequestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIssueCollectRequestResponse() *IssueCollectRequestResponse {
	this := IssueCollectRequestResponse{}
	return &this
}

// NewIssueCollectRequestResponseWithDefaults instantiates a new IssueCollectRequestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIssueCollectRequestResponseWithDefaults() *IssueCollectRequestResponse {
	this := IssueCollectRequestResponse{}
	return &this
}

// GetDecentroTxnId returns the DecentroTxnId field value if set, zero value otherwise.
func (o *IssueCollectRequestResponse) GetDecentroTxnId() string {
	if o == nil || isNil(o.DecentroTxnId) {
		var ret string
		return ret
	}
	return *o.DecentroTxnId
}

// GetDecentroTxnIdOk returns a tuple with the DecentroTxnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCollectRequestResponse) GetDecentroTxnIdOk() (*string, bool) {
	if o == nil || isNil(o.DecentroTxnId) {
    return nil, false
	}
	return o.DecentroTxnId, true
}

// HasDecentroTxnId returns a boolean if a field has been set.
func (o *IssueCollectRequestResponse) HasDecentroTxnId() bool {
	if o != nil && !isNil(o.DecentroTxnId) {
		return true
	}

	return false
}

// SetDecentroTxnId gets a reference to the given string and assigns it to the DecentroTxnId field.
func (o *IssueCollectRequestResponse) SetDecentroTxnId(v string) {
	o.DecentroTxnId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *IssueCollectRequestResponse) GetStatus() string {
	if o == nil || isNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCollectRequestResponse) GetStatusOk() (*string, bool) {
	if o == nil || isNil(o.Status) {
    return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *IssueCollectRequestResponse) HasStatus() bool {
	if o != nil && !isNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *IssueCollectRequestResponse) SetStatus(v string) {
	o.Status = &v
}

// GetResponseCode returns the ResponseCode field value if set, zero value otherwise.
func (o *IssueCollectRequestResponse) GetResponseCode() string {
	if o == nil || isNil(o.ResponseCode) {
		var ret string
		return ret
	}
	return *o.ResponseCode
}

// GetResponseCodeOk returns a tuple with the ResponseCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCollectRequestResponse) GetResponseCodeOk() (*string, bool) {
	if o == nil || isNil(o.ResponseCode) {
    return nil, false
	}
	return o.ResponseCode, true
}

// HasResponseCode returns a boolean if a field has been set.
func (o *IssueCollectRequestResponse) HasResponseCode() bool {
	if o != nil && !isNil(o.ResponseCode) {
		return true
	}

	return false
}

// SetResponseCode gets a reference to the given string and assigns it to the ResponseCode field.
func (o *IssueCollectRequestResponse) SetResponseCode(v string) {
	o.ResponseCode = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *IssueCollectRequestResponse) GetMessage() string {
	if o == nil || isNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCollectRequestResponse) GetMessageOk() (*string, bool) {
	if o == nil || isNil(o.Message) {
    return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *IssueCollectRequestResponse) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *IssueCollectRequestResponse) SetMessage(v string) {
	o.Message = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *IssueCollectRequestResponse) GetData() IssueCollectRequestResponseData {
	if o == nil || isNil(o.Data) {
		var ret IssueCollectRequestResponseData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IssueCollectRequestResponse) GetDataOk() (*IssueCollectRequestResponseData, bool) {
	if o == nil || isNil(o.Data) {
    return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *IssueCollectRequestResponse) HasData() bool {
	if o != nil && !isNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given IssueCollectRequestResponseData and assigns it to the Data field.
func (o *IssueCollectRequestResponse) SetData(v IssueCollectRequestResponseData) {
	o.Data = &v
}

func (o IssueCollectRequestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.DecentroTxnId) {
		toSerialize["decentroTxnId"] = o.DecentroTxnId
	}
	if !isNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !isNil(o.ResponseCode) {
		toSerialize["responseCode"] = o.ResponseCode
	}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableIssueCollectRequestResponse struct {
	value *IssueCollectRequestResponse
	isSet bool
}

func (v NullableIssueCollectRequestResponse) Get() *IssueCollectRequestResponse {
	return v.value
}

func (v *NullableIssueCollectRequestResponse) Set(val *IssueCollectRequestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableIssueCollectRequestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableIssueCollectRequestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIssueCollectRequestResponse(val *IssueCollectRequestResponse) *NullableIssueCollectRequestResponse {
	return &NullableIssueCollectRequestResponse{value: val, isSet: true}
}

func (v NullableIssueCollectRequestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIssueCollectRequestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


