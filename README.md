# Weekcase Metric Library

An open-source library of standardized key business metrics, with human-readable definitions and machine-translatable specifications for key data sources.

Weekcase Metric Library (WML) reduces ambiguity around common business metrics and their extraction from different data sources by providing a stable, versioned viewpoint.

## Philosophy

- **High signal-to-noise ratio**: Only key metrics businesses depend on
- **Human and machine readable**: Clear definitions paired with practical specifications
- **Source agnostic**: Multiple implementation options for each metric
- **Dependency aware**: Explicit relationships between derived metrics
- **Strictly versioned**: Immutable schemas ensure consistent interpretation across systems

## Structure

### Metric Index (`index.json`)

Each metric contains:

- **name**: Display name
- **id**: Unique identifier that matches the file tree
- **sources**: Supported systems for this metric's data, with an empty array marking a derived metric
- **dependencies**: For derived metrics, this array contains the metrics required for its calculation

Example:
```json
{
    "name": "Monthly Recurring Revenue",
    "id": "monthly_recurring_revenue",
    "sources": ["stripe"],
    "dependencies": []
}
```

### Metric Definitions (`metrics/{id}/definition.json`)

Each metric has a definition file conforming to [`schemas/definition.json`](schemas/definition.json):

- **version**: Semantic version of the definition
- **id**: Unique identifier
- **name**: Display name
- **abbreviation**: Standard abbreviation (or null)
- **description**: What the metric represents

Example:
```json
{
    "version": "1.0.0",
    "id": "monthly_recurring_revenue",
    "name": "Monthly Recurring Revenue",
    "abbreviation": "MRR",
    "description": "The normalized monthly value of all active recurring revenue streams."
}
```

### Source Specifications (`metrics/{id}/sources/{source}.json`)

For metrics with data sources, specifications conform to [`schemas/source.json`](schemas/source.json):

- **metric**: The metric identifier
- **source**: Data source system (e.g., stripe)
- **version**: Semantic version of this specification
- **url**: API endpoint
- **filters**: Query parameters
- **formula**: Mathematical definition of the calculation
- **normalization**: Rules for converting raw values to standard units

Example:
```json
{
    "metric": "monthly_recurring_revenue",
    "source": "stripe",
    "version": "1.0.0",
    "url": "https://api.stripe.com/v1/subscriptions",
    "filters": {
        "status": ["active", "past_due"]
    },
    "formula": "Sum of (subscription.items.data[].price.unit_amount x subscription.items.data[].quantity) normalized to monthly equivalent.",
    "normalization": {
        "period": "...",
        "currency": "..."
    }
}
```

## Schemas

All definitions and source specifications conform to strict JSON schemas in [`schemas/`](schemas/):

- [`definition.json`](schemas/definition.json): Schema for metric definitions
- [`source.json`](schemas/source.json): Schema for source specifications

Both schemas enforce immutability (`additionalProperties: false`) to ensure consistent interpretation across systems and versions.

## License

Apache 2.0 (see [LICENSE](LICENSE))
