# GoRestAPI

This repository contains the basic implementation of RestAPI using go.

# To run this project you require the following tools and configration setting:-
1. GoLand 2019.1.3
2. go 1.11.3
3. Google Postman Native App

RestAPI is implemented using 2 methods:
1. Using "http/net".
2. Using mux router.

***To run restAPI using http method:->***
-> Navigate to GoRestAPI -> http method, then run main.go
-> Go to URL localhost:8081 to check if you hit the endpoint or not

***To run restAPI using mux router:->***
-> Navigate to GoRestAPI -> mux router, then run main.go
-> Open URL localhost:8082/{Handle func} in google's postman native app to send "Get", "Post", "Delete" and "Put" requests
with their respective handle functions in URL
