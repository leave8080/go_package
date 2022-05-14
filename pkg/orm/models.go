package orm

import (
	"time"
)

type MxAccount struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	TeamId    int       `xorm:"not null default 0 comment('所属子公司id') index unique(uk_team_id_is_master) INT(10)"`
	RoleId    int       `xorm:"not null default 0 comment('所属角色id') INT(10)"`
	Openid    string    `xorm:"not null default '' comment('账户唯一openid') unique VARCHAR(32)"`
	IsMaster  int       `xorm:"not null default 0 comment('是否子公司主管理员，0-否，1-是') unique(uk_team_id_is_master) TINYINT(4)"`
	Username  string    `xorm:"not null default '' comment('登录账户') unique VARCHAR(32)"`
	Password  string    `xorm:"not null default '' comment('Hash密码') VARCHAR(64)"`
	Realname  string    `xorm:"not null default '' comment('真实姓名') VARCHAR(8)"`
	Phone     string    `xorm:"not null default '' comment('手机号码') index VARCHAR(16)"`
	Email     string    `xorm:"not null default '' comment('邮箱地址') VARCHAR(64)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxAccount) TableName() string {
	return "mx_account"
}

type MxAdminUser struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	Username  string    `xorm:"not null default '' comment('登录账户') unique VARCHAR(32)"`
	Password  string    `xorm:"not null default '' comment('Hash密码') VARCHAR(64)"`
	Salt      string    `xorm:"not null default '' comment('md5盐值') VARCHAR(32)"`
	Realname  string    `xorm:"not null default '' comment('真实姓名') VARCHAR(8)"`
	Phone     string    `xorm:"not null default '' comment('手机号码') VARCHAR(16)"`
	IsMaster  int       `xorm:"not null default 0 comment('主账户标示') TINYINT(4)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxAdminUser) TableName() string {
	return "mx_admin_user"
}

type MxApp struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	ProjectId int       `xorm:"not null default 0 comment('所属项目id') index INT(10)"`
	Name      string    `xorm:"not null default '' comment('App名') unique VARCHAR(16)"`
	Appid     string    `xorm:"not null default '' comment('App唯一标识') unique VARCHAR(32)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxApp) TableName() string {
	return "mx_app"
}

type MxAppExtra struct {
	Id          int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	AppId       int       `xorm:"not null default 0 comment('所属app的id') index INT(10)"`
	Platform    int       `xorm:"not null default 0 comment('设备平台：0-未知，1-Android，2-iOS') TINYINT(4)"`
	Package     string    `xorm:"not null default '' comment('包名') VARCHAR(128)"`
	OauthEnable int       `xorm:"not null default 1 comment('第三方登录开关，0-关闭，1-开启') TINYINT(4)"`
	IsDeleted   int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxAppExtra) TableName() string {
	return "mx_app_extra"
}

type MxAppPanel struct {
	Id          int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	AppId       int       `xorm:"not null default 0 comment('所属app的id') index INT(10)"`
	AdminUserId int       `xorm:"not null default 0 comment('添加admin账户的id') INT(10)"`
	Version     string    `xorm:"not null default '' comment('app版本号') VARCHAR(32)"`
	ReleaseType int       `xorm:"not null default 0 comment('升级方式：0-未发布，1-灰度发布，2-正式发布') TINYINT(4)"`
	IosUrl      string    `xorm:"not null default '' comment('ios 面板下载路径') VARCHAR(128)"`
	AndroidUrl  string    `xorm:"not null default '' comment('android 面板下载路径') VARCHAR(128)"`
	Remark      string    `xorm:"not null default '' comment('备注信息') VARCHAR(1024)"`
	IsDeleted   int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index TIMESTAMP"`
	Mtime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxAppPanel) TableName() string {
	return "mx_app_panel"
}

