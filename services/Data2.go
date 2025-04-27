package services
import(
	"strings"
	)


type payload2 struct {
	Title string `json:"title"`
	Content []string `json:"content"`
}

type rsp2 struct {
	Status  string  `json:"status"`
	Payload payload2 `json:"payload"`
}

type response2 struct{
	Rsp rsp2 `json:"rsp"`
}

	
func Data2( ) (int, interface{}) {

	text2 :="shark,whale,crab,dolphin,turtle,squids,prawns,kingfish,octopus"
	parts2:= strings.Split(text2, ",")

	resp := response2{}

	resp.Rsp.Status="Ok"

	resp.Rsp.Payload.Content=parts2
	resp.Rsp.Payload.Title="Aquatic Animal"
	

	

	return 200, resp

}
