package ecpay

import "github.com/Laysi/go-ecpay-sdk/base"

type AioCheckOutGeneralOptional struct {
	// **特店旗下店舖代號** 提供特店填入分店代號使用，僅可用英數字大小寫混合。
	StoreID *string `json:"StoreID,omitempty"`
	// **Client端返回特店的按鈕連結**   消費者點選此按鈕後，會將頁面導回到此設定的網址   注意事項：   導回時不會帶付款結果到此網址，只是將頁面導回而已。   設定此參數，綠界會在付款完成或取號完成頁面上顯示[返回商店]的按鈕。   設定此參數，發生簡訊 OTP 驗證失敗時，頁面上會顯示[返回商店]的按鈕。   若未設定此參數，則綠界付款完成頁或取號完成頁面，不會顯示[返回商店]的按鈕。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。
	ClientBackURL *string `json:"ClientBackURL,omitempty"`
	// **商品銷售網址**
	ItemURL *string `json:"ItemURL,omitempty"`
	// **備註欄位**
	Remark *string `json:"Remark,omitempty"`
	// **付款子項目**   若設定此參數，建立訂單將轉導至綠界訂單成立頁，依設定的付款方式及付款子項目帶入訂單，無法選擇其他付款子項目。   請參考付款方式一覽表
	ChooseSubPayment *ecpayBase.ChooseSubPaymentEnum `json:"ChooseSubPayment,omitempty"`
	// **Client端回傳付款結果網址**     當消費者付款完成後，綠界會將付款結果參數以幕前(Client POST)回傳到該網址。   詳細說明請參考付款結果通知   注意事項：   1. 若與[ClientBackURL]同時設定，將會以此參數為主。   2. 銀聯卡及非即時交易(ATM、CVS、BARCODE)不支援此參數。
	OrderResultURL *string `json:"OrderResultURL,omitempty"`
	// **是否需要額外的付款資訊**
	//   額外的付款資訊：
	//   若不回傳額外的付款資訊時，參數值請傳：`Ｎ`；
	//   若要回傳額外的付款資訊時，參數值請傳：`Ｙ`，付款完成後綠界會以 Server POST 方式回傳額外付款資訊。
	//   注意事項：
	//   回傳額外付款資訊參數請參考-額外回傳的參數
	NeedExtraPaidInfo *ecpayBase.NeedExtraPaidInfoEnum `json:"NeedExtraPaidInfo,omitempty"`
	// **特約合作平台商代號(由綠界提供)**   為專案合作的平台商使用。   一般特店或平台商本身介接，則參數請帶放空值。   若為專案合作平台商的特店使用時，則參數請帶平台商所綁的特店編號 `MerchantID`。
	PlatformID *string `json:"PlatformID,omitempty"`
	// **自訂名稱欄位1**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField1 *string `json:"CustomField1,omitempty"`
	// **自訂名稱欄位2**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField2 *string `json:"CustomField2,omitempty"`
	// **自訂名稱欄位3**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField3 *string `json:"CustomField3,omitempty"`
	// **自訂名稱欄位4**   提供合作廠商使用記錄用客製化使用欄位   注意事項：   特殊符號只支援 `,.#()$[];%{}:/?&@<>!`
	CustomField4 *string `json:"CustomField4,omitempty"`
	// **語系設定**
	//  預設語系為中文，若要變更語系參數值請帶：
	//  - 英語：`ENG`
	//  - 韓語：`KOR`
	//  - 日語：`JPN`
	//  - 簡體中文：`CHI`
	Language *ecpayBase.LanguageEnum `json:"Language,omitempty"`
}

type AioCheckOutAtmOptional struct {
	// **允許繳費有效天數**   若需設定最長 60 天，最短 1 天。   未設定此參數則預設為 3 天   注意事項：以天為單位
	ExpireDate *int `json:"ExpireDate,omitempty"`
	// **Server端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：銀行代碼、繳費虛擬帳號繳費期限…等)。   請參考[ATM、CVS 或 BARCODE 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。
	PaymentInfoURL *string `json:"PaymentInfoURL,omitempty"`
	// **Client端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Client 端回傳消費者付款方式相關資訊(例：銀行代碼、繳費虛擬帳號繳費期限…等)且將頁面轉到特店指定的頁面。請參考[ATM、CVS 或 BARCODE 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。
	ClientRedirectURL *string `json:"ClientRedirectURL,omitempty"`
}

