package main

import "github.com/codegangsta/cli"

var flagAgentUuid = cli.StringFlag{"agent-uuid", "", "the uuid of the agent", "GOCD_AGENT_UUID"}

var agent = cli.Command{
	Name:  "agent",
	Usage: "commands for agents",
	Subcommands: []cli.Command{
		{
			Name:   "list",
			Usage:  "list the connected agents",
			Flags:  flags,
			Action: agentList,
		},
		{
			Name:   "enable",
			Usage:  "enable a specific agent",
			Flags:  append(flags, flagAgentUuid),
			Action: agentEnable,
		},
		{
			Name:   "disable",
			Usage:  "disable a specific agent",
			Flags:  append(flags, flagAgentUuid),
			Action: agentDisable,
		},
		{
			Name:   "delete",
			Usage:  "delete a specific agent",
			Flags:  append(flags, flagAgentUuid),
			Action: agentDelete,
		},
	},
}

func agentList(c *cli.Context) {
	client := newClient(c)
	agents, err := client.AgentList()
	assert(err)
	print(agents)
}

func agentEnable(c *cli.Context) {
	client := newClient(c)
	uuid := c.String("agent-uuid")
	err := client.AgentEnable(uuid)
	assert(err)

	agentList(c)
}

func agentDisable(c *cli.Context) {
	client := newClient(c)
	uuid := c.String("agent-uuid")
	err := client.AgentDisable(uuid)
	assert(err)

	agentList(c)
}

func agentDelete(c *cli.Context) {
	client := newClient(c)
	uuid := c.String("agent-uuid")
	err := client.AgentDelete(uuid)
	assert(err)

	agentList(c)
}
