# AioCheckOutInvoiceOption

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RelateNumber** | Pointer to **string** | **特店自訂編號** 此為特店自訂編號，編號均為唯一值不可重複使用。  | 
**CustomerID** | Pointer to **string** | **客戶編號** 該參數有值時，僅接受『英文、數字、下底線』等字元。  | [optional] 
**CustomerIdentifier** | Pointer to **string** | **統一編號** 該參數有值時，請帶固定長度為數字 8 碼。  | [optional] 
**CustomerName** | Pointer to **string** | **客戶名稱**    當列印註記&#x60;Print&#x60;為 &#x60;1&#x60;(列印)時，則該參數必須有值。   該參數有值時，僅接受『中、英文及數字』等字元。   請將參數值做 UrlEncode 方式編碼。    | [optional] 
**CustomerAddr** | Pointer to **string** | **客戶地址**    當列印註記&#x60;Print&#x60;為 &#x60;1&#x60;(列印)時，則該參數必須有值。   當該參數有值時，請注意特殊字元轉換 。    請將參數值做 UrlEncode 方式編碼。  | [optional] 
**CustomerPhone** | Pointer to **string** | **客戶手機號碼**   當客戶電子信箱&#x60;CustomerEmail&#x60;為空字串時，則該參數必須有值。   當該參數有值時，則格式為數字。   注意事項：   請填手機號碼，不能填市話因為要收簡訊通知用    | [optional] 
**CustomerEmail** | Pointer to **string** | **客戶電子信箱**   當客戶手機號碼&#x60;CustomerPhone&#x60;為空字串時，則該參數必須有值。   當該參數有值時，則格式需符合 EMAIL格式。   請將參數值做 UrlEncode 方式編碼。    | [optional] 
**ClearanceMark** | Pointer to [**ClearanceMarkEnum**](ClearanceMarkEnum.md) |  | [optional] 
**TaxType** | Pointer to [**TaxTypeEnum**](TaxTypeEnum.md) |  | 
**CarruerType** | Pointer to [**CarruerTypeEnum**](CarruerTypeEnum.md) |  | [optional] 
**CarruerNum** | Pointer to **string** | **載具編號**   1. 當載具類別 &#x60;CarruerType&#x60;&#x3D;&#x60;&#x60;無載具)，請帶空字串。   2. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;1&#x60;(綠界科技電子發票載具)時，請帶空字串，系統會自動帶入值，為合作特店載具統一編號+自訂編號(RelateNumber)。   3. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;2&#x60;(買受人之自然人憑證)時，則請帶固定長度為16且格式 為2碼大寫英文字母加上14碼數字。   4. 當載具類別&#x60;CarruerType&#x60;&#x3D;&#x60;3&#x60;(買受人之手機條碼)時，則請帶固定長度為 8且格式為 1 碼斜線「/」加上由 7 碼數字及大寫英文字母及+-.符號組成。    注意事項：   1. 若手機條碼中有加號，可能在介接驗證時 發生錯誤，請將加號改為空白字元，產生 驗證碼。   2. 英文、數字、符號僅接受半形字   3. 若載具編號為手機條碼載具時，請先呼叫B2C電子發票介接技術文件手機條碼載驗證ＡＰＩ進行檢核    | [optional] 
**Donation** | Pointer to [**InvoiceDonationEunm**](InvoiceDonationEunm.md) |  | 
**LoveCode** | Pointer to **string** | **捐贈碼**   消費者選擇捐贈發票則於此欄位須填入受贈單位之捐贈碼。   1. 若捐贈註記 &#x60;Donation&#x60;&#x3D; &#x60;1&#x60; (捐贈)時，此欄位須有值。   2. 捐贈碼以阿拉伯數字為限，最少三碼，最多七碼。內容定位採「文字格式」，首位可以為零。     | [optional] 
**Print** | Pointer to [**InvoicePrintEnum**](InvoicePrintEnum.md) |  | 
**InvoiceItemName** | Pointer to **string** | **商品名稱**   預設不可為空字串且格式為 名稱 1 | 名稱 2 | 名稱 3 | … | 名稱 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。   將參數值以 UrlEncode 方式編碼。    | 
**InvoiceItemCount** | Pointer to **string** | **商品數量**   預設不可為空字串且格式為 數量 1 | 數量 2 | 數量 3 | … | 數量 n，當含有二筆或以上的商品名稱時，則以「|」符號區隔。  | 
**InvoiceItemWord** | Pointer to **string** | **商品單位**   商品單位若超過二筆以上請以「|」符號區隔單位最大長度為 6 碼。   請將參數做 UrlEncode 方式編碼。  | 
**InvoiceItemPrice** | Pointer to **string** | **商品價格**   預設不可為空字串且格式為 價格 1 | 價格 2 | 價格 3 | … | 價格 n，當含有二筆或以上的商品價格時，則以「|」符號區隔。  | 
**InvoiceItemTaxType** | Pointer to **string** | **商品課稅別**   1：應稅   2：零稅率   3：免稅   注意事項：   1. 預設為空字串，當課稅類別 [TaxType] &#x3D; 9 時，此欄位不可為空。   2. 格式為課稅 類別 1 | 課稅類別 2 | 課稅類別 3 | … | 課稅類別 n。當含有二筆或以上的商品課稅類別時，則以「|」符號區隔。   3. 課稅類別為混合稅率時，需含二筆或 以 上 的 商 品 課 稅   別[InvoiceItemTaxType]，且至少需有一筆商品課稅別為應稅及至少需有一筆商品課稅別為免稅或零稅率，即混稅發票只能 1.應稅+免稅 2.應稅+零稅率，免稅和零稅率發票不能同時開立。  | [optional] 
**InvoiceRemark** | Pointer to **string** | **備註** 當該參數有值時，請將參數值做UrlEncode 方式編碼。  | [optional] 
**DelayDay** | Pointer to **int** | **延遲天數**   本參數值請帶 0~15(天)，當天數為 0 時，則付款完成後立即開立發票。  | 
**InvType** | Pointer to **string** | **字軌類別**   若為一般稅額時，請帶 07。   預設值：07    | 

