Dev-Guide Overview
===

## Build Your Own Release 

It's highly recommended that build your own Zai release.

A storage system should be closed to your business model. Zai is just a core of
an Object Storage System, it won't be hard to make a change.

[Let's do it!](build_your_distribution.md)

## Codes Layout

Major Parts:

- api: Protocol definition files.
- cmd: Main applications for this project.

## Start with Settings

Settings of Zai are in [settings](https://github.com/zaibyte/pkg/blob/master/config/settings/settings.go).

## Optimization

How could I make Zai suit for my business?

See [Optimization](optimization.md)
