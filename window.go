package glfw

//#include <stdlib.h>
//#include <GLFW/glfw3.h>
//void glfwSetWindowPosCallbackCB(GLFWwindow *window);
//void glfwSetWindowSizeCallbackCB(GLFWwindow *window);
//void glfwSetWindowCloseCallbackCB(GLFWwindow *window);
//void glfwSetWindowRefreshCallbackCB(GLFWwindow *window);
//void glfwSetWindowFocusCallbackCB(GLFWwindow *window);
//void glfwSetWindowIconifyCallbackCB(GLFWwindow *window);
import "C"

import "unsafe"

const (
	Focused   = C.GLFW_FOCUSED
	Iconified = C.GLFW_ICONIFIED
	Visible   = C.GLFW_VISIBLE
	Resizable = C.GLFW_RESIZABLE
	Decorated = C.GLFW_DECORATED
)

const (
	ClientApi           = C.GLFW_CLIENT_API
	ContextVersionMajor = C.GLFW_CONTEXT_VERSION_MAJOR
	ContextVersionMinor = C.GLFW_CONTEXT_VERSION_MINOR
	ContextRobustness   = C.GLFW_CONTEXT_ROBUSTNESS
	OpenglForwardCompat = C.GLFW_OPENGL_FORWARD_COMPAT
	OpenglDebugContext  = C.GLFW_OPENGL_DEBUG_CONTEXT
	OpenglProfile       = C.GLFW_OPENGL_PROFILE
)

const (
	OpenglApi   = C.GLFW_OPENGL_API
	OpenglEsApi = C.GLFW_OPENGL_ES_API
)

const (
	NoRobustness        = C.GLFW_NO_ROBUSTNESS
	NoResetNotification = C.GLFW_NO_RESET_NOTIFICATION
	LoseContextOnReset  = C.GLFW_LOSE_CONTEXT_ON_RESET
)

const (
	OpenglNoProfile     = C.GLFW_OPENGL_NO_PROFILE
	OpenglCoreProfile   = C.GLFW_OPENGL_CORE_PROFILE
	OpenglCompatProfile = C.GLFW_OPENGL_COMPAT_PROFILE
)

type Window struct {
	data *C.GLFWwindow
}

type (
	goPositionFunc func(*Window, int, int)
	goSizeFunc     func(*Window, int, int)
	goCloseFunc    func(*Window)
	goRefreshFunc  func(*Window)
	goFocusFunc    func(*Window, int)
	goIconifyFunc  func(*Window, int)
)

var (
	fPositionHolder goPositionFunc
	fSizeHolder     goSizeFunc
	fCloseHolder    goCloseFunc
	fRefreshHolder  goRefreshFunc
	fFocusHolder    goFocusFunc
	fIconifyHolder  goIconifyFunc
)

//export goPositionCB
func goPositionCB(window unsafe.Pointer, xpos, ypos C.int) {
	fPositionHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(xpos), int(ypos))
}

//export goSizeCB
func goSizeCB(window unsafe.Pointer, width, height C.int) {
	fSizeHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(width), int(height))
}

//export goCloseCB
func goCloseCB(window unsafe.Pointer) {
	fCloseHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))})
}

//export goRefreshCB
func goRefreshCB(window unsafe.Pointer) {
	fRefreshHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))})
}

//export goFocusCB
func goFocusCB(window unsafe.Pointer, focused C.int) {
	fFocusHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(focused))
}

//export goIconifyCB
func goIconifyCB(window unsafe.Pointer, iconified C.int) {
	fIconifyHolder(&Window{(*C.GLFWwindow)(unsafe.Pointer(window))}, int(iconified))
}

//DefaultWindowHints resets all window hints to their default values.
func DefaultWindowHints() {
	C.glfwDefaultWindowHints()
}

//WindowHint function sets hints for the next call to CreateWindow. The hints,
//once set, retain their values until changed by a call to WindowHint or
//DefaultWindowHints, or until the library is terminated with Terminate.
func WindowHint(target, hint int) {
	C.glfwWindowHint(C.int(target), C.int(hint))
}

