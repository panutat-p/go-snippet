package profiling

import ( 
	"log"
	"net/http"
	_ "net/http/pprof"
)

// https://medium.com/@ravikumarray92/profiling-in-go-with-pprof-e45656df033e
// pprof will expose endpoints
// http://localhost:8080/debug/pprof
func main() { 
	log.Println(“booting on localhost:8080”) 
	log.Fatal(http.ListenAndServe(“:8080”, nil)) 
}
