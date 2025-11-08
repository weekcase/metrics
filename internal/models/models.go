package models

type MetricDefinition struct {
	ID           string   `json:"id"`
	Name         string   `json:"name"`
	Abbreviation string   `json:"abbreviation"`
	Version      string   `json:"version"`
	Definition   string   `json:"definition"`
	Dependencies []string `json:"dependencies"`
	Sources      []string `json:"sources"`
}

type MetricListItem struct {
	ID           string
	Name         string
	Abbreviation string
}

type SourceSpec struct {
	Metric        string                 `json:"metric"`
	Source        string                 `json:"source"`
	Version       string                 `json:"version"`
	URL           string                 `json:"url"`
	Filters       map[string]interface{} `json:"filters"`
	Formula       string                 `json:"formula"`
	Normalization Normalization          `json:"normalization"`
}

type Normalization struct {
	Period   string `json:"period"`
	Currency string `json:"currency"`
}
