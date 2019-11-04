set -Eeumo pipefail
DIR=$(dirname "$(command -v greadlink >/dev/null 2>&1 && greadlink -f "$0" || readlink -f "$0")")

curl "http://localhost:8080/input/webhook-test" -v \
  -H "Content-Type: application/json" -X POST \
  -d '{"key1":"value1", "name": "baaaaaa", "UC": "upperCase!"}'
