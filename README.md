# 🧩 Sudoku Solver in Go

Welcome to the **Sudoku Solver** project! This Go-based program efficiently solves 9x9 Sudoku puzzles using a backtracking algorithm. If you love puzzles and Go, this project is for you! 🚀

---

## 📜 Table of Contents

- [✨ Features](#-features)
- [📂 Project Structure](#-project-structure)
- [⚡ Installation & Setup](#-installation--setup)
- [🚀 Usage](#-usage)
- [🔬 Example Inputs & Outputs](#-example-inputs--outputs)
- [🛠 How It Works](#-how-it-works)
- [📌 Contributing](#-contributing)
- [📜 License](#-license)

---

## ✨ Features

✅ Solves valid 9x9 Sudoku puzzles instantly
✅ Detects and rejects invalid or unsolvable puzzles
✅ Lightweight and runs fast
✅ Simple command-line interface
✅ Built with pure Go (no dependencies!)

---

## 📂 Project Structure

```
sudoku/
│── main.go       # The core Sudoku solver program
│── README.md     # Documentation (You're reading it now!)
```

---

## ⚡ Installation & Setup

### Prerequisites

Ensure you have Go installed on your system:

```sh
go version  # Should return a Go version (e.g., go1.20)
```

### Clone the Repository

```sh
git clone https://github.com/your-username/sudoku_solver.git
cd sudoku_solver
```

### Run the Program

```sh
go run main.go ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

---

## 🚀 Usage

Pass a 9x9 Sudoku board as arguments where:

- `.` represents an empty cell
- Numbers (1-9) represent pre-filled values

**Example Command:**

```sh
go run main.go "..96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

**Example Output:**

```
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

**Handling Errors:**

- If input is incorrect or the puzzle is unsolvable, the program prints:

```
Error
```

---

## 🔬 Example Inputs & Outputs

### ✅ Valid Sudoku

**Input:**

```sh
go run main.go ".96.4...1" "1...6...4" "5.481.39." "..795..43" ".3..8...." "4.5.23.18" ".1.63..59" ".59.7.83." "..359...7"
```

**Output:**

```
3 9 6 2 4 5 7 8 1
1 7 8 3 6 9 5 2 4
5 2 4 8 1 7 3 9 6
2 8 7 9 5 1 6 4 3
9 3 1 4 8 6 2 7 5
4 6 5 7 2 3 9 1 8
7 1 2 6 3 8 4 5 9
6 5 9 1 7 4 8 3 2
8 4 3 5 9 2 1 6 7
```

### ❌ Invalid Input

```sh
go run main.go 1 2 3 4
```

**Output:**

```
Error
```

---

## 🛠 How It Works

1. Reads the Sudoku grid from command-line arguments
2. Validates the input format
3. Uses **backtracking** to find a solution
4. Prints the solved Sudoku or an error message

---

## 📌 Contributing

Want to improve this project? Feel free to fork the repo and submit a PR!

1. Fork the repository
2. Create a feature branch (`git checkout -b feature-xyz`)
3. Commit your changes (`git commit -m 'Added cool feature'`)
4. Push to your branch (`git push origin feature-xyz`)
5. Open a pull request

---

## 📜 License

This project is licensed under the MIT License. Feel free to use and modify it as needed.

💡 **Happy Coding!** 🎯
