package main

import (
	"golang-restful-api/app"
)

func main() {
	//handleRequests()

	app := &app.App{}
	app.Init()
	app.Run(":9999")
}

//func handleRequests() {
//	http.HandleFunc("/", homePage)
//	http.HandleFunc("/students", getAllStudent)
//	log.Fatal(http.ListenAndServe(":10000", nil))
//}
//
//func homePage(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Welcome to the HomePage!")
//	fmt.Println("Endpoint Hit: homePage")
//}
//
//func getAllStudent(w http.ResponseWriter, r *http.Request) {
//	Students := []model.Student{
//		{Id: 1, FullName: "Student 1"},
//		{Id: 2, FullName: "Student 2"},
//		{Id: 3, FullName: "Student 3"},
//	}
//	fmt.Println("Endpoint Hit: returnAllArticles")
//	json.NewEncoder(w).Encode(Students)
//}
