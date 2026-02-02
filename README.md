# PassVault Fyne

PassVault is a lightweight, secure, cross-platform password manager built with Go and Fyne.

## Features

- **Secure Storage**: Uses Argon2id for key derivation and AES-256-GCM for encryption.
- **Local Database**: Stores secrets in an encrypted SQLite database.
- **Cross-Platform UI**: Native look and feel on Linux, macOS, and Windows using Fyne.
- **Clipboard Integration**: Auto-clearing clipboard for copied secrets.
- **Dynamic Fields**: Add custom fields to your secrets.

## Prerequisites

### Linux
You need to install the Fyne dependencies (X11 development headers):

```bash
sudo apt-get install golang gcc libgl1-mesa-dev xorg-dev libxxf86vm-dev
```

### macOS / Windows
No special system dependencies required beyond Go and a C compiler (like Xcode or MinGW).

## Installation

1. Clone the repository
2. Initialize dependencies:
   ```bash
   go mod tidy
   ```

## Running the Application

```bash
go run cmd/passvault/main.go
```

## Building

```bash
go build -o passvault cmd/passvault/main.go
```

## Usage

1. **Unlock**: On first run, enter a master password to initialize your vault.
2. **Add Secret**: Click "Add Secret" in the sidebar to create a new entry.
3. **View/Edit**: Select a secret from the list to view details or edit them.
4. **Copy**: Click the "Copy" button next to a field to copy it to clipboard (clears after 30s).

## Project Structure

- `cmd/passvault`: Entry point
- `internal/crypto`: Encryption and key derivation logic
- `internal/database`: SQLite database operations and models
- `internal/state`: Secure memory state management
- `ui`: Fyne GUI components
