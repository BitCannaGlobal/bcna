# Please replace current `v1.4.2` /  `v1.4.3-patch` or `v1.4.4` by `v1.4.5`
## Key features
- SDK v0.45.10
- Tendermint v0.34.22
- IAVL v0.19.4 - fix AppHASHes (forks)
- IBC v.3.3.1
- No chain halt (consensus) required
- Apply it ASAP and monitoring your node.

> Stop your current daemon (bcnad or Cosmovisor) 
```
sudo service bcnad stop || sudo service cosmovisor stop
```

> Make a copy of the bd files:
```
cp -R $HOME/.bcna/data $HOME/.bcna/data_update
``` 

# Instructions -  **very important**
## 1_Previously to replace the binary or build from source you need to edit the file `app.toml`
you should to declare and put in `false` this var: `iavl-disable-fastnode = false`

Put this content in the main section, just before the `[telemetry]` section:
```
# IavlCacheSize set the size of the iavl tree cache. 
# Default cache size is 50mb.
iavl-cache-size = 781250

# IAVLDisableFastNode enables or disables the fast node feature of IAVL. 
# Default is true.
iavl-disable-fastnode = false  
###############################################################################
###                         Telemetry Configuration                         ###
###############################################################################

[telemetry]
```

## 2_Build from source or download it

### Option 1. Download
```
rm  -rf ./bcna_linux_amd64.tar.gz  # delete old versions, check also bcnad in this folder
wget https://github.com/BitCannaGlobal/bcna/releases/download/v1.4.5/bcna_linux_amd64.tar.gz
tar zxvf bcna_linux_amd64.tar.gz
./bcnad version
 >> result should be `1.4.5`

```
### Option 2. Compile from source
```
git clone https://github.com/BitCannaGlobal/bcna.git
cd bcna
git checkout v1.4.5
make build  ## move  build/bcnad manually to service path
make install ## it install in $GOBIN
```

## 3_Replace any previous version with this `v1.4.5` 

### For Cosmovisor:
```
mv ./bcnad  ~/.bcna/cosmovisor/current/bin/
sudo service cosmovisor restart
sudo journalctl -fu cosmovisor #check the logs
```

### For BCNAD daemon:
```
sudo mv ./bcnad $(which bcnad)
sudo service bcnad restart
sudo journalctl -fu bcnad #check the logs
```
