# JSON Parser

### Overview

This is a simple JSON parser project written in Go. The purpose of this project is to provide a tool to parse JSON data efficiently and extract information from it. JSON (JavaScript Object Notation) is a lightweight data interchange format commonly used for transmitting data between a server and a web application. This parser allows users to parse JSON strings and access the data within them programmatically.

### Features

- Parses JSON strings into a structured format.
- Supports nested objects and arrays.
- Provides methods to access and manipulate parsed JSON data.
- Lightweight and easy to use.

### Development Steps

This JSON parser was developed incrementally in several steps to add new functionality gradually. Each step focused on parsing different types of JSON data.

#### Step 1: Parse Empty Object

In this step, the goal was to parse a valid simple JSON object ('{}') and an invalid JSON file, correctly reporting which is which. A very simple lexer and parser were built for this step.

#### Step 2: Parse String Key-Value Pairs

In this step, the parser was extended to parse a simple JSON object containing string keys and string values, e.g., {"key": "value"}.

#### Step 3: Parse Various Data Types

In this step, the parser was extended to parse a JSON object containing string, numeric, boolean, and null values, e.g.,

```json
{
  "key1": true,
  "key2": false,
  "key3": null,
  "key4": "value",
  "key5": 101
}
```

#### Step 4: Parse Objects and Arrays

In this step, the parser was further extended to parse a JSON object with object and array values, e.g.,
```json
{
  "key": "value",
  "key-n": 101,
  "key-o": {},
  "key-l": []
}
```

### Tests
The tests are divided into directories named for the corresponding step to show the progression in the development of the parser's functionality.