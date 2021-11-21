package types

type GetOAuthAPPTokenReq struct {
	ComponentAppID    string `json:"component_appid"`    // 应用开发商的appid
	AuthorizationCode string `json:"authorization_code"` // 应用授权code,会在授权成功时返回给应用开发者，详见应用授权code获取说明
}

type AuthorizationInfo struct {
	AuthorizerAppID        string `json:"authorizer_appid"`         // 应用授权方appid
	AuthorizerAccessToken  string `json:"authorizer_access_token"`  // 应用授权方的接口调用凭证（企业授权token） (长度最长为196)
	ExpiresIn              int    `json:"expires_in"`               // 过期时间，单位秒（s）
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"` // 授权方的刷新令牌（应用授权刷新token),该令牌的有效期跟随应用授权关系，如果应用授权关系不终止则永久有效 (长度最长为128)
	ApplicationID          string `json:"applicationId"`            // 应用 id
}

type GetOAuthAPPTokenRsp struct {
	AuthorizationInfo *AuthorizationInfo `json:"authorization_info"`
}

type GetComponentTokenReq struct {
	ComponentAppID        string `json:"component_appid"`
	ComponentAppSecret    string `json:"component_appsecret"`
	ComponentVerifyTicket string `json:"component_verify_ticket"`
}

type GetComponentTokenRsp struct {
	ComponentAccessToken string `json:"component_access_token"` // 应用开发商token
	ExpiresIn            int    `json:"expires_in"`             // 有效时长(s)
}

type RefreshAuthorizerTokenReq struct {
	ComponentAppID         string `json:"component_appid"`          // 第三方应用开发商appid
	AuthorizerAppID        string `json:"authorizer_appid"`         // 应用授权方appid
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"` // 授权方的刷新令牌（应用授权刷新token),该令牌的有效期跟随应用授权关系，如果应用授权关系不终止则永久有效 (长度最长为128),应用没有被卸载将保持不变
	SID                    string `json:"sid"`                      // 应用id
}

type RefreshAuthorizerTokenRsp struct {
	AuthorizerAccessToken  string `json:"authorizer_access_token"`
	ExpiresIn              int    `json:"expires_in"`
	AuthorizerRefreshToken string `json:"authorizer_refresh_token"`
}
