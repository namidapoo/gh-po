A lightweight GitHub CLI extension that seamlessly combines `gh pr list` and `gh pr checkout` into a single interactive workflow.

## Overview

`gh-po` (PR + Checkout) is a simple wrapper around GitHub CLI commands that streamlines the process of viewing and checking out pull requests. It displays all available pull requests in your repository and allows you to immediately checkout any PR by entering its number.

## Motivation

When working with multiple pull requests, developers often need to:

1. Run `gh pr list` to see available PRs
2. Identify the PR number they want to work on
3. Run `gh pr checkout <number>` to switch to that PR

This extension combines these steps into a single command, reducing context switching and making PR management more efficient. The interactive prompt makes it easy to quickly jump between different pull requests during code reviews or when switching between multiple work streams.

## Installation

### Prerequisites

- [GitHub CLI](https://cli.github.com/) must be installed and authenticated

### Install as a GitHub CLI extension

```bash
gh extension install namidapoo/gh-po
```

## Usage

```bash
gh po
```

This will:

1. Display a list of all pull requests in the current repository
2. Prompt you to enter the number of the PR you want to checkout
3. Automatically checkout the selected PR

### Example

```bash
$ gh po
#123  Feature: Add user authentication  feature-auth
#122  Fix: Memory leak in data parser   fix-memory-leak
#121  Update: Documentation             update-docs

? Type the number of PR to checkout? ... 122
✓ Checked out pull request #122 (fix-memory-leak)
```

The extension inherits all authentication and repository context from the GitHub CLI, so it works seamlessly with your existing `gh` configuration.
