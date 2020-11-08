package main

import (
	"goby/pkg/conf"
	"goby/pkg/httpx"
	"goby/pkg/template"
	api "goby/services/helloworld/api"
	"goby/services/helloworld/model"
	"path/filepath"

	"golang.org/x/net/context"
)

var (
	allSvcConfs = conf.GetAllServiceList()
	basicConf   = conf.GetBasic()
)

func indexPageHandler(ctx context.Context, req *httpx.Request, resp *httpx.Response) error {
	resp.Stream = true

	tmpl, err := template.New(
		filepath.Join(staticResourcePathWebapp, "index.html"),
		filepath.Join(staticResourcePathWebapp, "pages/header.html"),
		filepath.Join(staticResourcePathWebapp, "pages/footer.html"),
	)
	if err != nil {
		return err
	}
	err = tmpl.ExecuteTemplate(resp.GetWriter(), "index.html", template.Data{
		"page":    "index",
		"message": "Hello, world!",
	})
	if err != nil {
		return err
	}

	return nil
}

func helloHandler(ctx context.Context, req *httpx.Request, resp *httpx.Response) error {
	// var msg string
	// if err := req.ParseParams("msg", &msg); err != nil {
	// 	return err
	// }

	resp.WriteJSON("Hello, world !")

	return nil
}

func saveAdviceHandler(ctx context.Context, req *httpx.Request, resp *httpx.Response) error {
	advice := new(model.Advice)

	if err := req.ParseParams(advice); err != nil {
		resp.Status = httpx.StatusParseParamsError
		resp.Message = httpx.StatusTextWithExtra(httpx.StatusParseParamsError, err.Error())
		return nil
	}

	id, err := api.SaveAdvice(ctx, advice)
	if err != nil {
		resp.Status = httpx.StatusSaveDataError
		resp.Message = "save advice error"
		return nil
	}

	resp.WriteJSON(id)

	return nil
}

func getAdviceListHandler(ctx context.Context, req *httpx.Request, resp *httpx.Response) error {
	var start, size int64

	if err := req.ParseParams("start", &start, "size", &size); err != nil {
		start = 0
		size = 10
	}

	ovList, err := api.GetAdviceList(ctx, &model.AdviceListQueryObj{
		Start: start,
		Size:  size,
	})

	if err != nil {
		resp.Status = httpx.StatusQueryDataError
		resp.Message = "query advice list error"
		return nil
	}

	resp.WriteJSON(ovList)

	return nil
}
