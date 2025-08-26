package core

const (
	ProgramName = "pie"
)
const (
	ColorReset  = "\033[0m"
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
)
const (
	StatusOK    = ColorGreen + "[+]" + ColorReset
	StatusWarn  = ColorYellow + "[!]" + ColorReset
	StatusError = ColorRed + "[-]" + ColorReset
)

const (
	UserDirectory      = "/data/" + ProgramName + "/"
	RepositoryFileName = UserDirectory + "repository.json"
	ModulesPath        = "/data/adb/modules"
	LogFileName        = UserDirectory + ProgramName + ".log"
	LockFile           = UserDirectory + ProgramName + "pie.lock"
	CacheDirectory     = UserDirectory + "cache/"
)

const (
	InstallCharacter    = "-I"
	RemoveCharacter     = "-R"
	UpdateRepoCharacter = "u"
	AutoPromptCharacter = "p"
	SearchCharacter     = "-S"
)
