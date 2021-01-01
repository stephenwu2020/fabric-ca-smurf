const proxy = import('http-proxy-middleware')

module.exports = {
  devServer: {
    host: '127.0.0.1',
    port: 8080,
    proxy: {
      '/api': {
        target: 'http://127.0.0.1:7054',
        changeOrigin: true,
        pathRewrite: {
          '^/api': '/api'
        }
      }
    }
  },
  pluginOptions: {
    autoRouting: {
      chunkNamePrefix: "page-"
    },
    i18n: {
      locale: "en",
      fallbackLocale: "en",
      localeDir: "locales",
      enableInSFC: false
    }
  }
};
