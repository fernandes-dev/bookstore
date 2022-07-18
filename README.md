<h1 align="center">Welcome to bookstore 👋</h1>
<p>
  <a href="https://www.npmjs.com/package/bookstore" target="_blank">
    <img alt="Version" src="https://img.shields.io/npm/v/bookstore.svg">
  </a>
  <a href="#" target="_blank">
    <img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-yellow.svg" />
  </a>
  <a href="https://twitter.com/fernandesdotts" target="_blank">
    <img alt="Twitter: fernandesdotts" src="https://img.shields.io/twitter/follow/fernandesdotts.svg?style=social" />
  </a>
</p>

> Simple CRUD API Rest using GORM and Postgresql

## Usage

```sh
go run cmd/main/main.go
```

### Endpoints

- List all books
    ```http request
    GET http://localhost:3333/book
    ```

- Get book by id
    ```http request
    GET http://localhost:3333/book/{book_id}
    ```

- Create a new book
    ```http request
    POST http://localhost:3333/book
    {
     "name": "Clean Architecture",
     "author": "Robert C. Martin",
     "publication": "Alta Books"
    }
    ```

- Update book by id
    ```http request
    PUT http://localhost:3333/book/{book_id}
    {
     "name": "Clean Architecture",
     "author": "Robert C. Martin",
     "publication": "Alta Books"
    }
    ```

- Delete book by id
    ```http request
    DELETE http://localhost:3333/book/{book_id}
    ```

## Author

👤 **Eduardo Fernandes**

* Twitter: [@fernandesdotts](https://twitter.com/fernandesdotts)
* Github: [@fernandes-dev](https://github.com/fernandes-dev)
* LinkedIn: [@fernandes-dev](https://linkedin.com/in/fernandes-dev)

## Show your support

Give a ⭐️ if this project helped you!

***
_This README was generated with ❤️ by [readme-md-generator](https://github.com/kefranabg/readme-md-generator)_