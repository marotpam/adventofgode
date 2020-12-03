package _2019

import (
	"github.com/marotpam/adventofgode/2019/intcode"
)

type Packet struct {
	x, y int
}

type packetSender interface {
	send(id int, packet Packet)
}

type robot struct {
	id            int
	interpreter   *intcode.Interpreter
	input, output []int
	packetSender  packetSender
}

func newRobot(id int, network *network) *robot {
	return &robot{
		id:           id,
		input:        []int{id},
		output:       make([]int, 0),
		interpreter:  intcode.NewInterpreter(),
		packetSender: network,
	}
}

func (r *robot) Write(n int) {
	r.output = append(r.output, n)
	if len(r.output) == 3 {
		p := Packet{
			x: r.output[1],
			y: r.output[2],
		}
		r.packetSender.send(r.output[0], p)
		r.output = r.output[3:]
	}
}

func (r *robot) Read() int {
	x := -1

	if len(r.input) > 0 {
		x, r.input = r.input[0], r.input[1:]
	}

	return x
}

func (r *robot) run(instructions []int) {
	for {
		r.interpreter.Run(instructions, r, r)
	}
}

type network struct {
	robots map[int]*robot
	done   chan Packet
}

func (n network) send(id int, packet Packet) {
	r, ok := n.robots[id]
	if !ok {
		n.done <- packet
		return
	}
	r.input = append(r.input, packet.x, packet.y)
}

func newNetwork(robotCount int) *network {
	network := &network{
		done: make(chan Packet, 1),
	}

	robots := make(map[int]*robot, robotCount+1)
	for i := 0; i < robotCount; i++ {
		robots[i] = newRobot(i, network)
	}
	network.robots = robots

	return network
}

func GetFirstPacketSentToNAT(instructions []int) Packet {
	n := newNetwork(50)

	for _, r := range n.robots {
		prog := append([]int{}, instructions...)
		go r.run(prog)
	}

	return <-n.done
}
