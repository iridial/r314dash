package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"r314dash/config"
	"r314dash/model"
	"sort"
	"strings"
	"text/template"
)

func ConvertMetricBytes(bytes float64, time bool) model.MetricValue {
	var metric model.MetricValue
	// metric.Value = "GB"
	const (
		_  = iota
		KB = 1 << (10 * iota) // 1 << 10 = 1024
		MB
		GB
	)

	var seconds string
	if time {
		seconds = "/s"
	} else {
		seconds = ""
	}

	switch {
	case bytes >= GB:
		metric.Value = fmt.Sprintf("%.2f", float64(bytes)/GB)
		metric.Unit = fmt.Sprintf("GB%s", seconds)
	case bytes >= MB:
		metric.Value = fmt.Sprintf("%.2f", float64(bytes)/MB)
		metric.Unit = fmt.Sprintf("MB%s", seconds)
	//case bytes >= KB:
	default:
		metric.Value = fmt.Sprintf("%.2f", float64(bytes)/KB)
		metric.Unit = fmt.Sprintf("KB%s", seconds)
	}
	return metric
}

func StatsHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	reqStat := query.Get("type")
	var tmpl *template.Template
	var tmplerr error
	var data any
	var dataerr error
	switch reqStat {
	case "vpn":
		data, dataerr = CollectVpnStatus()
		tmpl, tmplerr = template.ParseFiles("templates/vpn.html")
	default:
		data, dataerr = CollectMetrics()
		tmpl, tmplerr = template.ParseFiles("templates/metrics.html")
	}
	if dataerr != nil {
		http.Error(w, dataerr.Error(), http.StatusInternalServerError)
		return
	}
	if tmplerr != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("template load error:", tmplerr.Error())
		http.Error(w, "<p>STATS TEMPLATE LOAD ERROR</p>", http.StatusInternalServerError)
		return
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		//http.Error(w, err.Error(), http.StatusInternalServerError)
		fmt.Println("template exec error:", err.Error())
		http.Error(w, "<p>STATS TEMPLATE EXEC ERROR</p>", http.StatusInternalServerError)
	}
}

