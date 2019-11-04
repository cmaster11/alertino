set -Eeumo pipefail
DIR=$(dirname "$(command -v greadlink >/dev/null 2>&1 && greadlink -f "$0" || readlink -f "$0")")

# Run a mongodb instance
docker run \
	-p 27417:27017 \
	-it --rm \
	mongo:4.2.1