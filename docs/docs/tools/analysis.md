# Analysis Tools

Analysis tools fetch and aggregate large datasets from the Tekmetric API, then return structured data with intelligent formatting guidance for Claude. These tools handle server-side pagination and data processing, while Claude applies its domain knowledge to categorize, format, and interpret the results.

## Available Analysis Tools

### Vehicle Service Analysis

Get a complete service history timeline for any vehicle, formatted as clean markdown tables with intelligent categorization.

**Tool Name:** `vehicle_service_analysis`

**What It Does:**
- Fetches all repair orders for a vehicle (server-side pagination)
- Extracts services, parts, costs, labor hours, and customer concerns
- Returns concise timeline data
- Claude categorizes services (oil changes, brakes, tires, etc.) using automotive knowledge
- Presents as scannable markdown tables with spending analysis

**Use Cases:**
- Review complete maintenance history before recommending services
- Identify recurring issues or chronic problems
- Assess maintenance adherence (oil changes, scheduled services)
- Understand spending patterns and vehicle health
- Prepare for customer conversations with full context

**Parameters:**

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `vehicle_id` | number | Yes | Vehicle ID to analyze |
| `shop_id` | number | No | Shop ID (uses default if not specified) |
| `start_date` | string | No | Filter history from date (YYYY-MM-DD) |
| `end_date` | string | No | Filter history to date (YYYY-MM-DD) |
| `max_pages` | number | No | Maximum pages to fetch (default 10, max 50) |

**Example Usage:**

```
Analyze service history for vehicle 12345
```

```
Show me the complete maintenance timeline for vehicle 67890
```

```
Get service history for vehicle 12345 from 2024-01-01
```

**Example Output:**

The tool returns a complete timeline that Claude formats like this:

---

## Service Timeline

| Date | Miles | Services & Parts | Hours | Cost | Notes |
|------|-------|------------------|-------|------|-------|
| 2024-01-15 | 45,200 | Oil Change, Air Filter • Motorcraft Oil, K&N Filter | 1.5 | $89.50 | Check engine light |
| 2024-04-22 | 48,100 | Brake Service • Front Pads, Rotors | 2.0 | $425.00 | |
| 2024-07-10 | 51,300 | Oil Change, Tire Rotation | 1.0 | $75.00 | |

## Analysis Summary

- **Service Frequency**: 3 visits over 6 months, average $196.50 per visit
- **Common Services**: Oil changes (2x), brake work (1x), tire maintenance (1x)
- **Maintenance Schedule**: Oil changes at ~3,100 mile intervals (good adherence)
- **Notable Patterns**: No recurring issues identified

---

**What Makes This Tool Smart:**

1. **Server-Side Pagination**: Automatically fetches all repair orders across multiple pages
2. **Concise Data**: Returns only essential information (no bloat)
3. **Claude's Intelligence**: Categorizes services using automotive knowledge (not brittle keyword matching)
4. **Scannable Format**: Markdown tables make it easy to read
5. **Complete Context**: Includes dates, mileage, parts, concerns, costs

**Performance:**

- Fetches up to 50 pages (5,000 repair orders) by default
- Typical execution: 1-3 seconds for most vehicles
- Metadata included: records fetched, pages traversed, execution time

## How Analysis Tools Work

Analysis tools follow a collaborative pattern:

### 1. Tool Does Data Work
- Fetches paginated API data (handles pagination automatically)
- Aggregates and processes records
- Extracts relevant information
- Returns structured data + formatting guidance

### 2. Claude Does Intelligent Work
- Categorizes items using domain knowledge
- Formats as clean markdown tables
- Identifies patterns and insights
- Presents information clearly

This approach avoids:
- ❌ Hard-coded categorization logic (brittle)
- ❌ Keyword matching that breaks with variations
- ❌ Massive context dumps (inefficient)
- ❌ Loss of domain expertise (Claude knows automotive stuff!)

## Comparison: Regular Tools vs Analysis Tools

### Regular Tools (e.g., `get_repair_orders`)
- Return raw paginated data
- Limited to 25 results per call
- Requires multiple calls for complete history
- No aggregation or processing
- Good for: Quick lookups, recent data

### Analysis Tools (e.g., `vehicle_service_analysis`)
- Handle pagination automatically (server-side)
- Return complete datasets (up to configured limits)
- Aggregated and processed
- Formatted for readability
- Good for: Complete timelines, historical analysis, patterns

## Configuration

Analysis tools respect these configuration settings:

```bash
# Maximum pages to fetch per analysis
TEKMETRIC_ANALYSIS_MAX_PAGES=50

# Maximum records to process
TEKMETRIC_ANALYSIS_MAX_RECORDS=5000

# Timeout for analysis operations (seconds)
TEKMETRIC_ANALYSIS_TIMEOUT_SECONDS=120
```

## Troubleshooting

**Slow responses:**
- Reduce `max_pages` parameter
- Use date filters to limit scope
- Check if vehicle has thousands of repair orders

**Missing data:**
- Verify vehicle ID is correct
- Check shop ID if using multi-shop setup
- Ensure date filters aren't too restrictive

**Incomplete results:**
- Increase `max_pages` if needed (up to 50)
- Check metadata to see how many pages were fetched
- Consider splitting by date range for very active vehicles

## Related Documentation

- [Repair Orders Tool](repair-orders.md) - For quick repair order lookups
- [Vehicles Tool](vehicles.md) - For finding vehicle IDs
- [Configuration](../configuration/environment-variables.md) - Analysis tool settings
