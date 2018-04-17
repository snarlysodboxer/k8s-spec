package main

import (
	"github.com/snarlysodboxer/k8s-spec/apiserver"
	"github.com/snarlysodboxer/k8s-spec/engine"
	"github.com/snarlysodboxer/k8s-spec/state"
	"plugin"
)

func main() {
	statePluginPath := "./state/etcd/etcd.so"
	enginePluginPath := "./engine/kubectl/kubectl.so"

	statePlugin, err := plugin.Open(statePluginPath)
	if err != nil {
		panic(err)
	}
	stateStore, err := statePlugin.Lookup("Store")
	if err != nil {
		panic(err)
	}
	var store state.Interface
	store, ok := stateStore.(state.Interface)
	if !ok {
		panic("unexpected type from module symbol")
	}

	enginePlugin, err := plugin.Open(enginePluginPath)
	if err != nil {
		panic(err)
	}
	enginer, err := enginePlugin.Lookup("Engine")
	if err != nil {
		panic(err)
	}
	var k8sEngine engine.Interface
	k8sEngine, ok = enginer.(engine.Interface)
	if !ok {
		panic("unexpected type from module symbol")
	}

	server := apiserver.NewServer(store, k8sEngine)
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
