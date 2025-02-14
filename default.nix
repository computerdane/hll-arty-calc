{ buildGoModule }:

buildGoModule rec {
  pname = "hll-arty-calc";
  version = "1.0.0";
  src = ./.;
  vendorHash = null;
  ldflags = [ "-X main.Version=${version}" ];
}
