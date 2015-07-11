package core

import (
	"strings"
)

func applyrules(c *Client) {
	var class, instance string
	var i int
	var m *Monitor
	ch := XClassHint{}

	/* rule matching */
	c.isfloating = false
	c.tags = 0
	XGetClassHint(dpy, c.win, &ch)

	if ch.res_class != "" {
		class = ch.res_class
	} else {
		class = broken
	}
	if ch.res_name != "" {
		instance = ch.res_name
	} else {
		instance = broken
	}

	for i = 0; i < len(rules); i++ {
		r := &rules[i]

		if (r.title == "" || strings.Contains(c.name, r.title)) &&
			(r.class == "" || strings.Contains(class, r.class)) &&
			(r.instance == "" || strings.Contains(instance, r.instance)) {
			c.isfloating = r.isfloating
			c.tags = c.tags | r.tags
			for m = mons; m != nil && m.num != r.monitor; m = m.next {
				if m != nil {
					c.mon = m
				}
			}
		}
	}
	if ch.res_class != "" {
		XFree(ch.res_class)
	}

	if ch.res_name != "" {
		XFree(ch.res_name)
	}

	if c.tags&TAGMASK == 1 {
		c.tags = c.tags & TAGMASK
	} else {
		c.tags = c.mon.tagset[c.mon.seltags]
	}
}

// XEvent Handlers
func buttonpress(e *XEvent)      {}
func clientmessage(e *XEvent)    {}
func configurerequest(e *XEvent) {}
func configurenotify(e *XEvent)  {}
func destroynotify(e *XEvent)    {}
func enternotify(e *XEvent)      {}
func expose(e *XEvent)           {}
func focusin(e *XEvent)          {}
func keypress(e *XEvent)         {}
func mappingnotify(e *XEvent)    {}
func maprequest(e *XEvent)       {}
func motionnotify(e *XEvent)     {}
func propertynotify(e *XEvent)   {}
func unmapnotify(e *XEvent)      {}

// Window managing views
func tile(m *Monitor)    {}
func monocle(m *Monitor) {}

// Keyboard handlers
func spawn(arg *Arg)          {}
func togglebar(arg *Arg)      {}
func focusstack(arg *Arg)     {}
func incnmaster(arg *Arg)     {}
func setmfact(arg *Arg)       {}
func setlayout(arg *Arg)      {}
func view(arg *Arg)           {}
func tag(arg *Arg)            {}
func focusmon(arg *Arg)       {}
func tagmon(arg *Arg)         {}
func zoom(arg *Arg)           {}
func killclient(arg *Arg)     {}
func togglefloating(arg *Arg) {}
func quit(arg *Arg)           {}

//
func toggleview(arg *Arg)  {}
func toggletag(arg *Arg)   {}
func movemouse(arg *Arg)   {}
func resizemouse(arg *Arg) {}

//
//
// X Functions
//
//
func XGetClassHint(d *Display, w Window, xc *XClassHint) {}
func XFree(res string)                                   {}
