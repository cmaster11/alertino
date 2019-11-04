set -Eeumo pipefail
DIR=$(dirname "$(command -v greadlink >/dev/null 2>&1 && greadlink -f "$0" || readlink -f "$0")")

go run "$DIR/../." \
  -c "$DIR/config.yaml" \
  -m "mongodb://localhost:27417" \
  -v debug
