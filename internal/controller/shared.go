package controller

import (
	"chronosphere.io/chronosphere-operator/chronosphere_client"
	"github.com/go-openapi/errors"
	httptransport "github.com/go-openapi/runtime/client"
	"net/http"
	"os"
)

func GetChronosphereClient() (chronosphere_client.ChronosphereClient, error) {
	gatewayAddress := os.Getenv("GATEWAY_ADDRESS")
	if gatewayAddress == "" {
		return chronosphere_client.ChronosphereClient{}, errors.New(500, "GATEWAY_ADDRESS is not set")
	}
	apiToken := os.Getenv("API_TOKEN")
	if apiToken == "" {
		return chronosphere_client.ChronosphereClient{}, errors.New(500, "API_TOKEN is not set")
	}
	customTransport := httptransport.New(gatewayAddress, chronosphere_client.DefaultBasePath, chronosphere_client.DefaultSchemes)
	customTransport.Transport = &CustomTransport{
		rt: http.DefaultTransport,
	}
	customTransport.DefaultAuthentication = httptransport.BearerToken(apiToken)
	client := chronosphere_client.New(customTransport, nil)
	return *client, nil
}

type CustomTransport struct {
	rt http.RoundTripper
}

func (c *CustomTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	uiLockResources := os.Getenv("UI_LOCK_RESOURCES")
	if uiLockResources == "true" {
		req.Header.Set("User-Agent", "chrono-k8s-operator/v0.43.0-a1dab007")
	}
	return c.rt.RoundTrip(req)
}
