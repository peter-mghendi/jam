[![Build + Test](https://github.com/peter-mghendi/jam/actions/workflows/build.yml/badge.svg)](https://github.com/peter-mghendi/jam/actions/workflows/build.yml)

# jam
> **Alias manager for your shell**  
> Keep all your aliases in one structured, machine-managed file: `~/.jamrc`.

---

## What is jam?

jam is a small CLI tool that manages your shell aliases through a single declarative file: `~/.jamrc`.

Instead of scattering `alias` lines across `.bashrc`, `.zshrc`, and random scripts, jam treats `~/.jamrc` as the **source of truth** and keeps it consistent using a parser-aware, structured format under the hood.

jam:

- Stores aliases + metadata in a machine-friendly format.
- Regenerates `~/.jamrc` safely after every change using a robust AST-aware parser.
- Lets you preview changes before writing.

---

## Status

jam is experimental and currently supports:

- `init` - create a new `~/.jamrc`
- `debug` - print the parsed AST of `~/.jamrc`
- `list` - list aliases managed by jam
- `add` - add a new alias
- `remove` - remove an alias
- `enable` - mark an alias as enabled
- `disable` - mark an alias as disabled

More commands (export, edit, etc.) may come later.

---

## Installation

### Pre-built binaries

> While binaries are available for multiple operating systems, `jam` currently only supports bash and bash-compatible shells.

Grab the latest release from the [jam GitHub Releases](https://github.com/peter-mghendi/jam/releases).

Confirm:

```bash
jam --help
```

> It may help to put the binary on your `PATH` for easy access.

### From source

```bash
git clone https://github.com/peter-mghendi/jam.git
cd jam
go install ./cmd/jam
````

This should put `jam` in your `$GOBIN` (often `$HOME/go/bin`).

Make sure it’s on your `PATH`:

```bash
export PATH="$HOME/go/bin:$PATH"
```

Confirm:

```bash
jam --help
```

---

## How it works

* jam manages a single file: `~/.jamrc`
* Each alias has:

  * a **name** (what you type),
  * a **target** (what runs),
  * optional **description**,
  * **enabled/disabled** state.
* Internally, jam uses Bash’s own features and a real AST parser to keep the file well-structured.

You’ll typically:

1. Initialize `~/.jamrc` once.
2. Add / remove / enable / disable aliases using `jam`.
3. Source `~/.jamrc` from your shell config:

   ```bash
   # in ~/.bashrc or ~/.zshrc
   if [ -f "$HOME/.jamrc" ]; then
     . "$HOME/.jamrc"
   fi
   ```

---

## Usage

### `jam init`

Create a new, empty `~/.jamrc`.

```bash
# Create ~/.jamrc if it does not exist
jam init

# Overwrite existing ~/.jamrc
jam init --force

# Show what would be written, without touching the file
jam init --pretend
```

---

### `jam list`

List aliases managed by jam.

```bash
jam list
jam ls
```

Typical output might look like:

```text
# Says hello
# Added at: 2025-11-16 07:20:30 +0300 UTC
alias greet="$HOME/bin/greet.cs"

# Removes a directory
# Added at: 2025-11-17 02:13:01 +0300 UTC
# alias yeet="rm -rf --no-preserve-root"
```

---

### `jam add`

Add a new alias.

```bash
# Basic: name + target
jam add greet "$HOME/bin/greet.cs"

# With description
jam add greet "$HOME/bin/greet.cs" --desc "Says hello"

# Start disabled
jam add greet "$HOME/bin/greet.cs" --desc "Says hello" --disabled

# Preview changes without writing
jam add greet "$HOME/bin/greet.cs" --pretend
```

Flags:

* `--desc string` - optional description.
* `--disabled` - create the alias in a disabled state (default is enabled).
* `--pretend` - print the updated `.jamrc` to STDOUT without writing it.

---

### `jam remove`

Remove an existing alias.

```bash
# Remove the alias
jam remove greet

# Preview the updated .jamrc without writing
jam remove greet --pretend
```

This deletes both the alias definition and its metadata from `~/.jamrc`.
It does **not** delete any script or binary on disk.

---

### `jam enable`

Enable an alias.

```bash
jam enable greet
jam enable greet --pretend
```

This updates the metadata for `greet` to mark it as enabled.

---

### `jam disable`

Disable an alias.

```bash
jam disable greet
jam disable greet --pretend
```

This updates the metadata for `greet` to mark it as disabled. How this is reflected in the generated `.jamrc` is handled internally by jam.

---

### `jam debug`

Dump the parsed AST of `~/.jamrc` (for debugging jam itself or your file).

```bash
jam debug
```

This uses the underlying shell parser to print a structured view of the file.
It’s mainly aimed at development or troubleshooting.

---

## Notes & Assumptions

* jam currently assumes a Unix-like environment (tested primarily on Linux).
* `~/.jamrc` is considered **owned by jam**:

  * You can open and read it.
  * You *can* edit it by hand, but jam may overwrite your formatting.
* It’s recommended that you use jam commands (`add`, `remove`, `enable`, `disable`) for changes instead of manual editing.

---

## Design

* `.jamrc` is intended to be fully compatible with bash 4.0+. A sample `.jamrc` may look like:

```shell
#!/usr/bin/env bash

# jam: alias manager
# This file was generated by jam.
# Edit at your own risk - manual edits may be overwritten.
# Use `jam add`, `jam enable`, `jam disable`, etc. to modify aliases.

declare -A __jam__greet=([target]="$HOME/bin/greet.cs" [enabled]="true" [description]="Says hello" [added_at]="2025-11-16T07:20:30Z")
alias greet="$HOME/bin/greet.cs"

declare -A __jam__yeet=([target]="rm -rf random_dir" [enabled]="false" [description]="Lalala" [added_at]="2025-11-17T02:12:24Z") 
# alias yeet="rm -rf random_dir"
```

* jam uses bash associative arrays to store metadata, and it's possible to view alias metadata using bash:

```shell
# metadata is stored in a dictionary named __jam__<alias_name>
declare -p "__jam__greet"
```

## FAQ

**1. Can I change the location of `.jamrc`?**  
Not at the moment. jam will always look in `$HOME/.jamrc`

**2. Why?**  
Yes. Next question.

**3. What happens if `.jamrc` already exists/does not exist?**  
By default, jam will not perform any destructive actions and instead display a helpful error message. Follow the prompts to overwrite the existing file or create a new one.

**4. What does this cost?**  
Nothing. Can you imagine?
---

## License
[MIT License](LICENSE) © 2025 Peter Mghendi

