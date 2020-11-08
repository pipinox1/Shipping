package restclient


import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type RestClient struct {
	HTTPClient *http.Client
}

func NewRestClient(time time.Duration) *RestClient {
	return &RestClient{
		HTTPClient: &http.Client{
			Timeout: time,
		},
	}
}

func (r *RestClient)DoGet(ctx context.Context,url string,response interface{})error{
	res,err := r.HTTPClient.Get(url)
	if err != nil {
		return err
	}

	//Nos Aseguramos que cierre el body
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return err
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return nil
}

