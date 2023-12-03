# Superior Discord Webhook Spammer

Welcome to the fastest and most efficient Discord Webhook Spammer, meticulously crafted in Go. This tool is designed with superior performance and user experience in mind, providing a seamless and fluent interaction.

## Features

- **Fast and Efficient**: This tool is designed to be the fastest Discord Webhook Spammer available.
- **Customizable**: Personalize your spam messages with custom usernames and profile pictures. Make your messages stand out or blend in as you desire.
- **Text-to-Speech (TTS)**: Annoy or amuse with the TTS feature. Send your messages not just to the eyes, but also to the ears of the recipients.
- **Proxy Support**: Utilizes a list of HTTP/HTTPS proxies to send messages, ensuring efficient use of resources and improved performance.
- **User-Friendly**: Interactive prompts for user inputs such as webhook URL, spam message, number of threads, sleep duration, and Text-to-Speech (TTS) option.
- **Error Handling**: Robust error handling and retry mechanism for message sending.
- **Rate Limit Handling**: Intelligent handling of rate limits by sleeping for the duration specified in the `Retry-After` header.
- **Colored Console Output**: Enhanced user experience with colored console output for easy tracking of the spamming process.

## Demo

Watch this [video](https://streamable.com/m3hsbj) to see the tool in action.

## Usage

1. Clone the repository.
2. Ensure you have Go installed on your machine.
3. Populate the "proxies.txt" file with your list of HTTP/HTTPS proxies. Each proxy should be on a new line.
4. Run the program using `go run nuker.go`.
5. Follow the prompts to input the webhook URL, spam message, number of threads, sleep duration, and whether to use TTS. You can also customize the username and avatar URL for the messages.

## Contributing

We welcome contributions that can improve this tool. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This tool is intended for educational purposes only. Misuse can result in your account being banned from Discord. We are not responsible for any misuse of this tool. Always use responsibly and ensure you comply with Discord's API usage policies.
