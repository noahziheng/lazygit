package context

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

type SimpleContext struct {
	OnFocus     func(opts types.OnFocusOpts) error
	OnFocusLost func(opts types.OnFocusLostOpts) error
	OnRender    func() error

	*BaseContext
}

type ContextCallbackOpts struct {
	OnFocus     func(opts types.OnFocusOpts) error
	OnFocusLost func(opts types.OnFocusLostOpts) error
	OnRender    func() error
}

func NewSimpleContext(baseContext *BaseContext, opts ContextCallbackOpts) *SimpleContext {
	return &SimpleContext{
		OnFocus:     opts.OnFocus,
		OnFocusLost: opts.OnFocusLost,
		OnRender:    opts.OnRender,
		BaseContext: baseContext,
	}
}

var _ types.Context = &SimpleContext{}

// A Display context only renders a view. It has no keybindings and is not focusable.
func NewDisplayContext(key types.ContextKey, view *gocui.View, windowName string) types.Context {
	return NewSimpleContext(
		NewBaseContext(NewBaseContextOpts{
			Kind:       types.DISPLAY_CONTEXT,
			Key:        key,
			View:       view,
			WindowName: windowName,
			Focusable:  false,
			Transient:  false,
		}),
		ContextCallbackOpts{},
	)
}

func (self *SimpleContext) HandleFocus(opts types.OnFocusOpts) error {
	if self.highlightOnFocus {
		self.GetViewTrait().SetHighlight(true)
	}

	if self.OnFocus != nil {
		if err := self.OnFocus(opts); err != nil {
			return err
		}
	}

	if self.onRenderToMainFn != nil {
		if err := self.onRenderToMainFn(); err != nil {
			return err
		}
	}

	return nil
}

func (self *SimpleContext) HandleFocusLost(opts types.OnFocusLostOpts) error {
	if self.OnFocusLost != nil {
		return self.OnFocusLost(opts)
	}
	return nil
}

func (self *SimpleContext) HandleRender() error {
	if self.OnRender != nil {
		return self.OnRender()
	}
	return nil
}

func (self *SimpleContext) HandleRenderToMain() error {
	if self.onRenderToMainFn != nil {
		return self.onRenderToMainFn()
	}

	return nil
}
