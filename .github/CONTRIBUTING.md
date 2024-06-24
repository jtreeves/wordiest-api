# Contributing

## Orientation

- Peruse the [README](../README.md) at the root of this repo to familiarize yourself with the project, including its purpose and file structure
- Browse the [Issues](/issues) tab to find open tasks requiring work
- View the project's [Trello board](https://trello.com/b/gTQy2xQK/wordiest) to find tasks ready to be completed
- By participating in this project, you agree to abide by the terms specified in its [Code of Conduct](CODE_OF_CONDUCT.md)
- By contributing to this project, you agree that your contributions will be licensed based on terms specified in the [LICENSE](../LICENSE)

## Set Up

### Prerequisites

- Go (v1.22+)
- Echo (v4.12+)

### Installation

1.  Fork and clone the repository
2.  Install all dependencies:

```bash
go mod download
```

## Issues

### Reporting Bugs

- Use the Issues tab to report any bugs
- Check if the issue has already been reported before creating a new one
- Clearly describe the steps to reproduce the bug, and provide any relevant information about your environment
- Include screenshots or code snippets if they can help to illustrate the issue
- Use the `bug` issue tag

### Feature Requests

- Use the Issues tab to suggest new features or improvements
- Provide detailed information about the proposed feature and its use cases
- Be open to feedback and discussion from the community
- Use the `enhancement` issue tag (and any other relevant existing tags)

### Discussions

- Use the Issues tab to discuss anything else related to the project with the community at large
- Use the `question` issue tag (any any other relevant existing tags)

## Tests

Automated tests help to improve code quality. New code should add new tests, and all old tests should always pass. Tests are run as part of the CI/CD process to ensure errors do not creep in. All packages within the `pkg` folder have corresponding test suites.

To run all unit tests:

```bash
go test ./pkg/...
```

## Pull Requests

### Branching

- Write any new code in a branch off `main`
- Branch names should be short descriptions of the new changes, in kebab-case, optionally prefixed with `feature/` or `bug/` (e.g., `error-codes`)
- Branches should be short-lived, include only code affecting the new functionality, and only touch the files relevant to that new functionality

### Opening a Request

- Once all work is completed, open a PR
- Provide a descriptive title for the PR
- Fill out all fields in the [PR template](PULL_REQUEST_TEMPLATE.md), which includes:
  - Listing all major changes
  - Linking to any relevant Trello cards
  - Ensuring all tests pass
  - Verifying that all functionality works

### Automated Checks

Whenever a pull request is opened, the `Pre-Merge Checks` GitHub Action will be fired. It will run the following checks:

- No linting errors can be found
- All unit tests must pass
- The API must be able to be built

### Code Review

Currently, no approvals are required in order to merge, but any PRs should go through a standard code review process, and you should proactively request reviews from relevant contributors. At least one approval should be obtained, ideally from the project's owner (Jackson Reeves), prior to merging. The following issues are likely to be flagged during code review:

- Deviations from the [style guide](STYLE_GUIDE.md)
- New functionality lacking new tests or lacking sufficient coverage for edge cases

Be responsive to feedback and make any requested changes promptly.

### Merging

Once the pull request has been approved, perform a squash merge to merge it into the `main` branch, thus closing the PR and deleting the branch

## Resources

- [How to Contribute to Open Source](https://opensource.guide/how-to-contribute/)
- [Using Pull Requests](https://help.github.com/articles/about-pull-requests/)
- [GitHub Help](https://help.github.com)
