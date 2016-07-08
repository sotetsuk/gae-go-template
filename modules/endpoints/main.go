package main

import (
	"github.com/GoogleCloudPlatform/go-endpoints/endpoints"
	"golang.org/x/net/context"
	"log"
)

type SampleAPi struct {
}

type Res struct {
	Res string `json:"res"`
}

type Req struct {
	ReqParam string `json:"reqparam"`
}

func (h *SampleAPi) MethodConst(c context.Context) (*Res, error) {
	return &Res{Res: "Constant Response"}, nil
}

func (h *SampleAPi) MethodWithReqParam(c context.Context, r *Req) (*Res, error) {
	return &Res{Res: "Return given param: " + r.ReqParam}, nil
}

func init() {
	sample := &SampleAPi{}
	api, err := endpoints.RegisterService(sample, "sample", "v1", "Sample API", true)
	if err != nil {
		log.Fatalf("Register service: %v", err)
	}

	register := func(orig, name, method, path, desc string) {
		m := api.MethodByName(orig)
		if m == nil {
			log.Fatalf("Missing method %s", orig)
		}
		i := m.Info()
		i.Name, i.HTTPMethod, i.Path, i.Desc = name, method, path, desc
	}

	register("MethodConst", "methodConst", "GET", "methodConst", "Return constant response")
	register("MethodWithReqParam", "methodWIthReqparam", "GET", "methodWithReqparam", "Return the given parameter")
	endpoints.HandleHTTP()
}
