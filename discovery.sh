#!/bin/bash

session="discovery"

tmux new-session -d -s $session

tmux send-keys "q src/q/discovery.q -p 1800" C-m

sleep 1

tmux split-window -v

tmux send-keys "cd src/go;go run main.go" C-m

tmux a
