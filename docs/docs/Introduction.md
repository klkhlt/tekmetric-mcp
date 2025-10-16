---
sidebar_position: 1
slug: /
---

<div style={{marginBottom: '2rem', display: 'flex', alignItems: 'center', gap: '1rem'}}>
  <img src="/tekmetric-mcp/img/robot.png" alt="Tekmetric MCP Robot" style={{width: '120px', height: '120px'}} />
  <h1 style={{margin: 0, fontSize: '2.5rem'}}>Tekmetric MCP Server</h1>
</div>

# Introduction

Tekmetric MCP Server is a **Model Context Protocol (MCP)** server that provides tools for accessing the Tekmetric shop management API. Built with Go, it offers fast, reliable access to shop data, customers, vehicles, repair orders, and more.

<div style={{textAlign: 'center', margin: '2rem 0'}}>
  <img src="/tekmetric-mcp/img/example.png" alt="Example: Natural language query about repair orders" style={{maxWidth: '800px', width: '100%', borderRadius: '8px', boxShadow: '0 4px 6px rgba(0, 0, 0, 0.1)'}} />
  <p style={{fontSize: '0.9em', color: '#666', marginTop: '0.5rem'}}>Ask questions in plain English and get instant answers from your Tekmetric data</p>
</div>

:::info Independent Tool & AI-First Project
This is an **independent, community-built tool** that works with Tekmetric's API. It is **not an official Tekmetric product** and is not affiliated with, endorsed by, or supported by Tekmetric.

This is an **AI-first project** - built primarily with AI assistance (Claude). We welcome contributions in the form of clear requirements documents rather than code. See [Contributing](./contributing.md) for details.

For official Tekmetric products and support, visit [tekmetric.com](https://tekmetric.com).
:::

## What is MCP?

The [Model Context Protocol](https://modelcontextprotocol.io) is an open protocol that standardizes how applications provide context to Large Language Models (LLMs). MCP servers expose tools and data that AI assistants like Claude can use to perform tasks.

## What is Tekmetric?

[Tekmetric](https://www.tekmetric.com) is a cloud-based shop management software for automotive repair shops. It helps shops manage customers, vehicles, repair orders, appointments, inventory, and more.

## What Does This Server Do?

This MCP server acts as a bridge between Claude Desktop and the Tekmetric API, allowing you to:

- **Query shop data** - Access information about shops, customers, vehicles, and repair orders
- **Search records** - Find customers by name/email, vehicles by VIN/make/model, repair orders by status
- **Look up details** - Get contact info, vehicle specs, appointment times, job assignments
- **Read-only access** - Safely explore your data without risk of accidental modifications

## Key Features

- ✅ **8 Tool Categories** - Comprehensive access to all Tekmetric resources
- ✅ **OAuth 2.0 Authentication** - Automatic token management and renewal
- ✅ **Rate Limiting** - Built-in exponential backoff with jitter
- ✅ **Zero Config** - Works with environment variables
- ✅ **Lightweight** - Single binary, minimal dependencies
- ✅ **Cross-Platform** - Linux, macOS, Windows support
- ✅ **Read-Only** - Safe exploration without accidental modifications

## What This Is For

**Best for tactical, day-to-day queries:**

### Customer Lookups
```
Find customer John Smith's phone number
Look up email for customer ID 456
Search for customers named Mike
```

### Vehicle Information
```
Look up VIN 1HGBH41JXMN109186
Find all 2020 Honda Accords
What vehicles does customer 123 own?
```

### Daily Operations
```
Show me today's appointments
What repair orders are in progress?
Who's assigned to job 789?
```

### Specific Record Searches
```
Find repair order #1234
Show estimates from last week
Get details for appointment 567
```

**NOT designed for:**
- Financial reporting or accounting (use Tekmetric's built-in reports)
- Revenue calculations or profit analysis
- Large-scale data exports or bulk operations
- Business intelligence dashboards
- Comprehensive trend analysis across large datasets

:::info Result Limits
Results are limited to **25 records per query** to keep responses fast and manageable. For comprehensive reports, financial data, or bulk exports, use Tekmetric's native reporting and export tools.
:::

## Architecture

The server is structured as follows:

```
tekmetric-mcp
├── CLI entry point (main.go)
├── MCP server (internal/mcp/)
│   ├── Server implementation
│   └── Tool registry and handlers
├── Tekmetric API client (internal/tekmetric/)
│   ├── OAuth2 authentication
│   ├── API request handling
│   └── Data models
├── Configuration (internal/config/)
│   └── Environment and file-based config
└── Rate limiting (pkg/ratelimit/)
    └── Exponential backoff with jitter
```

## Technology Stack

- **Go 1.23+** - Fast, compiled, single binary
- **MCP SDK** - [mcp-go](https://github.com/mark3labs/mcp-go) v0.7.0
- **CLI Framework** - [Kong](https://github.com/alecthomas/kong) v0.9.0
- **Configuration** - [Viper](https://github.com/spf13/viper) v1.19.0
- **HTTP Client** - Standard library with custom OAuth implementation

## Security & Privacy

- **Read-only access** - This tool can only *view* your data, never change it
- **Secure connection** - All communication is encrypted
- **Your credentials** - Uses your official Tekmetric API credentials
- **Local execution** - Runs on your machine, not in the cloud

:::warning Disclaimer
This tool is provided "as-is" without warranty of any kind. While it only has read-only access to your Tekmetric data, you should:
- Test thoroughly in a sandbox environment first
- Verify all information before making business decisions
- Keep your API credentials secure
- Review the [security best practices](./configuration/index.md)

Use of this tool is at your own risk. The authors are not responsible for any data loss, security issues, or business decisions made based on information provided by this tool.
:::

## Getting Started

Ready to get started? Head over to the [Installation](./installation/index.md) guide to set up the server.

## Contributing

This is an AI-first project! We welcome contributions in the form of clear requirements documents. See the [Contributing Guide](./contributing.md) to learn how to help improve this project.

## Need Help?

- **Questions?** Check the [Tools Documentation](./tools/index.md) and [Examples](./examples/index.md)
- **Issues**: [GitHub Issues](https://github.com/beetlebugorg/tekmetric-mcp/issues)
- **Tekmetric API**: [api.tekmetric.com](https://api.tekmetric.com)
- **MCP Protocol**: [modelcontextprotocol.io](https://modelcontextprotocol.io)

## License

MIT License - see [LICENSE](https://github.com/beetlebugorg/tekmetric-mcp/blob/main/LICENSE) file for details.

**Robot illustration:** © 2019 Vijay Verma for [illlustrations.co](https://illlustrations.co) - [MIT License](https://illlustrations.co/license/)

**Not affiliated with Tekmetric, Inc.**
