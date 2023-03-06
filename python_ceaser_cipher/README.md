# Caesar Cipher and Leet Speak Converter

This program encrypts and decrypts text using the Caesar Cipher, and can also convert text to Leet (1337) speak. The Caesar Cipher is a simple encryption technique that shifts each letter of a message by a fixed number of positions down the alphabet.

## Installation

1. Clone the repository to your local machine.
2. Ensure that Python 3 is installed on your system.

## Usage

1. Navigate to the directory containing the program files.
2. Open a command prompt or terminal window in that directory.
3. Run the program by typing `python caesar.py` and pressing Enter.
4. Follow the prompts to select a mode (encrypt or decrypt) and enter text to be processed.

## Functions

### caesar_encrypt(plaintext, shift)

Encrypts a plaintext message using the Caesar Cipher, shifting each letter by a given number of positions down the alphabet.

#### Arguments

* `plaintext` - The text to be encrypted
* `shift` - The number of positions to shift each letter down the alphabet

### caesar_decrypt(ciphertext, shift)

Decrypts a ciphertext message using the Caesar Cipher, shifting each letter by a given number of positions up the alphabet.

#### Arguments

* `ciphertext` - The text to be decrypted
* `shift` - The number of positions to shift each letter up the alphabet

### leet(plaintext)

Converts a plaintext message to Leet (1337) speak.

#### Arguments

* `plaintext` - The text to be converted to Leet speak

### get_user_input(prompt)

Prompts the user for input until a valid choice is entered.

#### Arguments

* `prompt` - The text to be displayed to the user as a prompt

### main_menu()

Displays the main menu and handles user input.

### get_shift_value(note)

Prompts the user for the shift value to be used in encryption or decryption.

#### Arguments

* `note` - The text to be displayed to the user as a prompt

## Example

#### Encryption
```
Welcome to your C3473R Cipher
Select what you want to do:

----------
M41N M3NU
----------
1. Encrypt a new string
2. Decrypt an existing string
3. Exit

What would you like to do? 1
Enter string to encrypt: Hello world!
Enter a number to shift values by: 3
Your encrypted phrase is: Khoor zruog!
```

#### Decryption
Welcome to your C3473R Cipher
Select what you want to do:

----------
M41N M3NU
----------
1. Encrypt a new string
2. Decrypt an existing string
3. Exit

What would you like to do? 2
Enter string to decrypt: Khoor zruog!
Enter a number to shift values by: 3
Your decrypted phrase is: Hello world!
