<div align="center">
  <img src="https://www.quadrans.io/assets/brand/logo_quadrans_color.svg"><br>
</div>

-----------------

## Go Quadrans

This is the official Golang implementation of the Quadrans protocol.

[![Go Report Card](https://goreportcard.com/badge/github.com/quadrans/go-quadrans)](https://goreportcard.com/report/github.com/quadrans/go-quadrans)
[![Telegram](https://shields.io/badge/telegram-join-16173D?logo=telegram&style=flat)](https://t.me/quadrans)
[![Discord](https://shields.io/badge/discord-join-16173D?logo=discord&style=flat)](https://discord.gg/2bMpuU9Z)
[![Reddit](https://shields.io/badge/reddit-join-16173D?logo=reddit&style=flat)](https://reddit.com/r/quadrans)

**[Quadrans](https://quadrans.io)** is a public blockchain infrastructure that **runs Smart Contracts and decentralised applications (dApps)**, with particular focus on improving the major challenges that blockchain ecosystems are facing today such as security, interoperability, scalability and high operational costs. 

With a mind focused on the needs of **Industry, complex supply chains, IOT devices**, Quadrans finds its roots in a fork of the [Go-Ethereum's source code](https://github.com/ethereum/go-ethereum/). To keep the cost of operations balanced, the distribution of Quadrans Coin (QDC) required to write on the blockchain is carried out through a minting protocol.

The **full compatibility with the Ethereum environment** makes migrating a project to the Quadrans infrastructure effortless and results in a significant reduction in operating costs.

Binary archives are published at http://repo.quadrans.io/.

## Building the source

For prerequisites and detailed build instructions please read the
[Installation Instructions](https://docs.quadrans.io/nodes/build/source-code.html) on the Quadrans Documentation.

Building gqdc requires both a Go (version 1.10 or later) and a C compiler.
You can install them using your favourite package manager.
Once the dependencies are installed, run

    make gqdc

or, to build the full suite of utilities:

    make all

## Executables

The go-quadrans project comes with several wrappers/executables found in the `cmd` directory.

| Command    | Description |
|:----------:|-------------|
| **`gqdc`** | Our main Quadrans CLI client. It is the entry point into the Ethereum network (main-, test- or private net), capable of running as a full node (default), archive node (retaining all historical state) or a light node (retrieving data live). It can be used by other processes as a gateway into the Quadrans network via JSON RPC endpoints exposed on top of HTTP, WebSocket and/or IPC transports. `gqdc --help` and the [CLI Documentation page](https://docs.quadrans.io/nodes/management/command_line_options.html) for command line options. |
| `abigen` | Source code generator to convert Quadrans contract definitions into easy to use, compile-time type-safe Go packages. It operates on plain [Solidity Contract ABI Specification](https://docs.soliditylang.org/en/develop/abi-spec.html) with expanded functionality if the contract bytecode is also available. However, it also accepts Solidity source files, making development much more streamlined. Please see our [Native DApps](https://docs.quadrans.io/programming/native-bindings.html) documentation page for details. |
| `bootnode` | Stripped down version of our Quadrans client implementation that only takes part in the network node discovery protocol, but does not run any of the higher level application protocols. It can be used as a lightweight bootstrap node to aid in finding peers in private networks. |

## Running gqdc

Going through all the possible command line flags is out of scope here (please consult our
[CLI Documentation page](https://docs.quadrans.io/nodes/management/command_line_options.html)), but we've
enumerated a few common parameter combos to get you up to speed quickly on how you can run your
own Gqdc instance.

### Full node on the Quadrans mainnet network

By far the most common scenario is people wanting to simply interact with the Quadrans network: create accounts; transfer funds; deploy and interact with contracts. For this particular use-case the user doesn't care about years-old historical data, so we can fast-sync quickly to the current
state of the network. To do so:

```
$ gqdc console
```

This command will:

 * Start gqdc in fast sync mode (default, can be changed with the `--syncmode` flag), causing it to download more data in exchange for avoiding processing the entire history of the Quadrans network, which is CPU intensive.
 * Start up Gqdc's built-in interactive [JavaScript console](https://docs.quadrans.io/nodes/management/javascript-console.html), (via the trailing `console` subcommand) through which you can invoke all official [`web3` methods](https://web3js.readthedocs.io/) as well as Gqdc's [management APIs](https://docs.quadrans.io/programming/rpc-server.html).    This tool is optional and if you leave it out you can always attach to an already running Gqdc instance with `gqdc attach`.

### A Full node on the Quadrans testnet network

Transitioning towards developers, if you'd like to play around with creating Ethereum contracts, you almost certainly would like to do that without any real money involved until you get the hang of the entire system. In other words, instead of attaching to the main network, you want to join the **testnet** network with your node, which is fully equivalent to the main network, but with play-Quadrans only.

```
$ gqdc --testnet console
```

The `console` subcommand has the exact same meaning as above and they are equally useful on the testnet too. Please see above for their explanations if you've skipped here.

Specifying the `--testnet` flag, however, will reconfigure your Gqdc instance a bit:

 * Instead of using the default data directory (`~/.quadrans` on Linux for example), Gqdc will nest itself one level deeper into a `testnet` subfolder (`~/.quadrans/testnet` on Linux). Note, on OSX and Linux this also means that attaching to a running testnet node requires the use of a custom endpoint since `gqdc attach` will try to attach to a production node endpoint by default. E.g. `gqdc attach <datadir>/testnet/gqdc.ipc`. Windows users are not affected by this.
 * Instead of connecting the main Quadrans network, the client will connect to the testnet network, which uses different P2P bootnodes, different network IDs and genesis states.
 * Instead of using Mainnet Quadrans Coins (QDC) for paying transaction costs it will use free [faucet Quadrans Coins](https://faucetpage.quadrans.io).
   
*Note: Although there are some internal protective measures to prevent transactions from crossing over between the main network and test network, you should make sure to always use separate accounts for play-money and real-money. Unless you manually move accounts, Gqdc will by default correctly separate the two networks and will not make any accounts available between them.*

### Configuration

As an alternative to passing the numerous flags to the `gqdc` binary, you can also pass a configuration file via:

```
$ gqdc --config /path/to/your_config.toml
```

To get an idea how the file should look like you can use the `dumpconfig` subcommand to export your existing configuration:

```
$ gqdc --your-favourite-flags dumpconfig
```

### Programmatically interfacing Gqdc nodes

As a developer, sooner rather than later you'll want to start interacting with Gqdc and the Quadrans network via your own programs and not manually through the console. To aid this, Gqdc has built-in support for a JSON-RPC based APIs ([standard APIs](https://eth.wiki/json-rpc/API) and [Gqdc specific APIs](https://docs.quadrans.io/programming/rpc-server.html). These can be exposed via HTTP, WebSockets and IPC (UNIX sockets on UNIX based platforms, and named pipes on Windows).

The IPC interface is enabled by default and exposes all the APIs supported by Gqdc, whereas the HTTP and WS interfaces need to manually be enabled and only expose a subset of APIs due to security reasons. 
These can be turned on/off and configured as you'd expect.

HTTP based JSON-RPC API options:

  * `--rpc` Enable the HTTP-RPC server
  * `--rpcaddr` HTTP-RPC server listening interface (default: "localhost")
  * `--rpcport` HTTP-RPC server listening port (default: 8545)
  * `--rpcapi` API's offered over the HTTP-RPC interface (default: "eth,net,web3")
  * `--rpccorsdomain` Comma separated list of domains from which to accept cross origin requests (browser enforced)
  * `--ws` Enable the WS-RPC server
  * `--wsaddr` WS-RPC server listening interface (default: "localhost")
  * `--wsport` WS-RPC server listening port (default: 8546)
  * `--wsapi` API's offered over the WS-RPC interface (default: "eth,net,web3")
  * `--wsorigins` Origins from which to accept websockets requests
  * `--ipcdisable` Disable the IPC-RPC server
  * `--ipcapi` API's offered over the IPC-RPC interface (default: "admin,debug,eth,miner,net,personal,shh,txpool,web3")
  * `--ipcpath` Filename for IPC socket/pipe within the datadir (explicit paths escape it)

You'll need to use your own programming environments' capabilities (libraries, tools, etc) to connect
via HTTP, WS or IPC to a Gqdc node configured with the above flags and you'll need to speak [JSON-RPC](https://www.jsonrpc.org/specification)
on all transports. You can reuse the same connection for multiple requests!

**Note: Please understand the security implications of opening up an HTTP/WS based transport before doing so! Hackers on the internet are actively trying to subvert Quadrans nodes with exposed APIs! Further, all browser tabs can access locally running web servers, so malicious web pages could try to subvert locally available APIs!**

## Contribution

Thank you for considering to help out with the source code! We welcome contributions from anyone on the internet, and are grateful for even the smallest of fixes!

If you'd like to contribute to go-quadrans, please fork, fix, commit and send a pull request for the maintainers to review and merge into the main code base. If you wish to submit more complex changes though, please check up with the core devs first on our [Discord](https://discord.gg/2bMpuU9Z), [Telegram](https://t.me/quadrans) or [Reddit](https://reddit.com/r/quadrans) channels to ensure those changes are in line with the general philosophy of the project and/or get some early feedback which can make both your efforts much lighter as well as our review and merge procedures quick and simple.

Please make sure your contributions adhere to our coding guidelines:

 * Code must adhere to the official Go [formatting](https://golang.org/doc/effective_go.html#formatting) guidelines (i.e. uses [gofmt](https://golang.org/cmd/gofmt/)).
 * Code must be documented adhering to the official Go [commentary](https://golang.org/doc/effective_go.html#commentary) guidelines.
 * Pull requests need to be based on and opened against the `main` branch.
 * Commit messages should be prefixed with the package(s) they modify.
   * E.g. "eth, rpc: make trace configs optional"

Please see the [Developers' Guide](https://docs.quadrans.io/development/devguide.html) for more details on configuring your environment, managing project dependencies, and testing procedures.

## License

The go-quadrans library (i.e. all code outside of the `cmd` directory) is licensed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html), also included in our repository in the `COPYING.LESSER` file.

The go-quadrans binaries (i.e. all code inside of the `cmd` directory) is licensed under the [GNU General Public License v3.0](https://www.gnu.org/licenses/gpl-3.0.en.html), also included in our repository in the `COPYING` file.

