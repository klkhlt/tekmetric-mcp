---
sidebar_position: 3
---

# Cline Installation

**For developers using Cline in VS Code.** Cline (formerly Claude Dev) is a popular VS Code extension with built-in MCP support.

## What You'll Need

- [VS Code](https://code.visualstudio.com/) installed
- [Cline extension](https://marketplace.visualstudio.com/items?itemName=saoudrizwan.claude-dev) installed
- Your Tekmetric API credentials ([see here](../installation/index.md#tekmetric-api-credentials))
- The tekmetric-mcp binary (see download step below)

## Step 1: Download the Binary

1. Go to [Releases](https://github.com/beetlebugorg/tekmetric-mcp/releases)
2. Download the right version for your system:
   - **Mac (M1/M2/M3)**: `tekmetric-mcp-darwin-arm64`
   - **Mac (Intel)**: `tekmetric-mcp-darwin-amd64`
   - **Windows**: `tekmetric-mcp-windows-amd64.exe`
   - **Linux**: `tekmetric-mcp-linux-amd64`
3. Save it somewhere permanent (like `~/bin/` on Mac/Linux)
4. Rename it to just `tekmetric-mcp` (or `tekmetric-mcp.exe` on Windows)

### Make it Executable (Mac/Linux only)

```bash
chmod +x ~/bin/tekmetric-mcp
```

## Step 2: Configure Cline

### Open Cline MCP Settings

1. Open VS Code
2. Click the Cline icon in the sidebar (or press `Ctrl+Shift+P` and search for "Cline")
3. Click the **gear icon** ⚙️ in the Cline panel
4. Select **MCP Settings**

This will open Cline's MCP configuration file.

### Add the Tekmetric Server

Add this to your Cline MCP settings:

```json
{
  "mcpServers": {
    "tekmetric": {
      "command": "/absolute/path/to/tekmetric-mcp",
      "args": ["serve"],
      "env": {
        "TEKMETRIC_CLIENT_ID": "your_client_id_here",
        "TEKMETRIC_CLIENT_SECRET": "your_client_secret_here",
        "TEKMETRIC_BASE_URL": "https://shop.tekmetric.com",
        "TEKMETRIC_DEFAULT_SHOP_ID": "123"
      }
    }
  }
}
```

### Update the Configuration

**Replace these values:**

1. **Line 4**: Full path to the tekmetric-mcp binary
   - Mac example: `/Users/yourname/bin/tekmetric-mcp`
   - Windows example: `C:\\Program Files\\tekmetric-mcp.exe`
   - ⚠️ **Must be absolute** (not `~`)

2. **Line 7**: Your Tekmetric Client ID

3. **Line 8**: Your Tekmetric Client Secret

4. **Line 9**: Base URL
   - Production: `https://shop.tekmetric.com`
   - Testing: `https://sandbox.tekmetric.com`

5. **Line 10**: Your shop ID (or `2`/`3` for sandbox)

### Save and Reload

1. Save the file (`Ctrl+S` / `Cmd+S`)
2. Close and reopen the Cline panel

## Step 3: Verify It's Working

In the Cline panel, you should see **tekmetric** listed under available MCP servers.

Try asking Cline:
```
Show me my Tekmetric shops
```

If Cline can access your shop data, **you're all set!** 🎉

## Configuration Example

### Full Config with Optional Settings

```json
{
  "mcpServers": {
    "tekmetric": {
      "command": "/Users/yourname/bin/tekmetric-mcp",
      "args": ["serve"],
      "env": {
        "TEKMETRIC_CLIENT_ID": "abc123",
        "TEKMETRIC_CLIENT_SECRET": "secret789",
        "TEKMETRIC_BASE_URL": "https://shop.tekmetric.com",
        "TEKMETRIC_DEFAULT_SHOP_ID": "123",
        "TEKMETRIC_DEBUG": "false",
        "TEKMETRIC_TIMEOUT_SECONDS": "30",
        "TEKMETRIC_MAX_RETRIES": "3"
      }
    }
  }
}
```

## Troubleshooting

### ❌ "Command not found"

**Problem**: Cline can't find the binary

**Solutions**:
- Use absolute path (not `~` or `.`)
- Make sure file exists at that location
- Make sure file is executable: `chmod +x /path/to/tekmetric-mcp`

### ❌ Tekmetric server not showing in Cline

**Problem**: Configuration not loaded

**Solutions**:
1. Check JSON syntax is valid
2. Make sure you saved the file
3. Completely close and reopen Cline panel
4. Try reloading VS Code window

### ❌ "Authentication failed"

**Problem**: Wrong credentials

**Solutions**:
- Remove any extra spaces from Client ID/Secret
- Verify credentials in Tekmetric Settings → API Access
- Try regenerating credentials

### ❌ "Permission denied" (Mac/Linux)

**Problem**: File not executable

**Solution**:
```bash
chmod +x /path/to/tekmetric-mcp
```

### ❌ Server starts but queries fail

**Problem**: Wrong shop ID or base URL

**Solutions**:
- Verify shop ID in Tekmetric Settings
- For sandbox, use `2` or `3`
- Check base URL matches your environment
- Ask "Show me all my shops" to see available IDs

## Finding Your Settings File

Cline stores MCP settings here:

### Mac
```
~/Library/Application Support/Code/User/globalStorage/saoudrizwan.claude-dev/settings/cline_mcp_settings.json
```

### Windows
```
%APPDATA%\Code\User\globalStorage\saoudrizwan.claude-dev\settings\cline_mcp_settings.json
```

### Linux
```
~/.config/Code/User/globalStorage/saoudrizwan.claude-dev/settings/cline_mcp_settings.json
```

## Cline Tips

### Using Tools in Cline

When you ask Cline about Tekmetric data, it will automatically:
1. See the available Tekmetric tools
2. Choose the right tool for your query
3. Execute it and show you results

### Example Queries

```
Find customer John Smith in my shop
```

```
What repair orders are in progress?
```

```
Show me vehicles scheduled for service this week
```

Cline understands natural language and will use the appropriate Tekmetric tools automatically.

## Updating

When a new version is released:

1. Download the new binary
2. Replace the old file
3. Restart Cline (close and reopen the panel)
4. Configuration stays the same

## Next Steps

✅ **Installation complete!** Now you can:

- Try [Usage Examples](../examples/index.md) to see what to ask
- Learn about [Available Tools](../tools/index.md)
- Use Cline to query your shop data while coding!

## Security Notes

- Your API credentials are stored locally in Cline's settings
- Only sent to Tekmetric's API (nowhere else)
- Never commit settings files to version control
- Consider using environment variables for extra security
