package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//JSON을 담을 Struct
type User struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// 인스턴스를만듦  그다음 인터페이스 구현
type fooHandler struct{}

//ServeHTTP인터페이스
func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := new(User)
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Bad Request:", err)
		return
	}
	user.CreatedAt = time.Now()
	data, _ := json.Marshal(user)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
	//fmt.Fprint(w, "hello foo!")
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		name = "world"
	}
	fmt.Fprintf(w, "hello world bar %s ", name)
}

func main() {
	mux := http.NewServeMux()
	//func 을 직접 등록, 핸들러를 func 형태로 직접등록할때 쓰는 방법.
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello , world")
	})

	mux.HandleFunc("/bar", barHandler)
	//handler 즉 인스턴스형태 Handle 사용
	mux.Handle("/foo", &fooHandler{})
	http.ListenAndServe(":3000", mux) //localhost:3000
}
