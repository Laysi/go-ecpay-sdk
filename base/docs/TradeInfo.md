# TradeInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantID** | Pointer to **string** | **特店編號(由綠界提供)**  | 
**MerchantTradeNo** | Pointer to **string** | **特店交易編號(由特店提供)** 訂單產生時傳送給綠界的特店交易編號。  | 
**StoreID** | Pointer to **string** | **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。  | 
**TradeNo** | Pointer to **string** | **綠界的交易編號** 首次授權所產生的綠界交易編號  | 
**TradeAmt** | Pointer to **int** | **交易金額**  | 
**PaymentDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **付款時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
**PaymentType** | Pointer to [**ReturnPaymentTypeEnum**](ReturnPaymentTypeEnum.md) |  | 
**HandlingCharge** | Pointer to **int** | **手續費合計**  履約結束後才會計算，未計算前為 0  | 
**PaymentTypeChargeFee** | Pointer to **float32** | **通路費**  | 
**TradeDate** | Pointer to [**ECPayDateTime**](ECPayDateTime.md) | **訂單成立時間** 格式為 yyyy/MM/dd HH:mm:ss  | 
**TradeStatus** | Pointer to **string** | **交易狀態**    回傳值：   若為 0 時，代表交易訂單成立未付款   若為 1 時，代表交易訂單成立已付款   若為 10200095 時，代表消費者未選擇付款方式，故交易失敗。    | 
**ItemName** | Pointer to **string** | **商品名稱**     | 
**CustomField1** | Pointer to **string** | **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位    | 
**CustomField2** | Pointer to **string** | **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位    | 
**CustomField3** | Pointer to **string** | **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位    | 
**CustomField4** | Pointer to **string** | **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位    | 
**CheckMacValue** | Pointer to **string** | **檢查碼** 特店必須檢查檢查碼&#x60;CheckMacValue&#x60;來驗證，請參考附錄檢查碼機制。  | 

## Methods

### NewTradeInfo

`func NewTradeInfo(merchantID string, merchantTradeNo string, storeID string, tradeNo string, tradeAmt int, paymentDate ECPayDateTime, paymentType ReturnPaymentTypeEnum, handlingCharge int, paymentTypeChargeFee float32, tradeDate ECPayDateTime, tradeStatus string, itemName string, customField1 string, customField2 string, customField3 string, customField4 string, checkMacValue string, ) *TradeInfo`

NewTradeInfo instantiates a new TradeInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTradeInfoWithDefaults

`func NewTradeInfoWithDefaults() *TradeInfo`

NewTradeInfoWithDefaults instantiates a new TradeInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantID

`func (o *TradeInfo) GetMerchantID() string`

GetMerchantID returns the MerchantID field if non-nil, zero value otherwise.

### GetMerchantIDOk

`func (o *TradeInfo) GetMerchantIDOk() (*string, bool)`

GetMerchantIDOk returns a tuple with the MerchantID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantID

`func (o *TradeInfo) SetMerchantID(v string)`

SetMerchantID sets MerchantID field to given value.


### GetMerchantTradeNo

`func (o *TradeInfo) GetMerchantTradeNo() string`

GetMerchantTradeNo returns the MerchantTradeNo field if non-nil, zero value otherwise.

### GetMerchantTradeNoOk

`func (o *TradeInfo) GetMerchantTradeNoOk() (*string, bool)`

GetMerchantTradeNoOk returns a tuple with the MerchantTradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantTradeNo

`func (o *TradeInfo) SetMerchantTradeNo(v string)`

SetMerchantTradeNo sets MerchantTradeNo field to given value.


### GetStoreID

`func (o *TradeInfo) GetStoreID() string`

GetStoreID returns the StoreID field if non-nil, zero value otherwise.

### GetStoreIDOk

`func (o *TradeInfo) GetStoreIDOk() (*string, bool)`

GetStoreIDOk returns a tuple with the StoreID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreID

`func (o *TradeInfo) SetStoreID(v string)`

SetStoreID sets StoreID field to given value.


### GetTradeNo

`func (o *TradeInfo) GetTradeNo() string`

GetTradeNo returns the TradeNo field if non-nil, zero value otherwise.

### GetTradeNoOk

`func (o *TradeInfo) GetTradeNoOk() (*string, bool)`

GetTradeNoOk returns a tuple with the TradeNo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeNo

`func (o *TradeInfo) SetTradeNo(v string)`

SetTradeNo sets TradeNo field to given value.


### GetTradeAmt

`func (o *TradeInfo) GetTradeAmt() int`

GetTradeAmt returns the TradeAmt field if non-nil, zero value otherwise.

### GetTradeAmtOk

`func (o *TradeInfo) GetTradeAmtOk() (*int, bool)`

GetTradeAmtOk returns a tuple with the TradeAmt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeAmt

`func (o *TradeInfo) SetTradeAmt(v int)`

SetTradeAmt sets TradeAmt field to given value.


### GetPaymentDate

`func (o *TradeInfo) GetPaymentDate() ECPayDateTime`

GetPaymentDate returns the PaymentDate field if non-nil, zero value otherwise.

### GetPaymentDateOk

`func (o *TradeInfo) GetPaymentDateOk() (*ECPayDateTime, bool)`

GetPaymentDateOk returns a tuple with the PaymentDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentDate

`func (o *TradeInfo) SetPaymentDate(v ECPayDateTime)`

SetPaymentDate sets PaymentDate field to given value.


