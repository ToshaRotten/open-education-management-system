package image_keeper

type ImageKeeper struct {
	DockerImagesPath string
	CompleteRepoPath string
	ZipSourcePath    string
}

func New() *ImageKeeper {
	var k ImageKeeper
	k.DockerImagesPath = "../data/complete_data/docker_images/"
	k.CompleteRepoPath = "../data/complete_data/repo/"
	k.ZipSourcePath = "../data/source_data/zip/"
	return &k
}

func (k *ImageKeeper) CreateCompleteRepo(version string) {
	readZip(k.ZipSourcePath+"source_"+version+".zip", k.CompleteRepoPath+"/"+version)
}

func (k *ImageKeeper) CreateCompleteImages(version string) {
	buildDockerImages(k.CompleteRepoPath + "/" + version)
}
