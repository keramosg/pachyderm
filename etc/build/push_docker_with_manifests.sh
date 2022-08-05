#!/bin/bash

set +x -euo pipefail

#REPO=localhost:5001/pachyderm
REPO=pachyderm/

for product in pachd worker pachctl mount-server; do
    echo "push $product $VERSION..."
    echo ""
    docker push $REPO/$product-amd64:$VERSION
    docker push $REPO/$product-arm64:$VERSION

    docker manifest create --insecure $REPO/$product:$VERSION \
           $REPO/$product-amd64:$VERSION \
           $REPO/$product-arm64:$VERSION

    docker manifest annotate $REPO/$product:$VERSION $REPO/$product-amd64:$VERSION --arch amd64
    docker manifest annotate $REPO/$product:$VERSION $REPO/$product-arm64:$VERSION --arch arm64

    docker manifest inspect $REPO/$product:$VERSION
    docker manifest push $REPO/$product:$VERSION
    echo "ok"
    echo ""
done
