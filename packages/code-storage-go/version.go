package storage

const PackageName = "code-storage-go-sdk"
const PackageVersion = "0.2.0"

func userAgent() string {
	return PackageName + "/" + PackageVersion
}
