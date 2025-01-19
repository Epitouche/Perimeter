# Developer Documentation

## Table of Contents

- [Developer Documentation](#developer-documentation)
  - [Table of Contents](#table-of-contents)
  - [Main Document](#main-document)
  - [Project Documentation](#project-documentation)
  - [🛠️ How to Generate the Documentation](#️-how-to-generate-the-documentation)
  - [📖 How to Open the Documentation](#-how-to-open-the-documentation)
    - [1. **Using a Browser**](#1-using-a-browser)
    - [2. **Serving Locally (Optional)**](#2-serving-locally-optional)
      - [Option 1: Using `http-server` (Node.js)](#option-1-using-http-server-nodejs)
      - [Option 2: Using Python (if installed)](#option-2-using-python-if-installed)

## Main Document

[main documentation](../../README.md)

## Project Documentation

Welcome to the **Project Documentation** for the React Native app! This guide explains how to access and view the generated documentation.

## 🛠️ How to Generate the Documentation

To generate the documentation for the project, follow these steps:

1. Ensure you have all the necessary dependencies installed. You can install them using:

   ```bash
   npm install
   ```

2. Run the documentation generation script. This script will compile the documentation files into the `docs` directory:

   ```bash
   npm run generate-docs
   ```

3. Once the script completes, the generated documentation will be available in the `docs` directory.

---

## 📖 How to Open the Documentation

### 1. **Using a Browser**

- Navigate to the `docs` folder located in the project directory.
- Open the `index.html` file with your favorite web browser by:
  - Double-clicking the file.
  - Dragging and dropping the file into a browser window.

---

### 2. **Serving Locally (Optional)**

If you'd like to view the documentation via a local server for better navigation, follow these steps:

#### Option 1: Using `http-server` (Node.js)

1. Install `http-server` globally if you don’t have it:

   ```bash
   npm install -g http-server
   ```

2. Navigate to the `docs` directory:

   ```bash
   cd docs
   ```

3. Start the server:

   ```bash
   http-server
   ```

4. Open the URL displayed in the terminal (e.g., `http://localhost:8080`) in your browser.

#### Option 2: Using Python (if installed)

1. Navigate to the `docs` directory:

   ```bash
   cd docs
   ```

2. Start the server:

   ```bash
   python3 -m http.server
   ```

3. Open the URL displayed in the terminal (e.g., `http://localhost:8000`) in your browser.
