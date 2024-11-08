

## Best Practice Go Unit Test

This repository serves as a comprehensive guide to writing clean, maintainable, and effective unit tests in Go. It covers essential testing practices, setup configurations, and techniques to make testing easier and more reliable.


## Folder Structure Explanation
- **config** : This folder used for our configuration files such a database config.

- **entity** : Contains a struct definitions of database tables.

- **models** : Used to store models such as structs to hold a request from the client.

- **repository** : Used to store queries database like query for create a product.

- **services** : Where the business logic located.

- **test** : Used to test our function.

- **test/mocks** : Used to mocking our repository function.




## Running Tests

To run tests, run the following command

```bash
  go test ./test
```


Read my article about this repository [here](https://medium.com/@danarcahyadi21/unit-test-in-go-best-practices-for-easier-testing-79d194fe9a54).