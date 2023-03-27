# shogi-game-core

![build status](https://github.com/mainiak/shogi-game-core/actions/workflows/main.yml/badge.svg?event=push) - [click here for more details](https://github.com/mainiak/shogi-game-core/actions/workflows/main.yml)

## About

This is a game core logic library for board game Shogi

## Usage

```
import (
    game "github.com/mainiak/shogi-game-core/v1"
)

func main() {
    g := game.NewGame()
    g.SetDefaultBoard() // required
}
```

## Build

```
git clone git@github.com:mainiak/shogi-game-core.git
cd shogi-game-core
go build .
```

## Testing

### Basic

```
go test .
```

### More verbose

Install from https://onsi.github.io/ginkgo/

```
go get github.com/onsi/ginkgo/v2
go install github.com/onsi/ginkgo/v2/ginkgo

ginkgo -vv ./...
```

https://onsi.github.io/ginkgo/#ginkgo-cli-overview
