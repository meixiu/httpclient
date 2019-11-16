package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type (
	Response struct {
		OK          bool
		Error       error
		StatusCode  int
		Header      http.Header
		RawResponse *http.Response
		bodyInit    bool
		bodyBuffer  *bytes.Buffer
	}
)

func NewResponse(resp *http.Response) *Response {
	var ok bool
	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		ok = true
	}
	newResp := &Response{
		OK:          ok,
		StatusCode:  resp.StatusCode,
		Header:      resp.Header,
		RawResponse: resp,
		bodyInit:    false,
		bodyBuffer:  bytes.NewBuffer([]byte{}),
	}
	newResp.initBodyBuffer()
	return newResp
}

func (resp *Response) initBodyBuffer() {
	if resp.bodyInit {
		return
	}
	defer func() {
		resp.RawResponse.Body.Close()
	}()

	resp.bodyInit = true
	if resp.RawResponse.ContentLength == 0 {
		return
	}
	if resp.RawResponse.ContentLength > 0 {
		resp.bodyBuffer.Grow(int(resp.RawResponse.ContentLength))
	}
	if _, err := io.Copy(resp.bodyBuffer, resp.RawResponse.Body); err != nil && err != io.EOF {
		resp.Error = err
	}
}

func (resp *Response) Bytes() []byte {
	if resp.Error != nil {
		return nil
	}
	if resp.bodyBuffer.Len() == 0 {
		return nil
	}
	return resp.bodyBuffer.Bytes()
}

func (resp *Response) String() string {
	if resp.Error != nil {
		return ""
	}
	return resp.bodyBuffer.String()
}

func (resp *Response) Decode(v interface{}) error {
	if resp.Error != nil {
		return resp.Error
	}
	return json.NewDecoder(resp.bodyBuffer).Decode(v)
}
