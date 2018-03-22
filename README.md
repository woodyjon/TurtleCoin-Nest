# TurtleCoin Nest

The universal GUI wallet for TurtleCoin

![Logo](/turtlecoinnestlogo.png)

## Installation

[Windows](#windows) - [Mac](#mac) - [Linux](#linux)

### Windows

1. Go [here](https://github.com/woodyjon/TurtleCoin-Nest/releases) and download the latest release called **TurtleCoin-Nest-x.xx-Windows.zip**
2. Unzip the folder and launch **TurtleCoin-Nest.exe**. (Make sure you leave everything as is in the folder)

Important notes:

* Make sure *Walletd.exe* or *Turtlecoind.exe* are not running before you start *TurtleCoin-Nest*
* When you open a wallet in TurtleCoin Nest, you will see a black empty walletd window. You must keep it open. It will close automatically when you close your wallet. If it does not close automatically a few seconds after you close your wallet, you can close it manually.

### Mac

1. Go [here](https://github.com/woodyjon/TurtleCoin-Nest/releases) and download the latest release called **TurtleCoin-Nest-x.xx-Mac.zip**
2. Unzip the folder and launch **TurtleCoin-Nest**. (Make sure you leave everything as is in the folder. The file *walletd* should remain in the same folder as *TurtleCoin-Nest*)

Important notes:

* Make sure *Walletd* or *Turtlecoind* are not running before you start *TurtleCoin-Nest*
* If you encounter crashes, open the activity monitor (in your app > utilities), and force quit *walletd* (if it is running) before opening a wallet. (this bug is being worked on)

### Linux

1. Go [here](https://github.com/woodyjon/TurtleCoin-Nest/releases) and download the latest release called **TurtleCoin-Nest-x.xx-Linux.tar.gz**
2. extract it
`$ tar xvzf TurtleCoin-Nest-x.xx-Linux.tar.gz`
3. run **TurtleCoin-Nest.sh**. (Make sure you leave everything as is in the folder)

Important notes:

* Make sure *Walletd* or *Turtlecoind* are not running before you start *TurtleCoin-Nest*
* If you want the *copy address to clipboard* button to work, install *xclip* or *xsel* (on Debian/Ubuntu: `$ sudo apt install xclip`).
* If you encounter crashes, open an activity monitor (e.g. `$ htop`), and quit *walletd* (if it is running) before opening a wallet. (this bug is being worked on)

## Screenshots

![Main Screen](/Screenshots/MainScreen.png)

![Open Wallet](/Screenshots/OpenWallet.png)

## Donations

TRTLv3jzutiQwqHL3qFwsu5EVLWesxZr1AFQ4AuMR3SD56n3rkHDkwj79eKwvaiU1nYQWGydKoXM6fXyiiGKsPDnVCNXzNdusxx

## Build

(for developers only)

Install this binding: https://github.com/therecipe/qt and run `qtdeploy build desktop`