type AioCheckOutCvsBarcodeOptional struct {
	// **超商繳費截止時間**   注意事項：   `CVS`:以分鐘為單位   `BARCODE`:以天為單位   若未設定此參數，`CVS` 預設為 `10080` 分鐘(`7` 天)；BARCODE 預設為 `7` 天。   若需設定此參數，請於建立訂單時將此參數送給綠界。提醒您，CVS 帶入數值不可超過 `86400` 分鐘，超過時一律以 `86400` 分鐘計(`60` 天)   例：`08/01` 的 `20:15` 分購買商品，繳費期限為 `7` 天，表示 `8/08` 的 `20:15` 分前您必須前往超商繳費。   範例: `CVS`=`1440`(共 `1` 天)、`BARCODE`=`7`(共 `7` 天)
	StoreExpireDate *int `json:"StoreExpireDate,omitempty"`
	// **交易描述1** 會出現在超商繳費平台螢幕上
	Desc1 *string `json:"Desc_1,omitempty"`
	// **交易描述2** 會出現在超商繳費平台螢幕上
	Desc2 *string `json:"Desc_2,omitempty"`
	// **交易描述3** 會出現在超商繳費平台螢幕上
	Desc3 *string `json:"Desc_3,omitempty"`
	// **交易描述4** 會出現在超商繳費平台螢幕上
	Desc4 *string `json:"Desc_4,omitempty"`
	// **Server 端回傳付款相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會 Server 端背景回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)。   請參考[`ATM`、`CVS` 或 `BARCODE` 的取號結果通知.]   注意事項：   頁面將會停留在綠界，顯示繳費的相關資訊。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。
	PaymentInfoURL *string `json:"PaymentInfoURL,omitempty"`
	// **Client端回傳付款方式相關資訊**   若有設定此參數，訂單建立完成後(非付款完成)，綠界會從 Client 端回傳消費者付款方式相關資訊(例：繳費代碼與繳費超商)且將頁面轉到特店指定的頁面。   請參考[`ATM`、`CVS` 或 `BARCODE` 的取號結果通知.]   注意事項：   若設定此參數，將會使設定的返回特店的按鈕連結[ClientBackURL]失效。   若導回網址未使用 https 時，部份瀏覽器可能會出現警告訊息。   回傳只有三段號碼，並不會回傳條碼圖，需自行轉換成 code39 的三段條碼。
	ClientRedirectURL *string `json:"ClientRedirectURL,omitempty"`
}

type AioCheckOutCreditOptional struct {
	// **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 `MerchantID` + `廠商會員編號`)
	BindingCard *ecpayBase.BindingCardEnum `json:"BindingCard,omitempty"`
	// **記憶卡號識別碼** 記憶卡號識別碼 (特店代號 `MerchantID` + `廠商會員編號`)
	MerchantMemberID *string `json:"MerchantMemberID,omitempty"`
}

type AioCheckOutCreditOnetimeOptional struct {
	// **信用卡是否使用紅利折抵**
	//  設為 Y 時，當綠界特店選擇信用卡付款時，會進入紅利折抵的交易流程。
	//  注意事項：
	//  紅利折抵請參考信用卡紅利折抵辦法
	Redeem *ecpayBase.RedeemEnum `json:"Redeem,omitempty"`
	//  **銀聯卡交易選項**
	//   可帶入以下選項:
	//   `0`: 消費者於交易頁面可選擇是否使用銀聯交易。
	//   `1`: 只使用銀聯卡交易，且綠界會將交易頁面直接導到銀聯網站。
	//   `2`: 不可使用銀聯卡，綠界會將交易頁面隱藏銀聯選項。
	//   注意事項：
	//   1.若需使用銀聯卡服務，請與綠界提出申請方可使用，測試環境未提供銀聯卡服務。
	//   2.不支援信用卡分期付款及定期定額。
	//   3.不支援信用卡紅利折抵
	//   4.不支援信用卡記憶卡號功能
	UnionPay *ecpayBase.UnionPayEnum `json:"UnionPay,omitempty"`
}

type AioCheckOutCreditPeriodOptional struct {
	// **定期定額的執行結果回應URL**   若交易是信用卡定期定額的方式，則每次執行授權完，會將授權結果回傳到這個設定的 URL。   回覆內容請參考。
	//PeriodReturnURL *string `json:"PeriodReturnURL,omitempty"` //這個會自動填入
}

