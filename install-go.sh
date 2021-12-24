# let us download a file with curl on Linux command line #
: "${VERSION:=1.17.5}" # go version
: "${ARCH=amd64}"      # go archicture
tmpDir="$HOME/tmp"
mkdir -p "$tmpDir"
tarFile="go${VERSION}.linux-${ARCH}.tar.gz"

curl -L "https://golang.org/dl/$tarFile" --output "$tmpDir/$tarFile"
tar -xf "$tmpDir/$tarFile" -C "$tmpDir/go-$VERSION"

#sudo chown -R root:root "$tmpDir/go-$VERSION"

#sudo mv -v "$tmpDir/go-$VERSION" /usr/local
