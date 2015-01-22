package main
import (
	"github.com/coocood/jas"
	"fmt"
//	"html/template"
	"net/http"
//	"strings"
//	"log"
)
type Hello struct {}
func (*Hello) Get (ctx *jas.Context) { // `GET /v1/hello`
    ctx.Data = "hello world ,what are you doing?"
    //response: `{"data":"hello world","error":null}`
}
type Receiver struct{
    type_name string
    freqs []string

}
func (* Receiver) Get(ctx *jas.Context){
    ctx.Data = "you get a receiver!\n"
}
type Radio struct{
    type_name string
    produce_time string
    serve_time      string
    receiver Receiver
    count int32
}
type Radios struct {
    radio map[string]Radio
}
//var count  int = 0
func (r * Radio) Get(ctx * jas.Context){
	ctx.Data = "radio response"
	r.count =r.count+1
	if r.count>2{
		ctx.Data = "more response"
	}
}
func (*Radio) Post(ctx * jas.Context){
	ctx.Data = "POST sucessful"
}
func (* Radio) GetReceiver(ctx * jas.Context){
    ctx.Data = "receiver of Radio"
}
func (* Radio) GetReceiverTypename(ctx * jas.Context){
    ctx.Data = "receiver type name "
}
func (* Radio) Gap(ctx * jas.Context) string {
    return ":name"
}

func (* Radio) PostReceiverId(ctx * jas.Context) {
    id := ctx.Id
    ctx.Data= "post Receiver id "+string(id)
}
func (*Radio) Typename(ctx * jas.Context){
    name := ctx.GapSegment("")
    ctx.Data = string(name)+"type name"
    fmt.Println(name)
}
type RadioId struct{}
func (* RadioId) GetReceiver(ctx * jas.Context){
    id := ctx.Id
    _ = id
    ctx.Data = "ID  receiver "+string(id)
}

func (* RadioId) PostReceiverId(ctx * jas.Context) {
    id := ctx.Id
    ctx.Data= "post Receiver id "+string(id)
}
////////////////main
func main () {
    router := jas.NewRouter(new(Hello),new(Radio),new(RadioId))
    router.BasePath = "/"
    fmt.Println(router.HandledPaths(true))
    //output: `GET /v1/hello`
    http.Handle(router.BasePath, router)
    /*
    radio_router := jas.NewRouter(new(Radio))
    radio_router.BasePath = "/"
//    http.Handle("/radio",jas.NewRouter(new(Radio)))
    http.Handle(radio_router.BasePath,radio_router)
    */
    http.ListenAndServe(":8080", nil)
}
