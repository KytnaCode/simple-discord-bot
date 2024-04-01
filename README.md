# simple-discord-bot

A simple and small discord bot made in Go for learning purposes.

## 1.Overview

`simple discord bot` is built using Go standard library with one exception: the godotenv package is used to read bot's credentials fron a `.env` file. This bot verifies Discord interaction requests, handles ping responses and processes command responses. Additionally  there's a cli to register bot's commands from json files.

## 2.Commands

The bot has only one command `/ping`, when you use it, the bot replies with `pong!`.

## 3.Usage

### 3.1. Requirements

To run the bot, you'll need the following:

* A registered discord app (you can create one in Discord's developer portal)
* Docker and Docker Compose
* Go version 1.22.1 or higher
* `make` utily

If you want to run the bot in a local environment you also will need:

* `ngrok` (or an alrernative method to allow Discord API requests to reach your computer)

### 3.2 Configuration

Create a file named `.env` and add the following content:

```sh
export BOT_PUBLIC_KEY="<YOUR_BOT_PUBLIC_KEY>"
export BOT_TOKEN="<YOUR_BOT_TOKEN>"
export APP_ID="<YOUR_APP_ID>"
```

And replace `<YOUR_BOT_PUBLIC_KEY>`, `<YOUR_BOT_TOKEN>` and `<YOUR_APP_ID>` with your actual credentials, you can find them on your Discord app page in Discord's developers portal.

### 3.3 Register Commands

Before running the bot, ensure that the commands are registered, to do so, execute:

```sh
make register
```

### 3.4. Run the bot

Finally, start the bot with:

```sh
make run-dev
```

*Make sure docker daemon is running*

## 4. The Unlicense

This work is released under The UNLICENSE, also known as public domain. Feel free to use, copy, redistribute, or modify this code without any restrictions or permissions.
