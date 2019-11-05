# `alertino` TODOs

* [x] It should accept any kind of JSON payload (ideally considered an alert in another system)
* [ ] Proxy alerts via outputs
* [ ] Deduplicate alerts using templates to determine the hash key (requires DB)
* [ ] Configurable via YAML to accept/parse different types of alerts payloads
* [ ] Show list of current alerts (requires UI)
* [ ] Acknowledging alerts (prevents spamming, requires authentication)
* [ ] Complex alerts: observe last states of multiple inputs and raise alerts when specific conditions are met