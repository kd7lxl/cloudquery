package cmd

import (
	"github.com/cloudquery/cloudquery/cli/internal/specs/v0"
	pbSpecs "github.com/cloudquery/plugin-pb-go/specs"
)

func CLIBackendToPbBackend(backend specs.Backend) pbSpecs.Backend {
	switch backend {
	case specs.BackendLocal:
		return pbSpecs.BackendLocal
	case specs.BackendNone:
		return pbSpecs.BackendNone
	default:
		panic("unknown backend " + backend.String())
	}
}

func CLIRegistryToPbRegistry(registry specs.Registry) pbSpecs.Registry {
	switch registry {
	case specs.RegistryGithub:
		return pbSpecs.RegistryGithub
	case specs.RegistryLocal:
		return pbSpecs.RegistryLocal
	case specs.RegistryGrpc:
		return pbSpecs.RegistryGrpc
	default:
		panic("unknown registry " + registry.String())
	}
}

func CLISchedulerToPbScheduler(scheduler specs.Scheduler) pbSpecs.Scheduler {
	switch scheduler {
	case specs.SchedulerDFS:
		return pbSpecs.SchedulerDFS
	case specs.SchedulerRoundRobin:
		return pbSpecs.SchedulerRoundRobin
	default:
		panic("unknown scheduler " + scheduler.String())
	}
}

// This converts CLI configuration to a source spec prior to V3 version
// when our spec wasn't decoupled from the over the wire protocol
func CLISourceSpecToPbSpec(spec specs.Source) pbSpecs.Source {
	return pbSpecs.Source{
		Name:                spec.Name,
		Version:             spec.Version,
		Path:                spec.Path,
		Registry:            CLIRegistryToPbRegistry(spec.Registry),
		Concurrency:         spec.Concurrency,
		TableConcurrency:    spec.TableConcurrency,
		ResourceConcurrency: spec.ResourceConcurrency,
		Tables:              spec.Tables,
		SkipTables:          spec.SkipTables,
		SkipDependentTables: spec.SkipDependentTables,
		Destinations:        spec.Destinations,
		Backend:             CLIBackendToPbBackend(spec.Backend),
		BackendSpec:         spec.BackendSpec,
		Scheduler:           CLISchedulerToPbScheduler(spec.Scheduler),
		Spec:                spec.Spec,
		DeterministicCQID:   spec.DeterministicCQID,
	}
}

func CLIWriteModeToPbWriteMode(writeMode specs.WriteMode) pbSpecs.WriteMode {
	switch writeMode {
	case specs.WriteModeAppend:
		return pbSpecs.WriteModeAppend
	case specs.WriteModeOverwrite:
		return pbSpecs.WriteModeOverwrite
	case specs.WriteModeOverwriteDeleteStale:
		return pbSpecs.WriteModeOverwriteDeleteStale
	default:
		panic("unknown write mode " + writeMode.String())
	}
}

func CLIMigrateModeToPbMigrateMode(migrateMode specs.MigrateMode) pbSpecs.MigrateMode {
	switch migrateMode {
	case specs.MigrateModeSafe:
		return pbSpecs.MigrateModeSafe
	case specs.MigrateModeForced:
		return pbSpecs.MigrateModeForced
	default:
		panic("unknown migrate mode " + migrateMode.String())
	}
}

func CLIPkModeToPbPKMode(pkMode specs.PKMode) pbSpecs.PKMode {
	switch pkMode {
	case specs.PKModeCQID:
		return pbSpecs.PKModeCQID
	case specs.PKModeDefaultKeys:
		return pbSpecs.PKModeDefaultKeys
	default:
		panic("unknown pk mode " + pkMode.String())
	}
}

func CLIDestinationSpecToPbSpec(spec specs.Destination) pbSpecs.Destination {
	return pbSpecs.Destination{
		Name:           spec.Name,
		Version:        spec.Version,
		Path:           spec.Path,
		Registry:       CLIRegistryToPbRegistry(spec.Registry),
		WriteMode:      CLIWriteModeToPbWriteMode(spec.WriteMode),
		MigrateMode:    CLIMigrateModeToPbMigrateMode(spec.MigrateMode),
		BatchSize:      spec.BatchSize,
		BatchSizeBytes: spec.BatchSizeBytes,
		Spec:           spec.Spec,
		PKMode:         CLIPkModeToPbPKMode(spec.PKMode),
	}
}
