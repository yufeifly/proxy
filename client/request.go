package client

import (
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
