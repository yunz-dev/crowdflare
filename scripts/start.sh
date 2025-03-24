#!/bin/bash
# Start both processes

./server &
sudo cloudflared service install "$TUNNEL_TOKEN" &

# Wait for background processes to finish
wait
