{
  languages.go.enable = true;

  scripts = {
    build.exec = "go build -o bin/keisai ./cmd";
    run.exec = "go run ./cmd/main.go";
    start.exec = "go build -o bin ./cmd && ./bin";
    test-all.exec = "go test -v ./...";
    coverage.exec = "go test -cover ./...";
    clean.exec = "rm -rf bin";
    fmt.exec = "go fmt ./...";
  };
}
