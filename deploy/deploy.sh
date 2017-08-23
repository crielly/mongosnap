#!/bin/bash
binurl=$(curl https://api.github.com/repos/crielly/mongosnap/releases/latest | jq -r '.assets'[0]'.browser_download_url')
echo "Downloading release from $binurl"
rm -rf /usr/local/bin/mongosnap

wget $binurl -O /usr/local/bin/mongosnap
chmod 0755 /usr/local/bin/mongosnap
chown root:root /usr/local/bin/mongosnap
