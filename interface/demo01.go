package main

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
)

type Page struct {
	Size    int64
	Current int64
}

type DBService struct {
	DB        gdb.DB
	Tenant    string
	TableName string
}

type Person struct {
	Id   int64
	Name string
}

func (p Person) List(req interface{}, authorityId int64) (result *Result, err error) {
	result = &Result{}
	ps := make([]Person, 0)
	model := p.Model(req)
	if model != nil {
		resu, err := p.Model(req).FindAll()
		if err != nil {
			return nil, err
		}
		if !resu.IsEmpty() {
			resu.Structs(&ps)
			result.DataList = ps
			return result, nil
		}
	}
	return nil, err
}
func (p Person) Model(req interface{}) (model *gdb.Model) {
	reqs := req.(*Request)
	if reqs != nil {
		person := reqs.Data.(*Person)
		if person != nil {
			model = reqs.DBService.DB.Model(p.TableName()).Safe()
		}
	}
	return model
}

func (p Person) TableName() string {
	return "person"
}

type Entity interface {
	TableName() string
	List(req interface{}, authorityId int64) (result *Result, err error)
	Model(req interface{}) *gdb.Model
}

type Request struct {
	Data interface{}
	*DBService
	*Page
}

func (r *Request) NewService(tenant string, entity Entity) (*DBService, error) {
	dbService := &DBService{
		Tenant:    tenant,
		TableName: entity.TableName(),
	}
	dbService.DB = g.DB(tenant)
	return dbService, nil
}

func (r *Request) List(req interface{}, authorityId int64, entity Entity) (result *Result, err error) {
	return entity.List(req, authorityId)
}

type Result struct {
	DataList interface{}
	Total    int64
}

/*
	API method
*/
type ServiceOperator interface {
	NewService(tenant string, object Entity) (*DBService, error)
	List(req interface{}, authorityId int64, entity Entity) (result *Result, err error)
}

// 数据返回通用JSON数据结构
type JsonResponse struct {
	Code int32       `json:"code"` // 错误码((0:成功, 1:失败, >1:错误码))
	Msg  string      `json:"msg"`  // 提示信息
	Data interface{} `json:"data"` // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装。
func ResponseJson(r *ghttp.Request, code int32, message string, data ...interface{}) {
	responseData := interface{}(nil)
	if len(data) > 0 {
		responseData = data[0]
	}
	r.Response.WriteJson(JsonResponse{
		Code: code,
		Msg:  message,
		Data: responseData,
	})
	r.Exit()
}

func GetList(r *ghttp.Request, object ServiceOperator, entity Entity) {
	// 基础参数
	// 获取租户码
	tenant := r.Header.Get("Tenant")
	// 权限
	id := r.GetInt64("id")
	// 初始化service
	service, err := object.NewService(tenant, entity)
	if err != nil {
		ResponseJson(r, 1, err.Error())
	}
	// 分页
	page := &Page{
		Size:    r.GetInt64("size"),
		Current: r.GetInt64("current"),
	}
	req := &Request{
		DBService: service,
		Page:      page,
		Data:      entity,
	}
	if err := r.Parse(&req); err != nil {
		ResponseJson(r, -1, err.Error())
	}
	glog.Info("前端请求参数;", req.Data)
	dataList, err := object.List(req, id, entity)
	if err != nil {
		ResponseJson(r, 1, err.Error())
	}
	if dataList == nil {
		ResponseJson(r, 1, "")
	}
	ResponseJson(r, 0, "", dataList)

}

/*
	TestController
*/
type TestController struct{}

func (t TestController) GetList(r *ghttp.Request) {
	GetList(r, new(Request), new(Person))
}
func main() {
	s := g.Server()
	s.BindObject("/", new(TestController))
	s.SetPort(8082)
	s.Run()
}
