# checkers
**checkers** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).

## Get started

```
ignite chain serve
```

`serve` command installs dependencies, builds, initializes, and starts your blockchain in development.

### Configure

Your blockchain in development can be configured with `config.yml`. To learn more, see the [Ignite CLI docs](https://docs.ignite.com).

### Web Frontend

Ignite CLI has scaffolded a Vue.js-based web app in the `vue` directory. Run the following commands to install dependencies and start the app:

```
cd vue
npm install
npm run serve
```

The frontend app is built using the `@starport/vue` and `@starport/vuex` packages. For details, see the [monorepo for Ignite front-end development](https://github.com/ignite/web).

## Release
To release a new version of your blockchain, create and push a new tag with `v` prefix. A new draft release with the configured targets will be created.

```
git tag v0.1
git push origin v0.1
```

After a draft release is created, make your final changes from the release page and publish it.

### Install
To install the latest version of your blockchain node's binary, execute the following command on your machine:

```
curl https://get.ignite.com/username/checkers@latest! | sudo bash
```
`username/checkers` should match the `username` and `repo_name` of the Github repository to which the source code was pushed. Learn more about [the install process](https://github.com/allinbits/starport-installer).

## Learn more

- [Ignite CLI](https://ignite.com/cli)
- [Tutorials](https://docs.ignite.com/guide)
- [Ignite CLI docs](https://docs.ignite.com)
- [Cosmos SDK docs](https://docs.cosmos.network)
- [Developer Chat](https://discord.gg/ignite)


## Chain Initialization

```sh
ignite chain serve
```

```
Cosmos SDK's version is: stargate - v0.46.6

🛠  Building proto...
📦 Installing dependencies...
🛠  Building the blockchain...
💿 Initializing the app...
🗂  Initialize accounts...
🙂 Added account alice with address cosmos179kuaphscu08n2zkvlk8sxxu70au3zztupdsrm and mnemonic: thumb cheese cart fatigue movie poem pet snap erupt horror armor escape naive bulb huge depth crumble try fame cup quantum cupboard junk question
🙂 Added account bob with address cosmos1e3n4syctjgflkpf2vqje0ptytjhvasua6jnxvm and mnemonic: bread film bus glad moment script attract join essay rifle pistol kiss drama renew sand certain raven cash together hockey tumble slice piece rice
🌍 Tendermint node: http://0.0.0.0:26657
🌍 Blockchain API: http://0.0.0.0:1317
🌍 Token faucet: http://0.0.0.0:4500
```
### Accounts

### Alice

Added account alice with 

    - address cosmos14glrpzanps4s8qrfk6c7fxwksk234j4wd9n85z
    - mnemonic: left cause annual purity twenty kick dry usual gasp advice deal belt check slam caught jewel taste evoke involve used prison slot tag place

### Bob

Added account bob with 

    - address cosmos1um08kscm5yw4v9ch23ffzejz0nelgemey6eeus 
    - mnemonic: turtle cloud monkey bonus couch pulp car mesh execute velvet boy zoo robot cousin idle hole below strike topic insect adult frown tag feed

## Create a Message

```sh
ignite scaffold message createPost title body
```

### Modified Files

```
modify proto/checkers/checkers/tx.proto
modify x/checkers/client/cli/tx.go
modify x/checkers/module_simulation.go
modify x/checkers/types/codec.go
```

### Created Files

```
create x/checkers/client/cli/tx_create_post.go
create x/checkers/keeper/msg_server_create_post.go
create x/checkers/simulation/create_post.go
create x/checkers/types/message_create_post.go
create x/checkers/types/message_create_post_test.go

🎉 Created a message `createPost`.
```

### Store Object - Make a Checkers Blockchain


**SystemInfo**

```sh
ignite scaffold single systemInfo nextId:uint --module checkers --no-message
```

#### Modified Files

```
modify proto/checkers/checkers/genesis.proto
modify proto/checkers/checkers/query.proto
modify x/checkers/client/cli/query.go
modify x/checkers/genesis.go
modify x/checkers/genesis_test.go
modify x/checkers/types/genesis.go
modify x/checkers/types/genesis_test.go
modify x/checkers/types/keys.go
```

## Created Files

```
create proto/checkers/checkers/system_info.proto
create x/checkers/client/cli/query_system_info.go
create x/checkers/client/cli/query_system_info_test.go
create x/checkers/keeper/grpc_query_system_info.go
create x/checkers/keeper/grpc_query_system_info_test.go
create x/checkers/keeper/system_info.go
create x/checkers/keeper/system_info_test.go
```

🎉 systemInfo added.


**StoreGame**

```sh
ignite scaffold map storedGame board turn black red --index index --module checkers --no-message
```

#### Modified Files

```
modify proto/checkers/checkers/genesis.proto
modify proto/checkers/checkers/query.proto
modify x/checkers/client/cli/query.go
modify x/checkers/genesis.go
modify x/checkers/genesis_test.go
modify x/checkers/types/genesis.go
modify x/checkers/types/genesis_test.go
```

#### Created Files

```
create proto/checkers/checkers/stored_game.proto
create x/checkers/client/cli/query_stored_game.go
create x/checkers/client/cli/query_stored_game_test.go
create x/checkers/keeper/grpc_query_stored_game.go
create x/checkers/keeper/grpc_query_stored_game_test.go
create x/checkers/keeper/stored_game.go
create x/checkers/keeper/stored_game_test.go
create x/checkers/types/key_stored_game.go
```

🎉 storedGame added.

## Create Custom Messages

Checkers CRUD implementation

```sh
ignite scaffold message createGame black red --module checkers --response gameIndex
```

This creates & Modifies a certain files plus some GUI elements

#### Modified Files

```
modify proto/checkers/checkers/tx.proto
modify x/checkers/client/cli/tx.go
modify x/checkers/module_simulation.go
modify x/checkers/types/codec.go
```

#### Created Files

```
create x/checkers/client/cli/tx_create_game.go
create x/checkers/keeper/msg_server_create_game.go
create x/checkers/simulation/create_game.go
create x/checkers/types/message_create_game.go
create x/checkers/types/message_create_game_test.go
```

🎉 Created a message `createGame`.

## Add a Way to Make a Move


```sh
ignite scaffold message playMove gameIndex fromX:uint fromY:uint toX:uint toY:uint \
    --module checkers \
    --response capturedX:int,capturedY:int,winner
```

### Modified Files

```
modify proto/checkers/checkers/tx.proto
modify x/checkers/client/cli/tx.go
modify x/checkers/module_simulation.go
modify x/checkers/types/codec.go
```

### Created Files

```
create x/checkers/client/cli/tx_play_move.go
create x/checkers/keeper/msg_server_play_move.go
create x/checkers/simulation/play_move.go
create x/checkers/types/message_play_move.go
create x/checkers/types/message_play_move_test.go
```

🎉 Created a message `playMove`.