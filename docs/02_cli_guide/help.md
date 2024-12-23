---
title: Help
lang: en-US
---

# The `help` Command
You can use `help` command to know about the usage of a specific command. It shows a command's available args, flags and a small info about what that command does. For example:

```shell
$ chibi --help
# Chibi for AniList - A lightweight anime & manga tracker CLI app powered by AniList.

# Usage:
#   chibi [flags]
#   chibi [command]

# Available Commands:
#   add         Add a media to your list
#   help        Help about any command
#   login       Login with anilist
#   ls          List your current anime/manga list
#   profile     Get's your AniList profile (requires login)
#   search      Search for anime and manga
#   update      Update a list entry

# Flags:
#   -h, --help   help for chibi

# Use "chibi [command] --help" for more information about a command.
```

::: info NOTE
The above help command text varies from your output as we may add additional commands and flags in the future.
:::

Alternatively, you can also use the **shorthand** syntax to call the same command:
```shell
$ chibi -h
```

## Command specific help
The help command can be used not only on the top level, but also to individual commands. In other words, you can get the help text of a specific command using the syntax `chibi <command> --help`. For example;
```shell
$ chibi profile --help
# Get's your AniList profile (requires login)

# Usage:
#   chibi profile [flags]

# Flags:
#   -h, --help   help for profile
```

## Shorthand Syntax
The help command can be called on various formats. The following commands all does the same thing, which is printing the help text:

```shell
# All works the same
$ chibi profile --help

$ chibi profile -h

$ chibi help profile
```