# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-04-10

### Added

- initial release of dev quotes api
- public api endpoints for retrieving developer quotes
- random quote endpoint
- filtered quotes by author and tag
- quote by id endpoint
- rate limiting (50 requests/minute)
- swagger documentation
- docker support
- github actions ci/cd
- cloud run deployment
- health check endpoint
- cors support
- security headers
- pagination support
- error handling
- configuration management
- structured logging with zerolog
- request logging middleware with detailed metrics
- environment-based logging configuration
- makefile with development commands
- pull request template
- comprehensive api documentation with examples
- initial dataset of 20 developer quotes

### Security

- rate limiting per ip address
- cors configuration
- security headers (xss protection, hsts, etc)
- input validation
- request timeout configurations
- environment-specific security settings
