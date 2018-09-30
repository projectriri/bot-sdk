# Bot-SDK

The SDK that provides a common bot interface with several implementations.

## Implementations

| IM | Package | Base SDK |
| --- | --- | --- |
| Telegram | [github.com/projectriri/bot-sdk/sdks/telegram-bot-api](https://github.com/projectriri/bot-sdk/tree/master/sdks/telegram-bot-api) | [github.com/go-telegram-bot-api/telegram-bot-api](https://github.com/go-telegram-bot-api/telegram-bot-api) |
| QQ | [github.com/projectriri/bot-sdk/sdks/qq-bot-api](https://github.com/projectriri/bot-sdk/tree/master/sdks/qq-bot-api) | [github.com/catsworld/qq-bot-api](https://github.com/catsworld/qq-bot-api) |

## Usage

All methods are fairly self explanatory. For example use, checkout [example/main.go](https://github.com/projectriri/bot-sdk/tree/master/example/main.go)

## Contribute

You're welcomed to implement a Bot SDK derivation conforming to the interface.
If you would like to merge it into this repository, just open a pull request.
