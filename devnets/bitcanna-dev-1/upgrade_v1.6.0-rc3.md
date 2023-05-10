# Node operator instructions for "Vigorous-Grow"`v1.6.0-rc3` upgrade on `bitcanna-dev-1`
In this guide, you will find all the instructions required to upgrade a BitCanna node and help us test the latest software version on the DEVNET-1. 

We provide detailed instructions for **manual and automated upgrades**, as well as links to instructions for installing Cosmovisor or building the binaries from the source code. 

If you have any questions or find anything unclear when reading through this guide, don't hesitate to reach out to us on Discord. Also, if you have any suggestions for improvement, we welcome you to make a pull request!

Let's get started! 

### Governance proposal

The upgrade is scheduled to take place on **February 24nd around 16:00h CET / 15:00h UTC**
* **Halt-height 987.757**

https://testnet.ping.pub/bitcanna/gov/4

## Download the new binary 
1) Download the new binary or [compile it from the source code](#Build-the-binary-from-the-source-code)
```
cd ~
rm -f ./bcnad && rm -f ./bcna_linux_amd64.tar.gz #clean the previous downloads
wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v1.6.0-rc3/bcna_linux_amd64.tar.gz
```
2) Check the sha256sum. 
```
sha256sum bcna_linux_amd64.tar.gz
```
> It must return: `4ab89e9df3c340acc42b01d52167a676a6daa6749b70e3f0eee8f65808af8bc0`

3) Verify that the version is `v1.6.0-rc3`
```
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```

## Option 1. Manual (attended) upgrade.

This section of the guide shows how to perform a **manual** update of the binaries after a governance proposal has been approved for a chain upgrade.

Watch your logs when we aproach the halt-height and open an extra terminal to execute the commands. 
```
sudo journalctl -fu bcnad -o cat
```
1) Stop the bcnad service **after you see** `ERR UPGRADE "vigorous-grow-rc3" NEEDED at height:  987.757` in your logs.
```
sudo service bcnad stop
```
2) Move the new binary to your machine's PATH and overwrite the previous version.
```
sudo mv bcnad $(which bcnad)   #copy&paste don't replace anything
```
> If you know the exact destination you could also run: 
```
sudo mv bcnad /usr/local/bin/ #or wherever you have it
```
3) Start the bcnad service.
```
sudo service bcnad start
```

If you are familiar with the Cosmos environment, you can keep the daemon running while you are downloading or compiling the new binaries. Once the halt-height is reached, you can execute the upgrade in one command: 
```
sudo service bcnad stop && sudo mv bcnad $(which bcnad) && sudo service bcnad start
```
4) Verify that everything is running properly by inspecting the log files.
```
sudo journalctl -fu bcnad -o cat
```

## Option 2. Cosmovisor (unattended) upgrade. 
This section of the guide provides instructions on how to perform an **automated** upgrade of the binaries after a governance proposal has been approved for a chain update.

For detailed instructions on how to set up Cosmovisor for the first time, please follow [this link](https://github.com/BitCannaGlobal/bcna/blob/main/devnets/bitcanna-dev-1/cosmovisor_install.md).

This guide shows how to download the newest binary. If you want to build the binary from the source, please refer to [this link](https://github.com/BitCannaGlobal/bcna/blob/main/1.install-compile.md#option-2-compile-instructions) for detailed instructions.

### Setup Cosmovisor folder
Download the upgraded binary as mentioned in the beginning of this guide before continuing with this step.

1) Create a new directory designated for upgrades.
```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/vigorous-grow-rc3/bin
```
2) Move the newly downloaded binary to the directory.
```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/vigorous-grow-rc3/bin/
```
> If you build the binary from the code source, move it to the same folder.

3) If you want to know if Cosmovisor handles the correct binary file, execute this command after the upgrade:
```
sudo service cosmovisor status
```
And check the path of the binary file.

## Build the binary from the source code

For detailed instructions on how to build the binary from the source code, please refer to [this link](https://github.com/BitCannaGlobal/bcna/blob/main/1.install-compile.md#option-2-compile-instructions).
```
    git clone https://github.com/BitCannaGlobal/bcna.git
    cd bcna
    git checkout v1.6.0-rc3
    make build && make install 
```

###### tags: `upgrade` `upgrade_v1.6.0-rc3.md`
