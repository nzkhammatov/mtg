package telegram

import (
	"net"
	"time"

	"github.com/juju/errors"

	"github.com/9seconds/mtg/config"
	"github.com/9seconds/mtg/wrappers"
)

const telegramDialTimeout = 10 * time.Second

type tgDialer struct {
	net.Dialer

	conf *config.Config
}

func (t *tgDialer) dial(addr string) (net.Conn, error) {
	conn, err := t.Dialer.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Annotate(err, "Cannot connect to Telegram")
	}
	if err = config.SetSocketOptions(conn); err != nil {
		return nil, errors.Annotate(err, "Cannot set socket options")
	}

	return conn, nil
}

func (t *tgDialer) dialRWC(addr, connID string) (wrappers.StreamReadWriteCloser, error) {
	conn, err := t.dial(addr)
	if err != nil {
		return nil, err
	}
	tgConn := wrappers.NewConn(conn, connID, wrappers.ConnPurposeTelegram, t.conf.PublicIPv4, t.conf.PublicIPv6)

	return tgConn, nil
}
