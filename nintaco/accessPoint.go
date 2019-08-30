package nintaco

type accessPoint struct {
	listener        AccessPointListener
	accessPointType int
	minAddress      int
	maxAddress      int
	bank            int
}

func newAccessPoint(listener AccessPointListener, accessPointType, minAddress int) *accessPoint {
	return newAccessPoint3(listener, accessPointType, minAddress, -1, -1)
}

func newAccessPoint2(listener AccessPointListener, accessPointType, minAddress, maxAddress int) *accessPoint {
	return newAccessPoint3(listener, accessPointType, minAddress, maxAddress, -1)
}

func newAccessPoint3(listener AccessPointListener, accessPointType, minAddress, maxAddress, bank int) *accessPoint {
	a := &accessPoint{
		listener:        listener,
		accessPointType: accessPointType,
		bank:            bank,
	}
	if maxAddress < 0 {
		a.minAddress = minAddress
		a.maxAddress = minAddress
	} else if minAddress <= maxAddress {
		a.minAddress = minAddress
		a.maxAddress = maxAddress
	} else {
		a.minAddress = maxAddress
		a.maxAddress = minAddress
	}
	return a
}
