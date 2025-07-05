# Zena CLI - AI-Powered Command Line Assistant

![Zena CLI Logo](https://raw.githubusercontent.com/DEEPAKsingh74/zena/refs/heads/main/assets/zena_logo.jpg)

Zena is an intelligent command-line assistant that uses AI models (OpenAI, Anthropic, Gemini) to help you generate, understand, and run commands directly from your terminal.

## Features

- Natural language query processing
- Support for multiple AI providers (OpenAI, Anthropic, Gemini)
- Command explanations
- Configuration management for API keys
- Cross-platform support (Windows, Linux)
- Version management

## Installation

### Linux Installation

1. **Download the binary**:
   ```bash
   wget https://github.com/DEEPAKsingh74/zena/releases/download/v0.1.0/zena-linux.tar.gz
   ```

2. **Extract it**:
   ```bash
   tar -xvzf zena-linux.tar.gz
   ```

2. **Make it executable**:
   ```bash
   chmod +x zena
   ```

3. **Move to your PATH**:
   ```bash
   sudo mv zena /usr/local/bin/
   ```

4. **Verify installation**:
   ```bash
   zena version
   ```

### Windows Installation

1. **Download the binary**:
   - Visit the [releases page](https://github.com/DEEPAKsingh74/zena/releases) and download `zena-windows.zip`

2. **Extract the zip file**:


3. **Run the automation**:
   - right click on install-zena.ps1 and select run on powershell, or simple double click on it to run it.

4. **Verify installation**:
   ```powershell
   zena version
   ```

## Configuration

Before using Zena, you need to configure at least one AI provider API key:

```bash
zena config set openai <your-openai-api-key>
zena config set default openai
```

Available providers: `openai`, `anthropic`, `gemini`

### View current configuration

```bash
zena config --list
```

## Usage

### Basic Query

```bash
zena "how to create a zip file in linux"
```

### Command Generation

```bash
zena "give me a curl command to send a POST request with JSON"
```

### Directory Operations

```bash
zena "remove a directory recursively in bash"
```

### Configuration Management

Set API keys:
```bash
zena config set openai <your-api-key>
zena config set anthropic <your-api-key>
zena config set gemini <your-api-key>
```

Set default provider:
```bash
zena config set default openai
```

List configuration:
```bash
zena config --list
```

### Version Information

```bash
zena version
```

## Examples

1. **Simple command help**:
   ```bash
   zena "how to check disk space on linux"
   ```

2. **Complex command generation**:
   ```bash
   zena "create a bash script that monitors CPU usage and logs to a file when it exceeds 90%"
   ```

3. **Code explanation**:
   ```bash
   zena "explain this command: find . -type f -name '*.txt' -exec grep -l 'hello' {} \;"
   ```


## Contributing

Contributions are welcome! Please open an issue or submit a pull request on our [GitHub repository](https://github.com/DEEPAKsingh74/zena).

## License

MIT License. See [LICENSE](https://github.com/DEEPAKsingh74/zena/LICENSE) for details.
