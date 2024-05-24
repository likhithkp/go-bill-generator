# CLI Billing System in Go

This project is a simple Command Line Interface (CLI) application written in Go. It takes user input for names, items with their prices, and tips, then generates a bill and saves it in a `.txt` file.

## Features

- Accepts user input for customer names, items, and prices
- Calculates the total bill including tips
- Generates a formatted bill
- Saves the bill to a `.txt` file

## Prerequisites

Before you begin, ensure you have met the following requirements:

- You have [Go](https://golang.org/doc/install) installed on your machine.

## Installation

1. Clone this repository to your local machine:

    ```bash
    git clone https://github.com/likhithkp/go-bill-generator.git
    cd go-bill-generator
    ```

2. Build the application:

    ```bash
    go build -o go-bill-generator

    ```

## Usage

1. Run the application:

    ```bash
    ./go-bill-generator

    or

    go run main.go bills.go
    ```

2. Follow the on-screen prompts to enter the required information:

    - Enter the customer's name.
    - Enter the items and their prices.
    - Enter the tip amount.

3. After entering all the required information, the application will generate a bill and save it in a `.txt` file.

## Bill format 

Bill breakdown: 
 - Whiskers Biscuit:         ...$4.99
 - Shiba toys:               ...$5.89
 - CheeseCake:               ...$2.99
 - tip :                     ...$15.50
 - total :                   ...$29.37



