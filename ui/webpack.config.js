var path = require('path');

module.exports = {
  context: __dirname + "/app",
  entry: path.resolve(__dirname, "app/Main.js"),

  output: {
    path: path.resolve(__dirname, '..', "public"),
    filename: "app.js",
  },
  module: {
    loaders: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        loader: "babel",
        query:
          {
            presets:["es2015","react"]
          }
      }
    ]
  },
};
