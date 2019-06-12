package http_method
import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Student struct{
	Name string `json:"Name"`
	Age int `json:"Age"`
}

type Articles []Article
type Students []Student

func allStudents(w http.ResponseWriter, r *http.Request){
	students := Students{
		Student{Name:"Seedak", Age:20},
	}
	json.NewEncoder(w).Encode(students)
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title:"Test Title", Desc:"Test Description", Content:"Test Content: Hello World" },
	}

	fmt.Println("EndPoint Hit: All Articles EndPoint")

	json.NewEncoder(w).Encode(articles)
}


func homePage (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "HomePage endpoint hit")
}

func handleRequest(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/students", allStudents)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main(){
	handleRequest()
}