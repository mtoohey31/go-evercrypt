{ fetchFromGitHub, lib, stdenv, which }:

stdenv.mkDerivation rec {
  pname = "evercrypt";
  version = lib.strings.substring 0 7 src.rev;

  src = fetchFromGitHub {
    owner = "hacl-star";
    repo = "hacl-star";
    rev = "8904da86656cbb5d14b284ef359a57b2970864f7";
    hash = "sha256-/zp4fnFF6Em+pQJLjVGixmUrb5QjtSoWIBP27oeDpZc=";
  };
  sourceRoot = "source/dist/gcc-compatible";

  dontAddPrefix = true;

  installPhase = ''
    runHook preInstall

    mkdir -p $out/{include,lib}
    cp libevercrypt.a libevercrypt.so $out/lib

    for header in $(grep -Po '(?<=ALL_H_FILES=).*' Makefile.include) libintvector.h; do
      mkdir -p $(dirname $out/include/$header)
      cp $header $out/include/$header
    done

    pushd ../karamel/include
    for header in $(find . -type f); do
      mkdir -p $(dirname $out/include/$header)
      cp $header $out/include/$header
    done
    popd

    cp ../karamel/krmllib/dist/minimal/*.h $out/include

    runHook postInstall
  '';

  nativeBuildInputs = [ which ];

  meta = {
    homepage = "https://hacl-star.github.io";
    description = "A formally verified cryptographic library written in F*";
  };
}
