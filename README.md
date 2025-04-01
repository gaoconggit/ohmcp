# OMCP

OMCP is a command-line tool for managing MCP server installations and IDE configurations. It provides Docker-like commands for downloading and running MCP servers with various IDEs.

## Installation

```bash
go install github.com/yourusername/omcp@latest
```

## Usage

### Pull an MCP Server
```bash
omcp pull [version]    # Downloads a specific version of MCP server
omcp pull             # Downloads the latest version
```

### Run an MCP Server
```bash
omcp run [version] --ide cursor --port 8080    # Run a specific version with IDE and port
omcp run --ide windsurf                        # Run latest version with Windsurf IDE
omcp run --ide cline                           # Run latest version with Cline IDE
```

### Flags
- `--ide, -i`: Specify the IDE (cursor, windsurf, cline)
- `--port, -p`: Specify the port number (default: 8080)

## Supported IDEs
- Cursor
- Windsurf
- Cline
