# Random Number Generation Service Documentation

## Overview

The **Random Number Generation Service** is designed to provide high-quality pseudo-random numbers for various applications. This service utilizes the **Mersenne Twister (MT19937)** algorithm, which offers efficient and reliable randomness. The service now supports separate endpoints for requesting a **single random value** and **multiple random values** through different API paths.

## Features

- **High-Quality Randomness**: Powered by the Mersenne Twister 19937 algorithm.
- **Single and Batch Requests**: Request individual random values through `/api/single/...` and multiple random values through `/api/batch/...`.
- **Support for Various Formats**: Generate random integers, 64-bit integers, unsigned 64-bit integers, floating-point numbers, etc.
- **Automatic Reseeding**: The generator reseeds after every N requests to ensure high-quality randomness.
- **Thread-Safe**: The service is safe for concurrent requests.

## API Endpoints

All API requests are made to `/api/` for batch requests or `/api/single/` for single value requests.

### Single Value Endpoints

These endpoints are designed to generate **one** random value at a time.

### `GET /api/single/integer`

Generates a single random integer within a specified range.

#### Parameters:
- `min` (optional): The minimum value of the range (inclusive). Defaults to 0.
- `max` (optional): The maximum value of the range (inclusive). Defaults to 100.

#### Example Request:
```bash
GET /api/single/integer?min=1&max=1000
```

#### Example Response:
```json
{
  "random_integer": 457
}
```

### `GET /api/single/float`

Generates a single random floating-point number between 0 and 1.

#### Example Request:
```bash
GET /api/single/float
```

#### Example Response:
```json
{
  "random_float": 0.6785937
}
```

### `GET /api/single/int64`

Generates a single random signed 64-bit integer.

#### Example Request:
```bash
GET /api/single/int64
```

#### Example Response:
```json
{
  "random_int64": 9223372036854775807
}
```

### `GET /api/single/uint64`

Generates a single random unsigned 64-bit integer.

#### Example Request:
```bash
GET /api/single/uint64
```

#### Example Response:
```json
{
  "random_uint64": 18446744073709551615
}
```

---

### Batch Value Endpoints

These endpoints allow users to request **multiple random values** at once using the `count` parameter.

### `GET /api/batch/integer`

Generates multiple random integers within a specified range.

#### Parameters:
- `min` (optional): The minimum value of the range (inclusive). Defaults to 0.
- `max` (optional): The maximum value of the range (inclusive). Defaults to 100.
- `count` (optional): The number of random integers to generate. Defaults to 1.

#### Example Request:
```bash
GET /api/batch/integer?min=1&max=1000&count=5
```

#### Example Response:
```json
{
  "random_integers": [457, 329, 782, 123, 951]
}
```

### `GET /api/batch/float`

Generates multiple random floating-point numbers between 0 and 1.

#### Parameters:
- `count` (optional): The number of random floating-point numbers to generate. Defaults to 1.

#### Example Request:
```bash
GET /api/batch/float?count=3
```

#### Example Response:
```json
{
  "random_floats": [0.6785937, 0.2345682, 0.9571234]
}
```

### `GET /api/batch/int64`

Generates multiple signed 64-bit integers.

#### Parameters:
- `count` (optional): The number of signed 64-bit integers to generate. Defaults to 1.

#### Example Request:
```bash
GET /api/batch/int64?count=2
```

#### Example Response:
```json
{
  "random_int64": [9223372036854775807, -7234567890123456789]
}
```

### `GET /api/batch/uint64`

Generates multiple unsigned 64-bit integers.

#### Parameters:
- `count` (optional): The number of unsigned 64-bit integers to generate. Defaults to 1.

#### Example Request:
```bash
GET /api/batch/uint64?count=4
```

#### Example Response:
```json
{
  "random_uint64": [18446744073709551615, 12345678901234567890, 9876543210987654321, 5678901234567890123]
}
```

## Summary

The service now provides separate paths for requesting single and multiple random values. For single values, use `/api/single/`, and for batch requests, use `/api/` with the `count` parameter for multiple values. Both interfaces offer flexibility for different use cases, ensuring that random values can be efficiently requested and used.