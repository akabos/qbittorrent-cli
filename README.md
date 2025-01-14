# qbittorrent-cli

A cli to manage qBittorrent.

## Install

Download binary and put somewhere in $PATH.

Extract binary

    tar -xzvf qbt.tar.gz

Move to somewhere in `$PATH`. Need sudo if not already root. Or put it in your user `$HOME/bin` or similar.

    sudo mv qbt /usr/bin/

Verify that it runs

    qbt

This should print out basic usage.

## Configuration

Create a new configuration file `.qbt.toml` in `$HOME/.config/qbt/`.

    mkdir -p ~/.config/qbt && touch ~/.config/qbt/.qbt.toml

A bare minimum config.

```toml
[qbittorrent]
host     = "127.0.0.1" # qbittorrent webui-api hostname/ip
port     = 6776        # qbittorrent webui-api port
login    = "user"      # qbittorrent webui-api user
password = "password"  # qbittorrent webui-api password
```

### rutorrent-autodl-irssi setup

In rutrorrent, go to autodl-irssi `Preferences`, and then the `Action` tab. Put in the following for the global action. This can be set in a specific filter as well.

```
Choose .torrent action: Run Program
Command: /usr/bin/qbt
Arguments: add "$(TorrentPathName)"
```

## Usage

Use `qbt help` to find out more about how to use.

Commands:
  - add
  - list
  - version
  - help

Global flags:
  * `--config` -override config file
  
Use "qbt [command] --help" for more information about a command.

### Add

Add a new torrent to qBittorrent.

    qbt add my-torrent-file.torrent
