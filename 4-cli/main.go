package main

import (
	"cli/api"
	"cli/config"
	"cli/files"
	"cli/operations"
	"cli/storage"
	"flag"
	"fmt"
	"os"
)

const storagePath = "./storage.json"

func usage() {
	fmt.Fprintln(os.Stderr, "Необходимо указать команду: --create, --update, --delete, --get или --list")
}

type commandMap = map[string]func() error

func setupCommands(ops *operations.Operations) commandMap {
	commands := commandMap{}

	var createCmd = flag.NewFlagSet("--create", flag.ExitOnError)
	commands[createCmd.Name()] = func() error {
		file := createCmd.String("file", "", "файл для загрузки")
		name := createCmd.String("name", "", "имя bin")
		err := createCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		if *file == "" {
			return fmt.Errorf("необходимо указать файл для загрузки")
		}
		if *name == "" {
			return fmt.Errorf("необходимо указать имя bin")
		}
		return ops.Create(*file, *name)
	}
	var updateCmd = flag.NewFlagSet("--update", flag.ExitOnError)
	commands[updateCmd.Name()] = func() error {
		file := updateCmd.String("file", "", "файл для загрузки")
		id := updateCmd.String("id", "", "id bin")
		err := updateCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		if *file == "" {
			return fmt.Errorf("необходимо указать файл для загрузки")
		}
		if *id == "" {
			return fmt.Errorf("необходимо указать id bin")
		}
		return ops.Update(*file, *id)
	}
	var deleteCmd = flag.NewFlagSet("--delete", flag.ExitOnError)
	commands[deleteCmd.Name()] = func() error {
		id := deleteCmd.String("id", "", "id bin")
		err := deleteCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		if *id == "" {
			return fmt.Errorf("необходимо указать id bin")
		}
		return ops.Delete(*id)
	}
	var getCmd = flag.NewFlagSet("--get", flag.ExitOnError)
	commands[getCmd.Name()] = func() error {
		id := getCmd.String("id", "", "id bin")
		err := getCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		if *id == "" {
			return fmt.Errorf("необходимо указать id bin")
		}
		return ops.Get(*id)
	}
	var listCmd = flag.NewFlagSet("--list", flag.ExitOnError)
	commands[listCmd.Name()] = func() error {
		err := getCmd.Parse(os.Args[2:])
		if err != nil {
			return err
		}
		return ops.List()
	}
	return commands
}

func main() {
	config, err := config.NewConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка получения конфига: %v\n", err)
		os.Exit(1)
	}
	jsonBin := api.NewJsonBin(config)
	fileProvider, err := files.NewJsonFile(storagePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка инициализации провайдера: %v", err)
		os.Exit(1)
	}
	storage := storage.NewStorage(fileProvider)
	binList, err := storage.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Ошибка чтения списка бинов: %v", err)
		os.Exit(1)
	}
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	ops := operations.NewOperation(jsonBin, binList)
	commands := setupCommands(ops)
	command := commands[os.Args[1]]
	if command == nil {
		usage()
		os.Exit(1)
	}
	err = command()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = storage.Save(binList)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
