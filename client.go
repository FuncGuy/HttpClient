package main

type httpClient struct{}

func New() HttpClient {
	client := &httpClient{}
	return client
}

type HttpClient interface {
	Get()
	Post()
	Patch()
	Delete()
}

func (c *httpClient) Get() { println("you called GET method") }

func (c *httpClient) Post() {}

func (c *httpClient) Put() {}

func (c *httpClient) Patch() {}

func (c *httpClient) Delete() {}
