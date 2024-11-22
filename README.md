# BitCanna Blockchain

BitCanna provides a decentralized payment network, supply chain and trust network for the legal cannabis industry by utilizing the BCNA coin.

> [Current version (v3.1.0)](https://github.com/BitCannaGlobal/bcna/releases/tag/v3.1.0) of our development uses Cosmos SDK v0.47.11 & CometBFT v0.37.5 & IBC-go v7.4.0

These docs at Github are chain related. For more information about our coin, partners and roadmap visit:
* Our website: https://www.bitcanna.io
* Our docs platform: https://docs.bitcanna.io

## Hardware Requirements
Here are the minimal hardware configs required for running a validator/sentry node:

* 8GB RAM 
* 4vCPUs (8vCPUs is recommended)
* 200GB - 300GB SSD Disk space per year (NVMe disks are recommended)
* 400 Mbit/s bandwidth (800Mbit/s - 1Gbit/s recommended)

## Software Requirements
* Linux server (Ubuntu 20/22 server recommended)
* Go version v1.21.6

> Please avoid cheap VPS providers as a main validator (we suggest using it as a cheap backup). We advise to use a shared dedicated server or a high-end NVMe VPS.

## Genesis file
* [Link to Genesis file](https://raw.githubusercontent.com/BitCannaGlobal/bcna/main/genesis.json)

## Binary / Upgrades table

| Upgrade Date | Upgrade Height | Binary Path | Release GitHub | Release notes |
| -------- | -------- | -------- | ------- | ------- |
| 11th Aug 2022 15.00h CET    | 2.092.453    | [v1.3.1](https://github.com/BitCannaGlobal/bcna/releases/download/v.1.3.1/bcnad)    | [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v.1.3.1) | [-NA-]() |
| 11th Aug 2022 15.00h CET    | 4.490.420    | [v1.4.1](https://github.com/BitCannaGlobal/bcna/releases/download/v1.4.1/bcna_linux_amd64.tar.gz)    | [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.4.1) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md#v141-codename-strangebuddheads-halt-chain-on-11th-of-august-2022) |
| 7th Nov 2022 15.00h CET    | 5.787.420    | [v1.5.3](https://github.com/BitCannaGlobal/bcna/releases/download/v1.5.3/bcna_linux_amd64.tar.gz)    | [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.5.3) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md#v153-codename-trichomemonster-ica---from-v145) |
| 2nd Mar 2023 16.20h CET    | 7.467.420   | [v1.6.0-fix](https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.0-fix/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.0-fix) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md#v160-fix-codename-vigorous-grow---from-v153) |
| 10th Mar 2023 18.40h CET    | 7.585.420   | [v1.6.3](https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.3/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.3) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md#v161-codename-vigorous-grow-fix---from-v161) |
| 30th May 2023 15.40h CEST    | 8.771.420   | [v1.7.0](https://github.com/BitCannaGlobal/bcna/releases/download/v1.7.0/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.7.0) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md#v170-codename-vigorous-grow-huckleberry-from-v163) |
| 29th Jun 2023 16.20h CEST    | 9.209.420   | [v2.0.2](https://github.com/BitCannaGlobal/bcna/releases/download/v2.0.2/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v2.0.2) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md) |
| 25th Jan 2024 16.35h CET    | 12.288.420   | [v3.0.2](https://github.com/BitCannaGlobal/bcna/releases/download/v3.0.2/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v3.0.2) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md) |
| 10th May 2024 16.55h CET    | 13.846.420   | [v3.1.0](https://github.com/BitCannaGlobal/bcna/releases/download/v3.1.0/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v3.1.0) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md) |

> Current version: https://github.com/BitCannaGlobal/bcna/releases/tag/v3.1.0

# How to join `bitcanna-1` chain
> Tip: At the end of this doc, you will find links with important info

The necessary steps to join to `bitcanna-1` chain are:

## 1. Install / Compile last binary
Check this [link](https://github.com/BitCannaGlobal/bcna/blob/main/1.install-compile.md) to find instructions to install or compile our latest release.

##   2. Sync the chain
You can sync the chain in your server to run a validator or a simple node/peer.
**Select one method of the below:**
###   2.1. Sync using a StateSync snapshot server.
This is the recommended option for new nodes/validators.

By downloading the StateSync script, you will sync the last blocks of the chain from two BitCanna StateSync servers. 

The script will setup your folder and config if you start a fresh install. If your node/validator exist will make a backup, syncing and restoring the backup at the end of the script. Choose the best that fits you!
* https://github.com/BitCannaGlobal/bcna/blob/main/2.1.statesync.md

###   2.2. Sync using a snapshot file
This is an alternative way to get synced without having to download the entire chain block by block. Recommended for advanced users in Cosmos chains.
* https://github.com/BitCannaGlobal/bcna/blob/main/2.2.snapshot.md



##   3. Create a validator
Your node must be fully synced in order to send the TX of validator creation and start to validate the network. You can check if your node has fully synced by comparing your logs and the latest block in the explorer (https://explorer.bitcanna.io/)

**You will need coins:**
Send coins to your new address, you will need roughly 2 BCNA to run the validator (1 BCNA for self-delegation and a bit more for transactions).
* https://app.osmosis.zone/
* https://app.rango.exchange/
* https://coinmerce.io/



### 3.1. **Set the chain-id parameter** 
```
    bcnad config chain-id bitcanna-1
```
### 3.2. **MemPool custom settings**
   
   Before start you need to set this custom config for MemPool at `.bcna/config/config.toml/` to prevent Spam Storms: (reset the binary to apply)
   * max_tx_bytes = 524288
   * max_txs_bytes = 268435456
     
```bash
sed -i 's/^max_tx_bytes =.*/max_tx_bytes = 524288/' $HOME/.bcna/config/config.toml && \
sed -i 's/^max_txs_bytes =.*/max_txs_bytes = 268435456/' $HOME/.bcna/config/config.toml
```
       
### 3.3.  **Create a wallet**:
You may create a wallet with one or more keys (addresses) using `bcnad`; you can choose a name of your own liking (we strongly advice you use one word)
```
    bcnad keys add MyFirstAddress
```
```
      name: MyFirstAddress
      type: local
      address: bcna14shzreglay98us0hep44hhhuy7dm43snv38plr
      pubkey: bcnapub1addwnpepqvtpzyugupvcu773rzdcvhele6e22txy2zr235dn7uf8t2mlqcarcyx2gg9
      mnemonic: ""
      threshold: 0
      pubkeys: []

    deposit daring slim glide hello dolphin expire stoner cluster vivid orphan work pond section client friend yellow west hamster torch settle island opinion gloom
```
> It is very important to write this mnemonic phrase in a safe place.
It is the only way to recover your account if you ever forget/lose your password.
```
    deposit daring slim glide hello dolphin expire stoner cluster vivid orphan work pond section client friend yellow west hamster torch settle island opinion gloom
```

### 3.4. **Start the daemon**
If you are running the binary as a service use:
```bash
sudo service bcnad restart
```
If you are running the binary without a service (note that it is always advised to run the binary as a service):
```bash
bcnad stop (or use CTRL + C in the terminal window where the binary is running)
bcnad start
```  
### 3.5. **Send the _Create validator_ TX:**

> We recommend you read the [FAQ Chain's guide](https://github.com/BitCannaGlobal/bcna/blob/main/faq_chain.md) to understand all parameters - be aware that some values are permanent and cannot be changed at a later date.

When you have your **node synced** and your **wallet funded with coins**, send the TX to become _validator_ (change _wallet_name_ and _moniker_):
> You can use quotes to include spaces and more than two words
`--from "Royal Queen Seeds"`

```
bcnad tx staking create-validator \
    --amount 1000000ubcna \
    --commission-max-change-rate 0.10 \
    --commission-max-rate 0.2 \
    --commission-rate 0.1 \
    --from WALLET_NAME \
    --min-self-delegation 1 \
    --moniker YOUR_MONIKER \
    --pubkey $(bcnad tendermint show-validator) \
    --chain-id bitcanna-1 \
    --gas auto \
    --gas-adjustment 1.5 \
    --gas-prices 0.001ubcna
```

You can check the list of validators (also in [Explorer](https://explorer.bitcanna.io/validators)):

```
bcnad query staking validators --output json| jq
```

##   4. Backup the keys and config
Making a backup of the Validator private keys and node keys is very important. Store them encrypted also.
### 4.1. Backup your Validator_priv_key:

```
tar -czvf validator_key.tar.gz .bcna/config/*_key.json 
gpg -o validator_key.tar.gz.gpg -ca validator_key.tar.gz
rm validator_key.tar.gz
```
This will create a GPG encrypted file with both key files.
You can download the `validator_key.tar.gz.gpg`  file to your computer.

### 4.2. Export the wallet key (if you have backup the seeds keys is enough)
```
bcnad keys export MyFisrstAddress
```
```
Enter passphrase to encrypt the exported key: passwordForCryptThisKey
Enter keyring passphrase: TheWalletPassword 
-----BEGIN TENDERMINT PRIVATE KEY-----
salt: BEC519DA3C1A3BDFC74D799FE983CA6C
type: secp256k1
kdf: bcrypt

tfguD11614drcOrdnmujAm+c+FbPbAFkYbsHv/qbQ8O9nQjdQCPuXTOZsYLkYopK
FYgEgITfk980jjUfDTE25BPfJR22csjJM/qzx0Y=
=vchm
-----END TENDERMINT PRIVATE KEY-----
```
You can copy&paste the entire text above in a text file. 

##   5. Cosmovisor
Cosmovisor is a small process manager for Cosmos SDK application binaries to handle chain upgrades. It works for upgrades that has been approved through governance proposals. In these type of governance proposals the upgrade name and block height is included. If Cosmovisor finds the requirements for this upgrade in the logs of the chain binary, it stops the current binary, switch from the old binary to the new one, and finally restarts the node with the new binary.

Follow this guide to install Cosmovisor on your node. 
* https://github.com/BitCannaGlobal/bcna/blob/main/5.cosmovisor.md

# Links to important info

## Explorers:
* [MintScan based on Cosmostation](https://www.mintscan.io/bitcanna)
* [Ping-Pub explorer](https://ping.pub/bitcanna/uptime)
* [ATOMScan explorer](https://atomscan.com/bitcanna)

## Wallets
* [BitCanna web wallet](https://wallet.bitcanna.io)
* [BitCanna Mobile Android](https://play.google.com/store/search?q=bitcanna&c=apps)
* [Ping-Pub web wallet](https://ping.pub/wallet/accounts)
* [Cosmostation iOS](https://apps.apple.com/kr/app/cosmostation/id1459830339)
* [Cosmostation Android](https://play.google.com/store/apps/details?id=wannabit.io.cosmostaion)

## Genesis file
* [bitcanna-1](https://raw.githubusercontent.com/BitCannaGlobal/bcna/main/genesis.json)

## Peer, seeds and public service providers
* [BitCanna Chain Registry](https://github.com/BitCannaGlobal/bcna/blob/main/chain-registry.json) (Mandatory for Team's Delegation Program)
* [BitCanna DEVET-1 Chain Registry](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/chain-registry.json) (Mandatory for Team's Delegation Program)
* [Cosmos Chain Registry](https://github.com/cosmos/chain-registry/tree/master/bitcanna)

## Archived Guides, FAQs & docs
* [BitCanna Docs](https://docs.bitcanna.io)
* [GitHub Testnet & DevNet archive](https://github.com/BitCannaGlobal/testnet-bcna-cosmos/tree/main/instructions)
* [Current DevNet](https://github.com/BitCannaGlobal/bcna/tree/main/devnets/bitcanna-dev-1)
* [Old Instructions guide](https://github.com/BitCannaGlobal/bcna/blob/main/archived_guides/instructions.md)
* [How to perform a manual upgrade](https://github.com/BitCannaGlobal/bcna/blob/main/archived_guides/manual_update.md)
* [How to perform an upgrade with Cosmovisor (v5.0)](https://github.com/BitCannaGlobal/bcna/blob/main/5.cosmovisor.md)

###### tags: `doc` `github`