type AioCheckOutInvoiceOptional struct {
	// **客戶編號** 該參數有值時，僅接受『英文、數字、下底線』等字元。
	CustomerID *string `json:"CustomerID,omitempty"`
	// **統一編號** 該參數有值時，請帶固定長度為數字 8 碼。
	CustomerIdentifier *string `json:"CustomerIdentifier,omitempty"`
	// **客戶名稱**    當列印註記`Print`為 `1`(列印)時，則該參數必須有值。   該參數有值時，僅接受『中、英文及數字』等字元。   請將參數值做 UrlEncode 方式編碼。
	CustomerName *string `json:"CustomerName,omitempty"`
	// **客戶地址**    當列印註記`Print`為 `1`(列印)時，則該參數必須有值。   當該參數有值時，請注意特殊字元轉換 。    請將參數值做 UrlEncode 方式編碼。
	CustomerAddr *string `json:"CustomerAddr,omitempty"`
	// **客戶手機號碼**   當客戶電子信箱`CustomerEmail`為空字串時，則該參數必須有值。   當該參數有值時，則格式為數字。   注意事項：   請填手機號碼，不能填市話因為要收簡訊通知用
	CustomerPhone *string `json:"CustomerPhone,omitempty"`
	// **客戶電子信箱**   當客戶手機號碼`CustomerPhone`為空字串時，則該參數必須有值。   當該參數有值時，則格式需符合 EMAIL格式。   請將參數值做 UrlEncode 方式編碼。
	CustomerEmail *string `json:"CustomerEmail,omitempty"`
	// **通關方式**
	//  當課稅類別`TaxType`為 `2`(零稅率)時，則該參數請帶 `1`(非經海關出口)或 `2`(經海關出口)。
	ClearanceMark *ecpayBase.ClearanceMarkEnum `json:"ClearanceMark,omitempty"`
	// **載具類別**
	//  若為無載具時，則請帶空字串。
	//  若為特店載具時，則請帶 `1`。
	//  若為買受人之自然人憑證號碼時，則請帶 `2`。
	//  若為買受人之手機條碼資料時，則請帶`3`。
	//  若統一編號 `CustomerIdentifier` 有值時，則載具類別不可為特店載具或自然人憑證載具。
	//  注意事項：當`Print`有值時，載具類別不得有值。
	CarruerType *ecpayBase.CarruerTypeEnum `json:"CarruerType,omitempty"`
	// **載具編號**   1. 當載具類別 `CarruerType`=``無載具)，請帶空字串。   2. 當載具類別`CarruerType`=`1`(綠界科技電子發票載具)時，請帶空字串，系統會自動帶入值，為合作特店載具統一編號+自訂編號(RelateNumber)。   3. 當載具類別`CarruerType`=`2`(買受人之自然人憑證)時，則請帶固定長度為16且格式 為2碼大寫英文字母加上14碼數字。   4. 當載具類別`CarruerType`=`3`(買受人之手機條碼)時，則請帶固定長度為 8且格式為 1 碼斜線「/」加上由 7 碼數字及大寫英文字母及+-.符號組成。    注意事項：   1. 若手機條碼中有加號，可能在介接驗證時 發生錯誤，請將加號改為空白字元，產生 驗證碼。   2. 英文、數字、符號僅接受半形字   3. 若載具編號為手機條碼載具時，請先呼叫B2C電子發票介接技術文件手機條碼載驗證ＡＰＩ進行檢核
	CarruerNum *string `json:"CarruerNum,omitempty"`
	// **捐贈碼**   消費者選擇捐贈發票則於此欄位須填入受贈單位之捐贈碼。   1. 若捐贈註記 `Donation`= `1` (捐贈)時，此欄位須有值。   2. 捐贈碼以阿拉伯數字為限，最少三碼，最多七碼。內容定位採「文字格式」，首位可以為零。
	LoveCode *string `json:"LoveCode,omitempty"`
	// **商品課稅別**   1：應稅   2：零稅率   3：免稅   注意事項：   1. 預設為空字串，當課稅類別 [TaxType] = 9 時，此欄位不可為空。   2. 格式為課稅 類別 1 | 課稅類別 2 | 課稅類別 3 | … | 課稅類別 n。當含有二筆或以上的商品課稅類別時，則以「|」符號區隔。   3. 課稅類別為混合稅率時，需含二筆或 以 上 的 商 品 課 稅   別[InvoiceItemTaxType]，且至少需有一筆商品課稅別為應稅及至少需有一筆商品課稅別為免稅或零稅率，即混稅發票只能 1.應稅+免稅 2.應稅+零稅率，免稅和零稅率發票不能同時開立。
	//InvoiceItemTaxType *string `json:"InvoiceItemTaxType,omitempty"` //由TaxType控制，需要的時候會從別的地方填入
	// **備註** 當該參數有值時，請將參數值做UrlEncode 方式編碼。
	InvoiceRemark *string `json:"InvoiceRemark,omitempty"`
}
