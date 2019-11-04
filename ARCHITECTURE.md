# `alertino` architecture

## Inputs

You can setup multiple named inputs (e.g. `grafana-prod`, `graylog-prod`). Each input will trigger a new API endpoint, available with the same name.

Whatever JSON payload is sent to that input will be filtered/treated depending on the specific input config (e.g. to generate the deduplication key).