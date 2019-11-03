# Meereen

[![Go Report Card](https://goreportcard.com/badge/github.com/aerialls/meereen)](https://goreportcard.com/report/github.com/aerialls/meereen)
[![GitHub Actions](https://github.com/aerialls/meereen/workflows/Test/badge.svg)](https://github.com/aerialls/meereen/actions?query=workflow%3ATest)

Meereen is a lightweight monitoring tool with the following constraints to keep it that way.

* No database
* No web interface
* Easy to deploy
* Easy to automate checks

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
