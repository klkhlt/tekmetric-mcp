---
sidebar_position: 4
---

# Repair Order Tools

The `repair_orders` tool provides comprehensive access to repair order data with powerful search and filtering options.

## repair_orders

Search and filter repair orders, or get a specific RO by ID. Supports filtering by date range, status, customer, vehicle, and more.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No | Get specific repair order by ID | - |
| `search` | string | No | Search by RO#, customer name, or vehicle info | - |
| `shop` | number | No | Shop ID | Default shop |
| `start_date` | string | No | Filter by created after date (YYYY-MM-DD) | - |
| `end_date` | string | No | Filter by created before date (YYYY-MM-DD) | - |
| `status` | string | No | Filter by status (see below) | - |
| `customer_id` | number | No | Filter by customer ID | - |
| `vehicle_id` | number | No | Filter by vehicle ID | - |
| `limit` | number | No | Maximum results | 20 |

### Status Values

| Status | Description |
|--------|-------------|
| `estimate` | Quote given, not approved yet |
| `wip` | Work in progress |
| `complete` | Work done, ready for pickup |
| `saved` | Saved for later |
| `posted` | Closed and billed |
| `ar` | Accounts receivable |
| `deleted` | Deleted repair orders |

**Note:** You can provide multiple statuses separated by commas (e.g., `"estimate,wip"`).

### How to Use

#### Get a Specific Repair Order

```
Show me repair order 12345
```

```
Get details for RO 789
```

#### Search Repair Orders

```
Find repair orders for customer Sarah Johnson
```

```
Search for RO number 456
```

```
Find repair orders for 2020 Honda Accord
```

#### Filter by Status

```
What repair orders are in progress?
```

```
Show me all estimates
```

```
Find completed repair orders
```

#### Filter by Date

```
Show me repair orders from this month
```

```
Find ROs created in the last 7 days
```

```
What repair orders were completed between Jan 1 and Jan 31?
```

#### Advanced Filtering

```
Show me repair orders for customer 456 with status complete
```

```
Find all repair orders for vehicle 789 from last month
```

### What You'll Get

Repair order information includes:
- RO number and ID
- Customer information
- Vehicle details
- Status
- Services/jobs
- Parts used
- Labor charges
- Total cost
- Dates (created, updated, completed)
- Assigned technician

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

**Common statuses:**
- **Estimate** - Quote given, not approved yet
- **In Progress (WIP)** - Currently being worked on
- **Complete** - Work done, ready for pickup
- **Posted** - Closed and billed

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

### Search by Amount

```
Find repair orders over $1000
```

```
Show me estimates under $500
```

## What Information You'll Get

When you look up a repair order, you'll see:

- **RO Number** - The repair order ID
- **Customer** - Who owns the vehicle
- **Vehicle** - Make, model, year
- **Status** - Where it is in the process
- **Services** - What work is being done
- **Parts** - Parts needed or used
- **Labor** - Labor charges
- **Total** - Total cost
- **Dates** - Created, updated, completed dates
- **Technician** - Who's working on it

## Real-World Uses

### Morning Meeting

```
What repair orders are in progress today?
```

```
Show me yesterday's completed ROs
```

```
Which estimates need follow-up?
```

### Customer Calls

Customer: *"What's the status of my repair?"*

```
Find repair orders for John Smith
```

```
Show me open ROs for phone number 555-1234
```

### End of Day

```
How many repair orders did we complete today?
```

```
What's still in progress?
```

```
Show me tomorrow's scheduled work
```

### End of Month Reporting

```
Show me all completed ROs from last month
```

```
What was our total revenue in June?
```

```
How many estimates did we create this month?
```

## Search Tips

**By Customer:**
```
Find ROs for customer ID 456
Show me all repair orders for Sarah Johnson
```

**By Vehicle:**
```
Find repair orders for VIN 1HGBH41JXMN109186
Show me ROs for vehicle 789
```

**By Status:**
```
List all estimates
Show me work in progress
Find completed jobs
```

**By Date Range:**
```
ROs from January 1 to January 31
Show me this week's repair orders
Last month's completed work
```

## Understanding the Results

### RO Status Flow

1. **Estimate** → Customer approves →
2. **In Progress (WIP)** → Work completed →
3. **Complete** → Customer pays →
4. **Posted** → Closed out

### Common Scenarios

**Checking on a specific customer:**
```
Find Sarah Johnson's repair orders
```
Shows all ROs for that customer, past and present.

**Daily work queue:**
```
What's in progress at shop 123?
```
Shows everything currently being worked on.

**Follow up on estimates:**
```
Show me estimates older than 2 weeks
```
Find quotes that haven't been approved yet.

## Connecting to Other Information

Once you find an RO:

```
Who's the customer for RO 12345?
```

```
What vehicle is on repair order 789?
```

```
Show me the technician working on RO 456
```

## Need More Help?

- Learn about [customers](./customers.md)
- Look up [vehicles](./vehicles.md)
- See [jobs on repair orders](./jobs.md)
- Check [usage examples](../examples/index.md)
