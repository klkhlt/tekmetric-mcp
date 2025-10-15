---
sidebar_position: 5
---

# Usage Examples

Practical examples of using the Tekmetric MCP Server for **tactical, day-to-day queries**.

:::warning Not For Financial Reporting
This tool is designed for quick lookups and tactical queries, NOT for financial reporting, revenue calculations, or business intelligence. For comprehensive reports and analytics, use Tekmetric's built-in reporting tools.
:::

## Basic Queries

### Getting Specific Records

```
Get customer 12345
```

```
Show me vehicle ID 789
```

```
Get repair order 456
```

```
Show me details for employee 111
```

### Listing Recent Records

```
Show me today's appointments
```

```
Get repair orders from this week
```

```
List appointments for tomorrow
```

## Search Queries

### Customer Search

```
Find customers named John Smith
```

```
Search for customer with phone 555-1234
```

```
Find customer with email john@example.com
```

```
Look up customer Mike Johnson
```

### Vehicle Search

```
Find vehicle with VIN 1HGBH41JXMN109186
```

```
Search for 2020 Toyota Camry
```

```
Find vehicles with license plate ABC123
```

```
Look up Honda Accord
```

### Repair Order Search

```
Find repair order #12345
```

```
Show me repair orders with status "wip"
```

```
Find repair orders for customer 789
```

```
Get estimates from last week
```

## Daily Operations

### Appointments

```
Who's scheduled today?
```

```
What appointments does customer 456 have?
```

```
Show me next week's schedule
```

```
Find appointments for vehicle 789
```

### Work in Progress

```
What repair orders are in progress?
```

```
Show me all current estimates
```

```
Find jobs assigned to technician 123
```

```
Which repair orders are waiting for parts?
```

### Customer Service

```
What vehicles does customer 456 own?
```

```
Find repair history for vehicle VIN 1HGBH41JXMN109186
```

```
Show me customer John Smith's contact info
```

```
What's the status of repair order 789?
```

## Filtering Examples

### Date Ranges

```
Show me repair orders created last week
```

```
Find appointments between Jan 1 and Jan 7
```

```
Get estimates from yesterday
```

### Status Filtering

```
Show me all estimates (not just completed work)
```

```
Find repair orders with status "complete"
```

```
List jobs that are authorized
```

### Combined Filters

```
Find wip repair orders for customer 123
```

```
Show me completed repair orders from last week
```

```
Get appointments for customer 456 next week
```

## Tips for Better Results

### Be Specific

❌ **Too broad**: "Show me all customers"
✅ **Specific**: "Find customer named Mike Johnson"

### Use Names and Numbers

❌ **Vague**: "Who came in recently?"
✅ **Clear**: "Show me appointments from this week"

### One Thing at a Time

❌ **Too complex**: "Analyze revenue trends and customer patterns"
✅ **Tactical**: "Show me repair orders from last week"

### Avoid Financial Calculations

❌ **Don't ask**: "What's our total revenue this month?"
✅ **Instead**: Use Tekmetric's built-in reports

❌ **Don't ask**: "Calculate average repair order value"
✅ **Instead**: "Show me repair order #12345"

## What NOT to Ask

### ❌ Financial/Reporting Queries

These require comprehensive data and should use Tekmetric's reports:

- "Calculate total revenue for last month"
- "What's our average repair order value?"
- "Show me profit margins by service type"
- "Compare revenue year over year"
- "What's our average labor rate?"

### ❌ Large-Scale Analysis

Results are limited to 25 records, so avoid:

- "Analyze all repair orders from last year"
- "Show me customer lifetime values"
- "Which services generate the most revenue?"
- "Create a performance dashboard"
- "Generate a customer retention report"

### ❌ Complex Multi-Step Queries

These are too complex for tactical lookups:

- "Find all vehicles not serviced in 12 months and notify owners"
- "Analyze completion times and identify bottlenecks"
- "Compare shop performance across all locations"

## What TO Ask Instead

### ✅ Tactical Lookups

- "Find customer Mike Johnson's phone number"
- "Look up VIN 1HGBH41JXMN109186"
- "Show me repair order #12345"
- "What's on the schedule today?"
- "Find jobs for technician 456"

### ✅ Specific Searches

- "Get repair orders for customer 789"
- "Find vehicles owned by customer 123"
- "Show me estimates from last week"
- "What appointments are scheduled tomorrow?"

### ✅ Status Checks

- "What repair orders are in progress?"
- "Show me today's appointments"
- "Find jobs assigned to employee 111"
- "Get details for vehicle ID 456"

## Error Handling

### Common Errors and Solutions

**Error**: "Shop not found"
```
Make sure you're using the correct shop ID.
Try: "Show me all my shops" first
```

**Error**: "Customer not found"
```
Try searching instead:
"Find customers named [name]"
```

**Error**: "Too many results"
```
Add more specific filters:
- Use customer ID, vehicle ID, or date range
- Search for specific names or numbers
```

## Next Steps

- Review [Available Tools](../tools/index.md)
- See [Installation Guide](../installation/index.md)
- Check [Configuration Options](../configuration/index.md)

## Need Help?

- [GitHub Issues](https://github.com/beetlebugorg/tekmetric-mcp/issues)
- [Tekmetric API Documentation](https://api.tekmetric.com)
