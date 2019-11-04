# Meereen

[![Go Report Card](https://goreportcard.com/badge/github.com/aerialls/meereen)](https://goreportcard.com/report/github.com/aerialls/meereen)
[![GitHub Actions](https://github.com/aerialls/meereen/workflows/Test/badge.svg)](https://github.com/aerialls/meereen/actions?query=workflow%3ATest)

Meereen is a lightweight monitoring tool with the following constraints to keep it that way.

* No database
* No web interface
* Easy to deploy and to manage checks

Binaries for Linux, macOS and Windows are available from the releases page.

## Usage

Meereen uses checks, processors and notifiers. A check uses a processor to verify something and uses a notifier to notify th end-user.

```
meereen --config /etc/meereen/config.yml
```

The default config file can be this one.

```yaml
checks: /etc/meereen/checks.d
notifiers:
  - name: mynotifier
    kind: telegram
    data:
      token: ...
      chat_id: ...
  - ...
```

Checks will be loaded from the folder you specified. You can have multiple checks in one file.

```yaml
checks:
  - title: Google HTTP
    notifier: mynotifier
    processor:
      kind: http
      data:
        url: https://www.google.com/
  - title: Google DNS
    notifier: mynotifier
    processor:
      kind: dns
      data:
        domain: google.com
```

## Notifiers

The following notifiers are supported.

- Telegram (`telegram`)

## Processors

The following processors are supported.

- HTTP (`http`)
- DNS (`dns`)

## Roadmap

Here are the following items I would like to add in the near future.

* Dedicated scheduling for a check
* Retry system
* Reload when changes on checks.d
* More notifiers
  * Slack
  * E-Mail
  * Webhook
  * Mattermost
  * ...
* More processors
  * Certificates
  * Shell
  * ...
