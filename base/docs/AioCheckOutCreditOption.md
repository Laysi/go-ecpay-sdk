# AioCheckOutCreditOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BindingCard** | Pointer to [**BindingCardEnum**](BindingCardEnum.md) |  | [optional] 
**MerchantMemberID** | Pointer to **string** | **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 &#x60;MerchantID&#x60; + &#x60;廠商會員編號&#x60;)  | [optional] 

## Methods

### NewAioCheckOutCreditOption

`func NewAioCheckOutCreditOption() *AioCheckOutCreditOption`

NewAioCheckOutCreditOption instantiates a new AioCheckOutCreditOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutCreditOptionWithDefaults

`func NewAioCheckOutCreditOptionWithDefaults() *AioCheckOutCreditOption`

NewAioCheckOutCreditOptionWithDefaults instantiates a new AioCheckOutCreditOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBindingCard

`func (o *AioCheckOutCreditOption) GetBindingCard() BindingCardEnum`

GetBindingCard returns the BindingCard field if non-nil, zero value otherwise.

### GetBindingCardOk

`func (o *AioCheckOutCreditOption) GetBindingCardOk() (*BindingCardEnum, bool)`

GetBindingCardOk returns a tuple with the BindingCard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingCard

`func (o *AioCheckOutCreditOption) SetBindingCard(v BindingCardEnum)`

SetBindingCard sets BindingCard field to given value.

### HasBindingCard

`func (o *AioCheckOutCreditOption) HasBindingCard() bool`

HasBindingCard returns a boolean if a field has been set.

### GetMerchantMemberID

`func (o *AioCheckOutCreditOption) GetMerchantMemberID() string`

GetMerchantMemberID returns the MerchantMemberID field if non-nil, zero value otherwise.

### GetMerchantMemberIDOk

`func (o *AioCheckOutCreditOption) GetMerchantMemberIDOk() (*string, bool)`

GetMerchantMemberIDOk returns a tuple with the MerchantMemberID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantMemberID

`func (o *AioCheckOutCreditOption) SetMerchantMemberID(v string)`

SetMerchantMemberID sets MerchantMemberID field to given value.

### HasMerchantMemberID

`func (o *AioCheckOutCreditOption) HasMerchantMemberID() bool`

HasMerchantMemberID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


