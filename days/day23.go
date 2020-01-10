package days

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fivegreenapples/AOC2015/utils"
)

func (r *Runner) Day23Part1(in string) string {
	instructions := utils.Lines(in)
	vm := d23NewVM(instructions, r.verbose)
	vm.run()

	return strconv.Itoa(int(vm.registerVal("b")))
}

func (r *Runner) Day23Part2(in string) string {
	instructions := utils.Lines(in)
	vm := d23NewVM(instructions, r.verbose)
	vm.setRegisterVal("a", 1)
	vm.run()

	return strconv.Itoa(int(vm.registerVal("b")))
}

type d23VM struct {
	pc           int
	registers    map[string]uint
	instructions []string
	verbose      bool
}

func d23NewVM(instructions []string, verbose bool) *d23VM {

	return &d23VM{
		pc:           0,
		registers:    map[string]uint{},
		instructions: instructions,
		verbose:      verbose,
	}

}

func (vm *d23VM) run() {

	for {
		if vm.pc < 0 || vm.pc >= len(vm.instructions) {
			break
		}
		instr := vm.instructions[vm.pc]
		if vm.verbose {
			fmt.Println(vm.pc, instr, vm.registers["a"], vm.registers["b"])
		}
		bits := strings.Split(instr, " ")
		switch bits[0] {
		case "hlf":
			vm.registers[bits[1]] /= 2
			vm.pc++
		case "tpl":
			vm.registers[bits[1]] *= 3
			vm.pc++
		case "inc":
			vm.registers[bits[1]]++
			vm.pc++
		case "jmp":
			vm.pc += utils.MustAtoi(bits[1])
		case "jie":
			if vm.registers[string(bits[1][0])]%2 == 0 {
				vm.pc += utils.MustAtoi(bits[2])
			} else {
				vm.pc++
			}
		case "jio":
			if vm.registers[string(bits[1][0])] == 1 {
				vm.pc += utils.MustAtoi(bits[2])
			} else {
				vm.pc++
			}
		default:
			panic("unhandled instruction:" + bits[0])
		}

	}

}

func (vm *d23VM) registerVal(which string) uint {
	return vm.registers[which]
}
func (vm *d23VM) setRegisterVal(which string, val uint) {
	vm.registers[which] = val
}
