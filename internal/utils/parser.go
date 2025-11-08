package utils

import (
	"os"
	"path/filepath"
	"sort"

	"cuelang.org/go/cue/cuecontext"

	"metrics/internal/models"
)

func LoadMetrics() []models.MetricListItem {
	var metrics []models.MetricListItem
	ctx := cuecontext.New()

	dirs, err := os.ReadDir("./metrics")
	if err != nil {
		return metrics
	}

	for _, dir := range dirs {
		if !dir.IsDir() {
			continue
		}

		defPath := filepath.Join("./metrics", dir.Name(), "definition.cue")
		data, err := os.ReadFile(defPath)
		if err != nil {
			continue
		}

		// Parse CUE
		val := ctx.CompileBytes(data)
		if val.Err() != nil {
			continue
		}

		// Decode into struct
		var def models.MetricDefinition
		if err := val.Decode(&def); err != nil {
			continue
		}

		metrics = append(metrics, models.MetricListItem{
			ID:   dir.Name(),
			Name: def.Name,
		})
	}

	sort.Slice(metrics, func(i, j int) bool {
		return metrics[i].Name < metrics[j].Name
	})

	return metrics
}

func GetMetric(id string) *models.MetricDefinition {
	ctx := cuecontext.New()

	defPath := filepath.Join("./metrics", id, "definition.cue")
	data, err := os.ReadFile(defPath)
	if err != nil {
		return nil
	}

	val := ctx.CompileBytes(data)
	if val.Err() != nil {
		return nil
	}

	var def models.MetricDefinition
	if err := val.Decode(&def); err != nil {
		return nil
	}

	def.ID = id
	return &def
}

func GetMetricName(id string) string {
	metric := GetMetric(id)
	if metric == nil {
		return id // Fallback to ID if not found
	}
	return metric.Name
}

func GetSource(metricID, sourceName string) *models.SourceSpec {
	ctx := cuecontext.New()

	sourcePath := filepath.Join("./metrics", metricID, "sources", sourceName+".cue")
	data, err := os.ReadFile(sourcePath)
	if err != nil {
		return nil
	}

	val := ctx.CompileBytes(data)
	if val.Err() != nil {
		return nil
	}

	var source models.SourceSpec
	if err := val.Decode(&source); err != nil {
		return nil
	}

	return &source
}

func GetSourceName(metricID, sourceID string) string {
	source := GetSource(metricID, sourceID)
	if source == nil {
		return sourceID // Fallback to ID if not found
	}
	return source.Source
}