## Methods

### NewAioCheckOutInvoiceOption

`func NewAioCheckOutInvoiceOption(relateNumber string, taxType TaxTypeEnum, donation InvoiceDonationEunm, print InvoicePrintEnum, invoiceItemName string, invoiceItemCount string, invoiceItemWord string, invoiceItemPrice string, delayDay int, invType string, ) *AioCheckOutInvoiceOption`

NewAioCheckOutInvoiceOption instantiates a new AioCheckOutInvoiceOption object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAioCheckOutInvoiceOptionWithDefaults

`func NewAioCheckOutInvoiceOptionWithDefaults() *AioCheckOutInvoiceOption`

NewAioCheckOutInvoiceOptionWithDefaults instantiates a new AioCheckOutInvoiceOption object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRelateNumber

`func (o *AioCheckOutInvoiceOption) GetRelateNumber() string`

GetRelateNumber returns the RelateNumber field if non-nil, zero value otherwise.

### GetRelateNumberOk

`func (o *AioCheckOutInvoiceOption) GetRelateNumberOk() (*string, bool)`

GetRelateNumberOk returns a tuple with the RelateNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelateNumber

`func (o *AioCheckOutInvoiceOption) SetRelateNumber(v string)`

SetRelateNumber sets RelateNumber field to given value.


### GetCustomerID

