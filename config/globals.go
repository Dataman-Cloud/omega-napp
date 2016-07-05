package config

import "errors"

const RedisExpiration = 24 * 3600

const (
	HeaderToken   = "Authorization"
	CRON_ID       = "Cron_Id"
	STRATEGY_ID   = "Strategy-Id"
	Log_AlARM_ID  = "Log-Alarm-Id"
	SRY_SVC_TOKEN = "Sry-Svc-Token"
	CANALIAS      = "Can_Alias"
)
const (
	StatusDeploying = iota + 1
	StatusRunning
	StatusStopped
	StatusStopping
	StatusDeleting
	StatusScaling
	StatusStarting
	StatusCanceling
	StatusUnconnection
	StatusAbnormal
)
const (
	InstanceStarted = iota + 1
	InstanceDeploying
)
const (
	ExchangeCluster      = "cluster"
	ExchangeClusterInfo  = "cluster_info"
	ExchangeApplication  = "application"
	RoutingCluster       = "cluster"
	RoutingApplication   = "application"
	RoutingMarathonApp   = "marathon_apps"
	RoutingMarathonEvent = "marathon_event"
	RoutingHealth        = "health"
	LabelMarathon        = "app_shurenyun_com"
)

const (
	MethodGet    = "GET"
	MethodDelete = "DELETE"
	MethodPut    = "PUT"
	MethodPost   = "POST"
	MethodPatch  = "PATCH"
)

const (
	PathAppInfo      = "/v2/apps"
	PathDeploy       = "/v2/apps"
	PathUpdate       = "/v2/apps/%s"
	PathDeployment   = "/v2/deployments"
	PathMarathonApp  = "/v2/apps/%s"
	PathMarathonPing = "/ping"
)

const (
	PathHaWeight = "/api/weight"
)
const (
	PathAddJob         = "/scheduler/iso8601"             //add a job
	PathDependencyJob  = "/scheduler/dependency"          // add a dependency job
	PathListJobs       = "/scheduler/jobs"                //list all jobs
	PathDeleteJob      = "/scheduler/job/%s"              //delete a job
	PathDeleteJobTasks = "/scheduler/task/kill/%s"        //delete all tasks for a job
	PathStartJob       = "/scheduler/job/%s"              // lauch a job
	PathSearchJob      = "/scheduler/jobs/search?name=%s" // search job by jobName
	PathJobStatus      = "/scheduler/graph/csv"
)
const (
	KfApp     = "app:%d"
	KfCanary  = "canary:%d"
	KfVersion = "version:%d"
	KfNameCid = "%s%s"
)

const (
	KfAppName       = "name"
	KfAppClusterId  = "cid"
	KfAppStatus     = "status"
	KfAppInstances  = "instances"
	KfAppUpdateTime = "updateTime"
	KfAppVersionId  = "versionId"
	KfAppAliase     = "aliase"
)

const (
	KfVersionJson    = "json"
	KfVersionUpdated = "updated"
	CurrentDeploy    = 1
	UnCurrentDeploy  = 0
)

const (
	UpdateDeploying = 0
	UpdateOk        = 1
	UpdateAbnormal  = 2
	StatusCanary    = 3
)

const (
	// PortPrivate private service
	PortPrivate = 1
	// PortPublic public service
	PortPublic = 2
	// PortTCP tcp protocol
	PortTCP = 1
	// PortHTTP http protocol
	PortHTTP = 2
	// PortHasURI service bind by uri
	PortHasURI = 1
	// PortNoURI service without uri
	PortNoURI = 2
)

var (
	// ErrInvalidParameter invalid parameter
	ErrInvalidParameter = errors.New("invalid parameter")
)

const (
	BuildChronosJobsJson = `{"jobs":%s}`
)

const (
	// common
	NoError            int = 0
	SystemError        int = 10001
	InvalidToken       int = 10002
	DBOperationError   int = 10003 // 数据库操作错误
	NoPrivilege        int = 10010
	InvalidPerPagePara int = 10011
	InvalidPagePara    int = 10012
	InvalidSortByPara  int = 10013
	InvalidOrderByPara int = 10014
	JsonFormatError    int = 10008

	// application

	AppNameConflict                 = 14003
	AppPortConflict                 = 14004
	AppVersionConflict              = 14005
	AppIsLocked                     = 14006
	AppScaleFailed                  = 14007
	RollbackFailedUpdateSuccess     = 14008
	InvalidEnv                      = 14009
	IllegalClusterId                = 14010
	InvalidRequest              int = 14011
	IllegalAppId                    = 14012
	MarathonDisabled                = 14013
	CheckPermissionError            = 14014
	AddChronosJobError              = 14014
	CheckWeightFailed               = 14015
	NoneMasterError                 = 14016
	ScaleCompletely                 = 14017
	CanaryExistsError               = 14018

	AppNotFound = 14404
	// application compose
	RepoConfigFileNotFound  = 16000
	StackDeploymentError    = 16001
	StackNameConflict       = 16002
	StackDeletionNotAllowed = 16003
	StackNotInTheRightState = 16007

	// build
	JsonOperationError       = 15002 // json 操作错误
	GetHeaderError           = 15003 // 读取header错误
	GetPathError             = 15004 // 获取路径参数错误
	GetBodyError             = 15005 // 获取body中的参数错误
	TypeTransactionError     = 15006 // 类型转换错误
	BuildEntryError          = 15007 // 构建Entry错误
	PidNotExistError         = 15008 // project不存在
	ImgIdNotExistError       = 15009 // image不存在
	BuildStreamNotExistError = 15010 // image不存在
)

// in align with error reported from compose component
var ComposeErrorMsgCodeMap = map[string]int{
	"ErrStackDeployError":                            16001,
	"ErrCatalogSyntaxError":                          16004,
	"ErrDockerComposeSyntaxError":                    16005,
	"ErrMarathonConfigSyntaxError":                   16006,
	"ErrMarathonConfigAndDockerComposeNotMatchError": 16007,
}

const (
	DEFAULT_PER_PAGE_SIZE string = "2000"
	DEFAULT_PAGE          string = "1"
)

// path
const (
	PermissionPath   = "/api/v3/permissions"
	CheckClusterPath = "/api/v3/clusters"
)

// cluster read and write permission
const (
	CLUSTER_WRITE_PERMISSION = "cluster_write"
	CLUSTER_READ_PERMISSION  = "cluster_read"
	ROLE_ADMIN               = 1
	ROLE_USER                = 2
	ROLE_ILLEGAL_USER        = -1
)

const (
	CANCELANDSTOP = "cancelAndStop"
	CANCEL        = "cancel"
)

const (
	LAST_PULL_FAILURE = iota + 1
	LAST_START_FAILURE
)

const (
	SCALE_FOR_ALERT = 0
	SCALE_FOR_LOG   = 1
)
