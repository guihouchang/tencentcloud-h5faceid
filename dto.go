package faceid

type AccessTokenRequest struct {
	AppID     string `json:"app_id"`
	Secret    string `json:"secret"`
	Version   string `json:"version"`
	GrantType string `json:"grant_type"`
}

type AccessTokenResponse struct {
	Code            string `json:"code"`
	TransactionTime string `json:"transactionTime"`
	AccessToken     string `json:"access_token"`
	ExpireTime      string `json:"expire_time"`
	ExpireIn        int    `json:"expire_in"`
	Msg             string `json:"msg"`
}

type APITicketRequest struct {
	APPID       string `json:"app_id"`
	AccessToken string `json:"access_token"`
	Type        string `json:"type"`
	Version     string `json:"version"`
}

type Ticket struct {
	Value      string `json:"value"`
	ExpireTime string `json:"expire_time"`
	ExpireIn   int    `json:"expire_in"`
}

type APITicketResponse struct {
	Code            string   `json:"code"`
	Msg             string   `json:"msg"`
	TransactionTime string   `json:"transactionTime"`
	Tickets         []Ticket `json:"tickets"`
}

type NonceAPITicketRequest struct {
	APPID       string `json:"app_id"`
	AccessToken string `json:"access_token"`
	Type        string `json:"type"`
	Version     string `json:"version"`
	UserID      string `json:"user_id"`
}

type H5FaceIDRequest struct {
	OrderNO         string `json:"orderNo"`
	Name            string `json:"name"`
	IDNO            string `json:"idNo"`
	UserID          string `json:"userId"`
	SourcePhotoStr  string `json:"sourcePhotoStr"`
	SourcePhotoType string `json:"sourcePhotoType"`
	LiveInterType   string `json:"liveInterType"`
}

type H5FaceIDResponse struct {
	Code            string `json:"code"`
	Msg             string `json:"msg"`
	TransactionTime string `json:"transactionTime"`
	Result          struct {
		TransactionTime string `json:"transactionTime"`
		BizSeqNo        string `json:"bizSeqNo"`
		OrderNo         string `json:"orderNo"`
		H5FaceID        string `json:"h5faceId"`
		OptimalDomain   string `json:"optimalDomain"`
		Success         bool   `json:"success"`
	} `json:"result"`
}

type H5FaceRecordRequest struct {
	OrderNo string `json:"orderNo"`
	GetFile string `json:"getFile"`
}

type H5FaceRecordResponse struct {
	Code     string `json:"code"`
	Msg      string `json:"msg"`
	BizSeqNo string `json:"bizSeqNo"`
	Result   struct {
		OrderNo      string `json:"orderNo"`
		LiveRate     string `json:"liveRate"`
		Similarity   string `json:"similarity"`
		OccurredTime string `json:"occurredTime"`
		AppId        string `json:"appId"`
		Photo        string `json:"photo"`
		Video        string `json:"video"`
		BizSeqNo     string `json:"bizSeqNo"`
		SdkVersion   string `json:"sdkVersion"`
		TrtcFlag     string `json:"trtcFlag"`
	} `json:"result"`
}
