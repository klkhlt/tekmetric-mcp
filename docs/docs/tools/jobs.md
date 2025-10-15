---
sidebar_position: 5
---

# Jobs

Access and search individual service jobs on repair orders.

## What Are Jobs?

A **job** is a specific service task on a repair order:
- Oil change
- Brake pad replacement
- Engine diagnostic
- Tire rotation

One repair order can have multiple jobs.

## What You Can Ask

### Find Jobs by Repair Order

```
Show me all jobs on repair order 12345
```

```
Find jobs for RO 789
```

```
What services are on this repair order?
```

### Track Technician Workload

```
What jobs is Mike working on?
```

```
Show me jobs assigned to technician 456
```

```
Find all jobs for employee ID 123
```

```
What's Sarah working on right now?
```

### Find Specific Jobs

```
Show me job 456
```

```
Get details for job ID 789
```

### Browse and Filter Jobs

```
List all jobs from shop 123
```

```
Show me today's jobs
```

```
Find jobs in progress
```

```
List completed jobs from this week
```

### Search by Service Type

```
Find all oil change jobs
```

```
Show me brake jobs
```

```
List diagnostic jobs
```

### Examples by Role

**Shop Foreman:**
```
What jobs are in progress right now?
```

```
Show me jobs for technician Mike
```

**Service Advisor:**
```
Find jobs on repair order 12345
```

```
What's the status of job 789?
```

**Manager:**
```
Show me all jobs completed today
```

```
How many jobs are assigned to each technician?
```

## What You'll Get

Each job includes:
- Job ID and service name
- Repair order it belongs to
- Vehicle information
- Assigned technician
- Status (pending, in progress, completed)
- Labor time (estimated and actual)
- Labor cost
- Parts used
- Technician notes
- Created and updated dates

## Technical Reference

### Tool: `jobs`

Search and filter jobs, or get a specific job by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No | Get specific job by ID | - |
| `search` | string | No | Search by name or description | - |
| `shop` | number | No | Shop ID | Default shop |
| `repair_order_id` | number | No | Filter by repair order ID | - |
| `employee_id` | number | No | Filter by employee/technician ID | - |
| `status` | string | No | Filter by status | - |
| `sort` | string | No | Sort property | - |
| `sort_direction` | string | No | ASC or DESC | - |
| `limit` | number | No | Max results (max: 100) | 20 |
| `page` | number | No | Page number | 0 |

### Response Format

Results are paginated with `totalElements` and `totalPages` fields.

## Related Tools

- [Repair Orders](./repair-orders.md) - Parent repair orders
- [Employees](./employees.md) - Technician information
- [Vehicles](./vehicles.md) - Vehicle details
