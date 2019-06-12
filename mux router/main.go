package main
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
)

type Article struct {
	Id string `json:"Id"`
	Title string `json:"Title"`
	Desc string `json:"Desc"`
	Content string `json:"Content"`
}

type Articles []Article


var articles = Articles{
Article{Id:"1", Title:"Test Title", Desc:"Test Description", Content:"Test Content: Hello World" },
Article{Id:"2", Title:"Test Title 2", Desc:"Test Description 2", Content:"Test Content: Hello Again"},
}

func allArticles(w http.ResponseWriter, r*http.Request) {

	fmt.Println("EndPoint Hit: All Articles EndPoint")

	json.NewEncoder(w).Encode(articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["Id"]

	// Loop over all of our Articles
	// if the article.Id equals the key we pass in
	// return the article encoded as JSON
	for _, article := range articles{
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}
func createNewArticle(w http.ResponseWriter, r *http.Request){
	//Read the body of post request
	reqBody, _ := ioutil.ReadAll(r.Body)

	var article Article
	json.Unmarshal(reqBody, &article)
	// update our global Articles array to include
	// our new Article
	articles = append(articles, article)

	json.NewEncoder(w).Encode(article)

}

func deleteArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["Id"]

	//Loop through all the articles
	for index, article := range articles {
		if article.Id == key{
			articles = append(articles[:index], articles[index+1:]...)
		} // If current articles Id matches the key Id to delete
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	key := vars["Id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	for index , article := range articles {
		if article.Id == key{
			json.Unmarshal(reqBody, &article)
			articles[index] = article
			json.NewEncoder(w).Encode(article)
		}
	}
}

func homePage (w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "HomePage endpoint hit")
}

func handleRequests(){
	//create a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	//replace http.handleFunc myROuter.handleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", allArticles)
	myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{Id}", deleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/update/{Id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/article/{Id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8082", myRouter))
}

func main() {
	fmt.Println("Rest API v2.0 - Mux Routers")
	handleRequests()
}