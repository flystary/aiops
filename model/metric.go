package model

import (
	"fmt"
	"github.com/flystary/aiops/utils"
)


type MetricValue struct {
	Endpoint  string      `json:"endpoint"`
	Metric    string      `json:"metric"`
	Value     interface{} `json:"value"`
	Step      int64       `json:"step"`
	Type      string      `json:"counterType"`
	Tags      string      `json:"tags"`
	Timestamp int64       `json:"timestamp"`
}

func (m *MetricValue) String() string {
	return fmt.Sprintf(
		"<Endpoint:%s, Metric:%s, Type:%s, Tags:%s, Step:%d, Time:%d, Value:%v>",
		m.Endpoint,
		m.Metric,
		m.Type,
		m.Tags,
		m.Step,
		m.Timestamp,
		m.Value,
	)
}

type JsonMetaData struct {
	Metric      string      `json:"metric"`
	Endpoint    string      `json:"endpoint"`
	Timestamp   int64       `json:"timestamp"`
	Step        int64       `json:"step"`
	Value       interface{} `json:"value"`
	CounterType string      `json:"counterType"`
	Tags        string      `json:"tags"`
}

func (t *JsonMetaData) String() string {
	return fmt.Sprintf("<JsonMetaData Endpoint:%s, Metric:%s, Tags:%s, DsType:%s, Step:%d, Value:%v, Timestamp:%d>",
		t.Endpoint, t.Metric, t.Tags, t.CounterType, t.Step, t.Value, t.Timestamp)
}

type MetaData struct {
	Metric      string            `json:"metric"`
	Endpoint    string            `json:"endpoint"`
	Timestamp   int64             `json:"timestamp"`
	Step        int64             `json:"step"`
	Value       float64           `json:"value"`
	CounterType string            `json:"counterType"`
	Tags        map[string]string `json:"tags"`
}

func (t *MetaData) String() string {
	return fmt.Sprintf("<MetaData Endpoint:%s, Metric:%s, Timestamp:%d, Step:%d, Value:%f, Tags:%v>",
		t.Endpoint, t.Metric, t.Timestamp, t.Step, t.Value, t.Tags)
}

func (t *MetaData) PK() string {
	return utils.PK(t.Endpoint, t.Metric, t.Tags)
}
