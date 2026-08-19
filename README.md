# yell

> [!IMPORTANT]
> `yell` is experimental and has only been tested on a modern MacOS & Windows system. Please report issues.

## Why?

- More control over alert customization. Notificiation center for each OS is limiting.

## Why not?

- You like everything organised. `yell` is an unorthodox alert app.

## Install

```
TODO
```

## Usage

```sh
yell --title="Fuck off" --message="Time to fuck off boi" --icon="🖕🏽"
```

> <img width="379" height="148" alt="Screenshot 2026-08-19 at 6 52 16 PM" src="https://github.com/user-attachments/assets/182656f9-422a-4815-bdb6-0d75982b863b" />



You can also pipe the `message` from other UNIX utils

```sh
echo "$(lifespan yyyy-mm-dd) remaining" | yell --title="Break lele bhai" --icon="🍀"
```

> <img width="498" height="148" alt="Screenshot 2026-08-19 at 6 51 07 PM" src="https://github.com/user-attachments/assets/f56544b4-d592-40fd-8048-a4925b6f09a2" />