type MxAppVersion struct {
	Id          int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	AppId       int       `xorm:"not null default 0 comment('所属app的id') index INT(10)"`
	AdminUserId int       `xorm:"not null default 0 comment('添加账户的id') INT(10)"`
	Version     string    `xorm:"not null default '' comment('app版本号') VARCHAR(32)"`
	VersionCode int       `xorm:"not null default 0 comment('app版本code') INT(11)"`
	UpgradeType int       `xorm:"not null default 0 comment('升级方式：0-提醒升级，1-强制升级') TINYINT(4)"`
	Platform    int       `xorm:"not null default 0 comment('设备平台：0-未知，1-Android，2-iOS') TINYINT(4)"`
	Url         string    `xorm:"not null default '' comment('apk下载路径') VARCHAR(128)"`
	Remark      string    `xorm:"not null default '' comment('版本描述') VARCHAR(1024)"`
	IsDeleted   int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index TIMESTAMP"`
	Mtime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxAppVersion) TableName() string {
	return "mx_app_version"
}

type MxBinding struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	Type      int       `xorm:"not null default 0 comment('绑定关系种类，后期自定义：1-角色权限绑定关系(pid:角色,cid:权限)') index TINYINT(4)"`
	Pid       int       `xorm:"not null default 0 comment('父id') index INT(11)"`
	Cid       int       `xorm:"not null default 0 comment('子id') index INT(11)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxBinding) TableName() string {
	return "mx_binding"
}

type MxCity struct {
	Id          int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	PCode       string    `xorm:"not null default '' comment('所属省code') index VARCHAR(4)"`
	LocationKey string    `xorm:"not null default '' comment('Accu的LocationKey') unique VARCHAR(16)"`
	ZhName      string    `xorm:"not null default '' comment('城市中文名称') index VARCHAR(16)"`
	EnName      string    `xorm:"not null default '' comment('城市英文名称') VARCHAR(32)"`
	IsDeleted   int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxCity) TableName() string {
	return "mx_city"
}

type MxCompany struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	Name      string    `xorm:"not null default '' comment('公司名称') VARCHAR(32)"`
	Phone     string    `xorm:"not null default '' comment('电话') index VARCHAR(16)"`
	Address   string    `xorm:"not null default '' comment('地址') VARCHAR(128)"`
	Email     string    `xorm:"not null default '' comment('邮箱') VARCHAR(64)"`
	Logo      string    `xorm:"not null default '' comment('logo地址') VARCHAR(256)"`
	Remark    string    `xorm:"not null default '' comment('备注信息') VARCHAR(512)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxCompany) TableName() string {
	return "mx_company"
}

type MxCountry struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	Code      string    `xorm:"not null default '' comment('国家简写编号') unique VARCHAR(4)"`
	ZhName    string    `xorm:"not null default '' comment('国家中文名称') index VARCHAR(16)"`
	EnName    string    `xorm:"not null default '' comment('国家英文名称') VARCHAR(16)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxCountry) TableName() string {
	return "mx_country"
}

type MxDevice struct {
	Id           int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId       int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	HomeId       int       `xorm:"not null default 0 comment('所属家庭id') index INT(10)"`
	ProductId    int       `xorm:"not null default 0 comment('所属产品id') index INT(10)"`
	Name         string    `xorm:"not null default '' comment('云端设备名称') VARCHAR(32)"`
	Iotid        string    `xorm:"not null default '' comment('设备iotid') unique VARCHAR(32)"`
	Uuid         string    `xorm:"not null default '' comment('设备UUID') VARCHAR(36)"`
	MeshAddress  int       `xorm:"not null default 0 comment('设备mesh_address，设备端') INT(11)"`
	Hardware     string    `xorm:"not null default '' comment('硬件版本') VARCHAR(32)"`
	Software     string    `xorm:"not null default '' comment('固件版本') VARCHAR(32)"`
	DeviceName   string    `xorm:"not null default '' comment('飞燕设备名称') VARCHAR(32)"`
	DeviceSecret string    `xorm:"not null default '' comment('飞燕设备Secret') VARCHAR(32)"`
	Lon          string    `xorm:"not null default '' comment('经度，小数点后最多6位，例：116.481488') VARCHAR(16)"`
	Lat          string    `xorm:"not null default '' comment('纬度，小数点后最多6位，例：39.990466') VARCHAR(16)"`
	Country      string    `xorm:"not null default '' comment('国家') VARCHAR(8)"`
	Province     string    `xorm:"not null default '' comment('省级') VARCHAR(16)"`
	City         string    `xorm:"not null default '' comment('城市') VARCHAR(16)"`
	BindStatus   int       `xorm:"not null default 0 comment('绑定状态：0-未绑定，1-绑定') TINYINT(4)"`
	IsDeleted    int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime        time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index TIMESTAMP"`
	Mtime        time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxDevice) TableName() string {
	return "mx_device"
}

