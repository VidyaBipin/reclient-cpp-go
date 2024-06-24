// Copyright 2023 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package features defines features enabled conditionally via flags.
package features

import (
	"flag"
)

// Config is the feature configuration in use.
type Config struct {
	// CleanIncludePaths enables the cleaning of include paths, which
	// involves checking for absolute paths and making them relative
	// to the working directory.
	// It is temporary feature for nest build. b/157442013
	CleanIncludePaths bool

	// ExperimentalCacheMissRate is the cache miss rate simulated by
	// reproxy in an experimental build. Not to be used for production.
	ExperimentalCacheMissRate int

	// ExperimentalGomaDepsCache enables using the reproxy deps cache
	// with gomaIP instead of goma's deps cache. This has no effect on
	// clangscandeps
	ExperimentalGomaDepsCache bool

	// ExperimentalGomaDepsCacheSize is the maximum number of elements to
	// be stored in the cache.
	ExperimentalGomaDepsCacheSize int

	// ExperimentalSysrootDoNotUpload disables upload of the files/directories
	// under the directory specified by the --sysroot flag.
	ExperimentalSysrootDoNotUpload bool

	// ExperimentalExitOnStuckActions shutdowns reproxy with exit code 1 if rewrapper's command
	// didn't finish within 2*reclient_timeout. This is meant to generate the call stack when
	// closing reproxy to pin down where the execution gets stuck
	//TODO(b/284246561) Remove once we have a confirmation where action execution gets stuck
	ExperimentalExitOnStuckActions bool

	// If false, disables the credential cache even if the auth mechanism claims it is cachable.
	// It is enabled by default.
	EnableCredentialCache bool
}

var config = &Config{}

// GetConfig retrieves the singleton instance of the features config.
func GetConfig() *Config {
	return config
}

func init() {
	flag.Bool("shadow_header_detection", false, "Indicates whether to enable detection of shadow headers when building/verifying dependencies of c++ compilations in local execution remote cache mode")
	flag.BoolVar(&GetConfig().CleanIncludePaths, "clean_include_paths", false, "Indicates whether to clean include paths from -I arguments")
	flag.IntVar(&GetConfig().ExperimentalCacheMissRate, "experimental_cache_miss_rate", 0, "Indicates percent of actions to simulate cache misses. Integer [0,100).")
	flag.BoolVar(&GetConfig().ExperimentalSysrootDoNotUpload, "experimental_sysroot_do_not_upload", false, "Do not upload the the files/directories under the directory specified by the --sysroot flag.")
	flag.BoolVar(&GetConfig().ExperimentalGomaDepsCache, "experimental_goma_deps_cache", false, "Use go deps cache with goma instead of goma's deps cache")
	flag.BoolVar(&GetConfig().ExperimentalExitOnStuckActions, "experimental_exit_on_stuck_actions", false, "Stops reproxy with exit_code=1 if the command didn't finish within 2*reclient_timeout")
	flag.BoolVar(&GetConfig().EnableCredentialCache, "enable_creds_cache", true, "If false, disables the credentials cache even if used auth mechanism supports it")
	flag.IntVar(&GetConfig().ExperimentalGomaDepsCacheSize, "experimental_goma_deps_cache_size", 300000, "Maximum number of entries to hold in the experimental deps cache.")
}
