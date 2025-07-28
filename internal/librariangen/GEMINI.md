# Go development workflow rules

After any code change, do the following steps:

1. Run `go build .` If the build passes, go to step 2. If the build has any errors, repeat this step. If you can't fix the build in 3 tries, ask for help.
2. After implementing new functionality, add or update tests to ensure the new logic is covered.
3. Run `go test ./...`. If the tests pass, go to step 4. If you can't fix the build in 3 tries, ask for help.
4. After unit tests pass, ask to run the binary integration test script with this command: `run source ~/.zshrc && ./run-binary-integration-test.sh`
5. Check `librariangen.log` to verify the output. Analyze the output and ask if anything needs to be fixed.
6. If the integration test passes and the output is correct, then proceed to the git commit step.
