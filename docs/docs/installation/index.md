---
sidebar_position: 2
title: Installation
---

# Installation

Choose your AI assistant below for specific installation instructions. **No technical experience required!**

## Prerequisites

Before you start, you'll need:

### Tekmetric API Credentials

These are like a special password that lets your AI assistant access your shop data.

**Get your credentials:**

1. Log into your Tekmetric account
2. Click **Settings** in the menu
3. Click **API Access**
4. Click **Create New Application**
5. Give it a name like "AI Assistant"
6. Copy the **Client ID** and **Client Secret** (save these somewhere safe!)

❓ **Can't find API Access?** Ask your Tekmetric account manager - they can enable it for you.

## Choose Your Installation Method

Click on your AI assistant below:

### [Claude Desktop Installation](./claude-desktop.md)
**Recommended for most users**
- One-click Desktop Extension (.mcpb) installer
- Works on Mac and Windows
- Easiest setup

### [Claude Code Installation](./claude-code.md)
**For VS Code users**
- Configure via settings file
- Works within VS Code
- Great for developers

### [Cline Installation](./cline.md)
**For VS Code users**
- Popular VS Code extension
- MCP support built-in
- Good for technical users

### [Manual Installation](./manual.md)
**For advanced users or other MCP clients**
- Works with any MCP-compatible client
- Requires manual configuration
- Most flexible option

## Testing vs. Real Data

When you first set up, we recommend testing with fake data before using real customer information.

### Test Mode (Sandbox)

Use Tekmetric's test environment:
- **Base URL**: `https://sandbox.tekmetric.com`
- **Shop ID**: Usually `2` or `3`
- Get sandbox credentials from your Tekmetric rep

### Production Mode (Real Data)

Use your actual shop data:
- **Base URL**: `https://api.tekmetric.com`
- **Shop ID**: Your real shop ID number
- Use your production API credentials

**Always test first!** Make sure everything works with test data before switching to real data.

## Need Help?

- Check the troubleshooting section in your specific installation guide
- [Create a GitHub issue](https://github.com/beetlebugorg/tekmetric-mcp/issues)
- Contact Tekmetric support for help with API credentials

## Next Steps

After installation:
- Check out [Usage Examples](../examples/index.md) to see what you can ask
- Read about [Available Tools](../tools/index.md)
- Start asking questions about your shop!

## Privacy & Security

**Your data is safe:**
- ✅ This tool can only *read* your data, never change it
- ✅ All connections are encrypted
- ✅ Your credentials stay on your computer
- ✅ Your AI assistant doesn't store your shop data

**Think of it like this**: You're giving your AI assistant permission to look at your Tekmetric account, just like you'd let an employee log in and look around. They can see information but can't modify anything.
