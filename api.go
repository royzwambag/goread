package goread

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

// Client allows you to make api calls
type Client struct {
	http *http.Client
	key  string
}

// NewClient returns a client with a 10 second timeout and the developer key from your environment variables
func NewClient() Client {
	client := http.Client{
		Timeout: time.Duration(time.Second * 10),
	}
	return Client{
		http: &client,
		key:  os.Getenv("GOODREADS_DEVELOPER_KEY"),
	}
}

// Get makes a GET call to the goodreads api with given url and parameters
func (client *Client) get(url string, parameters map[string]string) ([]byte, error) {
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	query := request.URL.Query()
	query.Add("key", client.key)
	for key, value := range parameters {
		query.Add(key, value)
	}
	request.URL.RawQuery = query.Encode()

	response, err := client.http.Do(request)
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
