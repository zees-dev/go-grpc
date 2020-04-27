# 6. Godoc Support

Date: 2020-27-04

## Status

Accepted

## Context

Add support for GoDoc

## Decision

[godoc](https://blog.golang.org/godoc) can be used to automagically document Go code based on provided comments.

- Documentation for functions/structs/interfaces can be provided via writing a comment above decleration
- Documentation for packages can be provided via: ```// Package <package-name> ...```
- Documentation for known bugs can be provided via: ```// BUG(r) ...```
- Documentation for deprecated functionality can be provided via: ```// Deprecated: ...```

## Consequences

- Due to convention, code/packages within the **internal** directory are not visible
  - However, documentation is still viewable via visiting the respective endpoint in the godoc web UI
- Only **public** primitives are documented (uppercase declerations)
