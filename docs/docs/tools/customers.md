---
sidebar_position: 2
---

# Customer Tools

The `customers` tool provides access to customer records with flexible search and retrieval options.

## customers

Search for customers by name, email, phone, or get a specific customer by ID.

### Parameters

| Parameter | Type | Required | Description | Default |
|-----------|------|----------|-------------|---------|
| `id` | number | No* | Get specific customer by ID | - |
| `query` | string | No* | Search customers by name, email, or phone | - |
| `shop` | number | No | Shop ID | Default shop |
| `limit` | number | No | Maximum results to return | 10 |

**Note:** You must provide either `id` or `query` parameter.

### How to Use

#### Get a Specific Customer

When you know the customer ID:

```
Show me customer 12345
```

```
Get details for customer ID 456
```

#### Search for Customers

Search by name, email, or phone:

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
Show me customers at shop 123
```

### What You'll Get

Customer information includes:
- Personal information (first name, last name)
- Contact details (phone, email)
- Address
- Account settings
- Customer type
- Created/updated dates
- Associated vehicles

## Response Format

### Customer Object

```json
{
  "id": 12345,
  "shopId": 123,
  "firstName": "John",
  "lastName": "Smith",
  "email": "john.smith@example.com",
  "phone": [
    {
      "type": "mobile",
      "number": "(555) 123-4567",
      "primary": true
    }
  ],
  "address": {
    "street": "123 Main St",
    "city": "Springfield",
    "state": "IL",
    "zip": "62701",
    "fullAddress": "123 Main St, Springfield, IL 62701"
  },
  "eligibleForAccountsReceivable": true,
  "creditLimit": 100000,
  "okForMarketing": true,
  "createdDate": "2024-01-15T10:30:00Z",
  "updatedDate": "2024-03-20T14:45:00Z"
}
```

## Common Use Cases

### Finding Customers

```
Find all customers with last name Johnson
```

```
Search for customers in Springfield
```

### Customer Analysis

```
Show me customers created in the last 30 days
```

```
Which customers have email addresses?
```

### Customer Service

```
Find customer John Smith's contact information
```

```
What's the phone number for customer 789?
```

## Notes

- Email and phone fields may be empty if not provided
- Credit limit is in cents (100000 = $1000.00)
- Multiple phone numbers possible per customer
- Search is case-insensitive
- Partial matches are supported

## Next Steps

- [Vehicle Tools](./vehicles.md) - See customer vehicles
- [Repair Order Tools](./repair-orders.md) - See customer repair history
- [Examples](../examples/index.md) - More usage examples
