# Tekmetric MCP Server Documentation

This website is built using [Docusaurus](https://docusaurus.io/), a modern static website generator.

## Installation

```bash
npm install
```

Or using yarn:

```bash
yarn install
```

## Local Development

```bash
npm start
```

Or:

```bash
yarn start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

## Build

```bash
npm run build
```

Or:

```bash
yarn build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

## Deployment

### GitHub Pages

```bash
GIT_USER=<Your GitHub username> npm run deploy
```

Or:

```bash
GIT_USER=<Your GitHub username> yarn deploy
```

If you are using GitHub pages for hosting, this command is a convenient way to build the website and push to the `gh-pages` branch.

## Project Structure

```
docs/
├── docs/                   # Documentation pages
│   ├── Introduction.md    # Home page
│   ├── installation.md    # Installation guide
│   ├── tools/            # Tool documentation
│   ├── configuration/    # Configuration docs
│   └── examples/         # Usage examples
├── src/                   # React components and custom pages
│   └── css/              # Custom CSS
├── static/               # Static assets (images, etc)
│   └── img/
├── docusaurus.config.js  # Site configuration
├── sidebars.js           # Sidebar navigation
└── package.json          # Dependencies
```

## Customization

### Adding Pages

Add new markdown files to `docs/` directory. They will automatically appear in the sidebar.

### Modifying Navigation

Edit `sidebars.js` to customize sidebar structure.

### Styling

Edit `src/css/custom.css` for custom styles.

### Configuration

Edit `docusaurus.config.js` for site-wide configuration.
