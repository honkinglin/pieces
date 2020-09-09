var path = require('path');

module.exports = {
  mode: 'development',

  entry: {
    app: path.resolve('./index.js'),
  },

  module: {
    rules: [{
      test: /\.js$/,
      loader: 'babel-loader',
      options: {
        presets: ['@babel/preset-env'],
        plugins: [['@babel/plugin-transform-react-jsx', {
          pragma: 'createElement'
        }]]
      }
    }]
  }
}
