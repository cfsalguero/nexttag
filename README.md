# nexttag

`nexttag` prints the `git tag` command for the next semantic version in a Git
repository. Run it inside any repo to bump its latest `vX.Y.Z` tag.

## How it works

1. Lists tags with `git tag`.
2. Finds the highest tag matching the pattern `vX.Y.Z` (e.g. `v1.2.3`).
3. Bumps the version according to the argument you pass.
4. Prints the command to create the next tag, e.g.:

```
git tag release/v1.2.4
```

If the repo has no matching tags, it starts from `v0.0.0`.

## Installation

```sh
go install github.com/cfsalguero/nexttag@latest
```

Or build from source:

```sh
go build -o nexttag .
```

## Usage

```sh
nexttag <x|y|z>
```

| Arg | Bump  | Example (`v1.2.3` -> next) |
|-----|-------|----------------------------|
| `z` | patch | `v1.2.4`                   |
| `y` | minor | `v1.3.0`                   |
| `x` | major | `v2.0.0`                   |

Run it inside any Git repository:

```sh
cd /path/to/your/repo
nexttag z
# git tag release/v1.2.4
```

To create the tag directly:

```sh
$(nexttag z)
```

## Examples

Starting from the highest tag `v1.2.3`:

```sh
$ nexttag z
git tag release/v1.2.4

$ nexttag y
git tag release/v1.3.0

$ nexttag x
git tag release/v2.0.0
```

Only tags matching `vX.Y.Z` are considered; others (like `release/v1.2.3` or
`foo`) are ignored.
