# File Encryption/Decryption Program

This is a command-line program written in Go that allows you to encrypt and decrypt files using a custom password and supports multiple encryption algorithms.

## Usage

To use the program, run the following command:

```
go run main.go --file <file-path> --password <password> [--algorithm <algorithm>] [--decrypt]
```

where:

- `file-path`: Path to the file to encrypt/decrypt. Required.
- `password`: Password used to encrypt/decrypt the file. Required.
- `algorithm`: Encryption algorithm to use. Optional. Supported algorithms: aes-128-gcm, aes-192-gcm, or aes-256-gcm. Default: aes-256-gcm.
- `decrypt`: Flag indicating whether to decrypt the file. Optional. Default: encrypt the file.

## Examples

Encrypt a file using the default algorithm (aes-256-gcm):

```
go run main.go --file my-file.txt --password my-password
```

Encrypt a file using a specific algorithm (aes-128-gcm):

```
go run main.go --file my-file.txt --password my-password --algorithm aes-128-gcm
```

Decrypt a file:

```
go run main.go --file my-file.txt.enc --password my-password --decrypt
```


## Security Considerations

This program uses AES encryption in Galois/Counter Mode (GCM) to encrypt files. GCM is a secure encryption mode that provides confidentiality and authenticity. However, the security of the encryption depends on the strength of the password used to derive the encryption key. Therefore, it is important to use a strong, unique password for each file that is encrypted.

This program does not provide any protection against attacks such as brute-force password guessing or side-channel attacks. Therefore, it should not be used for encrypting files that contain sensitive information without proper auditing and testing.

## License

This program is released under the [MIT License](LICENSE).