package tests

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	sdk "github.com/tencentqidian/qidian-sdk-go"
	"github.com/tencentqidian/qidian-sdk-go/types"
)

var (
	componentAccessToken string
	appID                string
	secret               string
)

var client *sdk.Client

func init() {
	appID = os.Getenv("QD_APP_ID")
	if appID == "" {
		panic("environment variable `QD_APP_ID` is required")
	}
	secret = os.Getenv("QD_SECRET")
	if secret == "" {
		panic("environment variable `QD_SECRET` is required")
	}

	componentAccessToken = os.Getenv("QD_COMPONENT_ACCESS_TOKEN")
	if componentAccessToken == "" {
		panic("environment variable `QD_COMPONENT_ACCESS_TOKEN` is required")
	}

	client = sdk.NewClient(
		sdk.WithAppID(appID),
		sdk.WithSecret(secret),
		sdk.WithComponentAccessToken(componentAccessToken),
	)

}

func TestGetOAuthAPPToken(t *testing.T) {
	resp, err := client.GetOAuthAPPToken(&types.GetOAuthAPPTokenReq{
		ComponentAppID:    appID,
		AuthorizationCode: "",
	})
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.AuthorizationInfo.AuthorizerAccessToken)
	}
}

func TestDepartments(t *testing.T) {
	resp, err := client.Departments()
	if assert.Nil(t, err) {
		assert.NotNil(t, resp)
	}
}
