package protocol

import (
	"github.com/sartoopjj/thefeed/internal/protocol"
)

func DeriveKeys(passphrase string) (queryKey [32]byte, responseKey [32]byte, err error) {
	return protocol.DeriveKeys(passphrase)
}

func EncodeQuery(queryKey [protocol.KeySize]byte, channel, block uint16, domain string, mode protocol.QueryEncoding) (string, error) {
	return protocol.EncodeQuery(queryKey, channel, block, domain, mode)
}

func DecodeResponse(responseKey [protocol.KeySize]byte, encoded string) ([]byte, error) {
	return protocol.DecodeResponse(responseKey, encoded)
}
