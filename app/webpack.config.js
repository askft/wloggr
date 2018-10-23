const HtmlWebpackPlugin = require("html-webpack-plugin");

const htmlPlugin = new HtmlWebpackPlugin({
  template: "./src/index.html",
  filename: "./index.html"
});

module.exports = {
  entry: {
    app: "./src/index.js"
  },
  output: {
    path: __dirname + "/../dist",
    publicPath: "/",
    filename: "js/[name].js?v=[chunkhash]"
  },
  devServer: {
    contentBase: "./dev",
    historyApiFallback: true
  },
  plugins: [htmlPlugin],
  module: {
    rules: [
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: ["babel-loader"]
      },
      {
        test: /\.css$/,
        use: ["style-loader", "css-loader"]
      }
    ]
  },
  resolve: {
    extensions: ["*", ".js", ".jsx", ".css"]
  }
};
