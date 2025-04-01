# OHMCP

OHMCP is a command-line tool for managing MCP server installations and IDE configurations. It provides Docker-like commands for downloading and running MCP servers with various IDEs.

## Installation

```bash
go install github.com/gaoconggit/ohmcp@latest
```

## Usage

### Pull an MCP Server
```bash
ohmcp pull [version]    # Downloads a specific version of MCP server
ohmcp pull             # Downloads the latest version
```

### Run an MCP Server
```bash
ohmcp run [version] --ide cursor --port 8080    # Run a specific version with IDE and port
ohmcp run --ide windsurf                        # Run latest version with Windsurf IDE
ohmcp run --ide cline                           # Run latest version with Cline IDE
```

### Flags
- `--ide, -i`: Specify the IDE (cursor, windsurf, cline)
- `--port, -p`: Specify the port number (default: 8080)

## Supported IDEs
- Cursor
- Windsurf
- Cline
