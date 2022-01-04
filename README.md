# Bluepill

Have an app that's letting you down? Keep it up!

## What is this?

Do you have an app that needs to run continuously in the terminal or in the
background, but crashes once in a while? Bluepill can help you keep it up and
running!

You just need to run `bluepill your_app_executable arg1 arg2 etc` and it will
relaunch the app with the same arguments every time it crashes.

## Features

- Hooks up stdin/stdout/stderr so that your app can still receive keyboard input
  and print to the terminal.
- Restarts your app every time it exits in error, but will exit if the app
  finishes successfully.
- Will honor Ctrl+C and SIGTERM, although it won't pass these on to the app (see
  the TODO section).

## Install

Just download the latest release from the
[releases](https://github.com/tonyfg/bluepill/releases) page, put the executable
in your $PATH and start using it.

## Build

`go build`

## TODO

- Pass received signals to the app. This should help with graceful handling of
  Ctrl+C, or system shutdown (as well as apps that make use of other signals).
- Multi-arch builds
- Command line arguments to:
  - Allow restarting apps that exit without any error
  - Disabling grace period
  - Setting a custom grace period

## FAQ

### What's with the name?
I feel pretty cozy here in the matrix even if it's fake.

### Do you really need a FAQ for such a small piece of software?
No, but I like FAQs.

### Are there any useful questions or answers in the FAQ?
None at all.
