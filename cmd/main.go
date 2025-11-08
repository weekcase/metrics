package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"metrics/internal/utils"
)

func main() {
	r := gin.Default()
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("pages/*.html")

	// Health check endpoint
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	// Index page
	r.GET("/", func(c *gin.Context) {
		metrics := utils.LoadMetrics()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"Metrics": metrics,
		})
	})

	// Metric detail page
	r.GET("/:id", func(c *gin.Context) {
		metricID := c.Param("id")
		metric := utils.GetMetric(metricID)

		if metric == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		// Create a map of dependency names
		depNames := make([]string, len(metric.Dependencies))
		for i, depID := range metric.Dependencies {
			depNames[i] = utils.GetMetricName(depID)
		}

		// Create a map of source names
		sourceNames := make([]string, len(metric.Sources))
		for i, sourceID := range metric.Sources {
			sourceNames[i] = utils.GetSourceName(metricID, sourceID)
		}

		c.HTML(http.StatusOK, "metric.html", gin.H{
			"Metric":          metric,
			"DependencyNames": depNames,
			"SourceNames":     sourceNames,
		})
	})

	// Source detail page
	r.GET("/:id/:source", func(c *gin.Context) {
		metricID := c.Param("id")
		sourceID := c.Param("source")

		source := utils.GetSource(metricID, sourceID)

		if source == nil {
			c.Redirect(http.StatusFound, "/")
			return
		}

		c.HTML(http.StatusOK, "source.html", gin.H{
			"Source": source,
		})
	})

	r.Run(":8080")
}
