package utils

import (
	"github.com/yufeifly/proxy/api/types"
	"github.com/yufeifly/proxy/config"
	"github.com/yufeifly/proxy/cusErr"
	"regexp"
	"strings"
)

// ParseAddress 127.0.0.1:6789 -> ip , port, it will add default port if port does not exist
func ParseAddress(raw string) (types.Address, error) {
	if raw == "" {
		return types.Address{IP: "127.0.0.1", Port: config.DefaultMigratorListeningPort}, nil
	}
	addr := types.Address{}
	var ip, port string

	colonInd := strings.Index(raw, ":")
	if colonInd == -1 { // not found colon, means no port
		ip = raw
		addr.Port = config.DefaultMigratorListeningPort
	} else {
		ip = raw[:colonInd]
		port = raw[colonInd+1:]
	}

	matchedIP, err := regexp.MatchString("^((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", ip)
	if err != nil {
		return types.Address{}, err
	}
	if !matchedIP {
		return types.Address{}, cusErr.ErrBadAddress
	}
	addr.IP = ip
	if port != "" {
		portMatched, err := regexp.MatchString("^[1-9]\\d*$", port)
		if err != nil {
			return types.Address{}, err
		}
		if !portMatched {
			return types.Address{}, cusErr.ErrBadAddress
		}
		addr.Port = port
	}

	return addr, nil
}
