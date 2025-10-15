---
sidebar_position: 7
---

# Shop Tools

Access your shop information, settings, and details.

## What You Can Ask

### List All Shops

```
Show me all my shops
```

```
What shops do I have access to?
```

```
List all shops
```

### Find Specific Shop

```
Find shops with "Main Street" in the name
```

```
Search for shops in Springfield
```

```
What's my shop ID?
```

### Check Shop Details

```
Show me details for Main Street Auto
```

```
What are the hours for shop 123?
```

```
What's the labor rate at my shop?
```

### Multiple Locations

```
Show me all shops
```

```
Which shops are in California?
```

```
List shops by location
```

## What You'll Get

Each shop includes:
- **Shop ID** - Unique identifier (needed for other queries)
- Shop name
- Full address
- Phone and email
- Business hours
- Labor rates
- Tax rates
- Timezone

## Why Shop ID Matters

Your **Shop ID** is required for most other queries:
- Finding customers at a specific shop
- Viewing repair orders for a shop
- Listing appointments at a shop
- Checking inventory by location

## Tips

- **Save your shop ID** for future queries
- **Sandbox environment**: Shop IDs are usually `2` or `3`
- **Production environment**: Your real shop ID is assigned by Tekmetric

## Technical Reference

### Tool: `shops`

Search for shops by name or list all accessible Tekmetric shops.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `query` | string | No | Search shops by name | - |
| `limit` | number | No | Max results | 10 |

## Related Tools

- [Customers](./customers.md) - Customer records
- [Repair Orders](./repair-orders.md) - Service history
- [Employees](./employees.md) - Staff information
