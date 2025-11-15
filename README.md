# realtime-ledger
Real-time double-entry ledger backend with strict consistency, idempotent writes, and streaming balance updates.


```
realtime-ledger/
  cmd/
    server/
      main.go
  internal/
    config/
      config.go
    http/
      router.go
      middleware.go
      handlers/
        accounts.go
        ledger.go
        stream.go
        health.go
    ledger/
      ledger.go          # core posting logic and invariants
      money.go           # Money type, currency handling
      models.go          # domain objects (Account, Entry, Posting, BalanceSnapshot)
      errors.go
    store/
      postgres/
        tx.go            # transaction helpers
        accounts_repo.go
        ledger_repo.go
        idempotency_repo.go
        migrations/      # optional, or use /migrations at root
    stream/
      hub.go             # WebSocket/SSE hub
      subscription.go
    auth/
      api_key.go         # simple token auth
    observability/
      logging.go
      metrics.go
    util/
      time.go
  migrations/
    0001_create_accounts.sql
    0002_create_ledger_entries.sql
    ...
  api/
    openapi.yml          # later
  scripts/
    load_test.go
    seed_demo.go
```