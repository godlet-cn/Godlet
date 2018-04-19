package processor

import (
	"fmt"
	"github.com/godlet-cn/Godlet/request"
	"github.com/godlet-cn/Godlet/response"
)

type StaticResourceProcessor struct {
}

func (processor StaticResourceProcessor) Process(request request.Request, response response.Response) {
	fmt.Println("Handle static resource request")
	response.SendResponse()
}
