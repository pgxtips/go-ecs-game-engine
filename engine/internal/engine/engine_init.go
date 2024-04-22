package engine

import(
    "engine/internal/renderer"
    "engine/internal/window"
    "log/slog"
)

const (
    width = 640
    height = 480
)

func Engine_Start(exitChan *chan bool){
    windowObj := window.Window_Init("Engine Window", width, height, exitChan);
    go window.Start_Window(windowObj, exitChan);
    go start_Engine();
}

func start_Engine(){
    slog.Info("Initialising Renderer")
    count := 0;
    for true{
        count++;
        renderer.Update();
    }
}
