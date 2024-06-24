# Style Guide

## Naming Conventions

- Use descriptive and meaningful names for variables, functions, classes, and types
- Use camelCase for variables and function names (including constants)
- Use PascalCase for class and type names
- Use snake_case for file and folder names
- No unused variables, unless it's a parameter, in which case prefix it with an underscore (e.g., `_index`)

## Formatting

- Use 2 spaces for indentation
- Use single quotes for strings
- Use trailing commas

## Modern Dart

- Use `async`/`await` instead of `.then` chaining
- Use `try`/`catch` blocks instead of `.catchError` chaining
- Use regular comments (`//`) instead of API documentation comments (`///`)
- Prefer explicit types to inferred types unless doing so would be overly verbose
- Always provide explicit return types for functions

## Flutter Approach

- Prefer providers to local state
- Prefer `StatelessWidget`s to `StatefulWidget`s

## Code Structure

- Break down complex functions into smaller functions, each focusing on a single responsibility
- Isolate utilities into their own files
- Keep files short (goal of within 100 lines, but none should be longer than 500 lines)
- Use comments sparingly and prefer self-documenting code
