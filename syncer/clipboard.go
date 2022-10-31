package syncer

import (
	"context"
	"fmt"

	"golang.design/x/clipboard"
)

func (r *Syncer) StartClipboard() error {
	err := r.initClipboard()
	if err != nil {
		return err
	}

	err = r.listenClipboard()
	if err != nil {
		return err
	}

	return nil
}

func (r *Syncer) initClipboard() error {
	err := clipboard.Init()
	if err != nil {
		return err
	}
	return nil
}

func (r *Syncer) listenClipboard() error {
	ch := clipboard.Watch(context.TODO(), clipboard.FmtText)
	fmt.Println("Listening for clipboard changes...")
	for data := range ch {
		// print out clipboard data whenever it is changed
		r.publishClipboardMessage(string(data))
	}
	return nil
}

func (r *Syncer) writeClipboard(data string) error {
	clipboard.Write(clipboard.FmtText, []byte(data))
	return nil
}
