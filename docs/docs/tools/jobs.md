---
sidebar_position: 5
---

# Job Tools

The `jobs` tool provides access to individual job (service) details on repair orders.

## jobs

Get detailed information for a specific job by ID.

### Parameters

| Parameter | Type | Required | Description |
|-----------|------|----------|-------------|
| `id` | number | **Yes** | Job ID |

### What Are Jobs?

A **job** is a specific service on a repair order. For example:
- Oil change
- Brake pad replacement
- Engine diagnostic
- Tire rotation

One repair order can have multiple jobs.

### How to Use

```
Show me job 456
```

```
Get details for job ID 789
```

**Note:** Jobs are typically discovered through repair orders. Use the `repair_orders` tool to find jobs associated with a specific RO.

### What You'll Get

Job information includes:
- Job ID
- Service name/description
- Repair order ID
- Vehicle information
- Assigned technician
- Status
- Labor time and cost
- Parts used
- Technician notes

## What You Can Ask

### Find Specific Jobs

```
Show me job 456
```

```
Find jobs on repair order 12345
```

### Browse Jobs

```
List jobs from shop 123
```

```
Show me today's jobs
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
