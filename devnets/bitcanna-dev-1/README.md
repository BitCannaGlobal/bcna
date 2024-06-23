# Devnet-1: Setup up your validator and join *bitcanna-dev-1*

## Target of this DevNet..
We created a new testnet that follows the exact upgrade cycle as the current Mainnet. Our target is to test new features here before apply them at MainNet.

### We are going to work in three new testnets: 
* `bitcanna-dev-1`
* `innuendo-5` (with `quicksilverd` current version binary) YOU DON'T NEED TO RUN THIS CHAIN.
* `axelar-testnet-lisbon-3` (with `axelard` current version binary) YOU DON'T NEED TO RUN THIS CHAIN.

## P2P Network INFO
In the next section you will find the params for P2P and a *step by step* guide with a StateSync script ([Step by step guide](#running-a-validator-on-bitcanna-dev-5))

#### Seed server
* `471341f9befeab582e845d5e9987b7a4889c202f@144.91.89.66:26656`

#### Persistent peers
* `80ee9ed689bfb329cf21b94aa12978e073226db4@81.0.247.144:26656`
* `ba6c17d707cb0c4f81e0ef590f2e36152ff7dd1a@212.227.151.106:26656`


#### Genesis file
* [Link to Genesis file](https://raw.githubusercontent.com/BitCannaGlobal/bcna/main/devnets/bitcanna-dev-1/genesis.json)

#### Binary / Upgrades table

| Upgrade Date | Upgrade Height | Binary Path | Release GitHub | Release doc |
| -------- | -------- | -------- | ------- | ------- |
| 22nd Dec 2022     | 0    | [v1.5.3](https://github.com/BitCannaGlobal/bcna/releases/download/v1.5.3/bcna_linux_amd64.tar.gz)    | [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.5.3) | x |
| 22nd Feb 2023 15h UTC    | 956.837    | [v1.6.0-rc2](https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.0-rc2/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.0-rc2) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/upgrade_v1.6.0-rc2.md) |
| 24nd Feb 2023 15h UTC    | 987.757    | [v1.6.0-rc3](https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.0-rc3/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.0-rc3) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/upgrade_v1.6.0-rc3.md) |
| 7th Mar 2023 15h UTC    | 1.159.488   | [v1.6.1](https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.1/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.1) | [Doc](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/upgrade_v1.6.1.md) |
| No break consensus   | x   | [v1.6.2](https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.2/bcna_linux_amd64.tar.gz)   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.2) | x |
| No break consensus   | x   | v1.6.3-rc1   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.3-rc1/bcna_linux_amd64.tar.gz) | x |
| 19th Apr 2023 14h UTC   | 1.831.901   | v2.0.1-rc6   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/download/v2.0.1-rc6/bcna_linux_amd64.tar.gz) | [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v2.0.1-rc6) |
| 9th Nov 2023 14h UTC   | 5.001.067   | v3.0.0-rc3   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/download/v3.0.0-rc3/bcna_linux_amd64.tar.gz) | [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v3.0.0-rc3) |
| No break consensus   | x   | v3.0.1-rc1   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v3.0.1-rc1/bcna_linux_amd64.tar.gz) | x |
| No break consensus   | x   | v3.0.2-rc1   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v3.0.2-rc1/bcna_linux_amd64.tar.gz) | x |
| 17th Apr 2024 16h UTC  | 7.400.000   | v3.1.0-rc2   |  [Link](https://github.com/BitCannaGlobal/bcna/releases/download/v3.1.0-rc2/bcna_linux_amd64.tar.gz) | [Link](https://github.com/BitCannaGlobal/bcna/releases/tag/v3.1.0-rc2) |

#### More DEVNET-1 resources
* [Link to AWESOME doc](awesome.md)
* [Link to manual set-up by Aviaone Validator](awesome.md#Manual-setup)
* [Link to setup and other tools by KonsorTech](https://github.com/konsortech/Node/tree/main/Testnet/Bitcanna)
* [Link to setup Cosmovisor for DEVNET-1](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/cosmovisor_install.md)

## Running a validator on **bitcanna-dev-1** using StateSync
* Before you start, you want to ensure your system is updated.  Besides other utilities you can install `jq` which is a utility to read and navigate JSON files and output. Also remove the `.bcna` folder, take in consideration first NOTES in this doc.
    ```
    sudo apt update && sudo apt upgrade -y
    sudo apt install -y build-essential curl wget jq
    ```
* Increase the default open files limit. If we don't raise this value nodes will crash once the network grows large enough.
    ```
    sudo su -c "echo 'fs.file-max = 65536' >> /etc/sysctl.conf"
    sudo sysctl -p
    ```
## Step 1 - Download and run statesync script
By running the statesync script we download the latest binary (v1.5.3) and sync the chain to the latest block.
1. **Download the statesync script** for new peers from Github:
    ```
    cd ~
    wget https://raw.githubusercontent.com/BitCannaGlobal/cosmos-statesync_client/main/statesync_DEVNET-1_client_linux_new.sh
    ```
2. **Run the script**: 
    ```
    bash statesync_DEVNET-1_client_linux_new.sh
    ```
    Watch the output of the logs. When the chain is synced (wait 1 minute), press ctrl+C to stop the script and proceed with the next step.
    
    This is an example of the output you see when you can stop the script by pressing ctrl+C
    ```go
    4:39PM INF indexed block height=2920 module=txindex
    4:39PM INF Timed out dur=4988.140195 height=2921 module=consensus round=0 step=1
    4:39PM INF commit is for a block we do not know about; set ProposalBlock=nil commit=3E75B8B4371324172A860BBBB4BE8B5C2A2C96A7FA5F5507BB8457D0B40F00D2 commit_round=0 height=2921 module=consensus proposal={}
    4:39PM INF received complete proposal block hash=3E75B8B4371324172A860BBBB4BE8B5C2A2C96A7FA5F5507BB8457D0B40F00D2 height=2921 module=consensus
    4:39PM INF finalizing commit of block hash={} height=2921 module=consensus num_txs=0 root=E8705846BEAAA45BC87474D9ACBFBA074447ED8A680FAB5AD53516E7E0B2C7C7
    4:39PM INF minted coins from module account amount=4836690ubcna from=mint module=x/bank
    4:39PM INF executed block height=2921 module=state num_invalid_txs=0 num_valid_txs=0
    4:39PM INF commit synced commit=436F6D6D697449447B5B36352037332032303120393220313733203133203537203930203138352036342035382031323520323230203133392031313620313730203336203932203535203131382031303920363520323037203138382037312031333520313236203234352031343820353520313837203235305D3A4236397D
    4:39PM INF committed state app_hash=4149C95CAD0D395AB9403A7DDC8B74AA245C37766D41CFBC47877EF59437BBFA height=2921 module=state num_txs=0
    4:39PM INF indexed block height=2921 module=txindex
    ```
    
3. **Move the new `bcnad` binary** to your machine's PATH.
    ```bash
    sudo mv bcnad /usr/local/bin/ 
    ```
    **Optionally:**
    ```bash
    bcnad config chain-id bitcanna-dev-1
    ```
	3.1. **MemPool security settings:**
	Before start you need to set this custom config for MemPool at `.bcna/config/config.toml/` to prevent Spam Storms: (reset to apply)
	* max_tx_bytes = 524288
	* max_txs_bytes = 268435456
	```
	sed -i 's/^max_tx_bytes =.*/max_tx_bytes = 524288/' $HOME/.bcna/config/config.toml
	sed -i 's/^max_txs_bytes =.*/max_txs_bytes = 268435456/' $HOME/.bcna/config/config.toml
	```

     If you are running the binary as a service use:
    ```bash
    sudo service bcnad restart
    ```
    If you are running the binary without a service (note that it is always advised to run the binary as a service):
    ```bash
    bcnad stop (or use CTRL + C in the terminal window where the binary is running)
    bcnad start
    ```   
## Step 2 - Prepare the node
To create a validator you need a funded wallet. Once the wallet is created, go to the **#devnet-faucet** channel on [Discord](https://discord.com/channels/805725188355260436/847019574662922260) and claim your devnet coins. For example: `!claim bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`

1. **Create a wallet:**
You can create a wallet with one or more keys (addresses) using `bcnad`.  Replace **"MyFirstAddress"** with your desired name.
	```
    bcnad keys add MyFirstAddress

      name: MyFirstAddress
      type: local
      address: bcna14shzreglay98us0hep44hhhuy7dm43snv38plr
      pubkey: bcnapub1addwnpepqvtpzyugupvcu773rzdcvhele6e22txy2zr235dn7uf8t2mlqcarcyx2gg9
      mnemonic: ""
      threshold: 0
      pubkeys: []


     Important write this mnemonic phrase in a safe place.
    It is the only way to recover your account if you ever forget your password.

    deposit daring slim glide spend dolphin expire shadow cluster weed orphan work 420 section client friend yellow west hamster torch settle island opinion gloom
	```
	Your address will look something similar like this: `bcna14shzreglay98us0hep44hhhuy7dm43snv38plr`

2. **Service creation**
With all configurations ready you can set up `systemd` to run the node daemon with auto-restart.
Setup `bcnad` systemd service (copy and paste all to create the file service):
   ```
	cd $HOME
	echo "[Unit]
	Description=BitCanna Node
	After=network-online.target
	[Service]
	User=${USER}
	ExecStart=$(which bcnad) start
	Restart=always
	RestartSec=3
	LimitNOFILE=4096
	[Install]
	WantedBy=multi-user.target
	" >bcnad.service
   ```
	Enable and activate the BCNAD service.

	```
	sudo mv bcnad.service /lib/systemd/system/
	sudo systemctl enable bcnad.service && sudo systemctl start bcnad.service
	```

	Check the logs to see if everything is working correct:
   ```
    sudo journalctl -fu bcnad
   ```

## Step 3 - Create the validator
When your node is synced and your wallet funded it's time to send the TX to become validator:
(change **_WALLET_NAME_** and **_MONIKER_**)
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
	--moniker MONIKER \
	--pubkey $(bcnad tendermint show-validator) \
	--chain-id bitcanna-dev-1 \
	--gas auto \
	--gas-adjustment 1.5 \
	--gas-prices 0.001ubcna
```

You can check the list of validators (also in [Explorer](https://testnet.ping.pub/bitcanna/staking)):

   ```
   bcnad query staking validators --output json| jq
   ```

Another **IMPORTANT** but **optional** action is backup your Validator_priv_key:

   ```
    tar -czvf validator_key.tar.gz .bcna/config/*_key.json 
    gpg -o validator_key.tar.gz.gpg -ca validator_key.tar.gz
    rm validator_key.tar.gz
   ```
   This will create a GPG encrypted file with both key files.
