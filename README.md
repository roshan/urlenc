# UrlEnc

Trivial command-line so that you can make your URLs and GET params appropriately encoded at the CLI.

## Usage

```
> urlenc Hello World
Hello%20World
```

for things like

```
curl -L "google.com?q=$(./urlenc what\'s my name again)"
```

Use the binary renamed to `urldec` to have the opposite effect
