package main

import (
    "fmt"
)

func main() {
    fmt.Println("Welcome to crypter")
    for {
        fmt.Print("Enter a command: (encrypt, decrypt, exit): ")

        var command string

        _, err := fmt.Scanln(&command)

        if err != nil {
            fmt.Println("Error: Please enter a valid command.")
            fmt.Println()
            continue
        }

        switch command {
        case "encrypt":
            fmt.Println("You chose to encrypt")
            encrypt()
        case "decrypt":
            fmt.Println("You chose to decrypt")
        case "exit":
            fmt.Println("Exiting the program!")
            return
        default:
            fmt.Println("Invalid command. Please enter 'encrypt', 'decrypt', or 'exit'")
            fmt.Println()
        }

        fmt.Println("exiting program, goodbye!")
        return
    }
}

func encrypt() {
    fmt.Print("Which file to encrypt?: ")

    var encryptFile string

    _, err := fmt.Scanln(&encryptFile)

    if err != nil {
        fmt.Println("Invalid file error")
    }

    fmt.Printf("Encrypting file: %s\n", encryptFile)


}
