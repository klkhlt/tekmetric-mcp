---
sidebar_position: 8
---

# Inventory Tools

The `inventory` tool provides access to parts and inventory information.

:::info Beta Feature
Inventory tools use Tekmetric's Beta API. Some features may be limited or change over time.
:::

## inventory

Search inventory parts by name or part number, or list all available parts.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `query` | string | No | Search parts by name or part number | - |
| `shop` | number | No | Shop ID | Default shop |
| `limit` | number | No | Maximum results to return | 20 |

### How to Use

#### List All Parts

```
Show me all inventory parts at shop 123
```

```
List parts in stock
```

#### Search for Parts

```
Search for oil filters
```

```
Find part number ABC123
```

```
Do we have brake pads in stock?
```

### What You'll Get

Inventory information includes:
- Part number (SKU)
- Description
- Quantity in stock
- Cost
- Price
- Location in shop
- Supplier
- Reorder point

**Note:** The search queries against part description, part number, and brand.

## What You Can Ask

### Browse Inventory

```
Show me all inventory parts at shop 123
```

```
List parts in stock
```

```
What inventory do we have?
```

## What Information You'll Get

When you look up inventory, you'll see:

- **Part Number** - Manufacturer or internal part number
- **Description** - What the part is
- **Quantity** - How many in stock
- **Cost** - What we paid for it
- **Price** - What we sell it for
- **Location** - Where it's stored in the shop
- **Supplier** - Where we get it from
- **Reorder Point** - When to order more

## Real-World Uses

### Check Stock Levels

```
Do we have brake pads in stock?
```

```
Show me parts with low quantity
```

```
What needs to be reordered?
```

### Find Parts

```
Search for oil filters
```

```
Find part number ABC123
```

```
Do we have this part in stock?
```

### Inventory Management

```
Show me all parts under reorder point
```

```
List parts we're out of
```

```
What parts are overstocked?
```

## Examples

### Parts Counter Scenario

Customer: *"Do you have brake pads for a 2019 Honda Accord?"*

```
Search inventory for Honda Accord brake pads
```

### Ordering Parts

```
What parts need to be reordered?
```

```
Show me parts with quantity less than 5
```

```
List parts we're low on
```

### Cost Analysis

```
What's our total inventory value?
```

```
Show me most expensive parts
```

```
List parts by cost
```

## Tips

- Inventory data depends on how you use Tekmetric's inventory features
- Beta API means features may change
- Best for shops actively tracking parts in Tekmetric
- May have limited data if you don't use inventory management

## Limitations

Since this is a Beta feature:
- Not all shops use Tekmetric inventory
- Data may be limited
- Features may change as Tekmetric improves the API

## Need More Help?

- Learn about [repair orders](./repair-orders.md)
- Check [jobs requiring parts](./jobs.md)
- View [usage examples](../examples/index.md)
- Contact Tekmetric support about inventory features
