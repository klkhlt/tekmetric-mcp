---
sidebar_position: 3
---

# Employee Tools

Access and search employee and technician information.

## What You Can Ask

### List All Employees

```
Show me all employees at shop 123
```

```
List all technicians
```

```
Who works here?
```

### Find Specific Employee

```
Show me employee 456
```

```
Find employee Mike Johnson
```

```
Look up technician Sarah
```

### Search by Name or Email

```
Find employees named Mike
```

```
Search for employees with email @shop.com
```

```
Look up Sarah Johnson
```

### Filter by Status

```
Show me active employees only
```

```
List all active technicians
```

```
Find inactive employees
```

### Filter by Role

```
List all technicians
```

```
Show me service advisors
```

```
Find employees with manager role
```

```
Who are our shop foremen?
```

### Examples by Role

**Shop Manager:**
```
Show me all active employees
```

```
How many technicians do we have?
```

**Service Advisor:**
```
Find the technician assigned to this repair order
```

```
Who's available to work on brakes?
```

**HR/Admin:**
```
List all employees with contact info
```

```
Show me inactive employees
```

## What You'll Get

Each employee includes:
- Employee ID
- Full name
- Role (technician, service advisor, manager, etc.)
- Email address
- Phone number
- Active/inactive status
- Shop location
- Hire date
- Skills and certifications

## Technical Reference

### Tool: `employees`

Search and filter employees, or get a specific employee by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No | Get specific employee by ID | - |
| `search` | string | No | Search by name or email | - |
| `shop` | number | No | Shop ID | Default shop |
| `active` | boolean | No | Filter by active status | - |
| `role` | string | No | Filter by role | - |
| `sort` | string | No | Sort property | - |
| `sort_direction` | string | No | ASC or DESC | - |
| `limit` | number | No | Max results (max: 100) | 20 |
| `page` | number | No | Page number | 0 |

### Response Format

Results are paginated with `totalElements` and `totalPages` fields.

## Related Tools

- [Jobs](./jobs.md) - Work assigned to technicians
- [Repair Orders](./repair-orders.md) - Service history
- [Shops](./shops.md) - Shop information
