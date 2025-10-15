# Tekmetric MCP Server

A Model Context Protocol (MCP) server providing AI assistants with tools to access the Tekmetric shop management API. Built with Go for fast, reliable access to shop data, customers, vehicles, repair orders, and more.

> **⚠️ Independent Community Project**
>
> This is an **independent, community-built tool** that integrates with Tekmetric's API. It is **not an official Tekmetric product** and is not affiliated with, endorsed by, or supported by Tekmetric, Inc.
>
> For official Tekmetric products and support, visit [tekmetric.com](https://tekmetric.com).

> **🤖 AI-First Development**
>
> This project is built with AI assistance and embraces an AI-first development approach. We welcome contributions as **clear requirements documents** rather than code submissions.
>
> See our [Contributing Guide](./docs/docs/contributing.md) to learn how to effectively contribute using requirements instead of pull requests.

## Features

- **8 Tool Categories**: Comprehensive access to all Tekmetric resources
- **OAuth 2.0 Authentication**: Automatic token management and renewal
- **Rate Limiting**: Built-in exponential backoff with jitter
- **Zero Config**: Works with environment variables
- **Lightweight**: Single binary, minimal dependencies
- **Cross-Platform**: Linux, macOS, Windows support

## 🚀 Quick Start

### 1. Get Credentials

Get your API credentials from Tekmetric:

1. Log into your Tekmetric account
2. Navigate to **Settings → API Access**
3. Create a new API application
4. Copy your Client ID and Client Secret

### 2. Install

#### Option A: Desktop Extension (Recommended)

Download and install the pre-built Desktop Extension:

1. Download `tekmetric-mcp.mcpb` from [releases](https://github.com/beetlebugorg/tekmetric-mcp/releases)
2. Open the `.mcpb` file with Claude Desktop
3. Configure your API credentials
4. Start using immediately

#### Option B: Build from Source

```bash
go install github.com/beetlebugorg/tekmetric-mcp@latest
```

Or clone and build:

```bash
git clone https://github.com/beetlebugorg/tekmetric-mcp.git
cd tekmetric-mcp
make build
```

### 3. Configure

Add to your Claude Desktop configuration:

**macOS**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

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

**Important**: Use absolute paths, not relative paths or `~`

### 4. Restart Claude Desktop

Quit and reopen Claude Desktop. Look for the 🔌 icon to verify connection.

## 🔧 Available Tools

The server provides tools organized by resource type:

### Shop Tools

- `get_shops` - List all accessible shops
- `get_shop` - Get details for a specific shop

### Customer Tools

- `get_customers` - List customers with pagination
- `get_customer` - Get customer details by ID
- `search_customers` - Search customers by name, email, or phone

### Vehicle Tools

- `get_vehicles` - List vehicles with pagination
- `get_vehicle` - Get vehicle details by ID
- `search_vehicles` - Search vehicles by VIN, make, model, or year

### Repair Order Tools

- `get_repair_orders` - List repair orders with pagination
- `get_repair_order` - Get repair order details by ID
- `search_repair_orders` - Search by RO number, status, date range

### Job Tools

- `get_jobs` - List jobs with pagination
- `get_job` - Get job details by ID

### Appointment Tools

- `get_appointments` - List appointments with pagination
- `get_appointment` - Get appointment details by ID

### Employee Tools

- `get_employees` - List employees with pagination
- `get_employee` - Get employee details by ID

### Inventory Tools

- `get_inventory` - List inventory parts with pagination (Beta)

## 💡 Usage Examples

### Basic Queries

```
Show me all my shops
```

```
Get customers from shop 123
```

```
Find repair orders created in the last week
```

### Advanced Search

```
Search for customers with email containing "@example.com"
```

```
Find all vehicles with year 2020 or newer
```

```
Search repair orders with status "complete" from last month
```

### Data Analysis

```
Analyze customer purchase patterns for shop 123
```

```
Show me the top 10 most expensive repair orders this quarter
```

```
Find vehicles that haven't had service in 6 months
```

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Default | Required |
|----------|-------------|---------|----------|
| `TEKMETRIC_CLIENT_ID` | API Client ID | - | Yes |
| `TEKMETRIC_CLIENT_SECRET` | API Client Secret | - | Yes |
| `TEKMETRIC_BASE_URL` | API Base URL | `https://sandbox.tekmetric.com` | No |
| `TEKMETRIC_DEFAULT_SHOP_ID` | Default Shop ID | `0` | No |
| `TEKMETRIC_DEBUG` | Enable debug logging | `false` | No |

### Config File

Alternatively, create `~/.config/tekmetric-mcp/config.json`:

```json
{
  "tekmetric": {
    "base_url": "https://api.tekmetric.com",
    "client_id": "your_client_id",
    "client_secret": "your_client_secret",
    "default_shop_id": 123,
    "timeout_seconds": 30,
    "max_retries": 3,
    "max_backoff_sec": 60
  },
  "server": {
    "name": "tekmetric-mcp",
    "version": "0.1.0",
    "debug": false
  }
}
```

## 🌍 Environments

### Sandbox (Testing)

Use the sandbox for development and testing:

```bash
TEKMETRIC_BASE_URL=https://sandbox.tekmetric.com
TEKMETRIC_DEFAULT_SHOP_ID=2  # or 3
```

Sandbox shop IDs are typically `2` or `3`.

### Production

For production, update to your production credentials:

```bash
TEKMETRIC_BASE_URL=https://api.tekmetric.com
TEKMETRIC_DEFAULT_SHOP_ID=your_shop_id
```

## 💻 CLI Commands

```bash
# Show version information
tekmetric-mcp version

# Start MCP server (for Claude Desktop)
tekmetric-mcp serve

# Enable debug logging
tekmetric-mcp -d serve

# Override configuration
tekmetric-mcp serve \
  --client-id=your_id \
  --client-secret=your_secret \
  --base-url=https://api.tekmetric.com \
  --shop-id=123

# Show help
tekmetric-mcp --help
```

## 🛠️ Development

### Build

```bash
# Local build
make build

# Cross-platform builds (all architectures)
make build-all

# With version information
make VERSION=v1.0.0 build

# Install to $GOPATH/bin
make install
```

### Test

```bash
# Run tests (when available)
make test

# Run with debug logging
./tekmetric-mcp -d serve
```

### Clean

```bash
# Remove build artifacts
make clean
```

## 🏗️ Architecture

```
tekmetric-mcp/
├── main.go                      # CLI entry point
├── internal/
│   ├── config/
│   │   └── config.go           # Configuration management
│   ├── tekmetric/
│   │   ├── client.go           # API client with OAuth
│   │   └── models.go           # Data models
│   └── mcp/
│       ├── server.go           # MCP server
│       └── tools/
│           ├── tools.go        # Tool registry
│           ├── shops.go        # Shop tools
│           ├── customers.go    # Customer tools
│           ├── vehicles.go     # Vehicle tools
│           ├── repair_orders.go # Repair order tools
│           ├── jobs.go         # Job tools
│           ├── appointments.go # Appointment tools
│           ├── employees.go    # Employee tools
│           ├── inventory.go    # Inventory tools
│           ├── filters.go      # Filter utilities
│           └── helpers.go      # Helper functions
└── pkg/
    └── ratelimit/
        └── ratelimit.go        # Rate limiting with backoff
```

## 🔍 Troubleshooting

### Authentication Failed

**Problem**: "Authentication failed with status 401"

**Solutions**:
- Verify Client ID and Client Secret are correct
- Check you're using the correct base URL (sandbox vs production)
- Ensure credentials haven't been revoked in Tekmetric
- Try regenerating your API credentials

### Shop Not Found

**Problem**: "Shop ID not found" or "Access denied"

**Solutions**:
- Check the `scope` field in your access token
- Use `get_shops` tool to see available shop IDs
- For sandbox, use shop ID `2` or `3`
- Ensure your API application has permission for that shop

### Connection Issues

**Problem**: MCP server not connecting to Claude Desktop

**Solutions**:
- Verify absolute path in `claude_desktop_config.json`
- Check binary is executable: `chmod +x tekmetric-mcp`
- Test binary manually: `./tekmetric-mcp version`
- Look at Claude Desktop developer tools (View → Developer → Developer Tools)
- Check environment variables are set correctly

### Rate Limiting

**Problem**: "Too many requests" or "Rate limit exceeded"

**Solutions**:
- Server automatically retries with exponential backoff
- Wait a few minutes between heavy operations
- Reduce pagination size (use smaller `pageSize` values)
- Contact Tekmetric support for higher rate limits if needed

### Debug Mode

Enable debug logging to troubleshoot issues:

```bash
tekmetric-mcp -d serve
```

Or in Claude Desktop config:

```json
{
  "mcpServers": {
    "tekmetric": {
      "command": "/path/to/tekmetric-mcp",
      "args": ["-d", "serve"],
      "env": { ... }
    }
  }
}
```

## 🔒 Security Best Practices

- **Never commit credentials**: `.env` files are in `.gitignore`
- **Use environment variables**: Don't hardcode secrets
- **Separate environments**: Use different credentials for sandbox and production
- **Rotate regularly**: Update API credentials periodically
- **Minimal permissions**: Only grant necessary shop access
- **Secure storage**: Store credentials in encrypted environment or secrets manager

## ⏱️ API Rate Limits

The Tekmetric API enforces rate limits. This server handles rate limiting automatically with:

- **Exponential backoff**: Increases wait time after each failure
- **Jitter**: Adds randomness to prevent thundering herd
- **Configurable retries**: Set `max_retries` in configuration
- **Max backoff**: Caps maximum wait time (default 60 seconds)

Formula: `min(((2^n) + random_milliseconds), max_backoff)`

## 🚀 Technology

- **Go 1.23+**: Fast, compiled, single binary
- **MCP SDK**: [mcp-go](https://github.com/mark3labs/mcp-go) v0.7.0
- **CLI Framework**: [Kong](https://github.com/alecthomas/kong) v0.9.0
- **Configuration**: [Viper](https://github.com/spf13/viper) v1.19.0
- **HTTP Client**: Standard library with custom OAuth implementation

## 🤝 Contributing

**This is an AI-first project!** We believe clear requirements are more valuable than code. Instead of submitting pull requests, we prefer you describe what you want in a well-written requirements document.

### Why Requirements Over Code?

Good requirements can generate excellent code multiple times. They're reusable, testable, and help AI assistants build exactly what you need. Plus, you don't need to be a Go developer to contribute!

### How to Contribute

1. **📝 Write Requirements** - Describe what you want built (we have templates!)
2. **🐛 Report Issues** - Found a bug? Open an issue with reproduction steps
3. **💡 Request Features** - Share your ideas with example use cases
4. **💻 Submit Code** - Traditional pull requests are also welcome

### Learn More

Our [Contributing Guide](./docs/docs/contributing.md) includes:
- ✅ Requirements document template and examples
- ✅ How to write effective requirements for AI
- ✅ AI-first development principles
- ✅ Good vs bad requirements examples

## 📄 License

MIT License - see [LICENSE](LICENSE) file for details.

**Disclaimer**: This project is provided "as-is" without warranty of any kind. While the MCP server has read-only access to your Tekmetric data, you should test thoroughly in a sandbox environment and verify all information before making business decisions. Use at your own risk.

## ⚠️ Not Affiliated with Tekmetric

This is an **independent, community-built tool**. It is not an official Tekmetric product and is not affiliated with, endorsed by, or supported by Tekmetric, Inc.

For official Tekmetric products and support, visit [tekmetric.com](https://tekmetric.com).

## 📚 Support & Resources

- **📖 Documentation** - [Full documentation site](https://beetlebugorg.github.io/tekmetric-mcp/)
- **🐛 Issues** - [GitHub Issues](https://github.com/beetlebugorg/tekmetric-mcp/issues)
- **🔌 Tekmetric API** - [api.tekmetric.com](https://api.tekmetric.com)
- **🤖 MCP Protocol** - [modelcontextprotocol.io](https://modelcontextprotocol.io)

## 🙏 Acknowledgments

- Built for the [Model Context Protocol](https://modelcontextprotocol.io)
- Integrates with the [Tekmetric API](https://developer.tekmetric.com)
- Developed with AI assistance using Claude
- Inspired by the [go-dims](https://github.com/beetlebugorg/go-dims) architecture

---

**Built with 🤖 + 💻 | Not affiliated with Tekmetric, Inc.**
