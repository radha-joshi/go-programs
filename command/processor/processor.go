package processor

type Command struct {
	value int
}

type Processor struct {
	value       int
	commandList []Command
}

func NewProcessor() *Processor {
	return &Processor{
		value:       0,
		commandList: make([]Command, 0),
	}
}

func (p *Processor) AddValue(value int) {
	p.commandList = append(p.commandList, Command{
		value: value,
	})
}

func (p *Processor) GetValue() int {
	result := p.value
	for _, command := range p.commandList {
		result += command.value
	}
	return result
}

func (p *Processor) Undo() {
	if len(p.commandList) > 0 {
		p.commandList = p.commandList[:len(p.commandList)-1]
	}
}
