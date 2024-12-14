package models

type Metrics struct {
	Id          int `json:"id"`
	MetricId    int `json:"metric_id"`
	MetricValue int `json:"value"`
}
