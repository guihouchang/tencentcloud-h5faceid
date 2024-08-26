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
	OrderNO string `json:"orderNo"`
	Name    string `json:"name"`
	IDNO    string `json:"idNo"`
	// UserID 用户 ID ，用户的唯一标识 （不能带有特 殊字符），需要跟生成签名的 userId 保持一致
	UserID string `json:"userId"`
	// SourcePhotoStr  比对源照片，注意：原始图片不能超过 500k，且必须为 JPG 或 PNG、BMP 格式。 参数有值：使用合作伙伴提供的比对源照片进行比对，
	// 必须注照片是正脸可信照片，照片质量由合作方保证。参数为空 ：根据身份证号+ 姓名使用权威数据源比对
	SourcePhotoStr string `json:"sourcePhotoStr"`
	// SourcePhotoType 比对源照片类型，参数值为 1 时是：水纹正脸照。参数值为 2 时是：高清正脸照。重要提示：照片上无水波纹的为高清照，请勿传错，否则影响比对准确率。如有疑问，请联系腾讯云技术支持线下确认。
	SourcePhotoType string `json:"sourcePhotoType"`
	// LiveInterType  活体交互模式参数值为 1 时，表示仅使用实时检测模式，不兼容的情况下回调错误码 3005；参数值非 1 或不入参，表示优先使用实时检测模式，如遇不兼容情况，自动降级为视频录制模式。
	LiveInterType string `json:"liveInterType"`
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

type H5FaceURLRequest struct {
	OrderNo  string `json:"orderNo"`  // 订单号，由合作方上送，字母/数字组成的字符串,每次唯一，此信息为本次人脸核身上送的信息，不能超过 32 位
	H5FaceID string `json:"h5faceId"` // h5/geth5faceid接口返回的唯一标识。
	URL      string `json:"url"`      // H5人脸核身完成后回调的第三方URL，需要第三方提供完整URL且做URL Encode
	// ResultType 是否显示结果页面，参数值为“1”时直接跳转到url 回调地址，null 或其他值跳转提供的结果页面
	ResultType   string `json:"resultType"`
	UserID       string `json:"userId"`
	From         string `json:"from"`         // browser ：表示在浏览器启动刷脸 app ：表示在 app 里启动刷脸 默认值为 app
	RedirectType string `json:"redirectType"` // 跳转模式，参数值为“1”时，刷脸页面使用 replace方式跳转，不在浏览器history 中留下记录；传或其他值则正常跳转
}

type H5FaceURLResponse struct {
}
