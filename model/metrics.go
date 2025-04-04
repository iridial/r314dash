package model

type MetricResponse struct {
	Status string `json:"status"`
	Data   struct {
		Result []struct {
			Metric map[string]string `json:"metric"`
			//Values [][]interface{}   `json:"values"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

type MetricValue struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
	Warn  bool   `json:"warn"`
	Label string `json:"label"`
}

type InterfaceMetrics struct {
	Interface   string      `json:"interface"`
	Received    MetricValue `json:"received"`
	Transmitted MetricValue `json:"transmitted"`
}

type CpuMetrics struct {
	Usage  MetricValue `json:"usage"`
	Load1  MetricValue `json:"load1"`
	Load5  MetricValue `json:"load5"`
	Load15 MetricValue `json:"load15"`
}

type MemoryMetrics struct {
	Usage MetricValue `json:"usage"`
	Total MetricValue `json:"total"`
	Free  MetricValue `json:"free"`
	Cache MetricValue `json:"cache"`
}

type StorageMetrics struct {
	Name  string      `json:"name"`
	Usage MetricValue `json:"usage"`
	Free  MetricValue `json:"free"`
	Used  MetricValue `json:"used"`
}

type StorageMetricsCouple struct {
	Top StorageMetrics `json:"top"`
	Bot StorageMetrics `json:"bot"`
}

type MetricsCouple struct {
	Top  MetricValue `json:"top"`
	Bot  MetricValue `json:"bot"`
	Show bool        `json:"show"`
}

type SystemMetrics struct {
	Interfaces  []InterfaceMetrics     `json:"interfaces"`
	Cpu         CpuMetrics             `json:"cpu"`
	Memory      MemoryMetrics          `json:"memory"`
	Storage     []StorageMetricsCouple `json:"storage"`
	Temperature MetricsCouple          `json:"temperature"`
}
