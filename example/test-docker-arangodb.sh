set -Eeumo pipefail
DIR=$(dirname "$(command -v greadlink >/dev/null 2>&1 && greadlink -f "$0" || readlink -f "$0")")

# Run an ArangoDB instance
# docker create --name arangodb-persist arangodb true
docker run -e ARANGO_ROOT_PASSWORD=asdasd -p 8529:8529 -it --rm --volumes-from arangodb-persist arangodb:3.6
