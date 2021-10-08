package job

// CLIConfig contains all possible configurations required by the Nomad Pack
// CLI in order to render, plan, run, and destroy job templates.
type CLIConfig struct {
	RunConfig     *RunCLIConfig
}

// RunCLIConfig specifies the configuration that is used by the Nomad Pack run
// command.
type RunCLIConfig struct {
	CheckIndex      uint64
	ConsulToken     string
	ConsulNamespace string
	VaultToken      string
	VaultNamespace  string
	EnableRollback  bool
	HCL1            bool
	PreserveCounts  bool
	PolicyOverride  bool
}