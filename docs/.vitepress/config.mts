import { defineConfig } from 'vitepress'

// https://vitepress.dev/reference/site-config
export default defineConfig({
  title: "Chibi CLI Docs",
  description: "Docs and Guides for Chibi for AniList",
  themeConfig: {
    // https://vitepress.dev/reference/default-theme-config
    nav: [
      { text: 'Home', link: '/' },
      { text: 'Getting Started', link: '/00_getting_started/index' },
      { text: 'CLI Guide', link: '/02_cli_guide/index' }
    ],

    sidebar: [
      {
        text: "Getting Started",
        items: [
          { text: "Pre Requisites", link: "/00_getting_started/prereq" },
        ]
      },
      {
        text: "Installing Chibi",
        items: [
          { text: "Linux", link: "/01_installation/linux" },
          { text: "Windows", link: "/01_installation/windows" },
          { text: "Mac OS", link: "/01_installation/macos" }
        ]
      },
      {
        text: "CLI Guide",
        items: [
          { text: "help", link: "/02_cli_guide/help" }
        ]
      }
    ],

    footer: {
      message: 'Released under GNU GPL-3.0 license',
      copyright: 'Copyright © 2024-present Cosmic Predator'
    },

    search: {
      provider: 'local'
    },

    socialLinks: [
      { icon: 'github', link: 'https://github.com/CosmicPredator/chibi-cli' }
    ]
  }
})