//CreateWindow creates a window and its associated context. Most of the options
//controlling how the window and its context should be created are specified
//through WindowHint.
//
//Successful creation does not change which context is current. Before you can
//use the newly created context, you need to make it current using
//MakeContextCurrent.
//
//Note that the created window and context may differ from what you requested,
//as not all parameters and hints are hard constraints. This includes the size
//of the window, especially for full screen windows. To retrieve the actual
//attributes of the created window and context, use queries like
//GetWindowAttrib and GetWindowSize.
//
//To create the window at a specific position, make it initially invisible using
//the Visible window hint, set its position and then show it.
//
//If a fullscreen window is active, the screensaver is prohibited from starting.
//
//Windows: If the executable has an icon resource named GLFW_ICON, it will be
//set as the icon for the window. If no such icon is present, the IDI_WINLOGO
//icon will be used instead.
//
//Mac OS X: The GLFW window has no icon, as it is not a document window, but the
//dock icon will be the same as the application bundle's icon. Also, the first
//time a window is opened the menu bar is populated with common commands like
//Hide, Quit and About. The (minimal) about dialog uses information from the
//application's bundle. For more information on bundles, see the Bundle
//Programming Guide provided by Apple.
func CreateWindow(width, height int, title string, monitor *Monitor, share *Window) *Window {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	return &Window{C.glfwCreateWindow(C.int(width), C.int(height), t, monitor.data, share.data)}
}

//Destroy destroys the specified window and its context. On calling this
//function, no further callbacks will be called for that window.
func (w *Window) Destroy() {
	C.glfwDestroyWindow(w.data)
}

//ShouldClose returns the value of the close flag of the specified window.
func (w *Window) ShouldClose() bool {
	r := int(C.glfwWindowShouldClose(w.data))
	if r == C.GL_FALSE {
		return false
	}
	return true
}

//SetShouldClose sets the value of the close flag of the window. This can be
//used to override the user's attempt to close the window, or to signal that it
//should be closed.
func (w *Window) SetShouldClose(value int) {
	C.glfwSetWindowShouldClose(w.data, C.int(value))
}

//SetTitle sets the window title, encoded as UTF-8, of the window.
func (w *Window) SetTitle(title string) {
	t := C.CString(title)
	defer C.free(unsafe.Pointer(t))
	C.glfwSetWindowTitle(w.data, t)
}

//GetPosition returns the position, in screen coordinates, of the upper-left
//corner of the client area of the window.
func (w *Window) GetPosition() (int, int) {
	var (
		xpos C.int
		ypos C.int
	)

	C.glfwGetWindowPos(w.data, &xpos, &ypos)
	return int(xpos), int(ypos)
}

//SetPosition sets the position, in screen coordinates, of the upper-left corner
//of the client area of the window.
//
//If it is a full screen window, this function does nothing.
//
//If you wish to set an initial window position you should create a hidden
//window (using WindowHint and Visible), set its position and then show it.
//
//It is very rarely a good idea to move an already visible window, as it will
//confuse and annoy the user.
//
//The window manager may put limits on what positions are allowed.
func (w *Window) SetPosition(xpos, ypos int) {
	C.glfwSetWindowPos(w.data, C.int(xpos), C.int(ypos))
}

//GetSize returns the size, in screen coordinates, of the client area of the
//specified window.
func (w *Window) GetSize() (int, int) {
	var (
		width  C.int
		height C.int
	)

	C.glfwGetWindowSize(w.data, &width, &height)
	return int(width), int(height)
}

//SetSize sets the size, in screen coordinates, of the client area of the
//window.
//
//For full screen windows, this function selects and switches to the resolution
//closest to the specified size, without affecting the window's context. As the
//context is unaffected, the bit depths of the framebuffer remain unchanged.
//
//The window manager may put limits on what window sizes are allowed.
func (w *Window) SetSize(width, height int) {
	C.glfwSetWindowSize(w.data, C.int(width), C.int(height))
}

//Iconfiy iconifies/minimizes the window, if it was previously restored. If it
//is a full screen window, the original monitor resolution is restored until the
//window is restored. If the window is already iconified, this function does
//nothing.
func (w *Window) Iconify() {
	C.glfwIconifyWindow(w.data)
}

//Restore restores the window, if it was previously iconified/minimized. If it
//is a full screen window, the resolution chosen for the window is restored on
//the selected monitor. If the window is already restored, this function does
//nothing.
func (w *Window) Restore() {
	C.glfwRestoreWindow(w.data)
}

//Show makes the window visible, if it was previously hidden. If the window is
//already visible or is in full screen mode, this function does nothing.
func (w *Window) Show() {
	C.glfwShowWindow(w.data)
}