type MxDeviceAction struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	DeviceId  int       `xorm:"not null default 0 comment('所属设备id') index INT(10)"`
	Rule      string    `xorm:"not null comment('设备动作规则') JSON"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxDeviceAction) TableName() string {
	return "mx_device_action"
}

type MxDeviceOnline struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	DeviceId  int       `xorm:"not null default 0 comment('所属设备id') unique INT(10)"`
	Online    int       `xorm:"not null default 0 comment('在线状态：0-离线，1-在线') index TINYINT(4)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxDeviceOnline) TableName() string {
	return "mx_device_online"
}

type MxDeviceProperty struct {
	Id         int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId     int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	DeviceId   int       `xorm:"not null default 0 comment('所属设备id') index INT(10)"`
	Identifier string    `xorm:"not null default '' comment('属性标识key') index VARCHAR(32)"`
	Value      string    `xorm:"not null default '' comment('属性标识下的value') index VARCHAR(16)"`
	IsDeleted  int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxDeviceProperty) TableName() string {
	return "mx_device_property"
}

type MxDeviceRule struct {
	Id         int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId     int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	DeviceId   int       `xorm:"not null default 0 comment('所属设备id') index INT(10)"`
	SensorRule string    `xorm:"not null comment('传感器联动规则') JSON"`
	SwitchRule string    `xorm:"not null comment('单火线开关规则，数据格式如：1,2,3') JSON"`
	IsDeleted  int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxDeviceRule) TableName() string {
	return "mx_device_rule"
}

type MxDeviceStat struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	AppId     int       `xorm:"not null default 0 comment('所属app的id') index INT(10)"`
	DeviceId  int       `xorm:"not null default 0 comment('所属设备id') index unique(uk_device_id_date) INT(10)"`
	Date      time.Time `xorm:"not null comment('日期') index unique(uk_device_id_date) DATE"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxDeviceStat) TableName() string {
	return "mx_device_stat"
}

type MxDoc struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	TeamId    int       `xorm:"not null default 0 comment('所属团队id') index INT(10)"`
	Type      int       `xorm:"not null default 0 comment('文档类别：0-未知，1-隐私协议，2-服务协议，3-公司介绍，4-常见问题，5-推送内容') index TINYINT(4)"`
	Title     string    `xorm:"not null default '' comment('文档标题') VARCHAR(64)"`
	Content   string    `xorm:"not null comment('富文本内容') TEXT"`
	Path      string    `xorm:"not null default '' comment('oss路径') VARCHAR(128)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxDoc) TableName() string {
	return "mx_doc"
}

type MxEno struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	TeamId    int       `xorm:"not null default 0 comment('所属团队id') index unique(uk_team_code) INT(10)"`
	Code      int       `xorm:"not null default 0 comment('错误Code') unique(uk_team_code) INT(11)"`
	Message   string    `xorm:"not null default '' comment('错误Message') VARCHAR(64)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxEno) TableName() string {
	return "mx_eno"
}

type MxErrLog struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	Platform  int       `xorm:"not null default 0 comment('设备平台：0-未知，1-Android，2-iOS') index TINYINT(4)"`
	Version   string    `xorm:"not null default '' comment('软件版本') VARCHAR(32)"`
	Date      time.Time `xorm:"not null default '1970-01-01' comment('错误日期') index DATE"`
	Content   string    `xorm:"not null default '' comment('错误内容') VARCHAR(2048)"`
	Url       string    `xorm:"not null default '' comment('错误日志文件url') VARCHAR(256)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxErrLog) TableName() string {
	return "mx_err_log"
}

