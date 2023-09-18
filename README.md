# Sora Server SDK for Go

[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/shiguredo/sora-server-sdk-go.svg)](https://github.com/shiguredo/sora-server-sdk-go)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

## About Shiguredo's open source software

We will not respond to PRs or issues that have not been discussed on Discord. Also, Discord is only available in Japanese.

Please read https://github.com/shiguredo/oss/blob/master/README.en.md before use.

## 時雨堂のオープンソースソフトウェアについて

利用前に https://github.com/shiguredo/oss をお読みください。

## 概要

**これは現在開発中です**

WebRTC SFU Sora や Sora Cloud の API やウェブフックを利用するサーバーアプリケーションを開発するための Go 向けの SDK です。

## 対応内容

- [ ] API
- [x] AuthWebhook
- [x] SessionWebhook
- [x] EventWebhook

### Sora Cloud

- [ ] API
- [ ] Webhook

## 利用方法

v2023.2 は対応する Sora のバージョンを表しています。

```
go get -u github.com/shiguredo/sora-server-sdk-go/v2023.2
```

## ライセンス

Apache License 2.0

```
Copyright 2022-2023, Hiroshi Yoshida （Original Author)
Copyright 2022-2023, Shiguredo Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
