## Example Workflow: Creating a Simple Web Application

Here's a complete example of creating and running a simple Express.js application:

1. **Create a new terminal session:**

```bash
curl -X POST -H "X-API-Key: your-api-key" \
  http://localhost:8080/api/terminals
```

Response:

```json
{
  "id": "term-1",
  "pwd": "/"
}
```

2. **Create project directory and navigate to it:**

```bash
curl -X POST -H "X-API-Key: your-api-key" -H "Content-Type: application/json" \
  -d '{
    "type": "terminal",
    "payload": {
      "session_id": "term-1",
      "command": "mkdir -p myapp && cd myapp"
    }
  }' \
  http://localhost:8080/api/execute_action
```

3. **Initialize Node.js project:**

```bash
curl -X POST -H "X-API-Key: your-api-key" -H "Content-Type: application/json" \
  -d '{
    "type": "terminal",
    "payload": {
      "session_id": "term-1",
      "command": "npm init -y"
    }
  }' \
  http://localhost:8080/api/execute_action
```

4. **Install Express:**

```bash
curl -X POST -H "X-API-Key: your-api-key" -H "Content-Type: application/json" \
  -d '{
    "type": "terminal",
    "payload": {
      "session_id": "term-1",
      "command": "npm install express"
    }
  }' \
  http://localhost:8080/api/execute_action
```

5. **Create app.js:**

```bash
curl -X POST -H "X-API-Key: your-api-key" -H "Content-Type: application/json" \
  -d '{
    "type": "write_file",
    "payload": {
      "path": "myapp/app.js",
      "content": "const express = require('express');\nconst app = express();\n\napp.get('/', (req, res) => {\n  res.send('Hello, World!');\n});\n\napp.listen(3000, () => {\n  console.log('Server running on port 3000');\n});"
    }
  }' \
  http://localhost:8080/api/execute_action
```

6. **Start the server:**

```bash
curl -X POST -H "X-API-Key: your-api-key" -H "Content-Type: application/json" \
  -d '{
    "type": "terminal",
    "payload": {
      "session_id": "term-1",
      "command": "node app.js"
    }
  }' \
  http://localhost:8080/api/execute_action
```

7. **Check if the server is running:**

```bash
curl -X POST -H "X-API-Key: your-api-key" -H "Content-Type: application/json" \
  -d '{
    "type": "terminal",
    "payload": {
      "session_id": "term-1",
      "command": "ps aux | grep node"
    }
  }' \
  http://localhost:8080/api/execute_action
```

8. **When done, close the terminal session:**

```bash
curl -X DELETE -H "X-API-Key: your-api-key" \
  http://localhost:8080/api/terminals/term-1
```

This workflow demonstrates:

- Proper terminal session management
- Directory navigation
- Package installation
- File creation
- Server startup
- Process verification
- Resource cleanup

Each step builds upon the previous one, and the terminal session maintains its state (current directory, environment variables, etc.) throughout the process.