type MxFaq struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	TeamId    int       `xorm:"not null default 0 comment('所属团队id') index INT(10)"`
	DocId     int       `xorm:"not null default 0 comment('所属文档id') INT(10)"`
	ProductId int       `xorm:"not null default 0 comment('所属产品id') INT(10)"`
	Type      int       `xorm:"not null default 0 comment('问题类别：0-其他问题，1-设备问题，2-应用问题，3-网关问题') index TINYINT(4)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxFaq) TableName() string {
	return "mx_faq"
}

type MxFeedback struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	ProductId int       `xorm:"not null default 0 comment('所属产品id，当type为1、2时，此字段不可为空') index INT(10)"`
	Phone     string    `xorm:"not null default '' comment('用户手机号') VARCHAR(16)"`
	Content   string    `xorm:"not null default '' comment('反馈内容') VARCHAR(1024)"`
	Type      int       `xorm:"not null default 0 comment('反馈类型：0-其他问题，1-设备问题，2-配网问题，3-反馈故障，4-功能建议') index TINYINT(4)"`
	Status    int       `xorm:"not null default 0 comment('反馈状态：0-未解决，1-已解决') index TINYINT(4)"`
	Platform  int       `xorm:"not null default 0 comment('设备平台：0-未知，1-Android，2-iOS') index TINYINT(4)"`
	Version   string    `xorm:"not null default '' comment('软件版本') VARCHAR(32)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxFeedback) TableName() string {
	return "mx_feedback"
}

type MxFeedbackMsg struct {
	Id         int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId     int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	FeedbackId int       `xorm:"not null default 0 comment('所属反馈id') index INT(10)"`
	AccountId  int       `xorm:"not null default 0 comment('所属Web账号id，查询名字用') INT(10)"`
	Content    string    `xorm:"not null default '' comment('聊天内容') VARCHAR(1024)"`
	UserType   int       `xorm:"not null default 0 comment('消息所属方：0-未知，1-APP用户，2-Web用户') TINYINT(4)"`
	AppIsRead  int       `xorm:"not null default 0 comment('APP已读状态：0-未读，1-已读') TINYINT(4)"`
	WebIsRead  int       `xorm:"not null default 0 comment('Web已读状态：0-未读，1-已读') TINYINT(4)"`
	IsDeleted  int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxFeedbackMsg) TableName() string {
	return "mx_feedback_msg"
}

type MxHome struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	Homeid    string    `xorm:"not null default '' comment('家庭的id，家的唯一标识，飞燕返回字段') unique VARCHAR(32)"`
	Name      string    `xorm:"not null default '' comment('家庭名称') VARCHAR(32)"`
	Lon       string    `xorm:"not null default '' comment('经度，小数点后最多6位，例：116.481488') VARCHAR(16)"`
	Lat       string    `xorm:"not null default '' comment('纬度，小数点后最多6位，例：39.990466') VARCHAR(16)"`
	Country   string    `xorm:"not null default '' comment('国家') VARCHAR(8)"`
	Province  string    `xorm:"not null default '' comment('省级') index VARCHAR(16)"`
	City      string    `xorm:"not null default '' comment('城市') index VARCHAR(16)"`
	Address   string    `xorm:"not null default '' comment('详细地址') VARCHAR(128)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxHome) TableName() string {
	return "mx_home"
}

type MxHomeRule struct {
	Id         int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId     int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	HomeId     int       `xorm:"not null default 0 comment('所属家庭id') index INT(10)"`
	MeshRule   string    `xorm:"not null comment('家庭Mesh配置信息') JSON"`
	SensorRule string    `xorm:"not null comment('家庭下传感器联动规则') JSON"`
	IsDeleted  int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxHomeRule) TableName() string {
	return "mx_home_rule"
}

type MxLoginLog struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	LastIp    string    `xorm:"not null default '0.0.0.0' comment('最后登录IP地址') VARCHAR(16)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') index TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxLoginLog) TableName() string {
	return "mx_login_log"
}