func CollectVpnStatus() (model.VpnConnection, error) {
	req, err := http.NewRequest("GET", config.Main.Settings.Stats.Vpn.BackendUrl, nil)
	if err != nil {
		return model.VpnConnection{}, err
	}
	req.Header.Set("X-API-Key", config.Main.Settings.Stats.Vpn.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return model.VpnConnection{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return model.VpnConnection{}, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}
	var vpnResponse model.VpnResponse
	if err := json.NewDecoder(resp.Body).Decode(&vpnResponse); err != nil {
		return model.VpnConnection{}, err
	}

	return model.VpnConnection{
		Connected: len(vpnResponse.PublicIP) > 0,
		PublicIP:  vpnResponse.PublicIP,
		Region:    vpnResponse.Region,
		Country:   vpnResponse.Country,
		City:      vpnResponse.City,
	}, nil
}

func CollectMetrics() (model.SystemMetrics, error) {
	// TODO: auth to prometheus
	//end := time.Now().Unix()
	//start := end - 2 // match scraping time, get a single value
	//time := time.Now().Unix()
	metricsConfigMap := make([]string, len(config.Main.Settings.Stats.Metrics))
	for i := range config.Main.Settings.Stats.Metrics {
		metricsConfigMap[i] = fmt.Sprintf(
			"label_replace(%s,\"metric\",\"%s\",\"\",\"\")",
			config.Main.Settings.Stats.Metrics[i].Query,
			config.Main.Settings.Stats.Metrics[i].Label,
		)
	}
	//time=%d& / , time
	queryURL := fmt.Sprintf("%s?query=%s", config.Main.Settings.Stats.BackendUrl, url.QueryEscape(strings.Join(metricsConfigMap, " or ")))
	//fmt.Println("query", queryURL)

	resp, err := http.Get(queryURL)
	if err != nil {
		return model.SystemMetrics{}, fmt.Errorf("failed to query Prometheus")
	}
	defer resp.Body.Close()

	var metricResponse model.MetricResponse
	if err := json.NewDecoder(resp.Body).Decode(&metricResponse); err != nil {
		return model.SystemMetrics{}, fmt.Errorf("failed to decode Prometheus response")
	}
	if metricResponse.Status != "success" {
		return model.SystemMetrics{}, fmt.Errorf("frometheus query failed")
	}

	interfacesMetricsMap := make(map[string]*model.InterfaceMetrics)
	var cpuMetrics model.CpuMetrics
	var memMetrics model.MemoryMetrics
	storageMetricsMap := make(map[string]*model.StorageMetrics)
	var temperatureMetrics model.MetricsCouple

	for _, result := range metricResponse.Data.Result {
		device := result.Metric["device"]
		metricType := strings.Split(result.Metric["metric"], "_")

		if metricType[0] == "net" {
			if _, exists := interfacesMetricsMap[device]; !exists {
				interfacesMetricsMap[device] = &model.InterfaceMetrics{Interface: device}
			}

			if len(result.Value) == 2 {
				if bytes, err := json.Number(result.Value[1].(string)).Float64(); err == nil {
					humanReadableValue := ConvertMetricBytes(bytes, true)
					if metricType[1] == "receive" {
						interfacesMetricsMap[device].Received = humanReadableValue
					} else if metricType[1] == "transmit" {
						interfacesMetricsMap[device].Transmitted = humanReadableValue
					}
				}
			} else {
				if metricType[1] == "receive" {
					interfacesMetricsMap[device].Received.Value = "-"
					interfacesMetricsMap[device].Received.Unit = ""
				} else if metricType[1] == "transmit" {
					interfacesMetricsMap[device].Transmitted.Value = "-"
					interfacesMetricsMap[device].Transmitted.Unit = ""
				}
			}
			// TODO: if query_range endpoint is used, then return Values as [][]interface{}
		} else if metricType[0] == "cpu" {
			if usg, err := json.Number(result.Value[1].(string)).Float64(); err == nil {
				humanReadableValue := fmt.Sprintf("%.2f", usg)
				if metricType[1] == "usage" {
					cpuMetrics.Usage.Value = humanReadableValue
					cpuMetrics.Usage.Unit = "&percnt;"
					//cpuMetrics.Usage.Warn = true
				} else if metricType[1] == "load1" {
					cpuMetrics.Load1.Value = humanReadableValue
				} else if metricType[1] == "load5" {
					cpuMetrics.Load5.Value = humanReadableValue
				} else if metricType[1] == "load15" {
					cpuMetrics.Load15.Value = humanReadableValue
				}
			}
		} else if metricType[0] == "mem" {
			if metricType[1] == "usage" {
				if usg, err := json.Number(result.Value[1].(string)).Float64(); err == nil {
					memMetrics.Usage.Value = fmt.Sprintf("%.2f", usg)
					memMetrics.Usage.Unit = "&percnt;"
				}
			} else if bytes, err := json.Number(result.Value[1].(string)).Float64(); err == nil {
				humanReadableValue := ConvertMetricBytes(bytes, false)
				if metricType[1] == "total" {
					memMetrics.Total = humanReadableValue
				} else if metricType[1] == "free" {
					memMetrics.Free = humanReadableValue
				} else if metricType[1] == "cache" {
					memMetrics.Cache = humanReadableValue
				}
			}
		} else if metricType[0] == "disk" {
			if len(metricType) != 3 {
				fmt.Println("Invalid disk metric in response:", result.Metric["metric"])
				continue
			}
			storageName := metricType[1]
			if _, exists := storageMetricsMap[storageName]; !exists {
				storageMetricsMap[storageName] = &model.StorageMetrics{Name: strings.ToUpper(storageName)}
			}
			if metricType[2] == "usage" {
				if usg, err := json.Number(result.Value[1].(string)).Float64(); err == nil {
					storageMetricsMap[storageName].Usage.Value = fmt.Sprintf("%.2f", usg)
					storageMetricsMap[storageName].Usage.Unit = "&percnt;"
				}
			} else if bytes, err := json.Number(result.Value[1].(string)).Float64(); err == nil {
				humanReadableValue := ConvertMetricBytes(bytes, false)
				if metricType[2] == "free" {
					storageMetricsMap[storageName].Free = humanReadableValue
				} else if metricType[2] == "used" {
					storageMetricsMap[storageName].Used = humanReadableValue
				}
			}
		} else if metricType[0] == "temp" {
			tempLabel := metricType[2]
			if temp, err := json.Number(result.Value[1].(string)).Float64(); err == nil {
				tempMetricValue := model.MetricValue{
					Label: strings.ToUpper(tempLabel),
					Value: fmt.Sprintf("%.2f", temp),
					Unit:  "&deg;C",
				}
				if metricType[1] == "top" {
					temperatureMetrics.Top = tempMetricValue
				} else if metricType[1] == "bot" {
					temperatureMetrics.Bot = tempMetricValue
				}
				temperatureMetrics.Show = true
			}
		}
	}

	// for net interfaces, extract all device keys, sort them and create sorted array
	netdevs := make([]string, 0, len(interfacesMetricsMap))
	for netdev := range interfacesMetricsMap {
		netdevs = append(netdevs, netdev)
	}
	sort.Strings(netdevs)

	var interfaceMetrics []model.InterfaceMetrics
	for _, netdev := range netdevs {
		interfaceMetrics = append(interfaceMetrics, *interfacesMetricsMap[netdev])
	}

	// for disk metrics, pair 'em up
	stornames := make([]string, 0, len(storageMetricsMap))
	for storname := range storageMetricsMap {
		stornames = append(stornames, storname)
	}
	sort.Strings(stornames)

	var storageMetrics []model.StorageMetricsCouple
	storidx := 0
	for _, storname := range stornames {
		if storidx < len(storageMetrics) {
			storageMetrics[storidx].Bot = *storageMetricsMap[storname]
			storidx++
		} else {
			// no index, i make a new one
			storageMetrics = append(storageMetrics, model.StorageMetricsCouple{Top: *storageMetricsMap[storname]})
		}
	}

	return model.SystemMetrics{
		Interfaces:  interfaceMetrics,
		Cpu:         cpuMetrics,
		Memory:      memMetrics,
		Storage:     storageMetrics,
		Temperature: temperatureMetrics,
	}, nil
}
