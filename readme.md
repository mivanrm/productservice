# ProductService

A Go-based product management service that allows you to manage products, variants, and reviews.

## Table of Contents

- [Overview](#overview)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)


## Overview

ProductService is a Go-powered product management service designed to simplify product, variant, and review management.

Key features include:

- Create and manage products with variants.
- Collect and display product reviews.
- Calculate average product ratings.

## Getting Started

To get started with ProductService, follow the installation and usage instructions below.

### Prerequisites

Before you begin, ensure you have the following installed:

- Go (v1.16 or higher)
- PostgreSQL (v11 or higher)

### Installation

1. Clone the repository:

   ```
   git clone https://github.com/mivanrm/productservice.git
   ```
2. Set up your PostgreSQL database and update the configuration in config.yaml.
3. Navigate to the project directory
    ```
   cd productservice
   ```
4. Install the project dependencies:
    ```
    go mod tidy
    ```
5. Build and run the project:
    ```
    go run main.go
    ```