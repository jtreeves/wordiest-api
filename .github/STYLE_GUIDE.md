# Style Guide

Most styling is enforced via Go's linting and formatting rules. However, we have highlighted some of the specifics and notable deviations.

## Naming Conventions

- Use short yet descriptive names for variables, functions, and types; rely on documentation to provide any missing clarity
- Use short abbreviations for internal variables
- Use full words for more globally scoped variables
- Use camelCase for constant, function, and type names that will not be exported
- Use PascalCase for constant, function, and type names that will be exported
- Use snake_case for file names (the first word should indicate the area of focus, while the second word restates the layer indicted by the folder; e.g., `user_handler.go`)
- Limit folder names to only single words, in lowercase (e.g., `handler`)
- For unused variables that are syntactically required, use the underscore (i.e., `_`)

## Modern Go

- Follow the guidelines put forth in [Effective Go](https://go.dev/doc/effective_go)
- Take advantage of the various Go CLI tools

| command | usage                              | example            |
| ------- | ---------------------------------- | ------------------ |
| `fmt`   | reformat source code               | `go fmt ./pkg/...` |
| `fix`   | update packages to use new APIs    | `go fix ./pkg/...` |
| `vet`   | report likely mistakes in packages | `go vet ./pkg/...` |

These and others can all be run from the root of the repo to maintain code quality with the core functionality in the `pkg` folder.

## Echo Approach

- Use route groups
- Exit early with errors

## Code Structure

- Break down complex functions into smaller functions, each focusing on a single responsibility
- Limit the number of parameters to a function (preferably no more than 3)
- Isolate utilities into their own files
- Keep files short (goal of within 100 lines, but none should be longer than 500 lines)
- Put all public packages within the `pkg` folder
- Do not use the `internal` folder to hide functionality
- Ensure top-level executables exist in the `cmd` folder in a `main.go` file and package within a named subfolder
- Place helper functionality that does not need to be unit tested in the external `util` folder; it will call on functionality from the `pkg` packages, but it provides non-standard functionality that is not necessary for public exposure (even through it will be exposed)
- Folder structure within the `pkg` folder should be oriented based on the layers at work (e.g., `handler`, `service`); note the singular (not plural) names
- Tests should appear within the packages they test

## Error Handling

- Use Go's error handling idiom of returning an error as the last return value
- Prefer descriptive error messages and wrap errors using `fmt.Errorf` when adding context
- Handle errors gracefully and provide useful error messages to the API client

```go
if err != nil {
    return fmt.Errorf("failed to fetch user: %w", err)
}
```

## Logging

- Log at appropriate levels: `DEBUG`, `INFO`, `WARN`, `ERROR`, `FATAL`
- Include sufficient context in log messages to aid in debugging

## Documentation

- Maintain an up-to-date `README.md` with project overview, useful resources, and any other pertinent introductory material
- Use Go doc comments to provide single-sentence overviews of packages and functions, granular information about types, and multiline summaries for executables
- Use comments to explain the purpose of the code, not how it works
- Provide API documentation for all handlers via Swagger/OpenAPI
