---
sidebar_position: 1
---

# Shop Tools

The `shops` tool provides access to your shop information and settings.

## shops

Search for shops by name or list all accessible Tekmetric shops.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `query` | string | No | Search shops by name | - |
| `limit` | number | No | Maximum results to return | 10 |

### How to Use

#### List All Shops

```
Show me all my shops
```

```
List all shops
```

```
What shops do I have access to?
```

#### Search for Shops

```
Find shops with "Main Street" in the name
```

```
Search for shops in Springfield
```

### What You'll Get

Shop information includes:
- Shop ID (important for other queries)
- Shop name
- Address (street, city, state, zip)
- Contact info (phone, email)
- Business hours
- Settings (labor rates, tax rates, timezone)

## What You Can Ask

### See All Your Shops

```
Show me all my shops
```

```
What shops do I have access to?
```

```
List all shops
```

### Search for a Specific Shop

```
Find shops with "Main Street" in the name
```

```
Search for shops in Springfield
```

## What Information You'll Get

When you look up shops, you'll see:

- **Shop ID** - The unique number for this shop (you'll need this for other queries)
- **Shop Name** - The name of your shop
- **Address** - Street address, city, state, zip code
- **Contact Info** - Phone number and email
- **Hours** - When the shop is open
- **Settings** - Labor rates, tax rates, timezone

## Why This Matters

**Your Shop ID is important!** You'll need it when asking about:
- Customers at a specific shop
- Vehicles serviced at a shop
- Repair orders for a shop
- Appointments at a shop

## Examples

### Find Your Shop ID

```
What's my shop ID?
```

The AI will show you all shops you have access to with their IDs.

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

### Multiple Shops

If you manage multiple locations:

```
Show me all shops
```

```
Which shops are in California?
```

```
List shops by location
```

## Tips

- **Save your shop ID** - Write it down for future questions
- **Test environment** - Sandbox shop IDs are usually `2` or `3`
- **Production environment** - Your real shop ID is assigned by Tekmetric

## Need More Help?

- See examples of [customer queries](./customers.md)
- Learn about [repair orders](./repair-orders.md)
- Check out [common questions](../examples/index.md)
