package entity

type NightingaleMessage struct {
	Id              int    `json:"id"`
	Cate            string `json:"cate"`
	Cluster         string `json:"cluster"`
	DatasourceId    int    `json:"datasource_id"`
	GroupId         int    `json:"group_id"`
	GroupName       string `json:"group_name"`
	Hash            string `json:"hash"`
	RuleId          int    `json:"rule_id"`
	RuleName        string `json:"rule_name"`
	RuleNote        string `json:"rule_note"`
	RuleProd        string `json:"rule_prod"`
	RuleAlgo        string `json:"rule_algo"`
	Severity        int    `json:"severity"`
	PromForDuration int    `json:"prom_for_duration"`
	PromQl          string `json:"prom_ql"`
	RuleConfig      struct {
		Inhibit bool `json:"inhibit"`
		Queries []struct {
			Keys struct {
				LabelKey string `json:"labelKey"`
				ValueKey string `json:"valueKey"`
			} `json:"keys"`
			PromQl   string `json:"prom_ql"`
			Severity int    `json:"severity"`
		} `json:"queries"`
	} `json:"rule_config"`
	PromEvalInterval int           `json:"prom_eval_interval"`
	Callbacks        []interface{} `json:"callbacks"`
	RunbookUrl       string        `json:"runbook_url"`
	NotifyRecovered  int           `json:"notify_recovered"`
	NotifyChannels   []string      `json:"notify_channels"`
	NotifyGroups     []string      `json:"notify_groups"`
	NotifyGroupsObj  []struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Note     string `json:"note"`
		CreateAt int    `json:"create_at"`
		CreateBy string `json:"create_by"`
		UpdateAt int    `json:"update_at"`
		UpdateBy string `json:"update_by"`
	} `json:"notify_groups_obj"`
	TargetIdent  string   `json:"target_ident"`
	TargetNote   string   `json:"target_note"`
	TriggerTime  int      `json:"trigger_time"`
	TriggerValue string   `json:"trigger_value"`
	Tags         []string `json:"tags"`
	TagsMap      struct {
		Name     string `json:"__name__"`
		Ident    string `json:"ident"`
		Rulename string `json:"rulename"`
		Source   string `json:"source"`
	} `json:"tags_map"`
	Annotations struct {
	} `json:"annotations"`
	IsRecovered    bool `json:"is_recovered"`
	NotifyUsersObj []struct {
		Id       int      `json:"id"`
		Username string   `json:"username"`
		Nickname string   `json:"nickname"`
		Phone    string   `json:"phone"`
		Email    string   `json:"email"`
		Portrait string   `json:"portrait"`
		Roles    []string `json:"roles"`
		Contacts struct {
			WecomRobotToken string `json:"wecom_robot_token,omitempty"`
		} `json:"contacts"`
		Maintainer int    `json:"maintainer"`
		CreateAt   int    `json:"create_at"`
		CreateBy   string `json:"create_by"`
		UpdateAt   int    `json:"update_at"`
		UpdateBy   string `json:"update_by"`
		Admin      bool   `json:"admin"`
	} `json:"notify_users_obj"`
	LastEvalTime             int         `json:"last_eval_time"`
	LastEscalationNotifyTime int         `json:"last_escalation_notify_time"`
	LastSentTime             int         `json:"last_sent_time"`
	NotifyCurNumber          int         `json:"notify_cur_number"`
	FirstTriggerTime         int         `json:"first_trigger_time"`
	ExtraConfig              interface{} `json:"extra_config"`
	Status                   int         `json:"status"`
	Claimant                 string      `json:"claimant"`
	SubRuleId                int         `json:"sub_rule_id"`
}
