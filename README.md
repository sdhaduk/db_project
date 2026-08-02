# Code a Database in 45 Steps

Working through the [Trial of Code: Database](https://trialofcode.org/database/) project — a series of test-driven coding puzzles that build a database from scratch in Go, without external dependencies.

## About

This project covers implementing a database engine step by step, from a simple key-value store all the way to an LSM-Tree with SQL support. Each exercise builds on the last, with tests guiding the implementation.

## Sections

| # | Section | Exercises | Topics |
|---|---------|-----------|--------|
| 01 | Log-based Key-Value Storage | 0101–0105 | Serialization, log storage, fsync, checksums |
| 02 | Tables | 0201–0204 | Data types, schemas, CRUD operations |
| 03 | Simple SQL | 0301–0305 | Tokenizer, parser, statement execution |
| 04 | Range Queries | 0401–0405 | Sorting, searching, iterators, ordered encoding |
| 05 | Advanced SQL | 0501–0507 | Operators, expression evaluation, WHERE clauses |
| 06 | Log + Data Integration | 0601–0605 | Atomic updates, SSTable construction, merge sorting |
| 07 | LSM-Tree | 0701–0704 | Tree structure, multi-level management, merging |
| 08 | Indexing | 0801–0805 | Index structures |
| 09 | Concurrency | 0901–0904 | Concurrent transactions |
