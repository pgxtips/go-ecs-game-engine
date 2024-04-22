package window 

import (
    "log/slog"
	"github.com/go-gl/glfw/v3.3/glfw"
)
// Window_init is a function that initializes the window
// pass in the width and height of the window
func Window_Init(title string, width int, height int, exitChan *chan bool) *glfw.Window {
    slog.Info("Initialising Window")
    newWindow := initGlfw(title, width, height)
    return newWindow
}

func initGlfw(title string, width int, height int) *glfw.Window {
    if err := glfw.Init(); err != nil {
            panic(err)
    }
    
    glfw.WindowHint(glfw.Resizable, glfw.False)
    glfw.WindowHint(glfw.ContextVersionMajor, 4)
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
    glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

    window, err := glfw.CreateWindow(width, height, title, nil, nil)
    if err != nil { panic(err) }

    window.MakeContextCurrent()

    return window
}

func Start_Window(window *glfw.Window, exitChan *chan bool){
    for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
	glfw.Terminate()
    *exitChan<-true
}