//Hide hides the window, if it was previously visible. If the window is already
//hidden or is in full screen mode, this function does nothing.
func (w *Window) Hide() {
	C.glfwHideWindow(w.data)
}

//GetMonitor returns the handle of the monitor that the window is in
//fullscreen on.
func (w *Window) GetMonitor() *Monitor {
	return &Monitor{C.glfwGetWindowMonitor(w.data)}
}

//GetAttribute returns an attribute of the window. There are many attributes,
//some related to the window and others to its context.
func (w *Window) GetAttribute(attrib int) int {
	return int(C.glfwGetWindowAttrib(w.data, C.int(attrib)))
}

//SetUserPointer sets the user-defined pointer of the window. The current value
//is retained until the window is destroyed. The initial value is nil.
func (w *Window) SetUserPointer(pointer unsafe.Pointer) {
	C.glfwSetWindowUserPointer(w.data, pointer)
}

//GetUserPointer returns the current value of the user-defined pointer of the
//window. The initial value is nil.
func (w *Window) GetUserPointer() unsafe.Pointer {
	return C.glfwGetWindowUserPointer(w.data)
}

//SetPositionCallback sets the position callback of the window, which is called
//when the window is moved. The callback is provided with the screen position
//of the upper-left corner of the client area of the window.
func (w *Window) SetPositionCallback(cbfun goPositionFunc) {
	fPositionHolder = cbfun
	C.glfwSetWindowPosCallbackCB(w.data)
}

//SetSizeCallback sets the size callback of the window, which is called when
//the window is resized. The callback is provided with the size, in screen
//coordinates, of the client area of the window.
func (w *Window) SetSizeCallback(cbfun goSizeFunc) {
	fSizeHolder = cbfun
	C.glfwSetWindowSizeCallbackCB(w.data)
}

//SetCloseCallback sets the close callback of the window, which is called when
//the user attempts to close the window, for example by clicking the close
//widget in the title bar.
//
//The close flag is set before this callback is called, but you can modify it at
//any time with SetShouldClose.
//
//Mac OS X: Selecting Quit from the application menu will trigger the close
//callback for all windows.
func (w *Window) SetCloseCallback(cbfun goCloseFunc) {
	fCloseHolder = cbfun
	C.glfwSetWindowCloseCallbackCB(w.data)
}

//SetRefreshCallback sets the refresh callback of the window, which
//is called when the client area of the window needs to be redrawn, for example
//if the window has been exposed after having been covered by another window.
//
//On compositing window systems such as Aero, Compiz or Aqua, where the window
//contents are saved off-screen, this callback may be called only very
//infrequently or never at all.
func (w *Window) SetRefreshCallback(cbfun goRefreshFunc) {
	fRefreshHolder = cbfun
	C.glfwSetWindowRefreshCallbackCB(w.data)
}

//SetFocusCallback sets the focus callback of the window, which is called when
//the window gains or loses focus.
//
//After the focus callback is called for a window that lost focus, synthetic key
//and mouse button release events will be generated for all such that had been
//pressed. For more information, see SetKeyCallback and SetMouseButtonCallback.
func (w *Window) SetFocusCallback(cbfun goFocusFunc) {
	fFocusHolder = cbfun
	C.glfwSetWindowFocusCallbackCB(w.data)
}

//SetIconifyCallback sets the iconification callback of the window, which is
//called when the window is iconified or restored.
func (w *Window) SetIconifyCallback(cbfun goIconifyFunc) {
	fIconifyHolder = cbfun
	C.glfwSetWindowIconifyCallbackCB(w.data)
}

//PollEvents processes only those events that have already been received and
//then returns immediately. Processing events will cause the window and input
//callbacks associated with those events to be called.
//
//This function is not required for joystick input to work.
//
//This function may not be called from a callback.
func PollEvents() {
	C.glfwPollEvents()
}

//WaitEvents puts the calling thread to sleep until at least one event has been
//received. Once one or more events have been recevied, it behaves as if
//PollEvents was called, i.e. the events are processed and the function then
//returns immediately. Processing events will cause the window and input
//callbacks associated with those events to be called.
//
//Since not all events are associated with callbacks, this function may return
//without a callback having been called even if you are monitoring all
//callbacks.
//
//This function may not be called from a callback.
func WaitEvents() {
	C.glfwWaitEvents()
}
