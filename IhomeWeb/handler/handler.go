package handler

import (
	"context"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	GETAREA "gomicro/GetArea/proto/example"
	"gomicro/IhomeWeb/models"
	"net/http"
)

func GetArea(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//创建服务获取句柄
	server := grpc.NewService()
	//服务初始化
	server.Init()
	//调用服务返回句柄
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea", server.Client())
	//调用服务返回数据
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//接收数据
	//准备接收切片
	area_list := []models.Area{}
	//循环接收数据
	for _, value := range rsp.Data {
		tmp := models.Area{Id: int(value.Aid), Name: value.Aname}
		area_list = append(area_list, tmp)
	}

	//返回给前端的map
	response := map[string]interface{}{
		"errno":  rsp.Error,
		"errmsg": rsp.Errmsg,
		"data":   area_list,
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
