testing_dir := "testbench/basic-workspace/scenarios/tools/list.toml"

[group("aux")]
list:
  @just --list

[group("build")]
build:
  @go build -o ./bin/smriti .

[group("cli")]
run testfile:
  @./bin/smriti run {{testfile}}

# Not implemented!
[group("cli")]
discover:
  @./bin/smriti discover

[group("cli")]
version:
  @./bin/smriti version

# Not implemented!
[group("cli")]
servers:
  @./bin/smriti servers

# Not implemented!
[group("cli")]
scenarios:
  @./bin/smriti scenarios
