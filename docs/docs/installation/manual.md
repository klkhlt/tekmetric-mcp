---
sidebar_position: 4
---

# Manual Installation

**For advanced users and other MCP clients.** This guide shows you how to manually configure any MCP-compatible application.

## What You'll Need

- An MCP-compatible client application
- Your Tekmetric API credentials ([see here](../installation/index.md#tekmetric-api-credentials))
- The tekmetric-mcp binary

## Step 1: Download the Binary

1. Go to [Releases](https://github.com/beetlebugorg/tekmetric-mcp/releases)
2. Download the right version:
   - **Mac (Apple Silicon)**: `tekmetric-mcp-darwin-arm64`
   - **Mac (Intel)**: `tekmetric-mcp-darwin-amd64`
   - **Windows**: `tekmetric-mcp-windows-amd64.exe`
   - **Linux (x86_64)**: `tekmetric-mcp-linux-amd64`
   - **Linux (ARM64)**: `tekmetric-mcp-linux-arm64`

3. Save it to a permanent location
4. Make it executable (Mac/Linux):
```bash
chmod +x /path/to/tekmetric-mcp
```

## Step 2: Configure Your MCP Client

The configuration format follows the MCP standard. Add this to your client's MCP configuration file:

```json
{
  "mcpServers": {
    "tekmetric": {
      "command": "/absolute/path/to/tekmetric-mcp",
      "args": ["serve"],
      "env": {
        "TEKMETRIC_CLIENT_ID": "your_client_id",
        "TEKMETRIC_CLIENT_SECRET": "your_client_secret",
        "TEKMETRIC_BASE_URL": "https://shop.tekmetric.com",
        "TEKMETRIC_DEFAULT_SHOP_ID": "123"
      }
    }
  }
}
```

### Configuration Parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `command` | Yes | Absolute path to tekmetric-mcp binary |
| `args` | Yes | Must be `["serve"]` |
| `env.TEKMETRIC_CLIENT_ID` | Yes | Your Tekmetric API Client ID |
| `env.TEKMETRIC_CLIENT_SECRET` | Yes | Your Tekmetric API Client Secret |
| `env.TEKMETRIC_BASE_URL` | Yes | API endpoint URL |
| `env.TEKMETRIC_DEFAULT_SHOP_ID` | No | Default shop for queries |

### Optional Environment Variables

```json
{
  "env": {
    "TEKMETRIC_CLIENT_ID": "...",
    "TEKMETRIC_CLIENT_SECRET": "...",
    "TEKMETRIC_BASE_URL": "https://shop.tekmetric.com",
    "TEKMETRIC_DEFAULT_SHOP_ID": "123",
    "TEKMETRIC_DEBUG": "false",
    "TEKMETRIC_TIMEOUT_SECONDS": "30",
    "TEKMETRIC_MAX_RETRIES": "3",
    "TEKMETRIC_MAX_BACKOFF_SEC": "60"
  }
}
```

| Variable | Default | Description |
|----------|---------|-------------|
| `TEKMETRIC_DEBUG` | `false` | Enable debug logging |
| `TEKMETRIC_TIMEOUT_SECONDS` | `30` | Request timeout |
| `TEKMETRIC_MAX_RETRIES` | `3` | Max retry attempts |
| `TEKMETRIC_MAX_BACKOFF_SEC` | `60` | Max wait between retries |

## Step 3: Test the Server

### Command Line Test

Before configuring your MCP client, test the server directly:

```bash
# Set environment variables
export TEKMETRIC_CLIENT_ID="your_id"
export TEKMETRIC_CLIENT_SECRET="your_secret"
export TEKMETRIC_BASE_URL="https://shop.tekmetric.com"

# Run the server
/path/to/tekmetric-mcp serve
```

The server should start and wait for MCP connections on stdio.

### Verify Binary

Check the binary works:

```bash
/path/to/tekmetric-mcp version
```

Should output version information.

## Common MCP Client Config Locations

### Claude Desktop

**Mac**: `~/Library/Application Support/Claude/claude_desktop_config.json`
**Windows**: `%APPDATA%\Claude\claude_desktop_config.json`

### Claude Code

**Mac/Linux**: `~/.config/claude-code/mcp_settings.json`
**Windows**: `%APPDATA%\claude-code\mcp_settings.json`

### Cline

**Mac**: `~/Library/Application Support/Code/User/globalStorage/saoudrizwan.claude-dev/settings/cline_mcp_settings.json`

**Windows**: `%APPDATA%\Code\User\globalStorage\saoudrizwan.claude-dev\settings\cline_mcp_settings.json`

## Troubleshooting

### Server Won't Start

**Check permissions:**
```bash
chmod +x /path/to/tekmetric-mcp
```

**Test manually:**
```bash
export TEKMETRIC_CLIENT_ID="your_id"
export TEKMETRIC_CLIENT_SECRET="your_secret"
/path/to/tekmetric-mcp serve
```

Look for errors in the output.

### Authentication Fails

**Verify credentials:**
```bash
curl -X POST https://shop.tekmetric.com/api/v1/oauth/token \
  -u "CLIENT_ID:CLIENT_SECRET" \
  -d "grant_type=client_credentials"
```

Should return an access token.

### Wrong Architecture

**Check your system:**
```bash
# Mac
uname -m  # arm64 or x86_64

# Linux
uname -m  # x86_64 or aarch64
```

Download the matching binary.

### Path Issues

**Use absolute paths only:**
- ✅ `/Users/name/bin/tekmetric-mcp`
- ❌ `~/bin/tekmetric-mcp`
- ❌ `./tekmetric-mcp`

## Building from Source

If you want to build from source instead:

```bash
# Clone the repository
git clone https://github.com/beetlebugorg/tekmetric-mcp.git
cd tekmetric-mcp

# Build
go build -o tekmetric-mcp

# Test
./tekmetric-mcp version
```

Requires Go 1.23 or later.

## Environment File Alternative

Instead of JSON config, you can use a `.env` file:

```bash
# .env
TEKMETRIC_CLIENT_ID=your_client_id
TEKMETRIC_CLIENT_SECRET=your_client_secret
TEKMETRIC_BASE_URL=https://shop.tekmetric.com
TEKMETRIC_DEFAULT_SHOP_ID=123
```

Then configure your MCP client to run:

```json
{
  "command": "/path/to/tekmetric-mcp",
  "args": ["serve"],
  "env": {}
}
```

The server will read from the environment automatically.

## Advanced: Config File

Create `~/.config/tekmetric-mcp/config.json`:

```json
{
  "tekmetric": {
    "base_url": "https://shop.tekmetric.com",
    "client_id": "your_id",
    "client_secret": "your_secret",
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

Then your MCP client config only needs:

```json
{
  "command": "/path/to/tekmetric-mcp",
  "args": ["serve"]
}
```

## Debugging

Enable debug mode to see detailed logs:

```json
{
  "env": {
    "TEKMETRIC_DEBUG": "true"
  }
}
```

Or via command line:

```bash
/path/to/tekmetric-mcp -d serve
```

## Next Steps

✅ **Installation complete!** Now you can:

- Try [Usage Examples](../examples/index.md)
- Learn about [Available Tools](../tools/index.md)
- Start querying your shop data!

## Support

- [GitHub Issues](https://github.com/beetlebugorg/tekmetric-mcp/issues)
- [MCP Protocol Docs](https://modelcontextprotocol.io)
- [Tekmetric API](https://shop.tekmetric.com)
