package filters

func Commands(command []string, useUsername bool, aliasList ...string) FilterOperand {
	var listCompiledCommands []FilterOperand
	for _, cmd := range command {
		listCompiledCommands = append(listCompiledCommands, Command(cmd, useUsername, aliasList...))
	}
	return func(values *DataFilter) bool {
		for _, cmd := range listCompiledCommands {
			if cmd(values) {
				return true
			}
		}
		return false
	}
}
