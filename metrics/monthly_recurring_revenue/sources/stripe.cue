metric: "Monthly Recurring Revenue"
source: "Stripe"
version: "1.0.0"
url: "https://api.stripe.com/v1/subscriptions"
filters: {
    status: ["active", "past_due"]
}
formula: "Sum of (subscription.items.data[].price.unit_amount x subscription.items.data[].quantity) normalized to monthly equivalent."
normalization: {
    period: "price.recurring.interval normalized to monthly equivalent using 30.4375 days per month, divided by price.recurring.interval_count."
    currency: "price.unit_amount converted to base currency using subscription.items.data[].price.currency and adjusted to major units [Stripe reports values in the smallest currency unit (e.g., cents for USD, yen for JPY)]."
}