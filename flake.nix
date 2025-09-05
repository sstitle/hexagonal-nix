{
  description = "Development environment with nickel and mask";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
    flake-parts.url = "github:hercules-ci/flake-parts";
    treefmt-nix.url = "github:numtide/treefmt-nix";
  };

  outputs =
    inputs@{
      flake-parts,
      ...
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [
        "x86_64-linux"
        "aarch64-linux"
        "x86_64-darwin"
        "aarch64-darwin"
      ];

      perSystem =
        { config, self', inputs', pkgs, system, ... }:
        let
          treefmtEval = inputs.treefmt-nix.lib.evalModule pkgs ./treefmt.nix;
        in
        {
          # Go package for user profile system
          packages.hello = pkgs.writeShellApplication {
            name = "hello";
            text = ''
              echo "Hello, world from Flake!"
            '';
          };
          # App runner to execute the Go CLI (driving adapter)
          packages.greet = pkgs.writeShellApplication {
            name = "greet";
            runtimeInputs = [ pkgs.go ];
            text = ''
              # run from repo root
              exec go run ./src/adapters/driving "$@"
            '';
          };
          packages.default = self'.packages.hello;

          # nix run . [args] will run the Go CLI
          apps.greet = {
            type = "app";
            program = "${self'.packages.greet}/bin/greet";
          };
          apps.default = self'.apps.greet;

          # Development shell with nickel and mask
          devShells.default = pkgs.mkShell {
            buildInputs = with pkgs; [
              # Core tools
              git
              nickel
              mask
              
              # Go development
              go
              gopls
              gotools
              go-tools
            ];

            shellHook = ''
              echo "🚀 Development environment loaded!"
              echo "Available tools:"
              echo "  - nickel: Configuration language"
              echo "  - mask: Task runner"
              echo "  - go: Go development"
              echo "  - gopls: Go language server"
              echo ""
              echo "Run 'mask --help' to see available tasks."
              echo "Run 'nix fmt' to format all files."
              echo "Run 'mask greet Alice' to demo the Go hexagonal example."
              echo "Run 'nix run . -- Alice' to run via flake app."
            '';
          };

          # for `nix fmt`
          formatter = treefmtEval.config.build.wrapper;

          # for `nix flake check`
          checks = {
            formatting = treefmtEval.config.build.check inputs.self;
          };
        };
    };
}
