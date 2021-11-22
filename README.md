企点开放平台 SDK
====

警告: 该 SDK 目前处于开发中，请勿用于生产环境

## 使用方法

```shell
go get github.com/tencentqidian/qidian-sdk-go
```

```go
import (
    sdk "github.com/tencentqidian/qidian-sdk-go"
)

func main()  {
    // 创建服务商对象
    cmpt := sdk.NewComponent(
        sdk.WithComponentAppID("APP_ID"),
        sdk.WithComponentAppSecret("APP_SECRET"),
    )
	
    // 使用应用授权code换取应用授权token
    resp, err := cmpt.GetOAuthAPPToken(context.Background(), &types.GetOAuthAPPTokenReq{
        AuthorizationCode: "AUTHORIZATION_CODE",
    })
	
    // 获取应用开发商 token
    resp, err := cmpt.GetComponentToken(context.Background(), &types.GetComponentTokenReq{
        ComponentVerifyTicket: "COMPONENT_VERIFY_TICKET",
    })

    // 刷新 Token
    refreshResp, err := cmpt.RefreshAuthorizerToken(context.Background(), &types.RefreshAuthorizerTokenReq{
        AuthorizerAppID:        "AUTHORIZER_APP_ID",
        AuthorizerRefreshToken: "AUTHORIZER_REFRESH_TOKEN",
        SID:                    "APPLICATION_ID",
    })
	
	
    // 通过 ACCESS_TOKEN 调用服务端 API 
    client := sdk.NewClient(
        sdk.WithAccessToken("ACCESS_TOKEN"),
    )

    // 获取企业信息
    corpInfo, err := client.GetCorpInfo(context.Background())
	
    // 获取部门列表
    depts, err := client.Departments(context.Background())
	
    //....
}
```

## 更多

查看 [Demo](example/demo/main.go)

## 开发文档
https://api.qidian.qq.com/wiki
