package restclient

import (
	"apitest/internal/tools/errors"
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type RestClient struct {
	HTTPClient *http.Client
}

func NewRestClient(time time.Duration) RestClient {
	return RestClient{
		HTTPClient: &http.Client{
			Timeout: time,
		},
	}
}

func (r *RestClient) DoGet(ctx context.Context, url string, response interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", ctx.Value("user_logged").(string))

	res, err := r.HTTPClient.Do(req)
	if err != nil {
		return errors.NewRestError("rest_client_error", http.StatusServiceUnavailable)
	}

	//Nos Aseguramos que cierre el body
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		return errors.NewRestError("error_reading_body", http.StatusInternalServerError)
	}

	if res.StatusCode < http.StatusOK || res.StatusCode >= http.StatusBadRequest {
		return errors.NewRestError("rest_client_error", res.StatusCode)
	}

	return nil
}
