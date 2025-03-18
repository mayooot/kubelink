package model

type Cluster struct {
	tableName                struct{}          `sql:"cluster" pg:",discard_unknown_columns"`
	Id                       int               `sql:"id,pk"`
	ClusterName              string            `sql:"cluster_name"`
	ServerUrl                string            `sql:"server_url"`
	RemoteConnectionConfigId int               `sql:"remote_connection_config_id"`
	PrometheusEndpoint       string            `sql:"prometheus_endpoint"`
	Active                   bool              `sql:"active,notnull"`
	CdArgoSetup              bool              `sql:"cd_argo_setup,notnull"`
	Config                   map[string]string `sql:"config"`
	PUserName                string            `sql:"p_username"`
	PPassword                string            `sql:"p_password"`
	PTlsClientCert           string            `sql:"p_tls_client_cert"`
	PTlsClientKey            string            `sql:"p_tls_client_key"`
	AgentInstallationStage   int               `sql:"agent_installation_stage"`
	K8sVersion               string            `sql:"k8s_version"`
	ErrorInConnecting        string            `sql:"error_in_connecting"`
	IsVirtualCluster         bool              `sql:"is_virtual_cluster"`
	InsecureSkipTlsVerify    bool              `sql:"insecure_skip_tls_verify"`
	ProxyUrl                 string            `sql:"proxy_url"`
	ToConnectWithSSHTunnel   bool              `sql:"to_connect_with_ssh_tunnel"`
	SSHTunnelUser            string            `sql:"ssh_tunnel_user"`
	SSHTunnelPassword        string            `sql:"ssh_tunnel_password"`
	SSHTunnelAuthKey         string            `sql:"ssh_tunnel_auth_key"`
	SSHTunnelServerAddress   string            `sql:"ssh_tunnel_server_address"`
	RemoteConnectionConfig   *RemoteConnectionConfig
	AuditLog
}
