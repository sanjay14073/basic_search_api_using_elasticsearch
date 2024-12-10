package serach_controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Book struct {
	ID          string   `json:"id"`
	Title       string   `json:"title"`
	Author      string   `json:"author"`
	Categories  string `json:"categories"`
	Year        string   `json:"year"`
	Edition     string   `json:"edition"`
	Language    string   `json:"language"`
	Extension   string   `json:"extension"`
	Pages       string      `json:"pages"`
	Filesize    string      `json:"filesize"`
	MD5         string   `json:"md5"`
	URL         string   `json:"url"`
	ConvertTo   map[string]string `json:"convertTo"`
	Description string   `json:"description"`
	Cover      string   `json:"cover"`
}

type Response struct {
	Response []Book `json:"response"`
}

func AddData() gin.HandlerFunc{
	return func(ctx*gin.Context){
		err := godotenv.Load()
		if err != nil {
		  log.Fatal("Error loading .env file")
		}
	
		
		typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
				CloudID:os.Getenv("CLOUD_ID"),
				APIKey:os.Getenv("API_KEY"),
		})
		if err!=nil{
			log.Fatal(err);
		}
		data, err := os.ReadFile("data.json")
        if err != nil {
                fmt.Println("Error reading data.json:", err)
                return
        }
		var response Response
        err = json.Unmarshal(data, &response)
        if err != nil {
                fmt.Println("Error unmarshaling JSON:", err)
                return
        }

		for _, book := range response.Response {
			_, err := typedClient.Index("my_index").
					Id(book.ID).
					Document(book). 
					Do(context.Background())
			if err != nil {
					panic(err)
			}
		}


			ctx.JSON(http.StatusOK,gin.H{"message":"Working fine"})
		}
}

func GetData()gin.HandlerFunc{
	return func(ctx *gin.Context) {
		err := godotenv.Load()
		if err != nil {
		  log.Fatal("Error loading .env file")
		}
	
		
		typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
				CloudID:os.Getenv("CLOUD_ID"),
				APIKey:os.Getenv("API_KEY"),
		})
		if err!=nil{
			log.Fatal(err);
		}
		id:= ctx.DefaultQuery("id","1159142")
		y,err:=typedClient.Get("my_index", id).Do(context.TODO())
		if err!=nil{
			log.Fatal(err);
		}
		ctx.JSON(http.StatusOK,gin.H{"data":y})
	}
}