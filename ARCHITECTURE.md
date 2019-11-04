# `alertino` architecture

## Inputs

You can setup multiple named inputs (e.g. `grafana-prod`, `graylog-prod`). Each input will trigger a new API endpoint, available with the same name.

Whatever JSON payload is sent to that input will be filtered/treated depending on the specific input config (e.g. to generate the deduplication key).

## Alerts

Alerts are generated whenever events come through inputs, if these inputs are configured to treat the specific payload as an alert (by default this is always the case, so each event sent through the input raises an alert).

Alerts are object with a state, and can be marked as acknowledged. Non-acknowledged alerts will be re-thrown with a configured delay.

## Outputs

Outputs are endpoints where alerts are sent, e.g. webhooks/email. Each output can be configured to forward alerts states depending on configured filters (e.g. depending on the source key, resolved state, etc).

Alerts can be sent as clones of the last input event data, or as custom alerts (e.g. through templates). This offers the possibility of using alertino as a mere alert proxy, but with acknowledgment abilities to deduplicate alerts.