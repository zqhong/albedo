package constant

var Path *path

type path struct {
	RootDir string
	ConfDir string
}

func init() {
	Path = &path{}
}

func SetRootDir(root string) {
	Path.RootDir = root
}

func SetConfDir(conf string) {
	Path.ConfDir = conf
}
