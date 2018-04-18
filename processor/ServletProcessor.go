package processor

import (
	"fmt"
    "github.com/godlet-cn/Godlet/request"
    "github.com/godlet-cn/Godlet/response"
)


type ServletProcessor1 struct{

}

func(processor ServletProcessor1) Process (request request.Request , response response.Response) {
	fmt.Println("Handle servlet request");
	response.SendResponse();
}