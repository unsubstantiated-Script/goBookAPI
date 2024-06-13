# Go CRUD API Book Server

SIMPLE CRUD API linked up to MySQL

## Libraries Used

* Database - MySQL
* GORM
* json marshall, unmarshall
* Project Structure
* Gorilla Mux

## Project Structure
```mermaid
flowchart TD
    CMD --> main.go
    PKG --> config --> app.go
    PKG --> controllers --> book-controller.go
    PKG --> models --> book.go
    PKG --> routes --> bookstore-routes.go
    PKG --> utils --> utils.go
```

## Routes
```mermaid
    flowchart LR
        book1["/book/"]
        book2["/book/"]
        book_id1["/book/{id}"]
        book_id2["/book/{id}"]
        book_id3["/book/{id}"]
    getBooks --> book1 --> GET
    getBookById --> book_id1 --> GET
    createBook --> book2 --> POST
    deleteBook --> book_id2--> DELETE
    updateBook --> book_id3--> PUT
```

  