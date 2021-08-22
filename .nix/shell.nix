{ pkgs ? (import ./pkgs.nix {}) }:

# minitime is a mini-output time wrapper.
let minitime = pkgs.writeShellScriptBin "minitime"
		"command time --format $'%C -> %es\\n' \"$@\"";

	generate = pkgs.writeShellScriptBin "generate"
		"go generate";

	build = pkgs.writeShellScriptBin "build" ''
		cd pkg
		go build -v ./...
	'';

	publishCachix = with pkgs; writeShellScriptBin "publish-cachix" ''
		nix-build shell.nix -A inputDerivation | ${cachix}/bin/cachix push gotk4
	'';

in pkgs.stdenv.mkDerivation {
	name = "gotk4-shell";
	phases = [ "noopPhase" ];
	noopPhase = ":";

	# Runtime dependencies.
	buildInputs = with pkgs; [
		# Runtime dependencies.
		gobjectIntrospection
		glib
		graphene
		gdk-pixbuf
		gnome3.gtk
		gtk4
		vulkan-headers
	];

	nativeBuildInputs = with pkgs; [
		# Build dependencies.
		pkgconfig
		go

		# Tools.
		minitime
		generate
		build
		cachix
		publishCachix
	];

	CGO_ENABLED = "1";

	# Use /tmp, since /run/user/1000 (XDG_RUNTIME_DIRECTORY) might be too small.
	# See https://github.com/NixOS/nix/issues/395.
	TMP    = "/tmp";
	TMPDIR = "/tmp";
}
