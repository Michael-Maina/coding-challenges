# Rate Limiter Project
## Overview
This project implements a rate limiter in Go, supporting three different algorithms: token bucket, fixed window counter, and sliding window counter. Additionally, it enables rate limiting to work across multiple servers using Redis as a distributed cache.

## Components
  - **Token Bucket Algorithm**: Maintains a bucket of tokens replenished at a fixed rate. Requests are allowed if there are available tokens.

  - **Fixed Window Counter Algorithm**: Counts the number of requests within fixed time windows. Requests are allowed if they don't exceed a predefined threshold within a window.

  - **Sliding Window Counter Algorithm**: Maintains a sliding window of fixed duration to count requests. Requests are allowed if they don't exceed a predefined threshold within the window.

  - **Redis Integration**: Utilizes Redis as a distributed cache to synchronize rate limiting information across multiple servers.

## Project Structure
 - **token_bucket.go**: Implementation of the token bucket algorithm.
  - **fixed_window_counter.go**: Implementation of the fixed window counter algorithm.
 - **sliding_window_counter.go**: Implementation of the sliding window counter algorithm.
 - **redis_integration.go**: Implementation of Redis integration for distributed rate limiting.
 - **ratelimiter_test.go**: Unit tests for rate limiter algorithms and Redis integration.

## Conclusion
This project provides implementations of rate limiting algorithms and Redis integration in Go. It aims to facilitate understanding and experimentation with different rate limiting techniques and their application in distributed systems.