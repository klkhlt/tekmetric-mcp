---
sidebar_position: 1
---

# Appointments

Access and search scheduled appointments for your shop.

## What You Can Ask

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

```
List appointments for tomorrow
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

### By Status

```
Show me confirmed appointments
```

```
Find pending appointments
```

```
List cancelled appointments
```

### Examples by Role

**Service Advisor:**
```
Show me my appointments for today
```

```
Who's my 10am appointment?
```

**Front Desk:**
```
Who's scheduled in the next hour?
```

```
List today's arrivals
```

**Shop Manager:**
```
How many appointments do we have this week?
```

```
What's tomorrow's workload look like?
```

## What You'll Get

Each appointment includes:
- Customer name and contact info
- Date and time scheduled
- Vehicle information
- Service requested
- Duration
- Status (confirmed, pending, completed, cancelled)
- Service advisor assigned
- Special notes or instructions

## Technical Reference

### Tool: `appointments`

Search and filter appointments, or get a specific appointment by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No | Get specific appointment by ID | - |
| `search` | string | No | Search by customer name or vehicle info | - |
| `shop` | number | No | Shop ID | Default shop |
| `customer_id` | number | No | Filter by customer ID | - |
| `vehicle_id` | number | No | Filter by vehicle ID | - |
| `start_date` | string | No | Start date (YYYY-MM-DD) | - |
| `end_date` | string | No | End date (YYYY-MM-DD) | - |
| `updated_start` | string | No | Updated after date (YYYY-MM-DD) | - |
| `updated_end` | string | No | Updated before date (YYYY-MM-DD) | - |
| `status` | string | No | Filter by status | - |
| `sort` | string | No | Sort property | - |
| `sort_direction` | string | No | ASC or DESC | - |
| `limit` | number | No | Max results (max: 100) | 20 |
| `page` | number | No | Page number | 0 |

### Response Format

Results are paginated with `totalElements` and `totalPages` fields.

## Related Tools

- [Customers](./customers.md) - Customer information
- [Vehicles](./vehicles.md) - Vehicle details
- [Repair Orders](./repair-orders.md) - Service history
