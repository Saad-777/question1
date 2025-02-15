# Path Converter

A Go program that converts relative file paths to absolute paths. The program handles special cases like "./" (current directory), "../" (parent directory), and multiple slashes, ensuring the final absolute path is normalized.

## Features

- Converts relative paths to absolute paths
- Handles special directory notations (./ and ../)
- Normalizes paths by removing redundant slashes
- Interactive command-line interface
- Takes user input for current directory and relative path

## Prerequisites

- Go (1.16 or higher)

## Installation

1. Clone the repository:
```bash
git clone https://github.com/Saad-777/question1.git
```

2. Navigate to the project directory:
```bash
cd question1
```

## Usage

1. Run the program:
```bash
go run question1.go
```

2. When prompted, enter:
   - The current directory (e.g., /home/myhome/myfolder1)
   - The relative path (e.g., ../myfolder2/mydocument.txt)

3. The program will output the converted absolute path.

## Example

```
Enter current directory (e.g., /home/myhome/myfolder1): /home/myhome/myfolder1
Enter relative path (e.g., ../myfolder2/mydocument.txt): ../myfolder2/mydocument.txt

Absolute path: /home/myhome/myfolder2/mydocument.txt
```

## Input Guidelines

- Current directory should be an absolute path (starting with /)
- Relative path can include:
  - Current directory (./)
  - Parent directory references (../)
  - Multiple consecutive slashes (these will be normalized)
  - File or directory names

