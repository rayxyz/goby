package httpx

import (
	"net/http"
	"strings"

	"goby/pkg/dict"
	"goby/pkg/network"

	"golang.org/x/net/context"

	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
)

//
const (
	AccessTokenHeader               = "Ahz-Access-Token"
	LanguageHeader                  = "X-Ahz-Lang"
	HostkeyHeader                   = "X-Ahz-Hostkey"
	InternalCommAuthSignatureHeader = "Ahz-Internal-Comm-Signature"
	InternalCommAuthPayloadHeader   = "Ahz-Internal-Comm-Payload"
)

// Router :
type Router struct {
	*mux.Router
	CustomHandler                     http.Handler
	prefix                            string                                                                              // prefix of r path pattern
	checkAuth                         bool                                                                                // enable or disable authetication for accessing endpoints
	AuthCheckingMethod                func(ctx context.Context, accessToken string, permissions ...string) (bool, string) // Authetication method
	internalComm                      bool                                                                                // Are endpoints for internal communication
	AuthCheckingMethodForInternalComm func(ctx context.Context, signature string, payload string) bool                    // Authentication checking for internal service API calling
	PathList                          []string                                                                            // r path pattern list
	methods                           map[string]string
	AOPFuncs                          []func(context.Context, *Request) (bool, error) // AOP functions
}

// Group :
// internal: internal communication, if set true, and there is an `AuthCheckingInternal` auth checking method
// provided, it will check authentication.
func (rt *Router) Group(prefix string, checkAuth bool, internal bool, fns ...func()) {
	rt.checkAuth = checkAuth
	rt.prefix = prefix
	rt.internalComm = internal
	for _, fn := range fns {
		fn()
	}
}

// Wrap :
func (rt *Router) wrap(ctx context.Context, checkAuth bool, internalComm bool, handler func(ctx context.Context, req *Request, resp *Response) error, permissions ...string) func(w http.ResponseWriter, r *http.Request) {
	h := func(w http.ResponseWriter, r *http.Request) {
		req := &Request{
			Request: r,
		}
		resp := &Response{
			Status: http.StatusOK,
			Writer: w,
		}
		if checkAuth {
			if rt.AuthCheckingMethod == nil {
				log.Error("No authentication method set")
				resp.Status = http.StatusUnauthorized
				resp.Message = "Checking authentication, but no authentication method set"
				resp.WriteJSON(nil)
				return
			}
			if authorized, msg := rt.AuthCheckingMethod(ctx, r.Header.Get(AccessTokenHeader), permissions...); !authorized {
				resp.Status = http.StatusUnauthorized
				if strings.EqualFold(msg, dict.Blank) {
					resp.Message = strings.Join([]string{AccessTokenHeader, "is not present or expired"}, dict.Space)
				} else {
					resp.Message = msg
				}
				log.Error(resp.Message)
				resp.WriteJSON(nil)
				return
			}
		}
		if internalComm {
			if rt.AuthCheckingMethodForInternalComm == nil {
				log.Error("No internal communication authentication method set")
				http.Error(w, http.StatusText(http.StatusUnauthorized)+
					" <-> Checking internal communication authentication, but no authentication method set", http.StatusUnauthorized)
				return
			}
			if authorized := rt.AuthCheckingMethodForInternalComm(ctx, r.Header.Get(InternalCommAuthSignatureHeader),
				r.Header.Get(InternalCommAuthPayloadHeader)); !authorized {
				http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
				return
			}
		}
		if err := handler(ctx, req, resp); err != nil || resp.Status != http.StatusOK {
			if err != nil && resp.Message == "" {
				resp.Message = err.Error()
			}
			resp.handleErrrorStatuses(w)
		} else {
			if !resp.hasWritten {
				if resp.Stream {
					resp.hasWritten = true
					return
				}
				resp.WriteJSON(nil)
			}
		}
		// Execute AOP functions
		if len(rt.AOPFuncs) > 0 {
			go func() {
				for _, fn := range rt.AOPFuncs {
					fn(ctx, req)
				}
			}()
		}
	}
	return h
}

// Handle :
func (rt *Router) Handle(path string, handler func(ctx context.Context, req *Request, resp *Response) error, httpMethod string, permissions ...string) {
	path = rt.prefix + path
	for _, v := range rt.PathList {
		if strings.EqualFold(path, v) {
			log.Fatal("duplicated routing path pattern found => " + v)
		}
	}
	rt.PathList = append(rt.PathList, path)
	rt.Router.HandleFunc(path, rt.wrap(context.Background(), rt.checkAuth, rt.internalComm, handler, permissions...)).Methods(httpMethod)
}

// Get : HTTP GET
func (rt *Router) Get(path string, handler func(ctx context.Context, req *Request, resp *Response) error, permissions2Check ...string) {
	rt.Handle(path, handler, GET, permissions2Check...)
}

// Post : HTTP POST
func (rt *Router) Post(path string, handler func(ctx context.Context, req *Request, resp *Response) error, permissions2Check ...string) {
	rt.Handle(path, handler, POST, permissions2Check...)
}

// Put : HTTP PUT
func (rt *Router) Put(path string, handler func(ctx context.Context, req *Request, resp *Response) error, permissions2Check ...string) {
	rt.Handle(path, handler, PUT, permissions2Check...)
}

// Delete : HTTP DELETE
func (rt *Router) Delete(path string, handler func(ctx context.Context, req *Request, resp *Response) error, permissions2Check ...string) {
	rt.Handle(path, handler, DELETE, permissions2Check...)
}

// GetAvailablePort :
func (rt *Router) GetAvailablePort() int {
	port, err := network.ObtainAvailablePort(8000, 10000)
	if err != nil {
		log.Fatal(err)
	}

	return port
}
