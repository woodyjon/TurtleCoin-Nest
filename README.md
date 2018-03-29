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
* When you open a wallet in TurtleCoin Nest, you will see a black empty *walletd* window. You must keep it open. It will close automatically when you close your wallet. If it does not close automatically a few seconds after you close your wallet, you can close it manually.

### Mac

1. Go [here](https://github.com/woodyjon/TurtleCoin-Nest/releases) and download the latest release called **TurtleCoin-Nest-x.xx-Mac.dmg**.
2. Double click the .dmg and drag the application **TurtleCoin-Nest** into /Applications or any other folder.
3. Launch the application.

Important notes:

* The wallets you create or generate will be saved to your home folder. You can keep them there or move them wherever you want.
* Make sure *Walletd* or *Turtlecoind* are not running before you start *TurtleCoin-Nest*.
* If you encounter crashes, open the activity monitor (in your app > utilities), and force quit *walletd* (if it is running) before opening a wallet. (this bug is being worked on).
* The log files will be saved in ~/Library/Application Support/TurtleCoin-Nest/.

### Linux

1. Go [here](https://github.com/woodyjon/TurtleCoin-Nest/releases) and download the latest release called **TurtleCoin-Nest-x.xx-Linux.tar.gz**
2. extract it
`$ tar xvzf TurtleCoin-Nest-x.xx-Linux.tar.gz`
3. run **TurtleCoin-Nest.sh**. (Make sure you leave everything as is in the folder)

Important notes:

* Make sure *Walletd* or *Turtlecoind* are not running before you start *TurtleCoin-Nest*
* If you want the *copy address to clipboard* button to work, install *xclip* or *xsel* (on Debian/Ubuntu: `$ sudo apt install xclip`).
* If you encounter crashes, open an activity monitor (e.g. `$ htop`), and quit *walletd* (if it is running) before opening a wallet. (this bug is being worked on)

## Upgrade

Just download the new release and follow the same steps as [Installation](#installation). Just make sure you don't delete your wallet files in the old folder (.wallet files) and you copy them or move them to a new folder.

## Screenshots

![Main Screen](/Screenshots/MainScreen.png)

![Open Wallet](/Screenshots/OpenWallet.png)

## Donations

TRTLv3jzutiQwqHL3qFwsu5EVLWesxZr1AFQ4AuMR3SD56n3rkHDkwj79eKwvaiU1nYQWGydKoXM6fXyiiGKsPDnVCNXzNdusxx

## Build

(for developers only)

Install this binding: https://github.com/therecipe/qt and run `qtdeploy build desktop`
