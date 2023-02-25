import sys

# Encrypt 'plaintext' using Ceaser Cypher with key of 'shift'


def caesar_encrypt(plaintext, shift):
    ciphertext = ''
    for c in plaintext:
        if c.isalpha():
            if c.isupper():
                ciphertext += chr((ord(c) - 65 + shift) % 26 + 65)
            else:
                ciphertext += chr((ord(c) - 97 + shift) % 26 + 97)
        else:
            ciphertext += c
    return ciphertext

# Decrypt 'plaintext' using Ceaser Cypher with key of 'shift'


def caesar_decrypt(ciphertext, shift):
    plaintext = ''
    for c in ciphertext:
        if c.isalpha():
            if c.isupper():
                plaintext += chr((ord(c) - 65 - shift) % 26 + 65)
            else:
                plaintext += chr((ord(c) - 97 - shift) % 26 + 97)
        else:
            plaintext += c
    return plaintext

# Convert a string to leet (1337) speak


def leet(plaintext):
    leetMap = {
        'A': "4",
        'E': "3",
        'G': "6",
        'I': "1",
        'O': "0",
        'S': "5",
        'T': "7",
    }
    leetText = ""
    for char in plaintext.upper():
        if char in leetMap:
            leetText += leetMap[char]
        else:
            leetText += char
    return leetText


def get_user_input(prompt):
    while True:
        try:
            user_input = input(prompt)
            choice = int(user_input)
            if choice not in [1, 2, 3]:
                raise ValueError("Invalid input. Please enter 1, 2 or 3.")
            return choice
        except ValueError:
            print("Invalid input. Please enter 1, 2 or 3.")


# Prints functionality menu
def main_menu():
    print("-" * 10)
    print(leet("MAIN MENU"))
    print("-" * 10)
    while True:
        print("\n1. Encrypt a new string")
        print("2. Decrypt an existing string")
        print("3. Exit")
        userInput = getUserInput("\nWhat would you like to do?")
        try:
            choice = int(userInput)
        except:
            print("Invalid select.")
            print("Please select one of the following options using 1, 2 or 3.")
            continue
        if choice == 1:
            return 1
        elif choice == 2:
            return 2
        elif choice == 3:
            sys.exit(0)
        else:
            print("Invalid select.")
            print("Please select one of the following options using 1, 2 or 3.")


def get_shift_value(note):
    while True:
        input = input(note + "\n")
        try:
            shift = int(input)
        except:
            print("Invalid shift value. Please try again with a valid number.")
            continue
        return shift


if __name__ == '__main__':
    print("Welcome to your " + leet("Ceaser Cipher") +
          "\nSelect what you want to do:")
while True:
    choice = main_menu()
    match choice:
        case 1:
            plaintext = get_user_input("Enter string to encrypt")
            shift = get_shift_value("Enter a number to shift values by")
            print("Your encrypted phrase is: ")
            print(caesar_encrypt(plaintext, shift))
        case 2:
            ciphertext = get_user_input("Enter string to decrypt")
            shift = get_shift_value("Enter a number to shift values by")
            print("Your decrypted phrase is: ")
            print(caesar_decrypt(ciphertext, shift))
        case 3:
            sys.exit(0)
        case _:
            print("Invalid input. Please enter a number between 1 and 3.")

