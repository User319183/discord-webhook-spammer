# Discord Webhook Spammer

This project is a Discord webhook spammer written in Go. It uses a list of proxies to send messages to a specified webhook, with added functionality for Text-to-Speech (TTS) messages and rate limit handling.

## Features

- User input for webhook URL, spam message, number of threads, sleep duration, and whether to use TTS (NEW).
- Uses a list of proxies to send messages.
- Retries sending messages if an error occurs. (NEW)
- Handles rate limits by sleeping for the duration specified in the `Retry-After` header. (NEW)
- Prints colored console output for better user experience.

## Usage

1. Clone the repository.
2. Make sure you have Go installed on your machine.
3. Run the program using `go run nuker.go`.
4. Follow the prompts to input the webhook URL, spam message, number of threads, sleep duration, and whether to use TTS.

## Configuration

The program reads a list of proxies from a file named "proxies.txt". Each proxy should be on a new line. If you don't want to use proxies, you can modify the `spam` function to remove the proxy configuration.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This tool is for educational purposes only. The developer is not responsible for any misuse of this tool. Misusing the Discord API can result in your account being banned, so please use this tool responsibly.
