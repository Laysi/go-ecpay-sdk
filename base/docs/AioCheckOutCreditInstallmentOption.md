# AioCheckOutCreditInstallmentOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditInstallment** | Pointer to **string** | **刷卡分期期數**    提供刷卡分期期數   信用卡分期可用參數為:3,6,12,18,24   注意事項：   使用的期數必須先透過申請開通後方能使用，並以申請開通的期數為主。    | 

## Methods

### NewAioCheckOutCreditInstallmentOption

`func NewAioCheckOutCreditInstallmentOption(creditInstallment string, ) *AioCheckOutCreditInstallmentOption`

NewAioCheckOutCreditInstallmentOption instantiates a new AioCheckOutCreditInstallmentOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutCreditInstallmentOptionWithDefaults

`func NewAioCheckOutCreditInstallmentOptionWithDefaults() *AioCheckOutCreditInstallmentOption`

NewAioCheckOutCreditInstallmentOptionWithDefaults instantiates a new AioCheckOutCreditInstallmentOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditInstallment

`func (o *AioCheckOutCreditInstallmentOption) GetCreditInstallment() string`

GetCreditInstallment returns the CreditInstallment field if non-nil, zero value otherwise.

### GetCreditInstallmentOk

`func (o *AioCheckOutCreditInstallmentOption) GetCreditInstallmentOk() (*string, bool)`

GetCreditInstallmentOk returns a tuple with the CreditInstallment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditInstallment

`func (o *AioCheckOutCreditInstallmentOption) SetCreditInstallment(v string)`

SetCreditInstallment sets CreditInstallment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


