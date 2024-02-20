# File Encryption Tool

This is a simple web application for encrypting files using Triple DES encryption algorithm. It provides a user-friendly interface for uploading files and encrypting their contents securely.

## Features

- Upload any file and encrypt its contents using Triple DES encryption.
- User-friendly interface with file upload functionality.
- Real-time encryption without storing the file on the server.

## Tech Stack

- **Backend**: Go (Golang)
- **Frontend**: HTML, CSS, JavaScript
- **Encryption Algorithm**: Triple DES (Data Encryption Standard)

## How to Use

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/yourusername/file-encryption-tool.git
    ```

2. Navigate to the project directory:

    ```bash
    cd file-encryption-tool
    ```

3. Start the Go backend server:

    ```bash
    go run main.go
    ```

4. Open your web browser and navigate to `http://localhost:8080` to access the file encryption tool.

5. Choose a file to upload and click the "Encrypt" button. The encrypted data will be displayed on the page.
