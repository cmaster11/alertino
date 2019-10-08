# `alertino` architecture/concept

* It should accept any kind of JSON payload (ideally considered an alert in another system)
* Deduplicate alerts using templates to determine the hash key
* Configurable via YAML to accept/parse different types of alerts payloads

Advanced:

* Acknowledging alerts (prevents spamming, requires authentication)