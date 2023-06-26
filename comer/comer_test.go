package comer

import "testing"

func TestComer(t *testing.T) {
	comerIns := NewComer()
	if comerIns.Version() == `` {
		t.Errorf(`comer.Version() test failed`)
	}
	if comerIns.goVersion() == `` {
		t.Errorf(`comer.goVersion() test failed`)
	}
	if comerIns.showTips() != `showTips called` {
		t.Errorf(`showTips test failed`)
	}
	// comerIns.Start(nil, nil)
	comerIns.generateFrameworkDir()
	comerIns.generateFrameworkFiles()
	comerIns.generateAppDir()
	comerIns.generateAppFiles()
	comerIns.showAppTips()
	comerIns.addAppsDepend()
	comerIns.addAppRouterDepend()
	// comerIns.initApp(nil, nil)
	// comerIns.init(nil, nil)
	// comerIns.generateDirByName(`testdir`)
	// comerIns.generateFileByMap(`abc`, `a`, `d`)
	comerIns.showLogo()
}
