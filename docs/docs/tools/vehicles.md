---
sidebar_position: 3
---

# Vehicle Tools

The `vehicles` tool provides access to vehicle records with flexible search options.

## vehicles

Search for vehicles by VIN, license plate, make/model, or get a specific vehicle by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No* | Get specific vehicle by ID | - |
| `query` | string | No* | Search by VIN, license plate, or make/model | - |
| `shop` | number | No | Shop ID | Default shop |
| `limit` | number | No | Maximum results to return | 10 |

**Note:** You must provide either `id` or `query` parameter.

### How to Use

#### Get a Specific Vehicle

```
Show me vehicle ID 789
```

```
Get details for vehicle 456
```

#### Search for Vehicles

```
Find vehicle with VIN 1HGBH41JXMN109186
```

```
Look up license plate ABC123
```

```
Find all Toyota Camrys
```

```
Show me 2020 Honda Accords
```

```
Search for Ford F-150 trucks
```

### What You'll Get

Vehicle information includes:
- Make, model, year
- VIN (Vehicle Identification Number)
- License plate
- Color
- Owner information
- Mileage
- Service history

## What You Can Ask

### Find a Specific Vehicle

```
Find vehicle with VIN 1HGBH41JXMN109186
```

```
Show me vehicle ID 789
```

```
Look up license plate ABC123
```

### Search by Make/Model/Year

```
Find all Toyota Camrys
```

```
Show me 2020 Honda Accords
```

```
Search for Ford F-150 trucks
```

```
Find all vehicles from 2018 or newer
```

### Browse Vehicles

```
Show me the first 50 vehicles at shop 123
```

```
List all vehicles
```

## What Information You'll Get

When you look up a vehicle, you'll see:

- **Make, Model, Year** - What kind of vehicle it is
- **VIN** - Vehicle Identification Number
- **License Plate** - Tag number
- **Color** - Vehicle color
- **Owner** - Customer who owns it
- **Mileage** - Last recorded odometer reading
- **Service History** - Past repair orders for this vehicle

## Real-World Uses

### Customer Calls In

```
Customer says they drive a silver 2019 Toyota Camry,
what's their VIN?
```

```
Find the vehicle with license plate XYZ789
```

### Service Planning

```
Show me all vehicles due for service
```

```
Find vehicles that haven't been serviced in 6 months
```

```
What vehicles are scheduled this week?
```

### Inventory Analysis

```
How many Honda vehicles do we service?
```

```
What's the most common vehicle make in our system?
```

```
Show me all trucks
```

## Search Tips

**Be flexible with your questions:**

✅ Good:
- "Find Toyota Camry"
- "Show me vehicles from 2020"
- "Search for VIN ending in 1234"

❌ Too specific (might not find anything):
- "2020 Toyota Camry LE 4-Door Sedan Silver"

**VIN Searches:**
- Full VIN works best
- Partial VIN searches supported
- Case doesn't matter

**License Plates:**
- Enter as shown on the plate
- Spaces and dashes usually don't matter

## Examples by Situation

### Front Desk Scenario

Customer calls: *"I'm calling about my red Honda"*

```
Find red Honda vehicles
```

Then narrow it down by year or owner name.

### Service Advisor Scenario

Need to schedule an oil change for a regular customer:

```
Show me John Smith's vehicles
```

```
What was the last service date for vehicle 456?
```

### Shop Manager Scenario

Planning for the day:

```
What vehicles have appointments today?
```

```
Show me vehicles with open repair orders
```

## Connecting to Other Information

Once you find a vehicle, you can:

```
Show me repair history for vehicle 789
```

```
What's the owner of vehicle 456?
```

```
Find all repair orders for this VIN
```

## Need More Help?

- Look up [customer information](./customers.md)
- Check [repair orders](./repair-orders.md)
- See [usage examples](../examples/index.md)
