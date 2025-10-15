---
sidebar_position: 6
---

# Repair Orders

Access and search repair orders with powerful filtering by status, date, customer, and vehicle.

## What You Can Ask

### Find Specific Repair Orders

```
Show me repair order 12345
```

```
Find RO number 789
```

```
Look up repair order for customer Sarah Johnson
```

### Search by Status

```
What repair orders are in progress?
```

```
Show me all estimates
```

```
Find completed repair orders from last week
```

```
Which ROs are waiting for parts?
```

### Search by Date

```
Show me repair orders from this month
```

```
Find ROs created in the last 7 days
```

```
What repair orders were completed yesterday?
```

```
Show me ROs between Jan 1 and Jan 31
```

### Search by Customer or Vehicle

```
Find repair orders for customer 456
```

```
Show me all repair orders for Sarah Johnson
```

```
Find repair orders for VIN 1HGBH41JXMN109186
```

```
Show me ROs for vehicle 789
```

### Examples by Role

**Morning Meeting:**
```
What repair orders are in progress today?
```

```
Show me yesterday's completed ROs
```

**Customer Calls:**
```
Find repair orders for John Smith
```

```
Show me open ROs for phone number 555-1234
```

**End of Day:**
```
How many repair orders did we complete today?
```

```
What's still in progress?
```

**Monthly Reports:**
```
Show me all completed ROs from last month
```

```
How many estimates did we create this month?
```

## Status Values

| Status | What It Means |
|--------|---------------|
| **Estimate** | Quote given, awaiting approval |
| **WIP** | Work in progress |
| **Complete** | Work done, ready for pickup |
| **Posted** | Closed and billed |
| **Saved** | Saved for later |
| **AR** | Accounts receivable |

**Status Flow:** Estimate → WIP → Complete → Posted

## What You'll Get

Each repair order includes:
- RO number and ID
- Customer information
- Vehicle details
- Current status
- Services/jobs list
- Parts used
- Labor charges
- Total cost
- Created, updated, and completed dates
- Assigned technician

## Technical Reference

### Tool: `repair_orders`

Search and filter repair orders, or get a specific RO by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No | Get specific repair order by ID | - |
| `search` | string | No | Search by RO#, customer, or vehicle | - |
| `shop` | number | No | Shop ID | Default shop |
| `start_date` | string | No | Created after date (YYYY-MM-DD) | - |
| `end_date` | string | No | Created before date (YYYY-MM-DD) | - |
| `status` | string | No | Filter by status (comma-separated) | - |
| `customer_id` | number | No | Filter by customer ID | - |
| `vehicle_id` | number | No | Filter by vehicle ID | - |
| `limit` | number | No | Max results (max 25) | 10 |

### Result Limits

⚠️ **Important:** Results are limited to 25 repair orders maximum to prevent overwhelming Claude's context window.

If you query returns more than 25 results, you'll see a warning message like:
```
WARNING: ⚠️ SHOWING ONLY 10 OF 500 TOTAL REPAIR ORDERS ⚠️
```

**To see more results**, add filters to narrow your search:
- Use date ranges (`start_date`, `end_date`)
- Filter by status (`estimate`, `wip`, `complete`)
- Search for specific customers or vehicles

## Related Tools

- [Customers](./customers.md) - Customer information
- [Vehicles](./vehicles.md) - Vehicle details
- [Jobs](./jobs.md) - Individual services on ROs
- [Employees](./employees.md) - Technician assignments
