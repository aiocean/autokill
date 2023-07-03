# Auto Kill

Auto Kill is a command-line tool that allows you to automatically kill processes on your computer based on various criteria.

## Installation

To install Auto Kill, simply run:

```
go install github.com/aiocean/autokill
```

## Usage


The `autokill` command has several options that you can use to customize its behavior. To see a list of available options, run the following command:

```
autokill
```

Run as daemon:

```
daemonize $(which autokill) -max-percent 200 -period 5s -allowed-names goland,webstorm
```

You can install daemonize at: https://software.clapper.org/daemonize/
