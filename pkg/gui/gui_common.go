package gui

import (
	"github.com/jesseduffield/gocui"
	"github.com/jesseduffield/lazygit/pkg/commands/oscommands"
	"github.com/jesseduffield/lazygit/pkg/config"
	"github.com/jesseduffield/lazygit/pkg/gui/types"
)

// hacking this by including the gui struct for now until we split more things out
type guiCommon struct {
	gui *Gui
	types.IPopupHandler
}

var _ types.IGuiCommon = &guiCommon{}

func (self *guiCommon) LogAction(msg string) {
	self.gui.LogAction(msg)
}

func (self *guiCommon) LogCommand(cmdStr string, isCommandLine bool) {
	self.gui.LogCommand(cmdStr, isCommandLine)
}

func (self *guiCommon) Refresh(opts types.RefreshOptions) error {
	return self.gui.Refresh(opts)
}

func (self *guiCommon) PostRefreshUpdate(context types.Context) error {
	return self.gui.postRefreshUpdate(context)
}

func (self *guiCommon) RunSubprocessAndRefresh(cmdObj oscommands.ICmdObj) error {
	return self.gui.runSubprocessWithSuspenseAndRefresh(cmdObj)
}

func (self *guiCommon) RunSubprocess(cmdObj oscommands.ICmdObj) (bool, error) {
	return self.gui.runSubprocessWithSuspense(cmdObj)
}

func (self *guiCommon) PushContext(context types.Context, opts ...types.OnFocusOpts) error {
	return self.gui.pushContext(context, opts...)
}

func (self *guiCommon) PopContext() error {
	return self.gui.popContext()
}

func (self *guiCommon) ReplaceContext(context types.Context) error {
	return self.gui.replaceContext(context)
}

func (self *guiCommon) CurrentContext() types.Context {
	return self.gui.currentContext()
}

func (self *guiCommon) CurrentStaticContext() types.Context {
	return self.gui.currentStaticContext()
}

func (self *guiCommon) CurrentSideContext() types.Context {
	return self.gui.currentSideContext()
}

func (self *guiCommon) IsCurrentContext(c types.Context) bool {
	return self.CurrentContext().GetKey() == c.GetKey()
}

func (self *guiCommon) GetAppState() *config.AppState {
	return self.gui.Config.GetAppState()
}

func (self *guiCommon) SaveAppState() error {
	return self.gui.Config.SaveAppState()
}

func (self *guiCommon) RenderString(view *gocui.View, content string) error {
	return self.gui.renderString(view, content)
}

func (self *guiCommon) Render() {
	self.gui.render()
}

func (self *guiCommon) Views() types.Views {
	return self.gui.Views
}

func (self *guiCommon) OpenSearch() {
	_ = self.gui.handleOpenSearch(self.gui.currentViewName())
}

func (self *guiCommon) GocuiGui() *gocui.Gui {
	return self.gui.g
}

func (self *guiCommon) OnUIThread(f func() error) {
	self.gui.onUIThread(f)
}

func (self *guiCommon) RenderToMainViews(opts types.RefreshMainOpts) error {
	return self.gui.refreshMainViews(opts)
}

func (self *guiCommon) MainViewPairs() types.MainViewPairs {
	return types.MainViewPairs{
		Normal:         self.gui.normalMainContextPair(),
		Staging:        self.gui.stagingMainContextPair(),
		PatchBuilding:  self.gui.patchBuildingMainContextPair(),
		MergeConflicts: self.gui.mergingMainContextPair(),
	}
}
