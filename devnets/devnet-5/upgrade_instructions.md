# Upgrade instructions for `wakeandbake46.6 v2.0.1-rc3`
## Governance proposal: 
* halt-height `1.XXX.420` 
* on December 22th 15:00h CET - 14:00h UTC

https://testnet.ping.pub/bitcanna/gov/11

## Download the new binary 
1) Download the binary [or compile it from the source](#If-you-want-to-build-from-the-source)
```
cd ~
rm -f ./bcnad && rm -f ./bcna_linux_amd64.tar.gz # clean the previous downloads
wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v2.0.1-rc3/bcna_linux_amd64.tar.gz
```
2) Check the sha256sum. 
```
sha256sum bcna_linux_amd64.tar.gz
```
> It must return: `db17d8b37598bb1a575dfccf35b628744da2edf767629cbcc1ee1597cee432a2`

3) Verify that the version is `v2.0.1-rc3`
```
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```
> You can continue with Manual instructions or Cosmovisor ones.

## Option 1. Attended (manual) upgrade.

This section of the guide shows how to perform a **manual** update of the binaries after a governance proposal has been approved for a chain update.
1) Stop your bcnad service **after you see this** in your logs `ERR UPGRADE "wakeandbake46.6" NEEDED at height: 1.XXX.420`
```
sudo service bcnad stop
```
2) Move the new binary to your machine's PATH and overwrite the previous version
```
sudo mv bcnad $(which bcnad)   #copy&paste don't replace anything
```
> If you know the exact destination you could also run: 
```
sudo mv bcnad /usr/local/bin/ #or wherever you have it
```
3) Start the bcnad service
```
sudo service bcnad start
```

Of course, if you are familiar with the Cosmos environment, you can keep the daemon running while you are compiling/downloading and later make the upgrade in one command: 
```
sudo service bcnad stop && sudo mv bcnad $(which bcnad) && sudo service bcnad start
```
4) Ensure that everything is OK by checking the logs 
```
sudo journalctl -fu bcnad -o cat
```

## Option 2. Unattended (Cosmovisor) upgrade. 
This section of the guide shows how to perform a **automated** upgrade of the binaries after a governance proposal has been approved for a chain update.

For detailed instructions about setting up Cosmovisor from scratch, check this [guide](https://hackmd.io/jsJCqEyJSHKVOFKjScn3rw).

This guide shows how to download the binary. If you want to build the binary from the source, detailed instructions can be found in the [README](https://github.com/BitCannaGlobal/bcna/blob/main/README.md) of our GitHub (`git checkout v2.0.1-rc3`)

### Setup Cosmovisor folder
Binary should be download previously

1) Create new directory
```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/wakeandbake46.7/bin
```
2) Move the newly downloaded binary to the upgrades directory.
```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/wakeandbake46.7/bin/
```
> If you build the binary from the code source, move it to the same folder

3) If you want to know if Cosmovisor handles the correct binary file, exec:
```
sudo service cosmovisor status
```
And check the path of the binary file.

## If you want to build from the source

 If you want to build the binary from the source, detailed instructions can be found in the [README](https://github.com/BitCannaGlobal/bcna/blob/main/README.md) of our GitHub:
```
    git clone https://github.com/BitCannaGlobal/bcna.git
    cd bcna
    git checkout v2.0.1-rc3
    make build && make install 
```
---------


# Upgrade instructions for `wakeandbake46.6 v2.0.1-rc2`

## Governance proposal: 
* halt-height `1.032.049` 
* on November 28th 16:00h CET - 15:00h UTC

https://testnet.ping.pub/bitcanna/gov/10

## Download the new binary 
1) Download the binary [or compile it from the source](#If-you-want-to-build-from-the-source)
```
cd ~
rm -f ./bcnad && rm -f ./bcna_linux_amd64.tar.gz # clean the previous downloads
wget -nc https://github.com/BitCannaGlobal/bcna/releases/download/v2.0.1-rc2/bcna_linux_amd64.tar.gz
```
2) Check the sha256sum. 
```
sha256sum bcna_linux_amd64.tar.gz
```
> It must return: `a424ad37b301578370bce58134c78891798c9c0b519ec09954054184a2a868e1`

3) Verify that the version is `v2.0.1-rc2`
```
tar zxvf  bcna_linux_amd64.tar.gz
rm bcna_linux_amd64.tar.gz
./bcnad version
```
> You can continue with Manual instructions or Cosmovisor ones.

## Option 1. Attended (manual) upgrade.

This section of the guide shows how to perform a **manual** update of the binaries after a governance proposal has been approved for a chain update.
1) Stop your bcnad service **after you see this** in your logs `ERR UPGRADE "wakeandbake46.6" NEEDED at height: 1.032.049`
```
sudo service bcnad stop
```
2) Move the new binary to your machine's PATH and overwrite the previous version
```
sudo mv bcnad $(which bcnad)   #copy&paste don't replace anything
```
> If you know the exact destination you could also run: 
```
sudo mv bcnad /usr/local/bin/ #or wherever you have it
```
3) Start the bcnad service
```
sudo service bcnad start
```

Of course, if you are familiar with the Cosmos environment, you can keep the daemon running while you are compiling/downloading and later make the upgrade in one command: 
```
sudo service bcnad stop && sudo mv bcnad $(which bcnad) && sudo service bcnad start
```
4) Ensure that everything is OK by checking the logs 
```
sudo journalctl -fu bcnad -o cat
```

## Option 2. Unattended (Cosmovisor) upgrade. 
This section of the guide shows how to perform a **automated** upgrade of the binaries after a governance proposal has been approved for a chain update.

For detailed instructions about setting up Cosmovisor from scratch, check this [guide](https://hackmd.io/jsJCqEyJSHKVOFKjScn3rw).

This guide shows how to download the binary. If you want to build the binary from the source, detailed instructions can be found in the [README](https://github.com/BitCannaGlobal/bcna/blob/main/README.md) of our GitHub (`git checkout v2.0.1-rc2`)

### Setup Cosmovisor folder
Binary should be download previously

1) Create new directory
```
mkdir -p ${HOME}/.bcna/cosmovisor/upgrades/wakeandbake46.6/bin
```
2) Move the newly downloaded binary to the upgrades directory.
```
mv ./bcnad ${HOME}/.bcna/cosmovisor/upgrades/wakeandbake46.6/bin/
```
> If you build the binary from the code source, move it to the same folder

3) If you want to know if Cosmovisor handles the correct binary file, exec:
```
sudo service cosmovisor status
```
And check the path of the binary file.

## If you want to build from the source

 If you want to build the binary from the source, detailed instructions can be found in the [README](https://github.com/BitCannaGlobal/bcna/blob/main/README.md) of our GitHub:
```
    git clone https://github.com/BitCannaGlobal/bcna.git
    cd bcna
    git checkout v2.0.1-rc2
    make build && make install 
```
