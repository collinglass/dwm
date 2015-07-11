package core

// ENVIRONMENT variables
type Window struct{}
type Display struct{}
type Drawable struct{}
type GC struct{}
type XFontSet struct{}
type XFontStruct struct{}
type XEvent struct{}
type XErrorEvent struct{}
type XClassHint struct {
	res_class string
	res_name  string
}
type KeySym uint

// Union arg
type Arg interface{}

// Button
type Button struct {
	click  uint
	mask   uint
	button uint
	do     func(arg *Arg)
	arg    Arg
}

type Client struct {
	name                                                              string
	mina, maxa                                                        float64
	x, y, w, h                                                        int
	oldx, oldy, oldw, oldh                                            int
	basew, baseh, incw, inch, maxw, maxh, minw, minh                  int
	bw, oldbw                                                         int
	tags                                                              uint
	isfixed, isfloating, isurgent, neverfocus, oldstate, isfullscreen bool
	next                                                              *Client
	snext                                                             *Client
	mon                                                               *Monitor
	win                                                               Window
}

type DC struct {
	x, y, w, h int
	norm       [ColLast]uint64
	sel        [ColLast]uint64
	drawable   Drawable
	gc         GC
	font       struct {
		ascent  int
		descent int
		height  int
		set     XFontSet
		xfont   *XFontStruct
	}
}

type Key struct {
	mod    uint
	keysym KeySym
	do     func(arg *Arg)
	arg    Arg
}

type Layout struct {
	symbol string
	do     func(m *Monitor)
}

type Monitor struct {
	ltsymbol       string
	mfact          float64
	nmaster        int
	num            int
	by             int
	mx, my, mw, mh int
	wx, wy, ww, wh int
	seltags        uint
	sellt          uint
	tagset         [2]uint
	showbar        bool
	topbar         bool
	clients        *Client
	sel            *Client
	stack          *Client
	next           *Monitor
	barwin         Window
	lt             *[2]Layout
}

type Rule struct {
	class      string
	instance   string
	title      string
	tags       uint
	isfloating bool
	monitor    int
}

// ENUMS
type Cursor uint
type Color uint
type Atom uint

const (
	// cursor enum
	CurNormal Cursor = 1 + iota
	CurResize
	CurMove
	CurLast

	// color enum
	ColBorder Color = 1 + iota
	ColFG
	ColBG
	ColLast

	// EWMH atoms enum
	NetSupported Atom = 1 + iota
	NetWMName
	NetWMState
	NetWMFullscreen
	NetActiveWindow
	NetWMWindowType
	NetWMWindowTypeDialog
	NetLast

	// default atoms enum
	WMProtocols Atom = 1 + iota
	WMDelete
	WMState
	WMTakeFocus
	WMLast

	// click enum
	ClkTagBar uint = 1 + iota
	ClkLtSymbol
	ClkStatusText
	ClkWinTitle
	ClkClientWin
	ClkRootWin
	ClkLast

	broken = "broken"
)

// variables
var (
	TAGMASK = uint((1 << uint(len(tags))) - 1)

	stext       string
	screen      int
	sw, sh      int
	bh, blw     int
	xerrorxlib  *func(*Display, *XErrorEvent)
	numlockmask uint

	handler = map[*XEvent]func(*XEvent){
		ButtonPress:      buttonpress,
		ClientMessage:    clientmessage,
		ConfigureRequest: configurerequest,
		ConfigureNotify:  configurenotify,
		DestroyNotify:    destroynotify,
		EnterNotify:      enternotify,
		Expose:           expose,
		FocusIn:          focusin,
		KeyPress:         keypress,
		MappingNotify:    mappingnotify,
		MapRequest:       maprequest,
		MotionNotify:     motionnotify,
		PropertyNotify:   propertynotify,
		UnmapNotify:      unmapnotify,
	}

	wmatom  [WMLast]Atom
	netatom [NetLast]Atom
	running = true
	cursor  [CurLast]Cursor
	dpy     *Display
	dc      DC
	mons    *Monitor
	selmon  *Monitor
	root    Window

	// Mocked XEvents
	ButtonPress      *XEvent
	ClientMessage    *XEvent
	ConfigureRequest *XEvent
	ConfigureNotify  *XEvent
	DestroyNotify    *XEvent
	EnterNotify      *XEvent
	Expose           *XEvent
	FocusIn          *XEvent
	KeyPress         *XEvent
	MappingNotify    *XEvent
	MapRequest       *XEvent
	MotionNotify     *XEvent
	PropertyNotify   *XEvent
	UnmapNotify      *XEvent
)
