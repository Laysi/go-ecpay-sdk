# AioCheckOutCreditOnetimeOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Redeem** | Pointer to [**RedeemEnum**](RedeemEnum.md) |  | [optional] 
**UnionPay** | Pointer to [**UnionPayEnum**](UnionPayEnum.md) |  | [optional] [default to UNIONPAYENUM_NOT_SPECIFIED]

## Methods

### NewAioCheckOutCreditOnetimeOption

`func NewAioCheckOutCreditOnetimeOption() *AioCheckOutCreditOnetimeOption`

NewAioCheckOutCreditOnetimeOption instantiates a new AioCheckOutCreditOnetimeOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutCreditOnetimeOptionWithDefaults

`func NewAioCheckOutCreditOnetimeOptionWithDefaults() *AioCheckOutCreditOnetimeOption`

NewAioCheckOutCreditOnetimeOptionWithDefaults instantiates a new AioCheckOutCreditOnetimeOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRedeem

`func (o *AioCheckOutCreditOnetimeOption) GetRedeem() RedeemEnum`

GetRedeem returns the Redeem field if non-nil, zero value otherwise.

### GetRedeemOk

`func (o *AioCheckOutCreditOnetimeOption) GetRedeemOk() (*RedeemEnum, bool)`

GetRedeemOk returns a tuple with the Redeem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRedeem

`func (o *AioCheckOutCreditOnetimeOption) SetRedeem(v RedeemEnum)`

SetRedeem sets Redeem field to given value.

### HasRedeem

`func (o *AioCheckOutCreditOnetimeOption) HasRedeem() bool`

HasRedeem returns a boolean if a field has been set.

### GetUnionPay

`func (o *AioCheckOutCreditOnetimeOption) GetUnionPay() UnionPayEnum`

GetUnionPay returns the UnionPay field if non-nil, zero value otherwise.

### GetUnionPayOk

`func (o *AioCheckOutCreditOnetimeOption) GetUnionPayOk() (*UnionPayEnum, bool)`

GetUnionPayOk returns a tuple with the UnionPay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnionPay

`func (o *AioCheckOutCreditOnetimeOption) SetUnionPay(v UnionPayEnum)`

SetUnionPay sets UnionPay field to given value.

### HasUnionPay

`func (o *AioCheckOutCreditOnetimeOption) HasUnionPay() bool`

HasUnionPay returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


