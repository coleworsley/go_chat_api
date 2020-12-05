package main

import (
    // "fmt"
    // "log"
    // "encoding/json"

    "github.com/gin-gonic/gin"
  	"github.com/swaggo/files"
  	"github.com/swaggo/gin-swagger"

    // "./src"
    // "./controller"

    // My wish is to use swag + gin to gen my docs.
    // the generation of the docs actually works:
    // `swag init`
    // The issue I'm seeing is in docs/docs.go one of the packages is unable to be downloaded.
    // `docs/docs.go:11:2: cannot find package`
    // I've tried:
    // `go get github.com/alecthomas/template`
    // which should be adding the package to the go.mod file and downloading the required packages.
    // clearly, it's not working and I need to keep making progress on the init project.
  	// "./docs" // docs is generated by Swag CLI, you have to import it.
)


type User struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Users array
// that we can then populate in our main function
// to simulate a database
var Users []User


// func returnAllUsers(w http.ResponseWriter, r *http.Request){
  // fmt.Println("Endpoint Hit: returnAllArticles")
  // json.NewEncoder(w).Encode(Users)
// }

func handleRequests() {
    r := gin.Default()

  	c := controller.NewController()

  	v1 := r.Group("/api/v1")
  	{
  		accounts := v1.Group("/accounts")
  		{
  			accounts.GET(":id", c.ShowAccount)
  			accounts.GET("", c.ListAccounts)
  			accounts.POST("", c.AddAccount)
  			accounts.DELETE(":id", c.DeleteAccount)
  			accounts.PATCH(":id", c.UpdateAccount)
  			accounts.POST(":id/images", c.UploadAccountImage)
  		}
      //...
  	}
  	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
  	r.Run(":3050")




    // r := gin.Default()
  	// r.GET("/ping", func(c *gin.Context) {
  	// 	c.JSON(200, gin.H{
  	// 		"message": "pong",
  	// 	})
  	// })
  	// r.Run(":3050")



    // myRouter := mux.NewRouter().StrictSlash(true)
    // myRouter.HandleFunc("/", homePage)
    // myRouter.HandleFunc("/all", returnAllArticles)
    //
    // http.HandleFunc("/", homePage)
    // http.HandleFunc("/user", returnAllUsers)
    // log.Fatal(http.ListenAndServe(":10000", nil))
}


// @version 1.0

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

func main() {
    // programmatically set swagger info
    // docs.SwaggerInfo.Title = "Swagger Example API"
    // docs.SwaggerInfo.Description = "This is a sample server Petstore server."
    // docs.SwaggerInfo.Version = "1.0"
    // docs.SwaggerInfo.Host = "petstore.swagger.io"
    // docs.SwaggerInfo.BasePath = "/v2"
    // docs.SwaggerInfo.Schemes = []string{"http", "https"}


    Users = []User{
      User{Title: "Hello", Desc: "User Description", Content: "User Content"},
      User{Title: "Hello 2", Desc: "User Description", Content: "User Content"},
    }
    handleRequests()
}
