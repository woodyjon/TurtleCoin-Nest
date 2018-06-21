Functional tests to be done on windows, mac and linux before merging to master and deploying a release.

|  | Action | Result
|---|---|---
|  <ul><li>[ ] </li></ul> | Open wallet by clicking on the "Open" button | busy indicator for a few seconds then wallet opens and display balance, etc..., sync is ok
|  <ul><li>[ ] </li></ul> | Open wallet by typing return key while in password field | busy indicator for a few seconds then wallet opens and display balance, etc..., sync is ok
|  <ul><li>[ ] </li></ul> | Click on button to backup keys | backup keys are displayed (check the window is large enough)
|  <ul><li>[ ] </li></ul> | Click on button to copy keys to clipboard | Keys are copied to clipboard
|  <ul><li>[ ] </li></ul> | Click create wallet | dialog to confirm the password
|  <ul><li>[ ] </li></ul> | Click create wallet and enter wrong password confirmation | error
|  <ul><li>[ ] </li></ul> | Create wallet | busy indicator for a few seconds, then the wallet should be created in app folder and wallet open and display backup keys
|  <ul><li>[ ] </li></ul> | Create wallet with name already exist in folder | error: same filename already exist
|  <ul><li>[ ] </li></ul> | Open wallet -> Close wallet | Wallet closes
|  <ul><li>[ ] </li></ul> | Import wallet from keys | dialog to confirm the password
|  <ul><li>[ ] </li></ul> | Import wallet from keys and enter wrong password confirmation | error
|  <ul><li>[ ] </li></ul> | Import wallet from keys | wallet created in app folder. Check backup keys, address, balance.
|  <ul><li>[ ] </li></ul> | Open existing wallet -> close Nest -> reopen Nest | path to previously opened wallet should appear. Test opening it again without changing path. 
|  <ul><li>[ ] </li></ul> | Button copy address | Pop up "Copied!" and address should be copied to clipboard.
|  <ul><li>[ ] </li></ul> | Button copy tr. id | Pop up "Copied!" and tr. id should be copied to clipboard.
|  <ul><li>[ ] </li></ul> | Button explore transaction | Trtl explorer opens in default browser with the right transacton.
|  <ul><li>[ ] </li></ul> | Receive trtl | Balance updates automatically and new confirmed transaction added to list previous transactions
|  <ul><li>[ ] </li></ul> | Send TRTL, address diff. than 99 chars | error: address invalid
|  <ul><li>[ ] </li></ul> | Send TRTL, address not start by TRTL | error: address invalid
|  <ul><li>[ ] </li></ul> | Send TRTL, amount is 0 or less | error: amount should be greater than 0
|  <ul><li>[ ] </li></ul> | Send TRTL, amount is not a number | error: amount invalid
|  <ul><li>[ ] </li></ul> | Send TRTL, amount + fee is more than available balance | error: available balance is insufficient
|  <ul><li>[ ] </li></ul> | Send TRTL, small valid amount | dialog for confirming transfer. After confirmation, popup TRTL sent. amount is received on the other end
|  <ul><li>[ ] </li></ul> | Send TRTL, invalid payment id | error: wrong payment id format
|  <ul><li>[ ] </li></ul> | Send TRTL, valid payment id | popup TRTL sent. amount is received on the other end with correct payment id
|  <ul><li>[ ] </li></ul> | select the option local blockchain, restart | local blockchain should still be selected
|  <ul><li>[ ] </li></ul> | Test using the wallet with local blockchain selected | 
|  <ul><li>[ ] </li></ul> | Test using the wallet with remote node selected |
|  <ul><li>[ ] </li></ul> | Reduce window height and test scrolling |
|  <ul><li>[ ] </li></ul> | Check wallet filename is displayed above address |
|  <ul><li>[ ] </li></ul> | Click Settings | Settings screen appears
|  <ul><li>[ ] </li></ul> | Close Settings | Settings screen disappears
|  <ul><li>[ ] </li></ul> | Settings: Enable display in USD -> Close Nest -> Reopen | Setting is still enabled
|  <ul><li>[ ] </li></ul> | Display in USD enabled -> Open a wallet| Balance is displayed in USD (under TRTL balance)
|  <ul><li>[ ] </li></ul> | Display in USD enabled -> type an amount to be transfered| Value is displayed in USD
|  <ul><li>[ ] </li></ul> | Settings: change remote node address and port -> click save -> close and reopen Nest | new address in remote node radio button and in settings
|  <ul><li>[ ] </li></ul> | Settings: click remote node reset to default -> close and reopen Nest | default address in remote node radio button and in settings
|  <ul><li>[ ] </li></ul> | Win only: open a wallet | no black window (walletd) should open next to the nest window
|  <ul><li>[ ] </li></ul> | Win & Mac only: check executable file | its icon is the nest icon
|  <ul><li>[ ] </li></ul> | Win & Mac only: launch nest | icon in taskbar should be nest icon
|  <ul><li>[ ] </li></ul> | Launch on a small screen or resolution (800 x 600) | possibility to click full screen and to scroll in the full window
|  <ul><li>[ ] </li></ul> | Click button to send full balance | Full balance - fee is automatically displayed in transfer amount field
|  <ul><li>[ ] </li></ul> | Transfer with different fee | Check transfer was made with proper fee
|  <ul><li>[ ] </li></ul> | Modify transfer fee and mixin, then click "default values" | Default values should be displayed in corresponding field
|  <ul><li>[ ] </li></ul> | List transactions | 20 transactions are displayed
|  <ul><li>[ ] </li></ul> | Switch to display "all transactions" then switch back | Proper number of transactions are displayed
|  <ul><li>[ ] </li></ul> | Start application and check log file | version of Nest in "application started" log should be correct
|  <ul><li>[ ] </li></ul> | Connect to wrong remote node | error message about connection problem
|  <ul><li>[ ] </li></ul> | Click "?" to open Info screen from main screen | info dialog opens, check version
|  <ul><li>[ ] </li></ul> | Click "?" to open Info screen from wallet screen | info dialog opens
|  <ul><li>[ ] </li></ul> | Open info dialog, click link to chat | browser opens on chat url
|  <ul><li>[ ] </li></ul> | Open info dialog, click copy donation address | donation address copied to clipboard
