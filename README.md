# Hirego backend.

First clone the repository.

```bash
git clone https://github.com/mohit-051/hirego.git
```

Rename the .env.example file to .env file.

In Unix based system you can use the following:

```bash
mv .env.example .env
```

In windows you can manually do the same.

Next update the required values in the .env file.

Now run the application in local development server.

```bash
go run src/cmd/main.go
```

The backend will run on the following port:

http://localhost:8000
