package client

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/cloudquery/plugin-sdk/v4/plugin"
	"github.com/cloudquery/plugin-sdk/v4/writers"
	"github.com/rs/zerolog"

	"github.com/snowflakedb/gosnowflake"
)

type Client struct {
	plugin.UnimplementedSource
	db     *sql.DB
	logger zerolog.Logger
	spec   Spec
	writer *writers.BatchWriter
}

func New(ctx context.Context, logger zerolog.Logger, spec []byte) (plugin.Client, error) {
	var err error
	c := &Client{
		logger: logger.With().Str("module", "sf-dest").Logger(),
	}
	if err := json.Unmarshal(spec, &c.spec); err != nil {
		return nil, fmt.Errorf("failed to unmarshal snowflake spec: %w", err)
	}
	c.spec.SetDefaults()
	if err := c.spec.Validate(); err != nil {
		return nil, err
	}
	c.writer, err = writers.NewBatchWriter(c, writers.WithLogger(logger), writers.WithBatchSize(c.spec.BatchSize), writers.WithBatchSizeBytes(c.spec.BatchSizeBytes))
	if err != nil {
		return nil, err
	}
	cfg, err := gosnowflake.ParseDSN(c.spec.ConnectionString)
	if err != nil {
		return nil, err
	}
	binaryFormat := "BASE64"
	cfg.Params["BINARY_INPUT_FORMAT"] = &binaryFormat
	cfg.Params["BINARY_OUTPUT_FORMAT"] = &binaryFormat
	dsn, err := gosnowflake.DSN(cfg)
	if err != nil {
		return nil, err
	}
	db, err := sql.Open("snowflake", dsn)
	if err != nil {
		return nil, err
	}
	c.db = db
	if _, err := c.db.ExecContext(ctx, createOrReplaceFileFormat); err != nil {
		return nil, fmt.Errorf("failed to create file format %s: %w", createOrReplaceFileFormat, err)
	}
	if _, err := c.db.ExecContext(ctx, createOrReplaceStage); err != nil {
		return nil, fmt.Errorf("failed to create stage %s: %w", createOrReplaceStage, err)
	}
	return c, nil
}

func (c *Client) Close(context.Context) error {
	var err error
	if c.db == nil {
		return fmt.Errorf("client already closed or not initialized")
	}
	err = c.db.Close()
	c.db = nil
	return err
}
