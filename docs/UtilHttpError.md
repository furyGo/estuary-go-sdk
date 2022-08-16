# UtilHttpError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **int32** |  | [optional] 
**Details** | Pointer to **string** |  | [optional] 
**Reason** | Pointer to **string** |  | [optional] 

## Methods

### NewUtilHttpError

`func NewUtilHttpError() *UtilHttpError`

NewUtilHttpError instantiates a new UtilHttpError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUtilHttpErrorWithDefaults

`func NewUtilHttpErrorWithDefaults() *UtilHttpError`

NewUtilHttpErrorWithDefaults instantiates a new UtilHttpError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *UtilHttpError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UtilHttpError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UtilHttpError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *UtilHttpError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *UtilHttpError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UtilHttpError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UtilHttpError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *UtilHttpError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetReason

`func (o *UtilHttpError) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UtilHttpError) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UtilHttpError) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UtilHttpError) HasReason() bool`

HasReason returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


