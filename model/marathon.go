package model

type MarathonAppInfos struct {
	App *MarathonAppInfo `locationName:"app" type:"structure"`
}

type MarathonAppInfo struct {
	Id        *string            `locationName:"id", type:"string"`
	Cmd       *string            `locationName:"cmd", type:"string"`
	Args      *string            `locationName:"args", type:"string"`
	User      *string            `locationName:"user", type:"string"`
	Instance  *int64             `locationName:"instance", type:"integer"`
	Cpus      *float64           `locationName:"cpus", type:"float"`
	Men       *int64             `locationName:"mem", type:"integer"`
	Disk      *int64             `locationName:"disk", type:"integer"`
	Executor  *string            `locationName:"executor", type:"string"`
	Container *MarationContainer `locationName:"container" type:"structure"`
	Tasks     []*MarathonTask    `locationName:"tasks" type:"list"`
}

type MarationContainer struct {
	Type   *string         `locationName:"type", type:"string"`
	Docker *MarathonDocker `locationName:"docker", type:"structure"`
}

type MarathonDocker struct {
	Image          *string                `locationName:"image", type:"string"`
	Network        *string                `locationName:"network", type:"string"`
	PortMappings   []*MarathonPortMapping `locationName:"portmapping", type:"list"`
	Privileged     *bool                  `locationName:"privileged", type:"bool"`
	ForcePullImage *bool                  `locationName:"forcePullImage", type:"bool"`
	Parameters     []*DockerParameter     `locationName:"parameters", type:"list"`
}

type DockerParameter struct {
	Key   *string `locationName:"key", type:"string"`
	Value *string `locationName:"value", type:"string"`
}

type MarathonPortMapping struct {
	ContainerPort *int64  `locationName:"containerPort", type:"integer"`
	HostPort      *int64  `locationName:"hostPost", type:"integer"`
	ServicePort   *int64  `locationName:"servicePort", type:"integer"`
	Protocol      *string `locationName:"protocol", type:"string"`
}

type MarathonTask struct {
	Id        *string  `locationName:"id", type:"string"`
	Host      *string  `locationName:"host", type:"string"`
	Ports     []*int64 `locationName:"ports", type:"list"`
	StartedAt *string  `locationName:"startedAt" type:"string"`
	StagedAt  *string  `locationName:"stagedAt", type:"string"`
	Version   *string  `locationName:"version", type:"string"`
	AppId     *string  `locationName:"appId", type:"string"`
}

type MarathonApps struct {
	Apps []*MarathonApp `locationName:"apps" type:"list"`
}
type SingleMarathonApp struct {
	App *MarathonApp `locationName:"app" type:"structure"`
}

type MarathonApp struct {
	Id             *string       `locationName:"id" type:"string"`
	TasksStaged    *int64        `locationName:"tasksStaged" type:"integer"`
	TasksRunning   *int64        `locationName:"tasksRunning" type:"integer"`
	TasksHealthy   *int64        `locationName:"tasksHealthy" type:"integer"`
	TasksUnhealthy *int64        `locationName:"tasksUnhealthy" type:"integer"`
	Version        *string       `locationName:"version", type:"string"`
	Instances      *int64        `locationName:"instances" type:"integer"`
	Deployments    []*Deployment `locationName:"deployments" type:"list"`
}
