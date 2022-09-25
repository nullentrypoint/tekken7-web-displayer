
# init

```bash
npm init -y
npm install livegollection-client
npm install webpack webpack-cli --save-dev
```


# build client

```bash
npx webpack --mode production --entry ./src/index.js --output-filename bundle.js -o ./static
```