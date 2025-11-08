# Weekcase Metric Library

An open-source library of standardized key business metrics. It includes clear, versioned definitions and extraction specifications for key business systems.

Weekcase Metric Library (WML) exists to reduce ambiguity around common business metrics by providing a stable, trusted viewpoint for consistent reporting across tools and teams.

WML is published at [metrics.weekcase.net](https://metrics.weekcase.net). This repository contains both the metric definitions and related source specifications, and the web server that serves the published library.

## Philosophy

- **High signal-to-noise ratio**: Only key metrics businesses depend on
- **Human and machine readable**: Clear definitions paired with practical specifications
- **Source agnostic**: Multiple implementation options for each metric
- **Dependency aware**: Explicit relationships between derived metrics
- **Strictly versioned**: Serious focus towards reliability and stability

## Structure

```
metrics/
  {metric_id}/
    definition.cue          # Metric definition
    sources/
      {source_id}.cue       # Source-specific implementation specification
```

### Example: Monthly Recurring Revenue

```
metrics/
  monthly_recurring_revenue/
    definition.cue
    sources/
      stripe.cue
```

## Format

Metrics are defined using [CUE](https://cuelang.org/) for type-safe, human-readable configuration with built-in validation.

### Metric Definition Example

```cue
name: "Monthly Recurring Revenue"
definition: "The normalized monthly value of all active recurring revenue streams."
abbreviation: "MRR"
version: "1.0.0"
dependencies: []
sources: ["stripe"]
```

### Source Specification Example

```cue
metric: "Monthly Recurring Revenue"
source: "Stripe"
version: "1.0.0"
url: "https://api.stripe.com/v1/subscriptions"
filters: {
    status: ["active", "past_due"]
}
formula: "Sum of (subscription.items.data[].price.unit_amount x subscription.items.data[].quantity) normalized to monthly equivalent."
normalization: {
    period: "price.recurring.interval normalized to monthly equivalent..."
    currency: "price.unit_amount converted to base currency..."
}
```

## Tech Stack

- **Format**: [CUE](https://cuelang.org/)
- **Web**: Go, [Gin](https://github.com/gin-gonic/gin), HTML, CSS

## Running Locally

Start a local web server to open the library in your browser:

```bash
go mod tidy
go run cmd/main.go
```

Visit `http://localhost:8080` once the server is running.

## Contributing

Contributions are welcomed. Please ensure:
1. Metric definitions follow the established schema
2. Source specifications are complete and accurate
3. Version numbers follow [Semantic Versioning](https://semver.org)

## License

Apache 2.0 (see [LICENSE](LICENSE))
