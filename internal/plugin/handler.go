package plugin

import (
	"context"
	"net"

	"github.com/Sherlock-Holo/errors"
	"github.com/coredns/coredns/request"
	"github.com/miekg/dns"
)

type Handler struct {
	dns.Client

	cfg Config
}

func NewHandler(cfg Config) *Handler {
	return &Handler{cfg: cfg}
}

func (h *Handler) Name() string {
	return "china-list"
}

func (h *Handler) ServeDNS(ctx context.Context, rw dns.ResponseWriter, msg *dns.Msg) (int, error) {
	state := request.Request{W: rw, Req: msg}

	reqName := state.Name()
	reqName = reqName[:len(reqName)-1]

	match, rootWord := h.cfg.Forest.Get(reqName)

	var dnsAddress string

	switch match {
	case "":
		dnsAddress = h.cfg.ForeignDns

	default:
		if len(match) > len(rootWord) {
			dnsAddress = h.cfg.ChinaDns
		} else {
			dnsAddress = h.cfg.ForeignDns
		}
	}

	conn, err := h.Dial(dnsAddress)
	if err != nil {
		return dns.RcodeServerFailure, errors.Wrapf(err, "dial dns %s failed", dnsAddress)
	}

	if deadline, ok := ctx.Deadline(); ok {
		if err := conn.SetDeadline(deadline); err != nil {
			return dns.RcodeServerFailure, errors.Wrapf(err, "set deadline %s failed", deadline)
		}
	}

	if err := conn.WriteMsg(msg); err != nil {
		return dns.RcodeServerFailure, errors.Wrap(err, "write msg failed")
	}

	replyMsg, err := conn.ReadMsg()

	var netError net.Error
	switch {
	case errors.As(err, &netError):
		if netError.Timeout() {
			return dns.RcodeBadTime, errors.Wrap(err, "read reply msg timeout")
		}

		fallthrough

	default:
		return dns.RcodeServerFailure, errors.Wrap(err, "read reply msg failed")

	case err == nil:
	}

	err = rw.WriteMsg(replyMsg)
	switch {
	case errors.As(err, &netError):
		if netError.Timeout() {
			return dns.RcodeBadTime, errors.Wrap(err, "write reply msg timeout")
		}

		fallthrough

	default:
		return dns.RcodeServerFailure, errors.Wrap(err, "write reply msg failed")

	case err == nil:
	}

	return replyMsg.Rcode, nil
}
