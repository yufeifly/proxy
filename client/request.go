package client

import (
	"fmt"
	"github.com/levigross/grequests"
	"io"
	"io/ioutil"
)

func ensureReaderClosed(response *grequests.Response) {
	if body := response.RawResponse.Body; body != nil {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, body, 512)
		response.RawResponse.Body.Close()
	}
}

// getAPIPath path means webapi path, for example: /redis/set
func (cli *Client) getAPIPath(path string) string {
	return fmt.Sprintf("http://%s:%s%s", cli.addr.IP, cli.addr.Port, path)
}
