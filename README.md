# AICommitter

**AICommitter** is a CLI tool designed to generate commit messages using AI. This tool leverages various AI services to produce meaningful and well-structured commit messages based on the `git diff` output.

## Table of Contents

- [AICommitter](#aicommitter)
  - [Table of Contents](#table-of-contents)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Analyze Command](#analyze-command)
      - [Flags:](#flags)
    - [Doctor Command](#doctor-command)
  - [Configuration](#configuration)
    - [Example Configuration](#example-configuration)
  - [Architecture](#architecture)
    - [Domain](#domain)
    - [Use Cases](#use-cases)
    - [Adapters](#adapters)
    - [Interfaces](#interfaces)
  - [Contributing](#contributing)
  - [License](#license)

## Installation

To install AICommitter, ensure you have Go installed on your machine. Then, clone the repository and build the binary:

```bash
git clone https://github.com/guiyomh/aicommitter.git
cd aicommitter
go build -o aicommitter
```

You can then move the binary to your `$PATH` to use it globally.

## Usage

AICommitter provides several commands to interact with your git repository and generate commit messages. The primary commands are `analyze` and `doctor`.

### Analyze Command

The `analyze` command generates a commit message based on the current staged changes in your repository.

```bash
aicommitter analyze [flags]
```

#### Flags:

- `-s, --scope`: Force the scope of the commit.
- `-t, --type`: Force the type of the commit.
- `-i, --issue`: Add the issue number to the commit.
- `-l, --language`: Specify the language for the commit message.
- `-a, --adapter`: Specify the adapter to use for generating the commit message (default: `google_genai`).

### Doctor Command

The `doctor` command checks if your environment is properly set up, ensuring that required tools like `git` and `ollama` are installed.

```bash
aicommitter doctor
```

The `doctor` command will output the status of each check, indicating whether the required tools are installed and configured correctly.

## Configuration

AICommitter uses a configuration file located in the user's home directory at `~/.config/aicommitter.yaml`. The configuration file should include the API key required to interact with AI services like Google GenAI.

### Example Configuration

```yaml
api_key: your_api_key_here
```

## Architecture

AICommitter is built using the principles of Clean Architecture and the Hexagonal Architecture pattern, which ensures that the core domain is decoupled from external services and frameworks.

### Domain

The domain layer contains the core entities and business logic of the application. Key entities include:

- `CommitMessage`: Represents the structure of a commit message.
- `Diff`: Represents the changes detected by `git diff`.
- `Prompt`: Contains the data required to generate a commit message.

### Use Cases

Use cases define the application-specific business logic. In AICommitter, the primary use cases are:

- `AnalyzeUsecase`: Responsible for generating a commit message based on the diff.
- `DoctorUsecase`: Ensures that the environment is correctly configured.

### Adapters

Adapters handle the communication with external services and tools:

- `GoogleGenAIAdapter`: Interacts with Google's GenAI to generate commit messages.
- `OllamaAdapter`: Provides an alternative method for generating commit messages using the Ollama API.

### Interfaces

Interfaces define the contracts between different layers of the application. Key interfaces include:

- `MessageGenerator`: Defines how commit messages are generated.
- `DiffGenerator`: Defines how diffs are generated.
- `ConfigLoader`: Defines how configuration is loaded.

## Contributing

Contributions are welcome! Please fork the repository, create a new branch, and submit a pull request with your changes.

## License

This project is licensed under the GPL 3.0 License. See the [LICENSE](LICENSE.md) file for more details.
