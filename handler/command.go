package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/tencentqidian/qidian-sdk-go/types"
)

type CommandHandler struct {
	handle func(*types.Command) error
}

type CommandOption func(*CommandHandler)

// CommandHandle .
func CommandHandle(fn func(*types.Command) error) CommandOption {
	return func(h *CommandHandler) {
		h.handle = fn
	}
}

func NewCommandHandler(opt ...CommandOption) http.Handler {
	h := &CommandHandler{}
	for _, o := range opt {
		o(h)
	}
	return h
}

func (h *CommandHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()

	appID, _ := strconv.Atoi(q.Get("app_id"))
	sid, _ := strconv.Atoi(q.Get("sid"))

	err := h.handle(&types.Command{
		Code:  q.Get("code"),
		State: q.Get("state"),
		AppID: appID,
		SID:   sid,
	})

	if err != nil {
		log.Print("command action not allow: ", err)
		w.Write([]byte("failed"))
	} else {
		// required
		w.Write([]byte("success"))
	}
}
