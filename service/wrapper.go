package service

import (
	"fmt"
	"io/ioutil"
	"net/http"
	//"github.com/thanhlam/sso-service/config"
)

type RTokens struct {
	Accesstoken string `json:"access_token"`
	Tokentype   string `json:"token_type"`
	Expiresin   string `json:"expires_in"`
	Scope       string `json:"scope"`
}

func BasicAuth(token string) (s string, err error) {
	//config.ReadConfig()
	//link := viper.GetString(`sso.url`)
	link := "https://iotsso.vdc2.com.vn:8443/cas/oauth2.0"
	var username string = "exampleOauthClient"
	var passwd string = "exampleOauthClientSecret"
	client := &http.Client{}
	url := link + "/profile?access_token=" + token
	req, err := http.NewRequest("GET", url, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return "", err
	}
	bodyText, err := ioutil.ReadAll(resp.Body)
	s = string(bodyText)
	return s, nil
}
