---
sidebar_position: 5
---

# Job Tools

The `jobs` tool provides access to individual job (service) details on repair orders with flexible search and filtering options.

## jobs

Search and filter jobs, or get a specific job by ID. Supports filtering by repair order, employee, status, and more.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No* | Get specific job by ID | - |
| `search` | string | No | Search jobs by name or description | - |
| `shop` | number | No | Shop ID | Default shop |
| `repair_order_id` | number | No | Filter by repair order ID | - |
| `employee_id` | number | No | Filter by assigned employee/technician ID | - |
| `status` | string | No | Filter by job status | - |
| `sort` | string | No | Property to sort results by | - |
| `sort_direction` | string | No | Sort direction (ASC or DESC) | - |
| `limit` | number | No | Maximum results to return (max: 100) | 20 |
| `page` | number | No | Page number for pagination | 0 |

**Note:** If `id` is provided, it returns that specific job. Otherwise, it searches with the provided filters.

### What Are Jobs?

A **job** is a specific service on a repair order. For example:
- Oil change
- Brake pad replacement
- Engine diagnostic
- Tire rotation

One repair order can have multiple jobs.

### How to Use

#### Get a Specific Job

```
Show me job 456
```

```
Get details for job ID 789
```

#### Search by Repair Order

```
Find all jobs on repair order 12345
```

```
Show me jobs for RO 789
```

#### Search by Employee/Technician

```
Find jobs assigned to employee 456
```

```
Show me Mike's current jobs
```

```
What jobs is technician 123 working on?
```

#### Search by Status or Type

```
Find all oil change jobs
```

```
Show me jobs in progress
```

```
List completed brake jobs
```

### Response Format

Job information includes:
- **Job ID** - Unique identifier
- **Service Name/Description** - What service is being performed
- **Repair Order ID** - Which RO it belongs to
- **Vehicle** - What car it's for
- **Assigned Technician** - Who's working on it
- **Status** - Started, in progress, completed, etc.
- **Labor Time** - Estimated and actual hours
- **Labor Cost** - Cost for the labor
- **Parts Used** - Parts included in this job
- **Technician Notes** - Any notes from the tech
- **Created/Updated Dates** - When the job was created and last updated

### Pagination

When searching jobs, results are paginated:
- Use `limit` to control how many results per page (max 100)
- Use `page` to navigate through pages (0-indexed)
- Response includes `totalElements` and `totalPages`

## Common Use Cases

### Find Jobs by Repair Order

```
Show me all jobs on repair order 12345
```

```
Find jobs for RO 789
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

## What Information You'll Get

When you look up a job, you'll see:

- **Job ID** - The unique job number
- **Service Name** - What service is being performed
- **Repair Order** - Which RO it belongs to
- **Vehicle** - What car it's for
- **Technician** - Who's assigned to it
- **Status** - Started, in progress, completed
- **Labor** - Labor time and cost
- **Parts** - Parts used for this job
- **Notes** - Any technician notes

## Real-World Uses

### Check Technician Workload

```
What jobs is Mike working on?
```

```
Show me jobs assigned to technician 789
```

### Track Service Times

```
How long did job 456 take?
```

```
Show me completed jobs from today
```

### Find Specific Services

```
Show me all oil change jobs this month
```

```
Find brake jobs in progress
```

## Examples

### During Service

Shop foreman: *"What's Mike working on?"*

```
Show me jobs for technician Mike
```

```
What jobs are in progress right now?
```

### Quality Control

```
Show me jobs completed by Sarah this week
```

```
Find jobs that took over 3 hours
```

### Customer Questions

Customer: *"Is my oil change done yet?"*

```
Find jobs on repair order 12345
```

Shows all services on that RO and their status.

## Tips

- Jobs are tied to repair orders
- One RO can have many jobs
- Jobs show actual work being done
- Use for tracking technician productivity

## Need More Help?

- See [repair orders](./repair-orders.md)
- Check [employee information](./employees.md)
- View [usage examples](../examples/index.md)
