package ch

import (
	"context"
	"net"
	"testing"

	"github.com/go-faster/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
	"golang.org/x/sync/errgroup"
)

func TestServer_Serve(t *testing.T) {
	ln, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)

	lg := zaptest.NewLogger(t)

	done := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	g, ctx := errgroup.WithContext(ctx)
	s := NewServer(ServerOptions{
		Logger: lg.Named("srv"),
		OnError: func(err error) {
			assert.NoError(t, err, "server error")
			cancel()
		},
	})
	g.Go(func() error {
		defer close(done)
		c, err := Dial(ctx, ln.Addr().String(), Options{
			Logger: lg.Named("usr"),
		})
		if err != nil {
			return errors.Wrap(err, "dial")
		}
		if err := c.Ping(ctx); err != nil {
			return errors.Wrap(err, "ping")
		}
		if err := c.Do(ctx, Query{Body: "HELLO"}); err != nil {
			return errors.Wrap(err, "query")
		}
		return c.Close()
	})
	g.Go(func() error {
		<-done
		return ln.Close()
	})
	g.Go(func() error {
		if err := s.Serve(ln); err != nil {
			if errors.Is(err, net.ErrClosed) {
				return nil
			}
			return err
		}
		return nil
	})
	require.NoError(t, g.Wait())
}
