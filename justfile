testing_dir := "testdata/basic-workspace/scenarios/tools/list.toml"

[group("helpers")]
list:
  @just --list

[group("main")]
build:
  @go build -o ./bin/smriti .

[group("main")]
go cmd:
  @./bin/smriti {{cmd}}

[group("cli")]
run path:
  @./bin/smriti run {{path}}

[group("cli")]
discover:
  @./bin/smriti discover

[group("cli")]
version:
  @./bin/smriti version
