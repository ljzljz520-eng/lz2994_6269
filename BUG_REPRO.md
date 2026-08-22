# BUG_REPRO

The following failures were observed while validating the initial project state.
Each section records what failed, how to reproduce it, and the complete command output.
They are preserved intentionally; only failing build gates are omitted from the generated Dockerfile.

## Failure 1: Go test (.)

- Observed problem: `Go test (.)` failed in the initial project state.
- Working directory: `.`
- Command: `cd /app && GOTOOLCHAIN=local GOPROXY=off GOSUMDB=off go test -count=1 ./...`
- Exit status: `1`

```text
?   	ticketgate/cmd/gate-session-demo	[no test files]
--- FAIL: TestBusinessChain37 (0.01s)
    business_chain_test.go:30: expected complete validation record, got "37"
FAIL
FAIL	ticketgate	0.027s
ok  	ticketgate/internal/crypto	0.001s
ok  	ticketgate/internal/domain	0.001s
ok  	ticketgate/internal/fixture	0.002s
ok  	ticketgate/internal/logsummary	0.002s
ok  	ticketgate/internal/query	0.009s
ok  	ticketgate/internal/service	0.008s
ok  	ticketgate/internal/store	0.009s
FAIL
```

## Architecture reproduction

### linux/amd64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/gate-session-demo): exit `0`
### linux/arm64
- Go toolchain version: exit `0`
- Go build (.): exit `0`
- Go test (.): exit `1`
- Go run smoke (cmd/gate-session-demo): exit `0`
