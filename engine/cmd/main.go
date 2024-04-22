package main

import (
    "log/slog"
    "engine/internal/engine"
)

var exitChan = make(chan bool) 

func main(){
    engine.Engine_Start(&exitChan);
    <- exitChan
    slog.Info("Engine Exited")
}

func Exit(){exitChan<-false;}
