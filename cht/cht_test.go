package cht_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/go-faster/ch"
	"github.com/go-faster/ch/cht"
	"github.com/go-faster/ch/proto"
)

func TestConnect(t *testing.T) {
	ctx := context.Background()
	server := cht.New(t)

	client, err := ch.Dial(ctx, server.TCP, ch.Options{})
	require.NoError(t, err)

	t.Log("Connected", client.Location())
	t.Cleanup(func() {
		require.NoError(t, client.Close())
	})

	t.Run("CreateTable", func(t *testing.T) {
		// Create table, no data fetch.
		createTable := ch.Query{
			Body: "CREATE TABLE test_table (id UInt64) ENGINE = MergeTree ORDER BY id",
		}
		require.NoError(t, client.Do(ctx, createTable))
	})
	t.Run("SelectOne", func(t *testing.T) {
		// Select single row.
		var data proto.ColUInt8
		selectOne := ch.Query{
			Body: "SELECT 1 AS one",
			Result: proto.Results{
				{
					Name: "one",
					Data: &data,
				},
			},
		}
		require.NoError(t, client.Do(ctx, selectOne))
		require.Len(t, data, 1)
		require.Equal(t, byte(1), data[0])
	})
}
