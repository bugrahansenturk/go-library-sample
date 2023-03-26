# Library REST API

This is a simple library REST API implemented in Go. It allows users to manage books and borrows, and supports basic user authentication and role-based access control.

## Data Model

The data model consists of three entities: `Book`, `User`, and `Borrow`.

### Book

A `Book` entity represents a book in the library. Each book has the following fields:

- ID: Unique identifier for the book
- Title: The title of the book
- Author: The author of the book

### User

A `User` entity represents a library user. Each user has the following fields:

- ID: Unique identifier for the user
- FirstName: The first name of the user
- LastName: The last name of the user
- Email: The email address of the user
- Role: The user's role, which can be one of the following values: `Member` or `Expired`

### Borrow

A `Borrow` entity represents a book borrowed by a user. Each borrow has the following fields:

- ID: Unique identifier for the borrow
- UserID: The ID of the user who borrowed the book
- BookID: The ID of the book being borrowed
- DueDate: The due date for returning the borrowed book

## Server Structure

The server is organized into three main packages: `domain`, `services`, and `handlers`.

### domain

The `domain` package contains the data model entities (`Book`, `User`, `Borrow`) and their corresponding types and constants.

### services

The `services` package contains the business logic for managing books, users, and borrows. It includes a separate service for each entity (`BookService`, `UserService`, `BorrowService`). These services manage in-memory storage for the entities and handle CRUD operations.

### handlers

The `handlers` package contains the HTTP handlers for the API endpoints. Each handler is responsible for handling a specific entity and maps HTTP requests to corresponding service functions.

## Future Work Ideas

1. **Database integration**: Replace the in-memory storage with a persistent database, such as PostgreSQL or MongoDB, to store the books, users, and borrows data.

2. **Improved authentication**: Implement a more secure authentication system, such as JWT-based authentication, to protect the API endpoints.

3. **Fine management**: Add a fine management system for users who return books late. Calculate the fine based on the number of days a book is overdue and add a method to pay the fine.

4. **Search and filtering**: Implement search and filtering functionality for books, allowing users to find books by title, author, or other criteria.

5. **Advanced user management**: Add more user roles (e.g., librarian, administrator) and implement role-based access control for different API endpoints.
