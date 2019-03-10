# Repeat ProtonVPN

> Repeats ProtonVPN connection attempts until successful.

## Why

On some networks, my ProtonVPN connection attempts fail on the first few tries using the [protonvpn-cli](https://github.com/ProtonVPN/protonvpn-cli). It usually works after a random number of attempts. This program just repeats the connection command.

Note: this could be solved with a simple shell script, but I wanted to experiment with Go.

## Prerequisites

- Go
- [protonvpn-cli](https://github.com/ProtonVPN/protonvpn-cli)

## Install

```
go get github.com/rishispeets/repeat-pvpn
```

## Run

When you have installs available on `$PATH`:

```
sudo repeat-pvpn
```

## License

MIT