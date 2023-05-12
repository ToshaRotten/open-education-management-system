package image_keeper

type ImageKeeper struct {
	DockerImagesPath string
	CompleteRepoPath string
	ZipSourcePath    string
}

func New() *ImageKeeper {
	var k ImageKeeper
	k.DockerImagesPath = "complete_data/docker_images"
	k.CompleteRepoPath = "complete_data/repo"
	k.ZipSourcePath = "zip_source/"
	return &k
}

func (k *ImageKeeper) CreateCompleteRepo(version string) {
	readZip(k.ZipSourcePath+"source_"+version+".zip", k.CompleteRepoPath+"/"+version)
}

func (k *ImageKeeper) CreateCompleteImages(version string) {
	buildDockerImages(k.CompleteRepoPath + "/" + version)
}
