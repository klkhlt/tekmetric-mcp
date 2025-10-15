---
sidebar_position: 7
---

# Employee Tools

The `employees` tool provides access to employee and technician information with flexible search and filtering options.

## employees

Search and filter employees/technicians, or get a specific employee by ID. Supports filtering by active status, role, and more.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No* | Get specific employee by ID | - |
| `search` | string | No | Search employees by name or email | - |
| `shop` | number | No | Shop ID | Default shop |
| `active` | boolean | No | Filter by active status (true = active only, false = inactive only) | - |
| `role` | string | No | Filter by employee role (e.g., technician, service advisor, manager) | - |
| `sort` | string | No | Property to sort results by | - |
| `sort_direction` | string | No | Sort direction (ASC or DESC) | - |
| `limit` | number | No | Maximum results to return (max: 100) | 20 |
| `page` | number | No | Page number for pagination | 0 |

**Note:** If `id` is provided, it returns that specific employee. Otherwise, it searches with the provided filters.

### How to Use

#### Get a Specific Employee

```
Show me employee 456
```

```
Get details for employee ID 789
```

```
Look up technician 123
```

#### Search by Name or Email

```
Find employees named Mike
```

```
Search for employees with email containing @shop.com
```

```
Look up employee Sarah Johnson
```

#### Filter by Status

```
Show me all active employees
```

```
List active technicians at shop 123
```

```
Find inactive employees
```

#### Filter by Role

```
Show me all technicians
```

```
List service advisors
```

```
Find employees with role manager
```

### Response Format

Employee information includes:
- **Employee ID** - Unique identifier
- **Name** - Full name (first and last)
- **Role** - Position (technician, service advisor, manager, etc.)
- **Email** - Contact email address
- **Phone** - Phone number
- **Status** - Active or inactive
- **Shop ID** - Which location they work at
- **Hire Date** - When they started
- **Skills/Certifications** - Any special qualifications

### Pagination

When searching employees, results are paginated:
- Use `limit` to control how many results per page (max 100)
- Use `page` to navigate through pages (0-indexed)
- Response includes `totalElements` and `totalPages`

## Common Use Cases

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

## What Information You'll Get

When you look up an employee, you'll see:

- **Name** - Employee's full name
- **Employee ID** - Their unique ID number
- **Role** - Technician, service advisor, manager, etc.
- **Email** - Contact email
- **Phone** - Phone number
- **Status** - Active or inactive
- **Shop** - Which location they work at

## Real-World Uses

### Find Contact Info

```
What's Mike's phone number?
```

```
Show me Sarah's email
```

### Check Staffing

```
Who's working today?
```

```
List all active technicians
```

```
How many employees do we have?
```

### Assign Work

```
Show me available technicians
```

```
Who can work on this repair order?
```

## Examples

### Manager Scenarios

```
List all employees at shop 123
```

```
Show me technician contact information
```

```
Who's scheduled to work tomorrow?
```

### Service Advisor Scenarios

```
Find the technician working on RO 789
```

```
Show me Mike's current jobs
```

```
Who should I assign this brake job to?
```

## Tips

- Employee list shows who's at your shop
- Use for assigning work to technicians
- Look up contact info when needed
- Check who's available for scheduling

## Connecting to Other Information

From employee information:

```
What jobs is technician 456 working on?
```

```
Show me Mike's completed work this week
```

```
Which repair orders is Sarah assigned to?
```

## Need More Help?

- Check [jobs and work assignments](./jobs.md)
- See [repair orders](./repair-orders.md)
- View [usage examples](../examples/index.md)
