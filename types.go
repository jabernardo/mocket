package main

type Header struct {
	Name  string `yaml:name`
	Value string `yaml:value`
}

type Endpoint struct {
	Path    string `yaml:path`
	Method  string `yaml:method`
	Status  int    `yaml:status`
	Delay   int    `yaml:delay`
	Headers []Header
	Body    interface{} `yaml:body,inline`
}

type Config struct {
	Name      string `yaml:name`
	Port      int    `yaml:port`
	Endpoints []Endpoint
}
