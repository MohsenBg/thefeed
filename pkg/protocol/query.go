package protocol

import "github.com/sartoopjj/thefeed/internal/protocol"

func EncodeQuerySimple(queryKey [32]byte, channel, block uint16, domain string, mode int) (string, error) {
	return EncodeQuery(queryKey, channel, block, domain, protocol.QueryEncoding(mode))
}