### GetPaymentType

`func (o *TradeInfo) GetPaymentType() ReturnPaymentTypeEnum`

GetPaymentType returns the PaymentType field if non-nil, zero value otherwise.

### GetPaymentTypeOk

`func (o *TradeInfo) GetPaymentTypeOk() (*ReturnPaymentTypeEnum, bool)`

GetPaymentTypeOk returns a tuple with the PaymentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentType

`func (o *TradeInfo) SetPaymentType(v ReturnPaymentTypeEnum)`

SetPaymentType sets PaymentType field to given value.


### GetHandlingCharge

`func (o *TradeInfo) GetHandlingCharge() int`

GetHandlingCharge returns the HandlingCharge field if non-nil, zero value otherwise.

### GetHandlingChargeOk

`func (o *TradeInfo) GetHandlingChargeOk() (*int, bool)`

GetHandlingChargeOk returns a tuple with the HandlingCharge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingCharge

`func (o *TradeInfo) SetHandlingCharge(v int)`

SetHandlingCharge sets HandlingCharge field to given value.


### GetPaymentTypeChargeFee

`func (o *TradeInfo) GetPaymentTypeChargeFee() float32`

GetPaymentTypeChargeFee returns the PaymentTypeChargeFee field if non-nil, zero value otherwise.

### GetPaymentTypeChargeFeeOk

`func (o *TradeInfo) GetPaymentTypeChargeFeeOk() (*float32, bool)`

GetPaymentTypeChargeFeeOk returns a tuple with the PaymentTypeChargeFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentTypeChargeFee

`func (o *TradeInfo) SetPaymentTypeChargeFee(v float32)`

SetPaymentTypeChargeFee sets PaymentTypeChargeFee field to given value.


### GetTradeDate

`func (o *TradeInfo) GetTradeDate() ECPayDateTime`

GetTradeDate returns the TradeDate field if non-nil, zero value otherwise.

### GetTradeDateOk

`func (o *TradeInfo) GetTradeDateOk() (*ECPayDateTime, bool)`

GetTradeDateOk returns a tuple with the TradeDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeDate

`func (o *TradeInfo) SetTradeDate(v ECPayDateTime)`

SetTradeDate sets TradeDate field to given value.


### GetTradeStatus

`func (o *TradeInfo) GetTradeStatus() string`

GetTradeStatus returns the TradeStatus field if non-nil, zero value otherwise.

### GetTradeStatusOk

`func (o *TradeInfo) GetTradeStatusOk() (*string, bool)`

GetTradeStatusOk returns a tuple with the TradeStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradeStatus

`func (o *TradeInfo) SetTradeStatus(v string)`

SetTradeStatus sets TradeStatus field to given value.


### GetItemName

`func (o *TradeInfo) GetItemName() string`

GetItemName returns the ItemName field if non-nil, zero value otherwise.

### GetItemNameOk

`func (o *TradeInfo) GetItemNameOk() (*string, bool)`

GetItemNameOk returns a tuple with the ItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemName

`func (o *TradeInfo) SetItemName(v string)`

SetItemName sets ItemName field to given value.


### GetCustomField1

`func (o *TradeInfo) GetCustomField1() string`

GetCustomField1 returns the CustomField1 field if non-nil, zero value otherwise.

### GetCustomField1Ok

`func (o *TradeInfo) GetCustomField1Ok() (*string, bool)`

GetCustomField1Ok returns a tuple with the CustomField1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField1

`func (o *TradeInfo) SetCustomField1(v string)`

SetCustomField1 sets CustomField1 field to given value.


### GetCustomField2

`func (o *TradeInfo) GetCustomField2() string`

GetCustomField2 returns the CustomField2 field if non-nil, zero value otherwise.

### GetCustomField2Ok

`func (o *TradeInfo) GetCustomField2Ok() (*string, bool)`

GetCustomField2Ok returns a tuple with the CustomField2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField2

`func (o *TradeInfo) SetCustomField2(v string)`

SetCustomField2 sets CustomField2 field to given value.


### GetCustomField3

`func (o *TradeInfo) GetCustomField3() string`

GetCustomField3 returns the CustomField3 field if non-nil, zero value otherwise.

### GetCustomField3Ok

`func (o *TradeInfo) GetCustomField3Ok() (*string, bool)`

GetCustomField3Ok returns a tuple with the CustomField3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField3

`func (o *TradeInfo) SetCustomField3(v string)`

SetCustomField3 sets CustomField3 field to given value.


### GetCustomField4

`func (o *TradeInfo) GetCustomField4() string`

GetCustomField4 returns the CustomField4 field if non-nil, zero value otherwise.

### GetCustomField4Ok

`func (o *TradeInfo) GetCustomField4Ok() (*string, bool)`

GetCustomField4Ok returns a tuple with the CustomField4 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomField4

`func (o *TradeInfo) SetCustomField4(v string)`

SetCustomField4 sets CustomField4 field to given value.


### GetCheckMacValue

`func (o *TradeInfo) GetCheckMacValue() string`

GetCheckMacValue returns the CheckMacValue field if non-nil, zero value otherwise.

### GetCheckMacValueOk

`func (o *TradeInfo) GetCheckMacValueOk() (*string, bool)`

GetCheckMacValueOk returns a tuple with the CheckMacValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckMacValue

`func (o *TradeInfo) SetCheckMacValue(v string)`

SetCheckMacValue sets CheckMacValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


