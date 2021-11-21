package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"

	sdk "github.com/tencentqidian/qidian-sdk-go"
	"github.com/tencentqidian/qidian-sdk-go/handler"
	"github.com/tencentqidian/qidian-sdk-go/types"
)

var (
	token     string
	appID     string
	appSecret string
	aesKey    string
)

func init() {
	token = os.Getenv("QD_MESSAGE_TOKEN")
	appID = os.Getenv("QD_APP_ID")
	appSecret = os.Getenv("QD_APP_SECRET")
	aesKey = os.Getenv("QD_AES_KEY")
}

/**
export QD_MESSAGE_TOKEN=test123
export QD_APP_ID=123
export QD_APP_SECRET=qwwwwww
export QD_AES_KEY=vvvvv
*/
func main() {
	if token == "" || appID == "" || appSecret == "" || aesKey == "" {
		fmt.Println("environment variables `QD_MESSAGE_TOKEN | QD_APP_ID | QD_APP_SECRET | QD_AES_KEY` empty")
		os.Exit(1)
	}

	crypto, err := handler.NewCrypto(aesKey, appID)
	if err != nil {
		log.Print(err)
		return
	}

	cmpt := sdk.NewComponent(
		sdk.WithComponentAppID(appID),
		sdk.WithComponentAppSecret(appSecret),
	)

	http.Handle("/", WrapDumper(nil))
	http.Handle("/verify_receiver", WrapDumper(handler.NewVerifyHandler(token)))
	http.Handle("/command_receiver", WrapDumper(
		handler.NewCommandHandler(handler.CommandHandle(func(cmd *types.Command) error {
			resp, err := cmpt.GetOAuthAPPToken(context.Background(), &types.GetOAuthAPPTokenReq{
				AuthorizationCode: cmd.Code,
			})
			if err != nil {
				log.Print("GetOAuthAPPToken err: ", err)
				return nil
			}

			log.Printf("GetOAuthAPPToken->AuthorizationInfo: %+v", resp.AuthorizationInfo)

			// 刷新 Token
			refreshResp, err := cmpt.RefreshAuthorizerToken(context.Background(), &types.RefreshAuthorizerTokenReq{
				AuthorizerAppID:        resp.AuthorizationInfo.AuthorizerAppID,
				AuthorizerRefreshToken: resp.AuthorizationInfo.AuthorizerRefreshToken,
				SID:                    resp.AuthorizationInfo.ApplicationID,
			})
			if err != nil {
				log.Print("RefreshAuthorizerToken err: ", err)
				return nil
			}
			log.Printf("RefreshAuthorizerToken: %+v", refreshResp)

			// 获取企业信息
			client := sdk.NewClient(
				sdk.WithAccessToken(refreshResp.AuthorizerAccessToken),
			)
			corpInfoResp, err := client.GetCorpInfo(context.Background())
			if err != nil {
				log.Print("GetCorpInfo err: ", err)
				return nil
			}
			log.Printf("GetCorpInfo: %+v", corpInfoResp)

			// 获取部门列表
			depts, err := client.Departments(context.Background())
			if err != nil {
				log.Print("Departments err: ", err)
				return nil
			}
			log.Printf("Departments: %+v", depts)

			return nil
		})),
	))
	http.Handle("/message_receiver", WrapDumper(
		handler.NewMessageHandler(
			crypto,
			handler.On("EnableApplication", func(m *types.Message) {
				log.Printf("启用APP: %s\n", m.ApplicationId)
			}),
			handler.On("UnauthorizeApplication", func(m *types.Message) {
				log.Printf("卸载了APP: %s\n", m.ApplicationId)
			}),
		),
	))

	http.Handle("/ticket_receiver", WrapDumper(
		handler.NewTicketHandler(crypto, func(tm *handler.TicketMessage) {
			log.Printf("TicketMessage: %+v", tm)

			// check whether ComponentAccessToken has expired?
			resp, err := cmpt.GetComponentToken(context.Background(), &types.GetComponentTokenReq{
				ComponentVerifyTicket: tm.ComponentVerifyTicket,
			})
			if err != nil {
				log.Print("GetComponentToken err: ", err)
				return
			}
			log.Printf("GetComponentToken resp: %+v", resp)

			cmpt.Init(sdk.WithComponentAccessToken(resp.ComponentAccessToken))
		}),
	))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

type dumper struct {
	h http.Handler
}

func WrapDumper(h http.Handler) http.Handler {
	return &dumper{h: h}
}

func (d *dumper) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dump, _ := httputil.DumpRequest(r, true)
	log.Print(string(dump))

	if d.h != nil {
		d.h.ServeHTTP(w, r)
	}
}

type methodHandler struct {
	handlers map[string]http.Handler
}

func matchMethods(h map[string]http.Handler) http.Handler {
	return &methodHandler{handlers: h}
}

func (m *methodHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if h, ok := m.handlers[r.Method]; ok {
		h.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, http.StatusText(http.StatusNotFound))
	}
}
