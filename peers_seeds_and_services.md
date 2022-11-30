# BitCanna peers, seed & services
For the **Chain Registry** from Cosmos [GitHub](https://github.com/cosmos/chain-registry/) we'd like to gather some endpoints and persistent peers from our most valuable users.

We'd like to ask you to provide us your ***endpoints*** and/or ***persistent_peers*** that we can use. It is greatly appreciated!

# Persistent_peers

## Stakely
3cdfe02efd4432280707d2949e064a9d8db412b3@178.62.98.158:26656 
d806bb39349751c142627a547c23c586a787ef26@138.68.78.210:26656

## BloClick
be87c9abf1c54e1cc2f37e68d21fcd61679abb4c@65.21.196.90:46656 

## BitCanna
7c00beb4956bc40cd33ced6e2c2ffe07d4fa32e7@95.216.242.82:36656
21484408a7bcf0134689ddf52a7d9c8299cb65b5@176.9.139.74:36656

## Paranormal Brothers
df99de6cec9152c517990317b340b8b9a307493c@193.34.144.156:26656

## Tuzem
ec283da64f69f8c0dee4671021edc419bbeb4034@157.90.179.34:56656

## Polkachu
282fe52546750ca86e92e3f567cbdb2897694b5b@207.244.225.205:26656

## Nodejumper 🚀
45589e6147e36dda9e429668484d7614fb25b142@bitcanna.nodejumper.io:27656

## Inter Blockchain Services (former 3Tekos)
ec12bf44fd3c64db457f45f7d0111735c559a37d@185.218.126.71:26657
79d9c3aa19f61e06c33c54d80d0cda4fa535b28a@161.97.156.216:26657

## STAVR
0bf629f4e055af47f7c35bb444cb9013d18b9941@141.95.124.151:21326

## Panthea EU
0a658df9d9fab096983a12e6f878e87281a15ce6@bitcanna-peer.panthea.eu:27656

## AlxVoy
803fc66e3bd7b724921ef9c40636067f36e880c6@65.108.199.222:26357

# Seeds
## BitCanna (reseted everyday)
d6aa4c9f3ccecb0cc52109a95962b4618d69dd3f@seed1.bitcanna.io:26656
23671067d0fd40aec523290585c7d8e91034a771@seed2.bitcanna.io:26656

## Panthea EU
f0e6c86d769bf5c52f78e01864091690e731643f@bitcanna-seed.panthea.eu:37656

# StateSync Servers and instructions:
## BitCanna oficial:
  * https://github.com/BitCannaGlobal/bcna/blob/main/2.1.statesync.md

## Polkachu
  * https://polkachu.com/state_sync/bitcanna

## Panthea EU StateSync Server
  * https://bitcanna-rpc.panthea.eu:443

## Inter Blockchain Services (former 3Tekos) StateSync Server
  * https://bcna-rpc.ibs.team:443

## mintthemoon
  * https://docs.mintthemoon.xyz/bitcanna/statesync

## STAVR StateSync
  * https://github.com/obajay/StateSync-snapshots/tree/main/Bitcanna#statesync

## AlxVoy StateSync
  * https://github.com/Voynitskiy/Voynitskiy/blob/main/mainnet/BitCanna/README.md#state-sync

# Info, doc and other services from BitCanna validators & partners (feel free to make a PR)
### BitCanna Docs
* https://docs.bitcanna.io

### Stakely Faucet
* https://stakely.io/faucet/bitcanna-bcna

### Paranormal Brothers: current snapshot, RPC, seed, peer service, validators monitoring and alerting tool. 
* https://bc.paranorm.pro/

### Polkachu: snapshot service info
* https://polkachu.com/tendermint_snapshots/bitcanna

### Nodejumper 🚀: on-chain analytics, snapshot & state sync services, installation scripts (manual/automation), cheat sheet and more.
* https://nodejumper.io/bitcanna

### DappLocker: Stats service
* https://dapplooker.com/dapp/bitcanna-mainnet-120038

### mintthemoon
* Resource list: https://docs.mintthemoon.xyz/bitcanna/resources
* Node setup guide: https://docs.mintthemoon.xyz/bitcanna/node-guide
* Validator setup guide: https://docs.mintthemoon.xyz/bitcanna/validator-guide

### STAVR : snapshot and state sync services, installation scripts (manual/automation)
* https://github.com/obajay/nodes-Guides/blob/main/Bitcanna/README.md#1-auto_install-script

### AlxVoy : snapshot and state sync services, guide, RPC, API, peer
* https://github.com/Voynitskiy/Voynitskiy/blob/main/mainnet/BitCanna/README.md


# RPCs, LCD, GRPC and other endpoints/services 
```
  "peers": {
    "seeds": [
      {
        "id": "d6aa4c9f3ccecb0cc52109a95962b4618d69dd3f",
        "address": "seed1.bitcanna.io: 26656",
        "provider": "bitcanna"
      },
      {
        "id": "23671067d0fd40aec523290585c7d8e91034a771",
        "address": "seed2.bitcanna.io: 26656",
        "provider": "bitcanna"
      },
      {
        "id": "f0e6c86d769bf5c52f78e01864091690e731643f",
        "address": "bitcanna-seed.panthea.eu:37656",
        "provider": "Panthea EU"
      }
    ],
    "persistent_peers": [
      {
        "id": "21484408a7bcf0134689ddf52a7d9c8299cb65b5",
        "address": "176.9.139.74:36656",
        "provider": "BitCanna"
      },
      {
        "id": "a4c1e46441164c350f721cf142d52c136215e05c",
        "address": "135.181.176.55:36656",
        "provider": "BitCanna"
      },
      {
        "id": "3cdfe02efd4432280707d2949e064a9d8db412b3",
        "address": "178.62.98.158:26656",
        "provider": "Stakely"
      },
      {
        "id": "d806bb39349751c142627a547c23c586a787ef26",
        "address": "138.68.78.210:26656",
        "provider": "Stakely"
      },
      {
        "id": "df99de6cec9152c517990317b340b8b9a307493c",
        "address": "193.34.144.156:26656",
        "provider": "ParanormalBrothers"
      },
      {
        "id": "0a658df9d9fab096983a12e6f878e87281a15ce6",
        "address": "bitcanna-peer.panthea.eu:27656",
        "provider": "Panthea EU"
      },
      {
        "id": "803fc66e3bd7b724921ef9c40636067f36e880c6",
        "address": "65.108.199.222:26357",
        "provider": "AlxVoy"
      }
    ]
  },
  "apis": {
    "rpc": [
      {
        "address": "https://rpc.bitcanna.io/",
        "provider": "bitcanna"
      },
      {
        "address": "http://bcna.paranorm.pro/",
        "provider": "ParanormalBrothers"
      },
      {
        "address": "https://bcna-rpc.ibs.team/",
        "provider": "Inter Blockchain Services (former 3Tekos)"
      },      
      {
        "address": "https://bitcanna-rpc.panthea.eu",
        "provider": "Panthea EU"
      }
      {
        "address": "https://rpc.bitcanna.sgtstake.com/",
        "provider": "SGTstake"
      },
      {
        "address": "https://rpc-bitcanna.mintthemoon.xyz",
        "provider": "mintthemoon"
      },
      {
        "address": "https://bitcanna.rpc.m.anode.team",
        "provider": "AlxVoy"
      }
    ],
    "grpc": [
      {
        "address": "https://grpc.bitcanna.io",
        "provider": "bitcanna"
      }
    ],
    "rest": [
      {
        "address": "https://lcd.bitcanna.io",
        "provider": "bitcanna"
      },
      {
        "address": "https://bitcanna-api.panthea.eu",
        "provider": "Panthea EU"
      }
      {
        "address": "https://api.bitcanna.sgtstake.com/",
        "provider": "SGTstake"
      },
      {
        "address": "https://lcd-bitcanna.mintthemoon.xyz",
        "provider": "mintthemoon"
      },
      {
        "address": "https://bcna-api.ibs.team/",
        "provider": "Inter Blockchain Services (former 3Tekos)"
      },
      {
        "address": "https://bitcanna.api.m.anode.team",
        "provider": "AlxVoy"
      }
    ]
  }
  ```
