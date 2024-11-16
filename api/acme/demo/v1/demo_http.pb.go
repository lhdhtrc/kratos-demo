// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.8.2
// - protoc             (unknown)
// source: acme/demo/v1/demo.proto

package demov1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDemoServiceCreate = "/acme.demo.v1.DemoService/Create"
const OperationDemoServiceDelete = "/acme.demo.v1.DemoService/Delete"
const OperationDemoServiceGet = "/acme.demo.v1.DemoService/Get"
const OperationDemoServiceGetList = "/acme.demo.v1.DemoService/GetList"
const OperationDemoServiceUpdate = "/acme.demo.v1.DemoService/Update"

type DemoServiceHTTPServer interface {
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	Get(context.Context, *GetRequest) (*GetResponse, error)
	GetList(context.Context, *GetListRequest) (*GetListResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
}

func RegisterDemoServiceHTTPServer(s *http.Server, srv DemoServiceHTTPServer) {
	r := s.Route("/")
	r.POST("/v1/demo", _DemoService_Create0_HTTP_Handler(srv))
	r.GET("/v1/demo/{id}", _DemoService_Get0_HTTP_Handler(srv))
	r.GET("/v1/demo/list", _DemoService_GetList0_HTTP_Handler(srv))
	r.PUT("/v1/demo", _DemoService_Update0_HTTP_Handler(srv))
	r.DELETE("/v1/demo/{id}", _DemoService_Delete0_HTTP_Handler(srv))
}

func _DemoService_Create0_HTTP_Handler(srv DemoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in CreateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoServiceCreate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Create(ctx, req.(*CreateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*CreateResponse)
		return ctx.Result(200, reply)
	}
}

func _DemoService_Get0_HTTP_Handler(srv DemoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoServiceGet)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Get(ctx, req.(*GetRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetResponse)
		return ctx.Result(200, reply)
	}
}

func _DemoService_GetList0_HTTP_Handler(srv DemoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetListRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoServiceGetList)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetList(ctx, req.(*GetListRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetListResponse)
		return ctx.Result(200, reply)
	}
}

func _DemoService_Update0_HTTP_Handler(srv DemoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoServiceUpdate)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Update(ctx, req.(*UpdateRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateResponse)
		return ctx.Result(200, reply)
	}
}

func _DemoService_Delete0_HTTP_Handler(srv DemoServiceHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in DeleteRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoServiceDelete)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Delete(ctx, req.(*DeleteRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*DeleteResponse)
		return ctx.Result(200, reply)
	}
}

type DemoServiceHTTPClient interface {
	Create(ctx context.Context, req *CreateRequest, opts ...http.CallOption) (rsp *CreateResponse, err error)
	Delete(ctx context.Context, req *DeleteRequest, opts ...http.CallOption) (rsp *DeleteResponse, err error)
	Get(ctx context.Context, req *GetRequest, opts ...http.CallOption) (rsp *GetResponse, err error)
	GetList(ctx context.Context, req *GetListRequest, opts ...http.CallOption) (rsp *GetListResponse, err error)
	Update(ctx context.Context, req *UpdateRequest, opts ...http.CallOption) (rsp *UpdateResponse, err error)
}

type DemoServiceHTTPClientImpl struct {
	cc *http.Client
}

func NewDemoServiceHTTPClient(client *http.Client) DemoServiceHTTPClient {
	return &DemoServiceHTTPClientImpl{client}
}

func (c *DemoServiceHTTPClientImpl) Create(ctx context.Context, in *CreateRequest, opts ...http.CallOption) (*CreateResponse, error) {
	var out CreateResponse
	pattern := "/v1/demo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDemoServiceCreate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DemoServiceHTTPClientImpl) Delete(ctx context.Context, in *DeleteRequest, opts ...http.CallOption) (*DeleteResponse, error) {
	var out DeleteResponse
	pattern := "/v1/demo/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoServiceDelete))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "DELETE", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DemoServiceHTTPClientImpl) Get(ctx context.Context, in *GetRequest, opts ...http.CallOption) (*GetResponse, error) {
	var out GetResponse
	pattern := "/v1/demo/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoServiceGet))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DemoServiceHTTPClientImpl) GetList(ctx context.Context, in *GetListRequest, opts ...http.CallOption) (*GetListResponse, error) {
	var out GetListResponse
	pattern := "/v1/demo/list"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoServiceGetList))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func (c *DemoServiceHTTPClientImpl) Update(ctx context.Context, in *UpdateRequest, opts ...http.CallOption) (*UpdateResponse, error) {
	var out UpdateResponse
	pattern := "/v1/demo"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationDemoServiceUpdate))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, nil
}
