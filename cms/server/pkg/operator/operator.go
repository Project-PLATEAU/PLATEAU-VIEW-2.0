package operator

type Operator struct {
	user        *UserID
	integration *IntegrationID
	isMachine   bool
}

func OperatorFromUser(user UserID) Operator {
	return Operator{
		user: &user,
	}
}

func OperatorFromIntegration(integration IntegrationID) Operator {
	return Operator{
		integration: &integration,
	}
}

func OperatorFromMachine() Operator {
	return Operator{
		isMachine: true,
	}
}

func (o Operator) User() *UserID {
	return o.user.CloneRef()
}

func (o Operator) Integration() *IntegrationID {
	return o.integration.CloneRef()
}

func (o Operator) Machine() bool {
	return o.isMachine
}

func (o Operator) Validate() bool {
	return !o.user.IsNil() || !o.integration.IsNil() || o.Machine()
}
