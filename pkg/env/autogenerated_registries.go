package env

var (
	// AutogeneratedRegistriesDisabled being set to true ignores the autogenerated registries sent by Sensors
	AutogeneratedRegistriesDisabled = RegisterBooleanSetting("ROX_DISABLE_AUTOGENERATED_REGISTRIES", false)

	// AutogenerateGlobalPullSecRegistries when false disables creating autogenerated registries from the OCP global pull secret.
	// If autogenerated registries are disabled (via AutogeneratedRegistriesDisabled) - this setting has no effect.
	AutogenerateGlobalPullSecRegistries = RegisterBooleanSetting("ROX_AUTOGENERATE_GLOBAL_PULLSEC_REGISTRIES", true)
)