type MxMenu struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	Level     int       `xorm:"not null default 1 comment('菜单分级：1、2、3 . . .') TINYINT(4)"`
	Pid       int       `xorm:"not null default 0 comment('父菜单id') INT(11)"`
	Name      string    `xorm:"not null default '' comment('菜单名称') VARCHAR(16)"`
	Code      string    `xorm:"not null default '' comment('菜单编号') index VARCHAR(32)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxMenu) TableName() string {
	return "mx_menu"
}

type MxProduct struct {
	Id            int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	ProjectId     int       `xorm:"not null default 0 comment('所属项目id') index INT(10)"`
	Name          string    `xorm:"not null default '' comment('产品名称') VARCHAR(32)"`
	ProductKey    string    `xorm:"not null default '' comment('飞燕产品Key') unique VARCHAR(16)"`
	ProductSecret string    `xorm:"not null default '' comment('飞燕产品Secret') VARCHAR(16)"`
	ProductType   int       `xorm:"not null default 0 comment('云端自定义产品类型：0-未知产品类型 1-WiFi设备 2-BLE设备 3-BLE MESH设备 4-BLE MESH低功耗设备 5-网关设备') index TINYINT(4)"`
	SendCount     int       `xorm:"not null default 1 comment('报文发送次数') INT(11)"`
	DataFormat    string    `xorm:"not null default '' comment('数据格式') VARCHAR(16)"`
	NetType       string    `xorm:"not null default '' comment('入网类型') VARCHAR(16)"`
	NodeType      string    `xorm:"not null default '' comment('节点类型') VARCHAR(16)"`
	Region        string    `xorm:"not null default '' comment('地域') VARCHAR(16)"`
	Categoryid    int       `xorm:"not null default 0 comment('飞燕归属品类id') INT(11)"`
	RbacTenantId  string    `xorm:"not null default '' comment('飞燕租户id') VARCHAR(32)"`
	Status        string    `xorm:"not null default '' comment('产品状态') VARCHAR(32)"`
	SyncStatus    int       `xorm:"not null default 0 comment('数据同步状态：0-不同步，1-同步') TINYINT(4)"`
	GmtCreate     time.Time `xorm:"not null default '1970-01-01 00:00:00' comment('设备创建时间') DATETIME"`
	GmtModified   time.Time `xorm:"not null default '1970-01-01 00:00:00' comment('设备修改时间') DATETIME"`
	IsDeleted     int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime         time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime         time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxProduct) TableName() string {
	return "mx_product"
}

type MxProductCategory struct {
	Id          int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	CategoryId  int       `xorm:"not null default 0 comment('品类id') INT(11)"`
	CategoryPid int       `xorm:"not null default 0 comment('品类父id') INT(11)"`
	Name        string    `xorm:"not null default '' comment('品类名称') VARCHAR(32)"`
	IsDeleted   int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxProductCategory) TableName() string {
	return "mx_product_category"
}

type MxProject struct {
	Id            int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	TeamId        int       `xorm:"not null default 0 comment('所属团队id') INT(10)"`
	CloudPlatform int       `xorm:"not null default 0 comment('设备平台：0-未知, 1-飞燕') TINYINT(4)"`
	Name          string    `xorm:"not null default '' comment('项目名称') VARCHAR(32)"`
	Projectid     string    `xorm:"not null default '' comment('项目ID唯一标识') unique VARCHAR(32)"`
	ProdKey       string    `xorm:"not null default '' comment('线上环境AppKey') VARCHAR(16)"`
	ProdSecret    string    `xorm:"not null default '' comment('线上环境AppSecret') VARCHAR(32)"`
	TestKey       string    `xorm:"not null default '' comment('测试环境AppKey') VARCHAR(16)"`
	TestSecret    string    `xorm:"not null default '' comment('测试环境AppSecret') VARCHAR(32)"`
	SyncStatus    int       `xorm:"not null default 0 comment('数据同步状态：0-关闭，1-打开') TINYINT(4)"`
	SyncKey       string    `xorm:"not null default '' comment('数据同步AppKey') VARCHAR(16)"`
	SyncSecret    string    `xorm:"not null default '' comment('数据同步AppSecret') VARCHAR(32)"`
	IsDeleted     int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime         time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime         time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxProject) TableName() string {
	return "mx_project"
}

type MxProvince struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	PCode     string    `xorm:"not null default '' comment('所属国家code') index VARCHAR(4)"`
	Code      string    `xorm:"not null default '' comment('省简写编号') unique VARCHAR(4)"`
	ZhName    string    `xorm:"not null default '' comment('省中文名称') index VARCHAR(16)"`
	EnName    string    `xorm:"not null default '' comment('省英文名称') VARCHAR(32)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxProvince) TableName() string {
	return "mx_province"
}

