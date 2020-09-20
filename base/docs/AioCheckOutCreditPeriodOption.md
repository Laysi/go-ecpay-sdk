# AioCheckOutCreditPeriodOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PeriodAmount** | Pointer to **int** | **每次授權金額**   每次要授權(扣款)的金額。   注意事項：   綠界會依此次授權金額&#x60;PeriodAmount&#x60;所設定的金額做為之後固定授權的金額。   交易金額&#x60;TotalAmount&#x60;設定金額必須和授權金額&#x60;PeriodAmount&#x60;相同。   請帶整數，不可有小數點。僅限新台幣。  | 
**PeriodType** | Pointer to [**CreditPeriodTypeEnum**](CreditPeriodTypeEnum.md) |  | 
**Frequency** | Pointer to **int** | **執行頻率**   此參數用來定義多久要執行一次   注意事項：   至少要大於等於 1 次以上。   當 &#x60;PeriodType&#x60; 設為 &#x60;D&#x60; 時，最多可設 &#x60;365&#x60; 次。    當 &#x60;PeriodType&#x60; 設為 &#x60;M&#x60; 時，最多可設 &#x60;12&#x60; 次。     當 &#x60;PeriodType&#x60; 設為 &#x60;Y&#x60; 時，最多可設 &#x60;1&#x60; 次。  | 
**ExecTimes** | Pointer to **int** | **執行次數**   總共要執行幾次。   注意事項：   至少要大於 1 次以上。   當 &#x60;PeriodType&#x60; 設為 &#x60;D&#x60; 時，最多可設 &#x60;999&#x60;次。   當 &#x60;PeriodType&#x60; 設為 &#x60;M&#x60; 時，最多可設 &#x60;99&#x60;次。   當 &#x60;PeriodType&#x60; 設為 &#x60;Y&#x60; 時，最多可設 &#x60;9&#x60; 次。    例 1：   當信用卡定期定額扣款為每個月扣 1 次500 元，總共要扣 12次   &#x60;TotalAmount&#x60;參數請帶 &#x60;500&#x60; &#x60;PeriodAmount&#x60;&#x3D;&#x60;500&#x60;   &#x60;PeriodType&#x60;&#x3D;&#x60;M&#x60;   &#x60;Frequency&#x60;&#x3D;&#x60;1&#x60;   &#x60;ExecTimes&#x60;&#x3D;&#x60;12&#x60;    例 2：   當信用卡定期定額扣款為 6000 元，每 6 個月扣 1 次，總共要扣 2 次時    交易金額&#x60;TotalAmount&#x60;參數請帶 &#x60;6000&#x60;   &#x60;PeriodType&#x60;&#x3D;&#x60;M&#x60;   &#x60;Frequency&#x60;&#x3D;&#x60;6&#x60;   &#x60;ExecTimes&#x60;&#x3D;&#x60;2&#x60;    | 
**PeriodReturnURL** | Pointer to **string** | **定期定額的執行結果回應URL**   若交易是信用卡定期定額的方式，則每次執行授權完，會將授權結果回傳到這個設定的 URL。   回覆內容請參考。    | [optional] 

## Methods

### NewAioCheckOutCreditPeriodOption

`func NewAioCheckOutCreditPeriodOption(periodAmount int, periodType CreditPeriodTypeEnum, frequency int, execTimes int, ) *AioCheckOutCreditPeriodOption`

NewAioCheckOutCreditPeriodOption instantiates a new AioCheckOutCreditPeriodOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutCreditPeriodOptionWithDefaults

`func NewAioCheckOutCreditPeriodOptionWithDefaults() *AioCheckOutCreditPeriodOption`

NewAioCheckOutCreditPeriodOptionWithDefaults instantiates a new AioCheckOutCreditPeriodOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPeriodAmount

`func (o *AioCheckOutCreditPeriodOption) GetPeriodAmount() int`

GetPeriodAmount returns the PeriodAmount field if non-nil, zero value otherwise.

### GetPeriodAmountOk

`func (o *AioCheckOutCreditPeriodOption) GetPeriodAmountOk() (*int, bool)`

GetPeriodAmountOk returns a tuple with the PeriodAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodAmount

`func (o *AioCheckOutCreditPeriodOption) SetPeriodAmount(v int)`

SetPeriodAmount sets PeriodAmount field to given value.


### GetPeriodType

`func (o *AioCheckOutCreditPeriodOption) GetPeriodType() CreditPeriodTypeEnum`

GetPeriodType returns the PeriodType field if non-nil, zero value otherwise.

### GetPeriodTypeOk

`func (o *AioCheckOutCreditPeriodOption) GetPeriodTypeOk() (*CreditPeriodTypeEnum, bool)`

GetPeriodTypeOk returns a tuple with the PeriodType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodType

`func (o *AioCheckOutCreditPeriodOption) SetPeriodType(v CreditPeriodTypeEnum)`

SetPeriodType sets PeriodType field to given value.


### GetFrequency

`func (o *AioCheckOutCreditPeriodOption) GetFrequency() int`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *AioCheckOutCreditPeriodOption) GetFrequencyOk() (*int, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *AioCheckOutCreditPeriodOption) SetFrequency(v int)`

SetFrequency sets Frequency field to given value.


### GetExecTimes

`func (o *AioCheckOutCreditPeriodOption) GetExecTimes() int`

GetExecTimes returns the ExecTimes field if non-nil, zero value otherwise.

### GetExecTimesOk

`func (o *AioCheckOutCreditPeriodOption) GetExecTimesOk() (*int, bool)`

GetExecTimesOk returns a tuple with the ExecTimes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecTimes

`func (o *AioCheckOutCreditPeriodOption) SetExecTimes(v int)`

SetExecTimes sets ExecTimes field to given value.


### GetPeriodReturnURL

`func (o *AioCheckOutCreditPeriodOption) GetPeriodReturnURL() string`

GetPeriodReturnURL returns the PeriodReturnURL field if non-nil, zero value otherwise.

### GetPeriodReturnURLOk

`func (o *AioCheckOutCreditPeriodOption) GetPeriodReturnURLOk() (*string, bool)`

GetPeriodReturnURLOk returns a tuple with the PeriodReturnURL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodReturnURL

`func (o *AioCheckOutCreditPeriodOption) SetPeriodReturnURL(v string)`

SetPeriodReturnURL sets PeriodReturnURL field to given value.

### HasPeriodReturnURL

`func (o *AioCheckOutCreditPeriodOption) HasPeriodReturnURL() bool`

HasPeriodReturnURL returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


