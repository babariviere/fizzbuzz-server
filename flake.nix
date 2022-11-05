{
  description = "A very basic flake";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }:
    utils.lib.eachDefaultSystem (system:
      let pkgs = import nixpkgs { inherit system; };
      in {
        packages = {
          default = self.packages.${system}.fizzbuzz-server;
          fizzbuzz-server = pkgs.buildGoModule {
            name = "fizzbuzz-server";
            src = ./.;
            vendorSha256 = "sha256-pQpattmS9VmO3ZIQUFn66az8GSmB4IvYhTTCFn6SUmo=";
          };
        };

        apps = {
          default = self.apps.${system}.fizzbuzz-server;
          fizzbuzz-server = {
            type = "app";
            program =
              "${self.packages.${system}.fizzbuzz-server}/bin/fizzbuzz-server";
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [ go gotools go-tools gopls ];
        };
      });
}
