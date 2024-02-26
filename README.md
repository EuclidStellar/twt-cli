
# TWT - Twitter CLI

TWT is a command-line tool written in Go that allows you to post tweets to Twitter directly from your terminal.

## Installation

### Pre-built Binaries

Download the appropriate binary for your operating system from the [Releases](https://github.com/EuclidStellar/twitter-cli-golang/releases/tag/V1.0.2) page.

#### Linux

```bash
sudo curl -L -o /usr/local/bin/twt https://github.com/EuclidStellar/twitter-cli-golang/releases/download/V1.0.2/twt_linux
sudo chmod +x /usr/local/bin/twt
```

#### macOS

```bash
sudo curl -L -o /usr/local/bin/twt https://github.com/EuclidStellar/twitter-cli-golang/releases/download/V1.0.2/twt_macos
sudo chmod +x /usr/local/bin/twt
```

#### Windows

1. Download the `twt.exe` binary from the [Releases](https://github.com/EuclidStellar/twitter-cli-golang/releases/tag/V1.0.2) page and add it to your system's PATH.
2. 2. Move the downloaded `twt.exe` file to a directory of your choice. For example, you can create a directory named `C:\twt` and move `twt.exe` there.
3. Right-click on the Start button and select "System".
4. In the System window, click on "Advanced system settings" on the left side.
5. In the System Properties window, click on the "Environment Variables..." button.
6. In the Environment Variables window, under "System variables", select the "Path" variable and click on "Edit...".
7. In the Edit Environment Variable window, click on "New" and enter the path to the directory where you placed `twt.exe` (e.g., `C:\twt`).
8. Click "OK" on all windows to save the changes.
9. Open a new Command Prompt window and you should be able to run `twt` from anywhere.

### From Source

If you have Go installed, you can build `twt` from source:

```bash
go install github.com/EuclidStellar/twitter-cli-golang@latest
```
if path is not in usr/local/bin then use this command 

```bash
sudo mv twt /usr/local/bin
```

## Usage

1. Obtain your Twitter API keys from the [Twitter Developer Portal](https://developer.twitter.com/).
2. Run `twt` in your terminal.
3. Enter your Twitter API keys when prompted.
4. Enter your tweet content when prompted.


# Contributing to TWT - Twitter CLI

Thank you for considering contributing to TWT - Twitter CLI! This document outlines some guidelines to help you contribute to this project effectively.

## Code of Conduct

Before contributing, please read and adhere to our [Code of Conduct](CODE_OF_CONDUCT.md). We aim to foster an inclusive and welcoming community for all.

## How Can I Contribute?

### Reporting Bugs

If you encounter any bugs or issues with TWT, please search the [issue tracker](https://github.com/yourusername/twt/issues) to see if the issue has already been reported. If not, please open a new issue and provide detailed information about the problem, including steps to reproduce it and any relevant error messages.

### Requesting Features

We welcome feature requests! If you have an idea for a new feature or improvement, please search the [issue tracker](https://github.com/yourusername/twt/issues) to see if it has already been requested. If not, feel free to open a new issue and describe the feature you would like to see.

### Contributing Code

1. Fork the repository and create a new branch for your contribution.
2. Make your changes, following the [Go coding style](https://github.com/golang/go/wiki/CodeReviewComments) and ensuring that your code is well-documented.
3. Write tests for your code to ensure its correctness and prevent regressions.
4. Run `go test` to ensure that all tests pass.
5. Commit your changes and create a pull request. Please provide a clear and descriptive title for your pull request, along with a detailed description of the changes you've made.

### Improving Documentation

Improvements to documentation are always welcome! If you notice any errors or areas where the documentation could be clearer, please open an issue or pull request with your suggestions.

## Code Review Process

All contributions will be reviewed by project maintainers. Feedback will be provided, and changes may be requested before your contribution is accepted. We appreciate your patience and understanding throughout the review process.

## License

By contributing to this project, you agree that your contributions will be licensed under the [MIT License](LICENSE).
