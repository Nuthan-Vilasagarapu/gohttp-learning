package main

import "net/http"

func FormHandler(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("s_name")
	message := name + ", thank you filling this form!"
	res.Write([]byte(message))
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/form", FormHandler)
	http.ListenAndServe(
		":8000",
		nil,
	)
}
