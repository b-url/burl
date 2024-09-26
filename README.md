# burl

b(ookmark)url is a developer first bookmark management tool written in Go.


## Development

### Getting Started

To get started with development, clone the repository and run the following commands:

```bash
./scripts/setup.sh
```

This will install the necessary dependencies and set up the project for development.

### Tasks

This project uses [Taskfile](https://taskfile.dev) to manage tasks. To see a list of available tasks, run:

```bash
task --list
```

To enable auto-completion for the available tasks follow the Homebrew shell completion instructions [here](https://docs.brew.sh/Shell-Completion).

### Database migrations
For migrations we use [atlas](https://atlasgo.io/). Here follow some commands to get you started.

Create a new migration:

```bash
atlas migrate new name
```

After adding a migration we need a new hash.

```bash
atlas migrate hash
```
