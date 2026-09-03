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

Find your theme from `yell list themes`

```sh
yell --title="Lunch Time" --message="Time to eat boi" --icon="🥗" --theme=warm
```

> <img width="391" height="148" alt="Screenshot 2026-08-19 at 11 52 29 PM" src="https://github.com/user-attachments/assets/40fe0eb9-bb33-4442-a21c-a416a83acbf8" />

Markdown Support on `message`

```sh
yell --title="Deploy Complete" --message=$'**Build #142** finished in `2m 34s`\n\n- ✅ Tests passed\n- ✅ Lint clean\n- ⚠️ *1 warning* in `auth.go`\n\nRun `git log -1` for details.' --icon="🚀" --sound=bubble
```

> <img width="456" height="229" alt="Screenshot 2026-08-26 at 12 53 44 AM" src="https://github.com/user-attachments/assets/4837e556-e10c-4229-aaf2-8a26535ebd2b" />

Markdown also supports JPEG images, featuring an OG Sketch by yours truly

```sh
echo 'Go to bed ![horse-dog-snuggies.jpg](https://lh3.googleusercontent.com/d/1U5UcySisWg_PoUbM5PeTGZaQaL21tllW)' | yell --title="Sleepyy Time!" --icon="💤" --theme=white
```

> <img width="460" height="360" alt="Screenshot 2026-09-03 at 8 22 41 PM" src="https://github.com/user-attachments/assets/ff32c0c1-6d0f-4f62-a489-05520eacbc28" />



Play a background sound, find yours via `yell list sounds`

```sh
yell --title="Lunch Time" --message="Time to eat boi" --icon="🥗" --theme=warm --sound=chime
```