`func (o *AioCheckOutInvoiceOption) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *AioCheckOutInvoiceOption) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *AioCheckOutInvoiceOption) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *AioCheckOutInvoiceOption) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetCustomerIdentifier

`func (o *AioCheckOutInvoiceOption) GetCustomerIdentifier() string`

GetCustomerIdentifier returns the CustomerIdentifier field if non-nil, zero value otherwise.

### GetCustomerIdentifierOk

`func (o *AioCheckOutInvoiceOption) GetCustomerIdentifierOk() (*string, bool)`

GetCustomerIdentifierOk returns a tuple with the CustomerIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerIdentifier

`func (o *AioCheckOutInvoiceOption) SetCustomerIdentifier(v string)`

SetCustomerIdentifier sets CustomerIdentifier field to given value.

### HasCustomerIdentifier

`func (o *AioCheckOutInvoiceOption) HasCustomerIdentifier() bool`

HasCustomerIdentifier returns a boolean if a field has been set.

### GetCustomerName

`func (o *AioCheckOutInvoiceOption) GetCustomerName() string`

GetCustomerName returns the CustomerName field if non-nil, zero value otherwise.

### GetCustomerNameOk

`func (o *AioCheckOutInvoiceOption) GetCustomerNameOk() (*string, bool)`

GetCustomerNameOk returns a tuple with the CustomerName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerName

`func (o *AioCheckOutInvoiceOption) SetCustomerName(v string)`

SetCustomerName sets CustomerName field to given value.

### HasCustomerName

`func (o *AioCheckOutInvoiceOption) HasCustomerName() bool`

HasCustomerName returns a boolean if a field has been set.

### GetCustomerAddr

`func (o *AioCheckOutInvoiceOption) GetCustomerAddr() string`

GetCustomerAddr returns the CustomerAddr field if non-nil, zero value otherwise.

### GetCustomerAddrOk

`func (o *AioCheckOutInvoiceOption) GetCustomerAddrOk() (*string, bool)`

GetCustomerAddrOk returns a tuple with the CustomerAddr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerAddr

`func (o *AioCheckOutInvoiceOption) SetCustomerAddr(v string)`

SetCustomerAddr sets CustomerAddr field to given value.

### HasCustomerAddr

`func (o *AioCheckOutInvoiceOption) HasCustomerAddr() bool`

HasCustomerAddr returns a boolean if a field has been set.

### GetCustomerPhone

`func (o *AioCheckOutInvoiceOption) GetCustomerPhone() string`

GetCustomerPhone returns the CustomerPhone field if non-nil, zero value otherwise.

### GetCustomerPhoneOk

`func (o *AioCheckOutInvoiceOption) GetCustomerPhoneOk() (*string, bool)`

GetCustomerPhoneOk returns a tuple with the CustomerPhone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerPhone

`func (o *AioCheckOutInvoiceOption) SetCustomerPhone(v string)`

SetCustomerPhone sets CustomerPhone field to given value.

### HasCustomerPhone

`func (o *AioCheckOutInvoiceOption) HasCustomerPhone() bool`

HasCustomerPhone returns a boolean if a field has been set.

### GetCustomerEmail

`func (o *AioCheckOutInvoiceOption) GetCustomerEmail() string`

GetCustomerEmail returns the CustomerEmail field if non-nil, zero value otherwise.

### GetCustomerEmailOk

`func (o *AioCheckOutInvoiceOption) GetCustomerEmailOk() (*string, bool)`

GetCustomerEmailOk returns a tuple with the CustomerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerEmail

`func (o *AioCheckOutInvoiceOption) SetCustomerEmail(v string)`

SetCustomerEmail sets CustomerEmail field to given value.

### HasCustomerEmail

`func (o *AioCheckOutInvoiceOption) HasCustomerEmail() bool`

HasCustomerEmail returns a boolean if a field has been set.

### GetClearanceMark

`func (o *AioCheckOutInvoiceOption) GetClearanceMark() ClearanceMarkEnum`

GetClearanceMark returns the ClearanceMark field if non-nil, zero value otherwise.

### GetClearanceMarkOk

`func (o *AioCheckOutInvoiceOption) GetClearanceMarkOk() (*ClearanceMarkEnum, bool)`

GetClearanceMarkOk returns a tuple with the ClearanceMark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClearanceMark

`func (o *AioCheckOutInvoiceOption) SetClearanceMark(v ClearanceMarkEnum)`

SetClearanceMark sets ClearanceMark field to given value.

### HasClearanceMark

`func (o *AioCheckOutInvoiceOption) HasClearanceMark() bool`

HasClearanceMark returns a boolean if a field has been set.

### GetTaxType

`func (o *AioCheckOutInvoiceOption) GetTaxType() TaxTypeEnum`

GetTaxType returns the TaxType field if non-nil, zero value otherwise.

### GetTaxTypeOk

`func (o *AioCheckOutInvoiceOption) GetTaxTypeOk() (*TaxTypeEnum, bool)`

GetTaxTypeOk returns a tuple with the TaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxType

`func (o *AioCheckOutInvoiceOption) SetTaxType(v TaxTypeEnum)`

SetTaxType sets TaxType field to given value.


### GetCarruerType

`func (o *AioCheckOutInvoiceOption) GetCarruerType() CarruerTypeEnum`

GetCarruerType returns the CarruerType field if non-nil, zero value otherwise.

### GetCarruerTypeOk

`func (o *AioCheckOutInvoiceOption) GetCarruerTypeOk() (*CarruerTypeEnum, bool)`

GetCarruerTypeOk returns a tuple with the CarruerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarruerType

`func (o *AioCheckOutInvoiceOption) SetCarruerType(v CarruerTypeEnum)`

SetCarruerType sets CarruerType field to given value.

### HasCarruerType

`func (o *AioCheckOutInvoiceOption) HasCarruerType() bool`

HasCarruerType returns a boolean if a field has been set.

### GetCarruerNum

`func (o *AioCheckOutInvoiceOption) GetCarruerNum() string`

GetCarruerNum returns the CarruerNum field if non-nil, zero value otherwise.

### GetCarruerNumOk

`func (o *AioCheckOutInvoiceOption) GetCarruerNumOk() (*string, bool)`

GetCarruerNumOk returns a tuple with the CarruerNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarruerNum

`func (o *AioCheckOutInvoiceOption) SetCarruerNum(v string)`

SetCarruerNum sets CarruerNum field to given value.

### HasCarruerNum

`func (o *AioCheckOutInvoiceOption) HasCarruerNum() bool`

HasCarruerNum returns a boolean if a field has been set.

### GetDonation

`func (o *AioCheckOutInvoiceOption) GetDonation() InvoiceDonationEunm`

GetDonation returns the Donation field if non-nil, zero value otherwise.

### GetDonationOk

`func (o *AioCheckOutInvoiceOption) GetDonationOk() (*InvoiceDonationEunm, bool)`

GetDonationOk returns a tuple with the Donation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDonation

`func (o *AioCheckOutInvoiceOption) SetDonation(v InvoiceDonationEunm)`

SetDonation sets Donation field to given value.


### GetLoveCode

`func (o *AioCheckOutInvoiceOption) GetLoveCode() string`

GetLoveCode returns the LoveCode field if non-nil, zero value otherwise.

### GetLoveCodeOk

`func (o *AioCheckOutInvoiceOption) GetLoveCodeOk() (*string, bool)`

GetLoveCodeOk returns a tuple with the LoveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoveCode

`func (o *AioCheckOutInvoiceOption) SetLoveCode(v string)`

SetLoveCode sets LoveCode field to given value.

### HasLoveCode

`func (o *AioCheckOutInvoiceOption) HasLoveCode() bool`

HasLoveCode returns a boolean if a field has been set.

### GetPrint

`func (o *AioCheckOutInvoiceOption) GetPrint() InvoicePrintEnum`

GetPrint returns the Print field if non-nil, zero value otherwise.

### GetPrintOk

`func (o *AioCheckOutInvoiceOption) GetPrintOk() (*InvoicePrintEnum, bool)`

GetPrintOk returns a tuple with the Print field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrint

`func (o *AioCheckOutInvoiceOption) SetPrint(v InvoicePrintEnum)`

SetPrint sets Print field to given value.


### GetInvoiceItemName

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemName() string`

