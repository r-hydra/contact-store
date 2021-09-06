const mix = require('laravel-mix');

mix.setPublicPath('static').ts('src/app.tsx', 'static/app.js').
    react().
    less('src/app.less', 'static/app.css').
    sourceMaps(false).
    webpackConfig({
      resolve: {
        extensions: ['.ts', '.tsx', '.less'],
      },
    }).
    version();
