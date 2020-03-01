package main

import (
	"encoding/json"
	//"zcm_crm_zj/models/crm"

	"github.com/astaxie/beego"
)

// -------------------
//      Interface
// -------------------

func main() {
	realData := `[{"response_message":"123:testmess","status":"01","overall_status":"01","payment_amount":"100","app_seq":"12112121212","payment_time":"2017-06-16 10:16:00"}]`
	var m []ZJPayResponseParam
	json.Unmarshal([]byte(realData), &m)
	beego.Debug("m=", []byte(realData))

}

// 放款接口返回参数
type ZJPayResponseParam struct {
	App_seq          string  `json:"app_seq"`          //申请流水号
	Overall_status   string  `json:"overall_status"`   //  放款总状态 一个贷款申请的放款可能包括几个步骤，此处是这些步骤的总的状态。00：未做01：成功02：处理中03：失败04：等候
	Payment_time     string  `json:"payment_time"`     //放款时间
	Payment_no       int     `json:"payment_no"`       //放款步骤序号
	Payment_amount   float64 `json:"payment_amount"`   //本步骤金额
	Status           string  `json:"status"`           //本步骤放款结果 00：未做01：成功02：处理中03：失败 04：等候
	Response_message string  `json:"response_message"` //结果消息 格式：交易流水号：消息示例：20170301115959：银行发生异常
}
