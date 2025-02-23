# 🛠️ Simple CLI File Encryptor (Early Development)

Welcome to the **Simple CLI File Encryptor**! This project is a beginner-friendly command-line tool written in **Go** that allows users to encrypt and decrypt files.

🚧 **Currently, the app focuses on basic file handling (reading and writing). Encryption functionality will be added soon.** 🚧

---

## 📜 **Features (Current Stage)**
1. Accept user commands (`encrypt`, `decrypt`, `exit`).
2. Read content from an input file.
3. Write content to an output file.
4. Handle errors gracefully (missing files, invalid commands).

---

## 🚀 **Installation**

1. **Clone the Repository:**
```bash
git clone https://github.com/zbinx/crypter-go.git
cd crypter-go
```

2. **Ensure Go is Installed:**
Check if Go is installed:
```bash
go version
```
If not, [download Go here](https://golang.org/dl/).

3. **Build the Program:**
```bash
go build -o encryptor
```

4. **Run the Program:**
```bash
./encryptor
```

---

## 📖 **Usage**

1. **Start the App:**
```bash
go run main.go
```

2. **Example Run:**
```
Enter a command (encrypt, decrypt, exit): encrypt

Enter the input file name for encrypt: plain.txt
Enter the output file name for encrypt: encrypted.txt

File successfully encrypted from 'plain.txt' to 'encrypted.txt'.
```

3. **Available Commands:**
- `encrypt` → Copy file content (encryption coming soon).
- `decrypt` → Copy file content (decryption coming soon).
- `exit` → Exit the program.

---

## 📝 **Project Structure**
```
.
├── main.go        // Main CLI app
├── README.md      // Project documentation
└── plain.txt      // Example input file
```

---

## ⚙️ **Planned Features (Coming Soon!)**
1. AES encryption for file content.
2. Secure key-based encryption and decryption.
3. Progress indicators for large files.
4. Improved error handling and validation.

---

## 🤝 **Contributing**
Contributions are welcome! If you have suggestions or improvements, feel free to open an issue or pull request.

---

## 📄 **License**
This project is licensed under the [MIT License](LICENSE).

---

## 🌟 **Acknowledgments**
- [The Go Programming Language](https://golang.org/)
- Community tutorials and documentation

---

💡 **Happy coding!** 😊 If you find this project helpful, consider giving it a ⭐ on GitHub!

