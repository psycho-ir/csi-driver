package driver

const (
	PluginName    = "cloud.hetzner.csi.dobs"
	PluginVersion = "1.2.2"

	MaxVolumesPerNode = 16
	MinVolumeSize     = 10 // GB
	DefaultVolumeSize = MinVolumeSize

	TopologySegmentLocation = PluginName + "/location"
)
