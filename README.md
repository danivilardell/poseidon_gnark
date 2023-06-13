# Poseidon Implementation using gnark

This project implements the Poseidon hash function using gnark, a zero-knowledge proof (ZKP) circuit generator library in Go.

## Introduction

Poseidon is a cryptographic hash function that is widely used in various applications and widely used in zero-knowledge proofs. This implementation provides a straightforward way to generate Poseidon hash circuits using the gnark library.

My circuit and tests are based on [this](https://github.com/iden3/go-iden3-crypto) other Go implementation.

## Features

- Generate Poseidon hash circuits with customizable parameters.
- Efficient and secure implementation using gnark's circuit generation capabilities.

## Installation

1. Install Go by following the official [Go installation guide](https://golang.org/doc/install).
2. Install the gnark library by running the following command:
   ```shell
   go get github.com/consensys/gnark
   go get github.com/consensys/gnark-crypto
   ```