type MxPush struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	AppId     int       `xorm:"not null default 0 comment('所属App的id') index INT(10)"`
	AccountId int       `xorm:"not null default 0 comment('所属Web账号id，查询名字用') INT(10)"`
	DocId     int       `xorm:"not null default 0 comment('所属文档id') INT(10)"`
	Type      int       `xorm:"not null default 0 comment('推送类型：0-未知，1-推送所有，2-指定产品') index TINYINT(4)"`
	Summary   string    `xorm:"not null default '' comment('推送概览（摘要）') VARCHAR(256)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxPush) TableName() string {
	return "mx_push"
}

type MxPushRecord struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	PushId    int       `xorm:"not null default 0 comment('所属团队id') index INT(10)"`
	Status    int       `xorm:"not null default 0 comment('推送已读状态：0-未读，1-已读') index TINYINT(4)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxPushRecord) TableName() string {
	return "mx_push_record"
}

type MxRole struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	TeamId    int       `xorm:"not null default 0 comment('所属子公司id') INT(10)"`
	Name      string    `xorm:"not null default '' comment('角色名称') VARCHAR(16)"`
	Desc      string    `xorm:"not null default '' comment('角色描述') VARCHAR(64)"`
	IsSys     int       `xorm:"not null default 0 comment('0-非系统，1-系统') TINYINT(4)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxRole) TableName() string {
	return "mx_role"
}

type MxTeam struct {
	Id          int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	CompanyId   int       `xorm:"not null default 0 comment('所属公司id') index INT(10)"`
	AdminUserId int       `xorm:"not null default 0 comment('总后台用户id') INT(10)"`
	Name        string    `xorm:"not null default '' comment('团队名称') VARCHAR(32)"`
	Phone       string    `xorm:"not null default '' comment('电话') index VARCHAR(16)"`
	Address     string    `xorm:"not null default '' comment('地址') VARCHAR(128)"`
	Email       string    `xorm:"not null default '' comment('邮箱') VARCHAR(64)"`
	Logo        string    `xorm:"not null default '' comment('logo地址') VARCHAR(256)"`
	Remark      string    `xorm:"not null default '' comment('备注信息') VARCHAR(512)"`
	ExpireAt    time.Time `xorm:"not null default '1970-01-01' comment('会员过期日期') DATE"`
	IsDeleted   int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime       time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxTeam) TableName() string {
	return "mx_team"
}

type MxToken struct {
	Id           int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	AccountId    int       `xorm:"not null default 0 comment('所属Web账号id') INT(10)"`
	AccessToken  string    `xorm:"not null default '' comment('AccessToken') index VARCHAR(64)"`
	RefreshToken string    `xorm:"not null default '' comment('RefreshToken') index VARCHAR(64)"`
	ExpiresAt    time.Time `xorm:"not null default '1970-01-01 00:00:00' comment('到期时间') DATETIME"`
	IsDeleted    int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime        time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime        time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxToken) TableName() string {
	return "mx_token"
}

type MxUser struct {
	Id         int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	AppId      int       `xorm:"not null default 0 comment('所属App的id') index unique(uk_openid) unique(uk_username) INT(10)"`
	Openid     string    `xorm:"not null default '' comment('用户唯一openid') unique(uk_openid) VARCHAR(32)"`
	Identityid string    `xorm:"not null default '' comment('飞燕平台用户身份id') VARCHAR(64)"`
	Username   string    `xorm:"not null default '' comment('登录账户') index unique(uk_username) VARCHAR(32)"`
	Password   string    `xorm:"not null default '' comment('Hash密码') VARCHAR(64)"`
	Phone      string    `xorm:"not null default '' comment('手机号码') index VARCHAR(16)"`
	Email      string    `xorm:"not null default '' comment('邮箱地址') VARCHAR(64)"`
	IsDeleted  int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxUser) TableName() string {
	return "mx_user"
}

type MxUserClient struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	Platform  int       `xorm:"not null default 0 comment('设备平台：0-未知，1-Android，2-iOS') index TINYINT(4)"`
	Clientid  string    `xorm:"not null default '' comment('推送SDK生成的唯一clientid') VARCHAR(32)"`
	Switch    int       `xorm:"not null default 1 comment('client推送开关：1-开，0-关') index TINYINT(4)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxUserClient) TableName() string {
	return "mx_user_client"
}

