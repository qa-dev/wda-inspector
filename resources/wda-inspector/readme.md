## Сборка

### Установка зависимостей

```
cd resources/wda-inspector/src/
npm install
```

### Сборка javascript

```
node_modules/browserify/bin/cmd.js -t [ stringify ] main.js -o ../../static/wda-inspector/main.js
```

### Сборка CSS

```
node_modules/postcss-cli/bin/postcss -u postcss-import main.css > ../../static/wda-inspector/main.css
```