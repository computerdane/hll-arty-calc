{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.11";
    utils.url = "github:numtide/flake-utils";
  };
  outputs =
    { nixpkgs, utils, ... }:
    utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
        hll-arty-calc = pkgs.callPackage ./default.nix { };
      in
      {
        devShell = pkgs.mkShell { buildInputs = [ pkgs.go ]; };
        packages.default = hll-arty-calc;
        apps.default = utils.lib.mkApp { drv = hll-arty-calc; };
      }
    );
}