type MxUserInfo struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	Nickname  string    `xorm:"not null default '' comment('昵称') VARCHAR(16)"`
	Avatar    string    `xorm:"not null default '' comment('头像地址') VARCHAR(256)"`
	Gender    int       `xorm:"not null default 0 comment('性别：0-未知，1-男，2-女') TINYINT(4)"`
	Age       int       `xorm:"not null default 0 comment('年龄') TINYINT(4)"`
	Country   string    `xorm:"not null default '' comment('国家') VARCHAR(8)"`
	Province  string    `xorm:"not null default '' comment('省级') VARCHAR(16)"`
	City      string    `xorm:"not null default '' comment('城市') VARCHAR(16)"`
	Address   string    `xorm:"not null default '' comment('详细地址') VARCHAR(128)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxUserInfo) TableName() string {
	return "mx_user_info"
}

type MxUserStat struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index unique(uk_user_id_date) INT(10)"`
	AppId     int       `xorm:"not null default 0 comment('所属app的id') index INT(10)"`
	Date      time.Time `xorm:"not null comment('日期') index unique(uk_user_id_date) DATE"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxUserStat) TableName() string {
	return "mx_user_stat"
}

type MxUserThird struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	Platform  int       `xorm:"not null default 0 comment('第三方登录平台：1-微信，2-QQ') index TINYINT(4)"`
	Openid    string    `xorm:"not null default '' comment('用户三方登录唯一openid') index VARCHAR(32)"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxUserThird) TableName() string {
	return "mx_user_third"
}

type MxVirtual struct {
	Id         int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId     int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	HomeId     int       `xorm:"not null default 0 comment('所属家庭id') index unique(uk_home_id_vid) INT(10)"`
	Name       string    `xorm:"not null default '' comment('虚拟按钮名称') VARCHAR(32)"`
	Vid        string    `xorm:"not null default '' comment('虚拟按钮唯一标识，类似iotid') index unique(uk_home_id_vid) VARCHAR(4)"`
	Icon       string    `xorm:"not null default '' comment('虚拟按钮图标') VARCHAR(256)"`
	LabelIcon  string    `xorm:"not null default '' comment('标签icon图标') VARCHAR(256)"`
	LabelColor string    `xorm:"not null default '' comment('标签颜色，例：#000000') VARCHAR(8)"`
	IsDeleted  int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime      time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') index TIMESTAMP"`
}

func (m *MxVirtual) TableName() string {
	return "mx_virtual"
}

type MxVirtualAction struct {
	Id        int       `xorm:"not null pk autoincr comment('自增长ID') INT(10)"`
	UserId    int       `xorm:"not null default 0 comment('所属用户id') index INT(10)"`
	VirtualId int       `xorm:"not null default 0 comment('所属虚拟按钮id') index INT(10)"`
	DeviceId  int       `xorm:"not null default 0 comment('所属设备id') index INT(10)"`
	Rule      string    `xorm:"not null comment('虚拟按钮动作规则') JSON"`
	IsDeleted int       `xorm:"not null default 0 comment('0-正常，1-删除') TINYINT(4)"`
	Ctime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('创建时间') TIMESTAMP"`
	Mtime     time.Time `xorm:"not null default CURRENT_TIMESTAMP comment('最后修改时间') TIMESTAMP"`
}

func (m *MxVirtualAction) TableName() string {
	return "mx_virtual_action"
}
