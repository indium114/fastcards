{
  description = "fastcards devshell and package";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = import nixpkgs { inherit system; };
      in {
        devShells.default = pkgs.mkShell {
          name = "fastcards-devshell";

          packages = with pkgs; [
            go
            gopls
            gotools
            delve
          ];
        };

        packages.fastcards = pkgs.buildGoModule {
          pname = "fastcards";
          version = "2026.01.30-a";

          src = self;

          vendorHash = "sha256-FBb8RFndx6wwdu08B95d+qas4VWs6Vr7UrXVGLyCW0g=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A minimal to-do list program with a few amenities";
            license = licenses.mit;
            platforms = platforms.linux;
          };
        };

        apps.default = {
          type = "app";
          program = "${self.packages.${system}.fastcards}/bin/fastcards";
        };
      });
}
