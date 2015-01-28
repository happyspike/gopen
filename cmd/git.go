package cmd

func GitBranch() (string, error) {
	return Exec("git rev-parse --abbrev-ref HEAD")
}

func GitStatus() (string, error) {
	return Exec("git status")
}

func GitOriginUrl() (string, error) {
	return Exec("git config --get remote.origin.url")
}
