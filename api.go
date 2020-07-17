package goread

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

var developerKey string = os.Getenv("GOODREADS_DEVELOPER_KEY")

// Get makes a GET call to the goodreads api with given url and parameters
func get(url string, parameters map[string]string) ([]byte, error) {
	client := http.Client{
		Timeout: time.Duration(time.Second * 15),
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	query := request.URL.Query()
	query.Add("key", developerKey)
	for key, value := range parameters {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.Do(request)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	return body, nil
}
