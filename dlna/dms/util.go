package dms

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func PrintXML(obj interface{}) {
	encoder := xml.NewEncoder(os.Stderr)
	encoder.Indent("", "  ")
	encoder.Encode(obj)
	os.Stderr.WriteString("\n")
}

func dumpReqest(req *http.Request)  {
	var b bytes.Buffer

	// By default, print out the unmodified req.RequestURI, which
	// is always set for incoming server requests. But because we
	// previously used req.URL.RequestURI and the docs weren't
	// always so clear about when to use DumpRequest vs
	// DumpRequestOut, fall back to the old way if the caller
	// provides a non-server Request.
	reqURI := req.RequestURI
	if reqURI == "" {
		reqURI = req.URL.RequestURI()
	}

	fmt.Fprintf(&b, "%s %s HTTP/%d.%d\r\n", req.Method, reqURI, req.ProtoMajor, req.ProtoMinor)

	absRequestURI := strings.HasPrefix(req.RequestURI, "http://") || strings.HasPrefix(req.RequestURI, "https://")
	if !absRequestURI {
		host := req.Host
		if host == "" && req.URL != nil {
			host = req.URL.Host
		}
		if host != "" {
			fmt.Fprintf(&b, "Host: %s\r\n", host)
		}
	}

	req.Header.WriteSubset(&b, map[string]bool{})
	io.WriteString(os.Stdout, b.String())
}