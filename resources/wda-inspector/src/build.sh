#!/usr/bin/env bash

cd "$( dirname "${BASH_SOURCE[0]}" )"

npm update\
&& node_modules/browserify/bin/cmd.js -t [ stringify ] main.js -o ../../static/wda-inspector/main.js \
&& node_modules/postcss-cli/bin/postcss -u postcss-import main.css > ../../static/wda-inspector/main.css

echo 'done'