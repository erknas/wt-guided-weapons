package versionparser

import (
	"context"
	"fmt"
	"strings"

	csvreader "github.com/erknas/wt-guided-weapons/internal/lib/csv-reader"
	"github.com/erknas/wt-guided-weapons/internal/types"
)

const vUrl = "https://docs.google.com/spreadsheets/d/1SsOpw9LAKOs0V5FBnv1VqAlu3OssmX7DJaaVAUREw78/export?format=csv&gid=1624345539"

type CSVVersionParser struct {
	reader csvreader.Reader
}

func New(reader csvreader.Reader) *CSVVersionParser {
	return &CSVVersionParser{
		reader: reader,
	}
}

func (p *CSVVersionParser) Parse(ctx context.Context) (types.VersionInfo, error) {
	data, err := p.reader.Read(ctx, vUrl)
	if err != nil {
		return types.VersionInfo{}, fmt.Errorf("failed to read CSV: %w", err)
	}

	fields := strings.Fields(data[3][0])

	return types.VersionInfo{Version: fields[len(fields)-1]}, nil
}
