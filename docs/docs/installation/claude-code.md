---
sidebar_position: 2
---

# Claude Code Installation

**For developers using VS Code with Claude Code.** This guide shows you how to configure the Tekmetric assistant in your settings file.

## What You'll Need

- [VS Code](https://code.visualstudio.com/) installed
- [Claude Code extension](https://marketplace.visualstudio.com/items?itemName=Anthropic.claude-code) installed in VS Code
- Your Tekmetric API credentials ([see here](../installation/index.md#tekmetric-api-credentials))
- The tekmetric-mcp binary (see download step below)

## Step 1: Download the Binary

1. Go to [Releases](https://github.com/beetlebugorg/tekmetric-mcp/releases)
2. Download the right version for your system:
   - **Mac (M1/M2/M3)**: `tekmetric-mcp-darwin-arm64`
   - **Mac (Intel)**: `tekmetric-mcp-darwin-amd64`
   - **Windows**: `tekmetric-mcp-windows-amd64.exe`
   - **Linux**: `tekmetric-mcp-linux-amd64`
3. Save it somewhere permanent (like `~/bin/` on Mac/Linux or `C:\Program Files\` on Windows)
4. Rename it to just `tekmetric-mcp` (or `tekmetric-mcp.exe` on Windows)

### Make it Executable (Mac/Linux only)

```bash
chmod +x ~/bin/tekmetric-mcp
```

## Step 2: Configure Claude Code

### Open Claude Code Settings

1. Open VS Code
2. Press `Ctrl+Shift+P` (Windows/Linux) or `Cmd+Shift+P` (Mac)
3. Type "Claude Code: Open MCP Settings"
4. Press Enter

This will open your `claude_code_config.json` file.

### Add the Tekmetric Server

Add this configuration to your settings file:

```json
{
  "mcpServers": {
    "tekmetric": {
      "command": "/absolute/path/to/tekmetric-mcp",
      "args": ["serve"],
      "env": {
        "TEKMETRIC_CLIENT_ID": "your_client_id_here",
        "TEKMETRIC_CLIENT_SECRET": "your_client_secret_here",
        "TEKMETRIC_BASE_URL": "https://api.tekmetric.com",
        "TEKMETRIC_DEFAULT_SHOP_ID": "123"
      }
    }
  }
}
```

### Update the Configuration

**Replace these values:**

1. **Line 4**: Change `/absolute/path/to/tekmetric-mcp` to where YOU saved the file
   - Mac example: `/Users/yourname/bin/tekmetric-mcp`
   - Windows example: `C:\\Program Files\\tekmetric-mcp.exe`
   - ⚠️ **Must be an absolute path** (not `~` or relative)

2. **Line 7**: Paste your Client ID from Tekmetric

3. **Line 8**: Paste your Client Secret from Tekmetric

4. **Line 9**: Set your base URL
   - Production: `https://api.tekmetric.com`
   - Testing: `https://sandbox.tekmetric.com`

5. **Line 10**: Put your shop ID number (or `2`/`3` for sandbox)

### Save the File

Press `Ctrl+S` (Windows/Linux) or `Cmd+S` (Mac) to save.

## Step 3: Restart Claude Code

1. Press `Ctrl+Shift+P` / `Cmd+Shift+P`
2. Type "Developer: Reload Window"
3. Press Enter

Claude Code will restart and connect to the Tekmetric server.

## Step 4: Verify It's Working

Look for the **MCP connection indicator** in the Claude Code status bar (usually bottom-right).

Try asking Claude Code:
```
Show me my Tekmetric shops
```

If you see your shop information, **you're all set!** 🎉

## Configuration Reference

### Full Configuration Example

```json
{
  "mcpServers": {
    "tekmetric": {
      "command": "/Users/yourname/bin/tekmetric-mcp",
      "args": ["serve"],
      "env": {
        "TEKMETRIC_CLIENT_ID": "abc123def456",
        "TEKMETRIC_CLIENT_SECRET": "secretxyz789",
        "TEKMETRIC_BASE_URL": "https://api.tekmetric.com",
        "TEKMETRIC_DEFAULT_SHOP_ID": "123",
        "TEKMETRIC_DEBUG": "false"
      }
    }
  }
}
```

### Optional Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `TEKMETRIC_DEBUG` | Enable debug logging | `false` |
| `TEKMETRIC_TIMEOUT_SECONDS` | Request timeout | `30` |
| `TEKMETRIC_MAX_RETRIES` | Max retry attempts | `3` |

## Troubleshooting

### ❌ "Command not found" or "Cannot find module"

**Problem**: Claude Code can't find the tekmetric-mcp binary

**Solutions**:
- Make sure you're using an **absolute path** (not `~` or relative)
- On Mac, right-click the file → Get Info → copy full path
- On Windows, hold Shift → right-click → "Copy as path"
- Make sure the file is executable (Mac/Linux): `chmod +x /path/to/tekmetric-mcp`

### ❌ "Authentication failed"

**Problem**: Your credentials aren't working

**Solutions**:
- Check for extra spaces in your Client ID or Secret
- Make sure the credentials are for the correct Tekmetric account
- Try regenerating credentials in Tekmetric

### ❌ "Permission denied" (Mac/Linux)

**Problem**: File isn't executable

**Solution**:
```bash
chmod +x /path/to/tekmetric-mcp
```

### ❌ Server not connecting

**Problem**: Configuration issue

**Solutions**:
1. Check JSON syntax (commas, quotes, brackets)
2. Validate JSON: Paste into [jsonlint.com](https://jsonlint.com)
3. Check Claude Code logs:
   - Press `Ctrl+Shift+P` / `Cmd+Shift+P`
   - Type "Developer: Show Logs"
   - Look for errors related to "tekmetric"

### ❌ "Shop not found"

**Problem**: Wrong shop ID

**Solutions**:
- Check your shop ID in Tekmetric Settings
- For sandbox, use `2` or `3`
- Ask Claude "Show me all my shops" to see available IDs

## Finding Your Settings File

If you need to manually locate the settings file:

### Mac
```
~/.config/claude-code/mcp_settings.json
```

### Windows
```
%APPDATA%\claude-code\mcp_settings.json
```

### Linux
```
~/.config/claude-code/mcp_settings.json
```

## Updating to a New Version

When a new version is released:

1. Download the new binary
2. Replace the old one in the same location
3. Restart Claude Code (Developer: Reload Window)
4. Your configuration stays the same

## Next Steps

✅ **Installation complete!** Now you can:

- Try [Usage Examples](../examples/index.md) to see what questions to ask
- Learn about [Available Tools](../tools/index.md)
- Start asking about your shop data right in VS Code!

## Security Notes

- Your credentials are stored locally in VS Code settings
- They're only sent to Tekmetric's API (never anywhere else)
- Consider using VS Code's secure storage for sensitive values
- Never commit your `mcp_settings.json` to version control
