---
sidebar_position: 5
---

# Usage Examples

This page provides practical examples of using the Tekmetric MCP Server with Claude Desktop.

## Basic Queries

### Listing Resources

```
Show me all my shops
```

```
Get customers from shop 123
```

```
List all vehicles for shop 2
```

```
Show me repair orders from the last week
```

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

## Search Queries

### Customer Search

```
Find customers named John Smith
```

```
Search for customers with email containing "@gmail.com"
```

```
Find customer with phone number 555-1234
```

```
Show me all customers in Springfield
```

### Vehicle Search

```
Find vehicles with VIN 1HGBH41JXMN109186
```

```
Search for 2020 Toyota Camry
```

```
Find all Honda vehicles
```

```
Show me vehicles from 2018 or newer
```

### Repair Order Search

```
Find repair orders with status "complete"
```

```
Search for RO number 12345
```

```
Show me repair orders created this month
```

```
Find repair orders for customer 789
```

## Data Analysis

### Customer Analytics

```
Analyze customer purchase patterns for shop 123
```

```
Show me customers who haven't visited in 6 months
```

```
What's the average customer lifetime value?
```

```
Which customers have the most repair orders?
```

### Repair Order Analytics

```
What are the most common repair types this quarter?
```

```
Show me average repair order values by month
```

```
Analyze repair order completion times
```

```
Find repair orders over $1000
```

### Shop Performance

```
How many repair orders were completed this week?
```

```
What's the shop's average labor rate?
```

```
Show me appointment volume by day of week
```

```
Compare repair volumes across all shops
```

## Business Insights

### Revenue Analysis

```
Calculate total revenue for last month
```

```
Show me top 10 customers by spend
```

```
What's the average invoice amount?
```

```
Compare revenue year over year
```

### Customer Service

```
Find customers with pending estimates
```

```
Show me vehicles due for service
```

```
Which repair orders are overdue?
```

```
Find customers with multiple vehicles
```

### Inventory Management

```
Show me inventory parts below reorder point
```

```
List all parts with quantity > 0
```

```
Find parts used in recent repair orders
```

## Complex Queries

### Multi-Step Analysis

```
Find all repair orders for customer John Smith's vehicles,
calculate the total spend, and identify the most common repair types
```

```
Get all vehicles for shop 123, find which haven't been serviced
in 12 months, and list their owners
```

```
Analyze all completed repair orders from last quarter,
calculate average completion time, and identify bottlenecks
```

### Reporting

```
Create a summary report of all repair activity last month
including total orders, revenue, and top customers
```

```
Generate a customer retention report showing visit frequency
and spending patterns
```

```
Prepare a shop performance dashboard with key metrics
```

## Filtering and Sorting

### Date Ranges

```
Show me repair orders created between Jan 1 and Jan 31
```

```
Find customers added in the last 90 days
```

```
Get appointments scheduled for next week
```

### Status Filtering

```
Show me all estimates (status: estimate)
```

```
Find work in progress (status: WIP)
```

```
List completed but not posted orders
```

### Combined Filters

```
Find repair orders for shop 123, status complete,
created this month, over $500
```

```
Search for vehicles: make Toyota, year 2020+,
with repair orders in last 6 months
```

## Tips for Better Results

### Be Specific

❌ **Vague**: "Show me customers"
✅ **Specific**: "Show me the first 50 customers from shop 123"

### Use Natural Language

❌ **Too technical**: "Execute get_repair_orders with status=2"
✅ **Natural**: "Show me repair orders in progress"

### Break Down Complex Tasks

❌ **Too complex**: "Analyze everything about shop performance"
✅ **Broken down**:
1. "Show me total repair orders this month"
2. "Calculate average repair order value"
3. "List top 5 customers by spend"

### Specify Shops

❌ **Ambiguous**: "Get customers"
✅ **Clear**: "Get customers from shop 123"

### Use Pagination

❌ **Too broad**: "Show me all vehicles"
✅ **Paginated**: "Show me the first 100 vehicles"

## Error Handling

### Common Errors and Solutions

**Error**: "Shop not found"
```
Make sure you're using the correct shop ID.
Try: "Show me all my shops" first
```

**Error**: "Authentication failed"
```
Check your credentials and try restarting Claude Desktop
```

**Error**: "Customer not found"
```
Verify the customer ID exists:
"Search for customers named [name]" to find the ID
```

## Next Steps

- Explore [Available Tools](../tools/index.md)
- Review [Configuration Options](../configuration/index.md)
- See [Installation Guide](../installation/index.md)

## Need Help?

- [GitHub Issues](https://github.com/beetlebugorg/tekmetric-mcp/issues)
- [Tekmetric API Documentation](https://api.tekmetric.com)
- [MCP Protocol Docs](https://modelcontextprotocol.io)
