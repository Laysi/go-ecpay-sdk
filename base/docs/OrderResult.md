# OrderResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號** 訂單產生時傳送給綠界的特店交易編號。英數字大小寫混合  | 
**StoreID** | Pointer to **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
**RtnCode** | Pointer to **int** | **交易狀態**   若回傳值為 1 時，為付款成功   其餘代碼皆為交易異常，請至廠商管理後台確認後再出貨。  | 
**RtnMsg** | Pointer to **string** | **交易訊息** Server POST 成功回傳:交易成功   Server POST 補送通知回傳:paid   Client POST 成功回傳:Succeeded    | 
**TradeNo** | Pointer to **string** | **綠界的交易編號** 請保存綠界的交易編號與特店交易編號[MerchantTradeNo]的關連。  | 
**TradeAmt** | Pointer to **int** | **交易金額**  | 
**PaymentDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **付款時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
**PaymentType** | Pointer to [**ReturnPaymentTypeEnum**](ReturnPaymentTypeEnum.md) |  | 
**PaymentTypeChargeFee** | Pointer to **int** | **通路費**  | 
**TradeDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 
**SimulatePaid** | Pointer to [**SimulatePaidEnum**](SimulatePaidEnum.md) |  | 
**CustomField1** | Pointer to **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField2** | Pointer to **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField3** | Pointer to **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位  | 
**CustomField4** | Pointer to **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位  | 

## Methods

### NewOrderResult

`func NewOrderResult(merchantID string, merchantTradeNo string, storeID string, rtnCode int, rtnMsg string, tradeNo string, tradeAmt int, paymentDate ECPayDateTime, paymentType ReturnPaymentTypeEnum, paymentTypeChargeFee int, tradeDate ECPayDateTime, checkMacValue string, simulatePaid SimulatePaidEnum, customField1 string, customField2 string, customField3 string, customField4 string, ) *OrderResult`

NewOrderResult instantiates a new OrderResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderResultWithDefaults

`func NewOrderResultWithDefaults() *OrderResult`

NewOrderResultWithDefaults instantiates a new OrderResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *OrderResult) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *OrderResult) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *OrderResult) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *OrderResult) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *OrderResult) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *OrderResult) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetStoreID

`func (o *OrderResult) GetStoreID() string`

GetStoreID returns the StoreID field if non-nil, zero value otherwise.

### GetStoreIDOk

`func (o *OrderResult) GetStoreIDOk() (*string, bool)`

GetStoreIDOk returns a tuple with the StoreID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreID

`func (o *OrderResult) SetStoreID(v string)`

SetStoreID sets StoreID field to given value.


### GetRtnCode

`func (o *OrderResult) GetRtnCode() int`

GetRtnCode returns the RtnCode field if non-nil, zero value otherwise.

### GetRtnCodeOk

`func (o *OrderResult) GetRtnCodeOk() (*int, bool)`

GetRtnCodeOk returns a tuple with the RtnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnCode

`func (o *OrderResult) SetRtnCode(v int)`

SetRtnCode sets RtnCode field to given value.


### GetRtnMsg

`func (o *OrderResult) GetRtnMsg() string`

GetRtnMsg returns the RtnMsg field if non-nil, zero value otherwise.

### GetRtnMsgOk

`func (o *OrderResult) GetRtnMsgOk() (*string, bool)`

GetRtnMsgOk returns a tuple with the RtnMsg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRtnMsg

`func (o *OrderResult) SetRtnMsg(v string)`

SetRtnMsg sets RtnMsg field to given value.


### GetTradeNo

`func (o *OrderResult) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *OrderResult) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *OrderResult) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.


### GetTradeAmt

`func (o *OrderResult) GetTradeAmt() int`

GetTradeAmt returns the TradeAmt field if non-nil, zero value otherwise.

### GetTradeAmtOk

`func (o *OrderResult) GetTradeAmtOk() (*int, bool)`

GetTradeAmtOk returns a tuple with the TradeAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeAmt

`func (o *OrderResult) SetTradeAmt(v int)`

SetTradeAmt sets TradeAmt field to given value.


### GetPaymentDate

`func (o *OrderResult) GetPaymentDate() ECPayDateTime`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *OrderResult) GetPaymentDateOk() (*ECPayDateTime, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *OrderResult) SetPaymentDate(v ECPayDateTime)`

SetPaymentDate sets PaymentDate field to given value.


### GetPaymentType

`func (o *OrderResult) GetPaymentType() ReturnPaymentTypeEnum`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *OrderResult) GetPaymentTypeOk() (*ReturnPaymentTypeEnum, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *OrderResult) SetPaymentType(v ReturnPaymentTypeEnum)`

SetPaymentType sets PaymentType field to given value.


### GetPaymentTypeChargeFee

`func (o *OrderResult) GetPaymentTypeChargeFee() int`

GetPaymentTypeChargeFee returns the PaymentTypeChargeFee field if non-nil, zero value otherwise.

### GetPaymentTypeChargeFeeOk

`func (o *OrderResult) GetPaymentTypeChargeFeeOk() (*int, bool)`

GetPaymentTypeChargeFeeOk returns a tuple with the PaymentTypeChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeChargeFee

`func (o *OrderResult) SetPaymentTypeChargeFee(v int)`

SetPaymentTypeChargeFee sets PaymentTypeChargeFee field to given value.


### GetTradeDate

`func (o *OrderResult) GetTradeDate() ECPayDateTime`

GetTradeDate returns the TradeDate field if non-nil, zero value otherwise.

### GetTradeDateOk

`func (o *OrderResult) GetTradeDateOk() (*ECPayDateTime, bool)`

GetTradeDateOk returns a tuple with the TradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDate

`func (o *OrderResult) SetTradeDate(v ECPayDateTime)`

SetTradeDate sets TradeDate field to given value.


### GetCheckMacValue

`func (o *OrderResult) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *OrderResult) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *OrderResult) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.


### GetSimulatePaid

`func (o *OrderResult) GetSimulatePaid() SimulatePaidEnum`

GetSimulatePaid returns the SimulatePaid field if non-nil, zero value otherwise.

### GetSimulatePaidOk

`func (o *OrderResult) GetSimulatePaidOk() (*SimulatePaidEnum, bool)`

GetSimulatePaidOk returns a tuple with the SimulatePaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimulatePaid

`func (o *OrderResult) SetSimulatePaid(v SimulatePaidEnum)`

SetSimulatePaid sets SimulatePaid field to given value.


### GetCustomField1

`func (o *OrderResult) GetCustomField1() string`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *OrderResult) GetCustomField1Ok() (*string, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *OrderResult) SetCustomField1(v string)`

SetCustomField1 sets CustomField1 field to given value.


### GetCustomField2

`func (o *OrderResult) GetCustomField2() string`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *OrderResult) GetCustomField2Ok() (*string, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *OrderResult) SetCustomField2(v string)`

SetCustomField2 sets CustomField2 field to given value.


### GetCustomField3

`func (o *OrderResult) GetCustomField3() string`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *OrderResult) GetCustomField3Ok() (*string, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *OrderResult) SetCustomField3(v string)`

SetCustomField3 sets CustomField3 field to given value.


### GetCustomField4

`func (o *OrderResult) GetCustomField4() string`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *OrderResult) GetCustomField4Ok() (*string, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *OrderResult) SetCustomField4(v string)`

SetCustomField4 sets CustomField4 field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


