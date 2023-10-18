package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {

	router := gin.Default()

	router.GET("/login", func(context *gin.Context) {
		//redirect to strava
		context.Redirect(302, "https://www.strava.com/oauth/authorize?client_id=114422&redirect_uri=http%3A%2F%2Flocalhost%3A8080%2FstravaRedirected&response_type=code&scope=profile%3Aread_all")
	})

	router.GET("/stravaRedirected", func(context *gin.Context) {
		//get code from strava
		code := context.Query("code")

		newUri := "https://www.strava.com/oauth/token"
		newUri += "?client_id=114422"
		newUri += "&client_secret=462e590fb0b4352d092b3ada5f3a00fa5e9f49a3&code=4107d1a272e89646f76428201f4c9fbce685ce3f"
		newUri += "&grant_type=authorization_code"
		newUri += "&code=" + code

		r, err := http.Post(newUri, "application/json", nil)

		if err != nil {
			context.JSON(400, gin.H{"error": "Bad request"})
			println(err.Error())
			return
		}

		var jsonBody map[string]interface{}

		err = json.NewDecoder(r.Body).Decode(&jsonBody)

		if err != nil {
			context.JSON(400, gin.H{"error": "Bad request"})
			println(err.Error())
			return
		}

		println(jsonBody)
		context.JSON(200, jsonBody)
	})

	err := router.Run(":8080")
	if err != nil {
		println(err.Error())
		return
	}
}
