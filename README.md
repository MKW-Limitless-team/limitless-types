# Publishing Go Module

## 1. Check current version

`git fetch --tags`

`git tag`

## 2. Publish module

`go mod tidy`

`git tag [version]`

`git push origin [version]`

`go list -m github.com/MKW-Limitless-team/limitless-types@[version]`

## 3. Importing

**This should be done in what ever go project you want to import from**

`go get github.com/MKW-Limitless-team/limitless-types@[version]`
