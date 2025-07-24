# Go development workflow rules

After any code change, do the following steps:

1. Run `go build .` If the build passes, go to step 2. If the build has any errors, repeat this step. If you can't fix the build in 3 tries, ask for help.
2. Run `go test $(go list ./... | grep -v /build_out/)`. If the tests pass, go to step 3.  If you can't fix the build in 3 tries, ask for help.
3. Commit your changes with `git`.