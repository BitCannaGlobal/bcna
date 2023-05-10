# Cosmovisor instructions for `bitcanna-dev-1`

In this guide, you will find step-by-step instructions for downloading, compiling, and installing Cosmovisor on your BitCanna node on the DEV-NET. We have provided the simplest methods for installation, so you can quickly and easily get Cosmovisor running on your node.

If you have any questions or find anything unclear when reading through this guide, don’t hesitate to reach out to us on Discord. Also, if you have any suggestions for improvement, we welcome you to make a pull request!

Let’s get started!

## About Cosmovisor
Cosmovisor is a small process manager for Cosmos SDK application binaries that monitors the governance module for incoming chain upgrade proposals. If it sees a proposal that gets approved, it stops the current binary, switch from the old binary to the new one, and finally restarts the node with the new binary.

This guide will explain how to install Cosmovisor and prepare for a future chain upgrade. A full guide about Cosmovisor can be found [here](https://github.com/cosmos/cosmos-sdk/tree/main/tools/cosmovisor).

## Step 1. Download or compile Cosmovisor

The easiest option to install Cosmovisor is by downloading the pre-compiled binary. Alternatively, users with more advanced technical knowledge can install Go on their system and compile Cosmovisor from the source. For a full guide to install Go, please refer to [this link](https://github.com/BitCannaGlobal/bcna/blob/main/1.install-compile.md#option-2-compile-instructions).

**Option 1:** Download the pre-compiled Cosmovisor binary:

```
cd ~
wget https://github.com/BitCannaGlobal/bcna/releases/download/v1.5.3/cosmovisor
```
Make Cosmovisor executable and move it to the bin directory.
```
chmod +x ./cosmovisor
sudo mv cosmovisor /usr/local/bin/
```
**Option 2:** Download the source code and compile Cosmovisor using Go: 

`go install cosmossdk.io/tools/cosmovisor/cmd/cosmovisor@latest`

## Step 2. Setup Cosmovisor
> Version v1.2 of Cosmovisor and higher now include a new command that will automatically create the required folder structure, reducing the number of steps required to complete the task.

### 1.) Download the genesis binary of BitCanna devnet-1.

For Cosmovisor to initiate the required folders automatically, the genesis version of the `bcnad` binary needs to be present in your home directory.

Download `bcnad` version `v1.5.3`.
```
cd ~
rm -f bcnad #deletes previously downloaded binary if it exists to avoid version mixing
wget https://github.com/BitCannaGlobal/bcna/releases/download/v1.5.3/bcna_linux_amd64.tar.gz
```
Check the sha256sum. 
```
sha256sum bcna_linux_amd64.tar.gz
```
> It must return: `8a43bdbea31c299db2ca849f232189374286c2168264072358d48c2a6f6aa2da`

Verify that you have the correct version.
```
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```
> It must return: `v1.5.3`

### 2.) Initiate the Cosmovisor folders.

Edit your `.profile` file to add the required environment variables for Cosmovisor to start.
```
cd ~
nano .profile
```
Add these lines to the bottom of the file:
```
export DAEMON_NAME=bcnad
export DAEMON_RESTART_AFTER_UPGRADE=true
export DAEMON_HOME=${HOME}/.bcna
export DAEMON_RESTART_DELAY=30s
export UNSAFE_SKIP_BACKUP=true
export DAEMON_LOG_BUFFER_SIZE=512

#add this to continue to use bcnad for commands, this is optional
PATH="${HOME}/.bcna/cosmovisor/current/bin:$PATH" 
```
Reload the configuration of your `.profile` file.
```
source .profile
```
Start the initial Cosmovisor configuration.
```
cosmovisor init ./bcnad
```
The output should contain lines similar to these:
```
11:46AM INF checking on the genesis/bin directory module=cosmovisor
11:46AM INF creating directory (and any parents): "/Users/test/.bcna/cosmovisor/genesis/bin" module=cosmovisor
11:46AM INF checking on the genesis/bin executable module=cosmovisor
11:46AM INF copying executable into place: "/Users/test/.bcna/cosmovisor/genesis/bin/bcnad" module=cosmovisor
11:46AM INF making sure "/Users/test/.bcna/cosmovisor/genesis/bin/bcnad" is executable module=cosmovisor
11:46AM INF checking on the current symlink and creating it if needed module=cosmovisor
11:46AM INF the current symlink points to: "/Users/t/.bcna/cosmovisor/genesis/bin/bcnad" module=cosmovisor
```
### 3.) Create designated directories for the current version and next upgrade.

```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/vigorous-grow/bin
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/vigorous-grow-rc3/bin
```
### 4.) Download the current and updated `bcnad` binaries and move it to the designated directory.

This guide explains how to download the pre-compiled binary. If you want to build the binary from the source, please refer to [this link](https://github.com/BitCannaGlobal/bcna/blob/main/1.install-compile.md#option-2-compile-instructions)

### 4.1) Download the current `bcnad` version `v1.6.0-rc2`.
> Alternatively, you could also copy your currently running version from the bin directory.
```
cd ~
rm -f bcnad #remove the previous downloads
wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.0-rc2/bcna_linux_amd64.tar.gz
```
Check the sha256sum.
```
sha256sum ./bcnad
```
> It must return: `1df53fc3e0f7d7d5ee0e6e5368634ebc5994f5f43ac05118aad65502bf26e723`

Verify that you have the correct version.
```
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```
> It must return: **`v1.6.0-rc2`**

Move the newly built binary to the designated upgrade directory.
> If you build the binary from the source, move it to the same folder.
```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/vigorous-grow/bin
```
### 4.2) Download the upgrade `bcnad` version `v1.6.0-rc3`.
```
cd ~
rm -f bcnad #remove the previous downloads
wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.0-rc3/bcna_linux_amd64.tar.gz
```
Check the sha256sum.
```
sha256sum ./bcnad
```
> It must return: `4ab89e9df3c340acc42b01d52167a676a6daa6749b70e3f0eee8f65808af8bc0`

Verify that you have the correct version.
```
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```
> It must return: **`v1.6.0-rc3`**

Move the newly built binary to the designated upgrade directory.
> If you build the binary from the source, move it to the same folder.
```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/vigorous-grow-rc3/bin
```
### 5.) Setup Cosmovisor's current version link.

Next step is to let Cosmovisor know in which directory the current version of `bcnad` software is located.

This command creates a symbolic link, pointing `current` to the currently active directory.
```
ln -sfn  ${HOME}/.bcna/cosmovisor/upgrades/vigorous-grow ${HOME}/.bcna/cosmovisor/current
```
### 6.) To see if everything is OK, run:

```
ls .bcna/cosmovisor/ -lh
```
The output should look like this:
```
total 8.0K
lrwxrwxrwx 1 user user   35 Feb 23 20:16 current -> /home/user/.bcna/cosmovisor/upgrades/vigorous-grow
drwxrwxr-x 3 user user 4.0K Feb 23 20:09 genesis
drwxrwxr-x 4 user user 4.0K Feb 23 20:15 upgrades
```
## Step 3. Finish the installation and start Cosmovisor

### 1.) Create a systemd servicefile for Cosmovisor.

Simply copy and paste everything to create the service unit.
> Optionally enable backup, set the folder location, and name the folder.
```
echo "[Unit]
Description=Cosmovisor BitCanna Service
After=network-online.target
[Service]
User=${USER}
Environment=DAEMON_NAME=bcnad
Environment=DAEMON_RESTART_AFTER_UPGRADE=true
Environment=DAEMON_HOME=${HOME}/.bcna
Environment=UNSAFE_SKIP_BACKUP=true
Environment=DAEMON_RESTART_DELAY=30s
Environment=DAEMON_LOG_BUFFER_SIZE=512
#Optional export DAEMON_DATA_BACKUP_DIR=${HOME}/your_chain_backup_folder
ExecStart=$(which cosmovisor) run start
Restart=always
RestartSec=3
LimitNOFILE=4096
[Install]
WantedBy=multi-user.target
" >cosmovisor.service
```
### 2.) Change `bcnad` service for `cosmovisor` service.
> You can skip the 3rd line if you're doing a clean installation and the `bcnad`service doesn't exist.
```
sudo mv cosmovisor.service /lib/systemd/system/
sudo systemctl daemon-reload
sudo systemctl stop bcnad.service && sudo systemctl disable bcnad.service 
sudo systemctl enable cosmovisor.service && sudo systemctl start cosmovisor.service
```
### 3.) Check the logs.** (ctrl + C to stop).
```
sudo journalctl -fu cosmovisor -o cat
```
> You can speed up the synchronization using a StateSync Server or a snapshot file.

### 4.) Now let's try Cosmovisor.
If everything is OK, Cosmovisor will take control of the binaries. 
Instead of `bcnad` you must use `cosmosvisor run` in your service related commands. 

Show Cosmovisor's version.
```
cosmovisor run version
``` 
This will be **`v1.6.0-rc2` before** the upgrade and **`v1.6.0-rc3` after** the upgrade.

The output should look like this:
 ```
20:27PM INF running app args=["version"] module=cosmovisor path=/home/user/.bcna/cosmovisor/upgrades/vigorous-grow
1.6.0-rc2
```
* Show BitCanna version: `bcnad version` Must show the same version as above
* Show Cosmovisor sync info and status: `cosmovisor run status` 

## Reminder
In the future, you must use the `cosmovisor` command instead of the `bcnad` command if you want to perform **service related** commands.

For example: 
* Start the service: `sudo service cosmovisor start`
* Stop the service: `sudo service cosmovisor stop`
* Restart the service: `sudo service cosmovisor restart`
* Check the logs: `sudo journalctl -u cosmovisor -f`
