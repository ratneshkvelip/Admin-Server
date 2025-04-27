package services
import(
	"strings"
	)

type Row struct{
	Title string `json:"title"`
	Content []string `json:"content"`
}
	
type payload1 struct {
	Data []Row `json:"data"`
}

type rsp1 struct {
	Status  string  `json:"status"`
	Payload payload1 `json:"payload"`
}

type response struct{
	Rsp rsp1 `json:"rsp"`
}

	
func Data1( ) (int, interface{}) {

	text := "apple,banana,mango,cashew,lemon,papaya"
	text2 :="dog,cat,cow,lion,tiger,deer,monkey,giraff,fox,wolf"
	parts := strings.Split(text, ",")
	parts2:= strings.Split(text2, ",")

	resp := response{}

	resp.Rsp.Status="Ok"

	var data []Row

	for _, value := range parts {
		temp:= Row{}
		temp.Title=value
		temp.Content=parts2
		data = append(data, temp)
	}

	resp.Rsp.Payload.Data=data
	

	

	return 200, resp

}
