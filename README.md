# Scooplint

## What is this

This is a tiny tool that lints a minimal scoop manifest. This tool is rather
unflexible and minimal. I am currently using it only to check the manifest of
[cordless](https://github.com/Bios-Marcel/cordless).

## Usage

Usage is fairly easy:

```shell
scooplint -urlpattern=https://target.url/manifest.json -versionpattern=\\d{4}-\\d{2}-\\d{2} manifest.json
```

Defining the `urlpattern` and the `versionpattern` is optional. By default
both values are `.+`.