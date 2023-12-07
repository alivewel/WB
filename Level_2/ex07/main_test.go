package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	t.Run("OneChannelCloses", func(t *testing.T) {
		c1 := make(chan interface{})
		close(c1)

		orDone := or(c1)
		_, ok := <-orDone
		if ok {
			t.Error("Expected orDone to be closed")
		}
	})

	t.Run("MultipleChannelsOneCloses", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})
		close(c1)

		orDone := or(c1, c2)
		_, ok := <-orDone
		if ok {
			t.Error("Expected orDone to be closed")
		}
	})

	t.Run("AllChannelsClose", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})
		c3 := make(chan interface{})
		close(c1)
		close(c2)
		close(c3)

		orDone := or(c1, c2, c3)
		_, ok := <-orDone
		if ok {
			t.Error("Expected orDone to be closed")
		}
	})

	t.Run("NoChannels", func(t *testing.T) {
		orDone := or()
		_, ok := <-orDone
		if ok {
			t.Error("Expected orDone to be closed")
		}
	})

	t.Run("BlockingCase", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})

		go func() {
			time.Sleep(100 * time.Millisecond)
			close(c1)
		}()

		orDone := or(c1, c2)

		select {
		case _, ok := <-orDone:
			if !ok {
				t.Error("Expected orDone to be open")
			}
		case <-time.After(200 * time.Millisecond):
			t.Error("Expected orDone to not block indefinitely")
		}
	})
}
