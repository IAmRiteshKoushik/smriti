# Smriti: MCP Testing Client -- unreleased version

Smriti is a CLI tool for testing and interacting with Model Context Protocol (MCP)
servers over Streamable HTTP. Inspired by tools like Bruno, Smriti uses TOML 
files to define reusable MCP scenarios that can be version controlled, shared 
and executed from the command line through declarative configurations.

## Features

- TOML-based scenario definitions
- MCP tool discovery `tools/list`
- MCP tool invocation `tools/call`
- Workspace-based project structure
- Environment-specific variables
- Streamable HTTP transport
- Git-friendly collections

## Workspace structure

```
workspace/
├── mcpx.toml
├── scenarios/
├── servers/
├── vars/
└── captures/
```

## Example server

```toml
name = "local"

[transport]
type = "streamable-http"
url = "http://localhost:8080/mcp"

[arguments]
account_id = "{{checking_account}}"
```

## Commands

```
smriti run <scenario-path>
smriti servers
smriti scenarios
```

## Author

Ritesh Koushik
