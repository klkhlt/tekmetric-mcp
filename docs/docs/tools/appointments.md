---
sidebar_position: 6
---

# Appointment Tools

The `appointments` tool provides access to scheduled appointment information with flexible search and filtering options.

## appointments

Search and filter appointments, or get a specific appointment by ID. Supports filtering by date range, customer, vehicle, status, and more.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No* | Get specific appointment by ID | - |
| `search` | string | No | Search appointments by customer name or vehicle info | - |
| `shop` | number | No | Shop ID | Default shop |
| `customer_id` | number | No | Filter by customer ID | - |
| `vehicle_id` | number | No | Filter by vehicle ID | - |
| `start_date` | string | No | Filter appointments starting from this date (YYYY-MM-DD) | - |
| `end_date` | string | No | Filter appointments up to this date (YYYY-MM-DD) | - |
| `updated_start` | string | No | Filter by appointments updated after this date (YYYY-MM-DD) | - |
| `updated_end` | string | No | Filter by appointments updated before this date (YYYY-MM-DD) | - |
| `status` | string | No | Filter by appointment status | - |
| `sort` | string | No | Property to sort results by | - |
| `sort_direction` | string | No | Sort direction (ASC or DESC) | - |
| `limit` | number | No | Maximum results to return (max: 100) | 20 |
| `page` | number | No | Page number for pagination | 0 |

**Note:** If `id` is provided, it returns that specific appointment. Otherwise, it searches with the provided filters.

### How to Use

#### Get a Specific Appointment

```
Show me appointment 789
```

```
Get details for appointment ID 456
```

#### Search by Date Range

```
Show me appointments for today
```

```
Find appointments scheduled this week
```

```
What appointments do we have between January 1 and January 15?
```

#### Search by Customer or Vehicle

```
Find appointments for customer ID 456
```

```
Show me appointments for vehicle 789
```

### Response Format

Appointment information includes:
- **Appointment ID** - Unique identifier
- **Customer Name** - Who's coming in
- **Date & Time** - When they're scheduled
- **Vehicle** - What they're bringing in
- **Service Requested** - What work they need
- **Duration** - How long it's scheduled for
- **Status** - Confirmed, pending, completed, cancelled
- **Notes** - Any special instructions
- **Service Advisor** - Who's handling it
- **Shop ID** - Which location

### Pagination

When searching appointments, results are paginated:
- Use `limit` to control how many results per page (max 100)
- Use `page` to navigate through pages (0-indexed)
- Response includes `totalElements` and `totalPages`

## Common Use Cases

### Today's Schedule

```
Show me today's appointments
```

```
What appointments do we have this morning?
```

```
Who's scheduled for today at shop 123?
```

### Date Range Queries

```
Show me appointments for this week
```

```
Find appointments between Jan 1 and Jan 31
```

```
What's scheduled for next Monday?
```

### Customer-Specific Appointments

```
Find all appointments for customer 456
```

```
Show me Sarah Johnson's appointments
```

```
What time is customer 123 scheduled?
```

### Vehicle Appointments

```
Show me appointments for vehicle 789
```

```
Find all appointments for this VIN
```

## What Information You'll Get

When you look up an appointment, you'll see:

- **Customer Name** - Who's coming in
- **Date & Time** - When they're scheduled
- **Vehicle** - What they're bringing in
- **Service Requested** - What work they need
- **Duration** - How long it's scheduled for
- **Status** - Confirmed, pending, completed
- **Notes** - Any special instructions
- **Service Advisor** - Who's handling it

## Real-World Uses

### Morning Planning

```
What appointments do we have today?
```

```
Show me this morning's schedule
```

```
Who's our first appointment?
```

### Customer Calls

Customer: *"When is my appointment?"*

```
Find John Smith's appointment
```

```
Show me appointments for phone 555-1234
```

### Weekly Planning

```
How many appointments next week?
```

```
Show me Monday's schedule
```

```
What days are we fully booked?
```

### Capacity Planning

```
Do we have any openings tomorrow?
```

```
What time slots are available Friday?
```

```
How many appointments this week?
```

## Examples by Role

### Service Advisor

```
Show me my appointments for today
```

```
Who's my 10am appointment?
```

```
Find appointments that need confirmation
```

### Front Desk

```
List today's arrivals
```

```
Who's scheduled in the next hour?
```

```
Show me late arrivals
```

### Shop Manager

```
How many appointments do we have this week?
```

```
What's tomorrow's workload look like?
```

```
Show me no-shows from last week
```

## Search Tips

**By Date:**
- "today"
- "tomorrow"
- "next Monday"
- "January 15"

**By Customer:**
- "Find Sarah's appointment"
- "Show me appointments for customer 456"

**By Time:**
- "this morning"
- "this afternoon"
- "after 2pm"

## Connecting to Other Information

From an appointment, you can find:

```
What vehicle is appointment 789 for?
```

```
Show me customer details for this appointment
```

```
Has this customer been here before?
```

## Need More Help?

- Look up [customers](./customers.md)
- Check [vehicles](./vehicles.md)
- See [repair orders](./repair-orders.md)
- View [usage examples](../examples/index.md)
