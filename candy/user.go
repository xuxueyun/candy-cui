package candy

import (
	"github.com/zeazen/candy-cui/meta"
	//"github.com/zeazen/candy-cui/util/log"
	"golang.org/x/net/context"
)

// Register 用户注册接口
func (c *CandyClient) Register(user, passwd string) (int64, error) {
	if code, err := CheckUserName(user); err != nil {
		return -1, NewError(code, err.Error())
	}
	if code, err := CheckUserPassword(passwd); err != nil {
		return -1, NewError(code, err.Error())
	}
	req := &meta.GateRegisterRequest{User: user, Password: passwd}
	var resp *meta.GateRegisterResponse
	var err error
	c.service(func(ctx context.Context, api meta.GateClient) error {
		if resp, err = api.Register(ctx, req); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return -1, err
	}
	//log.Debugf("resp:%+v", resp)
	return resp.ID, resp.Header.JsonError()
}
