#/bin/bash
# you need docker to run this
docker run -w /mnt -v `pwd`:/mnt dqneo/ubuntu-build-essential:go ./test_linux.sh
