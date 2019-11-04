# `alertino` concepts

* It should accept any kind of JSON payload (ideally considered an alert in another system)
* Deduplicate alerts using templates to determine the hash key (requires DB)
* Configurable via YAML to accept/parse different types of alerts payloads
* Show list of current alerts (requires UI)
* Acknowledging alerts (prevents spamming, requires authentication)