# Gollege

<img align="right" width="180px" src="https://raw.githubusercontent.com/swaggo/swag/master/assets/swaggo.png">

**Gollege** is an API built with [Gin](https://gin-gonic.com/), a web framework written in [Go](https://go.dev/), designed to be extremely fast and efficient.

This project provides all the high school entrance exams in Brazil, including question sets, answers, verifications, and image handling.

## Table of Contents
- [About](#about)
- [Features](#features)
- [Technologies Used](#technologies-used)
- [Installation](#installation)
- [API Documentation](#api-documentation)
- [Authors](#authors)

## About

**Gollege** is an API designed to manage high school entrance exam questions, specifically for Brazilian exams such as Enem (National High School Exam). The API serves questions categorized by year and subject, with easy-to-use routes to access these data.

## Features
- **Question Listing:**  
  Access to a variety of questions organized by year and category (e.g., languages, natural sciences).
  
- **Fast and Secure Routes:**  
  Optimized endpoints for quick access and efficient data handling.

- **Multiple Year Support:**  
  The API provides access to historical data of past entrance exam questions by academic year (e.g., 2023, 2022).

- **Swagger Integration:**  
  Automatic documentation for easy understanding and testing of API endpoints.

## Technologies Used
Gollege uses several technologies to ensure high performance and security:
- **Go**: The main programming language used for building the API backend.
- **Gin**: A web framework for Go, used for routing HTTP requests efficiently.
- **Swagger**: A tool for auto-generating API documentation, integrated for better usability.
- **Gin-Swagger**: Provides Swagger-based documentation specifically for Gin-based APIs.

## Installation

## To run the project locally, follow these steps:

### 1. Clone the repository:
```bash
git clone https://github.com/reallease/gollege.git
2. Navigate to the project directory:
cd gollege
3. Install the dependencies:
go mod tidy
4. Run the application:
go run main.go
API Documentation
The API documentation is generated using Swagger and is accessible at:

Swagger UI

Routes
GET /v1/enem/:year/:lang
Retrieves questions for the specified year and language category.

year: The year of the exam (e.g., 2023).
lang: The subject category (e.g., languages, natural_sciences).
License
Gollege is licensed under the MIT License. See the LICENSE file for more information.
```

## 👨‍💻 Authors

### 🌟 [Storm](https://github.com/reallease)  
**🔧 Lead Developer**  
Passionate about creating powerful solutions with Go.  
🔗 [GitHub Profile](https://github.com/reallease)

---

### 🙌 [Felipe](https://github.com/GabriewF)  
**💻 Contributing Developer**  
Focused on building scalable applications with clean code.  
🔗 [GitHub Profile](https://github.com/GabriewF)

---

> "We are grateful for the contributions of these talented individuals who made this project possible!" 🌟
