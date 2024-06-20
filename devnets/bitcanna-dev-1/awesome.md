# bitcanna-dev-1 peers and services

## Persistent Peers
- 80ee9ed689bfb329cf21b94aa12978e073226db4@212.227.151.143:26656
- 20ca909b49106aacbf516ba28fa8a2409f825a82@212.227.151.106:26656
- b0c7e5c69aaf00626baaf7c59370029b587a91a4@bitcannadev.peers.stavr.tech:30006
- c4277e884bbbf33ef1d8f13cbf26adb3b12336fd@testnet-bitcanna.konsortech.xyz:26656

## Seed Peer
This will disconnect after passing P2P info so not include in your _persistent_peers_
- 471341f9befeab582e845d5e9987b7a4889c202f@144.91.89.66:26656
- 3f472746f46493309650e5a033076689996c8881@bitcanna-testnet.rpc.kjnodes.com:42659
- 2a369fb20a5827104efafa901916e97516e829fb@testnet-bitcanna-seed.genznodes.dev:50656
- 496ac0ba20188f70f41e0a814dfd4d9a617338f8@bcnadev-seed.ibs.team:16656
- b85358e035343a3b15e77e1102857dcdaf70053b@seeds.bluestake.net:44356
- 9f7ebc84e78187421e637627bbf608a54040fb17@seeds.bitcanna-dev.hexnodes.one:04656

## Explorers
- https://testnets-cosmos.mintthemoon.xyz/bitcanna/staking
- https://testnet.ping.pub/bitcanna/uptime
- https://explorer.stavr.tech/Bitcanna-DEV
- https://explorer.kjnodes.com/bitcanna-testnet
- https://testnet-explorer.konsortech.xyz/bitcanna
- https://explorer.hexnodes.one/BITCANNA-TESTNET

## RPC endpoints
- https://rpc-testnet.bitcanna.io
- https://rpc.bitcanna-dev-1.bitcanna.aviaone.com/ (TX index = off)
- https://testnet-bitcanna-rpc.konsortech.xyz/
- https://bitcanna-testnet.rpc.kjnodes.com
- https://bcna-testnet-rpc.ibs.team ( StateSync enable)
- https://testnet-bitcanna-rpc.konsortech.xyz
- https://bitcanna-testnet-rpc.bluestake.net
- https://rpc.bitcanna-dev.hexnodes.one:443

## GRPC endpoints
- bitcanna-testnet.grpc.kjnodes.com:42090
- testnet-bitcanna-grpc.konsortech.xyz:9090
- grpc.bitcanna-dev.hexnodes.one:4090

## LCD Enpoints
- http://lcd-testnet.bitcanna.io
- https://api.bitcanna-dev-1.bitcanna.aviaone.com/
- https://bitcanna.api.dev.stavr.tech/
- https://testnet-bitcanna-api.konsortech.xyz/
- https://bitcanna-testnet.api.kjnodes.com
- https://bcna-testnet-api.ibs.team
- https://testnet-bitcanna-api.konsortech.xyz
- https://bitcanna-testnet-api.bluestake.net
- https://lcd.bitcanna-dev.hexnodes.one

## Wallets / Dashboard
- https://wallet-testnet.bitcanna.io
- https://aviaone.com/keplr-bitcanna-dev-1/
- http://212.227.151.143:1338/
- http://212.227.151.143:1337/ 

## State sync
- https://services.kjnodes.com/testnet/bitcanna/state-sync
- https://docs.konsortech.xyz/node/Testnet/Bitcanna/statesync
- https://github.com/Inter-Blockchain-Service/Cosmos-StateSync/tree/main/Testnet/Bitcanna
- https://bluestake.net/bitcanna/testnet/statesync
- https://github.com/hexskrt/testnet_installation/tree/main/Bitcanna

## Snapshots
- https://services.kjnodes.com/testnet/bitcanna/snapshot
- https://docs.konsortech.xyz/node/Testnet/Bitcanna/snapshot
- https://testnet-snapshot.ibs.team/Bitcanna/
- https://bluestake.net/bitcanna/testnet/snapshots
- https://github.com/hexskrt/testnet_installation/tree/main/Bitcanna

