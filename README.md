![Intercom](sample/Intercom_logo-github.png)


[![Go Report Card](https://goreportcard.com/badge/github.com/intercom/speakeasy)](https://goreportcard.com/report/github.com/intercom/speakeasy)
[![Apache License](http://img.shields.io/badge/license-APACHE2-blue.svg?style=flat)](https://www.apache.org/licenses/LICENSE-2.0.html)
[![CircleCI](https://circleci.com/gh/intercom/speakeasy.svg?style=svg)](https://circleci.com/gh/intercom/speakeasy)

# Speakeasy
A tool for creating web servers that you can embed in iOS or Android apps.

## Install

* Install the `speakeasy` tool using the following commands:

    ```bash
    go get -d -u github.com/intercom/speakeasy/...
    $GOPATH/src/github.com/intercom/speakeasy/setup.sh
    ```

### Sample Engine
* From the `sample/backend` directory run the following commands:

    ```bash
    ./encode-assets.sh
    speakeasy build github.com/intercom/speakeasy/sample/backend
    ```

### Sample app
* Build the sample engine
* Open the Android project in `sample/android` and build it in Android studio

### About

This sample uses [Twemoji](https://github.com/twitter/twemoji) images.

### License

Speakeasy is released under the Apache 2.0 license. See LICENSE for details.
