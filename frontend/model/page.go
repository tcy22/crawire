package model

import "crawier/engine"

type SearchResult struct {
	Hits int
	Start int
	Items []engine.Item
}
