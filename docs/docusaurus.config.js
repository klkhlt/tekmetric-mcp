// @ts-check
import {themes as prismThemes} from 'prism-react-renderer';

/** @type {import('@docusaurus/types').Config} */
const config = {
  title: 'Tekmetric MCP Server',
  tagline: '🔧 Model Context Protocol server for Tekmetric shop management API',
  favicon: 'img/favicon.ico',
  url: 'https://beetlebugorg.github.io/',
  baseUrl: '/tekmetric-mcp/',
  organizationName: 'beetlebugorg',
  projectName: 'tekmetric-mcp',
  onBrokenLinks: 'throw',
  onBrokenMarkdownLinks: 'warn',
  i18n: {
    defaultLocale: 'en',
    locales: ['en'],
  },
  presets: [
    [
      'classic',
      ({
        docs: {
          routeBasePath: '/',
          sidebarPath: './sidebars.js',
          editUrl:
            'https://github.com/beetlebugorg/tekmetric-mcp/tree/main/docs/',
        },
        blog: false,
        theme: {
          customCss: './src/css/custom.css',
        },
      }),
    ],
  ],
  themeConfig: ({
    image: 'img/tekmetric-mcp.jpg',
    navbar: {
      title: '',
      logo: {
        alt: 'Tekmetric MCP Logo',
        src: 'img/tekmetric-mcp-logo.png',
      },
      items: [
        {
          href: 'https://github.com/beetlebugorg/tekmetric-mcp',
          position: 'right',
          className: 'header-github-link',
          'aria-label': 'GitHub repository',
        },
      ],
    },
    footer: {
      style: 'dark',
      links: [],
      copyright: `Copyright © ${new Date().getFullYear()} Jeremy Collins. Not affiliated with Tekmetric, Inc.`,
    },
    prism: {
      theme: prismThemes.github,
      darkTheme: prismThemes.dracula,
      additionalLanguages: ['bash', 'json', 'go'],
    },
  }),
};

export default config;
