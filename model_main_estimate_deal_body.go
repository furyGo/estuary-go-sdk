/*
Estuary API

This is the API for the Estuary application.

API version: 0.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// MainEstimateDealBody struct for MainEstimateDealBody
type MainEstimateDealBody struct {
	DurationBlks *int32 `json:"durationBlks,omitempty"`
	Replication *int32 `json:"replication,omitempty"`
	Size *int32 `json:"size,omitempty"`
	Verified *bool `json:"verified,omitempty"`
}

// NewMainEstimateDealBody instantiates a new MainEstimateDealBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMainEstimateDealBody() *MainEstimateDealBody {
	this := MainEstimateDealBody{}
	return &this
}

// NewMainEstimateDealBodyWithDefaults instantiates a new MainEstimateDealBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMainEstimateDealBodyWithDefaults() *MainEstimateDealBody {
	this := MainEstimateDealBody{}
	return &this
}

// GetDurationBlks returns the DurationBlks field value if set, zero value otherwise.
func (o *MainEstimateDealBody) GetDurationBlks() int32 {
	if o == nil || o.DurationBlks == nil {
		var ret int32
		return ret
	}
	return *o.DurationBlks
}

// GetDurationBlksOk returns a tuple with the DurationBlks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainEstimateDealBody) GetDurationBlksOk() (*int32, bool) {
	if o == nil || o.DurationBlks == nil {
		return nil, false
	}
	return o.DurationBlks, true
}

// HasDurationBlks returns a boolean if a field has been set.
func (o *MainEstimateDealBody) HasDurationBlks() bool {
	if o != nil && o.DurationBlks != nil {
		return true
	}

	return false
}

// SetDurationBlks gets a reference to the given int32 and assigns it to the DurationBlks field.
func (o *MainEstimateDealBody) SetDurationBlks(v int32) {
	o.DurationBlks = &v
}

// GetReplication returns the Replication field value if set, zero value otherwise.
func (o *MainEstimateDealBody) GetReplication() int32 {
	if o == nil || o.Replication == nil {
		var ret int32
		return ret
	}
	return *o.Replication
}

// GetReplicationOk returns a tuple with the Replication field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainEstimateDealBody) GetReplicationOk() (*int32, bool) {
	if o == nil || o.Replication == nil {
		return nil, false
	}
	return o.Replication, true
}

// HasReplication returns a boolean if a field has been set.
func (o *MainEstimateDealBody) HasReplication() bool {
	if o != nil && o.Replication != nil {
		return true
	}

	return false
}

// SetReplication gets a reference to the given int32 and assigns it to the Replication field.
func (o *MainEstimateDealBody) SetReplication(v int32) {
	o.Replication = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *MainEstimateDealBody) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainEstimateDealBody) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *MainEstimateDealBody) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *MainEstimateDealBody) SetSize(v int32) {
	o.Size = &v
}

// GetVerified returns the Verified field value if set, zero value otherwise.
func (o *MainEstimateDealBody) GetVerified() bool {
	if o == nil || o.Verified == nil {
		var ret bool
		return ret
	}
	return *o.Verified
}

// GetVerifiedOk returns a tuple with the Verified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MainEstimateDealBody) GetVerifiedOk() (*bool, bool) {
	if o == nil || o.Verified == nil {
		return nil, false
	}
	return o.Verified, true
}

// HasVerified returns a boolean if a field has been set.
func (o *MainEstimateDealBody) HasVerified() bool {
	if o != nil && o.Verified != nil {
		return true
	}

	return false
}

// SetVerified gets a reference to the given bool and assigns it to the Verified field.
func (o *MainEstimateDealBody) SetVerified(v bool) {
	o.Verified = &v
}

func (o MainEstimateDealBody) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DurationBlks != nil {
		toSerialize["durationBlks"] = o.DurationBlks
	}
	if o.Replication != nil {
		toSerialize["replication"] = o.Replication
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}
	if o.Verified != nil {
		toSerialize["verified"] = o.Verified
	}
	return json.Marshal(toSerialize)
}

type NullableMainEstimateDealBody struct {
	value *MainEstimateDealBody
	isSet bool
}

func (v NullableMainEstimateDealBody) Get() *MainEstimateDealBody {
	return v.value
}

func (v *NullableMainEstimateDealBody) Set(val *MainEstimateDealBody) {
	v.value = val
	v.isSet = true
}

func (v NullableMainEstimateDealBody) IsSet() bool {
	return v.isSet
}

func (v *NullableMainEstimateDealBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMainEstimateDealBody(val *MainEstimateDealBody) *NullableMainEstimateDealBody {
	return &NullableMainEstimateDealBody{value: val, isSet: true}
}

func (v NullableMainEstimateDealBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMainEstimateDealBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

