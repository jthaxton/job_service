package main

type Job struct {
	ID int `json:"id"`
	CustomId string `json:"custom_id"`
	Kind string `json:"kind"`
	DataJson string `json:"data_json"`
	CreatedAt string `json:"created_at"`
}