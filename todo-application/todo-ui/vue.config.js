const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
});

module.exports = {
  devServer: {
    proxy: {
      '^/todo': {
        target: 'http://localhost:3000',
        ws: true,
        changeOrigin: true
      },
    }
  }
}