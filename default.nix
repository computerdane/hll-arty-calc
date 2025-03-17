{ buildGoModule }:

buildGoModule rec {
  pname = "hll-arty-calc";
  version = "1.0.1";
  src = ./.;
  vendorHash = null;
  ldflags = [ "-X main.Version=${version}" ];
}
