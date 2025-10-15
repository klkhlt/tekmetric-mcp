# Tekmetric MCP Server

Ask questions about your shop data using natural language. Get instant answers about customers, vehicles, repair orders, appointments, and more - directly in Claude.

> **⚠️ Independent Community Project**
>
> This is an **independent, community-built tool** that integrates with Tekmetric's API. It is **not an official Tekmetric product** and is not affiliated with, endorsed by, or supported by Tekmetric, Inc.
>
> For official Tekmetric products and support, visit [tekmetric.com](https://tekmetric.com).

## What You Can Do

### Ask About Your Shop

```
Show me today's appointments
```

```
What repair orders are in progress?
```

```
How many customers do I have?
```

### Find Customers & Vehicles

```
Find customers named John Smith
```

```
Search for a 2020 Honda Accord
```

```
Look up vehicle with VIN 1HGBH41JXMN109186
```

### Track Work & Revenue

```
Show me all estimates from last week
```

```
What's our total revenue for the month?
```

```
Find repair orders over $1000
```

### Manage Your Schedule

```
Who's scheduled for tomorrow?
```

```
What appointments does customer 456 have?
```

```
Show me this week's workload
```

### Analyze Your Business

```
Which customers haven't been in for 6 months?
```

```
What's the average repair order value this quarter?
```

```
Show me the most common services we perform
```

## Features

- **Natural Language** - Just ask questions, no commands to memorize
- **Real-Time Data** - Direct access to your Tekmetric data
- **Smart Search** - Find anything by name, VIN, date, status, and more
- **Automatic Enrichment** - Appointments show customer names and vehicle details, not just IDs
- **Proper Formatting** - Prices display as $255.55, not 25555

## 🚀 Quick Start

### 1. Get API Credentials

1. Log into your Tekmetric account
2. Navigate to **Settings → API Access**
3. Create a new API application
4. Copy your Client ID and Client Secret

### 2. Install

#### Desktop Extension (Easiest)

1. Download `tekmetric-mcp.mcpb` from [releases](https://github.com/beetlebugorg/tekmetric-mcp/releases)
2. Open the `.mcpb` file with Claude Desktop
3. Configure your API credentials
4. Start asking questions!

#### Manual Installation

Add to your Claude Desktop config (`~/Library/Application Support/Claude/claude_desktop_config.json` on macOS):

```json
{
  "mcpServers": {
    "tekmetric": {
      "command": "/absolute/path/to/tekmetric-mcp",
      "args": ["serve"],
      "env": {
        "TEKMETRIC_CLIENT_ID": "your_client_id",
        "TEKMETRIC_CLIENT_SECRET": "your_client_secret",
        "TEKMETRIC_BASE_URL": "https://api.tekmetric.com",
        "TEKMETRIC_DEFAULT_SHOP_ID": "123"
      }
    }
  }
}
```

### 3. Restart Claude Desktop

Quit and reopen Claude Desktop. Look for the 🔌 icon to verify connection.

## What Data You Can Access

**Shops** - View shop details, hours, labor rates, and settings

**Customers** - Search by name, email, phone; view contact info and service history

**Vehicles** - Find by VIN, license plate, make/model; see service records

**Repair Orders** - Filter by status, date, customer, vehicle; track work and revenue

**Jobs** - View individual services, labor, parts, and technician assignments

**Appointments** - Search by date, customer, vehicle; manage scheduling

**Employees** - List technicians and staff, filter by role and status

**Inventory** - Check parts stock levels and availability (Beta)

## Example Queries

### Daily Operations

```
What's on the schedule today?
```

```
Show me repair orders that need to be completed this week
```

```
Find the customer with phone 555-1234
```

### Business Intelligence

```
How many new customers did we get last month?
```

```
What's our average repair order value?
```

```
Which services generate the most revenue?
```

### Customer Service

```
When was this customer last here?
```

```
What work have we done on this vehicle?
```

```
Show me all open estimates for customer John Smith
```

## 📚 Documentation

- **[Full Documentation](https://beetlebugorg.github.io/tekmetric-mcp/)** - Complete guide with examples
- **[Tool Reference](https://beetlebugorg.github.io/tekmetric-mcp/tools/)** - All available tools and parameters
- **[Contributing](./docs/docs/contributing.md)** - How to contribute (AI-first approach)

## 🔒 Security & Privacy

- **Read-Only Access** - Cannot modify your Tekmetric data
- **Your Credentials** - API keys stored locally, never sent to third parties
- **Direct Connection** - Communicates directly with Tekmetric API
- **Open Source** - Full code transparency

## 🤝 Contributing

This is an AI-first project! We prefer clear requirements documents over code submissions. See our [Contributing Guide](./docs/docs/contributing.md) to learn how to contribute using requirements instead of pull requests.

## 🐛 Found a Bug?

[Open an issue](https://github.com/beetlebugorg/tekmetric-mcp/issues) with:
- What you asked
- What you expected
- What actually happened

---

## Technical Details

### Requirements

- Go 1.23+ (for building from source)
- Tekmetric API credentials
- Claude Desktop app

### Build from Source

```bash
git clone https://github.com/beetlebugorg/tekmetric-mcp.git
cd tekmetric-mcp
make build
```

### Configuration

Environment variables:

| Variable | Description | Required |
|----------|-------------|----------|
| `TEKMETRIC_CLIENT_ID` | API Client ID | Yes |
| `TEKMETRIC_CLIENT_SECRET` | API Client Secret | Yes |
| `TEKMETRIC_BASE_URL` | API URL (sandbox or production) | No |
| `TEKMETRIC_DEFAULT_SHOP_ID` | Default shop ID | No |
| `TEKMETRIC_DEBUG` | Enable debug logging | No |

### Technology Stack

- **Go 1.23+** - Fast, compiled binary
- **MCP SDK** - Model Context Protocol
- **OAuth 2.0** - Secure authentication
- **Rate Limiting** - Automatic retry with backoff

### Troubleshooting

**Connection Issues:**
- Verify absolute paths in config
- Check binary is executable
- View Claude Desktop Developer Tools

**Authentication Failed:**
- Verify Client ID and Client Secret
- Check correct base URL (sandbox vs production)
- Try regenerating API credentials

**Rate Limiting:**
- Server automatically retries with backoff
- Reduce request frequency if needed

Enable debug mode:
```bash
tekmetric-mcp -d serve
```

## 📄 License

MIT License - see [LICENSE](LICENSE) file

**Disclaimer**: Provided "as-is" without warranty. Test in sandbox first. Not affiliated with Tekmetric, Inc.

---

**Built with 🤖 + 💻 | [GitHub](https://github.com/beetlebugorg/tekmetric-mcp) | [Docs](https://beetlebugorg.github.io/tekmetric-mcp/)**
