---
sidebar_position: 8
---

# Vehicle Tools

Access and search vehicle records by VIN, license plate, or make/model.

## What You Can Ask

### Find by VIN or License Plate

```
Find vehicle with VIN 1HGBH41JXMN109186
```

```
Look up license plate ABC123
```

```
Search for VIN ending in 1234
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

### Get Specific Vehicle

```
Show me vehicle ID 789
```

```
Get details for vehicle 456
```

### Browse Vehicles

```
Show me the first 50 vehicles at shop 123
```

```
List all vehicles
```

### Customer Service

**Customer calls in:**
```
Customer says they drive a silver 2019 Toyota Camry, what's their VIN?
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

### Examples by Role

**Front Desk:**
```
Find red Honda vehicles
```

```
Customer calling about a 2020 Camry, find it
```

**Service Advisor:**
```
Show me John Smith's vehicles
```

```
What was the last service date for vehicle 456?
```

**Shop Manager:**
```
What vehicles have appointments today?
```

```
Show me vehicles with open repair orders
```

## What You'll Get

Each vehicle includes:
- Vehicle ID
- Make, model, year
- VIN (Vehicle Identification Number)
- License plate
- Color
- Owner information
- Last recorded mileage
- Service history

## Search Tips

**VIN Searches:**
- Full VIN works best
- Partial VIN searches supported
- Case doesn't matter

**License Plates:**
- Enter as shown on the plate
- Spaces and dashes usually don't matter

**Make/Model:**
- Be flexible: "Find Toyota Camry" works better than "2020 Toyota Camry LE 4-Door Sedan Silver"
- Partial matches supported

## Technical Reference

### Tool: `vehicles`

Search for vehicles by VIN, license plate, make/model, or get a specific vehicle by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No* | Get specific vehicle by ID | - |
| `query` | string | No* | Search by VIN, plate, or make/model | - |
| `shop` | number | No | Shop ID | Default shop |
| `limit` | number | No | Max results | 10 |

**Note:** Provide either `id` or `query`.

## Related Tools

- [Customers](./customers.md) - Vehicle owners
- [Repair Orders](./repair-orders.md) - Service history
- [Appointments](./appointments.md) - Scheduled service
