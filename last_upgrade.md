#  MAINNET `bitcanna-1` Upgrades

## v2.0.2 codename `wakeandbake` from v1.7.0

> (Update scheduled for 29th June 2023 ~ 16.20h CEST - 14.20h UTC)
>
> Note: this is a planned upgrade, the chain **will halt** at block **[YYYYYYYYYY](https://www.mintscan.io/bitcanna/blocks/YYYYYY)**
>
> **Important note: The tag name required for the upgrade is v2.0.2**

On chain proposal: 

After 8 months testing this migration and adding new features (back-ported to v1.x branch) we decided to merge source code to produce a new BCNAD v2 based on Cosmos SDK v0.46.13 & CometBFT v0.34.28. This is a bridge to reach in a near future the SDK v0.47.

Those are the highlighted features/improvements/bug fixes:

- Chore: Lint errors fixing by @RaulBernal in Chore: Lint errors fixing #172
- Chore: some refactor coding task for BCNA v2 (Cosmos-SDK v0.46.12) for coming merging in main branch (current BCNA v1 repo (Cosmos-SDK v0.45.15))
- Check and avoid duplicates at Bitcannaid.bcnaid by @RaulBernal in Check and avoid duplicates at Bitcannaid.bcnaid #171
- New endpoint for BCNA module /BitCannaGlobal/bcna/bcna/bitcannaid/bcnaid/{bcnaid}
- Code refactor to deprecate Ignite OpenAPIConsole deps
- Refactor Gov Spam filter to reject v1 proposals, only v1beta1 are allowed for SDK v0.46
- Ante handler to prevent spam proposals at Gov module
- Bump IBC go version to v6.1
- Bump IAVL version to v0.19.5
- Fixed BCNA module path & bug with CLI
- Using Cosmos Keyring v1.2.0
- SDK v0.46 bech32fix
- Endpoint nodeconfig by @RaulBernal in Sdk46 bech32fix golang19 nodeservice #132
- Upgrades separated of App.go by @RaulBernal in Upgrades separated of App.go #133
- Deprecate ignite/cli cosmoscmd package by @RaulBernal in Deprecate ignite/cli cosmoscmd package #102
- Refactor tests and root & delete comments by @RaulBernal in Refactor tests and root & delete comments #112

### For detailed instructions check the release doc:
* https://github.com/BitCannaGlobal/bcna/releases/tag/v2.0.2
### For detailed Cosmovisor install guide:
* [Updated guide](https://github.com/BitCannaGlobal/bcna/blob/main/2.3.cosmovisor.md)


https://github.com/BitCannaGlobal/bcna/blob/main/last_upgrade.md#v170-codename-vigorous-grow-huckleberry-from-v163

## v1.7.0 codename `vigorous-grow-huckleberry` from v1.6.3
On chain proposal:
What will be upgrade & implement: 
- Huckleberry patch  
- Upgrade from IBC v3.4.0 to IBC v4.4.1   
- Upgrade Cosmos-SDK from v0.45.15 to v0.45.16  
- Upgrade CometBFT from v0.34.27 to v0.34.28 
 
Description:
A safety patch has been released for IBC Go which is advised to apply as soon as possible. Where we intended to update our blockchain with the upcoming upgrade to Cosmos-SDK v0.46.x we will now do an upgrade to be able to apply the patch. For that reason the chain needs to be halted. 
 
With this proposal we set a new upgrade height at block 8.771.420 (30-05-2023 around 15:45h CEST) with the codename vigorous-grow-huckleberry (v1.7.0). At the given height the chain will halt and all validators have to change the binary to the new version. After a brief downtime the chain will resume operations. 
 
Upgrade height: 
https://www.mintscan.io/bitcanna/blocks/8771420

## v1.6.1 codename `vigorous-grow-fix` from v1.6.0-fix 

> (Update scheduled for 10th March 2023 ~ 18.40h CET - 17.40h UTC)
>
> Note: this is a planned upgrade, the chain **will halt** at block **[7585420](https://www.mintscan.io/bitcanna/blocks/7585420)**
>
> **Important note: The tag name required for the upgrade is v1.6.1**

- Fix Ante handler to prevent spam proposals at Gov module
- Remove OpenAPI dep from Ignite; now included locally.
- Bump Cosmos SDK version to v0.45.14 to mitigate a possible way to [DDoS](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.14) a node.
- Bump Tendermint to v0.34.26 ([InformalSystem fork](https://github.com/informalsystems/tendermint/blob/v0.34.26/CHANGELOG.md#v03426)). 
- Bump IBC go version to v.3.4.0 (it will break consensus so a new start/stop is required)
- Bump IAVL version to v0.19.5
- Fixed BCNA module path & bug with CLI
- Using Cosmos Keyring v1.2.0

### For detailed instructions check the release doc:
* https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.1
### For detailed Cosmovisor install guide:
* [Updated guide](https://github.com/BitCannaGlobal/bcna/blob/main/2.3.cosmovisor.md)


----------
## v1.6.0-fix codename `Vigorous-grow` from v1.5.3 

> (Update scheduled for 2nd March 2023 ~ 16.20h CET - 15.20h UTC)
>
> Note: this is a planned upgrade, the chain **will halt** at block **7467420**
>
> **Important note: The tag name required for the upgrade is v1.6.0-fix because tag v1.6.0 contains a typo and won't match the Plan (Upgrade) Name**

- Ante handler to prevent spam proposals at Gov module
- Remove OpenAPI dep from Ignite; now included locally.
- Bump Cosmos SDK version to v0.45.14 to mitigate a possible way to [DDoS](https://github.com/cosmos/cosmos-sdk/releases/tag/v0.45.14) a node.
- Bump Tendermint to v0.34.26 ([InformalSystem fork](https://github.com/informalsystems/tendermint/blob/v0.34.26/CHANGELOG.md#v03426)). 
- Bump IBC go version to v.3.4.0 (it will break consensus so a new start/stop is required)
- Bump IAVL version to v0.19.5
- Fixed BCNA module path & bug with CLI
- Using Cosmos Keyring v1.2.0

### For detailed instructions check the release doc:
* https://github.com/BitCannaGlobal/bcna/releases/tag/v1.6.0-fix
### For detailed Cosmovisor install guide:
* [Updated guide](https://github.com/BitCannaGlobal/bcna/blob/main/2.3.cosmovisor.md)


----------
## v1.5.3 codename `TrichomeMonster-ica`   from v1.4.5 
> (Update scheduled for 7 November 2022 ~ 14.00h UTC)
* Source code refactored to deprecate `CosmosCMD` package from Ignite/CLI.
* InterChain Accounts module intregration.
* Last security and stability updates.

> Note: this is a planned upgrade, chain **will halts** at block **5787420**
### For detailed instructions check the release doc:
* https://github.com/BitCannaGlobal/bcna/releases/tag/v1.5.3

## v1.4.5 from v1.4.2 (from Oct. 31th 2022)
We've developed a new version with some updates (regarding SDK, Tendermint & IBC/go) to improve the security (Dragonberry) and stability (forks in IAVL) in our chain.
* Cosmos SDK v0.45.10 
* Tendermint v0.34.22
* IAVL v0.19.4
* IBC/go v3.3.1
* Ledger-go v0.9.3 (now support Ledger Nano Plus)

> Note: this is not a planned upgrade, chain won't be halted

### For detailed instructions check the release doc:
* https://github.com/BitCannaGlobal/bcna/releases/tag/v1.4.5

## v1.4.2 from v1.4.1 (from Sept. 7th 2022)
We've developed a new version with some updates (regarding SDK, Tendermint & IBC/go) to improve the security and stability in our chain.
* Cosmos SDK v.0.45.8 
* Tendermint v.0.34.21 (fixed `unsafe-reset-all` subcommand of `tendermint` command
* IBC/go v.3.2.0
* Ledger-go v0.9.3 (now support Ledger Nano Plus)

> Note: this is not a planned upgrade, chain won't be halted

Steps are: 
- Download or compile the new binary
- Stop the daemon (bcnad or cosmovisor)
- Replace binary in linux path or cosmovisor folder
- Start the daemo again

1. Download or compile the new binary

    *1.1. Download the new binary*
    ```
    cd $HOME
    wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v1.4.2/bcna_linux_amd64.tar.gz
    ```
    1.1.2. Check the sha256sum for the downloaded file. 
    ```
    sha256sum bcna_linux_amd64.tar.gz
    ```
    It must return: `903c63b9f668bf5208566955648279bdf0c15e73aab415d5ea5efc09ec1fc890`

    1.1.3. Verify that the version is `1.4.2`
    ```
    rm -f ./bcnad #delete old file if exist
    tar zxvf  bcna_linux_amd64.tar.gz
    rm bcna_linux_amd64.tar.gz
    chmod +x bcnad
    ./bcnad version
    ```

    *1.2. Alternatively you can compile the new binary, GoLang v1.18 required*
    
    1.2.1. Download the source code using `git`
    ```
    git clone https://github.com/BitCannaGlobal/bcna.git
    cd bcna
    git checkout v1.4.2
    make build   #it build the binary in build/ folder
    ```
    1.2.2. Verify the version:
    ```
    build/bcnad version
    ```
    The output must be `1.4.2`


2. Stop the daemon (bcnad or cosmovisor)
    ```
    sudo service bcnad stop  # for bcnad service
    sudo service cosmovisor stop #for cosmovisor
    ```

3. Replace binary in linux path or cosmovisor folder
    3.1. For Cosmovisor
    ```
    mv ./bcnad $HOME/.bcna/cosmovisor/current/bin/  #dowloaded
    mv $HOME/bcna/build/bcnad $HOME/.bcna/cosmovisor/current/bin #compiled
    ```    
    3.2. Without Cosmovisor 
    ```
    sudo mv bcnad $(which bcnad)   #downloaed
    sudo mv $HOME/bcna/build/bcnad $(which bcnad) #compiled
    ```  
4. Start the daemon again
    ```
    sudo service bcnad start  # for bcnad service
    sudo service cosmovisor start #for cosmovisor
    ```



## v1.4.1 codename `strangeBuddheads` (halt chain on 11th of August 2022)
>Outdated: this was a planed upgrade

We've developed a new version with some updates (regarding SDK, Tendermint & IBC/go) called `strangeBuddheads` to improve the security and stability in our chain.
* Cosmos SDK v.0.45.7 
* Tendermint v.0.34.20
* IBC/go v.3.1.1
* IAVL v.0.19.1

> Note: When using Cosmos SDK 0.45.7, the restart of the binary **could take up to 15 minutes** in order to build the fast cache of IAVL. The fast cache will provide a performance improvement over 0.45.6.

Additionally, we are including the SDK `authz` module which introduces the "Restake" functionality and some other on-chain functionality such as "Grants". 

Also we have fixed the `bcna` module & we have added a new extra field for "Supply Chain" tracking. 

We will upgrade to `strangeBuddheads` version by governance by forking the current chain. 


## Governance proposal: **halt-height `4490420`**
Upgrade schedule:
* Proposal: today, 8th of August, 2022 11:00 CEST
* Voting period: today, 8th August, 2022 11:00 CEST - 11th of August, 11:00 CEST
* Upgrade height: **block 4490420 (~16:05 CEST 11th of August 2022)**

https://www.mintscan.io/bitcanna/proposals/4
https://explorer.bitcanna.io/proposals/4

## Attended (manually) upgrade.

This section of the guide shows how to perform a **manual** update of the binaries after the governance proposal has been approved for the chain update. 


1. Stop your bcnad service **after you see this** in your logs `ERR UPGRADE "strangeBuddheads" NEEDED at height: XXXXXXXX`
```
sudo service bcnad stop
```
2. Compile the binary from the source (2.1) or download it (2.2).

    2.1. If you want to BUILD the binary from the source, detailed instructions can be found in this [DOC](https://github.com/BitCannaGlobal/bcna/blob/main/1.install-compile.md) of our GitHub:
    ```
    git clone https://github.com/BitCannaGlobal/bcna.git
    cd bcna && git checkout  v1.4.1
    make build
    ```
    This will produce a `bcnad` file in the `build` folder. Move to the current binary path:
    ```
    mv build/bcnad $(which bcnad)   #copy&paste don't replace anything
    ```
   *You can jump to step 3.*
    
    2.2 If you want to DOWNLOAD the compiled binary:
    ```
    cd ~
    rm bcna_linux_amd64.tar.gz #delete old file if exist
    wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v1.4.1/bcna_linux_amd64.tar.gz
    ```
    2.2.1. Check the sha256sum for the downloaded file. 
    ```
    sha256sum bcna_linux_amd64.tar.gz
    ```
    It must return: `305ff854cc892e5bd4de6f2b3d8bd745525878244f78478be6c579d56316fe9d`

    2.2.2. Verify that the version is `1.4.1`
    ```
    rm -f ./bcnad #delete old file if exist
    tar zxvf  bcna_linux_amd64.tar.gz
    rm bcna_linux_amd64.tar.gz
    chmod +x bcnad
    ./bcnad version
    ```
    2.2.3. Move the new binary to your machine's PATH and overwrite the previous version
    ```
    sudo mv bcnad $(which bcnad)   #copy&paste don't replace anything
    ```
    If you know the exact destination you could also run: 
    ```
    sudo mv bcnad /usr/local/bin/ #or wherever you have it
    ```
3. Start the bcnad service
    ```
    sudo service bcnad start
    ```
4. Ensure that everything is OK by checking the logs 
    ```
    sudo journalctl -u bcnad -f
    ```

## Unattended (Cosmovisor) upgrade. Recommended.
Cosmovisor is an hypervisor that check the logs and replaces binaries when the upgrades happens. 
It run as service and replaces the `bcnad` service. A complete guide to config Cosmovisor with BitCanna chain is [here](https://github.com/BitCannaGlobal/bcna/blob/main/2.3.cosmovisor.md)

### Step 1. Setup Cosmovisor folder
> Attention! There  is a new official release of Cosmovisor ([check here v.1.2](https://github.com/cosmos/cosmos-sdk/releases/tag/cosmovisor%2Fv1.2.0))
> If you need instructions to upgrade check this [doc](https://github.com/BitCannaGlobal/bcna/blob/main/2.3.cosmovisor.md). 

1. Create a new directory to store the new binary:
```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/strangeBuddheads/bin
```

2. Compile the binary from the source (2.1) or download it (2.2).

    2.1. If you want to BUILD the binary from the source, detailed instructions can be found in this [DOC](https://github.com/BitCannaGlobal/bcna/blob/main/1.install-compile.md) of our GitHub:
    ```
    git clone https://github.com/BitCannaGlobal/bcna.git
    cd bcna && git checkout  v1.4.1
    make build
    ```
    This will produce a `bcnad` file in the `build` folder. Move to the NEW binary path:
    ```
    mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/strangeBuddheads/bin   #copy&paste don't replace anything
    ```
   *You can jump to step 3.*
    
    2.2. If you want to DOWNLOAD the compiled binary (instead of compiling by yourself):
    ```
    cd ~
    rm bcna_linux_amd64.tar.gz #delete old file if exist
    wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v1.4.1/bcna_linux_amd64.tar.gz
    ```
    2.2.1 Check the sha256sum for the downloaded file. 
    ```
    sha256sum bcna_linux_amd64.tar.gz
    ```
    It must return: `305ff854cc892e5bd4de6f2b3d8bd745525878244f78478be6c579d56316fe9d`

    2.2.2 Verify that the version is `1.4.1`
    ```
    rm -f ./bcnad #delete old file if exist
    tar zxvf  bcna_linux_amd64.tar.gz
    rm bcna_linux_amd64.tar.gz
    chmod +x bcnad
    ./bcnad version
    ```
    2.2.3 Move the newly downloaded/built binary to the upgrades directory.
    ```
    mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/strangeBuddheads/bin/
    ```
3. If you want to be sure that the proper binary is in the proper folder:
```
${HOME}/.bcna/cosmovisor/upgrades/strangeBuddheads/bin/bcnad version
    # --> should reply: 1.4.1
```


## EXTRA OPTIONAL STEP 1: <br>Install Cosmovisor v1.2
If you want to update Cosmovisor from v1.1 to v1.2 follow the next instructions, for a new installation and setup of BCNA params check this [doc](https://github.com/BitCannaGlobal/bcna/blob/main/2.3.cosmovisor.md).

> CAUTION! Cosmovisor v1.2 has a bug and it doesn't get the uppercase in the binary folder path

Temporal solution is: 
```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/strangebuddheads/bin
cp -R ${HOME}/.bcna/cosmovisor/upgrades/strangeBuddheads/ ${HOME}/.bcna/cosmovisor/upgrades/strangebuddheads/
```

```
wget https://github.com/cosmos/cosmos-sdk/releases/download/cosmovisor%2Fv1.2.0/cosmovisor-v1.2.0-linux-amd64.tar.gz
tar zxvf cosmovisor-v1.2.0-linux-amd64.tar.gz 
rm *.md && rm cosmovisor-v1.2.0-linux-amd64.tar.gz
sudo mv cosmovisor $(which cosmovisor)
```
> BUG description here: https://github.com/cosmos/cosmos-sdk/issues/12915


## EXTRA OPTIONAL STEP 2: <br>Review the service file if you are running Cosmovisor v1.2 

Cosmovisor v1.2 introduces several changes and improvements. 

* It's mandatory to execute all the commands with `run` prefixe (`cosmovisor run status`)
    That implies that you need to modify the service file if it doesn't include the `run` command in the Daemon Start command.
* Cosmovisor v1.2 includes a new option (env: ` `to delay the start of the new binary and make the automated backup. It prevents to start again without kill the previous daemon and make a bad automated backup (if is set), so we should configure this new feature to introduce 30 seconds of delay.

**Let's review this two tips:**


* With Cosmovisor service 

    1. Edit the file
    ```
    sudo nano /lib/systemd/system/cosmovisor.service
    ```
    Replace `cosmovisor start` by `cosmovisor run start` line and save it:
    ```
    ExecStart=/usr/local/bin/cosmovisor run start
    ```
    Add this options if you want an automated backup of chain before the upgrade takes action.
    ```
    Environment=UNSAFE_SKIP_BACKUP=false
    Environment=DAEMON_RESTART_DELAY=30s
    ```
    2. Reload service file and start Cosmovisor
    ```
    sudo systemctl daemon-reload
    sudo service cosmovisor restart 
    sudo journalctl -fu cosmovisor
    ```

## Simple test with Cosmovisor after the upgrade

If you want to know if Cosmovisor handles the right binary file exec:
* `sudo service cosmovisor status`
* `cosmovisor version`

And check the path of the binary file.    

###### tags: "bitcanna", "mainnet", "upgrade"
