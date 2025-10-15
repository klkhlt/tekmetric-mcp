---
sidebar_position: 2
---

# Customers

Access and search customer records and contact information.

## What You Can Ask

### Find Customers

```
Find customers named John Smith
```

```
Search for customers with last name Johnson
```

```
Find customer with email john@example.com
```

```
Search for customers with phone number 555-1234
```

### Get Specific Customer

```
Show me customer 12345
```

```
Get details for customer ID 456
```

### Browse Customers

```
Show me customers at shop 123
```

```
List all customers
```

```
Show me the first 50 customers
```

### Customer Analysis

```
Show me customers created in the last 30 days
```

```
Which customers have email addresses?
```

```
Find customers in Springfield
```

### Examples by Role

**Service Advisor:**
```
Find customer John Smith's contact information
```

```
What's the phone number for customer 789?
```

**Front Desk:**
```
Look up customer by phone 555-1234
```

```
Find customer with email sarah@gmail.com
```

**Manager:**
```
How many new customers this month?
```

```
Show me customers without email addresses
```

## What You'll Get

Each customer includes:
- Customer ID
- First and last name
- Email address
- Phone number(s)
- Street address
- Account settings and credit limit
- Customer type
- Created and updated dates
- Associated vehicles

## Technical Reference

### Tool: `customers`

Search for customers by name, email, phone, or get a specific customer by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No* | Get specific customer by ID | - |
| `query` | string | No* | Search by name, email, or phone | - |
| `shop` | number | No | Shop ID | Default shop |
| `limit` | number | No | Max results | 10 |

**Note:** Provide either `id` or `query`.

### Notes

- Search is case-insensitive
- Partial matches are supported
- Credit limit is in cents (100000 = $1000.00)
- Multiple phone numbers possible per customer

## Related Tools

- [Vehicles](./vehicles.md) - Customer vehicles
- [Repair Orders](./repair-orders.md) - Service history
- [Appointments](./appointments.md) - Scheduled visits