GetInvoiceItemName returns the InvoiceItemName field if non-nil, zero value otherwise.

### GetInvoiceItemNameOk

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemNameOk() (*string, bool)`

GetInvoiceItemNameOk returns a tuple with the InvoiceItemName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemName

`func (o *AioCheckOutInvoiceOption) SetInvoiceItemName(v string)`

SetInvoiceItemName sets InvoiceItemName field to given value.


### GetInvoiceItemCount

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemCount() string`

GetInvoiceItemCount returns the InvoiceItemCount field if non-nil, zero value otherwise.

### GetInvoiceItemCountOk

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemCountOk() (*string, bool)`

GetInvoiceItemCountOk returns a tuple with the InvoiceItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemCount

`func (o *AioCheckOutInvoiceOption) SetInvoiceItemCount(v string)`

SetInvoiceItemCount sets InvoiceItemCount field to given value.


### GetInvoiceItemWord

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemWord() string`

GetInvoiceItemWord returns the InvoiceItemWord field if non-nil, zero value otherwise.

### GetInvoiceItemWordOk

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemWordOk() (*string, bool)`

GetInvoiceItemWordOk returns a tuple with the InvoiceItemWord field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemWord

`func (o *AioCheckOutInvoiceOption) SetInvoiceItemWord(v string)`

SetInvoiceItemWord sets InvoiceItemWord field to given value.


### GetInvoiceItemPrice

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemPrice() string`

