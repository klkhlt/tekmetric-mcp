---
sidebar_position: 4
---

# Inventory Tools

:::info Beta Feature
Inventory tools use Tekmetric's Beta API. Features may be limited or change over time.
:::

Access and search parts inventory and stock levels.

## What You Can Ask

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

```
Search inventory for Honda Accord brake pads
```

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

### Examples by Role

**Parts Counter:**
```
Do you have brake pads for a 2019 Honda Accord?
```

**Manager:**
```
What parts need to be reordered?
```

```
Show me parts with quantity less than 5
```

## What You'll Get

Each inventory item includes:
- Part number (SKU)
- Description
- Quantity in stock
- Cost (what you paid)
- Price (what you sell for)
- Location in shop
- Supplier
- Reorder point

## Important Notes

- Inventory data depends on your Tekmetric setup
- Best for shops actively tracking parts in Tekmetric
- May have limited data if you don't use inventory management
- Beta API means features may change

## Technical Reference

### Tool: `inventory`

Search inventory parts by name or part number, or list all available parts.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `query` | string | No | Search by name or part number | - |
| `shop` | number | No | Shop ID | Default shop |
| `limit` | number | No | Max results | 20 |

**Note:** Search queries against part description, part number, and brand.

## Related Tools

- [Repair Orders](./repair-orders.md) - See parts usage
- [Jobs](./jobs.md) - Parts used on jobs
- [Shops](./shops.md) - Inventory by location
