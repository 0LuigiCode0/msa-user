package hubHelper

func InitHelper(H HandlerForHelper) Helper { return &help{HandlerForHelper: H} }
