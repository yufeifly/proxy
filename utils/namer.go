package utils

import (
	"fmt"
	"strconv"
	"strings"
)

/* RenameService give service a new name,
useful in migration*/
func RenameService(oldName string) string {
	var newName string
	dot := strings.Index(oldName, ".")
	newName = oldName[:dot+1] + adder(oldName[dot+1:])
	return newName
}

func adder(tail string) string {
	num, _ := strconv.Atoi(tail)
	num++
	return strconv.Itoa(num)
}

// NameServiceByProxyService give service a name by proxy service
func NameServiceByProxyService(proxyID string) string {
	return fmt.Sprintf("%s.1", proxyID)
}
