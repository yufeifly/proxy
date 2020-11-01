package utils

import (
	"github.com/yufeifly/proxy/config"
	"github.com/yufeifly/proxy/cusErr"
	"github.com/yufeifly/proxy/model"
	"regexp"
	"strings"
)

// 127.0.0.1:6789 -> ip , port, it will add default port if port does not exist
func ParseAddress(raw string) (model.Address, error) {
	if raw == "" {
		return model.Address{IP: "127.0.0.1", Port: config.DefaultMigratorListeningPort}, nil
	}
	addr := model.Address{}
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
		return model.Address{}, err
	}
	if !matchedIP {
		return model.Address{}, cusErr.ErrBadAddress
	}
	addr.IP = ip
	if port != "" {
		portMatched, err := regexp.MatchString("^[1-9]\\d*$", port)
		if err != nil {
			return model.Address{}, err
		}
		if !portMatched {
			return model.Address{}, cusErr.ErrBadAddress
		}
		addr.Port = port
	}

	return addr, nil
}
