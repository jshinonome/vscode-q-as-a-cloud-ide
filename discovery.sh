#!/bin/bash

test -d kdb-common || git clone https://github.com/BuaBook/kdb-common.git

session="discovery"

tmux new-session -d -s $session

tmux split-window -v

tmux split-window -v

tmux send-keys -t 1 "q src/q/discovery.q -p 1800" C-m
sleep 1
tmux send-keys -t 2 "cd src/go;go run main.go" C-m

tmux send-keys -t 3 "q kdb-common/boot.q -p 2190 --load-libs compress,rand,mail,file,tz" C-m
tmux send-keys -t 3 "\l data/dailypx" C-m

tmux select-layout tiled

tmux a
