{
  description = "go-evercrypt";

  inputs = {
    nixpkgs.url = "nixpkgs/nixpkgs-unstable";
    utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, utils }: {
    overlays.default = final: _: {
      evercrypt = final.callPackage pkgs/by-name/ev/evercrypt/package.nix { };

      go-evercrypt = final.callPackage
        ({ buildGoModule, evercrypt }: buildGoModule {
          pname = "go-evercrypt";
          version = "0.1.0";
          src = builtins.path { path = ./..; name = "go-evercrypt-src"; };
          vendorHash = null;
          buildInputs = [ evercrypt ];
        })
        { };
    };
  } // utils.lib.eachDefaultSystem (system:
    let
      pkgs = import nixpkgs {
        overlays = [ self.overlays.default ];
        inherit system;
      };
      inherit (pkgs) evercrypt go-evercrypt gopls mkShell;
    in
    {
      packages = { inherit evercrypt; };

      devShells.default = mkShell {
        inputsFrom = [ go-evercrypt ];
        packages = [ gopls ];
      };
    });
}
