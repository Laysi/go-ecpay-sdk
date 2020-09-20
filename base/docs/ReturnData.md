# ReturnData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合  | 
**StoreID** | Pointer to **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
**RtnCode** | Pointer to **int32** | **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。  | 
**RtnMsg** | Pointer to **string** | **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded    | 
**TradeNo** | Pointer to **string** | **綠界的交易編號** 請保存綠界的交易編號與特店交易編號[MerchantTradeNo]的關連。  | 
**TradeAmt** | Pointer to **int32** | **交易金額**  | 
**PaymentDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **付款時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
**PaymentType** | Pointer to [**PaymentTypeEnum**](PaymentTypeEnum.md) |  | 
**PaymentTypeChargeFee** | Pointer to **int32** | **通路費**  | 
**TradeDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 
**SimulatePaid** | Pointer to [**SimulatePaidEnum**](SimulatePaidEnum.md) |  | 
**CustomField1** | Pointer to **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField2** | Pointer to **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField3** | Pointer to **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField4** | Pointer to **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField5** | Pointer to **string** | **自訂名稱欄位5**   提供合作廠商使用記錄用客製化使用欄位  | 

## Methods

### NewReturnData

`func NewReturnData(merchantID string, merchantTradeNo string, storeID string, rtnCode int32, rtnMsg string, tradeNo string, tradeAmt int32, paymentDate ECPayDateTime, paymentType PaymentTypeEnum, paymentTypeChargeFee int32, tradeDate ECPayDateTime, checkMacValue string, simulatePaid SimulatePaidEnum, customField1 string, customField2 string, customField3 string, customField4 string, customField5 string, ) *ReturnData`

NewReturnData instantiates a new ReturnData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReturnDataWithDefaults

`func NewReturnDataWithDefaults() *ReturnData`

NewReturnDataWithDefaults instantiates a new ReturnData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *ReturnData) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *ReturnData) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *ReturnData) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *ReturnData) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *ReturnData) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *ReturnData) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetStoreID

`func (o *ReturnData) GetStoreID() string`

GetStoreID returns the StoreID field if non-nil, zero value otherwise.

### GetStoreIDOk

`func (o *ReturnData) GetStoreIDOk() (*string, bool)`

GetStoreIDOk returns a tuple with the StoreID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreID

`func (o *ReturnData) SetStoreID(v string)`

SetStoreID sets StoreID field to given value.


### GetRtnCode

`func (o *ReturnData) GetRtnCode() int32`

GetRtnCode returns the RtnCode field if non-nil, zero value otherwise.

### GetRtnCodeOk

`func (o *ReturnData) GetRtnCodeOk() (*int32, bool)`

GetRtnCodeOk returns a tuple with the RtnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnCode

`func (o *ReturnData) SetRtnCode(v int32)`

SetRtnCode sets RtnCode field to given value.


### GetRtnMsg

`func (o *ReturnData) GetRtnMsg() string`

GetRtnMsg returns the RtnMsg field if non-nil, zero value otherwise.

### GetRtnMsgOk

`func (o *ReturnData) GetRtnMsgOk() (*string, bool)`

GetRtnMsgOk returns a tuple with the RtnMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnMsg

`func (o *ReturnData) SetRtnMsg(v string)`

SetRtnMsg sets RtnMsg field to given value.


### GetTradeNo

`func (o *ReturnData) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *ReturnData) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *ReturnData) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.


### GetTradeAmt

`func (o *ReturnData) GetTradeAmt() int32`

GetTradeAmt returns the TradeAmt field if non-nil, zero value otherwise.

### GetTradeAmtOk

`func (o *ReturnData) GetTradeAmtOk() (*int32, bool)`

GetTradeAmtOk returns a tuple with the TradeAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeAmt

`func (o *ReturnData) SetTradeAmt(v int32)`

SetTradeAmt sets TradeAmt field to given value.


### GetPaymentDate

`func (o *ReturnData) GetPaymentDate() ECPayDateTime`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *ReturnData) GetPaymentDateOk() (*ECPayDateTime, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *ReturnData) SetPaymentDate(v ECPayDateTime)`

SetPaymentDate sets PaymentDate field to given value.


### GetPaymentType

`func (o *ReturnData) GetPaymentType() PaymentTypeEnum`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *ReturnData) GetPaymentTypeOk() (*PaymentTypeEnum, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *ReturnData) SetPaymentType(v PaymentTypeEnum)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentTypeChargeFee

`func (o *ReturnData) GetPaymentTypeChargeFee() int32`

GetPaymentTypeChargeFee returns the PaymentTypeChargeFee field if non-nil, zero value otherwise.

### GetPaymentTypeChargeFeeOk

`func (o *ReturnData) GetPaymentTypeChargeFeeOk() (*int32, bool)`

GetPaymentTypeChargeFeeOk returns a tuple with the PaymentTypeChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeChargeFee

`func (o *ReturnData) SetPaymentTypeChargeFee(v int32)`

SetPaymentTypeChargeFee sets PaymentTypeChargeFee field to given value.


### GetTradeDate

`func (o *ReturnData) GetTradeDate() ECPayDateTime`

GetTradeDate returns the TradeDate field if non-nil, zero value otherwise.

### GetTradeDateOk

`func (o *ReturnData) GetTradeDateOk() (*ECPayDateTime, bool)`

GetTradeDateOk returns a tuple with the TradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDate

`func (o *ReturnData) SetTradeDate(v ECPayDateTime)`

SetTradeDate sets TradeDate field to given value.


### GetCheckMacValue

`func (o *ReturnData) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *ReturnData) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *ReturnData) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.


### GetSimulatePaid

`func (o *ReturnData) GetSimulatePaid() SimulatePaidEnum`

GetSimulatePaid returns the SimulatePaid field if non-nil, zero value otherwise.

### GetSimulatePaidOk

`func (o *ReturnData) GetSimulatePaidOk() (*SimulatePaidEnum, bool)`

GetSimulatePaidOk returns a tuple with the SimulatePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatePaid

`func (o *ReturnData) SetSimulatePaid(v SimulatePaidEnum)`

SetSimulatePaid sets SimulatePaid field to given value.


### GetCustomField1

`func (o *ReturnData) GetCustomField1() string`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *ReturnData) GetCustomField1Ok() (*string, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *ReturnData) SetCustomField1(v string)`

SetCustomField1 sets CustomField1 field to given value.


### GetCustomField2

`func (o *ReturnData) GetCustomField2() string`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *ReturnData) GetCustomField2Ok() (*string, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *ReturnData) SetCustomField2(v string)`

SetCustomField2 sets CustomField2 field to given value.


### GetCustomField3

`func (o *ReturnData) GetCustomField3() string`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *ReturnData) GetCustomField3Ok() (*string, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *ReturnData) SetCustomField3(v string)`

SetCustomField3 sets CustomField3 field to given value.


### GetCustomField4

`func (o *ReturnData) GetCustomField4() string`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *ReturnData) GetCustomField4Ok() (*string, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *ReturnData) SetCustomField4(v string)`

SetCustomField4 sets CustomField4 field to given value.


### GetCustomField5

`func (o *ReturnData) GetCustomField5() string`

GetCustomField5 returns the CustomField5 field if non-nil, zero value otherwise.

### GetCustomField5Ok

`func (o *ReturnData) GetCustomField5Ok() (*string, bool)`

GetCustomField5Ok returns a tuple with the CustomField5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField5

`func (o *ReturnData) SetCustomField5(v string)`

SetCustomField5 sets CustomField5 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


