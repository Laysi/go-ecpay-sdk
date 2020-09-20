# AioCheckOutCreditOnceOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redeem** | Pointer to [**RedeemEnum**](RedeemEnum.md) |  | [optional] 
**UnionPay** | Pointer to [**UnionPayEnum**](UnionPayEnum.md) |  | [optional] [default to UNIONPAYENUM_NOT_SPECIFIED]

## Methods

### NewAioCheckOutCreditOnceOption

`func NewAioCheckOutCreditOnceOption() *AioCheckOutCreditOnceOption`

NewAioCheckOutCreditOnceOption instantiates a new AioCheckOutCreditOnceOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutCreditOnceOptionWithDefaults

`func NewAioCheckOutCreditOnceOptionWithDefaults() *AioCheckOutCreditOnceOption`

NewAioCheckOutCreditOnceOptionWithDefaults instantiates a new AioCheckOutCreditOnceOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedeem

`func (o *AioCheckOutCreditOnceOption) GetRedeem() RedeemEnum`

GetRedeem returns the Redeem field if non-nil, zero value otherwise.

### GetRedeemOk

`func (o *AioCheckOutCreditOnceOption) GetRedeemOk() (*RedeemEnum, bool)`

GetRedeemOk returns a tuple with the Redeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeem

`func (o *AioCheckOutCreditOnceOption) SetRedeem(v RedeemEnum)`

SetRedeem sets Redeem field to given value.

### HasRedeem

`func (o *AioCheckOutCreditOnceOption) HasRedeem() bool`

HasRedeem returns a boolean if a field has been set.

### GetUnionPay

`func (o *AioCheckOutCreditOnceOption) GetUnionPay() UnionPayEnum`

GetUnionPay returns the UnionPay field if non-nil, zero value otherwise.

### GetUnionPayOk

`func (o *AioCheckOutCreditOnceOption) GetUnionPayOk() (*UnionPayEnum, bool)`

GetUnionPayOk returns a tuple with the UnionPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionPay

`func (o *AioCheckOutCreditOnceOption) SetUnionPay(v UnionPayEnum)`

SetUnionPay sets UnionPay field to given value.

### HasUnionPay

`func (o *AioCheckOutCreditOnceOption) HasUnionPay() bool`

HasUnionPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


