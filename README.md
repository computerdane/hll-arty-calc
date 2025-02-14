# hll-arty-calc

Calculate artillery angles for the game Hell Let Loose.

# Usage

Use interactive mode:

```sh
  hll-arty-calc
```

Specify a team:

```sh
  hll-arty-calc germany
```

Shorthand:

```sh
  hll-arty-calc g
```

With distances (non-interactive):

```sh
  hll-arty-calc g 100 200 300
  # Example output:
  # 978
  # 954
  # 931
```

# Install

## Nix

Use in a Nix shell:

```nix
  nix shell github:computerdane/hll-arty-calc
  hll-arty-calc
```

Run without cloning:

```nix
  nix run github:computerdane/hll-arty-calc
```

Run from the cloned repo:

```nix
  nix run .
```

# Develop

Install [Nix](https://nixos.org/), and then you can use `nix develop` to enter a shell with the exact Go version needed to work on this project.
