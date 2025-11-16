# realtime-ledger

Minimal double-entry ledger API in Go with balance invariants and a pluggable store. The current build ships an in-memory store and REST endpoints for accounts, balances, journals, and entries.

## Features
- Double-entry validation: postings must balance per currency; zero-amount postings rejected.
- Balance safety: rejects operations that would make an account negative.
- Currency-aware amounts using a typed `Money` model.
- Simple auth middleware with static API tokens for quick testing.
- Health endpoint for liveness probes.

## Quickstart
Requirements: Go 1.22+

Run the server (uses in-memory store, listens on `:6767`):
```
go run ./cmd/server
```

Auth: send `Authorization: 00000000` (or `aaaaaaaa`, `05f717e5`, `deadbeef` from `internal/http/middleware.go`).

## API Cheatsheet
Replace `<ACCOUNT_ID>` with the ID returned by create.

Create account:
```
curl -X POST http://localhost:6767/api/v1/accounts \
  -H "Authorization: 00000000" -H "Content-Type: application/json" \
  -d '{"name":"Cash","currency":"USD"}'
```

Get account:
```
curl -H "Authorization: 00000000" \
  http://localhost:6767/api/v1/accounts/<ACCOUNT_ID>
```

Get balance:
```
curl -H "Authorization: 00000000" \
  http://localhost:6767/api/v1/accounts/<ACCOUNT_ID>/balance
```

Post a balanced journal (debit and credit on the same account):
```
curl -X POST http://localhost:6767/api/v1/journals \
  -H "Authorization: 00000000" -H "Content-Type: application/json" \
  -d '{"at":"2024-01-01T00:00:00Z","entries":[
        {"account_id":"<ACCOUNT_ID>","amount":1000,"side":"debit"},
        {"account_id":"<ACCOUNT_ID>","amount":1000,"side":"credit"}
      ]}'
```

List entries (optionally filter by account):
```
curl -H "Authorization: 00000000" \
  "http://localhost:6767/api/v1/entries?account_id=<ACCOUNT_ID>"
```

Health check:
```
curl http://localhost:6767/healthz
```

## Project Layout
- `cmd/server/` – main entrypoint wiring router, service, and in-memory store.
- `internal/ledger/` – domain types (Account, Entry, Posting, Money) and core invariants.
- `internal/service/` – service layer, store interface, and in-memory implementation.
- `internal/http/` – mux router, handlers, auth middleware.
- `internal/auth/` – placeholder for API key utilities.

## Roadmap
- Swap in persistent store (e.g., Postgres) implementing `LedgerStore`.
- Add idempotency keys and journal batching semantics.
- Tighten auth/config via env vars and remove static tokens.
- Publish OpenAPI and add integration/unit tests around posting rules.
