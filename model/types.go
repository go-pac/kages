package model

type (
	Array     []Object
	ArrayCSV  []CSV
	ArrayJson []Json
	ArrayKV   []KV
	CSV       Json
	Json      map[string]interface{}
	KV        map[interface{}]interface{}
	Object    interface{}
	Yaml      Json
)
