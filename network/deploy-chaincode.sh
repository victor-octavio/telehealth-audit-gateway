#!/bin/bash
set -e

REPO_DIR="$(cd "$(dirname "$0")/.." && pwd)"
CHAINCODE_PATH="$REPO_DIR/chaincode/diagnostic-logs"
CHAINCODE_NAME="diagnostic-log"
CHANNEL_NAME="mychannel"
ORDERER_ADDRESS="orderer-api.127-0-0-1.nip.io:8080"

# detect next version and sequence
CURRENT_SEQ=$(peer lifecycle chaincode querycommitted --channelID "$CHANNEL_NAME" --name "$CHAINCODE_NAME" 2>/dev/null | grep -o 'Sequence: [0-9]*' | awk '{print $2}')

if [ -z "$CURRENT_SEQ" ]; then
  VERSION="1.0"
  SEQUENCE=1
else
  SEQUENCE=$((CURRENT_SEQ + 1))
  VERSION="${SEQUENCE}.0"
fi

LABEL="${CHAINCODE_NAME}_${VERSION}"

echo "==> Deploying $CHAINCODE_NAME v$VERSION (sequence $SEQUENCE)"

echo "==> [1/6] Vendoring dependencies..."
cd "$CHAINCODE_PATH"
go mod tidy
go mod vendor
cd "$REPO_DIR"

echo "==> [2/6] Packaging chaincode..."
peer lifecycle chaincode package "${CHAINCODE_NAME}.tar.gz" \
  --path "$CHAINCODE_PATH" \
  --lang golang \
  --label "$LABEL"

echo "==> [3/6] Installing on peer..."
peer lifecycle chaincode install "${CHAINCODE_NAME}.tar.gz"

echo "==> [4/6] Capturing package ID..."
PACKAGE_ID=$(peer lifecycle chaincode queryinstalled | grep "$LABEL" | awk '{print $3}' | tr -d ',')

if [ -z "$PACKAGE_ID" ]; then
  echo "ERROR: Package ID not found for label $LABEL"
  exit 1
fi

echo "    Package ID: $PACKAGE_ID"

echo "==> [5/6] Approving for org..."
peer lifecycle chaincode approveformyorg \
  -o "$ORDERER_ADDRESS" \
  --channelID "$CHANNEL_NAME" \
  --name "$CHAINCODE_NAME" \
  --version "$VERSION" \
  --sequence "$SEQUENCE" \
  --package-id "$PACKAGE_ID"

echo "==> [6/6] Committing to channel..."
peer lifecycle chaincode commit \
  -o "$ORDERER_ADDRESS" \
  --channelID "$CHANNEL_NAME" \
  --name "$CHAINCODE_NAME" \
  --version "$VERSION" \
  --sequence "$SEQUENCE"

rm -f "${REPO_DIR}/${CHAINCODE_NAME}.tar.gz"

echo ""
echo "Chaincode $CHAINCODE_NAME v$VERSION deployed successfully!"
echo ""
echo "Test with:"
echo "  peer chaincode query -C $CHANNEL_NAME -n $CHAINCODE_NAME -c '{\"function\":\"ReadDiagnosis\",\"Args\":[\"diag001\"]}'"