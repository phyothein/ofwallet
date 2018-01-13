# OF钱包用户使用手册

福包，OFBANK的智能数字钱包，安全放心，简单好用，转账方便，私钥本地生成，安全保存，OF资产一目了然。内置OFBANK区块浏览器，便捷查询块高及交易信息。

OFBANK官网：www.ofbank.com

# 安装钱包
 
安装OFBank的钱包需要Go1.7以上和C编译器，下载或者通过Git clone 把钱包源码复制到本地电脑,进入源码文件夹进行钱包的编译安装

```
$ git clone https://github.com/ofproject/ofwallet.git
$ cd ofwallet
$ make ofbank   //钱包运行文件会被安装在 build/bin 下
```

# 运行钱包

```
$ cd build/bin
$ ./ofban console //启动钱包终端程序
```

# 钱包指令

| JS指令                     |  指令介绍  |
| :------------------:       | :----: |
| ofbank.register|   生成钱包地址（包括私钥文件）    |
| ofbank.registerWithMulKeys     |   生成多个钱包地址（包括多个私钥文件）  
| ofbank.getBalance    |  获取指定用户的余额   |
| ofbank.getBlockHeight    |  获取当前块高   |
| ofbank.getTransactionByHash    |  通过交易Hash 查询交易   |
| ofbank.getTransactionsByAddress    |  通过钱包地址查询该地址的所有交易来往|
| ofbank.sendTranToNode    |  发送交易   |
| ofbank.getBlock    |  获取指定区块信息   |
| ofbank.accounts   | 查看当前所有钱包地址   |

# 指令使用
- ofbank.register()
  - 参数:
     1. 钱包国家代码(String)
     2. 钱包密码(String)
  - 钱包存储在 build/bin/keystore   
```
ofbank.register("123")   
Passphrase:
Repeat passphrase:
"0x000006007b1ec0045ac3a76de32e17aab2849be2d55248574e"  //返回的地址
``` 
- ofbank.registerWithMulKeys()
  - 参数:
     1. 钱包国家代码(String)
     2. 钱包个数(Int)
     3. 钱包密码
  - 钱包存储在 build/bin/keystore   
```
ofbank.registerWithMulKeys("123",6)
Passphrase:
Repeat passphrase:
["0x000006007b6506f2f70833beab3ab7b78ebcd41c447b469112", "0x000006007bd4c7f728ee28c3839a6c016a0cd4a50d3b211bf6", "0x000006007be10db8cf11f9461852157292133b5dd2bcdd9f9f", "0x000006007b038ab5b51025a0e516865e5ace0bd184300c64e9", "0x000006007bb71cbe80cf28df3a7790249fd802478f306ed654", "0x000006007bf78dd6be600a97ec453e4e2283a6aff388c2acaa"]  //返回的多个地址
``` 
- ofbank.getBalance()
  - 参数:
     1. 钱包地址(String)
    
```
ofbank.getBalance("0x000006007b6506f2f70833beab3ab7b78ebcd41c447b469112")  
"0.00"  //返回的余额
``` 
- ofbank.getBlockHeight()
```
ofbank.getBlockHeight() 
"4056"  //返回的余额
``` 

- ofbank.sendTranToNode()
  - 参数(JSON 字符串):
     1. from(String)
     2. to (String)
     3. value (string)
     4. gas (int)
     5. gasPrice(int)
   - from地址需要先解锁
```
fbank.sendTranToNode({from:"0x000006007b1eb42c88ea5e3e45f3df16e37900fbb049aeba7f",to:ofbank.accounts[2],value:"5.056",gas:30000,gasPrice:190})

"0xf2cc289aff3d617e7c305b394e2f01958c6bddfddad2443c676eaf117bb61aac" //返回交易Hash
``` 

- ofbank.getTransactionByHash()
  - 参数:
     1. 交易Hash(String)
```
ofbank.getTransactionByHash("0xf2cc289aff3d617e7c305b394e2f01958c6bddfddad2443c676eaf117bb61aac")

{
  blockHash: "0xa613cd0d06b8e60bb16f28248f427b1796d0fd81c89938e75d6f3cdd6bd990a0",
  blockNumber: "0xfda",
  from: "0x000006007b1ec0045ac3a76de32e17aab2849be2d55248574e",
  gas: "0x7530",
  gasPrice: "0xbe",
  hash: "0xf2cc289aff3d617e7c305b394e2f01958c6bddfddad2443c676eaf117bb61aac",
  input: "0x",
  nonce: "0x0",
  r: "0xe2768390f2b7a0af5cf8ac0d204755890e202c652f42fb14dbc82b76f121893a",
  s: "0x61ea09fc6f5384e36019ca5342ee6f11ba6b7bc43e1895c0ec74512272dd8a7b",
  to: "0x000006007bd4c7f728ee28c3839a6c016a0cd4a50d3b211bf6",
  transactionIndex: "0x0",
  v: "0x1c",
  value: 5056000000000000000
}

```
- ofbank.getBlock()
  - 参数:
     1. 区块高度(int)
```
ofbank.getBlock(7)
{
  blockReward: "913.242",
  coinbase: "0x000006009cea48556921ac29f40329e07b791c9f1482d5ace5",
  difficulty: "0x20180",
  extraData: "0xe5b9b3e58e9f",
  gasLimit: 5278817,
  gasUsed: "0x0",
  hash: "0x238cdddae447655d5bffe65605a7a7a4e9d2b9ad516973e1e29da5cadfc27ede",
  miner: "0x000006009cea48556921ac29f40329e07b791c9f1482d5ace5",
  mixHash: "0xae5ead321b85590d484945b3d5b24cf9933f8c07c57e111b0c2ef2409ef73eb3",
  nonce: "0x5cfab1089f6816f7",
  number: "0x7",
  parentHash: "0xbc2ef48d7c3a2780132ed4f69751d62b94a8ae335b65c78d53753d70c1e8ef7d",
  receiptsRoot: "0x000100020003000400050006000700080009000a000b000c000d000e000f0000",
  sha3Uncles: "0x1dcc4de8dec75d7aab85b567b6ccd41ad312451b948a7413f0a142fd40d49347",
  size: "0x246",
  stateRoot: "0xecf65ef48889bd13dc3136e5d7fabbf1156731b243bcedb2e49628510c55678b",
  timeStamp: "0x5a571230",
  totalDifficulty: "0x5c406",
  transactions: [],
  transactionsRoot: "0x000100020003000400050006000700080009000a000b000c000d000e000f0000"
}
```
- ofbank.getBalance()
  - 参数:
     1. 钱包地址(String)
```
ofbank.getBalance(ofbank.accounts[0])

"2734.670681"  /／余额
```
- ofbank.accounts
 ```
 ofbank.accounts

 ["0x000006007b1ec0045ac3a76de32e17aab2849be2d55248574e", "0x000006007b6506f2f70833beab3ab7b78ebcd41c447b469112", "0x000006007bd4c7f728ee28c3839a6c016a0cd4a50d3b211bf6", "0x000006007be10db8cf11f9461852157292133b5dd2bcdd9f9f", "0x000006007b038ab5b51025a0e516865e5ace0bd184300c64e9", "0x000006007bb71cbe80cf28df3a7790249fd802478f306ed654", "0x000006007bf78dd6be600a97ec453e4e2283a6aff388c2acaa"]
 ```

 ofbank.getBlockHeight()
 
```
ofbank.getBlockHeight()
"4059"
 ```
