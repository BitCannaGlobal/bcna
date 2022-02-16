## Hardware requirements
The minimum required hardware specifications for running a validator or sentry node.

* 8GB RAM
* 4vCPUs
* 200GB - 300GB Disk space per year

## Software requirements
Install dependencies
``` 
sudo apt-get install build-essential jq
```

# Looking for easy instructions on how to deploy a node/validator?
The guide below describes how to compile the binaries from the source code on GitHub. If you want and easier way to install a BitCanna node you can follow [this guide](https://github.com/BitCannaGlobal/bcna/blob/main/instructions.md)

## Compile instructions: 
### Step 1) Install GoLang

Install Go 1.17.x 
The official instructions can be found in [these docs](https://golang.org/doc/install)

First remove any existing old GoLang installation as root.
```
sudo rm -rf /usr/local/go
``` 

Download the software:
```
curl https://dl.google.com/go/go1.17.7.linux-amd64.tar.gz | sudo tar -C/usr/local -zxvf -
```
Update environment variables to include Go. Simply copy and paste everything.
```
cat <<'EOF' >>$HOME/.profile
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export GO111MODULE=on
export GOBIN=$HOME/go/bin
export PATH=$PATH:/usr/local/go/bin:$GOBIN
EOF
source $HOME/.profile
```
To verify that Go is installed:
``` 
go version
```
This should return: `go version go1.17.7 linux/amd64`

### Step 2) Download and compile BCNAD source code
```
git clone https://github.com/BitCannaGlobal/bcna.git
cd bcna
git checkout v.1.3.1.pre1
make build   #it build the binary in build/ folder
```
Check if you have the correct version:
```
build/bcnad version
```
This should return: `1.3.1.pre1`

If you have the correct version, you have two options:
* a) Move the binary to the */usr/local/bin/* path with: `sudo mv build/bcnad /usr/local/bin/` or,
* b) Compile and install the binary in the *$GOPATH* path:  `make install`

If you are using Cosmovisor you will need to move the binary to the correct Cosmovisor folder.
`sudo mv ./bcnad ${HOME}/.bcna/cosmovisor/genesis/bin/` or,
`sudo mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/[update-name]/bin/`

# Cosmosvisor Quick Start
Cosmovisor is a small process manager for Cosmos SDK application binaries that monitors the governance module for incoming chain upgrade proposals. If it sees a proposal that gets approved, Cosmovisor can automatically stop the current binary, switch from the old binary to the new one, and finally restart the node with the new binary. It's not recommended but Cosmovisor could also automatically download the new binary for you.

## Installation
To install Cosmovisor, compile it from the source:
```
git clone git@github.com:cosmos/cosmos-sdk
cd cosmos-sdk
git checkout cosmovisor/v1.1.0
make cosmovisor
```
* [more info about Cosmovisor](https://github.com/cosmos/cosmos-sdk/tree/master/cosmovisor#readme)
* [how BitCanna implements Cosmovisor](https://hackmd.io/PXmANfhUSGS5YlKcHcc3GA?view)

# Join our MainNet: `bitcanna-1`
* Proceed from step 1 in [this guide](https://github.com/BitCannaGlobal/bcna/blob/main/instructions.md)