GetInvoiceItemPrice returns the InvoiceItemPrice field if non-nil, zero value otherwise.

### GetInvoiceItemPriceOk

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemPriceOk() (*string, bool)`

GetInvoiceItemPriceOk returns a tuple with the InvoiceItemPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemPrice

`func (o *AioCheckOutInvoiceOption) SetInvoiceItemPrice(v string)`

SetInvoiceItemPrice sets InvoiceItemPrice field to given value.


### GetInvoiceItemTaxType

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemTaxType() string`

GetInvoiceItemTaxType returns the InvoiceItemTaxType field if non-nil, zero value otherwise.

### GetInvoiceItemTaxTypeOk

`func (o *AioCheckOutInvoiceOption) GetInvoiceItemTaxTypeOk() (*string, bool)`

GetInvoiceItemTaxTypeOk returns a tuple with the InvoiceItemTaxType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceItemTaxType

`func (o *AioCheckOutInvoiceOption) SetInvoiceItemTaxType(v string)`

SetInvoiceItemTaxType sets InvoiceItemTaxType field to given value.

### HasInvoiceItemTaxType

`func (o *AioCheckOutInvoiceOption) HasInvoiceItemTaxType() bool`

HasInvoiceItemTaxType returns a boolean if a field has been set.

### GetInvoiceRemark

`func (o *AioCheckOutInvoiceOption) GetInvoiceRemark() string`

GetInvoiceRemark returns the InvoiceRemark field if non-nil, zero value otherwise.

### GetInvoiceRemarkOk

`func (o *AioCheckOutInvoiceOption) GetInvoiceRemarkOk() (*string, bool)`

GetInvoiceRemarkOk returns a tuple with the InvoiceRemark field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoiceRemark

`func (o *AioCheckOutInvoiceOption) SetInvoiceRemark(v string)`

SetInvoiceRemark sets InvoiceRemark field to given value.

### HasInvoiceRemark

`func (o *AioCheckOutInvoiceOption) HasInvoiceRemark() bool`

HasInvoiceRemark returns a boolean if a field has been set.

### GetDelayDay

`func (o *AioCheckOutInvoiceOption) GetDelayDay() int`

GetDelayDay returns the DelayDay field if non-nil, zero value otherwise.

### GetDelayDayOk

`func (o *AioCheckOutInvoiceOption) GetDelayDayOk() (*int, bool)`

GetDelayDayOk returns a tuple with the DelayDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelayDay

`func (o *AioCheckOutInvoiceOption) SetDelayDay(v int)`

SetDelayDay sets DelayDay field to given value.


### GetInvType

`func (o *AioCheckOutInvoiceOption) GetInvType() string`

GetInvType returns the InvType field if non-nil, zero value otherwise.

### GetInvTypeOk

`func (o *AioCheckOutInvoiceOption) GetInvTypeOk() (*string, bool)`

GetInvTypeOk returns a tuple with the InvType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvType

`func (o *AioCheckOutInvoiceOption) SetInvType(v string)`

SetInvType sets InvType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


