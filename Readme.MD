# Todo App in Go

A simple Todo application built using Go, demonstrating basic CRUD operations such as adding, deleting, listing, and marking items as complete. This application uses the `net/http` package for handling HTTP requests and `html/template` for rendering HTML templates.

## Features

- **Add Todo Item**: Create a new todo item.
- **Delete Todo Item**: Remove an existing todo item.
- **List All Items**: Display all todo items with their current status.
- **Mark as Complete**: Change the status of a single todo item to complete.
- **Mark All as Complete**: Change the status of all todo items to complete.

## Technologies Used

- Go
- MongoDB

## Installation

1. **Clone the repository**:

   ```bash
   git clone https://github.com/yourusername/todo-app.git
   cd todo-app
   ```

2. **Install dependencies**:

   ```bash
   go get .
   ```

3. **Run the application**:

   ```bash
   go run cmd/api/main.go
   ```

4. **Access the app**:
   Open your browser and navigate to http://localhost:8000.

## Usage

1. **Add a Todo Item**: Use the input form to add a new item and click "Add".
2. **Delete a Todo Item**: Click the "Delete" button next to the item you want to remove.
3. **List All Items**: All items are displayed on the main page.
4. **Mark as Complete**: Click the "Mark as done" button next to the item to mark it as done.
5. **Mark All as Complete**: Click the "Mark all as done" button to mark all items as complete.
