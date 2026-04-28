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
          version = "2026.04.28-a";

          src = self;

          vendorHash = "sha256-psfVdzWz3jU+QEliVA2dPY5nZDy2HMFJ0B9XP25jjxU=";

          subPackages = [ "." ];
          ldflags = [ "-s" "-w" ];

          meta = with pkgs.lib; {
            description = "A minimal to-do list program with a few amenities";
            license = licenses.mit;
            platforms = platforms.all;
          };
        };

        apps.fastcards = {
          type = "app";
          program = "${self.packages.${system}.fastcards}/bin/fastcards";
        };
      });
}
