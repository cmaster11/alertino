set -Eeumo pipefail
DIR=$(dirname "$(command -v greadlink >/dev/null 2>&1 && greadlink -f "$0" || readlink -f "$0")")

go run . \
  -c "$DIR/example/app-config.yaml" \
  -i "$DIR/example/io-config.yaml" \
  -v debug
