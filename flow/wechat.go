package flow

import (
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/eatmoreapple/openwechat"
	"github.com/pkg/errors"
	"github.com/skip2/go-qrcode"
)

const (
	url = "https://login.weixin.qq.com/l/"
)

type WeChat struct {
	Bot *openwechat.Bot
}

func (w *WeChat) Init(_ context.Context) error {
	w.Bot = openwechat.DefaultBot(openwechat.Desktop)
	w.Bot.UUIDCallback = w.show

	dispatcher := openwechat.NewMessageMatchDispatcher()
	dispatcher.OnImage(w.handle)
	dispatcher.OnMedia(w.handle)
	dispatcher.OnText(w.handle)
	dispatcher.OnVoice(w.handle)

	w.Bot.MessageHandler = dispatcher.AsMessageHandler()

	return nil
}

func (w *WeChat) Deinit(_ context.Context) error {
	if err := w.Bot.Logout(); err != nil {
		return errors.Wrap(err, "failed to logout\n")
	}

	return nil
}

func (w *WeChat) Run(_ context.Context) error {
	if err := w.Bot.Login(); err != nil {
		return errors.Wrap(err, "failed to login\n")
	}

	if err := w.Bot.Block(); err != nil {
		return errors.Wrap(err, "failed to block\n")
	}

	return nil
}

func (w *WeChat) show(uuid string) {
	q, _ := qrcode.New(url+uuid, qrcode.Low)

	fmt.Println(q.ToString(true))
}

func (w *WeChat) handle(ctx *openwechat.MessageContext) {
	var err error

	msg := ctx.Message

	if msg.IsPicture() {
		_, err = msg.ReplyImage(w.handleImage(msg))
	} else if msg.IsMedia() {
		_, err = msg.ReplyFile(w.handleFile(msg))
	} else if msg.IsText() {
		_, err = msg.ReplyText(w.handleText(msg))
	} else if msg.IsVoice() {
		_, err = msg.ReplyText(w.handleVoice(msg))
	} else {
		err = errors.New("invalid msg type\n")
	}

	if err != nil {
		_, _ = msg.ReplyText(err.Error())
	}
}

func (w *WeChat) handleImage(msg *openwechat.Message) io.Reader {
	// TBD: FIXME
	return strings.NewReader(msg.Content)
}

func (w *WeChat) handleFile(msg *openwechat.Message) io.Reader {
	// TBD: FIXME
	return strings.NewReader(msg.Content)
}

func (w *WeChat) handleText(msg *openwechat.Message) string {
	// TBD: FIXME
	var ret string
	ret = msg.Content
	return ret
}

func (w *WeChat) handleVoice(msg *openwechat.Message) string {
	// TBD: FIXME
	var ret string
	return ret
}
