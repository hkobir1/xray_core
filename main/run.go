package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"runtime/debug"
	"strings"
	"syscall"
	"time"

	"github.com/hkobir1/xray_core/common/cmdarg"
	"github.com/hkobir1/xray_core/common/errors"
	clog "github.com/hkobir1/xray_core/common/log"
	"github.com/hkobir1/xray_core/common/platform"
	"github.com/hkobir1/xray_core/core"
	"github.com/hkobir1/xray_core/main/commands/base"
)

var cmdRun = &base.Command{
	UsageLine: "{{.Exec}} run [-c config.json] [-confdir dir]",
	Short:     "Run Xray with config, the default command",
	Long: `
Run Xray with config, the default command.

The -config=file, -c=file flags set the config files for 
Xray. Multiple assign is accepted.

The -confdir=dir flag sets a dir with multiple json config

The -format=json flag sets the format of config files. 
Default "auto".

The -test flag tells Xray to test config files only, 
without launching the server.

The -dump flag tells Xray to print the merged config.
	`,
}

func init() {
	cmdRun.Run = executeRun // break init loop
}

var (
	configFiles cmdarg.Arg
	configDir   string
	dump        = cmdRun.Flag.Bool("dump", false, "Dump merged config only, without launching Xray server.")
	test        = cmdRun.Flag.Bool("test", false, "Test config file only, without launching Xray server.")
	format      = cmdRun.Flag.String("format", "auto", "Format of input file.")

	_ = func() bool {
		cmdRun.Flag.Var(&configFiles, "config", "Config path for Xray.")
		cmdRun.Flag.Var(&configFiles, "c", "Short alias of -config")
		cmdRun.Flag.StringVar(&configDir, "confdir", "", "A dir with multiple json config")

		return true
	}()
)

func executeRun(cmd *base.Command, args []string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if *dump {
		clog.ReplaceWithSeverityLogger(clog.Severity_Warning)
		errCode := dumpConfig()
		os.Exit(errCode)
	}

	printVersion()
	server, err := startXray(ctx)
	if err != nil {
		fmt.Println("Failed to start:", err)
		os.Exit(23)
	}

	if *test {
		fmt.Println("Configuration OK.")
		os.Exit(0)
	}

	if err := server.Start(); err != nil {
		fmt.Println("Failed to start:", err)
		os.Exit(-1)
	}
	defer server.Close()

	runtime.GC()
	debug.FreeOSMemory()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	select {
	case <-sigChan:
		cancel()
	case <-ctx.Done():
	}

	server.Close()
}

func dumpConfig() int {
	files := getConfigFilePath(false)
	if config, err := core.GetMergedConfig(files); err != nil {
		fmt.Println(err)
		time.Sleep(1 * time.Second)
		return 23
	} else {
		fmt.Print(config)
	}
	return 0
}

func fileExists(file string) bool {
	info, err := os.Stat(file)
	return err == nil && !info.IsDir()
}

func dirExists(file string) bool {
	if file == "" {
		return false
	}
	info, err := os.Stat(file)
	return err == nil && info.IsDir()
}

func getRegepxByFormat() string {
	switch strings.ToLower(*format) {
	case "json":
		return `^.+\.(json|jsonc)$`
	case "toml":
		return `^.+\.toml$`
	case "yaml", "yml":
		return `^.+\.(yaml|yml)$`
	default:
		return `^.+\.(json|jsonc|toml|yaml|yml)$`
	}
}

func readConfDir(dirPath string) {
	confs, err := os.ReadDir(dirPath)
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range confs {
		matched, err := regexp.MatchString(getRegepxByFormat(), f.Name())
		if err != nil {
			log.Fatalln(err)
		}
		if matched {
			configFiles.Set(path.Join(dirPath, f.Name()))
		}
	}
}

func getConfigFilePath(verbose bool) cmdarg.Arg {
	if dirExists(configDir) {
		if verbose {
			log.Println("Using confdir from arg:", configDir)
		}
		readConfDir(configDir)
	} else if envConfDir := platform.GetConfDirPath(); dirExists(envConfDir) {
		if verbose {
			log.Println("Using confdir from env:", envConfDir)
		}
		readConfDir(envConfDir)
	}

	if len(configFiles) > 0 {
		return configFiles
	}

	if workingDir, err := os.Getwd(); err == nil {
		configFile := filepath.Join(workingDir, "config.json")
		if fileExists(configFile) {
			if verbose {
				log.Println("Using default config: ", configFile)
			}
			return cmdarg.Arg{configFile}
		}
	}

	if configFile := platform.GetConfigurationPath(); fileExists(configFile) {
		if verbose {
			log.Println("Using config from env: ", configFile)
		}
		return cmdarg.Arg{configFile}
	}

	if verbose {
		log.Println("Using config from STDIN")
	}
	return cmdarg.Arg{"stdin:"}
}

func getConfigFormat() string {
	f := core.GetFormatByExtension(*format)
	if f == "" {
		f = "auto"
	}
	return f
}

func startXray(ctx context.Context) (core.Server, error) {
	configFiles := getConfigFilePath(true)
	c, err := core.LoadConfig(getConfigFormat(), configFiles)
	if err != nil {
		return nil, errors.New("failed to load config files: [", configFiles.String(), "]").Base(err)
	}

	server, err := core.New(c)
	if err != nil {
		return nil, errors.New("failed to create server").Base(err)
	}

	go func() {
		<-ctx.Done()
		server.Close()
	}()

	return server, nil
}
