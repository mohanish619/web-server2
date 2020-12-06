
package main
import (
"fmt"
	"io/ioutil"
	"net/http"
"github.com/gorilla/mux"
)

func main(){
	r :=mux.NewRouter()
	r.HandleFunc("/",home).Methods("GET")
	r.HandleFunc("/",shrikar).Methods("POST")
	r.HandleFunc("/mohanish",mahi).Methods("GET")
	http.ListenAndServe(":8080",r)
}
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hello World!")
}
func mahi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"i am mohanish!")
}
func shrikar(w http.ResponseWriter, r *http.Request) {
	a,err := ioutil.ReadAll(r.Body)
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Fprint(w,"i am shrikar",string(a))
}