## Docs
- [Automatic StateSync Join `bitcanna-dev-1`](README.md)
- [Manually Join `bitcanna-dev-1` by Avione Validator](awesome.md#Manual-setup)
- [Upgrades - Last Upgrade](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/upgrade_v1.6.0-rc2.md)
- [Setup Cosmovisor from zero](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/cosmovisor_install.md)
- [Node installation instructions by kjnodes](https://services.kjnodes.com/testnet/bitcanna/installation)
- [Node installation instructions Mainnet/DEVNET by 🔥STAVR🔥](https://github.com/obajay/nodes-Guides/tree/main/Projects/Bitcanna)
- [Node upgrade instructions by kjnodes](https://services.kjnodes.com/testnet/bitcanna/upgrade)
- [Useful commands for node and wallet operation by kjnodes](https://services.kjnodes.com/testnet/bitcanna/useful-commands)
- [Testnet Service by genznodes](https://genznodes.dev/testnet_services/#bitcanna)
- [Node Installation Guide Devnet by KonsorTech](https://docs.konsortech.xyz/node/Testnet/Bitcanna)
- [Node Installation/Automatic Installation Cosmovisor](https://github.com/hexskrt/testnet_installation/tree/main/Bitcanna)

### Manual-setup
> This instructions was written by Avione Validators, thanks a lot!

> Please remove `.bcna` folder and bcnad binary if you have participated in previous devnets
```
#############################################################################
############
############ BITCANNA --chain bitcanna-dev-1
############
#############################################################################

===================================================
# Update if needed
sudo apt update && sudo apt upgrade -y

===================================================
# Install packages
sudo apt install curl tar wget clang pkg-config libssl-dev jq build-essential bsdmainutils git make ncdu unzip -y

===================================================
# Install GOLANG version 1.19.4
cd $HOME
ver="1.19.4"
wget "https://golang.org/dl/go$ver.linux-amd64.tar.gz"
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf "go$ver.linux-amd64.tar.gz"
rm "go$ver.linux-amd64.tar.gz"
echo "export PATH=$PATH:/usr/local/go/bin:$HOME/go/bin" >> $HOME/.bash_profile
source $HOME/.bash_profile
go version

===================================================
# Install App and build
cd $HOME
git clone https://github.com/BitCannaGlobal/bcna.git
cd bcna
git checkout v1.5.3
make install

#check version and install binary ok or not ?
bcnad version --long

===================================================
# Common -> set variables (Chain | Moniker | Wallet_name)
BCNA_CHAIN="bitcanna-dev-1"
BCNA_MONIKER="YOUR_MONIKER"
BCNA_WALLET="Wallet_Testnet_Bitcanna"

echo 'export BCNA_CHAIN='${BCNA_CHAIN} >> $HOME/.bash_profile
echo 'export BCNA_MONIKER='${BCNA_MONIKER} >> $HOME/.bash_profile
echo 'export BCNA_WALLET='${BCNA_WALLET} >> $HOME/.bash_profile
source $HOME/.bash_profile

===================================================
#add seed
seeds="471341f9befeab582e845d5e9987b7a4889c202f@144.91.89.66:26656"
sed -i.bak -e "s/^seeds =.*/seeds = \"$seeds\"/" $HOME/.bcna/config/config.toml

#avoid error and set gas prices
sed -E -i 's/minimum-gas-prices = \".*\"/minimum-gas-prices = \"0.001ubcna\"/' $HOME/.bcna/config/app.toml

#download genesis
cd $HOME
wget -O $HOME/.bcna/config/genesis.json "https://raw.githubusercontent.com/bitcannaglobal/bcna/main/devnets/bitcanna-dev-1/genesis.json"

===================================================
# Clean all before to start the chain
bcnad tendermint unsafe-reset-all --home $HOME/.bcna

===================================================
# create service
sudo tee /etc/systemd/system/bcnad.service > /dev/null <<EOF
[Unit]
Description=BCNA\n
After=network.target
[Service]
Type=simple
User=$USER
ExecStart=$(which bcnad) start
Restart=on-failure
RestartSec=10
LimitNOFILE=65535
[Install]
WantedBy=multi-user.target
EOF

# Start service
sudo systemctl enable bcnad
sudo systemctl daemon-reload
sudo systemctl restart bcnad && journalctl -u bcnad.service -f

===================================================
curl localhost:26657/status

#must return
#"network": "bitcanna-dev-1",
#      "version": "v0.34.22",
#      "channels": "40202122233038606100",
#
#
    "sync_info": {
#
#
      "catching_up": false

===================================================
======== CREATE OR RESTORE YOUR WALLET  ===========
===================================================
#Restore wallet
bcnad keys add $BCNA_WALLET --recover

#create wallet
bcnad keys add $BCNA_WALLET

#RETURN
- name: Wallet_Testnet_Bitcanna
  type: local
  address: bcna1fvl95kvcn5y7292xvt2cy8jwcgu92j34tcu42d
  pubkey: '{"@type":"/cosmos.crypto.secp256k1.PubKey","key":"AizMWWksR+VB78sKNTcXJAOqfYQf9rLj53LYCM+SCfJy"}'
  mnemonic: ""

**Important** write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget your password.

soft car cigar razor buddy analyst home actress accuse ignore drastic error garment split draft inform swarm rely approve confirm cluster disagree copy gather

#go faucet to get token https://discord.com/channels/805725188355260436/847019574662922260
#Enter this command
#replace bcna1fvl95kvcn5y7292xvt2cy8jwcgu92j34tcu42d by your address just created above !
!claim bcna1fvl95kvcn5y7292xvt2cy8jwcgu92j34tcu42d

# Check balance
#replace bcna1fvl95kvcn5y7292xvt2cy8jwcgu92j34tcu42d by your address just created above !
bcnad q bank balances bcna1fvl95kvcn5y7292xvt2cy8jwcgu92j34tcu42d

===================================================
========          CREATE VALIDATOR      ===========
===================================================
bcnad tx staking create-validator \
 --amount=10000000000ubcna \
 --pubkey=$(bcnad tendermint show-validator) \
 --moniker=$BCNA_MONIKER \
 --chain-id=$BCNA_CHAIN \
 --commission-rate=0.1 \
 --commission-max-rate=0.2 \
 --commission-max-change-rate=0.01 \
 --min-self-delegation=1 \
 --fees="0.001ubcna" \
 --from=$BCNA_WALLET
 
===================================================
========             DONE               ===========
===================================================
```
